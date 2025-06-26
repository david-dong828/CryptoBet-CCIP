// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import "../src/FunctionsConsumer.sol";

contract DeployFunctionsConsumerScript is Script {
    function setUp() public {}

    function run() public {
        // Start broadcast
        vm.startBroadcast();

        string memory source = "";

        // Insert your deployed PredictionMarket address here
        address predictionMarketAddr = 0x4172386A4ad2D42f2ea4B5bED0D5cFd2e36f6a44;

        // The Chainlink Functions router for Sepolia
        address router = 0xb83E47C2bC239B3bf370bc41e1459A34b41238D0;

        // Deploy FunctionsConsumer
        new FunctionsConsumer(router, predictionMarketAddr,source);

        vm.stopBroadcast();
    }
}
