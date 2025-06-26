import React, { useState, useEffect } from 'react';
import { 
  AlertCircle, 
  TrendingUp, 
  TrendingDown, 
  DollarSign, 
  Users, 
  Clock,
  CheckCircle,
  XCircle,
  Loader2,
  RefreshCw,
  Wallet,
  Trophy,
  BarChart3,
  ArrowLeft
} from 'lucide-react';
import { BrowserProvider } from 'ethers';

// Import components
import MarketList from './MarketList';
import CountdownTimer from './CountdonwTimer.jsx';

// Import configuration and utilities
import { CONFIG, MARKET_CCIP_CONFIG } from './config';
import {
  connectWallet,
  createMarketContract,
  createFunctionsContract,
  fetchMarketData,
  requestOddsUpdate,
  placeBet,
  claimWinnings,
  calculatePotentialPayout,
  pollForOddsUpdate,
  formatTimeSince,
  checkAndCloseMarket,
  placeBetBasedOnChain, // this uses ETH or CCIP as needed
  getCurrentChainId,
  getChainCurrency,
  validateBetAmount,
  checkSufficientBalance,
  createCCIPBridgeContract
} from './contractUtils';

// Import CSS
import './App.css';

function App() {
  // State variables
  const [account, setAccount] = useState("");
  const [provider, setProvider] = useState(null);
  const [marketRegistry, setMarketRegistry] = useState(null);
  const [currentMarketContract, setCurrentMarketContract] = useState(null);
  const [currentFunctionsContract, setCurrentFunctionsContract] = useState(null);
  const [currentView, setCurrentView] = useState('marketList'); // 'marketList' or 'marketDetail'
  const [selectedMarketData, setSelectedMarketData] = useState(null);
  
  // Market data
  const [question, setQuestion] = useState("");
  const [oddsYes, setOddsYes] = useState(50);
  const [oddsNo, setOddsNo] = useState(50);
  const [status, setStatus] = useState("Open");
  const [totalYes, setTotalYes] = useState("0");
  const [totalNo, setTotalNo] = useState("0");
  const [marketCloseTime, setMarketCloseTime] = useState(null);
  const [marketSettleTime, setMarketSettleTime] = useState(null);

  const [totalYesETH, setTotalYesETH] = useState("0");
  const [totalNoETH, setTotalNoETH] = useState("0");
  const [totalYesToken, setTotalYesToken] = useState("0");
  const [totalNoToken, setTotalNoToken] = useState("0");
  const [tokenBettingEnabled, setTokenBettingEnabled] = useState(false);
  const [chainId, setChainId] = useState(null);

  
  // User interaction
  const [selectedBet, setSelectedBet] = useState(null);
  const [betAmount, setBetAmount] = useState("");
  const [userBets, setUserBet] = useState(null);
  const [loading, setLoading] = useState(false);
  const [updatingOdds, setUpdatingOdds] = useState(false);
  const [lastUpdate, setLastUpdate] = useState(Date.now());

  useEffect(() => {
    if (window.ethereum) {
      window.ethereum.on('chainChanged', (chainIdHex) => {
        setChainId(parseInt(chainIdHex, 16));
      });
    }
    return () => {
      if (window.ethereum?.removeListener) {
        window.ethereum.removeListener('chainChanged', () => {});
      }
    };
  }, []);

  // Connect wallet handler
  const handleConnectWallet = async () => {
    try {
      setLoading(true);
      
      const walletData = await connectWallet();
      
      setAccount(walletData.account);
      setProvider(walletData.provider);
      setMarketRegistry(walletData.marketRegistry);

      // Listen for account changes
      window.ethereum.on('accountsChanged', (accounts) => {
        if (accounts.length > 0) {
          setAccount(accounts[0]);
          if (currentView === 'marketDetail') {
            handleFetchData();
          }
        } else {
          setAccount("");
          setProvider(null);
          setMarketRegistry(null);
          setCurrentView('marketList');
        }
      });
      
    } catch (error) {
      console.error("Connection error:", error);
      alert("Failed to connect: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  // Handle market selection from list
  const handleSelectMarket = async (marketAddress, marketData) => {
    try {
      setLoading(true);
      
      // Create contract instances for the selected market
      const marketContract = await createMarketContract(marketAddress, provider);
      const functionsContract = await createFunctionsContract(marketData.functionsConsumer, provider);
      
      setCurrentMarketContract(marketContract);
      setCurrentFunctionsContract(functionsContract);
      setSelectedMarketData(marketData);
      setCurrentView('marketDetail');
      
      // Set up event listener for OddsUpdated
      marketContract.on("OddsUpdated", (yes, no) => {
        console.log("OddsUpdated event:", yes.toString(), no.toString());
        setOddsYes(parseInt(yes.toString()));
        setOddsNo(parseInt(no.toString()));
        setLastUpdate(Date.now());
      });
      
      // Set market times from contract data
      setMarketCloseTime(marketData.closeTime);
      setMarketSettleTime(marketData.settleTime);
      
    } catch (error) {
      console.error("Error selecting market:", error);
      alert("Failed to load market: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  // Handle countdown timer expiration for close time
  const handleMarketCloseTime = async () => {
    console.log("Market close time reached, checking if market needs closing...");
    
    if (currentMarketContract && status === 'Open') {
      try {
        await checkAndCloseMarket(currentMarketContract);
        setStatus("Closed");
        console.log("Market closed successfully");
      } catch (error) {
        console.error("Error closing market:", error);
      }
    }
  };

  // Handle settlement time (in production, this would be automated)
  const handleSettlementTime = async () => {
    console.log("Settlement time reached - market can now be settled");
    // In production, this would trigger automatic settlement via Chainlink Automation
    // For now, just log that settlement is possible
  };

  // Go back to market list
  const handleBackToList = () => {
    setCurrentView('marketList');
    setSelectedMarketData(null);
    setCurrentMarketContract(null);
    setCurrentFunctionsContract(null);
    setSelectedBet(null);
    setBetAmount("");
    
    // Clean up event listeners
    if (currentMarketContract) {
      currentMarketContract.removeAllListeners("OddsUpdated");
    }
  };

  // Fetch market data handler
  const handleFetchData = async () => {
    if (!currentMarketContract) return;
    
    try {
      const { marketData, userBets } = await fetchMarketData(
        currentMarketContract,
        account
      );
      
      setQuestion(marketData.question);
      setOddsYes(marketData.oddsYes);
      setOddsNo(marketData.oddsNo);
      setStatus(marketData.status);
      // setTotalYes(marketData.totalYes);
      // setTotalNo(marketData.totalNo);
      setUserBet(userBets);
      setLastUpdate(Date.now());
      
      setTotalYesETH(marketData.totalYesETH);
      setTotalNoETH(marketData.totalNoETH);
      setTotalYesToken(marketData.totalYesToken);
      setTotalNoToken(marketData.totalNoToken);
      setTokenBettingEnabled(marketData.tokenBettingEnabled);

      console.log('user bet in main:', userBets)
      // Set these when you fetch marketData
      setTotalYesETH(marketData.totalYesETH);
      setTotalNoETH(marketData.totalNoETH);
      setTotalYesToken(marketData.totalYesToken);
      setTotalNoToken(marketData.totalNoToken);
      setTokenBettingEnabled(marketData.tokenBettingEnabled);

      
      // Update timing data from contract
      setMarketCloseTime(marketData.closeTime);
      setMarketSettleTime(marketData.settleTime);
      
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  // Update odds handler
  const handleUpdateOdds = async () => {
    if (!currentFunctionsContract) return;

    try {
      setUpdatingOdds(true);

      await requestOddsUpdate(currentFunctionsContract);
      
      // Poll for updates
      await pollForOddsUpdate(handleFetchData);
      
    } catch (error) {
      console.error("Odds update error:", error);
      alert("Failed to update odds: " + error.message);
    } finally {
      setUpdatingOdds(false);
    }
  };

  // Place bet handler
  const handlePlaceBet = async () => {
  if (!currentMarketContract || !account || selectedBet === null || !betAmount) return;

  try {
    setLoading(true);

    const chainId = await getCurrentChainId();

    // Validate and check balance for this chain/currency
    const { valid, error } = validateBetAmount(betAmount, chainId);
    if (!valid) {
      alert(error);
      setLoading(false);
      return;
    }

    const freshProvider = new BrowserProvider(window.ethereum);
    const balanceCheck = await checkSufficientBalance(betAmount, chainId, account, freshProvider);
    console.log('account:', account)
    console.log('provider:', freshProvider)
    console.log('balanceCheck:', balanceCheck)
    if (!balanceCheck.sufficient) {
      alert(`Insufficient ${balanceCheck.currency} balance!`);
      setLoading(false);
      return;
    }

    // Instantiate bridge only for Fuji chain
    let ccipBridge = null;
    let fujiBridgeAddress = null;
    if (chainId === CONFIG.CHAINS.FUJI.chainId) {
      // Find the correct bridge for this market
      fujiBridgeAddress = MARKET_CCIP_CONFIG[selectedMarketData.address]?.fujiBridge;
      if (!fujiBridgeAddress) {
        alert("No Fuji CCIP bridge configured for this market.");
        setLoading(false);
        return;
      }
      ccipBridge = await createCCIPBridgeContract(fujiBridgeAddress, freshProvider);
    }

    console.log('currentMarketContract: ',currentMarketContract)
    console.log('selectedMarketData.address: ',selectedMarketData.address)

    await placeBetBasedOnChain(
      chainId,
      currentMarketContract,
      ccipBridge,
      selectedMarketData.address,
      selectedBet === 'yes',
      betAmount,
      freshProvider
    );

    setSelectedBet(null);
    setBetAmount("");
    await handleFetchData();

  } catch (error) {
    console.error("Bet error:", error);
    alert("Bet failed: " + error.message);
  } finally {
    setLoading(false);
  }
};


  // Claim winnings handler
  const handleClaimWinnings = async () => {
    if (!currentMarketContract || !account) return;
    
    try {
      setLoading(true);
      
      await claimWinnings(currentMarketContract);
      
      alert("Winnings claimed successfully!");
      await handleFetchData();
      
    } catch (error) {
      console.error("Claim error:", error);
      alert("Claim failed: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  // Calculate payout
  const getPotentialPayout = () => {
    return calculatePotentialPayout(betAmount, selectedBet, totalYes, totalNo);
  };

  // Auto-refresh effect for market detail view
  useEffect(() => {
    if (currentMarketContract && currentView === 'marketDetail') {
      handleFetchData();
      const interval = setInterval(handleFetchData, CONFIG.REFRESH_INTERVAL);
      return () => clearInterval(interval);
    }
  }, [currentMarketContract, account, currentView]);

  // Cleanup effect
  useEffect(() => {
    return () => {
      if (currentMarketContract) {
        currentMarketContract.removeAllListeners("OddsUpdated");
      }
    };
  }, [currentMarketContract]);

  // Render welcome screen if wallet not connected
    if (!account) {
    return (
      <div className="app-container">
        <header className="header">
          <div className="header-content">
            <div className="logo">
              <BarChart3 size={32} color="#a78bfa" />
              <h1 className="logo-text">CryptoBet</h1>
            </div>
            <button
              onClick={handleConnectWallet}
              disabled={loading}
              className="btn btn-primary"
            >
              {loading ? <Loader2 size={16} className="loading-spinner" /> : <Wallet size={16} />}
              Connect Wallet
            </button>
          </div>
        </header>

        <main className="main">
          <div className="welcome-container">
            <div className="welcome-content">
              <BarChart3 size={64} color="#a78bfa" className="welcome-icon" />
              <h2 className="welcome-title">Welcome to CryptoBet</h2>
              <p className="welcome-text">Connect your wallet to start betting on future events with ML-powered odds</p>
              <button
                onClick={handleConnectWallet}
                disabled={loading}
                className="btn btn-primary"
              >
                {loading ? <Loader2 size={20} className="loading-spinner" /> : <Wallet size={20} />}
                Connect Wallet
              </button>
            </div>
          </div>
        </main>
      </div>
    );
  }

  const hasUserBet = userBets && Array.isArray(userBets) && userBets.length > 0;
  return (
    <div className="app-container">
      {/* Header always visible */}
      <header className="header">
        <div className="header-content">
          <div className="logo">
            <BarChart3 size={32} color="#a78bfa" />
            <h1 className="logo-text">CryptoBet</h1>
          </div>
          <div className="account-display">
            {/* <button
              onClick={handleUpdateOdds}
              disabled={updatingOdds || status !== 'Open'}
              className="btn-icon"
              title="Update odds via Chainlink"
            >
              <RefreshCw size={20} className={updatingOdds ? "loading-spinner" : ""} />
              <p className="account-label">Update Odds</p>
            </button> */}
            <div className="account-info">
              <p className="account-label">Connected</p>
              <p className="account-address">
                {account.slice(0, 6)}...{account.slice(-4)}
              </p>
            </div>
          </div>
        </div>
      </header>

      {/* Market list always rendered in background */}
      <main className={`main ${currentView === 'marketDetail' ? 'dimmed' : ''}`}>
        <MarketList
          onSelectMarket={handleSelectMarket}
          connectedAccount={account}
          marketRegistry={marketRegistry}
          provider={provider}
        />
      </main>

      {/* Slide-in Market Detail Panel */}
      <div className={`detail-slide-panel ${currentView === 'marketDetail' ? 'open' : ''}`}>
        <header className="header" style={{ position: 'sticky', top: 0, background: 'rgba(0,0,0,0.4)', zIndex: 100 }}>
          <div className="header-content">
            <div className="logo">
              <button
                onClick={handleBackToList}
                className="btn-icon"
                title="Back to market list"
              >
                <ArrowLeft size={20} />
              </button>
              <BarChart3 size={32} color="#a78bfa" />
              <h1 className="logo-text">Market Detail</h1>
            </div>
          </div>
        </header>

        {/* Market detail content from your original "marketDetail" view */}
        <main className="main">
        {/* Countdown Timers */}
        {marketCloseTime && status === 'Open' && (
          <CountdownTimer 
            endTime={marketCloseTime}
            onTimeUp={handleMarketCloseTime}
            status={status}
          />
        )}

        {marketSettleTime && status === 'Closed' && (
          <div className="card">
            <div className="card-header">
              <div>
                <h3 className="market-question">Settlement Countdown</h3>
                <p className="update-time">Market closed - awaiting settlement</p>
              </div>
              <Clock size={32} color="#eab308" />
            </div>
            <CountdownTimer 
              endTime={marketSettleTime}
              onTimeUp={handleSettlementTime}
              status={status}
            />
          </div>
        )}

        {/* Market Question and Stats */}
        <div className="card">
          <div className="card-header">
            <div>
              <h2 className="market-question">{question || "Loading..."}</h2>
              <div className="market-meta">
                <span className={`status-badge status-${status.toLowerCase()}`}>
                  {status === 'Open' && <Clock size={12} />}
                  {status === 'Closed' && <XCircle size={12} />}
                  {status === 'Settled' && <CheckCircle size={12} />}
                  {status}
                </span>
                <span className="update-time">
                  Last update: {formatTimeSince(lastUpdate)}
                </span>
              </div>
            </div>
            <Trophy size={48} color="#eab308" />
          </div>

          <div className="stats-grid">
            
            <div className="stat-card">
              <p className="stat-label">ETH Pool (Yes/No)</p>
              <p className="stat-value">{totalYesETH} / {totalNoETH} ETH</p>
              </div>
              {tokenBettingEnabled && (
                <div className="stat-card">
                  <p className="stat-label">Token Pool (Yes/No)</p>
                  <p className="stat-value">{totalYesToken} / {totalNoToken} CCIP-BnM</p>
                </div>
              )}
            <div className="stat-card">
              <p className="stat-label">Participants</p>
              <p className="stat-value stat-value-flex">
                <Users size={20} color="#a78bfa" />
                {Math.floor(Math.random() * 100) + 50}
              </p>
            </div>
            <div className="stat-card">
  <p className="stat-label">Your Position</p>
  <div className="stat-value" style={{ fontSize: "0.93em", lineHeight: 1.2 }}>
    {userBets && userBets.length > 0 ? (
      userBets.map((bet, idx) => (
        <span key={bet.index} style={{ marginRight: "0.6em", whiteSpace: "nowrap", display: "inline-block" }}>
          {bet.amount} {bet.asset} 
          <span style={{ color: bet.prediction ? "#22c55e" : "#ef4444", fontWeight: 600, marginLeft: 2 }}>
            {bet.prediction ? "YES" : "NO"}
          </span>
          {bet.claimed && (
            <span style={{ color: "#a78bfa", marginLeft: 2 }}>(claimed)</span>
          )}
        </span>
      ))
    ) : (
      <span style={{ color: "#888" }}>No bet</span>
    )}
  </div>
</div>

            {selectedMarketData && (
              <div className="stat-card">
                <p className="stat-label">Market Address</p>
                <p className="stat-value" style={{ fontSize: '0.875rem', fontFamily: 'monospace' }}>
                  {selectedMarketData.address.slice(0, 8)}...
                </p>
              </div>
            )}
          </div>
        </div>

        {/* Odds Cards */}
        <div className="odds-grid">
          {/* YES */}
          <div 
            className={`odds-card odds-card-yes ${selectedBet === 'yes' ? 'selected' : ''} ${status !== 'Open' || hasUserBet ? 'disabled' : ''}`}
            onClick={() => status === 'Open' && !hasUserBet && setSelectedBet('yes')}
          >
            <div className="odds-header">
              <div>
                <h3 className="odds-title">
                  <CheckCircle size={32} color="#22c55e" /> YES
                </h3>
                <p className="odds-percentage odds-percentage-yes">{oddsYes}%</p>
              </div>
              <TrendingUp size={32} color="#22c55e" />
            </div>
            <div className="odds-details">
              <div className="odds-detail-row">
                <span className="odds-detail-label">Pool Size</span>
                <span className="odds-detail-value">{totalYes} ETH</span>
              </div>
              <div className="odds-detail-row">
                <span className="odds-detail-label">Implied Probability</span>
                <span className="odds-detail-value">{oddsYes}%</span>
              </div>
              {selectedBet === 'yes' && betAmount && (
                <div className="odds-detail-row odds-detail-payout">
                  <span className="odds-detail-label">Potential Payout</span>
                  <span className="odds-detail-value odds-detail-payout-yes">
                    {getPotentialPayout()} ETH
                  </span>
                </div>
              )}
            </div>
          </div>

          {/* NO */}
          <div 
            className={`odds-card odds-card-no ${selectedBet === 'no' ? 'selected' : ''} ${status !== 'Open' || hasUserBet ? 'disabled' : ''}`}
            onClick={() => status === 'Open' && !hasUserBet && setSelectedBet('no')}
          >
            <div className="odds-header">
              <div>
                <h3 className="odds-title">
                  <XCircle size={32} color="#ef4444" /> NO
                </h3>
                <p className="odds-percentage odds-percentage-no">{oddsNo}%</p>
              </div>
              <TrendingDown size={32} color="#ef4444" />
            </div>
            <div className="odds-details">
              <div className="odds-detail-row">
                <span className="odds-detail-label">Pool Size</span>
                <span className="odds-detail-value">{totalNo} ETH</span>
              </div>
              <div className="odds-detail-row">
                <span className="odds-detail-label">Implied Probability</span>
                <span className="odds-detail-value">{oddsNo}%</span>
              </div>
              {selectedBet === 'no' && betAmount && (
                <div className="odds-detail-row odds-detail-payout">
                  <span className="odds-detail-label">Potential Payout</span>
                  <span className="odds-detail-value odds-detail-payout-no">
                    {getPotentialPayout()} ETH
                  </span>
                </div>
              )}
            </div>
          </div>
        </div>

        {/* Betting Interface */}
        {status === 'Open' && !hasUserBet && (
          <div className="card">
            <h3 className="odds-title">Place Your Bet</h3>
            {selectedBet ? (
              <div className="bet-form">
                <div className="input-group">
                  <label className="input-label">Bet Amount ({getChainCurrency(chainId)})</label>
                  <div className="input-wrapper">
                    <DollarSign size={20} className="input-icon" />
                    <input
                      type="number"
                      step="0.01"
                      min="0.01"
                      value={betAmount}
                      onChange={(e) => setBetAmount(e.target.value)}
                      placeholder="0.00"
                      className="input"
                    />
                  </div>
                </div>
                <div className="button-group">
                  <button
                    onClick={() => {setSelectedBet(null); setBetAmount('');}}
                    className="btn btn-secondary"
                  >
                    Cancel
                  </button>
                  <button
                    onClick={handlePlaceBet}
                    disabled={loading || !betAmount || parseFloat(betAmount) <= 0}
                    className="btn btn-primary"
                  >
                    {loading ? <Loader2 size={16} className="loading-spinner" /> : null}
                    Place Bet on {selectedBet.toUpperCase()}
                  </button>
                </div>
              </div>
            ) : (
              <p className="empty-state">
                Select YES or NO above to place your bet
              </p>
            )}
          </div>
        )}

        {/* User Position */}
        {userBets && userBets.length > 0 && (
  <div className="card position-card">
    <h3 className="position-header">
      <Trophy size={24} color="#eab308" /> Your Positions
    </h3>
    <div className="position-grid">
      {userBets.map((bet, idx) => (
        <div key={bet.index} className="position-item" style={{minWidth: 120}}>
          <p className="position-label">Your Bet</p>
          <p className="position-value">{bet.amount} {bet.asset}</p>
          <p className="position-label">Your Prediction</p>
          <p className="position-value">
            {bet.prediction
              ? <><CheckCircle size={18} color="#22c55e" /> YES</>
              : <><XCircle size={18} color="#ef4444" /> NO</>}
          </p>
          <p className="position-label">Status</p>
          <p className="position-value">
            {bet.claimed ? 'Claimed'
              : status === 'Settled' ? 'Ready to Claim' : 'Active'}
          </p>
          {status === 'Settled' && !bet.claimed && (
            <button
              onClick={handleClaimWinnings}
              disabled={loading}
              className="btn btn-primary btn-claim"
              style={{marginTop: 3, fontSize: '0.9em'}}
            >
              {loading ? <Loader2 size={16} className="loading-spinner" /> : <DollarSign size={16} />}
              Claim
            </button>
          )}
        </div>
      ))}
    </div>
  </div>
)}


        {/* Info Box */}
        <div className="info-box">
          <AlertCircle size={20} color="#3b82f6" className="info-box-icon" />
          <div className="info-box-content">
            <p className="info-box-title">Powered by Chainlink Functions & Registry</p>
            <p>
              Odds are calculated by ML models and updated via Chainlink's decentralized oracle network.
              Markets are stored in an on-chain registry with automated timing controls.
              {updatingOdds && " Updating odds now..."}
            </p>
          </div>
        </div>
      </main>

      </div>

      {/* Clickable backdrop */}
      <div
        className={`detail-overlay ${currentView === 'marketDetail' ? '' : 'hidden'}`}
        onClick={handleBackToList}
      />
    </div>
  );

}

export default App;