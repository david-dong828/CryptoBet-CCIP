// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "../src/PredictionMarket.sol";
import "forge-std/console.sol";

// CAN use this to specifically test with details: forge test --match-test "testCannotClaimTwice|testClaimWinningsYesWins" -vvvv
contract PredictionMarketTest is Test {
    PredictionMarket market;

    // IMPORTANT! Cannot use 0x1 to 0x9 as Ethereum reserves addresses 0x1 through 0x9 for precompiled contracts
    address player1 = address(0x123);  
    address player2 = address(0x234);

    function setUp() public {
        uint256 closeTime = block.timestamp + 1 days;
        uint256 settleTime = block.timestamp + 2 days;
        market = new PredictionMarket("Will ETH be above $4000 on Jul 1?",closeTime,settleTime);
    }

    function testPlaceBetYes() public {
        vm.deal(player1, 1 ether);
        vm.prank(player1);
        market.placeBet{value: 1 ether}(true);
        (uint256 amount, bool prediction, ) = market.bets(player1);
        assertEq(amount, 1 ether);
        assertTrue(prediction);
        assertEq(market.totalYes(), 1 ether);
    }

    function testPlaceBetNo() public {
        vm.deal(player2, 2 ether);
        vm.prank(player2);
        market.placeBet{value: 2 ether}(false);
        (uint256 amount, bool prediction, ) = market.bets(player2);
        assertEq(amount, 2 ether);
        assertFalse(prediction);
        assertEq(market.totalNo(), 2 ether);
    }

    function testClaimWinningsYesWins() public {
        vm.deal(player1, 1 ether);
        vm.prank(player1);
        market.placeBet{value: 1 ether}(true);

        vm.deal(player2, 2 ether);
        vm.prank(player2);
        market.placeBet{value: 2 ether}(false);

        // Close and settle market with "Yes" as winner
        market.closeMarket();
        market.setResult(true);

        // Check balances before claim
        uint256 balBefore1 = player1.balance;

        // Player1 (YES) claims winnings
        vm.prank(player1);
        market.claimWinnings();

        // Check that player1 received payout (should be 3 ether)
        uint256 balAfter1 = player1.balance;
        assertEq(balAfter1 - balBefore1, 3 ether);

        // Player2 (NO) tries to claim winnings (should succeed but get nothing)
        uint256 balBefore2 = player2.balance;
        vm.prank(player2);
        market.claimWinnings();
        uint256 balAfter2 = player2.balance;
        assertEq(balAfter2 - balBefore2, 0);
        
        // Verify player2's bet is marked as claimed
        (, , bool claimed) = market.bets(player2);
        assertTrue(claimed);
    }

    function testCannotBetTwice() public {
        vm.deal(player1, 2 ether);
        vm.prank(player1);
        market.placeBet{value: 1 ether}(true);
        vm.prank(player1);
        vm.expectRevert("Already bet");
        market.placeBet{value: 1 ether}(true);
    }

    function testCannotClaimTwice() public {
        vm.deal(player1, 1 ether);
        vm.prank(player1);
        market.placeBet{value: 1 ether}(true);

        market.closeMarket();
        market.setResult(true);

        // first claim should succeed
        vm.prank(player1);
        market.claimWinnings();

        // second claim should revert with 'Already claimed'
        vm.prank(player1);
        vm.expectRevert("Already claimed");
        market.claimWinnings();
        
        
    }

    function testSetOddsByUpdater() public {
        address updater = address(0x456);
        market.setOddsUpdater(updater);

        vm.prank(updater);
        market.setOdds(60, 40);

        assertEq(market.oddsYes(), 60);
        assertEq(market.oddsNo(), 40);
    }

    function testSetOddsFailByNonUpdater() public {
        address updater = address(0x456);
        market.setOddsUpdater(updater);

        vm.expectRevert("Not odds updater");
        market.setOdds(60, 40);
    }

    function testSetOddsInvalidSum() public {
        address updater = address(0x456);
        market.setOddsUpdater(updater);

        vm.prank(updater);
        vm.expectRevert("Odds must sum to 100");
        market.setOdds(70, 50); // Invalid: 70 + 50 = 120
    }

    function testSetOddsUpdaterByNonOwner() public {
        address attacker = address(0x999);
        vm.prank(attacker);
        vm.expectRevert("Not owner");
        market.setOddsUpdater(address(0x111));
    }


}
