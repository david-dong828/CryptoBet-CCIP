// contractUtils.js - Contract Interaction Utilities (Extended)

import { BrowserProvider, Contract } from 'ethers';
import { CONFIG, PREDICTION_MARKET_ABI, FUNCTIONS_CONSUMER_ABI, MARKET_REGISTRY_ABI, CCIP_BRIDGE_ABI, ERC20_ABI, MARKET_CCIP_CONFIG } from './config';

// Helper function to format wei to ether
export const formatEther = (wei) => {
  if (!wei) return "0";
  try {
    const eth = parseFloat(wei.toString()) / 1e18;
    return eth.toFixed(4);
  } catch (error) {
    console.error("Error formatting ether:", error);
    return "0";
  }
};

// Helper function to format token amounts based on decimals
export const formatToken = (amount, decimals = 18) => {
  if (!amount) return "0";
  try {
    const divisor = Math.pow(10, decimals);
    const tokenAmount = parseFloat(amount.toString()) / divisor;
    return tokenAmount.toFixed(4);
  } catch (error) {
    console.error("Error formatting token:", error);
    return "0";
  }
};

// Helper function to parse ether to wei
export const parseEther = (eth) => {
  try {
    return BigInt(Math.floor(parseFloat(eth) * 1e18));
  } catch (error) {
    console.error("Error parsing ether:", error);
    return BigInt(0);
  }
};

// Helper function to parse token amounts based on decimals
export const parseToken = (amount, decimals = 18) => {
  try {
    const multiplier = Math.pow(10, decimals);
    return BigInt(Math.floor(parseFloat(amount) * multiplier));
  } catch (error) {
    console.error("Error parsing token:", error);
    return BigInt(0);
  }
};

// Format time since last update
export const formatTimeSince = (timestamp) => {
  const seconds = Math.floor((Date.now() - timestamp) / 1000);
  if (seconds < 60) return `${seconds}s ago`;
  const minutes = Math.floor(seconds / 60);
  if (minutes < 60) return `${minutes}m ago`;
  const hours = Math.floor(minutes / 60);
  return `${hours}h ago`;
};

// Get current chain ID
export const getCurrentChainId = async () => {
  if (!window.ethereum) return null;
  const chainId = await window.ethereum.request({ method: 'eth_chainId' });
  return parseInt(chainId, 16);
};

// Check if user is on Sepolia or Fuji
export const isSepoliaChain = (chainId) => chainId === CONFIG.CHAINS.SEPOLIA.chainId;
export const isFujiChain = (chainId) => chainId === CONFIG.CHAINS.FUJI.chainId;

// Switch to specific chain
export const switchToChain = async (chainId) => {
  if (!window.ethereum) {
    throw new Error("MetaMask not found");
  }

  const chainConfig = Object.values(CONFIG.CHAINS).find(chain => chain.chainId === chainId);
  if (!chainConfig) {
    throw new Error("Unsupported chain");
  }

  try {
    await window.ethereum.request({
      method: 'wallet_switchEthereumChain',
      params: [{ chainId: `0x${chainId.toString(16)}` }],
    });
  } catch (switchError) {
    // This error code indicates that the chain has not been added to MetaMask
    if (switchError.code === 4902) {
      try {
        await window.ethereum.request({
          method: 'wallet_addEthereumChain',
          params: [{
            chainId: `0x${chainId.toString(16)}`,
            chainName: chainConfig.name,
            rpcUrls: [chainConfig.rpcUrl],
            nativeCurrency: {
              name: chainConfig.currency,
              symbol: chainConfig.currency,
              decimals: 18,
            },
          }],
        });
      } catch (addError) {
        throw new Error("Failed to add chain to MetaMask");
      }
    } else {
      throw switchError;
    }
  }
};

// Connect to wallet and return provider, signer, and contracts
export const connectWallet = async () => {
  if (!window.ethereum) {
    throw new Error("Please install MetaMask!");
  }

  const accounts = await window.ethereum.request({
    method: 'eth_requestAccounts'
  });
  
  const chainId = await getCurrentChainId();
  const provider = new BrowserProvider(window.ethereum);
  const signer = await provider.getSigner();
  
  // Only create market registry if on Sepolia
  let marketRegistry = null;
  if (isSepoliaChain(chainId)) {
    marketRegistry = new Contract(
      CONFIG.MARKET_REGISTRY_ADDRESS,
      MARKET_REGISTRY_ABI,
      signer
    );
  }
  
  return {
    account: accounts[0],
    chainId,
    provider,
    signer,
    marketRegistry
  };
};

