// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodestaking

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
)

// NodestakingMetaData contains all meta data concerning the Nodestaking contract.
var NodestakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimTime\",\"type\":\"uint256\"}],\"name\":\"PendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"}],\"name\":\"SlashedNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimTime\",\"type\":\"uint256\"}],\"name\":\"StakedNode\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeStakes\",\"outputs\":[{\"internalType\":\"enumDataTypes.STAKE_STATUS\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimTimeMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"configureSlashers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"finishStakeRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialStakeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialStakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_claimDelay\",\"type\":\"uint256\"}],\"name\":\"setClaimDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeTime\",\"type\":\"uint256\"}],\"name\":\"setInitialStakeTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"setStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"slashStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedAmountMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"startStakeRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NodestakingABI is the input ABI used to generate the binding from.
// Deprecated: Use NodestakingMetaData.ABI instead.
var NodestakingABI = NodestakingMetaData.ABI

// Nodestaking is an auto generated Go binding around an Ethereum contract.
type Nodestaking struct {
	NodestakingCaller     // Read-only binding to the contract
	NodestakingTransactor // Write-only binding to the contract
	NodestakingFilterer   // Log filterer for contract events
}

// NodestakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodestakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodestakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodestakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodestakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodestakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodestakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodestakingSession struct {
	Contract     *Nodestaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodestakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodestakingCallerSession struct {
	Contract *NodestakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodestakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodestakingTransactorSession struct {
	Contract     *NodestakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodestakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodestakingRaw struct {
	Contract *Nodestaking // Generic contract binding to access the raw methods on
}

// NodestakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodestakingCallerRaw struct {
	Contract *NodestakingCaller // Generic read-only contract binding to access the raw methods on
}

// NodestakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodestakingTransactorRaw struct {
	Contract *NodestakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodestaking creates a new instance of Nodestaking, bound to a specific deployed contract.
func NewNodestaking(address common.Address, backend bind.ContractBackend) (*Nodestaking, error) {
	contract, err := bindNodestaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nodestaking{NodestakingCaller: NodestakingCaller{contract: contract}, NodestakingTransactor: NodestakingTransactor{contract: contract}, NodestakingFilterer: NodestakingFilterer{contract: contract}}, nil
}

// NewNodestakingCaller creates a new read-only instance of Nodestaking, bound to a specific deployed contract.
func NewNodestakingCaller(address common.Address, caller bind.ContractCaller) (*NodestakingCaller, error) {
	contract, err := bindNodestaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodestakingCaller{contract: contract}, nil
}

// NewNodestakingTransactor creates a new write-only instance of Nodestaking, bound to a specific deployed contract.
func NewNodestakingTransactor(address common.Address, transactor bind.ContractTransactor) (*NodestakingTransactor, error) {
	contract, err := bindNodestaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodestakingTransactor{contract: contract}, nil
}

// NewNodestakingFilterer creates a new log filterer instance of Nodestaking, bound to a specific deployed contract.
func NewNodestakingFilterer(address common.Address, filterer bind.ContractFilterer) (*NodestakingFilterer, error) {
	contract, err := bindNodestaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodestakingFilterer{contract: contract}, nil
}

// bindNodestaking binds a generic wrapper to an already deployed contract.
func bindNodestaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodestakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodestaking *NodestakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodestaking.Contract.NodestakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodestaking *NodestakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodestaking.Contract.NodestakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodestaking *NodestakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodestaking.Contract.NodestakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodestaking *NodestakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodestaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodestaking *NodestakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodestaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodestaking *NodestakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodestaking.Contract.contract.Transact(opts, method, params...)
}

// ActiveStakes is a free data retrieval call binding the contract method 0x5c31f93c.
//
// Solidity: function activeStakes(address , uint256 ) view returns(uint8)
func (_Nodestaking *NodestakingCaller) ActiveStakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "activeStakes", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ActiveStakes is a free data retrieval call binding the contract method 0x5c31f93c.
//
// Solidity: function activeStakes(address , uint256 ) view returns(uint8)
func (_Nodestaking *NodestakingSession) ActiveStakes(arg0 common.Address, arg1 *big.Int) (uint8, error) {
	return _Nodestaking.Contract.ActiveStakes(&_Nodestaking.CallOpts, arg0, arg1)
}

