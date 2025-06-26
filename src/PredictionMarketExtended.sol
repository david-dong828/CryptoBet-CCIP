// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract PredictionMarketExtended {
    enum MarketStatus { Open, Closed, Settled }

    address public owner;
    address public oddsUpdater;

    string public question;
    uint256 public oddsYes; // in %
    uint256 public oddsNo;
    MarketStatus public status;
    bool public result;

    uint256 public closeTime;
    uint256 public settleTime;

    IERC20 public betToken;
    uint256 public tokenBetAmount;
    bool public tokenBettingEnabled;

    struct Bet {
        uint256 amount;
        bool prediction;
        bool claimed;
        bool isToken; // true if ERC20, false if ETH
    }

    // ************ SEPARATED POOLS FOR ETH AND TOKENS ***************
    mapping(address => Bet[]) public bets;
    
    // ETH pools
    uint256 public totalYesETH;
    uint256 public totalNoETH;
    
    // Token pools
    uint256 public totalYesToken;
    uint256 public totalNoToken;

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
    event BetPlaced(address indexed user, bool prediction, uint256 amount, bool isToken, uint256 betIndex);
    event WinningsClaimed(address indexed user, uint256 betIndex, uint256 amount, bool isToken);

    constructor(
        string memory _question,
        uint256 _closeTime,
        uint256 _settleTime,
        address _betToken,
        uint256 _tokenBetAmount,
        bool _enableTokenBetting
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

        // Set token betting configuration
        if (_betToken != address(0)) {
            betToken = IERC20(_betToken);
            tokenBetAmount = _tokenBetAmount;
            tokenBettingEnabled = _enableTokenBetting;
        }
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

    // *********** ETH Bet *************
    function placeBet(bool prediction) external payable inStatus(MarketStatus.Open) beforeCloseTime {
        require(msg.value > 0, "No bet");
        bets[msg.sender].push(Bet(msg.value, prediction, false, false));
        
        if (prediction) {
            totalYesETH += msg.value;
        } else {
            totalNoETH += msg.value;
        }
        
        emit BetPlaced(msg.sender, prediction, msg.value, false, bets[msg.sender].length - 1);
    }

    // *********** ERC20 Bet *************
    function placeBetWithToken(address user, bool prediction, uint256 amount)
        external
        inStatus(MarketStatus.Open)
        beforeCloseTime
    {
        require(tokenBettingEnabled, "Token betting disabled");
        require(betToken != IERC20(address(0)), "Token not set");
        require(amount > 0, "No bet");

        // Transfer tokens from the caller (e.g., CCIPBetBridge)
        betToken.transferFrom(msg.sender, address(this), amount);
        
        // Record the bet for the actual user
        bets[user].push(Bet(amount, prediction, false, true));

        if (prediction) {
            totalYesToken += amount;
        } else {
            totalNoToken += amount;
        }

        emit BetPlaced(user, prediction, amount, true, bets[user].length - 1);
    }

    // Anyone can call this to close the market after close time
    function closeMarket() external afterCloseTime inStatus(MarketStatus.Open) {
        status = MarketStatus.Closed;
        emit MarketClosed(block.timestamp);
    }

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

    function autoSettle(bool _result) external onlyOwner afterSettleTime inStatus(MarketStatus.Closed) {
        result = _result;
        status = MarketStatus.Settled;
        emit MarketSettled(_result, block.timestamp);
    }

    // *********** Claim for all winning bets - SEPARATED BY ASSET TYPE *************
    function claimWinnings() external inStatus(MarketStatus.Settled) {
        Bet[] storage userBets = bets[msg.sender];
        require(userBets.length > 0, "No bets");

        // Calculate totals for each asset type
        uint256 totalWinningETH = result ? totalYesETH : totalNoETH;
        uint256 totalWinningToken = result ? totalYesToken : totalNoToken;
        uint256 poolETH = totalYesETH + totalNoETH;
        uint256 poolToken = totalYesToken + totalNoToken;

        for (uint256 i = 0; i < userBets.length; i++) {
            Bet storage b = userBets[i];
            if (!b.claimed && b.amount > 0 && b.prediction == result) {
                b.claimed = true;
                uint256 payout = 0;

                if (b.isToken && totalWinningToken > 0) {
                    // Token payout calculation
                    payout = poolToken * b.amount / totalWinningToken;
                    if (payout > 0 && address(betToken) != address(0) && tokenBettingEnabled) {
                        betToken.transfer(msg.sender, payout);
                        emit WinningsClaimed(msg.sender, i, payout, true);
                    }
                } else if (!b.isToken && totalWinningETH > 0) {
                    // ETH payout calculation
                    payout = poolETH * b.amount / totalWinningETH;
                    if (payout > 0 && address(this).balance >= payout) {
                        payable(msg.sender).transfer(payout);
                        emit WinningsClaimed(msg.sender, i, payout, false);
                    }
                }
            }
        }
    }

    // *********** Claim specific bet indices to avoid gas issues *************
    function claimWinningsByIndices(uint256[] calldata betIndices) external inStatus(MarketStatus.Settled) {
        Bet[] storage userBets = bets[msg.sender];
        require(userBets.length > 0, "No bets");

        // Calculate totals for each asset type
        uint256 totalWinningETH = result ? totalYesETH : totalNoETH;
        uint256 totalWinningToken = result ? totalYesToken : totalNoToken;
        uint256 poolETH = totalYesETH + totalNoETH;
        uint256 poolToken = totalYesToken + totalNoToken;

        for (uint256 j = 0; j < betIndices.length; j++) {
            uint256 i = betIndices[j];
            require(i < userBets.length, "Invalid bet index");
            
            Bet storage b = userBets[i];
            if (!b.claimed && b.amount > 0 && b.prediction == result) {
                b.claimed = true;
                uint256 payout = 0;

                if (b.isToken && totalWinningToken > 0) {
                    // Token payout calculation
                    payout = poolToken * b.amount / totalWinningToken;
                    if (payout > 0 && address(betToken) != address(0) && tokenBettingEnabled) {
                        betToken.transfer(msg.sender, payout);
                        emit WinningsClaimed(msg.sender, i, payout, true);
                    }
                } else if (!b.isToken && totalWinningETH > 0) {
                    // ETH payout calculation
                    payout = poolETH * b.amount / totalWinningETH;
                    if (payout > 0 && address(this).balance >= payout) {
                        payable(msg.sender).transfer(payout);
                        emit WinningsClaimed(msg.sender, i, payout, false);
                    }
                }
            }
        }
    }

    function setBetToken(address _token, uint256 _amount) external onlyOwner {
        betToken = IERC20(_token);
        tokenBetAmount = _amount;
    }

    function enableTokenBetting(bool enabled) external onlyOwner {
        tokenBettingEnabled = enabled;
    }

    // *********** VIEW FUNCTIONS FOR FRONTEND *************
    function getPoolTotals() external view returns (
        uint256 ethYes,
        uint256 ethNo,
        uint256 tokenYes,
        uint256 tokenNo,
        uint256 totalETH,
        uint256 totalToken
    ) {
        return (
            totalYesETH,
            totalNoETH,
            totalYesToken,
            totalNoToken,
            totalYesETH + totalNoETH,
            totalYesToken + totalNoToken
        );
    }

    function getUserBetCount(address user) external view returns (uint256) {
        return bets[user].length;
    }

    function getUserBet(address user, uint256 index) external view returns (
        uint256 amount,
        bool prediction,
        bool claimed,
        bool isToken
    ) {
        require(index < bets[user].length, "Invalid bet index");
        Bet memory bet = bets[user][index];
        return (bet.amount, bet.prediction, bet.claimed, bet.isToken);
    }

    function getUserBets(address user) external view returns (Bet[] memory) {
        return bets[user];
    }

    function getUserUnclaimedWinnings(address user) external view returns (
        uint256 ethWinnings,
        uint256 tokenWinnings,
        uint256[] memory winningBetIndices
    ) {
        if (status != MarketStatus.Settled) {
            return (0, 0, new uint256[](0));
        }

        Bet[] memory userBets = bets[user];
        uint256 totalWinningETH = result ? totalYesETH : totalNoETH;
        uint256 totalWinningToken = result ? totalYesToken : totalNoToken;
        uint256 poolETH = totalYesETH + totalNoETH;
        uint256 poolToken = totalYesToken + totalNoToken;

        uint256 tempEthWinnings = 0;
        uint256 tempTokenWinnings = 0;
        uint256 winningCount = 0;

        // Count winning bets first
        for (uint256 i = 0; i < userBets.length; i++) {
            if (!userBets[i].claimed && userBets[i].amount > 0 && userBets[i].prediction == result) {
                winningCount++;
            }
        }

        uint256[] memory indices = new uint256[](winningCount);
        uint256 indexCounter = 0;

        for (uint256 i = 0; i < userBets.length; i++) {
            Bet memory b = userBets[i];
            if (!b.claimed && b.amount > 0 && b.prediction == result) {
                indices[indexCounter] = i;
                indexCounter++;

                if (b.isToken && totalWinningToken > 0) {
                    tempTokenWinnings += poolToken * b.amount / totalWinningToken;
                } else if (!b.isToken && totalWinningETH > 0) {
                    tempEthWinnings += poolETH * b.amount / totalWinningETH;
                }
            }
        }

        return (tempEthWinnings, tempTokenWinnings, indices);
    }

    function getTimeRemaining() external view returns (uint256, uint256) {
        uint256 closeRem = block.timestamp >= closeTime ? 0 : closeTime - block.timestamp;
        uint256 settleRem = block.timestamp >= settleTime ? 0 : settleTime - block.timestamp;
        return (closeRem, settleRem);
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

    receive() external payable {}
}