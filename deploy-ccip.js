// deploy-ccip.js - Complete deployment script for CCIP-enabled prediction markets
const { ethers } = require('ethers');
const fs = require('fs');
require('dotenv').config({ path: './new.env' });


// Chain configurations
const CHAIN_CONFIG = {
  sepolia: {
    chainId: 11155111,
    router: "0x0BF3dE8c5D3e8A2B34D2BEeB17ABfCeBaf363A59",
    link: "0x779877A7B0D9E8603169DdbD7836e478b4624789",
    ccipBnM: "0xFd57b4ddBf88a4e07fF4e34C487b99af2Fe82a05",
    chainSelector: "16015286601757825753"
  },
  fuji: {
    chainId: 43113,
    router: "0xF694E193200268f9a4868e4Aa017A0118C9a8177",
    link: "0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846",
    ccipBnM: "0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4",
    chainSelector: "14767482510784806043"
  }
};

// Read compiled contracts
function getContractArtifacts(contractName) {
  try {
    // Try different possible paths for Foundry artifacts
    const possiblePaths = [
      `./out/${contractName}.sol/${contractName}.json`,
      `./artifacts/contracts/${contractName}.sol/${contractName}.json`,
      `./artifacts/${contractName}.json`
    ];
    
    for (const artifactPath of possiblePaths) {
      if (fs.existsSync(artifactPath)) {
        const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'));
        return {
          abi: artifact.abi,
          bytecode: artifact.bytecode.object || artifact.bytecode
        };
      }
    }
    
    throw new Error(`Could not find artifacts for ${contractName}`);
  } catch (error) {
    console.error(`Error loading ${contractName} artifacts:`, error.message);
    throw error;
  }
}

async function deployContract(contractName, wallet, ...constructorArgs) {
  console.log(` Deploying ${contractName}...`);
  
  const { abi, bytecode } = getContractArtifacts(contractName);
  
  if (!bytecode || bytecode === '0x') {
    throw new Error(`No bytecode found for ${contractName}. Make sure the contract is compiled.`);
  }
  
  const factory = new ethers.ContractFactory(abi, bytecode, wallet);
  const contract = await factory.deploy(...constructorArgs);
  
  console.log(`   TX Hash: ${contract.deployTransaction.hash}`);
  console.log(`   Waiting for confirmation...`);
  
  await contract.deployed();
  console.log(` ${contractName} deployed at: ${contract.address}`);
  
  return contract;
}

async function deployPredictionMarketWithCCIP(network, config) {
  console.log(`\n Deploying to ${network}...`);
  
  const provider = new ethers.providers.JsonRpcProvider(
    network === 'sepolia' ? process.env.ALCHEMY_API_URL : process.env.FUJI_RPC_URL
  );
  const wallet = new ethers.Wallet(process.env.wallet_private_key, provider);
  
  console.log(` Deploying from: ${wallet.address}`);
  
  // Check wallet balance
  const balance = await wallet.getBalance();
  console.log(` Balance: ${ethers.utils.formatEther(balance)} ${network === 'sepolia' ? 'ETH' : 'AVAX'}`);
  
  if (balance.eq(0)) {
    throw new Error(`No balance on ${network}. Get test tokens from faucet.`);
  }
  
  // Deploy PredictionMarketExtended
  const closeTime = Math.floor(Date.now() / 1000) + (24 * 60 * 60); // 24 hours from now
  const settleTime = closeTime + (60 * 60); // 1 hour after close
  
  const predictionMarket = await deployContract(
    'PredictionMarketExtended',
    wallet,
    config.question,
    closeTime,
    settleTime,
    CHAIN_CONFIG[network].ccipBnM, // Bet token
    ethers.utils.parseEther("0.001"), // Token bet amount
    true // Enable token betting
  );
  
  // Deploy CCIPBetBridge with configuration
  // Configure allowed chains and senders based on network
  let allowedChains = [];
  let allowedSenders = [wallet.address]; // Always allow deployer
  
  if (network === 'sepolia') {
    // Sepolia receives from Fuji
    allowedChains = [CHAIN_CONFIG.fuji.chainSelector];
  } else if (network === 'fuji') {
    // Fuji sends to Sepolia (no receiving setup needed for sender)
    allowedChains = [];
  }
  
  const ccipBridge = await deployContract(
    'CCIPBetBridge',
    wallet,
    CHAIN_CONFIG[network].router,
    CHAIN_CONFIG[network].link,
    predictionMarket.address,
    allowedChains,
    allowedSenders
  );
  
  console.log(" Configuration complete!");
  
  return {
    predictionMarket: predictionMarket.address,
    ccipBridge: ccipBridge.address,
    config: {
      network,
      chainId: CHAIN_CONFIG[network].chainId,
      closeTime,
      settleTime
    }
  };
}

