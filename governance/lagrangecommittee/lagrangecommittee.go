// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lagrangecommittee

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

// RLPReaderRLPItem is an auto generated low-level Go binding around an user-defined struct.
type RLPReaderRLPItem struct {
	Len    *big.Int
	MemPtr *big.Int
}

// LagrangecommitteeMetaData contains all meta data concerning the Lagrangecommittee contract.
var LagrangecommitteeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"InitCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"next1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"next2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"next3\",\"type\":\"uint256\"}],\"name\":\"RotateCommittee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNT_CREATION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUTHORISE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_HEADER_NUMBER_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_CURRENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"COMMITTEE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"COMMITTEE_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERMEZ_NETWORK_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"}],\"name\":\"calculateBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"}],\"name\":\"checkAndDecodeRLP\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPtr\",\"type\":\"uint256\"}],\"internalType\":\"structRLPReader.RLPItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_blsPubKey\",\"type\":\"bytes\"}],\"name\":\"committeeAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epoch2committee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getCommitteeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getCommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getCommitteeStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getNextCommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"initCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon4Elements\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeCommitteeAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"rotateCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"comparisonNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"verifyBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LagrangecommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangecommitteeMetaData.ABI instead.
var LagrangecommitteeABI = LagrangecommitteeMetaData.ABI

// Lagrangecommittee is an auto generated Go binding around an Ethereum contract.
type Lagrangecommittee struct {
	LagrangecommitteeCaller     // Read-only binding to the contract
	LagrangecommitteeTransactor // Write-only binding to the contract
	LagrangecommitteeFilterer   // Log filterer for contract events
}

// LagrangecommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagrangecommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangecommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagrangecommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangecommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagrangecommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangecommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagrangecommitteeSession struct {
	Contract     *Lagrangecommittee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LagrangecommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagrangecommitteeCallerSession struct {
	Contract *LagrangecommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LagrangecommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagrangecommitteeTransactorSession struct {
	Contract     *LagrangecommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LagrangecommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagrangecommitteeRaw struct {
	Contract *Lagrangecommittee // Generic contract binding to access the raw methods on
}

// LagrangecommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagrangecommitteeCallerRaw struct {
	Contract *LagrangecommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// LagrangecommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagrangecommitteeTransactorRaw struct {
	Contract *LagrangecommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLagrangecommittee creates a new instance of Lagrangecommittee, bound to a specific deployed contract.
func NewLagrangecommittee(address common.Address, backend bind.ContractBackend) (*Lagrangecommittee, error) {
	contract, err := bindLagrangecommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lagrangecommittee{LagrangecommitteeCaller: LagrangecommitteeCaller{contract: contract}, LagrangecommitteeTransactor: LagrangecommitteeTransactor{contract: contract}, LagrangecommitteeFilterer: LagrangecommitteeFilterer{contract: contract}}, nil
}

// NewLagrangecommitteeCaller creates a new read-only instance of Lagrangecommittee, bound to a specific deployed contract.
func NewLagrangecommitteeCaller(address common.Address, caller bind.ContractCaller) (*LagrangecommitteeCaller, error) {
	contract, err := bindLagrangecommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeCaller{contract: contract}, nil
}

// NewLagrangecommitteeTransactor creates a new write-only instance of Lagrangecommittee, bound to a specific deployed contract.
func NewLagrangecommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*LagrangecommitteeTransactor, error) {
	contract, err := bindLagrangecommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeTransactor{contract: contract}, nil
}

// NewLagrangecommitteeFilterer creates a new log filterer instance of Lagrangecommittee, bound to a specific deployed contract.
func NewLagrangecommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*LagrangecommitteeFilterer, error) {
	contract, err := bindLagrangecommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeFilterer{contract: contract}, nil
}