// ActiveStakes is a free data retrieval call binding the contract method 0x5c31f93c.
//
// Solidity: function activeStakes(address , uint256 ) view returns(uint8)
func (_Nodestaking *NodestakingCallerSession) ActiveStakes(arg0 common.Address, arg1 *big.Int) (uint8, error) {
	return _Nodestaking.Contract.ActiveStakes(&_Nodestaking.CallOpts, arg0, arg1)
}

// ClaimDelay is a free data retrieval call binding the contract method 0x1c8ec299.
//
// Solidity: function claimDelay() view returns(uint256)
func (_Nodestaking *NodestakingCaller) ClaimDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "claimDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimDelay is a free data retrieval call binding the contract method 0x1c8ec299.
//
// Solidity: function claimDelay() view returns(uint256)
func (_Nodestaking *NodestakingSession) ClaimDelay() (*big.Int, error) {
	return _Nodestaking.Contract.ClaimDelay(&_Nodestaking.CallOpts)
}

// ClaimDelay is a free data retrieval call binding the contract method 0x1c8ec299.
//
// Solidity: function claimDelay() view returns(uint256)
func (_Nodestaking *NodestakingCallerSession) ClaimDelay() (*big.Int, error) {
	return _Nodestaking.Contract.ClaimDelay(&_Nodestaking.CallOpts)
}

// ClaimTimeMap is a free data retrieval call binding the contract method 0xea111354.
//
// Solidity: function claimTimeMap(address , uint256 ) view returns(uint256)
func (_Nodestaking *NodestakingCaller) ClaimTimeMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "claimTimeMap", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimTimeMap is a free data retrieval call binding the contract method 0xea111354.
//
// Solidity: function claimTimeMap(address , uint256 ) view returns(uint256)
func (_Nodestaking *NodestakingSession) ClaimTimeMap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Nodestaking.Contract.ClaimTimeMap(&_Nodestaking.CallOpts, arg0, arg1)
}

// ClaimTimeMap is a free data retrieval call binding the contract method 0xea111354.
//
// Solidity: function claimTimeMap(address , uint256 ) view returns(uint256)
func (_Nodestaking *NodestakingCallerSession) ClaimTimeMap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Nodestaking.Contract.ClaimTimeMap(&_Nodestaking.CallOpts, arg0, arg1)
}

// InitialStakeTime is a free data retrieval call binding the contract method 0xe9994cee.
//
// Solidity: function initialStakeTime() view returns(uint256)
func (_Nodestaking *NodestakingCaller) InitialStakeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "initialStakeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialStakeTime is a free data retrieval call binding the contract method 0xe9994cee.
//
// Solidity: function initialStakeTime() view returns(uint256)
func (_Nodestaking *NodestakingSession) InitialStakeTime() (*big.Int, error) {
	return _Nodestaking.Contract.InitialStakeTime(&_Nodestaking.CallOpts)
}

// InitialStakeTime is a free data retrieval call binding the contract method 0xe9994cee.
//
// Solidity: function initialStakeTime() view returns(uint256)
func (_Nodestaking *NodestakingCallerSession) InitialStakeTime() (*big.Int, error) {
	return _Nodestaking.Contract.InitialStakeTime(&_Nodestaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodestaking *NodestakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodestaking *NodestakingSession) Owner() (common.Address, error) {
	return _Nodestaking.Contract.Owner(&_Nodestaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodestaking *NodestakingCallerSession) Owner() (common.Address, error) {
	return _Nodestaking.Contract.Owner(&_Nodestaking.CallOpts)
}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_Nodestaking *NodestakingCaller) Slashers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "slashers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_Nodestaking *NodestakingSession) Slashers(arg0 common.Address) (bool, error) {
	return _Nodestaking.Contract.Slashers(&_Nodestaking.CallOpts, arg0)
}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_Nodestaking *NodestakingCallerSession) Slashers(arg0 common.Address) (bool, error) {
	return _Nodestaking.Contract.Slashers(&_Nodestaking.CallOpts, arg0)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_Nodestaking *NodestakingCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_Nodestaking *NodestakingSession) StakeAmount() (*big.Int, error) {
	return _Nodestaking.Contract.StakeAmount(&_Nodestaking.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_Nodestaking *NodestakingCallerSession) StakeAmount() (*big.Int, error) {
	return _Nodestaking.Contract.StakeAmount(&_Nodestaking.CallOpts)
}

// StakedAmountMap is a free data retrieval call binding the contract method 0x037b966b.
//
// Solidity: function stakedAmountMap(address , uint256 ) view returns(uint256)
func (_Nodestaking *NodestakingCaller) StakedAmountMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodestaking.contract.Call(opts, &out, "stakedAmountMap", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmountMap is a free data retrieval call binding the contract method 0x037b966b.
//
// Solidity: function stakedAmountMap(address , uint256 ) view returns(uint256)
func (_Nodestaking *NodestakingSession) StakedAmountMap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Nodestaking.Contract.StakedAmountMap(&_Nodestaking.CallOpts, arg0, arg1)
}

// StakedAmountMap is a free data retrieval call binding the contract method 0x037b966b.
//
// Solidity: function stakedAmountMap(address , uint256 ) view returns(uint256)
func (_Nodestaking *NodestakingCallerSession) StakedAmountMap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Nodestaking.Contract.StakedAmountMap(&_Nodestaking.CallOpts, arg0, arg1)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 chainID) payable returns()
func (_Nodestaking *NodestakingTransactor) AddStake(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "addStake", chainID)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 chainID) payable returns()
func (_Nodestaking *NodestakingSession) AddStake(chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.AddStake(&_Nodestaking.TransactOpts, chainID)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 chainID) payable returns()
func (_Nodestaking *NodestakingTransactorSession) AddStake(chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.AddStake(&_Nodestaking.TransactOpts, chainID)
}

