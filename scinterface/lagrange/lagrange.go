// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lagrange

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

// LagrangeServiceEvidence is an auto generated low-level Go binding around an user-defined struct.
type LagrangeServiceEvidence struct {
	Operator                    common.Address
	BlockHash                   [32]byte
	CorrectBlockHash            [32]byte
	CurrentCommitteeRoot        [32]byte
	CorrectCurrentCommitteeRoot [32]byte
	NextCommitteeRoot           [32]byte
	CorrectNextCommitteeRoot    [32]byte
	BlockNumber                 *big.Int
	EpochNumber                 *big.Int
	BlockSignature              []byte
	CommitSignature             []byte
	ChainID                     uint32
}

// LagrangeMetaData contains all meta data concerning the Lagrange contract.
var LagrangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"_slasher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"UploadEvidence\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestServeUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"recordLastStakeUpdateAndRevokeSlashingAbility\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"updateBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"prevElement\",\"type\":\"uint256\"}],\"name\":\"recordStakeUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"internalType\":\"structLagrangeService.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"uploadEvidence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600280546001600160401b031916905534801561002057600080fd5b50604051610c30380380610c3083398101604081905261003f916100a9565b61004833610059565b6001600160a01b03166080526100d9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bb57600080fd5b81516001600160a01b03811681146100d257600080fd5b9392505050565b608051610b19610117600039600081816101c40152818161025a0152818161055f015281816105e3015281816106fa01526107480152610b196000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063758f8dba11610071578063758f8dba146101825780638da5cb5b1461019a578063b1344271146101bf578063c747075b146101e6578063e5839836146101f9578063f2fde38b1461021c57600080fd5b80630ffabbce146100b9578063130d7906146100ce57806313e7c9d8146100e15780634495c7e914610142578063715018a61461015557806372d18e8d1461015d575b600080fd5b6100cc6100c73660046108dd565b61022f565b005b6100cc6100dc366004610910565b6102bb565b61011b6100ef366004610932565b6001602081905260009182526040909120805491015463ffffffff811690640100000000900460ff1683565b6040805193845263ffffffff90921660208401521515908201526060015b60405180910390f35b6100cc61015036600461094d565b610362565b6100cc610511565b60025461016d9063ffffffff1681565b60405163ffffffff9091168152602001610139565b60025461016d90640100000000900463ffffffff1681565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610139565b6101a77f000000000000000000000000000000000000000000000000000000000000000081565b6100cc6101f4366004610989565b610525565b61020c610207366004610932565b6105c1565b6040519015158152602001610139565b6100cc61022a366004610932565b610656565b6040516307fd5de760e11b81526001600160a01b03838116600483015263ffffffff831660248301527f00000000000000000000000000000000000000000000000000000000000000001690630ffabbce906044015b600060405180830381600087803b15801561029f57600080fd5b505af11580156102b3573d6000803e3d6000fd5b505050505050565b6102c533826106cf565b60408051606081018252600180825263ffffffff848116602080850182815260008688018181523380835287855291899020975188559151969095018054915115156401000000000264ffffffffff199092169690941695909517949094179091558351918252918101919091527f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a91015b60405180910390a150565b60006001816103746020850185610932565b6001600160a01b0316815260208101919091526040016000206001015463ffffffff16116103e95760405162461bcd60e51b815260206004820152601e60248201527f546865206f70657261746f72206973206e6f742072656769737465726564000060448201526064015b60405180910390fd5b600160006103fa6020840184610932565b6001600160a01b03168152602081019190915260400160002060010154640100000000900460ff161561046f5760405162461bcd60e51b815260206004820152601760248201527f546865206f70657261746f7220697320736c617368656400000000000000000060448201526064016103e0565b61048461047f6020830183610932565b610729565b7fa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f6104b26020830183610932565b6020830135606084013560a085013560e08601356101008701356104da6101208901896109d4565b6104e86101408b018b6109d4565b6104fa6101808d016101608e01610910565b6040516103579b9a99989796959493929190610a4b565b610519610803565b610523600061085d565b565b60405163c747075b60e01b81526001600160a01b03858116600483015263ffffffff808616602484015284166044830152606482018390527f0000000000000000000000000000000000000000000000000000000000000000169063c747075b90608401600060405180830381600087803b1580156105a357600080fd5b505af11580156105b7573d6000803e3d6000fd5b5050505050505050565b6040516372c1cc1b60e11b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e583983690602401602060405180830381865afa15801561062c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106509190610ac1565b92915050565b61065e610803565b6001600160a01b0381166106c35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103e0565b6106cc8161085d565b50565b60405163175d320560e01b81526001600160a01b03838116600483015263ffffffff831660248301527f0000000000000000000000000000000000000000000000000000000000000000169063175d320590604401610285565b604051630e323b9960e21b81526001600160a01b0382811660048301527f000000000000000000000000000000000000000000000000000000000000000016906338c8ee6490602401600060405180830381600087803b15801561078c57600080fd5b505af11580156107a0573d6000803e3d6000fd5b505050506001600160a01b038116600081815260016020818152604092839020909101805464ff00000000191664010000000017905590519182527fd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a30689101610357565b6000546001600160a01b031633146105235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103e0565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b03811681146108c457600080fd5b919050565b803563ffffffff811681146108c457600080fd5b600080604083850312156108f057600080fd5b6108f9836108ad565b9150610907602084016108c9565b90509250929050565b60006020828403121561092257600080fd5b61092b826108c9565b9392505050565b60006020828403121561094457600080fd5b61092b826108ad565b60006020828403121561095f57600080fd5b813567ffffffffffffffff81111561097657600080fd5b8201610180818503121561092b57600080fd5b6000806000806080858703121561099f57600080fd5b6109a8856108ad565b93506109b6602086016108c9565b92506109c4604086016108c9565b9396929550929360600135925050565b6000808335601e198436030181126109eb57600080fd5b83018035915067ffffffffffffffff821115610a0657600080fd5b602001915036819003821315610a1b57600080fd5b9250929050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061012060018060a01b038e1683528c60208401528b60408401528a60608401528960808401528860a08401528060c0840152610a8c818401888a610a22565b905082810360e0840152610aa1818688610a22565b91505063ffffffff83166101008301529c9b505050505050505050505050565b600060208284031215610ad357600080fd5b8151801515811461092b57600080fdfea264697066735822122059838f5c84cf44de24fca8dcfa5a939caafd94a09ae7ec1a0073dc298ac9af8764736f6c634300080c0033",
}

// LagrangeABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangeMetaData.ABI instead.
var LagrangeABI = LagrangeMetaData.ABI

// LagrangeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LagrangeMetaData.Bin instead.
var LagrangeBin = LagrangeMetaData.Bin

// DeployLagrange deploys a new Ethereum contract, binding an instance of Lagrange to it.
func DeployLagrange(auth *bind.TransactOpts, backend bind.ContractBackend, _slasher common.Address) (common.Address, *types.Transaction, *Lagrange, error) {
	parsed, err := LagrangeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LagrangeBin), backend, _slasher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lagrange{LagrangeCaller: LagrangeCaller{contract: contract}, LagrangeTransactor: LagrangeTransactor{contract: contract}, LagrangeFilterer: LagrangeFilterer{contract: contract}}, nil
}

// Lagrange is an auto generated Go binding around an Ethereum contract.
type Lagrange struct {
	LagrangeCaller     // Read-only binding to the contract
	LagrangeTransactor // Write-only binding to the contract
	LagrangeFilterer   // Log filterer for contract events
}

// LagrangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagrangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagrangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagrangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagrangeSession struct {
	Contract     *Lagrange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LagrangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagrangeCallerSession struct {
	Contract *LagrangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LagrangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagrangeTransactorSession struct {
	Contract     *LagrangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LagrangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagrangeRaw struct {
	Contract *Lagrange // Generic contract binding to access the raw methods on
}

// LagrangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagrangeCallerRaw struct {
	Contract *LagrangeCaller // Generic read-only contract binding to access the raw methods on
}

// LagrangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagrangeTransactorRaw struct {
	Contract *LagrangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLagrange creates a new instance of Lagrange, bound to a specific deployed contract.
func NewLagrange(address common.Address, backend bind.ContractBackend) (*Lagrange, error) {
	contract, err := bindLagrange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lagrange{LagrangeCaller: LagrangeCaller{contract: contract}, LagrangeTransactor: LagrangeTransactor{contract: contract}, LagrangeFilterer: LagrangeFilterer{contract: contract}}, nil
}

// NewLagrangeCaller creates a new read-only instance of Lagrange, bound to a specific deployed contract.
func NewLagrangeCaller(address common.Address, caller bind.ContractCaller) (*LagrangeCaller, error) {
	contract, err := bindLagrange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeCaller{contract: contract}, nil
}

// NewLagrangeTransactor creates a new write-only instance of Lagrange, bound to a specific deployed contract.
func NewLagrangeTransactor(address common.Address, transactor bind.ContractTransactor) (*LagrangeTransactor, error) {
	contract, err := bindLagrange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeTransactor{contract: contract}, nil
}

// NewLagrangeFilterer creates a new log filterer instance of Lagrange, bound to a specific deployed contract.
func NewLagrangeFilterer(address common.Address, filterer bind.ContractFilterer) (*LagrangeFilterer, error) {
	contract, err := bindLagrange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagrangeFilterer{contract: contract}, nil
}

// bindLagrange binds a generic wrapper to an already deployed contract.
func bindLagrange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LagrangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lagrange *LagrangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lagrange.Contract.LagrangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lagrange *LagrangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.Contract.LagrangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lagrange *LagrangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lagrange.Contract.LagrangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lagrange *LagrangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lagrange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lagrange *LagrangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lagrange *LagrangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lagrange.Contract.contract.Transact(opts, method, params...)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address operator) view returns(bool)
func (_Lagrange *LagrangeCaller) IsFrozen(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "isFrozen", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address operator) view returns(bool)
func (_Lagrange *LagrangeSession) IsFrozen(operator common.Address) (bool, error) {
	return _Lagrange.Contract.IsFrozen(&_Lagrange.CallOpts, operator)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address operator) view returns(bool)
