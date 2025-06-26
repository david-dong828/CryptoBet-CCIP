import React, { useState, useEffect } from 'react';
import { createFunctionsContract, requestOddsUpdate } from './contractUtils';
import { 
  Clock,
  TrendingUp,
  Users,
  DollarSign,
  ChevronRight,
  Loader2,
  Plus,
  BarChart3,
  AlertCircle,
  Activity,
  Zap,
  Trophy,
  Target,
  Search
} from 'lucide-react';
import ReactCountryFlag from "react-country-flag";
import { fetchAllMarkets, fetchMarketStats, getCountdownText } from './contractUtils';
import './MarketList.css';

// Function to extract countries from question in order
const getCountriesFromQuestion = (question) => {
  const countryMap = {
    'USA': 'US',
    'United States': 'US',
    'America': 'US',
    'Canada': 'CA',
    'UK': 'GB',
    'United Kingdom': 'GB',
    'England': 'GB',
    'Britain': 'GB',
    'Japan': 'JP',
    'Germany': 'DE',
    'France': 'FR',
    'Brazil': 'BR',
    'India': 'IN',
    'China': 'CN',
    'Australia': 'AU',
    'Spain': 'ES',
    'Italy': 'IT',
    'Netherlands': 'NL',
    'Mexico': 'MX',
    'Argentina': 'AR',
    'Portugal': 'PT',
    'Russia': 'RU',
    'South Korea': 'KR',
    'Korea': 'KR'
  };
  
  // Find all country mentions and their positions
  const foundCountries = [];
  
  for (const [country, code] of Object.entries(countryMap)) {
    const regex = new RegExp(`\\b${country}\\b`, 'i');
    const match = question.match(regex);
    if (match) {
      foundCountries.push({
        country,
        code,
        index: match.index
      });
    }
  }
  
  // Sort by position in the question to maintain order
  foundCountries.sort((a, b) => a.index - b.index);
  
  // Return the first two countries found
  if (foundCountries.length >= 2) {
    return [foundCountries[0].code, foundCountries[1].code];
  } else if (foundCountries.length === 1) {
    return [foundCountries[0].code, 'UN']; // Default second flag
  }
  
  return ['UN', 'UN']; // Default flags if no countries found
};