// bindLagrangecommittee binds a generic wrapper to an already deployed contract.
func bindLagrangecommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LagrangecommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lagrangecommittee *LagrangecommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lagrangecommittee.Contract.LagrangecommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lagrangecommittee *LagrangecommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.LagrangecommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lagrangecommittee *LagrangecommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.LagrangecommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lagrangecommittee *LagrangecommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lagrangecommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lagrangecommittee *LagrangecommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lagrangecommittee *LagrangecommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) ACCOUNTCREATIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "ACCOUNT_CREATION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.ACCOUNTCREATIONHASH(&_Lagrangecommittee.CallOpts)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.ACCOUNTCREATIONHASH(&_Lagrangecommittee.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) AUTHORISETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "AUTHORISE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.AUTHORISETYPEHASH(&_Lagrangecommittee.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.AUTHORISETYPEHASH(&_Lagrangecommittee.CallOpts)
}

// BLOCKHEADERNUMBERINDEX is a free data retrieval call binding the contract method 0x14501001.
//
// Solidity: function BLOCK_HEADER_NUMBER_INDEX() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) BLOCKHEADERNUMBERINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "BLOCK_HEADER_NUMBER_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKHEADERNUMBERINDEX is a free data retrieval call binding the contract method 0x14501001.
//
// Solidity: function BLOCK_HEADER_NUMBER_INDEX() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) BLOCKHEADERNUMBERINDEX() (*big.Int, error) {
	return _Lagrangecommittee.Contract.BLOCKHEADERNUMBERINDEX(&_Lagrangecommittee.CallOpts)
}

// BLOCKHEADERNUMBERINDEX is a free data retrieval call binding the contract method 0x14501001.
//
// Solidity: function BLOCK_HEADER_NUMBER_INDEX() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) BLOCKHEADERNUMBERINDEX() (*big.Int, error) {
	return _Lagrangecommittee.Contract.BLOCKHEADERNUMBERINDEX(&_Lagrangecommittee.CallOpts)
}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) COMMITTEECURRENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "COMMITTEE_CURRENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) COMMITTEECURRENT() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEECURRENT(&_Lagrangecommittee.CallOpts)
}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) COMMITTEECURRENT() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEECURRENT(&_Lagrangecommittee.CallOpts)
}

// COMMITTEEDURATION is a free data retrieval call binding the contract method 0x1c04e7f1.
//
// Solidity: function COMMITTEE_DURATION(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) COMMITTEEDURATION(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "COMMITTEE_DURATION", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEEDURATION is a free data retrieval call binding the contract method 0x1c04e7f1.
//
// Solidity: function COMMITTEE_DURATION(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) COMMITTEEDURATION(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEEDURATION(&_Lagrangecommittee.CallOpts, arg0)
}

// COMMITTEEDURATION is a free data retrieval call binding the contract method 0x1c04e7f1.
//
// Solidity: function COMMITTEE_DURATION(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) COMMITTEEDURATION(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEEDURATION(&_Lagrangecommittee.CallOpts, arg0)
}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) COMMITTEENEXT1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "COMMITTEE_NEXT_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) COMMITTEENEXT1() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEENEXT1(&_Lagrangecommittee.CallOpts)
}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) COMMITTEENEXT1() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEENEXT1(&_Lagrangecommittee.CallOpts)
}

// COMMITTEENEXT2 is a free data retrieval call binding the contract method 0xe04392f1.
//
// Solidity: function COMMITTEE_NEXT_2() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) COMMITTEENEXT2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "COMMITTEE_NEXT_2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEENEXT2 is a free data retrieval call binding the contract method 0xe04392f1.
//
// Solidity: function COMMITTEE_NEXT_2() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) COMMITTEENEXT2() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEENEXT2(&_Lagrangecommittee.CallOpts)
}

// COMMITTEENEXT2 is a free data retrieval call binding the contract method 0xe04392f1.
//
// Solidity: function COMMITTEE_NEXT_2() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) COMMITTEENEXT2() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEENEXT2(&_Lagrangecommittee.CallOpts)
}

// COMMITTEENEXT3 is a free data retrieval call binding the contract method 0x50643b60.
//
// Solidity: function COMMITTEE_NEXT_3() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) COMMITTEENEXT3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "COMMITTEE_NEXT_3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEENEXT3 is a free data retrieval call binding the contract method 0x50643b60.
//
// Solidity: function COMMITTEE_NEXT_3() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) COMMITTEENEXT3() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEENEXT3(&_Lagrangecommittee.CallOpts)
}

