// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voteweigher

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

// IVoteWeigherTokenMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IVoteWeigherTokenMultiplier struct {
	Token      common.Address
	Multiplier *big.Int
}

// VoteweigherMetaData contains all meta data concerning the Voteweigher contract.
var VoteweigherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeManager\",\"type\":\"address\",\"internalType\":\"contractIStakeManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WEIGHTING_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addQuorumMultiplier\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"multipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIVoteWeigher.TokenMultiplier[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTokenList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenListForQuorumNumbers\",\"inputs\":[{\"name\":\"quorumNumbers_\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumMultipliers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumbers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeQuorumMultiplier\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQuorumMultiplier\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiplier\",\"type\":\"tuple\",\"internalType\":\"structIVoteWeigher.TokenMultiplier\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"weightOfOperator\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumAdded\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"multipliers\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIVoteWeigher.TokenMultiplier[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumRemoved\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"multiplier\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIVoteWeigher.TokenMultiplier\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516117d43803806117d483398101604081905261002f91610108565b610037610048565b6001600160a01b0316608052610138565b600054610100900460ff16156100b45760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015610106576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561011a57600080fd5b81516001600160a01b038116811461013157600080fd5b9392505050565b60805161167a61015a600039600081816101cd0152610c19015261167a6000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063715018a61161008c578063a5ba381e11610066578063a5ba381e14610218578063c4d66de81461022b578063da4098581461023e578063f2fde38b1461026957600080fd5b8063715018a6146101c05780637542ff95146101c85780638da5cb5b1461020757600080fd5b806357d2ebf0116100c857806357d2ebf0146101585780635ce8f0671461016b5780635e5a67751461017e578063658096611461019b57600080fd5b80630855d567146100ef5780632202419a1461012e578063273cbaa014610143575b600080fd5b6101026100fd366004611224565b61027c565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390f35b61014161013c36600461124e565b6102c5565b005b61014b610484565b6040516101259190611269565b6101416101663660046112b6565b610502565b61014b61017936600461133c565b61075c565b61018d670de0b6b3a764000081565b604051908152602001610125565b6101ae6101a93660046113b1565b6107a3565b60405160ff9091168152602001610125565b6101416107d7565b6101ef7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610125565b6033546001600160a01b03166101ef565b6101416102263660046113ca565b6107eb565b610141610239366004611427565b610a6f565b61025161024c36600461144b565b610b82565b6040516001600160601b039091168152602001610125565b610141610277366004611427565b610d39565b6065602052816000526040600020818154811061029857600080fd5b6000918252602090912001546001600160a01b0381169250600160a01b90046001600160601b0316905082565b6102cd610db2565b60ff8116600090815260656020526040812054900361032a5760405162461bcd60e51b8152602060048201526014602482015273145d5bdc9d5b48191bd95cdb89dd08195e1a5cdd60621b60448201526064015b60405180910390fd5b60665460005b81811015610436578260ff166066828154811061034f5761034f611482565b60009182526020918290209181049091015460ff601f9092166101000a9004160361042e5760666103816001846114ae565b8154811061039157610391611482565b90600052602060002090602091828204019190069054906101000a900460ff16606682815481106103c4576103c4611482565b90600052602060002090602091828204019190066101000a81548160ff021916908360ff16021790555060668054806103ff576103ff6114c1565b60019003818190600052602060002090602091828204019190066101000a81549060ff02191690559055610436565b600101610330565b5060ff82166000908152606560205260408120610452916111dc565b60405160ff8316907f5ff139664dd1ec0405d47f20c0cbc4f03800d5ae401624fb58ae8233ef477e9190600090a25050565b60606104fd60668054806020026020016040519081016040528092919081815260200182805480156104f357602002820191906000526020600020906000905b825461010083900a900460ff168152602060019283018181049485019490930390920291018084116104c45790505b5050505050610e0c565b905090565b61050a610db2565b600081900361055b5760405162461bcd60e51b815260206004820152601960248201527f456d707479206c697374206f66206d756c7469706c69657273000000000000006044820152606401610321565b60ff8316600090815260656020526040902054156105b35760405162461bcd60e51b815260206004820152601560248201527451756f72756d20616c72656164792065786973747360581b6044820152606401610321565b60005b818110156106cb5760ff8416600090815260656020908152604080832080548251818502810185019093528083526106719492939192909184015b8282101561064057600084815260209081902060408051808201909152908401546001600160a01b0381168252600160a01b90046001600160601b0316818301528252600190920191016105f1565b5050505084848481811061065657610656611482565b61066c9260206040909202019081019150611427565b6110f5565b60ff8416600090815260656020526040902083838381811061069557610695611482565b835460018101855560009485526020909420604090910292909201929190910190506106c182826114ec565b50506001016105b6565b5060668054600181018255600091909152602081047f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e9435401805460ff808716601f9094166101000a8481029102199091161790556040517e91b236fa2c2c92dc35476c79f3942a160d346184707a42ed60b8f487de2a589061074f9085908590611566565b60405180910390a2505050565b606061079a838380806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610e0c92505050565b90505b92915050565b606681815481106107b357600080fd5b9060005260206000209060209182820401919006915054906101000a900460ff1681565b6107df610db2565b6107e9600061118a565b565b6107f3610db2565b60ff831660009081526065602052604090205482111561084b5760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b6044820152606401610321565b60ff831660009081526065602052604090205482900361092e5760ff8316600090815260656020908152604080832080548251818502810185019093528083526108f99492939192909184015b828210156108e757600084815260209081902060408051808201909152908401546001600160a01b0381168252600160a01b90046001600160601b031681830152825260019092019101610898565b5061066c925050506020840184611427565b60ff8316600090815260656020908152604082208054600181018255908352912082910161092782826114ec565b5050610a3a565b60ff8316600090815260656020526040812054905b818110156109f8578381146109f05761095f6020840184611427565b6001600160a01b0316606560008760ff1660ff168152602001908152602001600020828154811061099257610992611482565b6000918252602090912001546001600160a01b0316036109f05760405162461bcd60e51b81526020600482015260196024820152784d756c7469706c69657220616c72656164792065786973747360381b6044820152606401610321565b600101610943565b5060ff84166000908152606560205260409020805483919085908110610a2057610a20611482565b906000526020600020018181610a3691906114ec565b5050505b8260ff167fee9c9cc7dfc433e84d31f1d7a340b6aa0d46bd17055346a78997f3354a4250d4838360405161074f92919061159c565b600054610100900460ff1615808015610a8f5750600054600160ff909116105b80610aa95750303b158015610aa9575060005460ff166001145b610b0c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610321565b6000805460ff191660011790558015610b2f576000805461ff0019166101001790555b610b388261118a565b8015610b7e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b60ff82166000908152606560209081526040808320805482518185028101850190935280835284938493929190849084015b82821015610c0357600084815260209081902060408051808201909152908401546001600160a01b0381168252600160a01b90046001600160601b031681830152825260019092019101610bb4565b50505050905060005b8151811015610d1d5760007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663778e55f387858581518110610c5957610c59611482565b6020908102919091010151516040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401602060405180830381865afa158015610cb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd491906115b0565b9050828281518110610ce857610ce8611482565b6020026020010151602001516001600160601b031681610d0891906115c9565b610d1290856115e0565b935050600101610c0c565b50610d30670de0b6b3a7640000836115f3565b95945050505050565b610d41610db2565b6001600160a01b038116610da65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610321565b610daf8161118a565b50565b6033546001600160a01b031633146107e95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610321565b80516060906000805b82811015610e685760656000868381518110610e3357610e33611482565b602002602001015160ff1660ff1681526020019081526020016000208054905082610e5e91906115e0565b9150600101610e15565b5060008167ffffffffffffffff811115610e8457610e84611615565b604051908082528060200260200182016040528015610ead578160200160208202803683370190505b5090506000805b8481101561104b5760005b60656000898481518110610ed557610ed5611482565b602002602001015160ff1660ff16815260200190815260200160002080549050811015611042576000805b84811015610f9e57606560008b8681518110610f1e57610f1e611482565b602002602001015160ff1660ff1681526020019081526020016000208381548110610f4b57610f4b611482565b60009182526020909120015486516001600160a01b0390911690879083908110610f7757610f77611482565b60200260200101516001600160a01b031603610f965760019150610f9e565b600101610f00565b508061103957606560008a8581518110610fba57610fba611482565b602002602001015160ff1660ff1681526020019081526020016000208281548110610fe757610fe7611482565b60009182526020909120015485516001600160a01b039091169086908690811061101357611013611482565b6001600160a01b0390921660209283029190910190910152836110358161162b565b9450505b50600101610ebf565b50600101610eb4565b5060008167ffffffffffffffff81111561106757611067611615565b604051908082528060200260200182016040528015611090578160200160208202803683370190505b50905060005b828110156110ea578381815181106110b0576110b0611482565b60200260200101518282815181106110ca576110ca611482565b6001600160a01b0390921660209283029190910190910152600101611096565b509695505050505050565b815160005b8181101561118457826001600160a01b031684828151811061111e5761111e611482565b6020026020010151600001516001600160a01b03160361117c5760405162461bcd60e51b81526020600482015260196024820152784d756c7469706c69657220616c72656164792065786973747360381b6044820152606401610321565b6001016110fa565b50505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b5080546000825590600052602060002090810190610daf91905b8082111561120a57600081556001016111f6565b5090565b803560ff8116811461121f57600080fd5b919050565b6000806040838503121561123757600080fd5b6112408361120e565b946020939093013593505050565b60006020828403121561126057600080fd5b61079a8261120e565b6020808252825182820181905260009190848201906040850190845b818110156112aa5783516001600160a01b031683529284019291840191600101611285565b50909695505050505050565b6000806000604084860312156112cb57600080fd5b6112d48461120e565b9250602084013567ffffffffffffffff808211156112f157600080fd5b818601915086601f83011261130557600080fd5b81358181111561131457600080fd5b8760208260061b850101111561132957600080fd5b6020830194508093505050509250925092565b6000806020838503121561134f57600080fd5b823567ffffffffffffffff8082111561136757600080fd5b818501915085601f83011261137b57600080fd5b81358181111561138a57600080fd5b8660208260051b850101111561139f57600080fd5b60209290920196919550909350505050565b6000602082840312156113c357600080fd5b5035919050565b600080600083850360808112156113e057600080fd5b6113e98561120e565b9350602085013592506040603f198201121561140457600080fd5b506040840190509250925092565b6001600160a01b0381168114610daf57600080fd5b60006020828403121561143957600080fd5b813561144481611412565b9392505050565b6000806040838503121561145e57600080fd5b6114678361120e565b9150602083013561147781611412565b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561079d5761079d611498565b634e487b7160e01b600052603160045260246000fd5b6001600160601b0381168114610daf57600080fd5b81356114f781611412565b81546001600160a01b03199081166001600160a01b039290921691821783556020840135611524816114d7565b60a01b1617905550565b803561153981611412565b6001600160a01b031682526020810135611552816114d7565b6001600160601b0381166020840152505050565b602080825281018290526000604080830185835b868110156112aa5761158c838361152e565b918301919083019060010161157a565b82815260608101611444602083018461152e565b6000602082840312156115c257600080fd5b5051919050565b808202811582820484141761079d5761079d611498565b8082018082111561079d5761079d611498565b60008261161057634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052604160045260246000fd5b60006001820161163d5761163d611498565b506001019056fea26469706673582212205fc56164d2baff73a17b8f5e6d96ef06214c0c4edf6bfe40cffe14917848b27464736f6c63430008190033",
}

// VoteweigherABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteweigherMetaData.ABI instead.
var VoteweigherABI = VoteweigherMetaData.ABI

// VoteweigherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteweigherMetaData.Bin instead.
var VoteweigherBin = VoteweigherMetaData.Bin

// DeployVoteweigher deploys a new Ethereum contract, binding an instance of Voteweigher to it.
func DeployVoteweigher(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeManager common.Address) (common.Address, *types.Transaction, *Voteweigher, error) {
	parsed, err := VoteweigherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteweigherBin), backend, _stakeManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voteweigher{VoteweigherCaller: VoteweigherCaller{contract: contract}, VoteweigherTransactor: VoteweigherTransactor{contract: contract}, VoteweigherFilterer: VoteweigherFilterer{contract: contract}}, nil
}

// Voteweigher is an auto generated Go binding around an Ethereum contract.
type Voteweigher struct {
	VoteweigherCaller     // Read-only binding to the contract
	VoteweigherTransactor // Write-only binding to the contract
	VoteweigherFilterer   // Log filterer for contract events
}

// VoteweigherCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteweigherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteweigherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteweigherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteweigherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteweigherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteweigherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteweigherSession struct {
	Contract     *Voteweigher      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteweigherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteweigherCallerSession struct {
	Contract *VoteweigherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VoteweigherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteweigherTransactorSession struct {
	Contract     *VoteweigherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VoteweigherRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteweigherRaw struct {
	Contract *Voteweigher // Generic contract binding to access the raw methods on
}

// VoteweigherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteweigherCallerRaw struct {
	Contract *VoteweigherCaller // Generic read-only contract binding to access the raw methods on
}

// VoteweigherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteweigherTransactorRaw struct {
	Contract *VoteweigherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoteweigher creates a new instance of Voteweigher, bound to a specific deployed contract.
func NewVoteweigher(address common.Address, backend bind.ContractBackend) (*Voteweigher, error) {
	contract, err := bindVoteweigher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voteweigher{VoteweigherCaller: VoteweigherCaller{contract: contract}, VoteweigherTransactor: VoteweigherTransactor{contract: contract}, VoteweigherFilterer: VoteweigherFilterer{contract: contract}}, nil
}

// NewVoteweigherCaller creates a new read-only instance of Voteweigher, bound to a specific deployed contract.
func NewVoteweigherCaller(address common.Address, caller bind.ContractCaller) (*VoteweigherCaller, error) {
	contract, err := bindVoteweigher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteweigherCaller{contract: contract}, nil
}

// NewVoteweigherTransactor creates a new write-only instance of Voteweigher, bound to a specific deployed contract.
func NewVoteweigherTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteweigherTransactor, error) {
	contract, err := bindVoteweigher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteweigherTransactor{contract: contract}, nil
}

