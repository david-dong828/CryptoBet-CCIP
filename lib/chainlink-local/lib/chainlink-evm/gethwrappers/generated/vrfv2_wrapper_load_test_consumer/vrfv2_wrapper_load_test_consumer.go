// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2_wrapper_load_test_consumer

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

var VRFV2WrapperLoadTestConsumerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_link\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vrfV2Wrapper\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"requestTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilmentTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilmentBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_vrfV2Wrapper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractVRFV2WrapperInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"makeRequests\",\"inputs\":[{\"name\":\"_callbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_requestConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_numWords\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_requestCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rawFulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reset\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_averageFulfillmentInMillions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_fastestFulfillment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requestCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_requests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"paid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requestTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilmentTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilmentBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_responseCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_slowestFulfillment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLink\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WrappedRequestFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"payment\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WrapperRequestMade\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"paid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e0604052600060045560006005556103e76006553480156200002157600080fd5b50604051620017ce380380620017ce8339810160408190526200004491620001cb565b6001600160601b0319606083811b821660805282901b1660a0523380600081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000e857620000e88162000102565b50505060601b6001600160601b03191660c0525062000203565b6001600160a01b0381163314156200015d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001c657600080fd5b919050565b60008060408385031215620001df57600080fd5b620001ea83620001ae565b9150620001fa60208401620001ae565b90509250929050565b60805160601c60a05160601c60c05160601c6115766200025860003960006101600152600081816103ad015281816106cc01528181610ce20152610e0901526000818161054e0152610cb801526115766000f3fe6080604052600436106100f75760003560e01c80638da5cb5b1161008a578063d826f88f11610059578063d826f88f14610300578063d8a4676f1461032c578063dc1670db1461035f578063f2fde38b1461037557600080fd5b80638da5cb5b1461021e578063a168fa8914610249578063afacbf9c146102ca578063b1e21749146102ea57600080fd5b8063737144bc116100c6578063737144bc146101bd57806374dba124146101d357806379ba5097146101e95780637a8042bd146101fe57600080fd5b80631757f11c146101035780631fe543e31461012c5780632353f2381461014e578063557d2e92146101a757600080fd5b366100fe57005b600080fd5b34801561010f57600080fd5b5061011960055481565b6040519081526020015b60405180910390f35b34801561013857600080fd5b5061014c610147366004611183565b610395565b005b34801561015a57600080fd5b506101827f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610123565b3480156101b357600080fd5b5061011960035481565b3480156101c957600080fd5b5061011960045481565b3480156101df57600080fd5b5061011960065481565b3480156101f557600080fd5b5061014c610447565b34801561020a57600080fd5b5061014c610219366004611151565b610544565b34801561022a57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610182565b34801561025557600080fd5b5061029d610264366004611151565b600960205260009081526040902080546001820154600383015460048401546005850154600690950154939460ff909316939192909186565b604080519687529415156020870152938501929092526060840152608083015260a082015260c001610123565b3480156102d657600080fd5b5061014c6102e5366004611272565b61064c565b3480156102f657600080fd5b5061011960075481565b34801561030c57600080fd5b5061014c6000600481905560058190556103e76006556003819055600255565b34801561033857600080fd5b5061034c610347366004611151565b610892565b60405161012397969594939291906113c2565b34801561036b57600080fd5b5061011960025481565b34801561038157600080fd5b5061014c6103903660046110f2565b6109f9565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610439576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f6f6e6c792056524620563220777261707065722063616e2066756c66696c6c0060448201526064015b60405180910390fd5b6104438282610a0d565b5050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610430565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61054c610be9565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6105a760005473ffffffffffffffffffffffffffffffffffffffff1690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101849052604401602060405180830381600087803b15801561061457600080fd5b505af1158015610628573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610443919061112f565b60005b8161ffff168161ffff16101561088b57600061066c868686610c6c565b60078190559050600061067d610ead565b6040517f4306d35400000000000000000000000000000000000000000000000000000000815263ffffffff8916600482015290915060009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690634306d3549060240160206040518083038186803b15801561070e57600080fd5b505afa158015610722573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610746919061116a565b6040805160e08101825282815260006020808301828152845183815280830186528486019081524260608601526080850184905260a0850189905260c0850184905289845260098352949092208351815591516001830180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905592518051949550919390926107e6926002850192910190611067565b5060608201516003808301919091556080830151600483015560a0830151600583015560c0909201516006909101558054906000610823836114d2565b9091555050600083815260086020526040908190208390555183907f5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec49061086d9084815260200190565b60405180910390a25050508080610883906114b0565b91505061064f565b5050505050565b60008181526009602052604081205481906060908290819081908190610914576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e640000000000000000000000000000006044820152606401610430565b6000888152600960209081526040808320815160e08101835281548152600182015460ff1615158185015260028201805484518187028101870186528181529295939486019383018282801561098957602002820191906000526020600020905b815481526020019060010190808311610975575b505050505081526020016003820154815260200160048201548152602001600582015481526020016006820154815250509050806000015181602001518260400151836060015184608001518560a001518660c00151975097509750975097509750975050919395979092949650565b610a01610be9565b610a0a81610f4a565b50565b600082815260096020526040902054610a82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e640000000000000000000000000000006044820152606401610430565b6000610a8c610ead565b60008481526008602052604081205491925090610aa99083611499565b90506000610aba82620f424061145c565b9050600554821115610acc5760058290555b600654821015610adc5760068290555b600060025411610aec5780610b1f565b600254610afa906001611409565b81600254600454610b0b919061145c565b610b159190611409565b610b1f9190611421565b60045560028054906000610b32836114d2565b90915550506000858152600960209081526040909120600181810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790558551610b8a92600290920191870190611067565b5060008581526009602052604090819020426004820155600681018590555490517f6c84e12b4c188e61f1b4727024a5cf05c025fa58467e5eedf763c0744c89da7b91610bda9188918891611399565b60405180910390a15050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610c6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610430565b565b6040517f4306d35400000000000000000000000000000000000000000000000000000000815263ffffffff8416600482015260009073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691634000aea0917f00000000000000000000000000000000000000000000000000000000000000009190821690634306d3549060240160206040518083038186803b158015610d2757600080fd5b505afa158015610d3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5f919061116a565b6040805163ffffffff808b16602083015261ffff8a169282019290925290871660608201526080016040516020818303038152906040526040518463ffffffff1660e01b8152600401610db493929190611301565b602060405180830381600087803b158015610dce57600080fd5b505af1158015610de2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e06919061112f565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc2a88c36040518163ffffffff1660e01b815260040160206040518083038186803b158015610e6d57600080fd5b505afa158015610e81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea5919061116a565b949350505050565b600046610eb981611040565b15610f4357606473ffffffffffffffffffffffffffffffffffffffff1663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f0557600080fd5b505afa158015610f19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3d919061116a565b91505090565b4391505090565b73ffffffffffffffffffffffffffffffffffffffff8116331415610fca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610430565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061a4b1821480611054575062066eed82145b80611061575062066eee82145b92915050565b8280548282559060005260206000209081019282156110a2579160200282015b828111156110a2578251825591602001919060010190611087565b506110ae9291506110b2565b5090565b5b808211156110ae57600081556001016110b3565b803561ffff811681146110d957600080fd5b919050565b803563ffffffff811681146110d957600080fd5b60006020828403121561110457600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461112857600080fd5b9392505050565b60006020828403121561114157600080fd5b8151801515811461112857600080fd5b60006020828403121561116357600080fd5b5035919050565b60006020828403121561117c57600080fd5b5051919050565b6000806040838503121561119657600080fd5b8235915060208084013567ffffffffffffffff808211156111b657600080fd5b818601915086601f8301126111ca57600080fd5b8135818111156111dc576111dc61153a565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110858211171561121f5761121f61153a565b604052828152858101935084860182860187018b101561123e57600080fd5b600095505b83861015611261578035855260019590950194938601938601611243565b508096505050505050509250929050565b6000806000806080858703121561128857600080fd5b611291856110de565b935061129f602086016110c7565b92506112ad604086016110de565b91506112bb606086016110c7565b905092959194509250565b600081518084526020808501945080840160005b838110156112f6578151875295820195908201906001016112da565b509495945050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260006020848184015260606040840152835180606085015260005b8181101561135157858101830151858201608001528201611335565b81811115611363576000608083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160800195945050505050565b8381526060602082015260006113b260608301856112c6565b9050826040830152949350505050565b878152861515602082015260e0604082015260006113e360e08301886112c6565b90508560608301528460808301528360a08301528260c083015298975050505050505050565b6000821982111561141c5761141c61150b565b500190565b600082611457577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156114945761149461150b565b500290565b6000828210156114ab576114ab61150b565b500390565b600061ffff808316818114156114c8576114c861150b565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156115045761150461150b565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV2WrapperLoadTestConsumerABI = VRFV2WrapperLoadTestConsumerMetaData.ABI