// COMMITTEENEXT3 is a free data retrieval call binding the contract method 0x50643b60.
//
// Solidity: function COMMITTEE_NEXT_3() view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) COMMITTEENEXT3() (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEENEXT3(&_Lagrangecommittee.CallOpts)
}

// COMMITTEESTART is a free data retrieval call binding the contract method 0x946a55cf.
//
// Solidity: function COMMITTEE_START(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) COMMITTEESTART(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "COMMITTEE_START", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEESTART is a free data retrieval call binding the contract method 0x946a55cf.
//
// Solidity: function COMMITTEE_START(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) COMMITTEESTART(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEESTART(&_Lagrangecommittee.CallOpts, arg0)
}

// COMMITTEESTART is a free data retrieval call binding the contract method 0x946a55cf.
//
// Solidity: function COMMITTEE_START(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) COMMITTEESTART(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.COMMITTEESTART(&_Lagrangecommittee.CallOpts, arg0)
}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_Lagrangecommittee *LagrangecommitteeCaller) CommitteeMap(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "CommitteeMap", arg0, arg1)

	outstruct := new(struct {
		Addr      common.Address
		Stake     *big.Int
		BlsPubKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlsPubKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_Lagrangecommittee *LagrangecommitteeSession) CommitteeMap(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	return _Lagrangecommittee.Contract.CommitteeMap(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CommitteeMap(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	return _Lagrangecommittee.Contract.CommitteeMap(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) CommitteeMapKeys(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "CommitteeMapKeys", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) CommitteeMapKeys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeMapKeys(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CommitteeMapKeys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeMapKeys(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) CommitteeMapLength(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "CommitteeMapLength", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) CommitteeMapLength(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeMapLength(&_Lagrangecommittee.CallOpts, arg0)
}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CommitteeMapLength(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeMapLength(&_Lagrangecommittee.CallOpts, arg0)
}

// CommitteeNodes is a free data retrieval call binding the contract method 0x13e8f363.
//
// Solidity: function CommitteeNodes(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) CommitteeNodes(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "CommitteeNodes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeNodes is a free data retrieval call binding the contract method 0x13e8f363.
//
// Solidity: function CommitteeNodes(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) CommitteeNodes(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeNodes(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeNodes is a free data retrieval call binding the contract method 0x13e8f363.
//
// Solidity: function CommitteeNodes(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CommitteeNodes(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeNodes(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeRoot is a free data retrieval call binding the contract method 0xba1efb58.
//
// Solidity: function CommitteeRoot(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) CommitteeRoot(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "CommitteeRoot", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeRoot is a free data retrieval call binding the contract method 0xba1efb58.
//
// Solidity: function CommitteeRoot(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) CommitteeRoot(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeRoot(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// CommitteeRoot is a free data retrieval call binding the contract method 0xba1efb58.
//
// Solidity: function CommitteeRoot(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CommitteeRoot(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.CommitteeRoot(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Lagrangecommittee *LagrangecommitteeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Lagrangecommittee *LagrangecommitteeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Lagrangecommittee.Contract.DOMAINSEPARATOR(&_Lagrangecommittee.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Lagrangecommittee.Contract.DOMAINSEPARATOR(&_Lagrangecommittee.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "EIP712DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.EIP712DOMAINHASH(&_Lagrangecommittee.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.EIP712DOMAINHASH(&_Lagrangecommittee.CallOpts)
}

// EpochNumber is a free data retrieval call binding the contract method 0x2f374006.
//
// Solidity: function EpochNumber(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) EpochNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "EpochNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochNumber is a free data retrieval call binding the contract method 0x2f374006.
//
// Solidity: function EpochNumber(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) EpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.EpochNumber(&_Lagrangecommittee.CallOpts, arg0)
}

// EpochNumber is a free data retrieval call binding the contract method 0x2f374006.
//
// Solidity: function EpochNumber(uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) EpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.EpochNumber(&_Lagrangecommittee.CallOpts, arg0)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) HERMEZNETWORKHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "HERMEZ_NETWORK_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.HERMEZNETWORKHASH(&_Lagrangecommittee.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.HERMEZNETWORKHASH(&_Lagrangecommittee.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) NAMEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "NAME_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) NAMEHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.NAMEHASH(&_Lagrangecommittee.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) NAMEHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.NAMEHASH(&_Lagrangecommittee.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) VERSIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "VERSION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) VERSIONHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.VERSIONHASH(&_Lagrangecommittee.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) VERSIONHASH() ([32]byte, error) {
	return _Lagrangecommittee.Contract.VERSIONHASH(&_Lagrangecommittee.CallOpts)
}

// CalculateBlockHash is a free data retrieval call binding the contract method 0xf44c5c71.
//
// Solidity: function calculateBlockHash(bytes rlpData) pure returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCaller) CalculateBlockHash(opts *bind.CallOpts, rlpData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "calculateBlockHash", rlpData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateBlockHash is a free data retrieval call binding the contract method 0xf44c5c71.
//
// Solidity: function calculateBlockHash(bytes rlpData) pure returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeSession) CalculateBlockHash(rlpData []byte) ([32]byte, error) {
	return _Lagrangecommittee.Contract.CalculateBlockHash(&_Lagrangecommittee.CallOpts, rlpData)
}

// CalculateBlockHash is a free data retrieval call binding the contract method 0xf44c5c71.
//
// Solidity: function calculateBlockHash(bytes rlpData) pure returns(bytes32)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CalculateBlockHash(rlpData []byte) ([32]byte, error) {
	return _Lagrangecommittee.Contract.CalculateBlockHash(&_Lagrangecommittee.CallOpts, rlpData)
}

