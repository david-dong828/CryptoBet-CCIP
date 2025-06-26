export const CONFIG = {
  // Contract Addresses (Sepolia) - Update these after deployment
  MARKET_REGISTRY_ADDRESS: "0x2FeAFA8823B542928fd90eb4534C7EE27946B758", 
  // PREDICTION_MARKET_ADDRESS: "0x4172386A4ad2D42f2ea4B5bED0D5cFd2e36f6a44",
  // FUNCTIONS_CONSUMER_ADDRESS: "0x5251177a3D5d6323793856D9687926534639005F",

  // Chain configurations
  CHAINS: {
    SEPOLIA: {
      chainId: 11155111,
      name: "Sepolia",
      currency: "ETH",
      rpcUrl: "https://eth-sepolia.g.alchemy.com/v2/WYseGzoJeIJ1t8Gn35_fD",
      ccipSelector: "16015286601757825753" // Sepolia CCIP selector
    },
    FUJI: {
      chainId: 43113,
      name: "Avalanche Fuji",
      currency: "AVAX",
      rpcUrl: "https://avax-fuji.g.alchemy.com/v2/_pTIrjckMhqhmet1HBugg",
      ccipSelector: "14767482510784806043" // Fuji CCIP selector
    }
  },
  
  // CCIP Configuration
  CCIP: {
    BRIDGE_ADDRESS_FUJI: "0x6c12AC815F9d6383dcDC0318ca09c2B9aDF537Ad", // Deploy CCIPBetBridge on Fuji
    BRIDGE_ADDRESS_SEPOLIA: "0xbE6a2088f9BdDC0438cab1a54c8AA34B3Cb2cD7d",
    LINK_TOKEN_FUJI: "0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846", // LINK on Fuji
    CCIP_BNM_TOKEN_FUJI: "0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4", // CCIP-BnM on Fuji
    DEFAULT_TOKEN_AMOUNT: "1000000000000000", // 0.001 CCIP-BnM (18 decimals)
    GAS_LIMIT: 900000
  },
  
  // Chainlink Functions (from https://functions.chain.link)
  CHAINLINK_SUBSCRIPTION_ID: "4922",
  GAS_LIMIT: 300000,
  
  // Auto-refresh interval (ms)
  REFRESH_INTERVAL: 10000, // 10 seconds
  
  // Polling settings for odds updates
  POLLING_INTERVAL: 5000, // 5 seconds
  MAX_POLLING_ATTEMPTS: 12 // 60 seconds total
};