// NEW FEATURE: Deploy CCIP bridges for multiple existing prediction markets
async function deployMultiCCIPBridges(network) {
  console.log(`\n Deploying CCIP Bridges for Multiple Prediction Markets on ${network}...`);
  console.log("========================================================================");
  
  const provider = new ethers.providers.JsonRpcProvider(
    network === 'sepolia' ? process.env.ALCHEMY_API_URL : process.env.FUJI_RPC_URL
  );
  const wallet = new ethers.Wallet(process.env.wallet_private_key, provider);
  
  console.log(` Deploying from: ${wallet.address}`);
  
  // Check wallet balance
  const balance = await wallet.getBalance();
  console.log(` Balance: ${ethers.utils.formatEther(balance)} ${network === 'sepolia' ? 'ETH' : 'AVAX'}`);
  
  if (balance.eq(0)) {
    throw new Error(`No balance on ${network}. Get test tokens from faucet.`);
  }
  
  // Get prediction market addresses from environment
  const predictionMarkets = [];
  for (let i = 1; i <= 4; i++) {
    const envKey = `predictMarket_${i}`;
    const address = process.env[envKey];
    
    if (address && ethers.utils.isAddress(address)) {
      predictionMarkets.push({
        id: i,
        address: address,
        envKey: envKey
      });
      console.log(` Found ${envKey}: ${address}`);
    } else {
      console.log(`  ${envKey} not found or invalid address in .env file`);
    }
  }
  
  if (predictionMarkets.length === 0) {
    throw new Error('No valid prediction market addresses found in .env file');
  }
  
  console.log(`\n Found ${predictionMarkets.length} prediction market(s) to bridge`);
  
  // Configure allowed chains and senders based on network
  let allowedChains = [];
  let allowedSenders = [wallet.address]; // Always allow deployer
  
  if (network === 'sepolia') {
    // Sepolia receives from Fuji
    allowedChains = [CHAIN_CONFIG.fuji.chainSelector];
  } else if (network === 'fuji') {
    // Fuji sends to Sepolia (no receiving setup needed for sender)
    allowedChains = [];
  }
  
  const deployments = [];
  
  // Deploy CCIP bridge for each prediction market
  for (const market of predictionMarkets) {
    console.log(`\n Deploying CCIP Bridge for ${market.envKey} (${market.address})...`);
    
    try {
      // For Sepolia: Verify the prediction market exists
      // For Fuji: Skip verification since contracts are on Sepolia, just use the address for bridge config
      if (network === 'sepolia') {
        const code = await provider.getCode(market.address);
        if (code === '0x' || code === '0x0') {
          console.log(` Contract at ${market.address} not found on Sepolia`);
          continue;
        }
        console.log(` Verified prediction market exists on Sepolia`);
      } else {
        console.log(` Creating bridge for Sepolia market ${market.address} (sender on Fuji)`);
      }
      
      const ccipBridge = await deployContract(
        'CCIPBetBridge',
        wallet,
        CHAIN_CONFIG[network].router,
        CHAIN_CONFIG[network].link,
        market.address, // Use Sepolia market address even when deploying on Fuji
        allowedChains,
        allowedSenders
      );
      
      const deployment = {
        marketId: market.id,
        marketEnvKey: market.envKey,
        predictionMarket: market.address, // Sepolia market address
        ccipBridge: ccipBridge.address,   // Bridge address on current network
        network: network,
        chainId: CHAIN_CONFIG[network].chainId,
        bridgeType: network === 'sepolia' ? 'receiver' : 'sender'
      };
      
      deployments.push(deployment);
      console.log(` ${deployment.bridgeType} bridge deployed for ${market.envKey}: ${ccipBridge.address}`);
      
    } catch (error) {
      console.error(` Failed to deploy bridge for ${market.envKey}:`, error.message);
    }
  }
  
  if (deployments.length === 0) {
    throw new Error('No bridges were successfully deployed');
  }
  
  // Save deployment info
  const deploymentInfo = {
    timestamp: new Date().toISOString(),
    network: network,
    type: 'multi-ccip-bridges',
    totalDeployed: deployments.length,
    deployments: deployments
  };
  
  console.log("\n Multi-Bridge Deployment Complete!");
  console.log("===================================");
  console.log(JSON.stringify(deploymentInfo, null, 2));
  
  // Save to file
  const filename = `multi-bridge-deployment-${network}-${Date.now()}.json`;
  fs.writeFileSync(filename, JSON.stringify(deploymentInfo, null, 2));
  console.log(`\n Deployment info saved to: ${filename}`);
  
  console.log("\n Next Steps:");
  if (network === 'fuji') {
    console.log(`1. Fund each Fuji bridge with LINK for CCIP fees (sender bridges):`);
    deployments.forEach(d => {
      console.log(`   Send LINK to ${d.ccipBridge} (${d.bridgeType} for ${d.marketEnvKey})`);
    });
  } else {
    console.log(`1. Fund each bridge with LINK for CCIP fees:`);
    deployments.forEach(d => {
      console.log(`   Send LINK to ${d.ccipBridge} (${d.bridgeType} for ${d.marketEnvKey})`);
    });
  }
  
  console.log("\n2. Update .env with new bridge addresses:");
  deployments.forEach(d => {
    const envSuffix = network === 'sepolia' ? 'SEP' : 'FUJI';
    console.log(`   ${envSuffix}_CCIP_BRIDGE_${d.marketId}=${d.ccipBridge}`);
  });
  
  return deploymentInfo;
}