// CheckAndDecodeRLP is a free data retrieval call binding the contract method 0xaef524b0.
//
// Solidity: function checkAndDecodeRLP(bytes rlpData, bytes32 comparisonBlockHash) view returns((uint256,uint256)[])
func (_Lagrangecommittee *LagrangecommitteeCaller) CheckAndDecodeRLP(opts *bind.CallOpts, rlpData []byte, comparisonBlockHash [32]byte) ([]RLPReaderRLPItem, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "checkAndDecodeRLP", rlpData, comparisonBlockHash)

	if err != nil {
		return *new([]RLPReaderRLPItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]RLPReaderRLPItem)).(*[]RLPReaderRLPItem)

	return out0, err

}

// CheckAndDecodeRLP is a free data retrieval call binding the contract method 0xaef524b0.
//
// Solidity: function checkAndDecodeRLP(bytes rlpData, bytes32 comparisonBlockHash) view returns((uint256,uint256)[])
func (_Lagrangecommittee *LagrangecommitteeSession) CheckAndDecodeRLP(rlpData []byte, comparisonBlockHash [32]byte) ([]RLPReaderRLPItem, error) {
	return _Lagrangecommittee.Contract.CheckAndDecodeRLP(&_Lagrangecommittee.CallOpts, rlpData, comparisonBlockHash)
}

// CheckAndDecodeRLP is a free data retrieval call binding the contract method 0xaef524b0.
//
// Solidity: function checkAndDecodeRLP(bytes rlpData, bytes32 comparisonBlockHash) view returns((uint256,uint256)[])
func (_Lagrangecommittee *LagrangecommitteeCallerSession) CheckAndDecodeRLP(rlpData []byte, comparisonBlockHash [32]byte) ([]RLPReaderRLPItem, error) {
	return _Lagrangecommittee.Contract.CheckAndDecodeRLP(&_Lagrangecommittee.CallOpts, rlpData, comparisonBlockHash)
}

// Epoch2committee is a free data retrieval call binding the contract method 0x0c01dc12.
//
// Solidity: function epoch2committee(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) Epoch2committee(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "epoch2committee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch2committee is a free data retrieval call binding the contract method 0x0c01dc12.
//
// Solidity: function epoch2committee(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) Epoch2committee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.Epoch2committee(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// Epoch2committee is a free data retrieval call binding the contract method 0x0c01dc12.
//
// Solidity: function epoch2committee(uint256 , uint256 ) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) Epoch2committee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.Epoch2committee(&_Lagrangecommittee.CallOpts, arg0, arg1)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Lagrangecommittee *LagrangecommitteeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Lagrangecommittee *LagrangecommitteeSession) GetChainId() (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetChainId(&_Lagrangecommittee.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) GetChainId() (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetChainId(&_Lagrangecommittee.CallOpts)
}