// Create specific market contract instance
export const createMarketContract = async (marketAddress, provider) => {
  const signer = await provider.getSigner();
  return new Contract(marketAddress, PREDICTION_MARKET_ABI, signer);
};

// Create functions consumer contract instance
export const createFunctionsContract = async (functionsAddress, provider) => {
  const signer = await provider.getSigner();
  return new Contract(functionsAddress, FUNCTIONS_CONSUMER_ABI, signer);
};

// Create CCIP bridge contract instance
export const createCCIPBridgeContract = async (bridgeAddress, provider) => {
  const signer = await provider.getSigner();
  return new Contract(bridgeAddress, CCIP_BRIDGE_ABI, signer);
};

// Create ERC20 token contract instance
export const createTokenContract = async (tokenAddress, provider) => {
  const signer = await provider.getSigner();
  return new Contract(tokenAddress, ERC20_ABI, signer);
};

// Fetch all markets from registry
export const fetchAllMarkets = async (marketRegistry) => {
  try {
    const markets = await marketRegistry.getActiveMarkets();
    
    return markets.map(market => ({
      address: market.marketAddress,
      functionsConsumer: market.functionsConsumer,
      question: market.question,
      closeTime: parseInt(market.closeTime.toString()) * 1000, // Convert to milliseconds
      settleTime: parseInt(market.settleTime.toString()) * 1000,
      createdAt: parseInt(market.createdAt.toString()) * 1000,
      creator: market.creator,
      isActive: market.isActive
    }));
  } catch (error) {
    console.error("Error fetching markets:", error);
    throw error;
  }
};

// Fetch market data with timing information (Extended version)
export const fetchMarketData = async (predictionMarket, account) => {
  try {
    const [
      question, 
      oddsYes, 
      oddsNo, 
      status, 
      result,
      marketTimes,
      timeRemaining,
      poolTotals,
      tokenBettingEnabled,
      betToken
    ] = await Promise.all([
      predictionMarket.question(),
      predictionMarket.oddsYes(),
      predictionMarket.oddsNo(),
      predictionMarket.status(),
      predictionMarket.result().catch(() => false),
      predictionMarket.getMarketTimes(),
      predictionMarket.getTimeRemaining(),
      predictionMarket.getPoolTotals(),
      predictionMarket.tokenBettingEnabled(),
      predictionMarket.betToken()
    ]);
    
    console.log('poolTotals: ',poolTotals)

    const marketData = {
      question,
      oddsYes: parseInt(oddsYes.toString()),
      oddsNo: parseInt(oddsNo.toString()),
      status: ["Open", "Closed", "Settled"][status],
      result: result,
      closeTime: parseInt(marketTimes[0].toString()) * 1000, // Convert to milliseconds
      settleTime: parseInt(marketTimes[1].toString()) * 1000,
      closeTimeRemaining: parseInt(timeRemaining[0].toString()) * 1000,
      settleTimeRemaining: parseInt(timeRemaining[1].toString()) * 1000,
      
      // Extended pool data
      totalYesETH: formatEther(poolTotals[0]),
      totalNoETH: formatEther(poolTotals[1]),
      totalYesToken: formatToken(poolTotals[2], 18), // CCIP-BnM has 18 decimals
      totalNoToken: formatToken(poolTotals[3], 18),
      totalETH: formatEther(poolTotals[4]),
      totalToken: formatToken(poolTotals[5], 18),
      
      // Token configuration
      tokenBettingEnabled,
      betToken: betToken === "0x0000000000000000000000000000000000000000" ? null : betToken
      
      // Legacy compatibility
      // totalYes: formatEther(poolTotals[4]), // Total pool size for backwards compatibility
      // totalNo: formatEther(poolTotals[4])
    };
    
    // Fetch user bets if account is connected
    let userBets = [];
    if (account) {
      try {
        const bets = await predictionMarket.getUserBets(account);
        userBets = bets.map((bet, index) => ({
          index,
          amount: bet.isToken ? formatToken(bet.amount, 18) : formatEther(bet.amount),
          prediction: bet.prediction,
          claimed: bet.claimed,
          isToken: bet.isToken,
          asset: bet.isToken ? 'CCIP-BnM' : 'ETH'
        }));
      } catch (error) {
        console.warn("Error fetching user bets:", error);
      }
    }
    console.log('userBets: ',userBets)
    return { marketData, userBets };
  } catch (error) {
    console.error("Error fetching market data:", error);
    throw error;
  }
};