// Deploy multi-bridges for both networks
async function deployFullMultiCCIPSystem() {
  console.log(" Deploying Multi-CCIP Bridge System");
  console.log("======================================");
  
  try {
    // Check compiled contracts
    console.log("\n Checking compiled contracts...");
    try {
      getContractArtifacts('CCIPBetBridge');
      console.log(" CCIPBetBridge contract found");
    } catch (error) {
      console.log("  CCIPBetBridge not found. Compiling with Forge...");
      const { execSync } = require('child_process');
      try {
        execSync('forge build', { stdio: 'inherit' });
      } catch (e) {
        console.log("Could not compile with Forge. Make sure contracts are compiled.");
        throw e;
      }
    }
    
    // Deploy bridges on Sepolia (receiving chain)
    const sepoliaDeployment = await deployMultiCCIPBridges('sepolia');
    
    // Deploy bridges on Fuji (sending chain)  
    const fujiDeployment = await deployMultiCCIPBridges('fuji');
    
    // Create cross-chain configuration
    const crossChainConfig = {};
    for (let i = 1; i <= 4; i++) {
      const sepoliaBridge = sepoliaDeployment.deployments.find(d => d.marketId === i);
      const fujiBridge = fujiDeployment.deployments.find(d => d.marketId === i);
      
      if (sepoliaBridge && fujiBridge) {
        crossChainConfig[`market_${i}_fujiToSepolia`] = {
          source: fujiBridge.ccipBridge,
          destination: sepoliaBridge.ccipBridge,
          chainSelector: CHAIN_CONFIG.sepolia.chainSelector,
          predictionMarket: sepoliaBridge.predictionMarket
        };
      }
    }
    
    // Complete deployment info
    const fullDeploymentInfo = {
      timestamp: new Date().toISOString(),
      type: 'full-multi-ccip-system',
      sepolia: sepoliaDeployment,
      fuji: fujiDeployment,
      crossChainConfig: crossChainConfig
    };
    
    console.log("\n Full Multi-Bridge System Deployment Complete!");
    console.log("===============================================");
    console.log(JSON.stringify(fullDeploymentInfo, null, 2));
    
    // Save to file
    const filename = `full-multi-bridge-deployment-${Date.now()}.json`;
    fs.writeFileSync(filename, JSON.stringify(fullDeploymentInfo, null, 2));
    console.log(`\n📄 Full deployment info saved to: ${filename}`);
    
    return fullDeploymentInfo;
    
  } catch (error) {
    console.error("\n Multi-bridge deployment failed:", error.message);
    throw error;
  }
}

