// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_coordinator_v2

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type VRFCoordinatorV2FeeConfig struct {
	FulfillmentFlatFeeLinkPPMTier1 uint32
	FulfillmentFlatFeeLinkPPMTier2 uint32
	FulfillmentFlatFeeLinkPPMTier3 uint32
	FulfillmentFlatFeeLinkPPMTier4 uint32
	FulfillmentFlatFeeLinkPPMTier5 uint32
	ReqsForTier2                   *big.Int
	ReqsForTier3                   *big.Int
	ReqsForTier4                   *big.Int
	ReqsForTier5                   *big.Int
}

type VRFCoordinatorV2RequestCommitment struct {
	BlockNum         uint64
	SubId            uint64
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
}

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

var VRFCoordinatorV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"link\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockhashStore\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"linkEthFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLOCKHASH_STORE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBlockhashStoreInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LINK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractLinkTokenInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LINK_ETH_FEED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_CONSUMERS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_NUM_WORDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptSubscriptionOwnerTransfer\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addConsumer\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"consumer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelSubscription\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSubscription\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterProvingKey\",\"inputs\":[{\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structVRF.Proof\",\"components\":[{\"name\":\"pk\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gamma\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"c\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uWitness\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"sHashWitness\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"zInv\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"rc\",\"type\":\"tuple\",\"internalType\":\"structVRFCoordinatorV2.RequestCommitment\",\"components\":[{\"name\":\"blockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numWords\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCommitment\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stalenessSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentSubId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFallbackWeiPerUnitLink\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"fulfillmentFlatFeeLinkPPMTier1\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier2\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier3\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier4\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier5\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"reqsForTier2\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier3\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier4\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier5\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeTier\",\"inputs\":[{\"name\":\"reqCount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubscription\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"reqCount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consumers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOfKey\",\"inputs\":[{\"name\":\"publicKey\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onTokenTransfer\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oracleWithdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerCancelSubscription\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingRequestExists\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerProvingKey\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeConsumer\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"consumer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numWords\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSubscriptionOwnerTransfer\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stalenessSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structVRFCoordinatorV2.FeeConfig\",\"components\":[{\"name\":\"fulfillmentFlatFeeLinkPPMTier1\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier2\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier3\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier4\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier5\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"reqsForTier2\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier3\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier4\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier5\",\"type\":\"uint24\",\"internalType\":\"uint24\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"stalenessSeconds\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVRFCoordinatorV2.FeeConfig\",\"components\":[{\"name\":\"fulfillmentFlatFeeLinkPPMTier1\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier2\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier3\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier4\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfillmentFlatFeeLinkPPMTier5\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"reqsForTier2\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier3\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier4\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"reqsForTier5\",\"type\":\"uint24\",\"internalType\":\"uint24\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsRecovered\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProvingKeyDeregistered\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProvingKeyRegistered\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomWordsFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outputSeed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payment\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomWordsRequested\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"preSeed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"numWords\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionCanceled\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionConsumerAdded\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"consumer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionConsumerRemoved\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"consumer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionCreated\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionFunded\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionOwnerTransferRequested\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionOwnerTransferred\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceInvariantViolated\",\"inputs\":[{\"name\":\"internalBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"externalBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BlockhashNotInStore\",\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasLimitTooBig\",\"inputs\":[{\"name\":\"have\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"want\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"IncorrectCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasForConsumer\",\"inputs\":[{\"name\":\"have\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"want\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCalldata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsumer\",\"inputs\":[{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"consumer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidLinkWeiPrice\",\"inputs\":[{\"name\":\"linkWei\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestConfirmations\",\"inputs\":[{\"name\":\"have\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"min\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"max\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidSubscription\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeRequestedOwner\",\"inputs\":[{\"name\":\"proposedOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeSubOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NoCorrespondingRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSuchProvingKey\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NumWordsTooBig\",\"inputs\":[{\"name\":\"have\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"want\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableFromLink\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingRequestExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProvingKeyAlreadyRegistered\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Reentrant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooManyConsumers\",\"inputs\":[]}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005b9538038062005b958339810160408190526200003491620001b1565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e8565b5050506001600160601b0319606093841b811660805290831b811660a052911b1660c052620001fb565b6001600160a01b038116331415620001435760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001ac57600080fd5b919050565b600080600060608486031215620001c757600080fd5b620001d28462000194565b9250620001e26020850162000194565b9150620001f26040850162000194565b90509250925092565b60805160601c60a05160601c60c05160601c615930620002656000396000818161051901526139090152600081816106030152613ee801526000818161036d015281816114da0152818161237701528181612dae01528181612eea015261350f01526159306000f3fe608060405234801561001057600080fd5b506004361061025b5760003560e01c80636f64f03f11610145578063ad178361116100bd578063d2f9f9a71161008c578063e72f6e3011610071578063e72f6e30146106e0578063e82ad7d4146106f3578063f2fde38b1461071657600080fd5b8063d2f9f9a7146106ba578063d7ae1d30146106cd57600080fd5b8063ad178361146105fe578063af198b9714610625578063c3f909d414610655578063caf70c4a146106a757600080fd5b80638da5cb5b11610114578063a21a23e4116100f9578063a21a23e4146105c0578063a47c7696146105c8578063a4c0ed36146105eb57600080fd5b80638da5cb5b1461059c5780639f87fad7146105ad57600080fd5b80636f64f03f1461055b5780637341c10c1461056e57806379ba509714610581578063823597401461058957600080fd5b8063356dac71116101d85780635fbbc0d2116101a757806366316d8d1161018c57806366316d8d14610501578063689c45171461051457806369bcdb7d1461053b57600080fd5b80635fbbc0d2146103f357806364d51a2a146104f957600080fd5b8063356dac71146103a757806340d6bb82146103af5780634cb48a54146103cd5780635d3b1d30146103e057600080fd5b806308821d581161022f57806315c48b841161021457806315c48b841461030e578063181f5a77146103295780631b6b6d231461036857600080fd5b806308821d58146102cf57806312b58349146102e257600080fd5b80620122911461026057806302bcc5b61461028057806304c357cb1461029557806306bfa637146102a8575b600080fd5b610268610729565b604051610277939291906153e6565b60405180910390f35b61029361028e3660046151f6565b6107a5565b005b6102936102a3366004615211565b610837565b60055467ffffffffffffffff165b60405167ffffffffffffffff9091168152602001610277565b6102936102dd366004614ea0565b6109eb565b6005546801000000000000000090046bffffffffffffffffffffffff165b604051908152602001610277565b61031660c881565b60405161ffff9091168152602001610277565b604080518082018252601681527f565246436f6f7264696e61746f72563220312e302e30000000000000000000006020820152905161027791906153d3565b61038f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610277565b600a54610300565b6103b86101f481565b60405163ffffffff9091168152602001610277565b6102936103db3660046150a0565b610bb0565b6103006103ee366004614f7f565b610fa7565b600c546040805163ffffffff80841682526401000000008404811660208301526801000000000000000084048116928201929092526c010000000000000000000000008304821660608201527001000000000000000000000000000000008304909116608082015262ffffff740100000000000000000000000000000000000000008304811660a0830152770100000000000000000000000000000000000000000000008304811660c08301527a0100000000000000000000000000000000000000000000000000008304811660e08301527d01000000000000000000000000000000000000000000000000000000000090920490911661010082015261012001610277565b610316606481565b61029361050f366004614e58565b611385565b61038f7f000000000000000000000000000000000000000000000000000000000000000081565b6103006105493660046151dd565b60009081526009602052604090205490565b610293610569366004614d9d565b6115d4565b61029361057c366004615211565b611704565b610293611951565b6102936105973660046151f6565b611a1a565b6000546001600160a01b031661038f565b6102936105bb366004615211565b611be0565b6102b661201f565b6105db6105d63660046151f6565b612202565b604051610277949392919061558a565b6102936105f9366004614dd1565b612325565b61038f7f000000000000000000000000000000000000000000000000000000000000000081565b610638610633366004614fdd565b61257c565b6040516bffffffffffffffffffffffff9091168152602001610277565b600b546040805161ffff8316815263ffffffff6201000084048116602083015267010000000000000084048116928201929092526b010000000000000000000000909204166060820152608001610277565b6103006106b5366004614ebc565b612a16565b6103b86106c83660046151f6565b612a46565b6102936106db366004615211565b612c3b565b6102936106ee366004614d82565b612d75565b6107066107013660046151f6565b612fb2565b6040519015158152602001610277565b610293610724366004614d82565b6131d5565b600b546007805460408051602080840282018101909252828152600094859460609461ffff8316946201000090930463ffffffff1693919283919083018282801561079357602002820191906000526020600020905b81548152602001906001019080831161077f575b50505050509050925092509250909192565b6107ad6131e6565b67ffffffffffffffff81166000908152600360205260409020546001600160a01b0316610806576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600360205260409020546108349082906001600160a01b0316613242565b50565b67ffffffffffffffff821660009081526003602052604090205482906001600160a01b031680610893576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336001600160a01b038216146108e5576040517fd8a3fb520000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024015b60405180910390fd5b600b546601000000000000900460ff161561092c576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff84166000908152600360205260409020600101546001600160a01b038481169116146109e55767ffffffffffffffff841660008181526003602090815260409182902060010180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0388169081179091558251338152918201527f69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be91015b60405180910390a25b50505050565b6109f36131e6565b604080518082018252600091610a22919084906002908390839080828437600092019190915250612a16915050565b6000818152600660205260409020549091506001600160a01b031680610a77576040517f77f5b84c000000000000000000000000000000000000000000000000000000008152600481018390526024016108dc565b600082815260066020526040812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b600754811015610b67578260078281548110610aca57610aca61587d565b90600052602060002001541415610b55576007805460009190610aef9060019061570b565b81548110610aff57610aff61587d565b906000526020600020015490508060078381548110610b2057610b2061587d565b6000918252602090912001556007805480610b3d57610b3d61584e565b60019003818190600052602060002001600090559055505b80610b5f8161577b565b915050610aac565b50806001600160a01b03167f72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d83604051610ba391815260200190565b60405180910390a2505050565b610bb86131e6565b60c861ffff87161115610c0b576040517fa738697600000000000000000000000000000000000000000000000000000000815261ffff871660048201819052602482015260c860448201526064016108dc565b60008213610c48576040517f43d4cf66000000000000000000000000000000000000000000000000000000008152600481018390526024016108dc565b6040805160a0808201835261ffff891680835263ffffffff89811660208086018290526000868801528a831660608088018290528b85166080988901819052600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000001690971762010000909502949094177fffffffffffffffffffffffffffffffffff000000000000000000ffffffffffff166701000000000000009092027fffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffff16919091176b010000000000000000000000909302929092179093558651600c80549489015189890151938a0151978a0151968a015160c08b015160e08c01516101008d01519588167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009099169890981764010000000093881693909302929092177fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff1668010000000000000000958716959095027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16949094176c0100000000000000000000000098861698909802979097177fffffffffffffffffff00000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000096909416959095027fffffffffffffffffff000000ffffffffffffffffffffffffffffffffffffffff16929092177401000000000000000000000000000000000000000062ffffff92831602177fffffff000000000000ffffffffffffffffffffffffffffffffffffffffffffff1677010000000000000000000000000000000000000000000000958216959095027fffffff000000ffffffffffffffffffffffffffffffffffffffffffffffffffff16949094177a01000000000000000000000000000000000000000000000000000092851692909202919091177cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167d0100000000000000000000000000000000000000000000000000000000009390911692909202919091178155600a84905590517fc21e3bd2e0b339d2848f0dd956947a88966c242c0c0c582a33137a5c1ceb5cb291610f97918991899189918991899190615445565b60405180910390a1505050505050565b600b546000906601000000000000900460ff1615610ff1576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff85166000908152600360205260409020546001600160a01b031661104a576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600090815260026020908152604080832067ffffffffffffffff808a16855292529091205416806110ba576040517ff0019fe600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff871660048201523360248201526044016108dc565b600b5461ffff90811690861610806110d6575060c861ffff8616115b1561112657600b546040517fa738697600000000000000000000000000000000000000000000000000000000815261ffff8088166004830152909116602482015260c860448201526064016108dc565b600b5463ffffffff620100009091048116908516111561118d57600b546040517ff5d7e01e00000000000000000000000000000000000000000000000000000000815263ffffffff80871660048301526201000090920490911660248201526044016108dc565b6101f463ffffffff841611156111df576040517f47386bec00000000000000000000000000000000000000000000000000000000815263ffffffff841660048201526101f460248201526044016108dc565b60006111ec826001615670565b6040805160208082018c9052338284015267ffffffffffffffff808c16606084015284166080808401919091528351808403909101815260a08301845280519082012060c083018d905260e080840182905284518085039091018152610100909301909352815191012091925081611262613667565b60408051602081019390935282015267ffffffffffffffff8a16606082015263ffffffff8089166080830152871660a08201523360c082015260e00160408051808303601f19018152828252805160209182012060008681526009835283902055848352820183905261ffff8a169082015263ffffffff808916606083015287166080820152339067ffffffffffffffff8b16908c907f63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a97729060a00160405180910390a45033600090815260026020908152604080832067ffffffffffffffff808d16855292529091208054919093167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009091161790915591505095945050505050565b600b546601000000000000900460ff16156113cc576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000908152600860205260409020546bffffffffffffffffffffffff80831691161015611426576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600090815260086020526040812080548392906114539084906bffffffffffffffffffffffff16615722565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555080600560088282829054906101000a90046bffffffffffffffffffffffff166114aa9190615722565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb83836040518363ffffffff1660e01b81526004016115489291906001600160a01b039290921682526bffffffffffffffffffffffff16602082015260400190565b602060405180830381600087803b15801561156257600080fd5b505af1158015611576573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159a9190614f44565b6115d0576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6115dc6131e6565b60408051808201825260009161160b919084906002908390839080828437600092019190915250612a16915050565b6000818152600660205260409020549091506001600160a01b031615611660576040517f4a0b8fa7000000000000000000000000000000000000000000000000000000008152600481018290526024016108dc565b600081815260066020908152604080832080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0388169081179091556007805460018101825594527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688909301849055518381527fe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b89101610ba3565b67ffffffffffffffff821660009081526003602052604090205482906001600160a01b031680611760576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336001600160a01b038216146117ad576040517fd8a3fb520000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016108dc565b600b546601000000000000900460ff16156117f4576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff84166000908152600360205260409020600201546064141561184b576040517f05a48e0f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038316600090815260026020908152604080832067ffffffffffffffff80891685529252909120541615611885576109e5565b6001600160a01b038316600081815260026020818152604080842067ffffffffffffffff8a1680865290835281852080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001908117909155600384528286209094018054948501815585529382902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001685179055905192835290917f43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e091016109dc565b6001546001600160a01b031633146119ab5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016108dc565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600b546601000000000000900460ff1615611a61576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600360205260409020546001600160a01b0316611aba576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600360205260409020600101546001600160a01b03163314611b425767ffffffffffffffff8116600090815260036020526040908190206001015490517fd084e9750000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016108dc565b67ffffffffffffffff81166000818152600360209081526040918290208054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001909301805490931690925583516001600160a01b03909116808252928101919091529092917f6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0910160405180910390a25050565b67ffffffffffffffff821660009081526003602052604090205482906001600160a01b031680611c3c576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336001600160a01b03821614611c89576040517fd8a3fb520000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016108dc565b600b546601000000000000900460ff1615611cd0576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611cd984612fb2565b15611d10576040517fb42f66e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038316600090815260026020908152604080832067ffffffffffffffff808916855292529091205416611d91576040517ff0019fe600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526001600160a01b03841660248201526044016108dc565b67ffffffffffffffff8416600090815260036020908152604080832060020180548251818502810185019093528083529192909190830182828015611dff57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611de1575b50505050509050600060018251611e16919061570b565b905060005b8251811015611f8e57856001600160a01b0316838281518110611e4057611e4061587d565b60200260200101516001600160a01b03161415611f7c576000838381518110611e6b57611e6b61587d565b6020026020010151905080600360008a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206002018381548110611eb157611eb161587d565b600091825260208083209190910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03949094169390931790925567ffffffffffffffff8a168152600390915260409020600201805480611f1e57611f1e61584e565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550611f8e565b80611f868161577b565b915050611e1b565b506001600160a01b038516600081815260026020908152604080832067ffffffffffffffff8b168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001690555192835290917f182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b91015b60405180910390a2505050505050565b600b546000906601000000000000900460ff1615612069576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005805467ffffffffffffffff16906000612083836157b4565b82546101009290920a67ffffffffffffffff8181021990931691831602179091556005541690506000806040519080825280602002602001820160405280156120d6578160200160208202803683370190505b506040805180820182526000808252602080830182815267ffffffffffffffff888116808552600484528685209551865493516bffffffffffffffffffffffff9091167fffffffffffffffffffffffff0000000000000000000000000000000000000000948516176c01000000000000000000000000919093160291909117909455845160608101865233815280830184815281870188815295855260038452959093208351815483166001600160a01b03918216178255955160018201805490931696169590951790559151805194955090936121ba9260028501920190614be5565b505060405133815267ffffffffffffffff841691507f464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf9060200160405180910390a250905090565b67ffffffffffffffff8116600090815260036020526040812054819081906060906001600160a01b0316612262576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80861660009081526004602090815260408083205460038352928190208054600290910180548351818602810186019094528084526bffffffffffffffffffffffff8616966c01000000000000000000000000909604909516946001600160a01b0390921693909291839183018282801561230f57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116122f1575b5050505050905093509350935093509193509193565b600b546601000000000000900460ff161561236c576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123ce576040517f44b0e3c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208114612408576040517f8129bbcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612416828401846151f6565b67ffffffffffffffff81166000908152600360205260409020549091506001600160a01b0316612472576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8116600090815260046020526040812080546bffffffffffffffffffffffff16918691906124a98385615693565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555084600560088282829054906101000a90046bffffffffffffffffffffffff166125009190615693565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055508167ffffffffffffffff167fd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f88287846125679190615658565b6040805192835260208301919091520161200f565b600b546000906601000000000000900460ff16156125c6576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005a905060008060006125da87876136f7565b9250925092506000866060015163ffffffff1667ffffffffffffffff811115612605576126056158ac565b60405190808252806020026020018201604052801561262e578160200160208202803683370190505b50905060005b876060015163ffffffff168110156126a25760408051602081018590529081018290526060016040516020818303038152906040528051906020012060001c8282815181106126855761268561587d565b60209081029190910101528061269a8161577b565b915050612634565b506000838152600960205260408082208290555181907f1fe543e300000000000000000000000000000000000000000000000000000000906126ea908790869060240161553c565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252600b80547fffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffff166601000000000000179055908a015160808b015191925060009161279a9163ffffffff169084613a18565b600b80547fffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffff1690556020808c01805167ffffffffffffffff9081166000908152600490935260408084205492518216845290922080549394506c01000000000000000000000000918290048316936001939192600c9261281e928692900416615670565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060006128758a600b600001600b9054906101000a900463ffffffff1663ffffffff1661286f85612a46565b3a613a64565b6020808e015167ffffffffffffffff166000908152600490915260409020549091506bffffffffffffffffffffffff808316911610156128e1576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020808d015167ffffffffffffffff166000908152600490915260408120805483929061291d9084906bffffffffffffffffffffffff16615722565b82546101009290920a6bffffffffffffffffffffffff81810219909316918316021790915560008b8152600660209081526040808320546001600160a01b03168352600890915281208054859450909261297991859116615693565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550877f7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e48883866040516129fc939291909283526bffffffffffffffffffffffff9190911660208301521515604082015260600190565b60405180910390a299505050505050505050505b92915050565b600081604051602001612a2991906153a2565b604051602081830303815290604052805190602001209050919050565b6040805161012081018252600c5463ffffffff80821683526401000000008204811660208401526801000000000000000082048116938301939093526c010000000000000000000000008104831660608301527001000000000000000000000000000000008104909216608082015262ffffff740100000000000000000000000000000000000000008304811660a08301819052770100000000000000000000000000000000000000000000008404821660c08401527a0100000000000000000000000000000000000000000000000000008404821660e08401527d0100000000000000000000000000000000000000000000000000000000009093041661010082015260009167ffffffffffffffff841611612b64575192915050565b8267ffffffffffffffff168160a0015162ffffff16108015612b9957508060c0015162ffffff168367ffffffffffffffff1611155b15612ba8576020015192915050565b8267ffffffffffffffff168160c0015162ffffff16108015612bdd57508060e0015162ffffff168367ffffffffffffffff1611155b15612bec576040015192915050565b8267ffffffffffffffff168160e0015162ffffff16108015612c22575080610100015162ffffff168367ffffffffffffffff1611155b15612c31576060015192915050565b6080015192915050565b67ffffffffffffffff821660009081526003602052604090205482906001600160a01b031680612c97576040517f1f6a65b600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336001600160a01b03821614612ce4576040517fd8a3fb520000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016108dc565b600b546601000000000000900460ff1615612d2b576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612d3484612fb2565b15612d6b576040517fb42f66e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109e58484613242565b612d7d6131e6565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a082319060240160206040518083038186803b158015612df857600080fd5b505afa158015612e0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e309190614f66565b6005549091506801000000000000000090046bffffffffffffffffffffffff1681811115612e94576040517fa99da30200000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016108dc565b81811015612fad576000612ea8828461570b565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb90604401602060405180830381600087803b158015612f3057600080fd5b505af1158015612f44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f689190614f44565b50604080516001600160a01b0386168152602081018390527f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600910160405180910390a1505b505050565b67ffffffffffffffff81166000908152600360209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152879693958601939092919083018282801561304757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613029575b505050505081525050905060005b8160400151518110156131cb5760005b6007548110156131b8576000613181600783815481106130875761308761587d565b9060005260206000200154856040015185815181106130a8576130a861587d565b60200260200101518860026000896040015189815181106130cb576130cb61587d565b6020908102919091018101516001600160a01b03168252818101929092526040908101600090812067ffffffffffffffff808f16835293522054166040805160208082018790526001600160a01b03959095168183015267ffffffffffffffff9384166060820152919092166080808301919091528251808303909101815260a08201835280519084012060c082019490945260e080820185905282518083039091018152610100909101909152805191012091565b50600081815260096020526040902054909150156131a55750600195945050505050565b50806131b08161577b565b915050613065565b50806131c38161577b565b915050613055565b5060009392505050565b6131dd6131e6565b61083481613bbc565b6000546001600160a01b031633146132405760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016108dc565b565b600b546601000000000000900460ff1615613289576040517fed3ba6a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff82166000908152600360209081526040808320815160608101835281546001600160a01b0390811682526001830154168185015260028201805484518187028101870186528181529295939486019383018282801561331a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116132fc575b5050509190925250505067ffffffffffffffff80851660009081526004602090815260408083208151808301909252546bffffffffffffffffffffffff81168083526c01000000000000000000000000909104909416918101919091529293505b8360400151518110156134145760026000856040015183815181106133a2576133a261587d565b6020908102919091018101516001600160a01b03168252818101929092526040908101600090812067ffffffffffffffff8a168252909252902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001690558061340c8161577b565b91505061337b565b5067ffffffffffffffff8516600090815260036020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000908116825560018201805490911690559061346f6002830182614c62565b505067ffffffffffffffff8516600090815260046020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055600580548291906008906134df9084906801000000000000000090046bffffffffffffffffffffffff16615722565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb85836bffffffffffffffffffffffff166040518363ffffffff1660e01b815260040161357d9291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b15801561359757600080fd5b505af11580156135ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135cf9190614f44565b613605576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516001600160a01b03861681526bffffffffffffffffffffffff8316602082015267ffffffffffffffff8716917fe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815910160405180910390a25050505050565b60004661367381613c7e565b156136f05760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156136b257600080fd5b505afa1580156136c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136ea9190614f66565b91505090565b4391505090565b6040805180820182526000918291829161372a919087906002908390839080828437600092019190915250612a16915050565b6000818152600660205260409020549093506001600160a01b03168061377f576040517f77f5b84c000000000000000000000000000000000000000000000000000000008152600481018590526024016108dc565b838660c0013560405160200161379f929190918252602082015260400190565b60408051601f19818403018152918152815160209283012060008181526009909352912054909350806137fe576040517f3688124a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85516020808801516040808a015160608b015160808c0151925161386a968b96909594910195865267ffffffffffffffff948516602087015292909316604085015263ffffffff90811660608501529190911660808301526001600160a01b031660a082015260c00190565b6040516020818303038152906040528051906020012081146138b8576040517fd529142c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006138c78760000151613ca1565b9050806139d35786516040517fe9413d3800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e9413d389060240160206040518083038186803b15801561395357600080fd5b505afa158015613967573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061398b9190614f66565b9050806139d35786516040517f175dadad00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108dc565b6040805160c08a013560208083019190915281830184905282518083038401815260609092019092528051910120613a0b8982613da7565b9450505050509250925092565b60005a611388811015613a2a57600080fd5b611388810390508460408204820311613a4257600080fd5b50823b613a4e57600080fd5b60008083516020850160008789f1949350505050565b600080613a6f613e9d565b905060008113613aae576040517f43d4cf66000000000000000000000000000000000000000000000000000000008152600481018290526024016108dc565b6000613af06000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613fa492505050565b9050600082825a613b018b8b615658565b613b0b919061570b565b613b1590886156ce565b613b1f9190615658565b613b3190670de0b6b3a76400006156ce565b613b3b91906156ba565b90506000613b5463ffffffff881664e8d4a510006156ce565b9050613b6c816b033b2e3c9fd0803ce800000061570b565b821115613ba5576040517fe80fa38100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613baf8183615658565b9998505050505050505050565b6001600160a01b038116331415613c155760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016108dc565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061a4b1821480613c92575062066eed82145b80612a1057505062066eee1490565b600046613cad81613c7e565b15613d97576101008367ffffffffffffffff16613cc8613667565b613cd2919061570b565b1180613cef5750613ce1613667565b8367ffffffffffffffff1610155b15613cfd5750600092915050565b6040517f2b407a8200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152606490632b407a82906024015b60206040518083038186803b158015613d5857600080fd5b505afa158015613d6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d909190614f66565b9392505050565b505067ffffffffffffffff164090565b604080518082018252600091613e679190859060029083908390808284376000920191909152505060408051808201825291508087019060029083908390808284376000920191909152505050608086013560a087013586613e106101008a0160e08b01614d82565b604080518082018252906101008c019060029083908390808284376000920191909152505060408051808201825291506101408d0190600290839083908082843760009201919091525050506101808c013561407f565b600383604001604051602001613e7e929190615522565b60408051601f1981840301815291905280516020909101209392505050565b600b54604080517ffeaf968c0000000000000000000000000000000000000000000000000000000081529051600092670100000000000000900463ffffffff169182151591849182917f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163feaf968c9160048083019260a0929190829003018186803b158015613f3657600080fd5b505afa158015613f4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f6e919061523b565b509450909250849150508015613f925750613f89824261570b565b8463ffffffff16105b15613f9c5750600a545b949350505050565b600046613fb081613c7e565b15613fef57606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b158015613d5857600080fd5b613ff8816142ba565b156140765773420000000000000000000000000000000000000f6001600160a01b03166349948e0e846040518060800160405280604881526020016158dc6048913960405160200161404b9291906152e0565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401613d4091906153d3565b50600092915050565b61408889614301565b6140d45760405162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e20637572766500000000000060448201526064016108dc565b6140dd88614301565b6141295760405162461bcd60e51b815260206004820152601560248201527f67616d6d61206973206e6f74206f6e206375727665000000000000000000000060448201526064016108dc565b61413283614301565b61417e5760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e20637572766500000060448201526064016108dc565b61418782614301565b6141d35760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e2063757276650000000060448201526064016108dc565b6141df878a88876143da565b61422b5760405162461bcd60e51b815260206004820152601960248201527f6164647228632a706b2b732a6729213d5f755769746e6573730000000000000060448201526064016108dc565b60006142378a8761452b565b9050600061424a898b878b86898961458f565b9050600061425b838d8d8a866146bb565b9050808a146142ac5760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642070726f6f660000000000000000000000000000000000000060448201526064016108dc565b505050505050505050505050565b6000600a8214806142cc57506101a482145b806142d9575062aa37dc82145b806142e5575061210582145b806142f2575062014a3382145b80612a1057505062014a341490565b80516000906401000003d0191161435a5760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420782d6f7264696e617465000000000000000000000000000060448201526064016108dc565b60208201516401000003d019116143b35760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420792d6f7264696e617465000000000000000000000000000060448201526064016108dc565b60208201516401000003d0199080096143d38360005b60200201516146fb565b1492915050565b60006001600160a01b0382166144325760405162461bcd60e51b815260206004820152600b60248201527f626164207769746e65737300000000000000000000000000000000000000000060448201526064016108dc565b60208401516000906001161561444957601c61444c565b601b5b905060007ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03641418587600060200201510986517ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa158015614503573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b614533614c80565b6145606001848460405160200161454c93929190615381565b60405160208183030381529060405261471f565b90505b61456c81614301565b612a10578051604080516020810192909252614588910161454c565b9050614563565b614597614c80565b825186516401000003d01990819006910614156145f65760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e6374000060448201526064016108dc565b61460187898861476e565b61464d5760405162461bcd60e51b815260206004820152601660248201527f4669727374206d756c20636865636b206661696c65640000000000000000000060448201526064016108dc565b61465884868561476e565b6146a45760405162461bcd60e51b815260206004820152601760248201527f5365636f6e64206d756c20636865636b206661696c656400000000000000000060448201526064016108dc565b6146af8684846148b6565b98975050505050505050565b6000600286868685876040516020016146d99695949392919061530f565b60408051601f1981840301815291905280516020909101209695505050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b614727614c80565b6147308261497d565b81526147456147408260006143c9565b6149b8565b602082018190526002900660011415614769576020810180516401000003d0190390525b919050565b6000826147bd5760405162461bcd60e51b815260206004820152600b60248201527f7a65726f207363616c617200000000000000000000000000000000000000000060448201526064016108dc565b835160208501516000906147d3906002906157dc565b156147df57601c6147e2565b601b5b905060007ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03641418387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa158015614862573d6000803e3d6000fd5b50505060206040510351905060008660405160200161488191906152ce565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b6148be614c80565b8351602080860151855191860151600093849384936148df939091906149d8565b919450925090506401000003d01985820960011461493f5760405162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a0000000000000060448201526064016108dc565b60405180604001604052806401000003d0198061495e5761495e61581f565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d019811061476957604080516020808201939093528151808203840181529082019091528051910120614985565b6000612a108260026149d16401000003d0196001615658565b901c614ab8565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a0890506000614a1883838585614b78565b9098509050614a2988828e88614b9c565b9098509050614a3a88828c87614b9c565b90985090506000614a4d8d878b85614b9c565b9098509050614a5e88828686614b78565b9098509050614a6f88828e89614b9c565b9098509050818114614aa4576401000003d019818a0998506401000003d01982890997506401000003d0198183099650614aa8565b8196505b5050505050509450945094915050565b600080614ac3614c9e565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a0820152614af5614cbc565b60208160c08460057ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa925082614b6e5760405162461bcd60e51b815260206004820152601260248201527f6269674d6f64457870206661696c75726521000000000000000000000000000060448201526064016108dc565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b828054828255906000526020600020908101928215614c52579160200282015b82811115614c5257825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190614c05565b50614c5e929150614cda565b5090565b50805460008255906000526020600020908101906108349190614cda565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b80821115614c5e5760008155600101614cdb565b80356001600160a01b038116811461476957600080fd5b8060408101831015612a1057600080fd5b803561ffff8116811461476957600080fd5b803562ffffff8116811461476957600080fd5b803563ffffffff8116811461476957600080fd5b803567ffffffffffffffff8116811461476957600080fd5b805169ffffffffffffffffffff8116811461476957600080fd5b600060208284031215614d9457600080fd5b613d9082614cef565b60008060608385031215614db057600080fd5b614db983614cef565b9150614dc88460208501614d06565b90509250929050565b60008060008060608587031215614de757600080fd5b614df085614cef565b935060208501359250604085013567ffffffffffffffff80821115614e1457600080fd5b818701915087601f830112614e2857600080fd5b813581811115614e3757600080fd5b886020828501011115614e4957600080fd5b95989497505060200194505050565b60008060408385031215614e6b57600080fd5b614e7483614cef565b915060208301356bffffffffffffffffffffffff81168114614e9557600080fd5b809150509250929050565b600060408284031215614eb257600080fd5b613d908383614d06565b600060408284031215614ece57600080fd5b82601f830112614edd57600080fd5b6040516040810181811067ffffffffffffffff82111715614f0057614f006158ac565b8060405250808385604086011115614f1757600080fd5b60005b6002811015614f39578135835260209283019290910190600101614f1a565b509195945050505050565b600060208284031215614f5657600080fd5b81518015158114613d9057600080fd5b600060208284031215614f7857600080fd5b5051919050565b600080600080600060a08688031215614f9757600080fd5b85359450614fa760208701614d50565b9350614fb560408701614d17565b9250614fc360608701614d3c565b9150614fd160808701614d3c565b90509295509295909350565b600080828403610240811215614ff257600080fd5b6101a08082121561500257600080fd5b84935060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe608301121561503557600080fd5b61503d61560b565b915061504a818601614d50565b82525061505a6101c08501614d50565b602082015261506c6101e08501614d3c565b604082015261507e6102008501614d3c565b60608201526150906102208501614cef565b6080820152809150509250929050565b6000806000806000808688036101c08112156150bb57600080fd5b6150c488614d17565b96506150d260208901614d3c565b95506150e060408901614d3c565b94506150ee60608901614d3c565b935060808801359250610120807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608301121561512957600080fd5b615131615634565b915061513f60a08a01614d3c565b825261514d60c08a01614d3c565b602083015261515e60e08a01614d3c565b6040830152610100615171818b01614d3c565b6060840152615181828b01614d3c565b60808401526151936101408b01614d29565b60a08401526151a56101608b01614d29565b60c08401526151b76101808b01614d29565b60e08401526151c96101a08b01614d29565b818401525050809150509295509295509295565b6000602082840312156151ef57600080fd5b5035919050565b60006020828403121561520857600080fd5b613d9082614d50565b6000806040838503121561522457600080fd5b61522d83614d50565b9150614dc860208401614cef565b600080600080600060a0868803121561525357600080fd5b61525c86614d68565b9450602086015193506040860151925060608601519150614fd160808701614d68565b8060005b60028110156109e5578151845260209384019390910190600101615283565b600081518084526152ba81602086016020860161574f565b601f01601f19169290920160200192915050565b6152d8818361527f565b604001919050565b600083516152f281846020880161574f565b83519083019061530681836020880161574f565b01949350505050565b86815261531f602082018761527f565b61532c606082018661527f565b61533960a082018561527f565b61534660e082018461527f565b60609190911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166101208201526101340195945050505050565b838152615391602082018461527f565b606081019190915260800192915050565b60408101818360005b60028110156153ca5781518352602092830192909101906001016153ab565b50505092915050565b602081526000613d9060208301846152a2565b60006060820161ffff86168352602063ffffffff86168185015260606040850152818551808452608086019150828701935060005b818110156154375784518352938301939183019160010161541b565b509098975050505050505050565b60006101c08201905061ffff8816825263ffffffff808816602084015280871660408401528086166060840152846080840152835481811660a085015261549960c08501838360201c1663ffffffff169052565b6154b060e08501838360401c1663ffffffff169052565b6154c86101008501838360601c1663ffffffff169052565b6154e06101208501838360801c1663ffffffff169052565b62ffffff60a082901c811661014086015260b882901c811661016086015260d082901c1661018085015260e81c6101a090930192909252979650505050505050565b828152606081016040836020840137600081529392505050565b6000604082018483526020604081850152818551808452606086019150828701935060005b8181101561557d57845183529383019391830191600101615561565b5090979650505050505050565b6000608082016bffffffffffffffffffffffff87168352602067ffffffffffffffff8716818501526001600160a01b0380871660408601526080606086015282865180855260a087019150838801945060005b818110156155fb5785518416835294840194918401916001016155dd565b50909a9950505050505050505050565b60405160a0810167ffffffffffffffff8111828210171561562e5761562e6158ac565b60405290565b604051610120810167ffffffffffffffff8111828210171561562e5761562e6158ac565b6000821982111561566b5761566b6157f0565b500190565b600067ffffffffffffffff808316818516808303821115615306576153066157f0565b60006bffffffffffffffffffffffff808316818516808303821115615306576153066157f0565b6000826156c9576156c961581f565b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615615706576157066157f0565b500290565b60008282101561571d5761571d6157f0565b500390565b60006bffffffffffffffffffffffff83811690831681811015615747576157476157f0565b039392505050565b60005b8381101561576a578181015183820152602001615752565b838111156109e55750506000910152565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156157ad576157ad6157f0565b5060010190565b600067ffffffffffffffff808316818114156157d2576157d26157f0565b6001019392505050565b6000826157eb576157eb61581f565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfe307866666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666a164736f6c6343000806000a",
}