// Check if market needs to be closed and close it
export const checkAndCloseMarket = async (predictionMarket) => {
  try {
    const tx = await predictionMarket.checkAndCloseMarket();
    console.log("Market close check transaction:", tx.hash);
    await tx.wait();
    return tx.hash;
  } catch (error) {
    console.error("Error checking/closing market:", error);
    throw error;
  }
};

// Request odds update via Chainlink Functions
export const requestOddsUpdate = async (functionsConsumer) => {
  const tx = await functionsConsumer.requestOdds(
    [], 
    CONFIG.CHAINLINK_SUBSCRIPTION_ID, 
    CONFIG.GAS_LIMIT
  );
  
  console.log("Odds update transaction sent:", tx.hash);
  await tx.wait();
  console.log("Transaction confirmed");
  
  return tx.hash;
};

// Place a bet on Sepolia (ETH)
export const placeBet = async (predictionMarket, isYes, amount) => {
  const tx = await predictionMarket.placeBet(isYes, {
    value: parseEther(amount)
  });
  
  console.log("Bet placed:", tx.hash);
  await tx.wait();
  
  return tx.hash;
};

// Place a bet via CCIP from Fuji (Token)
export const placeBetViaCCIP = async (ccipBridge, marketAddress, isYes, tokenAmount, provider) => {
  try {
    // Get CCIP-BnM token contract
    const ccipBnmToken = await createTokenContract(CONFIG.CCIP.CCIP_BNM_TOKEN_FUJI, provider);
    const signer = await provider.getSigner();
    const account = await signer.getAddress();
    
    // Check CCIP-BnM balance
    const balance = await ccipBnmToken.balanceOf(account);
    const requiredAmount = parseToken(tokenAmount, 18);
    
    if (balance < requiredAmount) {
      throw new Error(`Insufficient CCIP-BnM balance. Required: ${tokenAmount}, Available: ${formatToken(balance, 18)}`);
    }
    
    // Approve CCIP bridge to spend CCIP-BnM
    console.log("Approving CCIP-BnM for CCIP bridge...");
    const approveTx = await ccipBnmToken.approve(await ccipBridge.getAddress(), requiredAmount);
    await approveTx.wait();
    console.log("CCIP-BnM approval confirmed");
    
    // Calculate CCIP fee (we'll send some AVAX for fees)
    // const feeAmount = parseEther("0.01"); // 0.01 AVAX for fees
    
    const mapping = MARKET_CCIP_CONFIG[marketAddress];
    if (!mapping) throw new Error("No CCIP bridge mapping for this market");

    const destinationBridge = mapping.sepoliaReceiver; // should be the Sepolia CCIPBetBridge
    const eventId = marketAddress; // or use mapping.eventId if you store a string

    // Send bet via CCIP
    console.log("Sending bet via CCIP...");
    const tx = await ccipBridge.sendBetWithToken(
      CONFIG.CHAINS.SEPOLIA.ccipSelector, // destination chain
      destinationBridge, // receiver (market contract on Sepolia)
      eventId, // eventId (using market address as event ID)
      isYes, // prediction
      CONFIG.CCIP.CCIP_BNM_TOKEN_FUJI, // token
      requiredAmount, // amount
      true // payWithLINK (ALWAYS true. but false = pay with native AVAX)
    );
    
    console.log("CCIP bet transaction sent:", tx.hash);
    await tx.wait();
    console.log("CCIP bet transaction confirmed");
    
    return tx.hash;
  } catch (error) {
    console.error("Error placing bet via CCIP:", error);
    throw error;
  }
};

// Get user's token balance
export const getTokenBalance = async (tokenAddress, userAddress, provider) => {
  try {
    const tokenContract = await createTokenContract(tokenAddress, provider);
    const balance = await tokenContract.balanceOf(userAddress);
    return formatToken(balance, 18); // CCIP-BnM has 18 decimals
  } catch (error) {
    console.error("Error getting token balance:", error);
    return "0";
  }
};