export const MARKET_REGISTRY_ABI = [
  {
    "inputs": [],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "internalType": "address", "name": "marketAddress", "type": "address"}
    ],
    "name": "MarketDeactivated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "internalType": "address", "name": "marketAddress", "type": "address"},
      {"indexed": true, "internalType": "address", "name": "creator", "type": "address"},
      {"indexed": false, "internalType": "string", "name": "question", "type": "string"},
      {"indexed": false, "internalType": "uint256", "name": "closeTime", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "settleTime", "type": "uint256"}
    ],
    "name": "MarketRegistered",
    "type": "event"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_marketAddress", "type": "address"}
    ],
    "name": "deactivateMarket",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getActiveMarkets",
    "outputs": [
      {
        "components": [
          {"internalType": "address", "name": "marketAddress", "type": "address"},
          {"internalType": "address", "name": "functionsConsumer", "type": "address"},
          {"internalType": "string", "name": "question", "type": "string"},
          {"internalType": "uint256", "name": "closeTime", "type": "uint256"},
          {"internalType": "uint256", "name": "settleTime", "type": "uint256"},
          {"internalType": "uint256", "name": "createdAt", "type": "uint256"},
          {"internalType": "address", "name": "creator", "type": "address"},
          {"internalType": "bool", "name": "isActive", "type": "bool"}
        ],
        "internalType": "struct MarketRegistry.MarketInfo[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getActiveMarketsCount",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getAllMarkets",
    "outputs": [
      {
        "components": [
          {"internalType": "address", "name": "marketAddress", "type": "address"},
          {"internalType": "address", "name": "functionsConsumer", "type": "address"},
          {"internalType": "string", "name": "question", "type": "string"},
          {"internalType": "uint256", "name": "closeTime", "type": "uint256"},
          {"internalType": "uint256", "name": "settleTime", "type": "uint256"},
          {"internalType": "uint256", "name": "createdAt", "type": "uint256"},
          {"internalType": "address", "name": "creator", "type": "address"},
          {"internalType": "bool", "name": "isActive", "type": "bool"}
        ],
        "internalType": "struct MarketRegistry.MarketInfo[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_marketAddress", "type": "address"}
    ],
    "name": "getMarketInfo",
    "outputs": [
      {
        "components": [
          {"internalType": "address", "name": "marketAddress", "type": "address"},
          {"internalType": "address", "name": "functionsConsumer", "type": "address"},
          {"internalType": "string", "name": "question", "type": "string"},
          {"internalType": "uint256", "name": "closeTime", "type": "uint256"},
          {"internalType": "uint256", "name": "settleTime", "type": "uint256"},
          {"internalType": "uint256", "name": "createdAt", "type": "uint256"},
          {"internalType": "address", "name": "creator", "type": "address"},
          {"internalType": "bool", "name": "isActive", "type": "bool"}
        ],
        "internalType": "struct MarketRegistry.MarketInfo",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getMarketsCount",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_user", "type": "address"}
    ],
    "name": "getUserMarkets",
    "outputs": [
      {
        "components": [
          {"internalType": "address", "name": "marketAddress", "type": "address"},
          {"internalType": "address", "name": "functionsConsumer", "type": "address"},
          {"internalType": "string", "name": "question", "type": "string"},
          {"internalType": "uint256", "name": "closeTime", "type": "uint256"},
          {"internalType": "uint256", "name": "settleTime", "type": "uint256"},
          {"internalType": "uint256", "name": "createdAt", "type": "uint256"},
          {"internalType": "address", "name": "creator", "type": "address"},
          {"internalType": "bool", "name": "isActive", "type": "bool"}
        ],
        "internalType": "struct MarketRegistry.MarketInfo[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "name": "markets",
    "outputs": [
      {"internalType": "address", "name": "marketAddress", "type": "address"},
      {"internalType": "address", "name": "functionsConsumer", "type": "address"},
      {"internalType": "string", "name": "question", "type": "string"},
      {"internalType": "uint256", "name": "closeTime", "type": "uint256"},
      {"internalType": "uint256", "name": "settleTime", "type": "uint256"},
      {"internalType": "uint256", "name": "createdAt", "type": "uint256"},
      {"internalType": "address", "name": "creator", "type": "address"},
      {"internalType": "bool", "name": "isActive", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_marketAddress", "type": "address"},
      {"internalType": "address", "name": "_functionsConsumer", "type": "address"},
      {"internalType": "string", "name": "_question", "type": "string"},
      {"internalType": "uint256", "name": "_closeTime", "type": "uint256"},
      {"internalType": "uint256", "name": "_settleTime", "type": "uint256"}
    ],
    "name": "registerMarket",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "name": "registeredMarkets",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "", "type": "address"},
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "name": "userMarkets",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  }
];

export const FUNCTIONS_CONSUMER_ABI = [
  {
    "inputs": [
      {"internalType": "address", "name": "router", "type": "address"},
      {"internalType": "address", "name": "_predictionMarket", "type": "address"}
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "inputs": [
      {"internalType": "bytes32", "name": "requestId", "type": "bytes32"}
    ],
    "name": "UnexpectedRequestID",
    "type": "error"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "bytes32", "name": "requestId", "type": "bytes32"},
      {"indexed": false, "internalType": "uint256", "name": "oddsYes", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "oddsNo", "type": "uint256"}
    ],
    "name": "OddsFulfilled",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "bytes32", "name": "requestId", "type": "bytes32"}
    ],
    "name": "OddsRequestSent",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "bytes32", "name": "requestId", "type": "bytes32"},
      {"indexed": false, "internalType": "bytes", "name": "error", "type": "bytes"}
    ],
    "name": "RequestFailed",
    "type": "event"
  },
  {
    "inputs": [],
    "name": "debugDecodeResponse",
    "outputs": [
      {"internalType": "string", "name": "decoded", "type": "string"},
      {"internalType": "uint256", "name": "oddsYes", "type": "uint256"},
      {"internalType": "uint256", "name": "oddsNo", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getLastError",
    "outputs": [
      {"internalType": "bytes", "name": "", "type": "bytes"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getLastResponse",
    "outputs": [
      {"internalType": "bytes", "name": "", "type": "bytes"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "latestRequestId",
    "outputs": [
      {"internalType": "bytes32", "name": "", "type": "bytes32"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint256", "name": "_oddsYes", "type": "uint256"},
      {"internalType": "uint256", "name": "_oddsNo", "type": "uint256"}
    ],
    "name": "manualSetOdds",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "string", "name": "oddsStr", "type": "string"}
    ],
    "name": "parseOddsString",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"},
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "pure",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "predictionMarket",
    "outputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "string[]", "name": "args", "type": "string[]"},
      {"internalType": "uint64", "name": "subscriptionId", "type": "uint64"},
      {"internalType": "uint32", "name": "gasLimit", "type": "uint32"}
    ],
    "name": "requestOdds",
    "outputs": [
      {"internalType": "bytes32", "name": "requestId", "type": "bytes32"}
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "s_lastError",
    "outputs": [
      {"internalType": "bytes", "name": "", "type": "bytes"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "s_lastResponse",
    "outputs": [
      {"internalType": "bytes", "name": "", "type": "bytes"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_pm", "type": "address"}
    ],
    "name": "setPredictionMarket",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
];

// Updated ABI for PredictionMarketExtended
export const PREDICTION_MARKET_ABI = [
  {
    "inputs": [
      {"internalType": "string", "name": "_question", "type": "string"},
      {"internalType": "uint256", "name": "_closeTime", "type": "uint256"},
      {"internalType": "uint256", "name": "_settleTime", "type": "uint256"},
      {"internalType": "address", "name": "_betToken", "type": "address"},
      {"internalType": "uint256", "name": "_tokenBetAmount", "type": "uint256"},
      {"internalType": "bool", "name": "_enableTokenBetting", "type": "bool"}
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "internalType": "address", "name": "user", "type": "address"},
      {"indexed": false, "internalType": "bool", "name": "prediction", "type": "bool"},
      {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"},
      {"indexed": false, "internalType": "bool", "name": "isToken", "type": "bool"},
      {"indexed": false, "internalType": "uint256", "name": "betIndex", "type": "uint256"}
    ],
    "name": "BetPlaced",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "uint256", "name": "timestamp", "type": "uint256"}
    ],
    "name": "MarketClosed",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "bool", "name": "result", "type": "bool"},
      {"indexed": false, "internalType": "uint256", "name": "timestamp", "type": "uint256"}
    ],
    "name": "MarketSettled",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "uint256", "name": "yes", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "no", "type": "uint256"}
    ],
    "name": "OddsUpdated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "internalType": "address", "name": "user", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "betIndex", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"},
      {"indexed": false, "internalType": "bool", "name": "isToken", "type": "bool"}
    ],
    "name": "WinningsClaimed",
    "type": "event"
  },
  {
    "inputs": [
      {"internalType": "bool", "name": "_result", "type": "bool"}
    ],
    "name": "autoSettle",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "", "type": "address"},
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "name": "bets",
    "outputs": [
      {"internalType": "uint256", "name": "amount", "type": "uint256"},
      {"internalType": "bool", "name": "prediction", "type": "bool"},
      {"internalType": "bool", "name": "claimed", "type": "bool"},
      {"internalType": "bool", "name": "isToken", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "betToken",
    "outputs": [
      {"internalType": "contract IERC20", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "checkAndCloseMarket",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "claimWinnings",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint256[]", "name": "betIndices", "type": "uint256[]"}
    ],
    "name": "claimWinningsByIndices",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "closeMarket",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "closeTime",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "bool", "name": "enabled", "type": "bool"}
    ],
    "name": "enableTokenBetting",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getMarketTimes",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"},
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getPoolTotals",
    "outputs": [
      {"internalType": "uint256", "name": "ethYes", "type": "uint256"},
      {"internalType": "uint256", "name": "ethNo", "type": "uint256"},
      {"internalType": "uint256", "name": "tokenYes", "type": "uint256"},
      {"internalType": "uint256", "name": "tokenNo", "type": "uint256"},
      {"internalType": "uint256", "name": "totalETH", "type": "uint256"},
      {"internalType": "uint256", "name": "totalToken", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getTimeRemaining",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"},
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "user", "type": "address"}
    ],
    "name": "getUserBetCount",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "user", "type": "address"},
      {"internalType": "uint256", "name": "index", "type": "uint256"}
    ],
    "name": "getUserBet",
    "outputs": [
      {"internalType": "uint256", "name": "amount", "type": "uint256"},
      {"internalType": "bool", "name": "prediction", "type": "bool"},
      {"internalType": "bool", "name": "claimed", "type": "bool"},
      {"internalType": "bool", "name": "isToken", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "user", "type": "address"}
    ],
    "name": "getUserBets",
    "outputs": [
      {
        "components": [
          {"internalType": "uint256", "name": "amount", "type": "uint256"},
          {"internalType": "bool", "name": "prediction", "type": "bool"},
          {"internalType": "bool", "name": "claimed", "type": "bool"},
          {"internalType": "bool", "name": "isToken", "type": "bool"}
        ],
        "internalType": "struct PredictionMarketExtended.Bet[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "user", "type": "address"}
    ],
    "name": "getUserUnclaimedWinnings",
    "outputs": [
      {"internalType": "uint256", "name": "ethWinnings", "type": "uint256"},
      {"internalType": "uint256", "name": "tokenWinnings", "type": "uint256"},
      {"internalType": "uint256[]", "name": "winningBetIndices", "type": "uint256[]"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "isMarketExpired",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "isSettlementTime",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "oddsNo",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "oddsUpdater",
    "outputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "oddsYes",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "bool", "name": "prediction", "type": "bool"}
    ],
    "name": "placeBet",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "user", "type": "address"},
      {"internalType": "bool", "name": "prediction", "type": "bool"},
      {"internalType": "uint256", "name": "amount", "type": "uint256"}
    ],
    "name": "placeBetWithToken",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "question",
    "outputs": [
      {"internalType": "string", "name": "", "type": "string"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "result",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_token", "type": "address"},
      {"internalType": "uint256", "name": "_amount", "type": "uint256"}
    ],
    "name": "setBetToken",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint256", "name": "_oddsYes", "type": "uint256"},
      {"internalType": "uint256", "name": "_oddsNo", "type": "uint256"}
    ],
    "name": "setOdds",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "_updater", "type": "address"}
    ],
    "name": "setOddsUpdater",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "bool", "name": "_result", "type": "bool"}
    ],
    "name": "setResult",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "settleTime",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "status",
    "outputs": [
      {"internalType": "enum PredictionMarketExtended.MarketStatus", "name": "", "type": "uint8"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "tokenBetAmount",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "tokenBettingEnabled",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalNoETH",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalNoToken",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalYesETH",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalYesToken",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  }
];