var VRFCoordinatorV2ABI = VRFCoordinatorV2MetaData.ABI

var VRFCoordinatorV2Bin = VRFCoordinatorV2MetaData.Bin

func DeployVRFCoordinatorV2(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, blockhashStore common.Address, linkEthFeed common.Address) (common.Address, *types.Transaction, *VRFCoordinatorV2, error) {
	parsed, err := VRFCoordinatorV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV2Bin), backend, link, blockhashStore, linkEthFeed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV2{address: address, abi: *parsed, VRFCoordinatorV2Caller: VRFCoordinatorV2Caller{contract: contract}, VRFCoordinatorV2Transactor: VRFCoordinatorV2Transactor{contract: contract}, VRFCoordinatorV2Filterer: VRFCoordinatorV2Filterer{contract: contract}}, nil
}

type VRFCoordinatorV2 struct {
	address common.Address
	abi     abi.ABI
	VRFCoordinatorV2Caller
	VRFCoordinatorV2Transactor
	VRFCoordinatorV2Filterer
}

type VRFCoordinatorV2Caller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2Transactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2Filterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2Session struct {
	Contract     *VRFCoordinatorV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2CallerSession struct {
	Contract *VRFCoordinatorV2Caller
	CallOpts bind.CallOpts
}

type VRFCoordinatorV2TransactorSession struct {
	Contract     *VRFCoordinatorV2Transactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2Raw struct {
	Contract *VRFCoordinatorV2
}

type VRFCoordinatorV2CallerRaw struct {
	Contract *VRFCoordinatorV2Caller
}

type VRFCoordinatorV2TransactorRaw struct {
	Contract *VRFCoordinatorV2Transactor
}

func NewVRFCoordinatorV2(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2, error) {
	abi, err := abi.JSON(strings.NewReader(VRFCoordinatorV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFCoordinatorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2{address: address, abi: abi, VRFCoordinatorV2Caller: VRFCoordinatorV2Caller{contract: contract}, VRFCoordinatorV2Transactor: VRFCoordinatorV2Transactor{contract: contract}, VRFCoordinatorV2Filterer: VRFCoordinatorV2Filterer{contract: contract}}, nil
}

func NewVRFCoordinatorV2Caller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2Caller, error) {
	contract, err := bindVRFCoordinatorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Caller{contract: contract}, nil
}

func NewVRFCoordinatorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2Transactor, error) {
	contract, err := bindVRFCoordinatorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Transactor{contract: contract}, nil
}

func NewVRFCoordinatorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2Filterer, error) {
	contract, err := bindVRFCoordinatorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Filterer{contract: contract}, nil
}

func bindVRFCoordinatorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Caller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Transactor.contract.Transfer(opts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Transactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.contract.Transfer(opts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) LINK() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.LINK(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.LINK(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.LINKETHFEED(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.LINKETHFEED(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2.Contract.MAXCONSUMERS(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2.Contract.MAXCONSUMERS(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2.Contract.MAXNUMWORDS(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2.Contract.MAXNUMWORDS(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetCommitment(opts *bind.CallOpts, requestId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getCommitment", requestId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetCommitment(requestId *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2.Contract.GetCommitment(&_VRFCoordinatorV2.CallOpts, requestId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetCommitment(requestId *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2.Contract.GetCommitment(&_VRFCoordinatorV2.CallOpts, requestId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetConfig(opts *bind.CallOpts) (GetConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getConfig")

	outstruct := new(GetConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.StalenessSeconds = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetConfig() (GetConfig,

	error) {
	return _VRFCoordinatorV2.Contract.GetConfig(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetConfig() (GetConfig,

	error) {
	return _VRFCoordinatorV2.Contract.GetConfig(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetCurrentSubId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getCurrentSubId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetCurrentSubId() (uint64, error) {
	return _VRFCoordinatorV2.Contract.GetCurrentSubId(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetCurrentSubId() (uint64, error) {
	return _VRFCoordinatorV2.Contract.GetCurrentSubId(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getFallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.GetFallbackWeiPerUnitLink(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.GetFallbackWeiPerUnitLink(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetFeeConfig(opts *bind.CallOpts) (GetFeeConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getFeeConfig")

	outstruct := new(GetFeeConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FulfillmentFlatFeeLinkPPMTier1 = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier2 = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier3 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier4 = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPMTier5 = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ReqsForTier2 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ReqsForTier3 = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ReqsForTier4 = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ReqsForTier5 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetFeeConfig() (GetFeeConfig,

	error) {
	return _VRFCoordinatorV2.Contract.GetFeeConfig(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetFeeConfig() (GetFeeConfig,

	error) {
	return _VRFCoordinatorV2.Contract.GetFeeConfig(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetFeeTier(opts *bind.CallOpts, reqCount uint64) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getFeeTier", reqCount)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetFeeTier(reqCount uint64) (uint32, error) {
	return _VRFCoordinatorV2.Contract.GetFeeTier(&_VRFCoordinatorV2.CallOpts, reqCount)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetFeeTier(reqCount uint64) (uint32, error) {
	return _VRFCoordinatorV2.Contract.GetFeeTier(&_VRFCoordinatorV2.CallOpts, reqCount)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2.Contract.GetRequestConfig(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2.Contract.GetRequestConfig(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetSubscription(opts *bind.CallOpts, subId uint64) (GetSubscription,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(GetSubscription)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetSubscription(subId uint64) (GetSubscription,

	error) {
	return _VRFCoordinatorV2.Contract.GetSubscription(&_VRFCoordinatorV2.CallOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetSubscription(subId uint64) (GetSubscription,

	error) {
	return _VRFCoordinatorV2.Contract.GetSubscription(&_VRFCoordinatorV2.CallOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getTotalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetTotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.GetTotalBalance(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetTotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.GetTotalBalance(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2.Contract.HashOfKey(&_VRFCoordinatorV2.CallOpts, publicKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2.Contract.HashOfKey(&_VRFCoordinatorV2.CallOpts, publicKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) Owner() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.Owner(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.Owner(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) PendingRequestExists(subId uint64) (bool, error) {
	return _VRFCoordinatorV2.Contract.PendingRequestExists(&_VRFCoordinatorV2.CallOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) PendingRequestExists(subId uint64) (bool, error) {
	return _VRFCoordinatorV2.Contract.PendingRequestExists(&_VRFCoordinatorV2.CallOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) TypeAndVersion() (string, error) {
	return _VRFCoordinatorV2.Contract.TypeAndVersion(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) TypeAndVersion() (string, error) {
	return _VRFCoordinatorV2.Contract.TypeAndVersion(&_VRFCoordinatorV2.CallOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AcceptOwnership(&_VRFCoordinatorV2.TransactOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AcceptOwnership(&_VRFCoordinatorV2.TransactOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2.TransactOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2.TransactOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AddConsumer(&_VRFCoordinatorV2.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AddConsumer(&_VRFCoordinatorV2.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.CancelSubscription(&_VRFCoordinatorV2.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.CancelSubscription(&_VRFCoordinatorV2.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.CreateSubscription(&_VRFCoordinatorV2.TransactOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.CreateSubscription(&_VRFCoordinatorV2.TransactOpts)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.DeregisterProvingKey(&_VRFCoordinatorV2.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.DeregisterProvingKey(&_VRFCoordinatorV2.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.FulfillRandomWords(&_VRFCoordinatorV2.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.FulfillRandomWords(&_VRFCoordinatorV2.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.OnTokenTransfer(&_VRFCoordinatorV2.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.OnTokenTransfer(&_VRFCoordinatorV2.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.OracleWithdraw(&_VRFCoordinatorV2.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.OracleWithdraw(&_VRFCoordinatorV2.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "ownerCancelSubscription", subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2.TransactOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2.TransactOpts, subId)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "recoverFunds", to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RecoverFunds(&_VRFCoordinatorV2.TransactOpts, to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RecoverFunds(&_VRFCoordinatorV2.TransactOpts, to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "registerProvingKey", oracle, publicProvingKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RegisterProvingKey(&_VRFCoordinatorV2.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RegisterProvingKey(&_VRFCoordinatorV2.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RemoveConsumer(&_VRFCoordinatorV2.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RemoveConsumer(&_VRFCoordinatorV2.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "requestRandomWords", keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestRandomWords(&_VRFCoordinatorV2.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestRandomWords(&_VRFCoordinatorV2.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.SetConfig(&_VRFCoordinatorV2.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.SetConfig(&_VRFCoordinatorV2.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.TransferOwnership(&_VRFCoordinatorV2.TransactOpts, to)
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.TransferOwnership(&_VRFCoordinatorV2.TransactOpts, to)
}

type VRFCoordinatorV2ConfigSetIterator struct {
	Event *VRFCoordinatorV2ConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2ConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2ConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2ConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2ConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2ConfigSet struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
	FallbackWeiPerUnitLink      *big.Int
	FeeConfig                   VRFCoordinatorV2FeeConfig
	Raw                         types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2ConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2ConfigSetIterator{contract: _VRFCoordinatorV2.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2ConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2ConfigSet)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseConfigSet(log types.Log) (*VRFCoordinatorV2ConfigSet, error) {
	event := new(VRFCoordinatorV2ConfigSet)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2FundsRecoveredIterator struct {
	Event *VRFCoordinatorV2FundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2FundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2FundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2FundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2FundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2FundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2FundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2FundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2FundsRecoveredIterator{contract: _VRFCoordinatorV2.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2FundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2FundsRecovered)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2FundsRecovered, error) {
	event := new(VRFCoordinatorV2FundsRecovered)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2OwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV2OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2OwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2OwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2OwnershipTransferRequestedIterator{contract: _VRFCoordinatorV2.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2OwnershipTransferRequested)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2OwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV2OwnershipTransferRequested)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2OwnershipTransferredIterator struct {
	Event *VRFCoordinatorV2OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2OwnershipTransferredIterator{contract: _VRFCoordinatorV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2OwnershipTransferred)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2OwnershipTransferred, error) {
	event := new(VRFCoordinatorV2OwnershipTransferred)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2ProvingKeyDeregisteredIterator struct {
	Event *VRFCoordinatorV2ProvingKeyDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2ProvingKeyDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2ProvingKeyDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2ProvingKeyDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2ProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2ProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2ProvingKeyDeregistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2ProvingKeyDeregisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2ProvingKeyDeregisteredIterator{contract: _VRFCoordinatorV2.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2ProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2ProvingKeyDeregistered)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV2ProvingKeyDeregistered, error) {
	event := new(VRFCoordinatorV2ProvingKeyDeregistered)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2ProvingKeyRegisteredIterator struct {
	Event *VRFCoordinatorV2ProvingKeyRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2ProvingKeyRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2ProvingKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2ProvingKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2ProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2ProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2ProvingKeyRegistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2ProvingKeyRegisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2ProvingKeyRegisteredIterator{contract: _VRFCoordinatorV2.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2ProvingKeyRegistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2ProvingKeyRegistered)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2ProvingKeyRegistered, error) {
	event := new(VRFCoordinatorV2ProvingKeyRegistered)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2RandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV2RandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2RandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2RandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2RandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2RandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2RandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2RandomWordsFulfilled struct {
	RequestId  *big.Int
	OutputSeed *big.Int
	Payment    *big.Int
	Success    bool
	Raw        types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV2RandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2RandomWordsFulfilledIterator{contract: _VRFCoordinatorV2.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2RandomWordsFulfilled)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2RandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV2RandomWordsFulfilled)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2RandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV2RandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2RandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2RandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2RandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2RandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2RandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2RandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       uint64
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	Sender                      common.Address
	Raw                         types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VRFCoordinatorV2RandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2RandomWordsRequestedIterator{contract: _VRFCoordinatorV2.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2RandomWordsRequested)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2RandomWordsRequested, error) {
	event := new(VRFCoordinatorV2RandomWordsRequested)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionCanceledIterator struct {
	Event *VRFCoordinatorV2SubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionCanceled struct {
	SubId  uint64
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionCanceledIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionCanceled, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionCanceled)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2SubscriptionCanceled, error) {
	event := new(VRFCoordinatorV2SubscriptionCanceled)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorV2SubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionConsumerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionConsumerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionConsumerAdded struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionConsumerAddedIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionConsumerAdded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionConsumerAdded)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2SubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorV2SubscriptionConsumerAdded)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorV2SubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionConsumerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionConsumerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionConsumerRemoved struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionConsumerRemovedIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionConsumerRemoved)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2SubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorV2SubscriptionConsumerRemoved)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionCreatedIterator struct {
	Event *VRFCoordinatorV2SubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionCreated struct {
	SubId uint64
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionCreatedIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionCreated, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionCreated)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2SubscriptionCreated, error) {
	event := new(VRFCoordinatorV2SubscriptionCreated)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionFundedIterator struct {
	Event *VRFCoordinatorV2SubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionFunded struct {
	SubId      uint64
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionFundedIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionFunded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionFunded)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2SubscriptionFunded, error) {
	event := new(VRFCoordinatorV2SubscriptionFunded)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorV2SubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionOwnerTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionOwnerTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionOwnerTransferRequested struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionOwnerTransferRequested)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2SubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorV2SubscriptionOwnerTransferRequested)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2SubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorV2SubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2SubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2SubscriptionOwnerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2SubscriptionOwnerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2SubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2SubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2SubscriptionOwnerTransferred struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2SubscriptionOwnerTransferredIterator{contract: _VRFCoordinatorV2.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2SubscriptionOwnerTransferred)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2SubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorV2SubscriptionOwnerTransferred)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetConfig struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}
type GetFeeConfig struct {
	FulfillmentFlatFeeLinkPPMTier1 uint32
	FulfillmentFlatFeeLinkPPMTier2 uint32
	FulfillmentFlatFeeLinkPPMTier3 uint32
	FulfillmentFlatFeeLinkPPMTier4 uint32
	FulfillmentFlatFeeLinkPPMTier5 uint32
	ReqsForTier2                   *big.Int
	ReqsForTier3                   *big.Int
	ReqsForTier4                   *big.Int
	ReqsForTier5                   *big.Int
}
type GetSubscription struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFCoordinatorV2.abi.Events["ConfigSet"].ID:
		return _VRFCoordinatorV2.ParseConfigSet(log)
	case _VRFCoordinatorV2.abi.Events["FundsRecovered"].ID:
		return _VRFCoordinatorV2.ParseFundsRecovered(log)
	case _VRFCoordinatorV2.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFCoordinatorV2.ParseOwnershipTransferRequested(log)
	case _VRFCoordinatorV2.abi.Events["OwnershipTransferred"].ID:
		return _VRFCoordinatorV2.ParseOwnershipTransferred(log)
	case _VRFCoordinatorV2.abi.Events["ProvingKeyDeregistered"].ID:
		return _VRFCoordinatorV2.ParseProvingKeyDeregistered(log)
	case _VRFCoordinatorV2.abi.Events["ProvingKeyRegistered"].ID:
		return _VRFCoordinatorV2.ParseProvingKeyRegistered(log)
	case _VRFCoordinatorV2.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFCoordinatorV2.ParseRandomWordsFulfilled(log)
	case _VRFCoordinatorV2.abi.Events["RandomWordsRequested"].ID:
		return _VRFCoordinatorV2.ParseRandomWordsRequested(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionCanceled"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionCanceled(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionConsumerAdded"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionConsumerAdded(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionConsumerRemoved"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionConsumerRemoved(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionCreated"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionCreated(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionFunded"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionFunded(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionOwnerTransferRequested"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionOwnerTransferRequested(log)
	case _VRFCoordinatorV2.abi.Events["SubscriptionOwnerTransferred"].ID:
		return _VRFCoordinatorV2.ParseSubscriptionOwnerTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFCoordinatorV2ConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc21e3bd2e0b339d2848f0dd956947a88966c242c0c0c582a33137a5c1ceb5cb2")
}

func (VRFCoordinatorV2FundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600")
}

func (VRFCoordinatorV2OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFCoordinatorV2OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFCoordinatorV2ProvingKeyDeregistered) Topic() common.Hash {
	return common.HexToHash("0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d")
}

func (VRFCoordinatorV2ProvingKeyRegistered) Topic() common.Hash {
	return common.HexToHash("0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8")
}

func (VRFCoordinatorV2RandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4")
}

func (VRFCoordinatorV2RandomWordsRequested) Topic() common.Hash {
	return common.HexToHash("0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772")
}

func (VRFCoordinatorV2SubscriptionCanceled) Topic() common.Hash {
	return common.HexToHash("0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815")
}

func (VRFCoordinatorV2SubscriptionConsumerAdded) Topic() common.Hash {
	return common.HexToHash("0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0")
}

func (VRFCoordinatorV2SubscriptionConsumerRemoved) Topic() common.Hash {
	return common.HexToHash("0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b")
}

func (VRFCoordinatorV2SubscriptionCreated) Topic() common.Hash {
	return common.HexToHash("0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf")
}

func (VRFCoordinatorV2SubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8")
}

func (VRFCoordinatorV2SubscriptionOwnerTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be")
}

func (VRFCoordinatorV2SubscriptionOwnerTransferred) Topic() common.Hash {
	return common.HexToHash("0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0")
}

func (_VRFCoordinatorV2 *VRFCoordinatorV2) Address() common.Address {
	return _VRFCoordinatorV2.address
}

type VRFCoordinatorV2Interface interface {
	BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKETHFEED(opts *bind.CallOpts) (common.Address, error)

	MAXCONSUMERS(opts *bind.CallOpts) (uint16, error)

	MAXNUMWORDS(opts *bind.CallOpts) (uint32, error)

	MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error)

	GetCommitment(opts *bind.CallOpts, requestId *big.Int) ([32]byte, error)

	GetConfig(opts *bind.CallOpts) (GetConfig,

		error)

	GetCurrentSubId(opts *bind.CallOpts) (uint64, error)

	GetFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error)

	GetFeeConfig(opts *bind.CallOpts) (GetFeeConfig,

		error)

	GetFeeTier(opts *bind.CallOpts, reqCount uint64) (uint32, error)

	GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error)

	GetSubscription(opts *bind.CallOpts, subId uint64) (GetSubscription,

		error)

	GetTotalBalance(opts *bind.CallOpts) (*big.Int, error)

	HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error)

	CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error)

	CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error)

	DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2RequestCommitment) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error)

	RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2FeeConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFCoordinatorV2ConfigSet, error)

	FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2FundsRecoveredIterator, error)

	WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2FundsRecovered) (event.Subscription, error)

	ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2FundsRecovered, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2OwnershipTransferred, error)

	FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2ProvingKeyDeregisteredIterator, error)

	WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2ProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV2ProvingKeyDeregistered, error)

	FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2ProvingKeyRegisteredIterator, error)

	WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2ProvingKeyRegistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2ProvingKeyRegistered, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV2RandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2RandomWordsFulfilled, error)

	FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VRFCoordinatorV2RandomWordsRequestedIterator, error)

	WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error)

	ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2RandomWordsRequested, error)

	FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionCanceledIterator, error)

	WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionCanceled, subId []uint64) (event.Subscription, error)

	ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2SubscriptionCanceled, error)

	FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionConsumerAddedIterator, error)

	WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionConsumerAdded, subId []uint64) (event.Subscription, error)

	ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2SubscriptionConsumerAdded, error)

	FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionConsumerRemovedIterator, error)

	WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error)

	ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2SubscriptionConsumerRemoved, error)

	FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionCreatedIterator, error)

	WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionCreated, subId []uint64) (event.Subscription, error)

	ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2SubscriptionCreated, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionFunded, subId []uint64) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2SubscriptionFunded, error)

	FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionOwnerTransferRequestedIterator, error)

	WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error)

	ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2SubscriptionOwnerTransferRequested, error)

	FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2SubscriptionOwnerTransferredIterator, error)

	WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2SubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error)

	ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2SubscriptionOwnerTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
