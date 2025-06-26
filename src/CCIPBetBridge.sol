// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {CCIPReceiver} from "@chainlink/contracts-ccip/contracts/applications/CCIPReceiver.sol";
import {OwnerIsCreator} from "@chainlink/contracts/src/v0.8/shared/access/OwnerIsCreator.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {IRouterClient} from "@chainlink/contracts-ccip/contracts/interfaces/IRouterClient.sol";
import {Client} from "@chainlink/contracts-ccip/contracts/libraries/Client.sol";

interface IPredictionMarketBet {
    function placeBetWithToken(address user, bool prediction, uint256 amount) external;
}

contract CCIPBetBridge is CCIPReceiver, OwnerIsCreator {
    using SafeERC20 for IERC20;

    error SourceChainNotAllowed(uint64 chain);
    error SenderNotAllowed(address sender);
    error NotEnoughFee(uint256 needed, uint256 current);

    event BetMessageSent(
        bytes32 messageId,
        uint64 destinationChainSelector,
        address receiver,
        string eventId,
        bool prediction,
        address token,
        uint256 amount,
        address feeToken,
        uint256 fee
    );

    event BetReceived(
        bytes32 messageId,
        string eventId,
        bool prediction,
        address token,
        uint256 amount,
        address bettor
    );

    struct BetData {
        string eventId;
        bool prediction;
    }

    mapping(uint64 => bool) public allowlistedSourceChains;
    mapping(address => bool) public allowlistedSenders;

    IERC20 public linkToken;
    IPredictionMarketBet public predictionMarket;

    constructor(
        address _router,
        address _link,
        address _predictionMarket,
        uint64[] memory _allowedChains,
        address[] memory _allowedSenders
    ) CCIPReceiver(_router) {
        linkToken = IERC20(_link);
        predictionMarket = IPredictionMarketBet(_predictionMarket);

        // Set up allowlisted chains
        for (uint256 i = 0; i < _allowedChains.length; i++) {
            allowlistedSourceChains[_allowedChains[i]] = true;
        }

        // Set up allowlisted senders
        for (uint256 i = 0; i < _allowedSenders.length; i++) {
            allowlistedSenders[_allowedSenders[i]] = true;
        }
    }

    function sendBetWithToken(
        uint64 destinationChainSelector,
        address receiver,
        string calldata eventId,
        bool prediction,
        address token,
        uint256 amount,
        bool payWithLINK
    ) external payable returns (bytes32 messageId) {
        require(receiver != address(0), "Receiver cannot be zero address");
        require(amount > 0, "Amount must be > 0");

        // First transfer tokens from the user to this contract
        IERC20(token).transferFrom(msg.sender, address(this), amount);
        

        bytes memory data = abi.encode(BetData(eventId, prediction));
        Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
        tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

        address feeToken = payWithLINK ? address(linkToken) : address(0);
        Client.EVM2AnyMessage memory msgToSend = Client.EVM2AnyMessage({
            receiver: abi.encode(receiver),
            data: data,
            tokenAmounts: tokenAmounts,
            extraArgs: Client._argsToBytes(Client.GenericExtraArgsV2({
                gasLimit: 900000,
                allowOutOfOrderExecution: true
            })),
            feeToken: feeToken
        });

        IRouterClient router = IRouterClient(this.getRouter());
        uint256 fee = router.getFee(destinationChainSelector, msgToSend);

        if (payWithLINK) {
            if (fee > linkToken.balanceOf(address(this)))
                revert NotEnoughFee(fee, linkToken.balanceOf(address(this)));
            linkToken.approve(address(router), fee);
        } else {
            if (fee > address(this).balance) revert NotEnoughFee(fee, address(this).balance);
        }

        // Then approve the router
        IERC20(token).approve(address(router), amount);

        messageId = router.ccipSend{value: payWithLINK ? 0 : fee}(destinationChainSelector, msgToSend);

        emit BetMessageSent(messageId, destinationChainSelector, receiver, eventId, prediction, token, amount, feeToken, fee);
    }

    function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
        // Decode the original sender address from the message
        address user = abi.decode(message.sender, (address));
        
        if (!allowlistedSourceChains[message.sourceChainSelector])
            revert SourceChainNotAllowed(message.sourceChainSelector);
        if (!allowlistedSenders[user]) revert SenderNotAllowed(user);

        BetData memory bet = abi.decode(message.data, (BetData));
        address token = message.destTokenAmounts[0].token;
        uint256 amount = message.destTokenAmounts[0].amount;

        emit BetReceived(message.messageId, bet.eventId, bet.prediction, token, amount, user);

        // Approve the PredictionMarket to spend the received tokens
        IERC20(token).approve(address(predictionMarket), amount);

        // Direct call to updated PredictionMarket with user address
        predictionMarket.placeBetWithToken(user, bet.prediction, amount);
    }

    function allowSourceChain(uint64 selector, bool allowed) external onlyOwner {
        allowlistedSourceChains[selector] = allowed;
    }

    function allowSender(address sender, bool allowed) external onlyOwner {
        allowlistedSenders[sender] = allowed;
    }

    receive() external payable {}
}