// GetCommitteeDuration is a free data retrieval call binding the contract method 0x64a83b4b.
//
// Solidity: function getCommitteeDuration(uint256 chainID) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) GetCommitteeDuration(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "getCommitteeDuration", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitteeDuration is a free data retrieval call binding the contract method 0x64a83b4b.
//
// Solidity: function getCommitteeDuration(uint256 chainID) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) GetCommitteeDuration(chainID *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetCommitteeDuration(&_Lagrangecommittee.CallOpts, chainID)
}

// GetCommitteeDuration is a free data retrieval call binding the contract method 0x64a83b4b.
//
// Solidity: function getCommitteeDuration(uint256 chainID) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) GetCommitteeDuration(chainID *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetCommitteeDuration(&_Lagrangecommittee.CallOpts, chainID)
}

// GetCommitteeRoot is a free data retrieval call binding the contract method 0xe00aab1a.
//
// Solidity: function getCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) GetCommitteeRoot(opts *bind.CallOpts, chainID *big.Int, _epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "getCommitteeRoot", chainID, _epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitteeRoot is a free data retrieval call binding the contract method 0xe00aab1a.
//
// Solidity: function getCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) GetCommitteeRoot(chainID *big.Int, _epoch *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetCommitteeRoot(&_Lagrangecommittee.CallOpts, chainID, _epoch)
}

// GetCommitteeRoot is a free data retrieval call binding the contract method 0xe00aab1a.
//
// Solidity: function getCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) GetCommitteeRoot(chainID *big.Int, _epoch *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetCommitteeRoot(&_Lagrangecommittee.CallOpts, chainID, _epoch)
}

// GetCommitteeStart is a free data retrieval call binding the contract method 0x4d008f4d.
//
// Solidity: function getCommitteeStart(uint256 chainID) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) GetCommitteeStart(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "getCommitteeStart", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitteeStart is a free data retrieval call binding the contract method 0x4d008f4d.
//
// Solidity: function getCommitteeStart(uint256 chainID) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) GetCommitteeStart(chainID *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetCommitteeStart(&_Lagrangecommittee.CallOpts, chainID)
}

// GetCommitteeStart is a free data retrieval call binding the contract method 0x4d008f4d.
//
// Solidity: function getCommitteeStart(uint256 chainID) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) GetCommitteeStart(chainID *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetCommitteeStart(&_Lagrangecommittee.CallOpts, chainID)
}

// GetNextCommitteeRoot is a free data retrieval call binding the contract method 0x2e253dcc.
//
// Solidity: function getNextCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCaller) GetNextCommitteeRoot(opts *bind.CallOpts, chainID *big.Int, _epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "getNextCommitteeRoot", chainID, _epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCommitteeRoot is a free data retrieval call binding the contract method 0x2e253dcc.
//
// Solidity: function getNextCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeSession) GetNextCommitteeRoot(chainID *big.Int, _epoch *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetNextCommitteeRoot(&_Lagrangecommittee.CallOpts, chainID, _epoch)
}

// GetNextCommitteeRoot is a free data retrieval call binding the contract method 0x2e253dcc.
//
// Solidity: function getNextCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(uint256)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) GetNextCommitteeRoot(chainID *big.Int, _epoch *big.Int) (*big.Int, error) {
	return _Lagrangecommittee.Contract.GetNextCommitteeRoot(&_Lagrangecommittee.CallOpts, chainID, _epoch)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrangecommittee *LagrangecommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrangecommittee *LagrangecommitteeSession) Owner() (common.Address, error) {
	return _Lagrangecommittee.Contract.Owner(&_Lagrangecommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) Owner() (common.Address, error) {
	return _Lagrangecommittee.Contract.Owner(&_Lagrangecommittee.CallOpts)
}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) view returns(bool)
func (_Lagrangecommittee *LagrangecommitteeCaller) VerifyBlockNumber(opts *bind.CallOpts, comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Lagrangecommittee.contract.Call(opts, &out, "verifyBlockNumber", comparisonNumber, rlpData, comparisonBlockHash, chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) view returns(bool)
func (_Lagrangecommittee *LagrangecommitteeSession) VerifyBlockNumber(comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	return _Lagrangecommittee.Contract.VerifyBlockNumber(&_Lagrangecommittee.CallOpts, comparisonNumber, rlpData, comparisonBlockHash, chainID)
}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) view returns(bool)
func (_Lagrangecommittee *LagrangecommitteeCallerSession) VerifyBlockNumber(comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	return _Lagrangecommittee.Contract.VerifyBlockNumber(&_Lagrangecommittee.CallOpts, comparisonNumber, rlpData, comparisonBlockHash, chainID)
}

