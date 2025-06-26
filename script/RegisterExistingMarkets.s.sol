// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/MarketRegistry.sol";

contract RegisterExistingMarketsScript is Script {
    
    struct ExistingMarket {
        address marketAddress;
        address functionsConsumer;
        string question;
        uint256 closeTimeFromNow;
        uint256 settleTimeFromNow;
    }
    
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Registering existing markets with account:", deployer);

        // UPDATE THESE ADDRESSES WITH YOUR DEPLOYED CONTRACTS
        address REGISTRY_ADDRESS = 0x1234567890123456789012345678901234567890; // Update this
        
        vm.startBroadcast(deployerPrivateKey);

        MarketRegistry marketRegistry = MarketRegistry(REGISTRY_ADDRESS);
        
        // Register your existing markets
        ExistingMarket[] memory existingMarkets = getExistingMarkets();
        
        for (uint256 i = 0; i < existingMarkets.length; i++) {
            console.log("\n=== Registering Existing Market", i + 1, "===");
            console.log("Address:", existingMarkets[i].marketAddress);
            console.log("Question:", existingMarkets[i].question);
            
            uint256 closeTime = block.timestamp + existingMarkets[i].closeTimeFromNow;
            uint256 settleTime = block.timestamp + existingMarkets[i].settleTimeFromNow;
            
            try marketRegistry.registerMarket(
                existingMarkets[i].marketAddress,
                existingMarkets[i].functionsConsumer,
                existingMarkets[i].question,
                closeTime,
                settleTime
            ) {
                console.log("Market registered successfully");
            } catch Error(string memory reason) {
                console.log("Registration failed:", reason);
            }
        }

        vm.stopBroadcast();
        
        // Verify registrations
        console.log("\n=== Verification ===");
        console.log("Total markets in registry:", marketRegistry.getMarketsCount());
        console.log("Active markets:", marketRegistry.getActiveMarketsCount());
    }
    
    function getExistingMarkets() internal pure returns (ExistingMarket[] memory) {
        ExistingMarket[] memory markets = new ExistingMarket[](2);
        
        // UPDATE THESE WITH YOUR ACTUAL DEPLOYED CONTRACT ADDRESSES
        markets[0] = ExistingMarket({
            marketAddress: 0x4172386A4ad2D42f2ea4B5bED0D5cFd2e36f6a44, // Your existing PredictionMarket
            functionsConsumer: 0x5251177a3D5d6323793856D9687926534639005F, // Your existing FunctionsConsumer
            question: "Will it rain in Toronto tomorrow?", // Update with actual question
            closeTimeFromNow: 1 days,
            settleTimeFromNow: 2 days
        });
        
        markets[1] = ExistingMarket({
            marketAddress: 0x0000000000000000000000000000000000000000, // Add more markets as needed
            functionsConsumer: 0x0000000000000000000000000000000000000000,
            question: "Another test question",
            closeTimeFromNow: 3 days,
            settleTimeFromNow: 4 days
        });
        
        return markets;
    }
}