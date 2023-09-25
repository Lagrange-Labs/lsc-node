// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mantle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
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

// AssertionMapMetaData contains all meta data concerning the AssertionMap contract.
var AssertionMapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStakers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"childInboxSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AssertionMapABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionMapMetaData.ABI instead.
var AssertionMapABI = AssertionMapMetaData.ABI

// AssertionMap is an auto generated Go binding around an Ethereum contract.
type AssertionMap struct {
	AssertionMapCaller     // Read-only binding to the contract
	AssertionMapTransactor // Write-only binding to the contract
	AssertionMapFilterer   // Log filterer for contract events
}

// AssertionMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionMapSession struct {
	Contract     *AssertionMap     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertionMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionMapCallerSession struct {
	Contract *AssertionMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AssertionMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionMapTransactorSession struct {
	Contract     *AssertionMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AssertionMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionMapRaw struct {
	Contract *AssertionMap // Generic contract binding to access the raw methods on
}

// AssertionMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionMapCallerRaw struct {
	Contract *AssertionMapCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionMapTransactorRaw struct {
	Contract *AssertionMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertionMap creates a new instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMap(address common.Address, backend bind.ContractBackend) (*AssertionMap, error) {
	contract, err := bindAssertionMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssertionMap{AssertionMapCaller: AssertionMapCaller{contract: contract}, AssertionMapTransactor: AssertionMapTransactor{contract: contract}, AssertionMapFilterer: AssertionMapFilterer{contract: contract}}, nil
}

// NewAssertionMapCaller creates a new read-only instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMapCaller(address common.Address, caller bind.ContractCaller) (*AssertionMapCaller, error) {
	contract, err := bindAssertionMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionMapCaller{contract: contract}, nil
}

// NewAssertionMapTransactor creates a new write-only instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMapTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionMapTransactor, error) {
	contract, err := bindAssertionMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionMapTransactor{contract: contract}, nil
}

// NewAssertionMapFilterer creates a new log filterer instance of AssertionMap, bound to a specific deployed contract.
func NewAssertionMapFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionMapFilterer, error) {
	contract, err := bindAssertionMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionMapFilterer{contract: contract}, nil
}

// bindAssertionMap binds a generic wrapper to an already deployed contract.
func bindAssertionMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssertionMapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionMap *AssertionMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionMap.Contract.AssertionMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionMap *AssertionMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionMap.Contract.AssertionMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionMap *AssertionMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionMap.Contract.AssertionMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionMap *AssertionMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionMap *AssertionMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionMap *AssertionMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionMap.Contract.contract.Transact(opts, method, params...)
}

// Assertions is a free data retrieval call binding the contract method 0x524232f6.
//
// Solidity: function assertions(uint256 ) view returns(bytes32 stateHash, uint256 inboxSize, uint256 parent, uint256 deadline, uint256 proposalTime, uint256 numStakers, uint256 childInboxSize)
func (_AssertionMap *AssertionMapCaller) Assertions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateHash      [32]byte
	InboxSize      *big.Int
	Parent         *big.Int
	Deadline       *big.Int
	ProposalTime   *big.Int
	NumStakers     *big.Int
	ChildInboxSize *big.Int
}, error) {
	var out []interface{}
	err := _AssertionMap.contract.Call(opts, &out, "assertions", arg0)

	outstruct := new(struct {
		StateHash      [32]byte
		InboxSize      *big.Int
		Parent         *big.Int
		Deadline       *big.Int
		ProposalTime   *big.Int
		NumStakers     *big.Int
		ChildInboxSize *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.InboxSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Parent = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ProposalTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NumStakers = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ChildInboxSize = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assertions is a free data retrieval call binding the contract method 0x524232f6.
//
// Solidity: function assertions(uint256 ) view returns(bytes32 stateHash, uint256 inboxSize, uint256 parent, uint256 deadline, uint256 proposalTime, uint256 numStakers, uint256 childInboxSize)
func (_AssertionMap *AssertionMapSession) Assertions(arg0 *big.Int) (struct {
	StateHash      [32]byte
	InboxSize      *big.Int
	Parent         *big.Int
	Deadline       *big.Int
	ProposalTime   *big.Int
	NumStakers     *big.Int
	ChildInboxSize *big.Int
}, error) {
	return _AssertionMap.Contract.Assertions(&_AssertionMap.CallOpts, arg0)
}

// Assertions is a free data retrieval call binding the contract method 0x524232f6.
//
// Solidity: function assertions(uint256 ) view returns(bytes32 stateHash, uint256 inboxSize, uint256 parent, uint256 deadline, uint256 proposalTime, uint256 numStakers, uint256 childInboxSize)
func (_AssertionMap *AssertionMapCallerSession) Assertions(arg0 *big.Int) (struct {
	StateHash      [32]byte
	InboxSize      *big.Int
	Parent         *big.Int
	Deadline       *big.Int
	ProposalTime   *big.Int
	NumStakers     *big.Int
	ChildInboxSize *big.Int
}, error) {
	return _AssertionMap.Contract.Assertions(&_AssertionMap.CallOpts, arg0)
}