// CommitteeAdd is a paid mutator transaction binding the contract method 0x40dda085.
//
// Solidity: function committeeAdd(uint256 chainID, uint256 stake, bytes _blsPubKey) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) CommitteeAdd(opts *bind.TransactOpts, chainID *big.Int, stake *big.Int, _blsPubKey []byte) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "committeeAdd", chainID, stake, _blsPubKey)
}

// CommitteeAdd is a paid mutator transaction binding the contract method 0x40dda085.
//
// Solidity: function committeeAdd(uint256 chainID, uint256 stake, bytes _blsPubKey) returns()
func (_Lagrangecommittee *LagrangecommitteeSession) CommitteeAdd(chainID *big.Int, stake *big.Int, _blsPubKey []byte) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.CommitteeAdd(&_Lagrangecommittee.TransactOpts, chainID, stake, _blsPubKey)
}

// CommitteeAdd is a paid mutator transaction binding the contract method 0x40dda085.
//
// Solidity: function committeeAdd(uint256 chainID, uint256 stake, bytes _blsPubKey) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) CommitteeAdd(chainID *big.Int, stake *big.Int, _blsPubKey []byte) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.CommitteeAdd(&_Lagrangecommittee.TransactOpts, chainID, stake, _blsPubKey)
}

// InitCommittee is a paid mutator transaction binding the contract method 0x40d810b4.
//
// Solidity: function initCommittee(uint256 _chainID, uint256 _duration) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) InitCommittee(opts *bind.TransactOpts, _chainID *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "initCommittee", _chainID, _duration)
}

// InitCommittee is a paid mutator transaction binding the contract method 0x40d810b4.
//
// Solidity: function initCommittee(uint256 _chainID, uint256 _duration) returns()
func (_Lagrangecommittee *LagrangecommitteeSession) InitCommittee(_chainID *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.InitCommittee(&_Lagrangecommittee.TransactOpts, _chainID, _duration)
}

// InitCommittee is a paid mutator transaction binding the contract method 0x40d810b4.
//
// Solidity: function initCommittee(uint256 _chainID, uint256 _duration) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) InitCommittee(_chainID *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.InitCommittee(&_Lagrangecommittee.TransactOpts, _chainID, _duration)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) Initialize(opts *bind.TransactOpts, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "initialize", _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Lagrangecommittee *LagrangecommitteeSession) Initialize(_poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.Initialize(&_Lagrangecommittee.TransactOpts, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) Initialize(_poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.Initialize(&_Lagrangecommittee.TransactOpts, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// RemoveCommitteeAddr is a paid mutator transaction binding the contract method 0x333edc3f.
//
// Solidity: function removeCommitteeAddr(uint256 chainID) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) RemoveCommitteeAddr(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "removeCommitteeAddr", chainID)
}

// RemoveCommitteeAddr is a paid mutator transaction binding the contract method 0x333edc3f.
//
// Solidity: function removeCommitteeAddr(uint256 chainID) returns()
func (_Lagrangecommittee *LagrangecommitteeSession) RemoveCommitteeAddr(chainID *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.RemoveCommitteeAddr(&_Lagrangecommittee.TransactOpts, chainID)
}

// RemoveCommitteeAddr is a paid mutator transaction binding the contract method 0x333edc3f.
//
// Solidity: function removeCommitteeAddr(uint256 chainID) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) RemoveCommitteeAddr(chainID *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.RemoveCommitteeAddr(&_Lagrangecommittee.TransactOpts, chainID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrangecommittee *LagrangecommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.RenounceOwnership(&_Lagrangecommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.RenounceOwnership(&_Lagrangecommittee.TransactOpts)
}