var VRFV2WrapperLoadTestConsumerBin = VRFV2WrapperLoadTestConsumerMetaData.Bin

func DeployVRFV2WrapperLoadTestConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _vrfV2Wrapper common.Address) (common.Address, *types.Transaction, *VRFV2WrapperLoadTestConsumer, error) {
	parsed, err := VRFV2WrapperLoadTestConsumerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV2WrapperLoadTestConsumerBin), backend, _link, _vrfV2Wrapper)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV2WrapperLoadTestConsumer{address: address, abi: *parsed, VRFV2WrapperLoadTestConsumerCaller: VRFV2WrapperLoadTestConsumerCaller{contract: contract}, VRFV2WrapperLoadTestConsumerTransactor: VRFV2WrapperLoadTestConsumerTransactor{contract: contract}, VRFV2WrapperLoadTestConsumerFilterer: VRFV2WrapperLoadTestConsumerFilterer{contract: contract}}, nil
}

type VRFV2WrapperLoadTestConsumer struct {
	address common.Address
	abi     abi.ABI
	VRFV2WrapperLoadTestConsumerCaller
	VRFV2WrapperLoadTestConsumerTransactor
	VRFV2WrapperLoadTestConsumerFilterer
}

type VRFV2WrapperLoadTestConsumerCaller struct {
	contract *bind.BoundContract
}

