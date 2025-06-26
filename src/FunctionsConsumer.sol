// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {FunctionsClient} from "@chainlink/contracts/src/v0.8/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "@chainlink/contracts/src/v0.8/functions/v1_0_0/libraries/FunctionsRequest.sol";

interface IPredictionMarket {
    function setOdds(uint256 _oddsYes, uint256 _oddsNo) external;
}

contract FunctionsConsumer is FunctionsClient {
    using FunctionsRequest for FunctionsRequest.Request;

    address public owner;
    address public predictionMarket;
    bytes32 public latestRequestId;
    bytes public s_lastError;
    bytes public s_lastResponse;
    
    // Make source code configurable
    string public source;

    error UnexpectedRequestID(bytes32 requestId);

    event OddsRequestSent(bytes32 requestId);
    event OddsFulfilled(bytes32 requestId, uint256 oddsYes, uint256 oddsNo);
    event RequestFailed(bytes32 requestId, bytes error);
    event SourceUpdated(string newSource);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    constructor(
        address router, 
        address _predictionMarket,
        string memory _source
    ) FunctionsClient(router) {
        owner = msg.sender;
        predictionMarket = _predictionMarket;
        source = _source;
    }

    // Allow owner to update the source code
    function updateSource(string memory _newSource) external onlyOwner {
        source = _newSource;
        emit SourceUpdated(_newSource);
    }

    function requestOdds(
        string[] memory args,
        uint64 subscriptionId,
        uint32 gasLimit
    ) external onlyOwner returns (bytes32 requestId) {
        FunctionsRequest.Request memory req;
        req.initializeRequestForInlineJavaScript(source);
        
        if (args.length > 0) {
            req.setArgs(args);
        }

        requestId = _sendRequest(
            req.encodeCBOR(),
            subscriptionId,
            gasLimit,
            0x66756e2d657468657265756d2d7365706f6c69612d3100000000000000000000
        );
        
        latestRequestId = requestId;
        emit OddsRequestSent(requestId);
    }

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {
        if (latestRequestId != requestId) {
            revert UnexpectedRequestID(requestId);
        }

        s_lastResponse = response;
        s_lastError = err;

        if (err.length > 0) {
            emit RequestFailed(requestId, err);
            return;
        }

        // Process the response
        string memory oddsStr = string(response);
        
        // Parse the comma-separated values
        (uint256 oddsYes, uint256 oddsNo) = parseOddsString(oddsStr);
        
        // Validate odds
        if (oddsYes == 0 || oddsNo == 0) {
            emit RequestFailed(requestId, "Invalid odds values");
            return;
        }
        
        if (oddsYes + oddsNo != 100) {
            emit RequestFailed(requestId, "Odds must sum to 100");
            return;
        }
        
        // Update the prediction market
        try IPredictionMarket(predictionMarket).setOdds(oddsYes, oddsNo) {
            emit OddsFulfilled(requestId, oddsYes, oddsNo);
        } catch {
            emit RequestFailed(requestId, "Failed to set odds in prediction market");
        }
    }

    // Parse odds string helper
    function parseOddsString(string memory oddsStr) public pure returns (uint256, uint256) {
        bytes memory b = bytes(oddsStr);
        uint256 commaIndex = 0;
        
        // Find comma position
        for (uint256 i = 0; i < b.length; i++) {
            if (b[i] == ",") {
                commaIndex = i;
                break;
            }
        }
        
        require(commaIndex > 0, "Invalid format");
        
        // Extract yes odds
        uint256 oddsYes = 0;
        for (uint256 i = 0; i < commaIndex; i++) {
            require(b[i] >= "0" && b[i] <= "9", "Invalid character");
            oddsYes = oddsYes * 10 + (uint8(b[i]) - 48);
        }
        
        // Extract no odds
        uint256 oddsNo = 0;
        for (uint256 i = commaIndex + 1; i < b.length; i++) {
            require(b[i] >= "0" && b[i] <= "9", "Invalid character");
            oddsNo = oddsNo * 10 + (uint8(b[i]) - 48);
        }
        
        return (oddsYes, oddsNo);
    }

    function setPredictionMarket(address _pm) external onlyOwner {
        predictionMarket = _pm;
    }

    // Emergency function to manually set odds
    function manualSetOdds(uint256 _oddsYes, uint256 _oddsNo) external onlyOwner {
        require(_oddsYes > 0 && _oddsNo > 0, "Invalid odds");
        require(_oddsYes + _oddsNo == 100, "Odds must sum to 100");
        IPredictionMarket(predictionMarket).setOdds(_oddsYes, _oddsNo);
        emit OddsFulfilled(latestRequestId, _oddsYes, _oddsNo);
    }

    // Debugging functions
    function getLastResponse() external view returns (bytes memory) {
        return s_lastResponse;
    }

    function getLastError() external view returns (bytes memory) {
        return s_lastError;
    }

    function debugDecodeResponse() external view returns (string memory decoded, uint256 oddsYes, uint256 oddsNo) {
        if (s_lastResponse.length == 0) return ("No response", 0, 0);
        
        decoded = string(s_lastResponse);
        try this.parseOddsString(decoded) returns (uint256 yes, uint256 no) {
            oddsYes = yes;
            oddsNo = no;
        } catch {
            oddsYes = 0;
            oddsNo = 0;
        }
    }

    // View function to get current source code
    function getSource() external view returns (string memory) {
        return source;
    }
}