// Claim winnings (works for both ETH and tokens)
export const claimWinnings = async (predictionMarket) => {
  const tx = await predictionMarket.claimWinnings();
  
  console.log("Claim transaction:", tx.hash);
  await tx.wait();
  
  return tx.hash;
};

// Calculate potential payout for ETH bets
export const calculatePotentialPayout = (betAmount, selectedBet, totalYesETH, totalNoETH) => {
  if (!betAmount || selectedBet === null) return 0;
  
  const amount = parseFloat(betAmount);
  const yesPool = parseFloat(totalYesETH);
  const noPool = parseFloat(totalNoETH);
  const totalPool = yesPool + noPool + amount;
  const selectedPool = selectedBet === 'yes' ? yesPool + amount : noPool + amount;
  
  return selectedPool > 0 ? (totalPool / selectedPool * amount).toFixed(4) : 0;
};

// Calculate potential payout for token bets
export const calculateTokenPayout = (betAmount, selectedBet, totalYesToken, totalNoToken) => {
  if (!betAmount || selectedBet === null) return 0;
  
  const amount = parseFloat(betAmount);
  const yesPool = parseFloat(totalYesToken);
  const noPool = parseFloat(totalNoToken);
  const totalPool = yesPool + noPool + amount;
  const selectedPool = selectedBet === 'yes' ? yesPool + amount : noPool + amount;
  
  return selectedPool > 0 ? (totalPool / selectedPool * amount).toFixed(4) : 0;
};

// Calculate actual winnings for a settled market
export const calculateWinnings = (userBet, totalYes, totalNo, result) => {
  if (!userBet) return 0;
  
  // Check if user won
  if (userBet.prediction !== result) return 0;
  
  const betAmount = parseFloat(userBet.amount);
  const totalPool = parseFloat(totalYes) + parseFloat(totalNo);
  const winningPool = result ? parseFloat(totalYes) : parseFloat(totalNo);
  
  if (winningPool === 0) return 0;
  
  return (totalPool * betAmount / winningPool).toFixed(4);
};

// Poll for odds updates
export const pollForOddsUpdate = async (fetchDataCallback, maxAttempts = CONFIG.MAX_POLLING_ATTEMPTS) => {
  let attempts = 0;
  
  return new Promise((resolve) => {
    const pollInterval = setInterval(async () => {
      attempts++;
      console.log(`Polling attempt ${attempts}/${maxAttempts}`);
      
      try {
        await fetchDataCallback();
      } catch (error) {
        console.error("Polling error:", error);
      }
      
      if (attempts >= maxAttempts) {
        clearInterval(pollInterval);
        console.log("Polling completed");
        resolve();
      }
    }, CONFIG.POLLING_INTERVAL);
  });
};

// Fetch market stats for list view (Extended)
export const fetchMarketStats = async (marketAddress, provider) => {
  try {
    const marketContract = await createMarketContract(marketAddress, provider);
    
    const [
      question,
      status,
      oddsYes,
      oddsNo,
      timeRemaining,
      poolTotals,
      tokenBettingEnabled
    ] = await Promise.all([
      marketContract.question(),
      marketContract.status(),
      marketContract.oddsYes(),
      marketContract.oddsNo(),
      marketContract.getTimeRemaining(),
      marketContract.getPoolTotals(),
      marketContract.tokenBettingEnabled()
    ]);
    
    const totalETHPool = parseFloat(formatEther(poolTotals[4]));
    const totalTokenPool = parseFloat(formatToken(poolTotals[5], 18));
    
    return {
      address: marketAddress,
      question,
      status: ["Open", "Closed", "Settled"][status],
      
      // Pool information
      totalETHPool: totalETHPool.toFixed(4),
      totalTokenPool: totalTokenPool.toFixed(4),
      totalYesETH: formatEther(poolTotals[0]),
      totalNoETH: formatEther(poolTotals[1]),
      totalYesToken: formatToken(poolTotals[2], 18),
      totalNoToken: formatToken(poolTotals[3], 18),
      
      // Combined for display
      totalPool: (totalETHPool + totalTokenPool).toFixed(4),
      participants: Math.floor(Math.random() * 100) + 20, // Mock participants for now
      
      // Odds
      oddsYes: parseInt(oddsYes.toString()),
      oddsNo: parseInt(oddsNo.toString()),
      
      // Timing
      closeTimeRemaining: parseInt(timeRemaining[0].toString()) * 1000,
      settleTimeRemaining: parseInt(timeRemaining[1].toString()) * 1000,
      
      // Token support
      tokenBettingEnabled
    };
  } catch (error) {
    console.error("Error fetching market stats:", error);
    return null;
  }
};