// NewVoteweigherFilterer creates a new log filterer instance of Voteweigher, bound to a specific deployed contract.
func NewVoteweigherFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteweigherFilterer, error) {
	contract, err := bindVoteweigher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteweigherFilterer{contract: contract}, nil
}

// bindVoteweigher binds a generic wrapper to an already deployed contract.
func bindVoteweigher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteweigherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voteweigher *VoteweigherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voteweigher.Contract.VoteweigherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voteweigher *VoteweigherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voteweigher.Contract.VoteweigherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voteweigher *VoteweigherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voteweigher.Contract.VoteweigherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voteweigher *VoteweigherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voteweigher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voteweigher *VoteweigherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voteweigher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voteweigher *VoteweigherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voteweigher.Contract.contract.Transact(opts, method, params...)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_Voteweigher *VoteweigherCaller) WEIGHTINGDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "WEIGHTING_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_Voteweigher *VoteweigherSession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _Voteweigher.Contract.WEIGHTINGDIVISOR(&_Voteweigher.CallOpts)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_Voteweigher *VoteweigherCallerSession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _Voteweigher.Contract.WEIGHTINGDIVISOR(&_Voteweigher.CallOpts)
}

// GetTokenList is a free data retrieval call binding the contract method 0x273cbaa0.
//
// Solidity: function getTokenList() view returns(address[])
func (_Voteweigher *VoteweigherCaller) GetTokenList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "getTokenList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenList is a free data retrieval call binding the contract method 0x273cbaa0.
//
// Solidity: function getTokenList() view returns(address[])
func (_Voteweigher *VoteweigherSession) GetTokenList() ([]common.Address, error) {
	return _Voteweigher.Contract.GetTokenList(&_Voteweigher.CallOpts)
}