// ConfigureSlashers is a paid mutator transaction binding the contract method 0x2bd09d1c.
//
// Solidity: function configureSlashers(address slasher, bool value) returns()
func (_Nodestaking *NodestakingTransactor) ConfigureSlashers(opts *bind.TransactOpts, slasher common.Address, value bool) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "configureSlashers", slasher, value)
}

// ConfigureSlashers is a paid mutator transaction binding the contract method 0x2bd09d1c.
//
// Solidity: function configureSlashers(address slasher, bool value) returns()
func (_Nodestaking *NodestakingSession) ConfigureSlashers(slasher common.Address, value bool) (*types.Transaction, error) {
	return _Nodestaking.Contract.ConfigureSlashers(&_Nodestaking.TransactOpts, slasher, value)
}

// ConfigureSlashers is a paid mutator transaction binding the contract method 0x2bd09d1c.
//
// Solidity: function configureSlashers(address slasher, bool value) returns()
func (_Nodestaking *NodestakingTransactorSession) ConfigureSlashers(slasher common.Address, value bool) (*types.Transaction, error) {
	return _Nodestaking.Contract.ConfigureSlashers(&_Nodestaking.TransactOpts, slasher, value)
}

// FinishStakeRemoval is a paid mutator transaction binding the contract method 0xc742826a.
//
// Solidity: function finishStakeRemoval(uint256 chainID) returns()
func (_Nodestaking *NodestakingTransactor) FinishStakeRemoval(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "finishStakeRemoval", chainID)
}

// FinishStakeRemoval is a paid mutator transaction binding the contract method 0xc742826a.
//
// Solidity: function finishStakeRemoval(uint256 chainID) returns()
func (_Nodestaking *NodestakingSession) FinishStakeRemoval(chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.FinishStakeRemoval(&_Nodestaking.TransactOpts, chainID)
}