async function deployFullCCIPSystem() {
  console.log(" Deploying Full CCIP Prediction Market System");
  console.log("==============================================");
  
  const marketConfig = {
    question: "Will Team A beat Team B in the championship?",
  };
  
  try {
    // First compile contracts if using Foundry
    console.log("\n Checking compiled contracts...");
    try {
      getContractArtifacts('PredictionMarketExtended');
      getContractArtifacts('CCIPBetBridge');
      console.log(" Contracts found");
    } catch (error) {
      console.log("  Contracts not found. Compiling with Forge...");
      const { execSync } = require('child_process');
      try {
        execSync('forge build', { stdio: 'inherit' });
      } catch (e) {
        console.log("Could not compile with Forge. Make sure contracts are compiled.");
      }
    }
    
    // Deploy on Sepolia (receiving chain)
    const sepoliaDeployment = await deployPredictionMarketWithCCIP('sepolia', marketConfig);
    
    // Deploy on Fuji (sending chain)
    const fujiDeployment = await deployPredictionMarketWithCCIP('fuji', marketConfig);
    
    // Save deployment info
    const deploymentInfo = {
      timestamp: new Date().toISOString(),
      sepolia: sepoliaDeployment,
      fuji: fujiDeployment,
      crossChainConfig: {
        fujiToSepolia: {
          source: fujiDeployment.ccipBridge,
          destination: sepoliaDeployment.ccipBridge,
          chainSelector: CHAIN_CONFIG.sepolia.chainSelector
        }
      }
    };
    
    console.log("\n Deployment Complete!");
    console.log("========================");
    console.log(JSON.stringify(deploymentInfo, null, 2));
    
    // Save to file
    const filename = `deployment-${Date.now()}.json`;
    fs.writeFileSync(filename, JSON.stringify(deploymentInfo, null, 2));
    console.log(`\n📄 Deployment info saved to: ${filename}`);
    
    console.log("\n📋 Next Steps:");
    console.log("1. Fund Fuji bridge with LINK for CCIP fees:");
    console.log(`   Send LINK to: ${fujiDeployment.ccipBridge}`);
    console.log("2. Update .env with new addresses:");
    console.log(`   FUJI_CROSS_SENDER=${fujiDeployment.ccipBridge}`);
    console.log(`   SEP_CCIP_RECEIVER=${sepoliaDeployment.ccipBridge}`);
    console.log(`   SEP_PREDICTION_MARKET=${sepoliaDeployment.predictionMarket}`);
    
  } catch (error) {
    console.error("\n Deployment failed:", error.message);
    if (error.message.includes('artifacts')) {
      console.error("\n Make sure to compile your contracts first:");
      console.error("   forge build");
    }
  }
}

// Standalone deployment functions
async function deploySepolia() {
  const config = {
    question: "Will Team A beat Team B?",
  };
  
  try {
    const deployment = await deployPredictionMarketWithCCIP('sepolia', config);
    console.log("\n Sepolia Deployment:", deployment);
  } catch (error) {
    console.error(" Deployment failed:", error.message);
  }
}

async function deployFuji() {
  const config = {
    question: "Will Team A beat Team B?",
  };
  
  try {
    const deployment = await deployPredictionMarketWithCCIP('fuji', config);
    console.log("\n Fuji Deployment:", deployment);
  } catch (error) {
    console.error(" Deployment failed:", error.message);
  }
}

