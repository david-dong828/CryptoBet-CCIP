// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MarketRegistry {
    struct MarketInfo {
        address marketAddress;
        address functionsConsumer;
        string question;
        uint256 closeTime;
        uint256 settleTime;
        uint256 createdAt;
        address creator;
        bool isActive;
    }
    
    address public owner;
    MarketInfo[] public markets;
    mapping(address => uint256[]) public userMarkets; // markets created by user
    mapping(address => bool) public registeredMarkets; // quick lookup
    
    event MarketRegistered(
        address indexed marketAddress,
        address indexed creator,
        string question,
        uint256 closeTime,
        uint256 settleTime
    );
    
    event MarketDeactivated(address indexed marketAddress);
    
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }
    
    constructor() {
        owner = msg.sender;
    }
    
    function registerMarket(
        address _marketAddress,
        address _functionsConsumer,
        string memory _question,
        uint256 _closeTime,
        uint256 _settleTime
    ) external {
        require(_marketAddress != address(0), "Invalid market address");
        require(_closeTime > block.timestamp, "Close time must be in future");
        require(_settleTime > _closeTime, "Settle time must be after close time");
        require(!registeredMarkets[_marketAddress], "Market already registered");
        
        MarketInfo memory newMarket = MarketInfo({
            marketAddress: _marketAddress,
            functionsConsumer: _functionsConsumer,
            question: _question,
            closeTime: _closeTime,
            settleTime: _settleTime,
            createdAt: block.timestamp,
            creator: msg.sender,
            isActive: true
        });
        
        markets.push(newMarket);
        userMarkets[msg.sender].push(markets.length - 1);
        registeredMarkets[_marketAddress] = true;
        
        emit MarketRegistered(_marketAddress, msg.sender, _question, _closeTime, _settleTime);
    }
    
    function deactivateMarket(address _marketAddress) external onlyOwner {
        for (uint256 i = 0; i < markets.length; i++) {
            if (markets[i].marketAddress == _marketAddress) {
                markets[i].isActive = false;
                emit MarketDeactivated(_marketAddress);
                break;
            }
        }
    }
    
    function getAllMarkets() external view returns (MarketInfo[] memory) {
        return markets;
    }
    
    function getActiveMarkets() external view returns (MarketInfo[] memory) {
        uint256 activeCount = 0;
        
        // Count active markets
        for (uint256 i = 0; i < markets.length; i++) {
            if (markets[i].isActive) {
                activeCount++;
            }
        }
        
        // Create array of active markets
        MarketInfo[] memory activeMarkets = new MarketInfo[](activeCount);
        uint256 currentIndex = 0;
        
        for (uint256 i = 0; i < markets.length; i++) {
            if (markets[i].isActive) {
                activeMarkets[currentIndex] = markets[i];
                currentIndex++;
            }
        }
        
        return activeMarkets;
    }
    
    function getMarketInfo(address _marketAddress) external view returns (MarketInfo memory) {
        for (uint256 i = 0; i < markets.length; i++) {
            if (markets[i].marketAddress == _marketAddress) {
                return markets[i];
            }
        }
        revert("Market not found");
    }
    
    function getUserMarkets(address _user) external view returns (MarketInfo[] memory) {
        uint256[] memory userMarketIndices = userMarkets[_user];
        MarketInfo[] memory userMarketsList = new MarketInfo[](userMarketIndices.length);
        
        for (uint256 i = 0; i < userMarketIndices.length; i++) {
            userMarketsList[i] = markets[userMarketIndices[i]];
        }
        
        return userMarketsList;
    }
    
    function getMarketsCount() external view returns (uint256) {
        return markets.length;
    }
    
    function getActiveMarketsCount() external view returns (uint256) {
        uint256 count = 0;
        for (uint256 i = 0; i < markets.length; i++) {
            if (markets[i].isActive) {
                count++;
            }
        }
        return count;
    }
}