// GetTokenList is a free data retrieval call binding the contract method 0x273cbaa0.
//
// Solidity: function getTokenList() view returns(address[])
func (_Voteweigher *VoteweigherCallerSession) GetTokenList() ([]common.Address, error) {
	return _Voteweigher.Contract.GetTokenList(&_Voteweigher.CallOpts)
}

// GetTokenListForQuorumNumbers is a free data retrieval call binding the contract method 0x5ce8f067.
//
// Solidity: function getTokenListForQuorumNumbers(uint8[] quorumNumbers_) view returns(address[])
func (_Voteweigher *VoteweigherCaller) GetTokenListForQuorumNumbers(opts *bind.CallOpts, quorumNumbers_ []uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "getTokenListForQuorumNumbers", quorumNumbers_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenListForQuorumNumbers is a free data retrieval call binding the contract method 0x5ce8f067.
//
// Solidity: function getTokenListForQuorumNumbers(uint8[] quorumNumbers_) view returns(address[])
func (_Voteweigher *VoteweigherSession) GetTokenListForQuorumNumbers(quorumNumbers_ []uint8) ([]common.Address, error) {
	return _Voteweigher.Contract.GetTokenListForQuorumNumbers(&_Voteweigher.CallOpts, quorumNumbers_)
}

// GetTokenListForQuorumNumbers is a free data retrieval call binding the contract method 0x5ce8f067.
//
// Solidity: function getTokenListForQuorumNumbers(uint8[] quorumNumbers_) view returns(address[])
func (_Voteweigher *VoteweigherCallerSession) GetTokenListForQuorumNumbers(quorumNumbers_ []uint8) ([]common.Address, error) {
	return _Voteweigher.Contract.GetTokenListForQuorumNumbers(&_Voteweigher.CallOpts, quorumNumbers_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voteweigher *VoteweigherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voteweigher *VoteweigherSession) Owner() (common.Address, error) {
	return _Voteweigher.Contract.Owner(&_Voteweigher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voteweigher *VoteweigherCallerSession) Owner() (common.Address, error) {
	return _Voteweigher.Contract.Owner(&_Voteweigher.CallOpts)
}

// QuorumMultipliers is a free data retrieval call binding the contract method 0x0855d567.
//
// Solidity: function quorumMultipliers(uint8 , uint256 ) view returns(address token, uint96 multiplier)
func (_Voteweigher *VoteweigherCaller) QuorumMultipliers(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
	Token      common.Address
	Multiplier *big.Int
}, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "quorumMultipliers", arg0, arg1)

	outstruct := new(struct {
		Token      common.Address
		Multiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Multiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuorumMultipliers is a free data retrieval call binding the contract method 0x0855d567.
//
// Solidity: function quorumMultipliers(uint8 , uint256 ) view returns(address token, uint96 multiplier)
func (_Voteweigher *VoteweigherSession) QuorumMultipliers(arg0 uint8, arg1 *big.Int) (struct {
	Token      common.Address
	Multiplier *big.Int
}, error) {
	return _Voteweigher.Contract.QuorumMultipliers(&_Voteweigher.CallOpts, arg0, arg1)
}

// QuorumMultipliers is a free data retrieval call binding the contract method 0x0855d567.
//
// Solidity: function quorumMultipliers(uint8 , uint256 ) view returns(address token, uint96 multiplier)
func (_Voteweigher *VoteweigherCallerSession) QuorumMultipliers(arg0 uint8, arg1 *big.Int) (struct {
	Token      common.Address
	Multiplier *big.Int
}, error) {
	return _Voteweigher.Contract.QuorumMultipliers(&_Voteweigher.CallOpts, arg0, arg1)
}

// QuorumNumbers is a free data retrieval call binding the contract method 0x65809661.
//
// Solidity: function quorumNumbers(uint256 ) view returns(uint8)
func (_Voteweigher *VoteweigherCaller) QuorumNumbers(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "quorumNumbers", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumNumbers is a free data retrieval call binding the contract method 0x65809661.
//
// Solidity: function quorumNumbers(uint256 ) view returns(uint8)
func (_Voteweigher *VoteweigherSession) QuorumNumbers(arg0 *big.Int) (uint8, error) {
	return _Voteweigher.Contract.QuorumNumbers(&_Voteweigher.CallOpts, arg0)
}

// QuorumNumbers is a free data retrieval call binding the contract method 0x65809661.
//
// Solidity: function quorumNumbers(uint256 ) view returns(uint8)
func (_Voteweigher *VoteweigherCallerSession) QuorumNumbers(arg0 *big.Int) (uint8, error) {
	return _Voteweigher.Contract.QuorumNumbers(&_Voteweigher.CallOpts, arg0)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Voteweigher *VoteweigherCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Voteweigher *VoteweigherSession) StakeManager() (common.Address, error) {
	return _Voteweigher.Contract.StakeManager(&_Voteweigher.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Voteweigher *VoteweigherCallerSession) StakeManager() (common.Address, error) {
	return _Voteweigher.Contract.StakeManager(&_Voteweigher.CallOpts)
}

// WeightOfOperator is a free data retrieval call binding the contract method 0xda409858.
//
// Solidity: function weightOfOperator(uint8 quorumNumber, address operator) view returns(uint96)
func (_Voteweigher *VoteweigherCaller) WeightOfOperator(opts *bind.CallOpts, quorumNumber uint8, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voteweigher.contract.Call(opts, &out, "weightOfOperator", quorumNumber, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightOfOperator is a free data retrieval call binding the contract method 0xda409858.
//
// Solidity: function weightOfOperator(uint8 quorumNumber, address operator) view returns(uint96)
func (_Voteweigher *VoteweigherSession) WeightOfOperator(quorumNumber uint8, operator common.Address) (*big.Int, error) {
	return _Voteweigher.Contract.WeightOfOperator(&_Voteweigher.CallOpts, quorumNumber, operator)
}

// WeightOfOperator is a free data retrieval call binding the contract method 0xda409858.
//
// Solidity: function weightOfOperator(uint8 quorumNumber, address operator) view returns(uint96)
func (_Voteweigher *VoteweigherCallerSession) WeightOfOperator(quorumNumber uint8, operator common.Address) (*big.Int, error) {
	return _Voteweigher.Contract.WeightOfOperator(&_Voteweigher.CallOpts, quorumNumber, operator)
}

// AddQuorumMultiplier is a paid mutator transaction binding the contract method 0x57d2ebf0.
//
// Solidity: function addQuorumMultiplier(uint8 quorumNumber, (address,uint96)[] multipliers) returns()
func (_Voteweigher *VoteweigherTransactor) AddQuorumMultiplier(opts *bind.TransactOpts, quorumNumber uint8, multipliers []IVoteWeigherTokenMultiplier) (*types.Transaction, error) {
	return _Voteweigher.contract.Transact(opts, "addQuorumMultiplier", quorumNumber, multipliers)
}

// AddQuorumMultiplier is a paid mutator transaction binding the contract method 0x57d2ebf0.
//
// Solidity: function addQuorumMultiplier(uint8 quorumNumber, (address,uint96)[] multipliers) returns()
func (_Voteweigher *VoteweigherSession) AddQuorumMultiplier(quorumNumber uint8, multipliers []IVoteWeigherTokenMultiplier) (*types.Transaction, error) {
	return _Voteweigher.Contract.AddQuorumMultiplier(&_Voteweigher.TransactOpts, quorumNumber, multipliers)
}

// AddQuorumMultiplier is a paid mutator transaction binding the contract method 0x57d2ebf0.
//
// Solidity: function addQuorumMultiplier(uint8 quorumNumber, (address,uint96)[] multipliers) returns()
func (_Voteweigher *VoteweigherTransactorSession) AddQuorumMultiplier(quorumNumber uint8, multipliers []IVoteWeigherTokenMultiplier) (*types.Transaction, error) {
	return _Voteweigher.Contract.AddQuorumMultiplier(&_Voteweigher.TransactOpts, quorumNumber, multipliers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Voteweigher *VoteweigherTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Voteweigher.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Voteweigher *VoteweigherSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Voteweigher.Contract.Initialize(&_Voteweigher.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Voteweigher *VoteweigherTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Voteweigher.Contract.Initialize(&_Voteweigher.TransactOpts, initialOwner)
}

// RemoveQuorumMultiplier is a paid mutator transaction binding the contract method 0x2202419a.
//
// Solidity: function removeQuorumMultiplier(uint8 quorumNumber) returns()
func (_Voteweigher *VoteweigherTransactor) RemoveQuorumMultiplier(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _Voteweigher.contract.Transact(opts, "removeQuorumMultiplier", quorumNumber)
}

// RemoveQuorumMultiplier is a paid mutator transaction binding the contract method 0x2202419a.
//
// Solidity: function removeQuorumMultiplier(uint8 quorumNumber) returns()
func (_Voteweigher *VoteweigherSession) RemoveQuorumMultiplier(quorumNumber uint8) (*types.Transaction, error) {
	return _Voteweigher.Contract.RemoveQuorumMultiplier(&_Voteweigher.TransactOpts, quorumNumber)
}

// RemoveQuorumMultiplier is a paid mutator transaction binding the contract method 0x2202419a.
//
// Solidity: function removeQuorumMultiplier(uint8 quorumNumber) returns()
func (_Voteweigher *VoteweigherTransactorSession) RemoveQuorumMultiplier(quorumNumber uint8) (*types.Transaction, error) {
	return _Voteweigher.Contract.RemoveQuorumMultiplier(&_Voteweigher.TransactOpts, quorumNumber)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Voteweigher *VoteweigherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voteweigher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Voteweigher *VoteweigherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Voteweigher.Contract.RenounceOwnership(&_Voteweigher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Voteweigher *VoteweigherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Voteweigher.Contract.RenounceOwnership(&_Voteweigher.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voteweigher *VoteweigherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voteweigher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voteweigher *VoteweigherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voteweigher.Contract.TransferOwnership(&_Voteweigher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voteweigher *VoteweigherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voteweigher.Contract.TransferOwnership(&_Voteweigher.TransactOpts, newOwner)
}

// UpdateQuorumMultiplier is a paid mutator transaction binding the contract method 0xa5ba381e.
//
// Solidity: function updateQuorumMultiplier(uint8 quorumNumber, uint256 index, (address,uint96) multiplier) returns()
func (_Voteweigher *VoteweigherTransactor) UpdateQuorumMultiplier(opts *bind.TransactOpts, quorumNumber uint8, index *big.Int, multiplier IVoteWeigherTokenMultiplier) (*types.Transaction, error) {
	return _Voteweigher.contract.Transact(opts, "updateQuorumMultiplier", quorumNumber, index, multiplier)
}

// UpdateQuorumMultiplier is a paid mutator transaction binding the contract method 0xa5ba381e.
//
// Solidity: function updateQuorumMultiplier(uint8 quorumNumber, uint256 index, (address,uint96) multiplier) returns()
func (_Voteweigher *VoteweigherSession) UpdateQuorumMultiplier(quorumNumber uint8, index *big.Int, multiplier IVoteWeigherTokenMultiplier) (*types.Transaction, error) {
	return _Voteweigher.Contract.UpdateQuorumMultiplier(&_Voteweigher.TransactOpts, quorumNumber, index, multiplier)
}

// UpdateQuorumMultiplier is a paid mutator transaction binding the contract method 0xa5ba381e.
//
// Solidity: function updateQuorumMultiplier(uint8 quorumNumber, uint256 index, (address,uint96) multiplier) returns()
func (_Voteweigher *VoteweigherTransactorSession) UpdateQuorumMultiplier(quorumNumber uint8, index *big.Int, multiplier IVoteWeigherTokenMultiplier) (*types.Transaction, error) {
	return _Voteweigher.Contract.UpdateQuorumMultiplier(&_Voteweigher.TransactOpts, quorumNumber, index, multiplier)
}

// VoteweigherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Voteweigher contract.
type VoteweigherInitializedIterator struct {
	Event *VoteweigherInitialized // Event containing the contract specifics and raw log

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
func (it *VoteweigherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteweigherInitialized)
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
		it.Event = new(VoteweigherInitialized)
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
func (it *VoteweigherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteweigherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteweigherInitialized represents a Initialized event raised by the Voteweigher contract.
type VoteweigherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voteweigher *VoteweigherFilterer) FilterInitialized(opts *bind.FilterOpts) (*VoteweigherInitializedIterator, error) {

	logs, sub, err := _Voteweigher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VoteweigherInitializedIterator{contract: _Voteweigher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voteweigher *VoteweigherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VoteweigherInitialized) (event.Subscription, error) {

	logs, sub, err := _Voteweigher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteweigherInitialized)
				if err := _Voteweigher.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Voteweigher *VoteweigherFilterer) ParseInitialized(log types.Log) (*VoteweigherInitialized, error) {
	event := new(VoteweigherInitialized)
	if err := _Voteweigher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteweigherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Voteweigher contract.
type VoteweigherOwnershipTransferredIterator struct {
	Event *VoteweigherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VoteweigherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteweigherOwnershipTransferred)
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
		it.Event = new(VoteweigherOwnershipTransferred)
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
func (it *VoteweigherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteweigherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteweigherOwnershipTransferred represents a OwnershipTransferred event raised by the Voteweigher contract.
type VoteweigherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Voteweigher *VoteweigherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VoteweigherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Voteweigher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoteweigherOwnershipTransferredIterator{contract: _Voteweigher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Voteweigher *VoteweigherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VoteweigherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Voteweigher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteweigherOwnershipTransferred)
				if err := _Voteweigher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Voteweigher *VoteweigherFilterer) ParseOwnershipTransferred(log types.Log) (*VoteweigherOwnershipTransferred, error) {
	event := new(VoteweigherOwnershipTransferred)
	if err := _Voteweigher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteweigherQuorumAddedIterator is returned from FilterQuorumAdded and is used to iterate over the raw logs and unpacked data for QuorumAdded events raised by the Voteweigher contract.
type VoteweigherQuorumAddedIterator struct {
	Event *VoteweigherQuorumAdded // Event containing the contract specifics and raw log

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
func (it *VoteweigherQuorumAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteweigherQuorumAdded)
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
		it.Event = new(VoteweigherQuorumAdded)
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
func (it *VoteweigherQuorumAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteweigherQuorumAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteweigherQuorumAdded represents a QuorumAdded event raised by the Voteweigher contract.
type VoteweigherQuorumAdded struct {
	QuorumNumber uint8
	Multipliers  []IVoteWeigherTokenMultiplier
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumAdded is a free log retrieval operation binding the contract event 0x0091b236fa2c2c92dc35476c79f3942a160d346184707a42ed60b8f487de2a58.
//
// Solidity: event QuorumAdded(uint8 indexed quorumNumber, (address,uint96)[] multipliers)
func (_Voteweigher *VoteweigherFilterer) FilterQuorumAdded(opts *bind.FilterOpts, quorumNumber []uint8) (*VoteweigherQuorumAddedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _Voteweigher.contract.FilterLogs(opts, "QuorumAdded", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &VoteweigherQuorumAddedIterator{contract: _Voteweigher.contract, event: "QuorumAdded", logs: logs, sub: sub}, nil
}

// WatchQuorumAdded is a free log subscription operation binding the contract event 0x0091b236fa2c2c92dc35476c79f3942a160d346184707a42ed60b8f487de2a58.
//
// Solidity: event QuorumAdded(uint8 indexed quorumNumber, (address,uint96)[] multipliers)
func (_Voteweigher *VoteweigherFilterer) WatchQuorumAdded(opts *bind.WatchOpts, sink chan<- *VoteweigherQuorumAdded, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _Voteweigher.contract.WatchLogs(opts, "QuorumAdded", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteweigherQuorumAdded)
				if err := _Voteweigher.contract.UnpackLog(event, "QuorumAdded", log); err != nil {
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

// ParseQuorumAdded is a log parse operation binding the contract event 0x0091b236fa2c2c92dc35476c79f3942a160d346184707a42ed60b8f487de2a58.
//
// Solidity: event QuorumAdded(uint8 indexed quorumNumber, (address,uint96)[] multipliers)
func (_Voteweigher *VoteweigherFilterer) ParseQuorumAdded(log types.Log) (*VoteweigherQuorumAdded, error) {
	event := new(VoteweigherQuorumAdded)
	if err := _Voteweigher.contract.UnpackLog(event, "QuorumAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteweigherQuorumRemovedIterator is returned from FilterQuorumRemoved and is used to iterate over the raw logs and unpacked data for QuorumRemoved events raised by the Voteweigher contract.
type VoteweigherQuorumRemovedIterator struct {
	Event *VoteweigherQuorumRemoved // Event containing the contract specifics and raw log

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
func (it *VoteweigherQuorumRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteweigherQuorumRemoved)
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
		it.Event = new(VoteweigherQuorumRemoved)
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
func (it *VoteweigherQuorumRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteweigherQuorumRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteweigherQuorumRemoved represents a QuorumRemoved event raised by the Voteweigher contract.
type VoteweigherQuorumRemoved struct {
	QuorumNumber uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumRemoved is a free log retrieval operation binding the contract event 0x5ff139664dd1ec0405d47f20c0cbc4f03800d5ae401624fb58ae8233ef477e91.
//
// Solidity: event QuorumRemoved(uint8 indexed quorumNumber)
func (_Voteweigher *VoteweigherFilterer) FilterQuorumRemoved(opts *bind.FilterOpts, quorumNumber []uint8) (*VoteweigherQuorumRemovedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _Voteweigher.contract.FilterLogs(opts, "QuorumRemoved", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &VoteweigherQuorumRemovedIterator{contract: _Voteweigher.contract, event: "QuorumRemoved", logs: logs, sub: sub}, nil
}

// WatchQuorumRemoved is a free log subscription operation binding the contract event 0x5ff139664dd1ec0405d47f20c0cbc4f03800d5ae401624fb58ae8233ef477e91.
//
// Solidity: event QuorumRemoved(uint8 indexed quorumNumber)
func (_Voteweigher *VoteweigherFilterer) WatchQuorumRemoved(opts *bind.WatchOpts, sink chan<- *VoteweigherQuorumRemoved, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _Voteweigher.contract.WatchLogs(opts, "QuorumRemoved", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteweigherQuorumRemoved)
				if err := _Voteweigher.contract.UnpackLog(event, "QuorumRemoved", log); err != nil {
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

// ParseQuorumRemoved is a log parse operation binding the contract event 0x5ff139664dd1ec0405d47f20c0cbc4f03800d5ae401624fb58ae8233ef477e91.
//
// Solidity: event QuorumRemoved(uint8 indexed quorumNumber)
func (_Voteweigher *VoteweigherFilterer) ParseQuorumRemoved(log types.Log) (*VoteweigherQuorumRemoved, error) {
	event := new(VoteweigherQuorumRemoved)
	if err := _Voteweigher.contract.UnpackLog(event, "QuorumRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteweigherQuorumUpdatedIterator is returned from FilterQuorumUpdated and is used to iterate over the raw logs and unpacked data for QuorumUpdated events raised by the Voteweigher contract.
type VoteweigherQuorumUpdatedIterator struct {
	Event *VoteweigherQuorumUpdated // Event containing the contract specifics and raw log

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
func (it *VoteweigherQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteweigherQuorumUpdated)
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
		it.Event = new(VoteweigherQuorumUpdated)
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
func (it *VoteweigherQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteweigherQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteweigherQuorumUpdated represents a QuorumUpdated event raised by the Voteweigher contract.
type VoteweigherQuorumUpdated struct {
	QuorumNumber uint8
	Index        *big.Int
	Multiplier   IVoteWeigherTokenMultiplier
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumUpdated is a free log retrieval operation binding the contract event 0xee9c9cc7dfc433e84d31f1d7a340b6aa0d46bd17055346a78997f3354a4250d4.
//
// Solidity: event QuorumUpdated(uint8 indexed quorumNumber, uint256 index, (address,uint96) multiplier)
func (_Voteweigher *VoteweigherFilterer) FilterQuorumUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*VoteweigherQuorumUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _Voteweigher.contract.FilterLogs(opts, "QuorumUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &VoteweigherQuorumUpdatedIterator{contract: _Voteweigher.contract, event: "QuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumUpdated is a free log subscription operation binding the contract event 0xee9c9cc7dfc433e84d31f1d7a340b6aa0d46bd17055346a78997f3354a4250d4.
//
// Solidity: event QuorumUpdated(uint8 indexed quorumNumber, uint256 index, (address,uint96) multiplier)
func (_Voteweigher *VoteweigherFilterer) WatchQuorumUpdated(opts *bind.WatchOpts, sink chan<- *VoteweigherQuorumUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _Voteweigher.contract.WatchLogs(opts, "QuorumUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteweigherQuorumUpdated)
				if err := _Voteweigher.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
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

// ParseQuorumUpdated is a log parse operation binding the contract event 0xee9c9cc7dfc433e84d31f1d7a340b6aa0d46bd17055346a78997f3354a4250d4.
//
// Solidity: event QuorumUpdated(uint8 indexed quorumNumber, uint256 index, (address,uint96) multiplier)
func (_Voteweigher *VoteweigherFilterer) ParseQuorumUpdated(log types.Log) (*VoteweigherQuorumUpdated, error) {
	event := new(VoteweigherQuorumUpdated)
	if err := _Voteweigher.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