type VRFV2WrapperLoadTestConsumerTransactor struct {
	contract *bind.BoundContract
}

type VRFV2WrapperLoadTestConsumerFilterer struct {
	contract *bind.BoundContract
}

type VRFV2WrapperLoadTestConsumerSession struct {
	Contract     *VRFV2WrapperLoadTestConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV2WrapperLoadTestConsumerCallerSession struct {
	Contract *VRFV2WrapperLoadTestConsumerCaller
	CallOpts bind.CallOpts
}

type VRFV2WrapperLoadTestConsumerTransactorSession struct {
	Contract     *VRFV2WrapperLoadTestConsumerTransactor
	TransactOpts bind.TransactOpts
}

type VRFV2WrapperLoadTestConsumerRaw struct {
	Contract *VRFV2WrapperLoadTestConsumer
}

type VRFV2WrapperLoadTestConsumerCallerRaw struct {
	Contract *VRFV2WrapperLoadTestConsumerCaller
}

type VRFV2WrapperLoadTestConsumerTransactorRaw struct {
	Contract *VRFV2WrapperLoadTestConsumerTransactor
}

func NewVRFV2WrapperLoadTestConsumer(address common.Address, backend bind.ContractBackend) (*VRFV2WrapperLoadTestConsumer, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV2WrapperLoadTestConsumerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV2WrapperLoadTestConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumer{address: address, abi: abi, VRFV2WrapperLoadTestConsumerCaller: VRFV2WrapperLoadTestConsumerCaller{contract: contract}, VRFV2WrapperLoadTestConsumerTransactor: VRFV2WrapperLoadTestConsumerTransactor{contract: contract}, VRFV2WrapperLoadTestConsumerFilterer: VRFV2WrapperLoadTestConsumerFilterer{contract: contract}}, nil
}

func NewVRFV2WrapperLoadTestConsumerCaller(address common.Address, caller bind.ContractCaller) (*VRFV2WrapperLoadTestConsumerCaller, error) {
	contract, err := bindVRFV2WrapperLoadTestConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerCaller{contract: contract}, nil
}

func NewVRFV2WrapperLoadTestConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV2WrapperLoadTestConsumerTransactor, error) {
	contract, err := bindVRFV2WrapperLoadTestConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerTransactor{contract: contract}, nil
}