// Helper function to verify deployment
async function verifyDeployment(network, addresses) {
  console.log(`\n Verifying ${network} deployment...`);
  
  const provider = new ethers.providers.JsonRpcProvider(
    network === 'sepolia' ? process.env.ALCHEMY_API_URL : process.env.FUJI_RPC_URL
  );
  
  // Check contract code
  const marketCode = await provider.getCode(addresses.predictionMarket);
  const bridgeCode = await provider.getCode(addresses.ccipBridge);
  
  console.log(`PredictionMarket deployed: ${marketCode !== '0x' && marketCode !== '0x0'}`);
  console.log(`CCIPBridge deployed: ${bridgeCode !== '0x' && bridgeCode !== '0x0'}`);
  
  // Check configurations
  const bridgeAbi = [
    "function allowlistedSourceChains(uint64) view returns (bool)",
    "function allowlistedSenders(address) view returns (bool)",
    "function predictionMarket() view returns (address)"
  ];
  
  const bridge = new ethers.Contract(addresses.ccipBridge, bridgeAbi, provider);
  
  if (network === 'sepolia') {
    const fujiAllowed = await bridge.allowlistedSourceChains(CHAIN_CONFIG.fuji.chainSelector);
    console.log(`Fuji chain allowed: ${fujiAllowed}`);
  }
  
  const marketAddress = await bridge.predictionMarket();
  console.log(`PredictionMarket configured: ${marketAddress === addresses.predictionMarket}`);
}

// Check if contracts are compiled
async function checkCompiled() {
  console.log("\n Checking compiled contracts...");
  
  const contracts = ['PredictionMarketExtended', 'CCIPBetBridge'];
  let allFound = true;
  
  for (const contract of contracts) {
    try {
      getContractArtifacts(contract);
      console.log(` ${contract} found`);
    } catch (error) {
      console.log(` ${contract} not found`);
      allFound = false;
    }
  }
  
  if (!allFound) {
    console.log("\n Compile contracts with: forge build");
  }
  
  return allFound;
}

// CLI interface
const command = process.argv[2];

switch (command) {
  case 'full':
    deployFullCCIPSystem();
    break;
  case 'multi-bridges':
    const network = process.argv[3];
    if (!network || !['sepolia', 'fuji'].includes(network)) {
      console.log("Usage: node deploy-ccip.js multi-bridges <network>");
      console.log("Network must be: sepolia or fuji");
      break;
    }
    deployMultiCCIPBridges(network);
    break;
  case 'multi-full':
    deployFullMultiCCIPSystem();
    break;
  case 'sepolia':
    deploySepolia();
    break;
  case 'fuji':
    deployFuji();
    break;
  case 'verify':
    const verifyNetwork = process.argv[3];
    const addresses = {
      predictionMarket: process.argv[4],
      ccipBridge: process.argv[5]
    };
    if (!verifyNetwork || !addresses.predictionMarket || !addresses.ccipBridge) {
      console.log("Usage: node deploy-ccip.js verify <network> <marketAddr> <bridgeAddr>");
      break;
    }
    verifyDeployment(verifyNetwork, addresses);
    break;
  case 'check':
    checkCompiled();
    break;
  default:
    console.log(`
Usage: node deploy-ccip.js <command>

Commands:
  full                - Deploy complete system on both chains
  multi-bridges <network> - Deploy CCIP bridges for multiple existing prediction markets
  multi-full          - Deploy CCIP bridges for all prediction markets on both chains
  sepolia             - Deploy only on Sepolia
  fuji                - Deploy only on Fuji
  check               - Check if contracts are compiled
  verify <network> <marketAddr> <bridgeAddr> - Verify deployment

Example:
  node deploy-ccip.js check
  node deploy-ccip.js full
  node deploy-ccip.js multi-bridges sepolia
  node deploy-ccip.js multi-full
  node deploy-ccip.js verify sepolia 0x123... 0x456...

Prerequisites:
  1. Compile contracts first: forge build
  2. Set up .env with RPC URLs and private key
  3. Have test ETH/AVAX on both chains
  4. For multi-bridges: Set predictMarket_1, predictMarket_2, predictMarket_3, predictMarket_4 in .env
    `);
}

module.exports = { deployPredictionMarketWithCCIP, deployMultiCCIPBridges, CHAIN_CONFIG };