// FinishStakeRemoval is a paid mutator transaction binding the contract method 0xc742826a.
//
// Solidity: function finishStakeRemoval(uint256 chainID) returns()
func (_Nodestaking *NodestakingTransactorSession) FinishStakeRemoval(chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.FinishStakeRemoval(&_Nodestaking.TransactOpts, chainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _stakeAmount, uint256 _initialStakeTime, uint256 _claimDelay) returns()
func (_Nodestaking *NodestakingTransactor) Initialize(opts *bind.TransactOpts, _stakeAmount *big.Int, _initialStakeTime *big.Int, _claimDelay *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "initialize", _stakeAmount, _initialStakeTime, _claimDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _stakeAmount, uint256 _initialStakeTime, uint256 _claimDelay) returns()
func (_Nodestaking *NodestakingSession) Initialize(_stakeAmount *big.Int, _initialStakeTime *big.Int, _claimDelay *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.Initialize(&_Nodestaking.TransactOpts, _stakeAmount, _initialStakeTime, _claimDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _stakeAmount, uint256 _initialStakeTime, uint256 _claimDelay) returns()
func (_Nodestaking *NodestakingTransactorSession) Initialize(_stakeAmount *big.Int, _initialStakeTime *big.Int, _claimDelay *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.Initialize(&_Nodestaking.TransactOpts, _stakeAmount, _initialStakeTime, _claimDelay)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodestaking *NodestakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodestaking *NodestakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nodestaking.Contract.RenounceOwnership(&_Nodestaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodestaking *NodestakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nodestaking.Contract.RenounceOwnership(&_Nodestaking.TransactOpts)
}

// SetClaimDelayTime is a paid mutator transaction binding the contract method 0x9ea7cdd2.
//
// Solidity: function setClaimDelayTime(uint256 _claimDelay) returns()
func (_Nodestaking *NodestakingTransactor) SetClaimDelayTime(opts *bind.TransactOpts, _claimDelay *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "setClaimDelayTime", _claimDelay)
}

// SetClaimDelayTime is a paid mutator transaction binding the contract method 0x9ea7cdd2.
//
// Solidity: function setClaimDelayTime(uint256 _claimDelay) returns()
func (_Nodestaking *NodestakingSession) SetClaimDelayTime(_claimDelay *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.SetClaimDelayTime(&_Nodestaking.TransactOpts, _claimDelay)
}

// SetClaimDelayTime is a paid mutator transaction binding the contract method 0x9ea7cdd2.
//
// Solidity: function setClaimDelayTime(uint256 _claimDelay) returns()
func (_Nodestaking *NodestakingTransactorSession) SetClaimDelayTime(_claimDelay *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.SetClaimDelayTime(&_Nodestaking.TransactOpts, _claimDelay)
}

// SetInitialStakeTime is a paid mutator transaction binding the contract method 0x4b193838.
//
// Solidity: function setInitialStakeTime(uint256 _stakeTime) returns()
func (_Nodestaking *NodestakingTransactor) SetInitialStakeTime(opts *bind.TransactOpts, _stakeTime *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "setInitialStakeTime", _stakeTime)
}

// SetInitialStakeTime is a paid mutator transaction binding the contract method 0x4b193838.
//
// Solidity: function setInitialStakeTime(uint256 _stakeTime) returns()
func (_Nodestaking *NodestakingSession) SetInitialStakeTime(_stakeTime *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.SetInitialStakeTime(&_Nodestaking.TransactOpts, _stakeTime)
}

// SetInitialStakeTime is a paid mutator transaction binding the contract method 0x4b193838.
//
// Solidity: function setInitialStakeTime(uint256 _stakeTime) returns()
func (_Nodestaking *NodestakingTransactorSession) SetInitialStakeTime(_stakeTime *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.SetInitialStakeTime(&_Nodestaking.TransactOpts, _stakeTime)
}

// SetStakeAmount is a paid mutator transaction binding the contract method 0x43808c50.
//
// Solidity: function setStakeAmount(uint256 _stakeAmount) returns()
func (_Nodestaking *NodestakingTransactor) SetStakeAmount(opts *bind.TransactOpts, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "setStakeAmount", _stakeAmount)
}

// SetStakeAmount is a paid mutator transaction binding the contract method 0x43808c50.
//
// Solidity: function setStakeAmount(uint256 _stakeAmount) returns()
func (_Nodestaking *NodestakingSession) SetStakeAmount(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.SetStakeAmount(&_Nodestaking.TransactOpts, _stakeAmount)
}

// SetStakeAmount is a paid mutator transaction binding the contract method 0x43808c50.
//
// Solidity: function setStakeAmount(uint256 _stakeAmount) returns()
func (_Nodestaking *NodestakingTransactorSession) SetStakeAmount(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.SetStakeAmount(&_Nodestaking.TransactOpts, _stakeAmount)
}

// SlashStake is a paid mutator transaction binding the contract method 0xca315ca6.
//
// Solidity: function slashStake(uint256 chainID, address user) returns()
func (_Nodestaking *NodestakingTransactor) SlashStake(opts *bind.TransactOpts, chainID *big.Int, user common.Address) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "slashStake", chainID, user)
}

// SlashStake is a paid mutator transaction binding the contract method 0xca315ca6.
//
// Solidity: function slashStake(uint256 chainID, address user) returns()
func (_Nodestaking *NodestakingSession) SlashStake(chainID *big.Int, user common.Address) (*types.Transaction, error) {
	return _Nodestaking.Contract.SlashStake(&_Nodestaking.TransactOpts, chainID, user)
}