// Get countdown display text
export const getCountdownText = (timeRemaining) => {
  if (timeRemaining <= 0) return 'Ended';
  
  const seconds = Math.floor(timeRemaining / 1000);
  const days = Math.floor(seconds / (24 * 60 * 60));
  const hours = Math.floor((seconds % (24 * 60 * 60)) / (60 * 60));
  const minutes = Math.floor((seconds % (60 * 60)) / 60);
  
  if (days > 0) return `${days}d ${hours}h`;
  if (hours > 0) return `${hours}h ${minutes}m`;
  if (minutes > 0) return `${minutes}m`;
  return `${seconds}s`;
};

// Chain-specific betting function
export const placeBetBasedOnChain = async (
  chainId, 
  predictionMarket, 
  ccipBridge, 
  marketAddress, 
  isYes, 
  amount, 
  provider
) => {
  if (isSepoliaChain(chainId)) {
    // Place ETH bet directly on Sepolia
    return await placeBet(predictionMarket, isYes, amount);
  } else if (isFujiChain(chainId)) {
    // Place token bet via CCIP from Fuji
    return await placeBetViaCCIP(ccipBridge, marketAddress, isYes, amount, provider);
  } else {
    throw new Error(`Unsupported chain. Please switch to Sepolia (${CONFIG.CHAINS.SEPOLIA.chainId}) or Fuji (${CONFIG.CHAINS.FUJI.chainId})`);
  }
};

// Get chain-specific currency symbol
export const getChainCurrency = (chainId) => {
  if (isSepoliaChain(chainId)) return 'ETH';
  if (isFujiChain(chainId)) return 'CCIP-BnM';
  return 'Unknown';
};

// Get chain-specific bet amount formatting
export const formatBetAmount = (amount, chainId) => {
  if (isSepoliaChain(chainId)) {
    return `${amount} ETH`;
  } else if (isFujiChain(chainId)) {
    return `${amount} CCIP-BnM`;
  }
  return amount;
};

// Get minimum bet amount for chain
export const getMinimumBetAmount = (chainId) => {
  if (isSepoliaChain(chainId)) return "0.01"; // 0.01 ETH
  if (isFujiChain(chainId)) return "0.001"; // 0.001 CCIP-BnM
  return "0";
};

// Validate bet amount for chain
export const validateBetAmount = (amount, chainId) => {
  const minAmount = parseFloat(getMinimumBetAmount(chainId));
  const betAmount = parseFloat(amount);
  
  if (isNaN(betAmount) || betAmount <= 0) {
    return { valid: false, error: "Please enter a valid amount" };
  }
  
  if (betAmount < minAmount) {
    const currency = getChainCurrency(chainId);
    return { valid: false, error: `Minimum bet is ${minAmount} ${currency}` };
  }
  
  return { valid: true };
};

// Check if user has sufficient balance
export const checkSufficientBalance = async (amount, chainId, account, provider) => {
  try {
    if (isSepoliaChain(chainId)) {
      // Check ETH balance
      const balance = await provider.getBalance(account);
      const ethBalance = parseFloat(formatEther(balance));
      const requiredAmount = parseFloat(amount);
      
      return {
        sufficient: ethBalance >= requiredAmount,
        balance: ethBalance.toFixed(4),
        currency: 'ETH'
      };
    } else if (isFujiChain(chainId)) {
      // Check CCIP-BnM balance
      const tokenBalance = await getTokenBalance(CONFIG.CCIP.CCIP_BNM_TOKEN_FUJI, account, provider);
      const balance = parseFloat(tokenBalance);
      const requiredAmount = parseFloat(amount);
      
      return {
        sufficient: balance >= requiredAmount,
        balance: balance.toFixed(4),
        currency: 'CCIP-BnM'
      };
    }
    
    return { sufficient: false, balance: "0", currency: "Unknown" };
  } catch (error) {
    console.error("Error checking balance:", error);
    return { sufficient: false, balance: "0", currency: "Error" };
  }
};