func (_Lagrange *LagrangeCallerSession) IsFrozen(operator common.Address) (bool, error) {
	return _Lagrange.Contract.IsFrozen(&_Lagrange.CallOpts, operator)
}

// LatestServeUntilBlock is a free data retrieval call binding the contract method 0x758f8dba.
//
// Solidity: function latestServeUntilBlock() view returns(uint32)
func (_Lagrange *LagrangeCaller) LatestServeUntilBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "latestServeUntilBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestServeUntilBlock is a free data retrieval call binding the contract method 0x758f8dba.
//
// Solidity: function latestServeUntilBlock() view returns(uint32)
func (_Lagrange *LagrangeSession) LatestServeUntilBlock() (uint32, error) {
	return _Lagrange.Contract.LatestServeUntilBlock(&_Lagrange.CallOpts)
}

// LatestServeUntilBlock is a free data retrieval call binding the contract method 0x758f8dba.
//
// Solidity: function latestServeUntilBlock() view returns(uint32)
func (_Lagrange *LagrangeCallerSession) LatestServeUntilBlock() (uint32, error) {
	return _Lagrange.Contract.LatestServeUntilBlock(&_Lagrange.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, uint32 serveUntilBlock, bool slashed)
func (_Lagrange *LagrangeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount          *big.Int
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		Amount          *big.Int
		ServeUntilBlock uint32
		Slashed         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ServeUntilBlock = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Slashed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, uint32 serveUntilBlock, bool slashed)
func (_Lagrange *LagrangeSession) Operators(arg0 common.Address) (struct {
	Amount          *big.Int
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	return _Lagrange.Contract.Operators(&_Lagrange.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, uint32 serveUntilBlock, bool slashed)
func (_Lagrange *LagrangeCallerSession) Operators(arg0 common.Address) (struct {
	Amount          *big.Int
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	return _Lagrange.Contract.Operators(&_Lagrange.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrange *LagrangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrange *LagrangeSession) Owner() (common.Address, error) {
	return _Lagrange.Contract.Owner(&_Lagrange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lagrange *LagrangeCallerSession) Owner() (common.Address, error) {
	return _Lagrange.Contract.Owner(&_Lagrange.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_Lagrange *LagrangeCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_Lagrange *LagrangeSession) Slasher() (common.Address, error) {
	return _Lagrange.Contract.Slasher(&_Lagrange.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_Lagrange *LagrangeCallerSession) Slasher() (common.Address, error) {
	return _Lagrange.Contract.Slasher(&_Lagrange.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_Lagrange *LagrangeCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_Lagrange *LagrangeSession) TaskNumber() (uint32, error) {
	return _Lagrange.Contract.TaskNumber(&_Lagrange.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_Lagrange *LagrangeCallerSession) TaskNumber() (uint32, error) {
	return _Lagrange.Contract.TaskNumber(&_Lagrange.CallOpts)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactor) RecordLastStakeUpdateAndRevokeSlashingAbility(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "recordLastStakeUpdateAndRevokeSlashingAbility", operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeSession) RecordLastStakeUpdateAndRevokeSlashingAbility(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.RecordLastStakeUpdateAndRevokeSlashingAbility(&_Lagrange.TransactOpts, operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactorSession) RecordLastStakeUpdateAndRevokeSlashingAbility(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.RecordLastStakeUpdateAndRevokeSlashingAbility(&_Lagrange.TransactOpts, operator, serveUntilBlock)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) returns()
func (_Lagrange *LagrangeTransactor) RecordStakeUpdate(opts *bind.TransactOpts, operator common.Address, updateBlock uint32, serveUntilBlock uint32, prevElement *big.Int) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "recordStakeUpdate", operator, updateBlock, serveUntilBlock, prevElement)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) returns()
func (_Lagrange *LagrangeSession) RecordStakeUpdate(operator common.Address, updateBlock uint32, serveUntilBlock uint32, prevElement *big.Int) (*types.Transaction, error) {
	return _Lagrange.Contract.RecordStakeUpdate(&_Lagrange.TransactOpts, operator, updateBlock, serveUntilBlock, prevElement)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) returns()
func (_Lagrange *LagrangeTransactorSession) RecordStakeUpdate(operator common.Address, updateBlock uint32, serveUntilBlock uint32, prevElement *big.Int) (*types.Transaction, error) {
	return _Lagrange.Contract.RecordStakeUpdate(&_Lagrange.TransactOpts, operator, updateBlock, serveUntilBlock, prevElement)
}

// Register is a paid mutator transaction binding the contract method 0x130d7906.
//
// Solidity: function register(uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactor) Register(opts *bind.TransactOpts, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "register", serveUntilBlock)
}

// Register is a paid mutator transaction binding the contract method 0x130d7906.
//
// Solidity: function register(uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeSession) Register(serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, serveUntilBlock)
}

// Register is a paid mutator transaction binding the contract method 0x130d7906.
//
// Solidity: function register(uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactorSession) Register(serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, serveUntilBlock)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrange *LagrangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrange *LagrangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lagrange.Contract.RenounceOwnership(&_Lagrange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lagrange *LagrangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lagrange.Contract.RenounceOwnership(&_Lagrange.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrange *LagrangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrange *LagrangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.TransferOwnership(&_Lagrange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lagrange *LagrangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.TransferOwnership(&_Lagrange.TransactOpts, newOwner)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0x4495c7e9.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32) evidence) returns()
func (_Lagrange *LagrangeTransactor) UploadEvidence(opts *bind.TransactOpts, evidence LagrangeServiceEvidence) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "uploadEvidence", evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0x4495c7e9.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32) evidence) returns()
func (_Lagrange *LagrangeSession) UploadEvidence(evidence LagrangeServiceEvidence) (*types.Transaction, error) {
	return _Lagrange.Contract.UploadEvidence(&_Lagrange.TransactOpts, evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0x4495c7e9.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32) evidence) returns()
func (_Lagrange *LagrangeTransactorSession) UploadEvidence(evidence LagrangeServiceEvidence) (*types.Transaction, error) {
	return _Lagrange.Contract.UploadEvidence(&_Lagrange.TransactOpts, evidence)
}

// LagrangeOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the Lagrange contract.
type LagrangeOperatorRegisteredIterator struct {
	Event *LagrangeOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *LagrangeOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOperatorRegistered)
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
		it.Event = new(LagrangeOperatorRegistered)
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
func (it *LagrangeOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOperatorRegistered represents a OperatorRegistered event raised by the Lagrange contract.
type LagrangeOperatorRegistered struct {
	Operator        common.Address
	ServeUntilBlock uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address operator, uint32 serveUntilBlock)
func (_Lagrange *LagrangeFilterer) FilterOperatorRegistered(opts *bind.FilterOpts) (*LagrangeOperatorRegisteredIterator, error) {

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return &LagrangeOperatorRegisteredIterator{contract: _Lagrange.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address operator, uint32 serveUntilBlock)
func (_Lagrange *LagrangeFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *LagrangeOperatorRegistered) (event.Subscription, error) {

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOperatorRegistered)
				if err := _Lagrange.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address operator, uint32 serveUntilBlock)
func (_Lagrange *LagrangeFilterer) ParseOperatorRegistered(log types.Log) (*LagrangeOperatorRegistered, error) {
	event := new(LagrangeOperatorRegistered)
	if err := _Lagrange.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the Lagrange contract.
type LagrangeOperatorSlashedIterator struct {
	Event *LagrangeOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *LagrangeOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOperatorSlashed)
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
		it.Event = new(LagrangeOperatorSlashed)
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
func (it *LagrangeOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOperatorSlashed represents a OperatorSlashed event raised by the Lagrange contract.
type LagrangeOperatorSlashed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a3068.
//
// Solidity: event OperatorSlashed(address operator)
func (_Lagrange *LagrangeFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*LagrangeOperatorSlashedIterator, error) {

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &LagrangeOperatorSlashedIterator{contract: _Lagrange.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a3068.
//
// Solidity: event OperatorSlashed(address operator)
func (_Lagrange *LagrangeFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *LagrangeOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOperatorSlashed)
				if err := _Lagrange.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a3068.
//
// Solidity: event OperatorSlashed(address operator)
func (_Lagrange *LagrangeFilterer) ParseOperatorSlashed(log types.Log) (*LagrangeOperatorSlashed, error) {
	event := new(LagrangeOperatorSlashed)
	if err := _Lagrange.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lagrange contract.
type LagrangeOwnershipTransferredIterator struct {
	Event *LagrangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LagrangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeOwnershipTransferred)
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
		it.Event = new(LagrangeOwnershipTransferred)
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
func (it *LagrangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeOwnershipTransferred represents a OwnershipTransferred event raised by the Lagrange contract.
type LagrangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrange *LagrangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LagrangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeOwnershipTransferredIterator{contract: _Lagrange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lagrange *LagrangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LagrangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeOwnershipTransferred)
				if err := _Lagrange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lagrange *LagrangeFilterer) ParseOwnershipTransferred(log types.Log) (*LagrangeOwnershipTransferred, error) {
	event := new(LagrangeOwnershipTransferred)
	if err := _Lagrange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeUploadEvidenceIterator is returned from FilterUploadEvidence and is used to iterate over the raw logs and unpacked data for UploadEvidence events raised by the Lagrange contract.
type LagrangeUploadEvidenceIterator struct {
	Event *LagrangeUploadEvidence // Event containing the contract specifics and raw log

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
func (it *LagrangeUploadEvidenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeUploadEvidence)
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
		it.Event = new(LagrangeUploadEvidence)
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
func (it *LagrangeUploadEvidenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeUploadEvidenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeUploadEvidence represents a UploadEvidence event raised by the Lagrange contract.
type LagrangeUploadEvidence struct {
	Operator             common.Address
	BlockHash            [32]byte
	CurrentCommitteeRoot [32]byte
	NextCommitteeRoot    [32]byte
	BlockNumber          *big.Int
	EpochNumber          *big.Int
	BlockSignature       []byte
	CommitSignature      []byte
	ChainID              uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUploadEvidence is a free log retrieval operation binding the contract event 0xa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f.
//
// Solidity: event UploadEvidence(address operator, bytes32 blockHash, bytes32 currentCommitteeRoot, bytes32 nextCommitteeRoot, uint256 blockNumber, uint256 epochNumber, bytes blockSignature, bytes commitSignature, uint32 chainID)
func (_Lagrange *LagrangeFilterer) FilterUploadEvidence(opts *bind.FilterOpts) (*LagrangeUploadEvidenceIterator, error) {

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "UploadEvidence")
	if err != nil {
		return nil, err
	}
	return &LagrangeUploadEvidenceIterator{contract: _Lagrange.contract, event: "UploadEvidence", logs: logs, sub: sub}, nil
}

// WatchUploadEvidence is a free log subscription operation binding the contract event 0xa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f.
//
// Solidity: event UploadEvidence(address operator, bytes32 blockHash, bytes32 currentCommitteeRoot, bytes32 nextCommitteeRoot, uint256 blockNumber, uint256 epochNumber, bytes blockSignature, bytes commitSignature, uint32 chainID)
func (_Lagrange *LagrangeFilterer) WatchUploadEvidence(opts *bind.WatchOpts, sink chan<- *LagrangeUploadEvidence) (event.Subscription, error) {

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "UploadEvidence")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeUploadEvidence)
				if err := _Lagrange.contract.UnpackLog(event, "UploadEvidence", log); err != nil {
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

// ParseUploadEvidence is a log parse operation binding the contract event 0xa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f.
//
// Solidity: event UploadEvidence(address operator, bytes32 blockHash, bytes32 currentCommitteeRoot, bytes32 nextCommitteeRoot, uint256 blockNumber, uint256 epochNumber, bytes blockSignature, bytes commitSignature, uint32 chainID)
func (_Lagrange *LagrangeFilterer) ParseUploadEvidence(log types.Log) (*LagrangeUploadEvidence, error) {
	event := new(LagrangeUploadEvidence)
	if err := _Lagrange.contract.UnpackLog(event, "UploadEvidence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
