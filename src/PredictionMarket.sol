// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract PredictionMarket {
    enum MarketStatus { Open, Closed, Settled }
    
    address public owner;
    address public oddsUpdater; // Chainlink Functions consumer contract
    string public question;
    uint256 public oddsYes; // Odds for "Yes" (e.g. 60 means 60%)
    uint256 public oddsNo;  // Odds for "No"
    MarketStatus public status;
    bool public result;
    
    // Timing variables
    uint256 public closeTime;  // When betting closes
    uint256 public settleTime; // When market gets settled
    
    mapping(address => Bet) public bets;
    uint256 public totalYes;
    uint256 public totalNo;

    struct Bet {
        uint256 amount;
        bool prediction;
        bool claimed;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    modifier onlyOddsUpdater() {
        require(msg.sender == oddsUpdater, "Not odds updater");
        _;
    }

    modifier inStatus(MarketStatus s) {
        require(status == s, "Invalid status");
        _;
    }
    
    modifier beforeCloseTime() {
        require(block.timestamp < closeTime, "Market closed");
        _;
    }
    
    modifier afterCloseTime() {
        require(block.timestamp >= closeTime, "Market still open");
        _;
    }
    
    modifier afterSettleTime() {
        require(block.timestamp >= settleTime, "Settlement time not reached");
        _;
    }

    event OddsUpdated(uint256 yes, uint256 no);
    event MarketClosed(uint256 timestamp);
    event MarketSettled(bool result, uint256 timestamp);
    event BetPlaced(address indexed user, bool prediction, uint256 amount);
    event WinningsClaimed(address indexed user, uint256 amount);

    constructor(
        string memory _question,
        uint256 _closeTime,
        uint256 _settleTime
    ) {
        require(_closeTime > block.timestamp, "Close time must be in future");
        require(_settleTime > _closeTime, "Settle time must be after close time");
        
        owner = msg.sender;
        question = _question;
        closeTime = _closeTime;
        settleTime = _settleTime;
        status = MarketStatus.Open;
        oddsYes = 50;
        oddsNo = 50;
    }

    function setOddsUpdater(address _updater) external onlyOwner {
        oddsUpdater = _updater;
    }

    // Chainlink Functions consumer will call this
    function setOdds(uint256 _oddsYes, uint256 _oddsNo) external onlyOddsUpdater inStatus(MarketStatus.Open) beforeCloseTime {
        require(_oddsYes + _oddsNo == 100, "Odds must sum to 100");
        oddsYes = _oddsYes;
        oddsNo = _oddsNo;
        emit OddsUpdated(_oddsYes, _oddsNo);
    }

    function placeBet(bool prediction) external payable inStatus(MarketStatus.Open) beforeCloseTime {
        require(msg.value > 0, "No bet");
        require(bets[msg.sender].amount == 0, "Already bet");
        
        bets[msg.sender] = Bet(msg.value, prediction, false);
        if (prediction) {
            totalYes += msg.value;
        } else {
            totalNo += msg.value;
        }
        
        emit BetPlaced(msg.sender, prediction, msg.value);
    }

    // Anyone can call this to close the market after close time
    function closeMarket() external afterCloseTime inStatus(MarketStatus.Open) {
        status = MarketStatus.Closed;
        emit MarketClosed(block.timestamp);
    }

    // Auto-close function that can be called by anyone
    function checkAndCloseMarket() external {
        if (block.timestamp >= closeTime && status == MarketStatus.Open) {
            status = MarketStatus.Closed;
            emit MarketClosed(block.timestamp);
        }
    }

    // Owner sets the result after settlement time
    function setResult(bool _result) external onlyOwner afterSettleTime inStatus(MarketStatus.Closed) {
        result = _result;
        status = MarketStatus.Settled;
        emit MarketSettled(_result, block.timestamp);
    }

    // Auto-settle function (would typically be called by Chainlink Automation or similar)
    function autoSettle(bool _result) external onlyOwner afterSettleTime inStatus(MarketStatus.Closed) {
        result = _result;
        status = MarketStatus.Settled;
        emit MarketSettled(_result, block.timestamp);
    }

    function claimWinnings() external inStatus(MarketStatus.Settled) {
        Bet storage b = bets[msg.sender];
        require(!b.claimed, "Already claimed");
        require(b.amount > 0, "No bet");
        
        b.claimed = true;
        if (b.prediction == result) {
            uint256 totalWinning = result ? totalYes : totalNo;
            if (totalWinning > 0) {
                // Winner: gets share of pool
                uint256 payout = (totalYes + totalNo) * b.amount / totalWinning;

                if (payout > 0 && address(this).balance >= payout) {
                    payable(msg.sender).transfer(payout);
                    emit WinningsClaimed(msg.sender, payout);
                }
            }
        } 
    }
    
    // View functions
    function getTimeRemaining() external view returns (uint256 closeTimeRemaining, uint256 settleTimeRemaining) {
        closeTimeRemaining = block.timestamp >= closeTime ? 0 : closeTime - block.timestamp;
        settleTimeRemaining = block.timestamp >= settleTime ? 0 : settleTime - block.timestamp;
    }
    
    function isMarketExpired() external view returns (bool) {
        return block.timestamp >= closeTime;
    }
    
    function isSettlementTime() external view returns (bool) {
        return block.timestamp >= settleTime;
    }
    
    function getMarketTimes() external view returns (uint256, uint256) {
        return (closeTime, settleTime);
    }
}