function MarketList({ onSelectMarket, connectedAccount, marketRegistry, provider }) {
  const [markets, setMarkets] = useState([]);
  const [marketStats, setMarketStats] = useState({});
  const [loading, setLoading] = useState(true);
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [searchTerm, setSearchTerm] = useState('');
  const [error, setError] = useState(null);

  // Mock data for KPI cards
  const kpiData = {
    totalVolume: "2,847.35",
    activeTraders: "14,329",
    avgWinRate: "67.8",
    totalMarkets: markets.length || 42
  };

  // Fetch markets from registry
  const loadMarkets = async () => {
    if (!marketRegistry || !provider) {
      console.log("Missing dependencies:", { marketRegistry: !!marketRegistry, provider: !!provider });
      return;
    }
    
    try {
      setLoading(true);
      setError(null);
      
      console.log("Fetching markets from registry...");
      const registryMarkets = await fetchAllMarkets(marketRegistry);
      console.log("Fetched markets:", registryMarkets);
      
      setMarkets(registryMarkets);
      
      if (registryMarkets.length === 0) {
        console.log("No markets found in registry");
        setLoading(false);
        return;
      }
      
      // Fetch stats for each market
      console.log("Fetching stats for markets...");
      const statsPromises = registryMarkets.map(market => 
        fetchMarketStats(market.address, provider)
      );
      
      const stats = await Promise.all(statsPromises);
      const statsMap = {};
      
      stats.forEach((stat, index) => {
        if (stat) {
          statsMap[registryMarkets[index].address] = stat;
        }
      });
      
      console.log("Market stats:", statsMap);
      setMarketStats(statsMap);
      
    } catch (error) {
      console.error("Error loading markets:", error);
      setError(`Failed to load markets: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  // Load markets on component mount and when dependencies change
  useEffect(() => {
    loadMarkets();
  }, [marketRegistry, provider]);

  // Auto-refresh market stats every 30 seconds
  useEffect(() => {
    if (!markets.length) return;
    
    const interval = setInterval(async () => {
      try {
        const statsPromises = markets.map(market => 
          fetchMarketStats(market.address, provider)
        );
        
        const stats = await Promise.all(statsPromises);
        const statsMap = {};
        
        stats.forEach((stat, index) => {
          if (stat) {
            statsMap[markets[index].address] = stat;
          }
        });
        
        setMarketStats(statsMap);
      } catch (error) {
        console.error("Error refreshing market stats:", error);
      }
    }, 30000); // 30 seconds
    
    return () => clearInterval(interval);
  }, [markets, provider]);

  // Get market status based on countdown
  const getMarketStatus = (market) => {
    const stats = marketStats[market.address];
    if (!stats) return 'Loading';
    
    const now = Date.now();
    
    if (stats.status === 'Settled') return 'Settled';
    if (stats.status === 'Closed') return 'Closed';
    if (now >= market.closeTime) return 'Expired';
    return 'Open';
  };

  // Get countdown for market
  const getMarketCountdown = (market) => {
    const stats = marketStats[market.address];
    if (!stats) return 'Loading...';
    
    const now = Date.now();
    const status = getMarketStatus(market);
    
    if (status === 'Settled') return 'Settled';
    if (status === 'Closed') return 'Awaiting Settlement';
    if (status === 'Expired') return 'Expired';
    
    const timeRemaining = market.closeTime - now;
    return getCountdownText(timeRemaining);
  };

  // Filter markets based on search and category
  const filteredMarkets = markets.filter(market => {
    const matchesSearch = market.question.toLowerCase().includes(searchTerm.toLowerCase());
    const status = getMarketStatus(market);
    
    let matchesCategory = true;
    if (selectedCategory === 'active') {
      matchesCategory = status === 'Open';
    } else if (selectedCategory === 'ended') {
      matchesCategory = status === 'Closed' || status === 'Settled' || status === 'Expired';
    }
    
    return matchesSearch && matchesCategory && market.isActive;
  });

  const handleMarketClick = (market) => {
    onSelectMarket(market.address, market);
  };

  const handleCreateMarket = () => {
    alert("Market creation coming soon!");
  };

  if (loading && markets.length === 0) {
    return (
      <div className="market-list-container">
        <div className="loading-container">
          <Loader2 size={48} className="loading-spinner" />
          <p>Loading markets from blockchain...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="market-list-container">
        <div className="info-box">
          <AlertCircle size={20} color="#ef4444" className="info-box-icon" />
          <div className="info-box-content">
            <p className="info-box-title">Error Loading Markets</p>
            <p>{error}</p>
            <button 
              onClick={loadMarkets}
              className="btn btn-primary"
              style={{ marginTop: '1rem' }}
            >
              Try Again
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="market-list-container">
      {/* Header */}
      <div className="market-list-header">
        <div className="market-list-title">
          <BarChart3 size={32} color="#a78bfa" />
          <h1>Crypto Bet - Football</h1>
        </div>
        <button className="btn-create-market" onClick={handleCreateMarket}>
          <Plus size={20} />
          Create Market
        </button>
      </div>

      {/* Search Bar */}
      <div className="search-section">
        <div className="search-wrapper">
          <Search size={20} className="search-icon" />
          <input
            type="text"
            placeholder="Search markets..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="search-input-enhanced"
          />
        </div>
        <div className="category-tabs">
          <button 
            className={`category-tab ${selectedCategory === 'all' ? 'active' : ''}`}
            onClick={() => setSelectedCategory('all')}
          >
            All Markets ({markets.length})
          </button>
          <button 
            className={`category-tab ${selectedCategory === 'active' ? 'active' : ''}`}
            onClick={() => setSelectedCategory('active')}
          >
            Active ({markets.filter(m => getMarketStatus(m) === 'Open').length})
          </button>
          <button 
            className={`category-tab ${selectedCategory === 'ended' ? 'active' : ''}`}
            onClick={() => setSelectedCategory('ended')}
          >
            Ended ({markets.filter(m => ['Closed', 'Settled', 'Expired'].includes(getMarketStatus(m))).length})
          </button>
        </div>
      </div>

      {/* KPI Cards */}
      <div className="kpi-grid">
        <div className="kpi-card">
          <div className="kpi-icon-wrapper kpi-volume">
            <DollarSign size={24} />
          </div>
          <div className="kpi-content">
            <div className="kpi-label">24h Volume</div>
            <div className="kpi-value">{kpiData.totalVolume} ETH</div>
            <div className="kpi-change positive">
              <TrendingUp size={14} />
              +12.3%
            </div>
          </div>
        </div>

        <div className="kpi-card">
          <div className="kpi-icon-wrapper kpi-traders">
            <Users size={24} />
          </div>
          <div className="kpi-content">
            <div className="kpi-label">Active Traders</div>
            <div className="kpi-value">{kpiData.activeTraders}</div>
            <div className="kpi-change positive">
              <TrendingUp size={14} />
              +5.7%
            </div>
          </div>
        </div>

        <div className="kpi-card">
          <div className="kpi-icon-wrapper kpi-winrate">
            <Trophy size={24} />
          </div>
          <div className="kpi-content">
            <div className="kpi-label">Avg Win Rate</div>
            <div className="kpi-value">{kpiData.avgWinRate}%</div>
            <div className="kpi-change neutral">
              <Activity size={14} />
              0.0%
            </div>
          </div>
        </div>

        <div className="kpi-card">
          <div className="kpi-icon-wrapper kpi-markets">
            <Target size={24} />
          </div>
          <div className="kpi-content">
            <div className="kpi-label">Total Markets</div>
            <div className="kpi-value">{kpiData.totalMarkets}</div>
            <div className="kpi-change positive">
              <TrendingUp size={14} />
              +2 today
            </div>
          </div>
        </div>
      </div>

      {/* Market Activity Section */}
      <div className="market-activity-section">
        <div className="activity-header">
          <h2 className="activity-title">
            <Zap size={20} />
            Live Market Activity
          </h2>
          <span className="activity-status">
            <span className="activity-dot"></span>
            Real-time analysis by AI
          </span>
        </div>
      </div>

      {/* Markets Stack */}
      {loading && markets.length === 0 ? (
        <div className="loading-container">
          <Loader2 size={48} className="loading-spinner" />
          <p>Loading markets from blockchain...</p>
        </div>
      ) : filteredMarkets.length > 0 ? (
        <div className="markets-stack">
          {filteredMarkets.map((market) => {
            const stats = marketStats[market.address];
            const status = getMarketStatus(market);
            const countdown = getMarketCountdown(market);
            const [countryCode1, countryCode2] = getCountriesFromQuestion(market.question);
            
            return (
              <div 
                key={market.address}
                className={`market-card-stack ${status !== 'Open' ? 'market-ended' : ''}`}
                onClick={() => handleMarketClick(market)}
              >
                <div className="market-card-left">
                  <div className="market-card-header-row">
                    <div className="market-flags-wrapper">
                      <div className="market-flag">
                        <ReactCountryFlag countryCode={countryCode1} svg style={{ width: '1.5em', height: '1.5em', borderRadius: '50%' }} />
                      </div>
                      <span className="vs-text">VS</span>
                      <div className="market-flag">
                        <ReactCountryFlag countryCode={countryCode2} svg style={{ width: '1.5em', height: '1.5em', borderRadius: '50%' }} />
                      </div>
                    </div>
                    <button
                      className="market-action-btn"
                      style={status !== 'Open' ? { backgroundColor: '#6b7280', cursor: 'not-allowed', pointerEvents: 'none' } : {}}
                      disabled={status !== 'Open'}
                      onClick={async (e) => {
                        e.stopPropagation();
                        try {
                          const fc = await createFunctionsContract(market.functionsConsumer, provider);
                          await requestOddsUpdate(fc);
                          alert("Odds updated for this contract!");
                        } catch (err) {
                          alert("Error updating: " + err.message);
                        }
                      }}
                    >
                      Update Odds
                    </button>
                  </div>
                  <div className="market-info">
                    <h3 className="market-question-stack">{market.question}</h3>
                    <div className="market-meta-stack">
                      <span className="market-meta-item">
                        <Clock size={14} />
                        {countdown}
                      </span>
                      {stats && (
                        <>
                          <span className="market-meta-item">
                            <Users size={14} />
                            {stats.participants} traders
                          </span>
                          <span className="market-meta-item">
                            <DollarSign size={14} />
                            {stats.totalYesETH} / {stats.totalNoETH} ETH
                          </span>
                          {stats.tokenBettingEnabled && (
                            <span className="market-meta-item">
                              <DollarSign size={14} />
                              {stats.totalYesToken} / {stats.totalNoToken} CCIP-BnM
                            </span>
                          )}
                        </>
                      )}
                    </div>
                  </div>
                </div>

                <div className="market-card-right">
                  {stats ? (
                    <>
                      <div className="odds-display">
                        <div className="odds-item yes">
                          <span className="odds-label">YES</span>
                          <span className="odds-value">{stats.oddsYes}%</span>
                        </div>
                        <div className="odds-divider"></div>
                        <div className="odds-item no">
                          <span className="odds-label">NO</span>
                          <span className="odds-value">{stats.oddsNo}%</span>
                        </div>
                      </div>
                      <button className="market-action-btn">
                        Trade
                        <ChevronRight size={16} />
                      </button>
                    </>
                  ) : (
                    <div className="market-loading">
                      <Loader2 size={20} className="loading-spinner" />
                    </div>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      ) : (
        <div className="empty-state">
          <BarChart3 size={48} color="#6b7280" style={{ margin: '0 auto 1rem' }} />
          <p>No markets found</p>
          {markets.length === 0 ? (
            <p>Be the first to create a prediction market!</p>
          ) : (
            <p>Try adjusting your search or filters</p>
          )}
        </div>
      )}

      {/* Market Insights */}
      <div className="market-insights">
        <h3 className="insights-title">
          <Activity size={20} />
          AI Market Insights
        </h3>
        <div className="insights-grid">
          <div className="insight-card">
            <div className="insight-icon pattern-detected">
              <Target size={20} />
            </div>
            <div className="insight-content">
              <h4>Pattern Detected</h4>
              <p>Bull flag on ETH/USDC with 89% success rate</p>
            </div>
          </div>
          <div className="insight-card">
            <div className="insight-icon market-alert">
              <AlertCircle size={20} />
            </div>
            <div className="insight-content">
              <h4>Market Alert</h4>
              <p>High volatility expected in next 4 hours</p>
            </div>
          </div>
        </div>
      </div>

      {/* Info about contract addresses */}
      {markets.length > 0 && (
        <div className="info-box" style={{ marginTop: '2rem' }}>
          <AlertCircle size={20} color="#3b82f6" className="info-box-icon" />
          <div className="info-box-content">
            <p className="info-box-title">Real-time Contract Data</p>
            <p>
              Showing {markets.length} market{markets.length !== 1 ? 's' : ''} from the blockchain registry. 
              Market data updates automatically every 30 seconds.
            </p>
          </div>
        </div>
      )}
    </div>
  );
}

export default MarketList;