// RotateCommittee is a paid mutator transaction binding the contract method 0x759f5b5a.
//
// Solidity: function rotateCommittee(uint256 chainID) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) RotateCommittee(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "rotateCommittee", chainID)
}

// RotateCommittee is a paid mutator transaction binding the contract method 0x759f5b5a.
//
// Solidity: function rotateCommittee(uint256 chainID) returns()
func (_Lagrangecommittee *LagrangecommitteeSession) RotateCommittee(chainID *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.RotateCommittee(&_Lagrangecommittee.TransactOpts, chainID)
}

// RotateCommittee is a paid mutator transaction binding the contract method 0x759f5b5a.
//
// Solidity: function rotateCommittee(uint256 chainID) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) RotateCommittee(chainID *big.Int) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.RotateCommittee(&_Lagrangecommittee.TransactOpts, chainID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lagrangecommittee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrangecommittee *LagrangecommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.TransferOwnership(&_Lagrangecommittee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrangecommittee *LagrangecommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lagrangecommittee.Contract.TransferOwnership(&_Lagrangecommittee.TransactOpts, newOwner)
}

// LagrangecommitteeInitCommitteeIterator is returned from FilterInitCommittee and is used to iterate over the raw logs and unpacked data for InitCommittee events raised by the Lagrangecommittee contract.
type LagrangecommitteeInitCommitteeIterator struct {
	Event *LagrangecommitteeInitCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangecommitteeInitCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangecommitteeInitCommittee)
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
		it.Event = new(LagrangecommitteeInitCommittee)
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
func (it *LagrangecommitteeInitCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangecommitteeInitCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangecommitteeInitCommittee represents a InitCommittee event raised by the Lagrangecommittee contract.
type LagrangecommitteeInitCommittee struct {
	ChainID  *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitCommittee is a free log retrieval operation binding the contract event 0x6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration)
func (_Lagrangecommittee *LagrangecommitteeFilterer) FilterInitCommittee(opts *bind.FilterOpts) (*LagrangecommitteeInitCommitteeIterator, error) {

	logs, sub, err := _Lagrangecommittee.contract.FilterLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeInitCommitteeIterator{contract: _Lagrangecommittee.contract, event: "InitCommittee", logs: logs, sub: sub}, nil
}

// WatchInitCommittee is a free log subscription operation binding the contract event 0x6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration)
func (_Lagrangecommittee *LagrangecommitteeFilterer) WatchInitCommittee(opts *bind.WatchOpts, sink chan<- *LagrangecommitteeInitCommittee) (event.Subscription, error) {

	logs, sub, err := _Lagrangecommittee.contract.WatchLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangecommitteeInitCommittee)
				if err := _Lagrangecommittee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
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

// ParseInitCommittee is a log parse operation binding the contract event 0x6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration)
func (_Lagrangecommittee *LagrangecommitteeFilterer) ParseInitCommittee(log types.Log) (*LagrangecommitteeInitCommittee, error) {
	event := new(LagrangecommitteeInitCommittee)
	if err := _Lagrangecommittee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangecommitteeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lagrangecommittee contract.
type LagrangecommitteeInitializedIterator struct {
	Event *LagrangecommitteeInitialized // Event containing the contract specifics and raw log

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
func (it *LagrangecommitteeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangecommitteeInitialized)
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
		it.Event = new(LagrangecommitteeInitialized)
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
func (it *LagrangecommitteeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangecommitteeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangecommitteeInitialized represents a Initialized event raised by the Lagrangecommittee contract.
type LagrangecommitteeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrangecommittee *LagrangecommitteeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LagrangecommitteeInitializedIterator, error) {

	logs, sub, err := _Lagrangecommittee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeInitializedIterator{contract: _Lagrangecommittee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrangecommittee *LagrangecommitteeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LagrangecommitteeInitialized) (event.Subscription, error) {

	logs, sub, err := _Lagrangecommittee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangecommitteeInitialized)
				if err := _Lagrangecommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Lagrangecommittee *LagrangecommitteeFilterer) ParseInitialized(log types.Log) (*LagrangecommitteeInitialized, error) {
	event := new(LagrangecommitteeInitialized)
	if err := _Lagrangecommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangecommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lagrangecommittee contract.
type LagrangecommitteeOwnershipTransferredIterator struct {
	Event *LagrangecommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LagrangecommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangecommitteeOwnershipTransferred)
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
		it.Event = new(LagrangecommitteeOwnershipTransferred)
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
func (it *LagrangecommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangecommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangecommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the Lagrangecommittee contract.
type LagrangecommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrangecommittee *LagrangecommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LagrangecommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lagrangecommittee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeOwnershipTransferredIterator{contract: _Lagrangecommittee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrangecommittee *LagrangecommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LagrangecommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lagrangecommittee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangecommitteeOwnershipTransferred)
				if err := _Lagrangecommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lagrangecommittee *LagrangecommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*LagrangecommitteeOwnershipTransferred, error) {
	event := new(LagrangecommitteeOwnershipTransferred)
	if err := _Lagrangecommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangecommitteeRotateCommitteeIterator is returned from FilterRotateCommittee and is used to iterate over the raw logs and unpacked data for RotateCommittee events raised by the Lagrangecommittee contract.
type LagrangecommitteeRotateCommitteeIterator struct {
	Event *LagrangecommitteeRotateCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangecommitteeRotateCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangecommitteeRotateCommittee)
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
		it.Event = new(LagrangecommitteeRotateCommittee)
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
func (it *LagrangecommitteeRotateCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangecommitteeRotateCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangecommitteeRotateCommittee represents a RotateCommittee event raised by the Lagrangecommittee contract.
type LagrangecommitteeRotateCommittee struct {
	ChainID *big.Int
	Current *big.Int
	Next1   *big.Int
	Next2   *big.Int
	Next3   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRotateCommittee is a free log retrieval operation binding the contract event 0x34b7d9871fbc7cb893d9c8368fc02e0e7909edfc64e1b57de99103770bedd598.
//
// Solidity: event RotateCommittee(uint256 chainID, uint256 current, uint256 next1, uint256 next2, uint256 next3)
func (_Lagrangecommittee *LagrangecommitteeFilterer) FilterRotateCommittee(opts *bind.FilterOpts) (*LagrangecommitteeRotateCommitteeIterator, error) {

	logs, sub, err := _Lagrangecommittee.contract.FilterLogs(opts, "RotateCommittee")
	if err != nil {
		return nil, err
	}
	return &LagrangecommitteeRotateCommitteeIterator{contract: _Lagrangecommittee.contract, event: "RotateCommittee", logs: logs, sub: sub}, nil
}

// WatchRotateCommittee is a free log subscription operation binding the contract event 0x34b7d9871fbc7cb893d9c8368fc02e0e7909edfc64e1b57de99103770bedd598.
//
// Solidity: event RotateCommittee(uint256 chainID, uint256 current, uint256 next1, uint256 next2, uint256 next3)
func (_Lagrangecommittee *LagrangecommitteeFilterer) WatchRotateCommittee(opts *bind.WatchOpts, sink chan<- *LagrangecommitteeRotateCommittee) (event.Subscription, error) {

	logs, sub, err := _Lagrangecommittee.contract.WatchLogs(opts, "RotateCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangecommitteeRotateCommittee)
				if err := _Lagrangecommittee.contract.UnpackLog(event, "RotateCommittee", log); err != nil {
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

// ParseRotateCommittee is a log parse operation binding the contract event 0x34b7d9871fbc7cb893d9c8368fc02e0e7909edfc64e1b57de99103770bedd598.
//
// Solidity: event RotateCommittee(uint256 chainID, uint256 current, uint256 next1, uint256 next2, uint256 next3)
func (_Lagrangecommittee *LagrangecommitteeFilterer) ParseRotateCommittee(log types.Log) (*LagrangecommitteeRotateCommittee, error) {
	event := new(LagrangecommitteeRotateCommittee)
	if err := _Lagrangecommittee.contract.UnpackLog(event, "RotateCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