// ERC20 ABI for token operations (CCIP-BnM and other tokens)
export const ERC20_ABI = [
  {
    "inputs": [
      {"internalType": "address", "name": "spender", "type": "address"},
      {"internalType": "uint256", "name": "amount", "type": "uint256"}
    ],
    "name": "approve",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "account", "type": "address"}
    ],
    "name": "balanceOf",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "decimals",
    "outputs": [
      {"internalType": "uint8", "name": "", "type": "uint8"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "name",
    "outputs": [
      {"internalType": "string", "name": "", "type": "string"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "symbol",
    "outputs": [
      {"internalType": "string", "name": "", "type": "string"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalSupply",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "to", "type": "address"},
      {"internalType": "uint256", "name": "amount", "type": "uint256"}
    ],
    "name": "transfer",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "from", "type": "address"},
      {"internalType": "address", "name": "to", "type": "address"},
      {"internalType": "uint256", "name": "amount", "type": "uint256"}
    ],
    "name": "transferFrom",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "owner", "type": "address"},
      {"internalType": "address", "name": "spender", "type": "address"}
    ],
    "name": "allowance",
    "outputs": [
      {"internalType": "uint256", "name": "", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "to", "type": "address"}
    ],
    "name": "drip",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
];

// CCIP Bridge ABI for cross-chain betting
export const CCIP_BRIDGE_ABI = [
  {
    "inputs": [
      {"internalType": "address", "name": "_router", "type": "address"},
      {"internalType": "address", "name": "_link", "type": "address"},
      {"internalType": "address", "name": "_predictionMarket", "type": "address"},
      {"internalType": "uint64[]", "name": "_allowedChains", "type": "uint64[]"},
      {"internalType": "address[]", "name": "_allowedSenders", "type": "address[]"}
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "bytes32", "name": "messageId", "type": "bytes32"},
      {"indexed": false, "internalType": "string", "name": "eventId", "type": "string"},
      {"indexed": false, "internalType": "bool", "name": "prediction", "type": "bool"},
      {"indexed": false, "internalType": "address", "name": "token", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"},
      {"indexed": false, "internalType": "address", "name": "bettor", "type": "address"}
    ],
    "name": "BetReceived",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "internalType": "bytes32", "name": "messageId", "type": "bytes32"},
      {"indexed": false, "internalType": "uint64", "name": "destinationChainSelector", "type": "uint64"},
      {"indexed": false, "internalType": "address", "name": "receiver", "type": "address"},
      {"indexed": false, "internalType": "string", "name": "eventId", "type": "string"},
      {"indexed": false, "internalType": "bool", "name": "prediction", "type": "bool"},
      {"indexed": false, "internalType": "address", "name": "token", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"},
      {"indexed": false, "internalType": "address", "name": "feeToken", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "fee", "type": "uint256"}
    ],
    "name": "BetMessageSent",
    "type": "event"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "sender", "type": "address"},
      {"internalType": "bool", "name": "allowed", "type": "bool"}
    ],
    "name": "allowSender",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint64", "name": "selector", "type": "uint64"},
      {"internalType": "bool", "name": "allowed", "type": "bool"}
    ],
    "name": "allowSourceChain",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "name": "allowlistedSenders",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint64", "name": "", "type": "uint64"}
    ],
    "name": "allowlistedSourceChains",
    "outputs": [
      {"internalType": "bool", "name": "", "type": "bool"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "linkToken",
    "outputs": [
      {"internalType": "contract IERC20", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {"internalType": "address", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "predictionMarket",
    "outputs": [
      {"internalType": "contract IPredictionMarketBet", "name": "", "type": "address"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "uint64", "name": "destinationChainSelector", "type": "uint64"},
      {"internalType": "address", "name": "receiver", "type": "address"},
      {"internalType": "string", "name": "eventId", "type": "string"},
      {"internalType": "bool", "name": "prediction", "type": "bool"},
      {"internalType": "address", "name": "token", "type": "address"},
      {"internalType": "uint256", "name": "amount", "type": "uint256"},
      {"internalType": "bool", "name": "payWithLINK", "type": "bool"}
    ],
    "name": "sendBetWithToken",
    "outputs": [
      {"internalType": "bytes32", "name": "messageId", "type": "bytes32"}
    ],
    "stateMutability": "payable",
    "type": "function"
  }
];


export const MARKET_CCIP_CONFIG = {
  "0xABc561948eFE3e0E37D28B483a5B9d1EbC8f0899": { // Market 1
    fujiBridge: "0xd716B3F4B000c4Dc8Adb5Da8491a32A99c97108A",
    sepoliaReceiver: "0xd0825e9db3BBb1Fd142C5cFCEE5b362F043f6260",
    predictionMarket: "0xABc561948eFE3e0E37D28B483a5B9d1EbC8f0899"
  },
  "0x252CA7101A32619CE922455f0a84AFdDE41F0843": { // Market 2
    fujiBridge: "0x44C3A367124DEAC71CaeB5c583bd3525A65A4BB0",
    sepoliaReceiver: "0x87EE37a5D3D266Ab6391FA95f1aCd3E532201412",
    predictionMarket: "0x252CA7101A32619CE922455f0a84AFdDE41F0843"
  },
  "0x119bdDbEb434447A0045966F23364E53EbE45270": { // Market 3
    fujiBridge: "0xc5DeC4B6E97f5A6Ce3451E62a3BC7126322AF8AF",
    sepoliaReceiver: "0x33C3b22998f74B1AA9ba016389D24cA6E6f437F3",
    predictionMarket: "0x119bdDbEb434447A0045966F23364E53EbE45270"
  },
  "0xCB287A95e48D651Ee5793E02A6923502B34cA7e6": { // Market 4
    fujiBridge: "0x239Ba0fc64B183E0b3cbaccBC4689E31F0001771",
    sepoliaReceiver: "0x02109d94B50662710f1DB07fB27A4D5BfD68b84d",
    predictionMarket: "0xCB287A95e48D651Ee5793E02A6923502B34cA7e6"
  }
};