func NewVRFV2WrapperLoadTestConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV2WrapperLoadTestConsumerFilterer, error) {
	contract, err := bindVRFV2WrapperLoadTestConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerFilterer{contract: contract}, nil
}

func bindVRFV2WrapperLoadTestConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV2WrapperLoadTestConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2WrapperLoadTestConsumer.Contract.VRFV2WrapperLoadTestConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.VRFV2WrapperLoadTestConsumerTransactor.contract.Transfer(opts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.VRFV2WrapperLoadTestConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2WrapperLoadTestConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.contract.Transfer(opts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

	error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(GetRequestStatus)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.RequestTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RequestBlockNumber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.GetRequestStatus(&_VRFV2WrapperLoadTestConsumer.CallOpts, _requestId)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.GetRequestStatus(&_VRFV2WrapperLoadTestConsumer.CallOpts, _requestId)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) IVrfV2Wrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "i_vrfV2Wrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) IVrfV2Wrapper() (common.Address, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.IVrfV2Wrapper(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) IVrfV2Wrapper() (common.Address, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.IVrfV2Wrapper(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) Owner() (common.Address, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.Owner(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) Owner() (common.Address, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.Owner(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SAverageFulfillmentInMillions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_averageFulfillmentInMillions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SAverageFulfillmentInMillions() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SAverageFulfillmentInMillions(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SAverageFulfillmentInMillions() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SAverageFulfillmentInMillions(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SFastestFulfillment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_fastestFulfillment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SFastestFulfillment() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SFastestFulfillment(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SFastestFulfillment() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SFastestFulfillment(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SLastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SLastRequestId() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SLastRequestId(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SLastRequestId() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SLastRequestId(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SRequestCount() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SRequestCount(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SRequestCount() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SRequestCount(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

	error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(SRequests)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RequestTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RequestBlockNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentBlockNumber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SRequests(&_VRFV2WrapperLoadTestConsumer.CallOpts, arg0)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SRequests(&_VRFV2WrapperLoadTestConsumer.CallOpts, arg0)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SResponseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_responseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SResponseCount() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SResponseCount(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SResponseCount() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SResponseCount(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCaller) SSlowestFulfillment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2WrapperLoadTestConsumer.contract.Call(opts, &out, "s_slowestFulfillment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) SSlowestFulfillment() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SSlowestFulfillment(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerCallerSession) SSlowestFulfillment() (*big.Int, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.SSlowestFulfillment(&_VRFV2WrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.AcceptOwnership(&_VRFV2WrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.AcceptOwnership(&_VRFV2WrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) MakeRequests(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.Transact(opts, "makeRequests", _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) MakeRequests(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.MakeRequests(&_VRFV2WrapperLoadTestConsumer.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) MakeRequests(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.MakeRequests(&_VRFV2WrapperLoadTestConsumer.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.Transact(opts, "rawFulfillRandomWords", _requestId, _randomWords)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.RawFulfillRandomWords(&_VRFV2WrapperLoadTestConsumer.TransactOpts, _requestId, _randomWords)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.RawFulfillRandomWords(&_VRFV2WrapperLoadTestConsumer.TransactOpts, _requestId, _randomWords)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.Transact(opts, "reset")
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) Reset() (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.Reset(&_VRFV2WrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) Reset() (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.Reset(&_VRFV2WrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.TransferOwnership(&_VRFV2WrapperLoadTestConsumer.TransactOpts, to)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.TransferOwnership(&_VRFV2WrapperLoadTestConsumer.TransactOpts, to)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) WithdrawLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.Transact(opts, "withdrawLink", amount)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) WithdrawLink(amount *big.Int) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.WithdrawLink(&_VRFV2WrapperLoadTestConsumer.TransactOpts, amount)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) WithdrawLink(amount *big.Int) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.WithdrawLink(&_VRFV2WrapperLoadTestConsumer.TransactOpts, amount)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.contract.RawTransact(opts, nil)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerSession) Receive() (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.Receive(&_VRFV2WrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerTransactorSession) Receive() (*types.Transaction, error) {
	return _VRFV2WrapperLoadTestConsumer.Contract.Receive(&_VRFV2WrapperLoadTestConsumer.TransactOpts)
}

type VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator struct {
	Event *VRFV2WrapperLoadTestConsumerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperLoadTestConsumerOwnershipTransferRequested)
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
		it.Event = new(VRFV2WrapperLoadTestConsumerOwnershipTransferRequested)
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

func (it *VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperLoadTestConsumerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator{contract: _VRFV2WrapperLoadTestConsumer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperLoadTestConsumerOwnershipTransferRequested)
				if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV2WrapperLoadTestConsumerOwnershipTransferRequested, error) {
	event := new(VRFV2WrapperLoadTestConsumerOwnershipTransferRequested)
	if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator struct {
	Event *VRFV2WrapperLoadTestConsumerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperLoadTestConsumerOwnershipTransferred)
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
		it.Event = new(VRFV2WrapperLoadTestConsumerOwnershipTransferred)
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

func (it *VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperLoadTestConsumerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator{contract: _VRFV2WrapperLoadTestConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperLoadTestConsumerOwnershipTransferred)
				if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV2WrapperLoadTestConsumerOwnershipTransferred, error) {
	event := new(VRFV2WrapperLoadTestConsumerOwnershipTransferred)
	if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator struct {
	Event *VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled)
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
		it.Event = new(VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled)
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

func (it *VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Payment     *big.Int
	Raw         types.Log
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) FilterWrappedRequestFulfilled(opts *bind.FilterOpts) (*VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator, error) {

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.FilterLogs(opts, "WrappedRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator{contract: _VRFV2WrapperLoadTestConsumer.contract, event: "WrappedRequestFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) WatchWrappedRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.WatchLogs(opts, "WrappedRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled)
				if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "WrappedRequestFulfilled", log); err != nil {
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

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) ParseWrappedRequestFulfilled(log types.Log) (*VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled, error) {
	event := new(VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled)
	if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "WrappedRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator struct {
	Event *VRFV2WrapperLoadTestConsumerWrapperRequestMade

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperLoadTestConsumerWrapperRequestMade)
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
		it.Event = new(VRFV2WrapperLoadTestConsumerWrapperRequestMade)
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

func (it *VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperLoadTestConsumerWrapperRequestMade struct {
	RequestId *big.Int
	Paid      *big.Int
	Raw       types.Log
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) FilterWrapperRequestMade(opts *bind.FilterOpts, requestId []*big.Int) (*VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.FilterLogs(opts, "WrapperRequestMade", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator{contract: _VRFV2WrapperLoadTestConsumer.contract, event: "WrapperRequestMade", logs: logs, sub: sub}, nil
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) WatchWrapperRequestMade(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerWrapperRequestMade, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFV2WrapperLoadTestConsumer.contract.WatchLogs(opts, "WrapperRequestMade", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperLoadTestConsumerWrapperRequestMade)
				if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "WrapperRequestMade", log); err != nil {
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

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumerFilterer) ParseWrapperRequestMade(log types.Log) (*VRFV2WrapperLoadTestConsumerWrapperRequestMade, error) {
	event := new(VRFV2WrapperLoadTestConsumerWrapperRequestMade)
	if err := _VRFV2WrapperLoadTestConsumer.contract.UnpackLog(event, "WrapperRequestMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetRequestStatus struct {
	Paid                  *big.Int
	Fulfilled             bool
	RandomWords           []*big.Int
	RequestTimestamp      *big.Int
	FulfilmentTimestamp   *big.Int
	RequestBlockNumber    *big.Int
	FulfilmentBlockNumber *big.Int
}
type SRequests struct {
	Paid                  *big.Int
	Fulfilled             bool
	RequestTimestamp      *big.Int
	FulfilmentTimestamp   *big.Int
	RequestBlockNumber    *big.Int
	FulfilmentBlockNumber *big.Int
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumer) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV2WrapperLoadTestConsumer.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV2WrapperLoadTestConsumer.ParseOwnershipTransferRequested(log)
	case _VRFV2WrapperLoadTestConsumer.abi.Events["OwnershipTransferred"].ID:
		return _VRFV2WrapperLoadTestConsumer.ParseOwnershipTransferred(log)
	case _VRFV2WrapperLoadTestConsumer.abi.Events["WrappedRequestFulfilled"].ID:
		return _VRFV2WrapperLoadTestConsumer.ParseWrappedRequestFulfilled(log)
	case _VRFV2WrapperLoadTestConsumer.abi.Events["WrapperRequestMade"].ID:
		return _VRFV2WrapperLoadTestConsumer.ParseWrapperRequestMade(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV2WrapperLoadTestConsumerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV2WrapperLoadTestConsumerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled) Topic() common.Hash {
	return common.HexToHash("0x6c84e12b4c188e61f1b4727024a5cf05c025fa58467e5eedf763c0744c89da7b")
}

func (VRFV2WrapperLoadTestConsumerWrapperRequestMade) Topic() common.Hash {
	return common.HexToHash("0x5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec4")
}

func (_VRFV2WrapperLoadTestConsumer *VRFV2WrapperLoadTestConsumer) Address() common.Address {
	return _VRFV2WrapperLoadTestConsumer.address
}

type VRFV2WrapperLoadTestConsumerInterface interface {
	GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

		error)

	IVrfV2Wrapper(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAverageFulfillmentInMillions(opts *bind.CallOpts) (*big.Int, error)

	SFastestFulfillment(opts *bind.CallOpts) (*big.Int, error)

	SLastRequestId(opts *bind.CallOpts) (*big.Int, error)

	SRequestCount(opts *bind.CallOpts) (*big.Int, error)

	SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

		error)

	SResponseCount(opts *bind.CallOpts) (*big.Int, error)

	SSlowestFulfillment(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	MakeRequests(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error)

	Reset(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperLoadTestConsumerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV2WrapperLoadTestConsumerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperLoadTestConsumerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV2WrapperLoadTestConsumerOwnershipTransferred, error)

	FilterWrappedRequestFulfilled(opts *bind.FilterOpts) (*VRFV2WrapperLoadTestConsumerWrappedRequestFulfilledIterator, error)

	WatchWrappedRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled) (event.Subscription, error)

	ParseWrappedRequestFulfilled(log types.Log) (*VRFV2WrapperLoadTestConsumerWrappedRequestFulfilled, error)

	FilterWrapperRequestMade(opts *bind.FilterOpts, requestId []*big.Int) (*VRFV2WrapperLoadTestConsumerWrapperRequestMadeIterator, error)

	WatchWrapperRequestMade(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperLoadTestConsumerWrapperRequestMade, requestId []*big.Int) (event.Subscription, error)

	ParseWrapperRequestMade(log types.Log) (*VRFV2WrapperLoadTestConsumerWrapperRequestMade, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