// SlashStake is a paid mutator transaction binding the contract method 0xca315ca6.
//
// Solidity: function slashStake(uint256 chainID, address user) returns()
func (_Nodestaking *NodestakingTransactorSession) SlashStake(chainID *big.Int, user common.Address) (*types.Transaction, error) {
	return _Nodestaking.Contract.SlashStake(&_Nodestaking.TransactOpts, chainID, user)
}

// StartStakeRemoval is a paid mutator transaction binding the contract method 0x7aef4be8.
//
// Solidity: function startStakeRemoval(uint256 chainID) returns()
func (_Nodestaking *NodestakingTransactor) StartStakeRemoval(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "startStakeRemoval", chainID)
}

// StartStakeRemoval is a paid mutator transaction binding the contract method 0x7aef4be8.
//
// Solidity: function startStakeRemoval(uint256 chainID) returns()
func (_Nodestaking *NodestakingSession) StartStakeRemoval(chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.StartStakeRemoval(&_Nodestaking.TransactOpts, chainID)
}

// StartStakeRemoval is a paid mutator transaction binding the contract method 0x7aef4be8.
//
// Solidity: function startStakeRemoval(uint256 chainID) returns()
func (_Nodestaking *NodestakingTransactorSession) StartStakeRemoval(chainID *big.Int) (*types.Transaction, error) {
	return _Nodestaking.Contract.StartStakeRemoval(&_Nodestaking.TransactOpts, chainID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodestaking *NodestakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nodestaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodestaking *NodestakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodestaking.Contract.TransferOwnership(&_Nodestaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodestaking *NodestakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodestaking.Contract.TransferOwnership(&_Nodestaking.TransactOpts, newOwner)
}

// NodestakingClaimedNodeIterator is returned from FilterClaimedNode and is used to iterate over the raw logs and unpacked data for ClaimedNode events raised by the Nodestaking contract.
type NodestakingClaimedNodeIterator struct {
	Event *NodestakingClaimedNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodestakingClaimedNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodestakingClaimedNode)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodestakingClaimedNode)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodestakingClaimedNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodestakingClaimedNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodestakingClaimedNode represents a ClaimedNode event raised by the Nodestaking contract.
type NodestakingClaimedNode struct {
	Node    common.Address
	ChainId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedNode is a free log retrieval operation binding the contract event 0x9465c52c6301cd7fc0a16603c1390eb858e570eea04e2651408a3febf2ed8028.
//
// Solidity: event ClaimedNode(address node, uint256 chainId, uint256 amount)
func (_Nodestaking *NodestakingFilterer) FilterClaimedNode(opts *bind.FilterOpts) (*NodestakingClaimedNodeIterator, error) {

	logs, sub, err := _Nodestaking.contract.FilterLogs(opts, "ClaimedNode")
	if err != nil {
		return nil, err
	}
	return &NodestakingClaimedNodeIterator{contract: _Nodestaking.contract, event: "ClaimedNode", logs: logs, sub: sub}, nil
}

// WatchClaimedNode is a free log subscription operation binding the contract event 0x9465c52c6301cd7fc0a16603c1390eb858e570eea04e2651408a3febf2ed8028.
//
// Solidity: event ClaimedNode(address node, uint256 chainId, uint256 amount)
func (_Nodestaking *NodestakingFilterer) WatchClaimedNode(opts *bind.WatchOpts, sink chan<- *NodestakingClaimedNode) (event.Subscription, error) {

	logs, sub, err := _Nodestaking.contract.WatchLogs(opts, "ClaimedNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodestakingClaimedNode)
				if err := _Nodestaking.contract.UnpackLog(event, "ClaimedNode", log); err != nil {
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

// ParseClaimedNode is a log parse operation binding the contract event 0x9465c52c6301cd7fc0a16603c1390eb858e570eea04e2651408a3febf2ed8028.
//
// Solidity: event ClaimedNode(address node, uint256 chainId, uint256 amount)
func (_Nodestaking *NodestakingFilterer) ParseClaimedNode(log types.Log) (*NodestakingClaimedNode, error) {
	event := new(NodestakingClaimedNode)
	if err := _Nodestaking.contract.UnpackLog(event, "ClaimedNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodestakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Nodestaking contract.
type NodestakingInitializedIterator struct {
	Event *NodestakingInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodestakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodestakingInitialized)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodestakingInitialized)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodestakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodestakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodestakingInitialized represents a Initialized event raised by the Nodestaking contract.
type NodestakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodestaking *NodestakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodestakingInitializedIterator, error) {

	logs, sub, err := _Nodestaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodestakingInitializedIterator{contract: _Nodestaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodestaking *NodestakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodestakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Nodestaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodestakingInitialized)
				if err := _Nodestaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodestaking *NodestakingFilterer) ParseInitialized(log types.Log) (*NodestakingInitialized, error) {
	event := new(NodestakingInitialized)
	if err := _Nodestaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodestakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nodestaking contract.
type NodestakingOwnershipTransferredIterator struct {
	Event *NodestakingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodestakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodestakingOwnershipTransferred)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodestakingOwnershipTransferred)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodestakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodestakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodestakingOwnershipTransferred represents a OwnershipTransferred event raised by the Nodestaking contract.
type NodestakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodestaking *NodestakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodestakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nodestaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodestakingOwnershipTransferredIterator{contract: _Nodestaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodestaking *NodestakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodestakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nodestaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodestakingOwnershipTransferred)
				if err := _Nodestaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodestaking *NodestakingFilterer) ParseOwnershipTransferred(log types.Log) (*NodestakingOwnershipTransferred, error) {
	event := new(NodestakingOwnershipTransferred)
	if err := _Nodestaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodestakingPendingNodeIterator is returned from FilterPendingNode and is used to iterate over the raw logs and unpacked data for PendingNode events raised by the Nodestaking contract.
type NodestakingPendingNodeIterator struct {
	Event *NodestakingPendingNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodestakingPendingNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodestakingPendingNode)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodestakingPendingNode)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodestakingPendingNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodestakingPendingNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodestakingPendingNode represents a PendingNode event raised by the Nodestaking contract.
type NodestakingPendingNode struct {
	Node      common.Address
	ChainId   *big.Int
	Amount    *big.Int
	ClaimTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPendingNode is a free log retrieval operation binding the contract event 0xf0a2b450365814b25d2f317ccb08e29c8d1249ee8446074d8112ce30045b6647.
//
// Solidity: event PendingNode(address node, uint256 chainId, uint256 amount, uint256 claimTime)
func (_Nodestaking *NodestakingFilterer) FilterPendingNode(opts *bind.FilterOpts) (*NodestakingPendingNodeIterator, error) {

	logs, sub, err := _Nodestaking.contract.FilterLogs(opts, "PendingNode")
	if err != nil {
		return nil, err
	}
	return &NodestakingPendingNodeIterator{contract: _Nodestaking.contract, event: "PendingNode", logs: logs, sub: sub}, nil
}

// WatchPendingNode is a free log subscription operation binding the contract event 0xf0a2b450365814b25d2f317ccb08e29c8d1249ee8446074d8112ce30045b6647.
//
// Solidity: event PendingNode(address node, uint256 chainId, uint256 amount, uint256 claimTime)
func (_Nodestaking *NodestakingFilterer) WatchPendingNode(opts *bind.WatchOpts, sink chan<- *NodestakingPendingNode) (event.Subscription, error) {

	logs, sub, err := _Nodestaking.contract.WatchLogs(opts, "PendingNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodestakingPendingNode)
				if err := _Nodestaking.contract.UnpackLog(event, "PendingNode", log); err != nil {
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

// ParsePendingNode is a log parse operation binding the contract event 0xf0a2b450365814b25d2f317ccb08e29c8d1249ee8446074d8112ce30045b6647.
//
// Solidity: event PendingNode(address node, uint256 chainId, uint256 amount, uint256 claimTime)
func (_Nodestaking *NodestakingFilterer) ParsePendingNode(log types.Log) (*NodestakingPendingNode, error) {
	event := new(NodestakingPendingNode)
	if err := _Nodestaking.contract.UnpackLog(event, "PendingNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodestakingSlashedNodeIterator is returned from FilterSlashedNode and is used to iterate over the raw logs and unpacked data for SlashedNode events raised by the Nodestaking contract.
type NodestakingSlashedNodeIterator struct {
	Event *NodestakingSlashedNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodestakingSlashedNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodestakingSlashedNode)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodestakingSlashedNode)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodestakingSlashedNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodestakingSlashedNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodestakingSlashedNode represents a SlashedNode event raised by the Nodestaking contract.
type NodestakingSlashedNode struct {
	Node    common.Address
	ChainId *big.Int
	Amount  *big.Int
	Slasher common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashedNode is a free log retrieval operation binding the contract event 0x14e421663d2498e0453f0b979bc783bc0147feda45f4aae7f3fff60420ec4b2a.
//
// Solidity: event SlashedNode(address node, uint256 chainId, uint256 amount, address slasher)
func (_Nodestaking *NodestakingFilterer) FilterSlashedNode(opts *bind.FilterOpts) (*NodestakingSlashedNodeIterator, error) {

	logs, sub, err := _Nodestaking.contract.FilterLogs(opts, "SlashedNode")
	if err != nil {
		return nil, err
	}
	return &NodestakingSlashedNodeIterator{contract: _Nodestaking.contract, event: "SlashedNode", logs: logs, sub: sub}, nil
}

// WatchSlashedNode is a free log subscription operation binding the contract event 0x14e421663d2498e0453f0b979bc783bc0147feda45f4aae7f3fff60420ec4b2a.
//
// Solidity: event SlashedNode(address node, uint256 chainId, uint256 amount, address slasher)
func (_Nodestaking *NodestakingFilterer) WatchSlashedNode(opts *bind.WatchOpts, sink chan<- *NodestakingSlashedNode) (event.Subscription, error) {

	logs, sub, err := _Nodestaking.contract.WatchLogs(opts, "SlashedNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodestakingSlashedNode)
				if err := _Nodestaking.contract.UnpackLog(event, "SlashedNode", log); err != nil {
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

// ParseSlashedNode is a log parse operation binding the contract event 0x14e421663d2498e0453f0b979bc783bc0147feda45f4aae7f3fff60420ec4b2a.
//
// Solidity: event SlashedNode(address node, uint256 chainId, uint256 amount, address slasher)
func (_Nodestaking *NodestakingFilterer) ParseSlashedNode(log types.Log) (*NodestakingSlashedNode, error) {
	event := new(NodestakingSlashedNode)
	if err := _Nodestaking.contract.UnpackLog(event, "SlashedNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodestakingStakedNodeIterator is returned from FilterStakedNode and is used to iterate over the raw logs and unpacked data for StakedNode events raised by the Nodestaking contract.
type NodestakingStakedNodeIterator struct {
	Event *NodestakingStakedNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodestakingStakedNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodestakingStakedNode)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodestakingStakedNode)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodestakingStakedNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodestakingStakedNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodestakingStakedNode represents a StakedNode event raised by the Nodestaking contract.
type NodestakingStakedNode struct {
	Node      common.Address
	ChainId   *big.Int
	Amount    *big.Int
	ClaimTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakedNode is a free log retrieval operation binding the contract event 0xe876baf068bed217b62c3542001f74ea8d83b5d96ad7e9aecbdbab2c659eff16.
//
// Solidity: event StakedNode(address node, uint256 chainId, uint256 amount, uint256 claimTime)
func (_Nodestaking *NodestakingFilterer) FilterStakedNode(opts *bind.FilterOpts) (*NodestakingStakedNodeIterator, error) {

	logs, sub, err := _Nodestaking.contract.FilterLogs(opts, "StakedNode")
	if err != nil {
		return nil, err
	}
	return &NodestakingStakedNodeIterator{contract: _Nodestaking.contract, event: "StakedNode", logs: logs, sub: sub}, nil
}

// WatchStakedNode is a free log subscription operation binding the contract event 0xe876baf068bed217b62c3542001f74ea8d83b5d96ad7e9aecbdbab2c659eff16.
//
// Solidity: event StakedNode(address node, uint256 chainId, uint256 amount, uint256 claimTime)
func (_Nodestaking *NodestakingFilterer) WatchStakedNode(opts *bind.WatchOpts, sink chan<- *NodestakingStakedNode) (event.Subscription, error) {

	logs, sub, err := _Nodestaking.contract.WatchLogs(opts, "StakedNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodestakingStakedNode)
				if err := _Nodestaking.contract.UnpackLog(event, "StakedNode", log); err != nil {
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

// ParseStakedNode is a log parse operation binding the contract event 0xe876baf068bed217b62c3542001f74ea8d83b5d96ad7e9aecbdbab2c659eff16.
//
// Solidity: event StakedNode(address node, uint256 chainId, uint256 amount, uint256 claimTime)
func (_Nodestaking *NodestakingFilterer) ParseStakedNode(log types.Log) (*NodestakingStakedNode, error) {
	event := new(NodestakingStakedNode)
	if err := _Nodestaking.contract.UnpackLog(event, "StakedNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
