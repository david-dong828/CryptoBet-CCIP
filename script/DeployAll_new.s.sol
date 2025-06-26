// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/MarketRegistry.sol";
import "../src/PredictionMarketExtended.sol";
import "../src/FunctionsConsumer.sol";
import "../src/CCIPBetBridge.sol";

contract DeployAllScript_new is Script {
    // =============== NETWORK CONFIGURATIONS ===============
    
    // Sepolia Configuration
    struct SepoliaConfig {
        address chainlinkRouter;
        address ccipRouter;
        address linkToken;
        uint64 chainSelector;
        address ccipBnM;
    }
    
    // Fuji Configuration  
    struct FujiConfig {
        address ccipRouter;
        address linkToken;
        uint64 chainSelector;
        address ccipBnM;
    }
    
    // Network configuration functions
    function getSepoliaConfig() internal pure returns (SepoliaConfig memory) {
        return SepoliaConfig({
            chainlinkRouter: 0xb83E47C2bC239B3bf370bc41e1459A34b41238D0,
            ccipRouter: 0x0BF3dE8c5D3e8A2B34D2BEeB17ABfCeBaf363A59,
            linkToken: 0x779877A7B0D9E8603169DdbD7836e478b4624789,
            chainSelector: 16015286601757825753,
            ccipBnM: 0xFd57b4ddBf88a4e07fF4e34C487b99af2Fe82a05 // Sepolia CCIP-BnM
        });
    }
    
    function getFujiConfig() internal pure returns (FujiConfig memory) {
        return FujiConfig({
            ccipRouter: 0xF694E193200268f9a4868e4Aa017A0118C9a8177,
            linkToken: 0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846,
            chainSelector: 14767482510784806043,
            ccipBnM: 0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4 // Fuji CCIP-BnM
        });
    }
    
    // API Configuration
    string constant API_BASE_URL = "https://upright-mongrel-virtually.ngrok-free.app/odds";
    
    // Deployment configuration
    struct MarketConfig {
        string question;
        string eventName;
        uint256 closeTimeFromNow; // seconds from now
        uint256 settleTimeFromNow; // seconds from now
        uint256 tokenBetAmount; // Amount for token betting (in CCIP-BnM)
    }
    
    // Deployment results storage
    struct DeploymentResult {
        address marketRegistry;
        address[] predictionMarkets;
        address[] functionsConsumers;
        address sepoliaBridge;
        address fujiBridge;
        uint256 deployedMarketsCount;
    }

    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("wallet_private_key");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Deploying contracts with account:", deployer);
        console.log("Account balance:", deployer.balance);

        DeploymentResult memory result;
        SepoliaConfig memory sepolia = getSepoliaConfig();
        FujiConfig memory fuji = getFujiConfig();

        // =============== DEPLOY ON SEPOLIA ===============
        console.log("\n DEPLOYING ON SEPOLIA NETWORK");
        console.log("=====================================");
        
        vm.createSelectFork("sepolia");
        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy MarketRegistry
        console.log("\n=== Deploying MarketRegistry ===");
        MarketRegistry marketRegistry = new MarketRegistry();
        result.marketRegistry = address(marketRegistry);
        console.log("MarketRegistry deployed at:", address(marketRegistry));

        // 2. Deploy prediction markets and functions consumers
        MarketConfig[] memory markets = getMarketConfigs();
        result.predictionMarkets = new address[](markets.length);
        result.functionsConsumers = new address[](markets.length);
        
        for (uint256 i = 0; i < markets.length; i++) {
            console.log("\n=== Deploying Market", i + 1, "===");
            console.log("Question:", markets[i].question);
            
            uint256 closeTime = block.timestamp + markets[i].closeTimeFromNow;
            uint256 settleTime = block.timestamp + markets[i].settleTimeFromNow;
            
            // Deploy PredictionMarketExtended with token support
            PredictionMarketExtended market = new PredictionMarketExtended(
                markets[i].question,
                closeTime,
                settleTime,
                sepolia.ccipBnM, // CCIP-BnM token
                markets[i].tokenBetAmount, // Token bet amount
                true // Enable token betting
            );
            result.predictionMarkets[i] = address(market);
            console.log("PredictionMarketExtended deployed at:", address(market));
            
            // Deploy FunctionsConsumer
            string memory source = generateSourceCode(markets[i].eventName);
            FunctionsConsumer functionsConsumer = new FunctionsConsumer(
                sepolia.chainlinkRouter,
                address(market),
                source
            );
            result.functionsConsumers[i] = address(functionsConsumer);
            console.log("FunctionsConsumer deployed at:", address(functionsConsumer));
            
            // Set the functions consumer as odds updater
            market.setOddsUpdater(address(functionsConsumer));
            console.log("Set FunctionsConsumer as odds updater");
            
            // Register market in registry
            marketRegistry.registerMarket(
                address(market),
                address(functionsConsumer),
                markets[i].question,
                closeTime,
                settleTime
            );
            console.log("Market registered in registry");
        }
        result.deployedMarketsCount = markets.length;

        // 3. Deploy Sepolia CCIPBetBridge
        console.log("\n=== Deploying Sepolia CCIPBetBridge ===");
        
        // Prepare allowed chains and senders for Sepolia bridge
        uint64[] memory allowedChains = new uint64[](1);
        allowedChains[0] = fuji.chainSelector; // Allow Fuji
        
        address[] memory allowedSenders = new address[](1);
        allowedSenders[0] = deployer; // Add deployer as allowed sender (will be updated later)
        
        CCIPBetBridge sepoliaBridge = new CCIPBetBridge(
            sepolia.ccipRouter,
            sepolia.linkToken,
            result.predictionMarkets[0], // Use first market as example
            allowedChains,
            allowedSenders
        );
        result.sepoliaBridge = address(sepoliaBridge);
        console.log("Sepolia CCIPBetBridge deployed at:", address(sepoliaBridge));

        vm.stopBroadcast();

        // =============== DEPLOY ON FUJI ===============
        console.log("\n DEPLOYING ON FUJI NETWORK");
        console.log("===================================");
        
        vm.createSelectFork("fuji");
        vm.startBroadcast(deployerPrivateKey);

        // 4. Deploy Fuji CCIPBetBridge
        console.log("\n=== Deploying Fuji CCIPBetBridge ===");
        
        // Prepare allowed chains for Fuji bridge
        uint64[] memory fujiAllowedChains = new uint64[](1);
        fujiAllowedChains[0] = sepolia.chainSelector; // Allow Sepolia
        
        address[] memory fujiAllowedSenders = new address[](1);
        fujiAllowedSenders[0] = deployer; // Add deployer as allowed sender
        
        CCIPBetBridge fujiBridge = new CCIPBetBridge(
            fuji.ccipRouter,
            fuji.linkToken,
            address(0), // No prediction market on Fuji - this is sending side
            fujiAllowedChains,
            fujiAllowedSenders
        );
        result.fujiBridge = address(fujiBridge);
        console.log("Fuji CCIPBetBridge deployed at:", address(fujiBridge));

        vm.stopBroadcast();

        // =============== PRINT DEPLOYMENT SUMMARY ===============
        console.log("\n DEPLOYMENT COMPLETED SUCCESSFULLY!");
        console.log("==========================================");
        
        // Print deployment summary
        printDeploymentSummary(result, sepolia, fuji);
        
        // Print manual configuration steps
        printManualConfigurationSteps(result);
    }
    
    function getMarketConfigs() internal pure returns (MarketConfig[] memory) {
        MarketConfig[] memory markets = new MarketConfig[](4);
        
        markets[0] = MarketConfig({
            question: "Will England win the football match against France scheduled for June 18th, 2025 at 20:00 UTC?",
            eventName: "England vs France",
            closeTimeFromNow: 1 days,
            settleTimeFromNow: 2 days,
            tokenBetAmount: 10 * 10**18 // 10 CCIP-BnM (18 decimals)
        });
        
        markets[1] = MarketConfig({
            question: "Will China PR defeat Japan in the AFC Asian Cup match on June 20th, 2025?",
            eventName: "China PR vs Japan",
            closeTimeFromNow: 2 days,
            settleTimeFromNow: 3 days,
            tokenBetAmount: 5 * 10**18 // 5 CCIP-BnM
        });
        
        markets[2] = MarketConfig({
            question: "Canada vs United States, who will win on June 23rd, 2025?",
            eventName: "Canada vs United States",
            closeTimeFromNow: 5 days,
            settleTimeFromNow: 6 days,
            tokenBetAmount: 20 * 10**18 // 20 CCIP-BnM
        });
        
        markets[3] = MarketConfig({
            question: "Will Japan beat India in the match on June 23rd, 2025?",
            eventName: "Japan vs India",
            closeTimeFromNow: 6 days, 
            settleTimeFromNow: 7 days,
            tokenBetAmount: 15 * 10**18 // 15 CCIP-BnM
        });
        
        return markets;
    }
    
    function generateSourceCode(string memory eventName) internal view returns (string memory) {
        return string(abi.encodePacked(
            "try {",
            "  const oddsRequest = await Functions.makeHttpRequest({",
            "    url: '", API_BASE_URL, "',",
            "    method: 'POST',",
            "    headers: { ",
            "      'Content-Type': 'application/json',",
            "      'ngrok-skip-browser-warning': 'true'",
            "    },",
            "    data: { 'event': '", eventName, "' }",
            "  });",
            "  ",
            "  if (oddsRequest.error) {",
            "    throw new Error(`HTTP Error: ${oddsRequest.error}`);",
            "  }",
            "  ",
            "  const responseData = oddsRequest.data;",
            "  const oddsYes = responseData.yes;",
            "  const oddsNo = responseData.no;",
            "  ",
            "  if (!oddsYes || !oddsNo) {",
            "    throw new Error('Missing odds data in response');",
            "  }",
            "  ",
            "  return Functions.encodeString(`${oddsYes},${oddsNo}`);",
            "} catch (error) {",
            "  throw new Error(`Failed to fetch odds: ${error.message}`);",
            "}"
        ));
    }
    
    function printDeploymentSummary(
        DeploymentResult memory result,
        SepoliaConfig memory sepolia,
        FujiConfig memory fuji
    ) internal view {
        console.log("\n ===============================================");
        console.log(" COMPLETE DEPLOYMENT SUMMARY");
        console.log("===============================================");
        console.log("Deployer:", vm.addr(vm.envUint("wallet_private_key")));
        console.log("");
        
        console.log(" SEPOLIA TESTNET:");
        console.log("MarketRegistry:", result.marketRegistry);
        console.log("Sepolia CCIPBetBridge:", result.sepoliaBridge);
        console.log("");
        
        console.log(" FUJI TESTNET:");
        console.log("Fuji CCIPBetBridge:", result.fujiBridge);
        console.log("");
        
        console.log(" PREDICTION MARKETS (", result.deployedMarketsCount, " deployed):");
        MarketConfig[] memory markets = getMarketConfigs();
        for (uint256 i = 0; i < result.deployedMarketsCount; i++) {
            console.log("Market", i + 1, ":");
            console.log("  Question:", markets[i].question);
            console.log("  Contract:", result.predictionMarkets[i]);
            console.log("  FunctionsConsumer:", result.functionsConsumers[i]);
            console.log("  Event:", markets[i].eventName);
            console.log("");
        }
        
        console.log(" CONFIGURATION NEEDED:");
        console.log("1. Update frontend config.js:");
        console.log("   MARKET_REGISTRY_ADDRESS:", result.marketRegistry);
        console.log("   SEPOLIA_BRIDGE_ADDRESS:", result.sepoliaBridge);
        console.log("   FUJI_BRIDGE_ADDRESS:", result.fujiBridge);
        console.log("");
        console.log("2. Complete manual configuration steps (see below)");
        console.log("");
        
        console.log(" USAGE FLOW:");
        console.log("- Users on Sepolia: Bet directly on PredictionMarketExtended contracts");
        console.log("- Users on Fuji: Use CCIPBetBridge to send bets to Sepolia markets");
        console.log("- All markets registered in MarketRegistry for frontend discovery");
        console.log("- Functions consumers update odds automatically via Chainlink");
        console.log("- Token betting uses CCIP-BnM (get from Chainlink faucet)");
        console.log("");
        
        console.log(" NETWORK SELECTORS:");
        console.log("Sepolia:", sepolia.chainSelector);
        console.log("Fuji:", fuji.chainSelector);
        console.log("===============================================");
    }
    
    function printManualConfigurationSteps(DeploymentResult memory result) internal view {
        console.log("\n MANUAL CONFIGURATION REQUIRED:");
        console.log("=======================================");
        console.log("Due to cross-chain complexity, complete these steps manually:");
        console.log("");
        
        console.log("1. Configure Sepolia Bridge to allow Fuji Bridge:");
        console.log("   forge script script/DeployAll_new.s.sol \\");
        console.log("     --sig \"configureBridge(address,uint64,bool,address,bool)\" \\");
        console.log("     --rpc-url sepolia --broadcast \\");
        console.log("     --private-key $wallet_private_key \\");
        console.log("    ", result.sepoliaBridge, "\\");
        console.log("     14767482510784806043 true \\");
        console.log("    ", result.fujiBridge, "true");
        console.log("");
        
        console.log("2. Fund bridges with LINK tokens:");
        console.log("   forge script script/DeployAll_new.s.sol \\");
        console.log("     --sig \"fundBridges(address,address,uint256)\" \\");
        console.log("     --broadcast \\");
        console.log("     --private-key $wallet_private_key \\");
        console.log("    ", result.sepoliaBridge, "\\");
        console.log("    ", result.fujiBridge, "\\");
        console.log("     1000000000000000000"); // 1 LINK
        console.log("");
        
        console.log("3. Get test tokens:");
        console.log("   - LINK: https://faucets.chain.link/sepolia & https://faucets.chain.link/fuji");
        console.log("   - CCIP-BnM: https://faucets.chain.link/sepolia & https://faucets.chain.link/fuji");
        console.log("");
        
        console.log("4. Fund Chainlink Functions subscriptions for each FunctionsConsumer");
        console.log("=======================================");
    }
    
    // =============== HELPER FUNCTIONS ===============
     function getDeployedMarkets(address registryAddress) external view returns (address[] memory) {
        MarketRegistry registry = MarketRegistry(registryAddress);
        MarketRegistry.MarketInfo[] memory markets = registry.getAllMarkets();
        
        address[] memory marketAddresses = new address[](markets.length);
        for (uint256 i = 0; i < markets.length; i++) {
            marketAddresses[i] = markets[i].marketAddress;
        }
        
        return marketAddresses;
    }
    
    // Helper for manual bridge configuration
    function configureBridge(
        address payable bridgeAddress,
        uint64 chainSelector,
        bool allowed,
        address senderAddress,
        bool allowSender
    ) external {
        vm.startBroadcast();
        
        CCIPBetBridge bridge = CCIPBetBridge(bridgeAddress);
        bridge.allowSourceChain(chainSelector, allowed);
        bridge.allowSender(senderAddress, allowSender);
        
        vm.stopBroadcast();
        
        console.log("Bridge configured:", bridgeAddress);
        console.log("Chain selector:", chainSelector, "allowed:", allowed);
        console.log("Sender:", senderAddress, "allowed:", allowSender);
    }
    
    // Fund bridges with LINK tokens
    function fundBridges(address sepoliaBridge, address fujiBridge, uint256 amount) external {
        SepoliaConfig memory sepolia = getSepoliaConfig();
        FujiConfig memory fuji = getFujiConfig();
        
        // Fund Sepolia bridge
        vm.createSelectFork("sepolia");
        vm.startBroadcast();
        
        IERC20 sepoliaLink = IERC20(sepolia.linkToken);
        uint256 sepoliaBalance = sepoliaLink.balanceOf(msg.sender);
        
        if (sepoliaBalance >= amount) {
            sepoliaLink.transfer(sepoliaBridge, amount);
            console.log(" Funded Sepolia bridge with", amount, "LINK");
        } else {
            console.log(" Insufficient LINK on Sepolia. Balance:", sepoliaBalance, "Needed:", amount);
            console.log("Get LINK from: https://faucets.chain.link/sepolia");
        }
        
        vm.stopBroadcast();
        
        // Fund Fuji bridge  
        vm.createSelectFork("fuji");
        vm.startBroadcast();
        
        IERC20 fujiLink = IERC20(fuji.linkToken);
        uint256 fujiBalance = fujiLink.balanceOf(msg.sender);
        
        if (fujiBalance >= amount) {
            fujiLink.transfer(fujiBridge, amount);
            console.log(" Funded Fuji bridge with", amount, "LINK");
        } else {
            console.log(" Insufficient LINK on Fuji. Balance:", fujiBalance, "Needed:", amount);
            console.log("Get LINK from: https://faucets.chain.link/fuji");
        }
        
        vm.stopBroadcast();
        
        console.log("\n Bridge Funding Summary:");
        console.log("Sepolia bridge:", sepoliaBridge);
        console.log("Fuji bridge:", fujiBridge);
        console.log("Amount per bridge:", amount);
    }
}