// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LagrangeCommittee

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

// LagrangeCommitteeCommitteeLeaf is an auto generated low-level Go binding around an user-defined struct.
type LagrangeCommitteeCommitteeLeaf struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}

// LagrangeCommitteeMetaData contains all meta data concerning the LagrangeCommittee contract.
var LagrangeCommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon4Elements\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"}],\"name\":\"InitCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"current\",\"type\":\"bytes32\"}],\"name\":\"UpdateCommittee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNT_CREATION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUTHORISE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_CURRENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeLeaves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERMEZ_NETWORK_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seqAddr\",\"type\":\"address\"}],\"name\":\"addSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addr2bls\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getCommittee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structLagrangeCommittee.CommitteeLeaf\",\"name\":\"cleaf\",\"type\":\"tuple\"}],\"name\":\"getLeafHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getNext1CommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getServeUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"hash2Elements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"stakedAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"name\":\"setSlashed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"comparisonNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"verifyBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106102065760003560e01c80639dfa4e941161011a578063d4b96c44116100ad578063f1f2fcab1161007c578063f1f2fcab1461020b578063f2fde38b14610612578063f5425bd514610625578063f96ff9f21461062d578063fe095d5c1461064057600080fd5b8063d4b96c44146105bd578063e04392f1146105d0578063e62f6b92146105d8578063f0d60817146105ff57600080fd5b8063c2b84126116100e9578063c2b8412614610538578063c364091e14610558578063c473af331461056b578063cf7aa2111461059257600080fd5b80639dfa4e94146104d85780639e4e7318146104eb578063ac407e1714610512578063adacd9921461052557600080fd5b806355c1c2bc1161019d5780637d99c8641161016c5780637d99c8641461046857806382ab890a146104705780638a336231146104835780638da5cb5b146104965780639743c7b7146104b657600080fd5b806355c1c2bc14610418578063620390221461042b578063715018a61461043e5780637987e1991461044857600080fd5b80633408e470116101d95780633408e470146102db5780633644e515146102e157806344a5c4bf1461038a57806344f5b6b4146103d157600080fd5b806304622c2e1461020b5780631300aff01461024557806313e7c9d81461026c5780632b846504146102c8575b600080fd5b6102327fbe287413178bfeddef8d9753ad4be825ae998706a6dabff23978b59dccaea0ad81565b6040519081526020015b60405180910390f35b6102327fff946cf82975b1a2b6e6d28c9a76a4b8d7a1fd0592b785cb92771933310f9ee781565b6102a661027a3660046119e2565b6072602052600090815260409020805460019091015463ffffffff811690640100000000900460ff1683565b6040805193845263ffffffff909216602084015215159082015260600161023c565b6102326102d63660046119fd565b610653565b46610232565b61023260007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7fbe287413178bfeddef8d9753ad4be825ae998706a6dabff23978b59dccaea0ad7fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646604080516020810195909552840192909252606083015260808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6103bc6103983660046119e2565b6001600160a01b031660009081526072602052604090206001015463ffffffff1690565b60405163ffffffff909116815260200161023c565b6104086103df3660046119e2565b6001600160a01b0316600090815260726020526040902060010154640100000000900460ff1690565b604051901515815260200161023c565b610232610426366004611ac2565b610684565b610408610439366004611b60565b6106ce565b6104466106e5565b005b61045b6104563660046119e2565b6106f9565b60405161023c9190611c04565b610232600181565b61044661047e366004611c17565b610793565b6104466104913660046119e2565b610a9a565b61049e610ac6565b6040516001600160a01b03909116815260200161023c565b6104c96104c43660046119fd565b610adf565b60405161023c93929190611c30565b6104466104e6366004611c57565b610b9f565b6102327fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc681565b610446610520366004611c83565b610c12565b610232610533366004611c17565b610c80565b610232610546366004611c17565b606d6020526000908152604090205481565b6102326105663660046119fd565b610dc3565b6102327f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6102326105a03660046119fd565b6000918252606c6020908152604080842092845291905290205490565b6102326105cb3660046119fd565b610df9565b610232600281565b6102327fafd642c6a37a2e6887dc4ad5142f84197828a904e53d3204ecb1100329231eaa81565b61044661060d366004611cbf565b610e15565b6104466106203660046119e2565b610ef9565b610232600081565b61044661063b366004611d2b565b610f72565b61023261064e3660046119fd565b610fda565b6070602052816000526040600020818154811061066f57600080fd5b90600052602060002001600091509150505481565b60006106c8604051806060016040528084600001516001600160a01b031681526020018460200151815260200184604001518051906020012060001c81525061100f565b92915050565b60006106dc858585856106ce565b95945050505050565b6106ed611081565b6106f760006110e0565b565b6071602052600090815260409020805461071290611db9565b80601f016020809104026020016040519081016040528092919081815260200182805461073e90611db9565b801561078b5780601f106107605761010080835404028352916020019161078b565b820191906000526020600020905b81548152906001019060200180831161076e57829003601f168201915b505050505081565b3360009081526068602052604090205460ff1615156001146107d05760405162461bcd60e51b81526004016107c790611df4565b60405180910390fd5b60006107dc8243610dc3565b6000838152606b6020526040812060010154919250906107fc9083611e56565b6000848152606b602052604090206002015490915061081b8183611e6e565b43116108835760405162461bcd60e51b815260206004820152603160248201527f426c6f636b206e756d626572206973207072696f7220746f20636f6d6d69747460448201527032b290333932b2bd32903bb4b73237bb9760791b60648201526084016107c7565b60005b6000858152606960205260409020548110156109895760008581526069602052604090208054610977918791849081106108c2576108c2611e85565b600091825260208083209091015433835260719091526040822080546001600160a01b0390921692916108f490611db9565b80601f016020809104026020016040519081016040528092919081815260200182805461092090611db9565b801561096d5780601f106109425761010080835404028352916020019161096d565b820191906000526020600020905b81548152906001019060200180831161095057829003601f168201915b5050505050611132565b8061098181611e9b565b915050610886565b5060005b6000858152606a60205260409020548110156109f5576000858152606a6020526040902080546109e3918791849081106109c9576109c9611e85565b6000918252602090912001546001600160a01b03166112da565b806109ed81611e9b565b91505061098d565b506000848152606960205260408120610a0d9161185a565b6000848152606a60205260408120610a249161185a565b610a2d84611542565b6000848152606c602052604081207fc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd90891869190610a6b600188611e56565b81526020808201929092526040908101600020548151938452918301919091520160405180910390a150505050565b610aa2611081565b6001600160a01b03166000908152606860205260409020805460ff19166001179055565b6000610ada6033546001600160a01b031690565b905090565b606f6020908152600092835260408084209091529082529020805460018201546002830180546001600160a01b03909316939192610b1c90611db9565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4890611db9565b8015610b955780601f10610b6a57610100808354040283529160200191610b95565b820191906000526020600020905b815481529060010190602001808311610b7857829003601f168201915b5050505050905083565b3360009081526068602052604090205460ff161515600114610bd35760405162461bcd60e51b81526004016107c790611df4565b6000918252606a6020908152604083208054600181018255908452922090910180546001600160a01b0319166001600160a01b03909216919091179055565b3360009081526068602052604090205460ff161515600114610c465760405162461bcd60e51b81526004016107c790611df4565b6001600160a01b03909116600090815260726020526040902060010180549115156401000000000264ff0000000019909216919091179055565b600081815260706020526040812054610c9e576106c8600080610fda565b60008281526070602052604090205460011415610ce65760008281526070602052604081208054909190610cd457610cd4611e85565b90600052602060002001549050919050565b60025b600083815260706020526040902054811015610d1157610d0a600282611eb6565b9050610ce9565b600083815260706020526040812080548290610d2f57610d2f611e85565b600091825260209091200154905060015b82811015610dbb57600085815260706020526040902054811015610d9b5760008581526070602052604090208054610d9491849184908110610d8457610d84611e85565b9060005260206000200154610fda565b9150610da9565b610da6826000610fda565b91505b80610db381611e9b565b915050610d40565b509392505050565b6000828152606b6020526040812080546001909101548281610de58487611e6e565b610def9190611ed5565b9695505050505050565b606e602052816000526040600020818154811061066f57600080fd5b3360009081526068602052604090205460ff161515600114610e495760405162461bcd60e51b81526004016107c790611df4565b600084815260696020908152604080832080546001810182559084528284200180546001600160a01b031916339081179091558352607182529091208451610e9392860190611878565b506040805160608101825292835263ffffffff9182166020808501918252600085840181815233825260729092529290922093518455516001939093018054915115156401000000000264ffffffffff1990921693909216929092179190911790555050565b610f01611081565b6001600160a01b038116610f665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107c7565b610f6f816110e0565b50565b610f7a611081565b610f85858383611590565b60005b83811015610fd257610fc086868684818110610fa657610fa6611e85565b9050602002016020810190610fbb91906119e2565b611740565b80610fca81611e9b565b915050610f88565b505050505050565b6000610ff9604051806040016040528085815260200184815250611829565b9392505050565b6001600160a01b03163b151590565b6066546040516304b98e1d60e31b81526000916001600160a01b0316906325cc70e890611040908590600401611ef7565b602060405180830381865afa15801561105d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c89190611f28565b3361108a610ac6565b6001600160a01b0316146106f75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107c7565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b3360009081526068602052604090205460ff1615156001146111665760405162461bcd60e51b81526004016107c790611df4565b6000848152606b60205260409020546111e75760405162461bcd60e51b815260206004820152603760248201527f4120636f6d6d697474656520666f72207468697320636861696e20494420686160448201527f73206e6f74206265656e20696e697469616c697a65642e00000000000000000060648201526084016107c7565b604080516060810182526001600160a01b038516815260208101849052908101829052600061121582610684565b6000878152606f60209081526040808320848452825291829020855181546001600160a01b0319166001600160a01b03909116178155858201516001820155918501518051939450859361126f9260028501920190611878565b5050506000868152606e60209081526040808320805460018101825590845282842001849055888352606d90915281208054916112ab83611e9b565b909155505060009586526070602090815260408720805460018101825590885296209095019490945550505050565b3360009081526068602052604090205460ff16151560011461130e5760405162461bcd60e51b81526004016107c790611df4565b60005b6000838152606e602052604090205481101561153d576000838152606e6020526040812080548390811061134757611347611e85565b6000918252602080832090910154868352606e9091526040822080549193509061137390600190611e6e565b8154811061138357611383611e85565b6000918252602080832090910154878352606f825260408084208685529092529120549091506001600160a01b039081169085161415611528576000858152606f602090815260408083208484529091528082208483529120815481546001600160a01b0319166001600160a01b0390911617815560018083015490820155600280830180549183019161141690611db9565b6114219291906118fc565b50506040805160608101825260008082526020808301828152845180830186528381528486019081528b8452606f8352858420888552835294909220835181546001600160a01b0319166001600160a01b039091161781559151600183015592518051929450909261149b92600285019290910190611878565b5050506000858152606e6020526040902080546114ba90600190611e6e565b815481106114ca576114ca611e85565b9060005260206000200154606e600087815260200190815260200160002084815481106114f9576114f9611e85565b6000918252602080832090910192909255868152606d9091526040812080549161152283611f41565b91905055505b5050808061153590611e9b565b915050611311565b505050565b600061154d82610c80565b9050600061155b8343610dc3565b6000848152606c6020526040812091925083919061157a600185611e56565b8152602081019190915260400160002055505050565b3360009081526068602052604090205460ff1615156001146115c45760405162461bcd60e51b81526004016107c790611df4565b6000838152606b6020526040902054156116305760405162461bcd60e51b815260206004820152602760248201527f436f6d6d69747465652068617320616c7265616479206265656e20696e69746960448201526630b634bd32b21760c91b60648201526084016107c7565b6040805160608101825243815260208082018581528284018581526000888152606b8452858120945185559151600180860191909155905160029094019390935583518085018552818152808301828152888352606c8452858320838052845285832091518255519301929092558251828152808201808552878452606e9092529290912091516116c2929190611977565b506000838152606d6020908152604080832083905580518381528083018083528785526070909352922091516116f9929190611977565b5060408051848152602081018490529081018290527fd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc49060600160405180910390a1505050565b3360009081526068602052604090205460ff1615156001146117745760405162461bcd60e51b81526004016107c790611df4565b60005b6000838152606960205260409020548110156117eb57600083815260696020526040902080546001600160a01b0384169190839081106117b9576117b9611e85565b6000918252602090912001546001600160a01b031614156117d957505050565b806117e381611e9b565b915050611777565b5060008281526069602090815260408220805460018101825590835291200180546001600160a01b0383166001600160a01b03199091161790555050565b6065546040516314d2f97b60e11b81526000916001600160a01b0316906329a5f2f690611040908590600401611f58565b5080546000825590600052602060002090810190610f6f91906119b1565b82805461188490611db9565b90600052602060002090601f0160209004810192826118a657600085556118ec565b82601f106118bf57805160ff19168380011785556118ec565b828001600101855582156118ec579182015b828111156118ec5782518255916020019190600101906118d1565b506118f89291506119b1565b5090565b82805461190890611db9565b90600052602060002090601f01602090048101928261192a57600085556118ec565b82601f1061193b57805485556118ec565b828001600101855582156118ec57600052602060002091601f016020900482015b828111156118ec57825482559160010191906001019061195c565b8280548282559060005260206000209081019282156118ec57916020028201828111156118ec5782518255916020019190600101906118d1565b5b808211156118f857600081556001016119b2565b80356001600160a01b03811681146119dd57600080fd5b919050565b6000602082840312156119f457600080fd5b610ff9826119c6565b60008060408385031215611a1057600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611a4657600080fd5b813567ffffffffffffffff80821115611a6157611a61611a1f565b604051601f8301601f19908116603f01168101908282118183101715611a8957611a89611a1f565b81604052838152866020858801011115611aa257600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215611ad457600080fd5b813567ffffffffffffffff80821115611aec57600080fd5b9083019060608286031215611b0057600080fd5b604051606081018181108382111715611b1b57611b1b611a1f565b604052611b27836119c6565b815260208301356020820152604083013582811115611b4557600080fd5b611b5187828601611a35565b60408301525095945050505050565b60008060008060808587031215611b7657600080fd5b84359350602085013567ffffffffffffffff811115611b9457600080fd5b611ba087828801611a35565b949794965050505060408301359260600135919050565b6000815180845260005b81811015611bdd57602081850181015186830182015201611bc1565b81811115611bef576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610ff96020830184611bb7565b600060208284031215611c2957600080fd5b5035919050565b60018060a01b03841681528260208201526060604082015260006106dc6060830184611bb7565b60008060408385031215611c6a57600080fd5b82359150611c7a602084016119c6565b90509250929050565b60008060408385031215611c9657600080fd5b611c9f836119c6565b915060208301358015158114611cb457600080fd5b809150509250929050565b60008060008060808587031215611cd557600080fd5b84359350602085013567ffffffffffffffff811115611cf357600080fd5b611cff87828801611a35565b93505060408501359150606085013563ffffffff81168114611d2057600080fd5b939692955090935050565b600080600080600060808688031215611d4357600080fd5b85359450602086013567ffffffffffffffff80821115611d6257600080fd5b818801915088601f830112611d7657600080fd5b813581811115611d8557600080fd5b8960208260051b8501011115611d9a57600080fd5b9699602092909201985095966040810135965060600135945092505050565b600181811c90821680611dcd57607f821691505b60208210811415611dee57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602c908201527f4f6e6c792073657175656e636572206e6f6465732063616e2063616c6c20746860408201526b34b990333ab731ba34b7b71760a11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008219821115611e6957611e69611e40565b500190565b600082821015611e8057611e80611e40565b500390565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611eaf57611eaf611e40565b5060010190565b6000816000190483118215151615611ed057611ed0611e40565b500290565b600082611ef257634e487b7160e01b600052601260045260246000fd5b500490565b60608101818360005b6003811015611f1f578151835260209283019290910190600101611f00565b50505092915050565b600060208284031215611f3a57600080fd5b5051919050565b600081611f5057611f50611e40565b506000190190565b60408101818360005b6002811015611f1f578151835260209283019290910190600101611f6156fea264697066735822122022466a86931767e6f1255e9137593ee04d6d77893abde232117d35b34bc087d564736f6c634300080c0033",
}

// LagrangeCommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangeCommitteeMetaData.ABI instead.
var LagrangeCommitteeABI = LagrangeCommitteeMetaData.ABI

// LagrangeCommitteeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LagrangeCommitteeMetaData.Bin instead.
var LagrangeCommitteeBin = LagrangeCommitteeMetaData.Bin

// DeployLagrangeCommittee deploys a new Ethereum contract, binding an instance of LagrangeCommittee to it.
func DeployLagrangeCommittee(auth *bind.TransactOpts, backend bind.ContractBackend, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (common.Address, *types.Transaction, *LagrangeCommittee, error) {
	parsed, err := LagrangeCommitteeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LagrangeCommitteeBin), backend, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LagrangeCommittee{LagrangeCommitteeCaller: LagrangeCommitteeCaller{contract: contract}, LagrangeCommitteeTransactor: LagrangeCommitteeTransactor{contract: contract}, LagrangeCommitteeFilterer: LagrangeCommitteeFilterer{contract: contract}}, nil
}

// LagrangeCommittee is an auto generated Go binding around an Ethereum contract.
type LagrangeCommittee struct {
	LagrangeCommitteeCaller     // Read-only binding to the contract
	LagrangeCommitteeTransactor // Write-only binding to the contract
	LagrangeCommitteeFilterer   // Log filterer for contract events
}

// LagrangeCommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagrangeCommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeCommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagrangeCommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeCommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagrangeCommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeCommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagrangeCommitteeSession struct {
	Contract     *LagrangeCommittee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LagrangeCommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagrangeCommitteeCallerSession struct {
	Contract *LagrangeCommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LagrangeCommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagrangeCommitteeTransactorSession struct {
	Contract     *LagrangeCommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LagrangeCommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagrangeCommitteeRaw struct {
	Contract *LagrangeCommittee // Generic contract binding to access the raw methods on
}

// LagrangeCommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagrangeCommitteeCallerRaw struct {
	Contract *LagrangeCommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// LagrangeCommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagrangeCommitteeTransactorRaw struct {
	Contract *LagrangeCommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLagrangeCommittee creates a new instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommittee(address common.Address, backend bind.ContractBackend) (*LagrangeCommittee, error) {
	contract, err := bindLagrangeCommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommittee{LagrangeCommitteeCaller: LagrangeCommitteeCaller{contract: contract}, LagrangeCommitteeTransactor: LagrangeCommitteeTransactor{contract: contract}, LagrangeCommitteeFilterer: LagrangeCommitteeFilterer{contract: contract}}, nil
}

// NewLagrangeCommitteeCaller creates a new read-only instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommitteeCaller(address common.Address, caller bind.ContractCaller) (*LagrangeCommitteeCaller, error) {
	contract, err := bindLagrangeCommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeCaller{contract: contract}, nil
}

// NewLagrangeCommitteeTransactor creates a new write-only instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*LagrangeCommitteeTransactor, error) {
	contract, err := bindLagrangeCommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeTransactor{contract: contract}, nil
}

// NewLagrangeCommitteeFilterer creates a new log filterer instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*LagrangeCommitteeFilterer, error) {
	contract, err := bindLagrangeCommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeFilterer{contract: contract}, nil
}

// bindLagrangeCommittee binds a generic wrapper to an already deployed contract.
func bindLagrangeCommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LagrangeCommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LagrangeCommittee *LagrangeCommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LagrangeCommittee.Contract.LagrangeCommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LagrangeCommittee *LagrangeCommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.LagrangeCommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LagrangeCommittee *LagrangeCommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.LagrangeCommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LagrangeCommittee *LagrangeCommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LagrangeCommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LagrangeCommittee *LagrangeCommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LagrangeCommittee *LagrangeCommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) ACCOUNTCREATIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "ACCOUNT_CREATION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.ACCOUNTCREATIONHASH(&_LagrangeCommittee.CallOpts)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.ACCOUNTCREATIONHASH(&_LagrangeCommittee.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) AUTHORISETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "AUTHORISE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.AUTHORISETYPEHASH(&_LagrangeCommittee.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.AUTHORISETYPEHASH(&_LagrangeCommittee.CallOpts)
}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) COMMITTEECURRENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "COMMITTEE_CURRENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) COMMITTEECURRENT() (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEECURRENT(&_LagrangeCommittee.CallOpts)
}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) COMMITTEECURRENT() (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEECURRENT(&_LagrangeCommittee.CallOpts)
}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) COMMITTEENEXT1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "COMMITTEE_NEXT_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) COMMITTEENEXT1() (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEENEXT1(&_LagrangeCommittee.CallOpts)
}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) COMMITTEENEXT1() (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEENEXT1(&_LagrangeCommittee.CallOpts)
}

// COMMITTEENEXT2 is a free data retrieval call binding the contract method 0xe04392f1.
//
// Solidity: function COMMITTEE_NEXT_2() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) COMMITTEENEXT2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "COMMITTEE_NEXT_2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEENEXT2 is a free data retrieval call binding the contract method 0xe04392f1.
//
// Solidity: function COMMITTEE_NEXT_2() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) COMMITTEENEXT2() (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEENEXT2(&_LagrangeCommittee.CallOpts)
}

// COMMITTEENEXT2 is a free data retrieval call binding the contract method 0xe04392f1.
//
// Solidity: function COMMITTEE_NEXT_2() view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) COMMITTEENEXT2() (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEENEXT2(&_LagrangeCommittee.CallOpts)
}

// CommitteeLeaves is a free data retrieval call binding the contract method 0x2b846504.
//
// Solidity: function CommitteeLeaves(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeLeaves(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "CommitteeLeaves", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeLeaves is a free data retrieval call binding the contract method 0x2b846504.
//
// Solidity: function CommitteeLeaves(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeLeaves(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeLeaves(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeLeaves is a free data retrieval call binding the contract method 0x2b846504.
//
// Solidity: function CommitteeLeaves(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeLeaves(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeLeaves(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeMap(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "CommitteeMap", arg0, arg1)

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
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeMap(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	return _LagrangeCommittee.Contract.CommitteeMap(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeMap(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	return _LagrangeCommittee.Contract.CommitteeMap(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeMapKeys(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "CommitteeMapKeys", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeMapKeys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeMapKeys(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeMapKeys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeMapKeys(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeMapLength(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "CommitteeMapLength", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeMapLength(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeMapLength(&_LagrangeCommittee.CallOpts, arg0)
}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeMapLength(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeMapLength(&_LagrangeCommittee.CallOpts, arg0)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_LagrangeCommittee *LagrangeCommitteeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_LagrangeCommittee *LagrangeCommitteeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LagrangeCommittee.Contract.DOMAINSEPARATOR(&_LagrangeCommittee.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LagrangeCommittee.Contract.DOMAINSEPARATOR(&_LagrangeCommittee.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "EIP712DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.EIP712DOMAINHASH(&_LagrangeCommittee.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.EIP712DOMAINHASH(&_LagrangeCommittee.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) HERMEZNETWORKHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "HERMEZ_NETWORK_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.HERMEZNETWORKHASH(&_LagrangeCommittee.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.HERMEZNETWORKHASH(&_LagrangeCommittee.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) NAMEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "NAME_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) NAMEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.NAMEHASH(&_LagrangeCommittee.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) NAMEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.NAMEHASH(&_LagrangeCommittee.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) VERSIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "VERSION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) VERSIONHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.VERSIONHASH(&_LagrangeCommittee.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) VERSIONHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.VERSIONHASH(&_LagrangeCommittee.CallOpts)
}

// Addr2bls is a free data retrieval call binding the contract method 0x7987e199.
//
// Solidity: function addr2bls(address ) view returns(bytes)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Addr2bls(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "addr2bls", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr2bls is a free data retrieval call binding the contract method 0x7987e199.
//
// Solidity: function addr2bls(address ) view returns(bytes)
func (_LagrangeCommittee *LagrangeCommitteeSession) Addr2bls(arg0 common.Address) ([]byte, error) {
	return _LagrangeCommittee.Contract.Addr2bls(&_LagrangeCommittee.CallOpts, arg0)
}

// Addr2bls is a free data retrieval call binding the contract method 0x7987e199.
//
// Solidity: function addr2bls(address ) view returns(bytes)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Addr2bls(arg0 common.Address) ([]byte, error) {
	return _LagrangeCommittee.Contract.Addr2bls(&_LagrangeCommittee.CallOpts, arg0)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetChainId() (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetChainId(&_LagrangeCommittee.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetChainId() (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetChainId(&_LagrangeCommittee.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0xc364091e.
//
// Solidity: function getEpochNumber(uint256 chainID, uint256 blockNumber) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetEpochNumber(opts *bind.CallOpts, chainID *big.Int, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getEpochNumber", chainID, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0xc364091e.
//
// Solidity: function getEpochNumber(uint256 chainID, uint256 blockNumber) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetEpochNumber(chainID *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetEpochNumber(&_LagrangeCommittee.CallOpts, chainID, blockNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0xc364091e.
//
// Solidity: function getEpochNumber(uint256 chainID, uint256 blockNumber) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetEpochNumber(chainID *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetEpochNumber(&_LagrangeCommittee.CallOpts, chainID, blockNumber)
}

// GetLeafHash is a free data retrieval call binding the contract method 0x55c1c2bc.
//
// Solidity: function getLeafHash((address,uint256,bytes) cleaf) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetLeafHash(opts *bind.CallOpts, cleaf LagrangeCommitteeCommitteeLeaf) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getLeafHash", cleaf)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLeafHash is a free data retrieval call binding the contract method 0x55c1c2bc.
//
// Solidity: function getLeafHash((address,uint256,bytes) cleaf) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetLeafHash(cleaf LagrangeCommitteeCommitteeLeaf) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetLeafHash(&_LagrangeCommittee.CallOpts, cleaf)
}

// GetLeafHash is a free data retrieval call binding the contract method 0x55c1c2bc.
//
// Solidity: function getLeafHash((address,uint256,bytes) cleaf) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetLeafHash(cleaf LagrangeCommitteeCommitteeLeaf) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetLeafHash(&_LagrangeCommittee.CallOpts, cleaf)
}

// GetNext1CommitteeRoot is a free data retrieval call binding the contract method 0xadacd992.
//
// Solidity: function getNext1CommitteeRoot(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetNext1CommitteeRoot(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getNext1CommitteeRoot", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNext1CommitteeRoot is a free data retrieval call binding the contract method 0xadacd992.
//
// Solidity: function getNext1CommitteeRoot(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetNext1CommitteeRoot(chainID *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetNext1CommitteeRoot(&_LagrangeCommittee.CallOpts, chainID)
}

// GetNext1CommitteeRoot is a free data retrieval call binding the contract method 0xadacd992.
//
// Solidity: function getNext1CommitteeRoot(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetNext1CommitteeRoot(chainID *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetNext1CommitteeRoot(&_LagrangeCommittee.CallOpts, chainID)
}

// Hash2Elements is a free data retrieval call binding the contract method 0xfe095d5c.
//
// Solidity: function hash2Elements(uint256 a, uint256 b) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Hash2Elements(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "hash2Elements", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hash2Elements is a free data retrieval call binding the contract method 0xfe095d5c.
//
// Solidity: function hash2Elements(uint256 a, uint256 b) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) Hash2Elements(a *big.Int, b *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Hash2Elements(&_LagrangeCommittee.CallOpts, a, b)
}

// Hash2Elements is a free data retrieval call binding the contract method 0xfe095d5c.
//
// Solidity: function hash2Elements(uint256 a, uint256 b) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Hash2Elements(a *big.Int, b *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Hash2Elements(&_LagrangeCommittee.CallOpts, a, b)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, uint32 serveUntilBlock, bool slashed)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount          *big.Int
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "operators", arg0)

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
func (_LagrangeCommittee *LagrangeCommitteeSession) Operators(arg0 common.Address) (struct {
	Amount          *big.Int
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	return _LagrangeCommittee.Contract.Operators(&_LagrangeCommittee.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, uint32 serveUntilBlock, bool slashed)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Operators(arg0 common.Address) (struct {
	Amount          *big.Int
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	return _LagrangeCommittee.Contract.Operators(&_LagrangeCommittee.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeSession) Owner() (common.Address, error) {
	return _LagrangeCommittee.Contract.Owner(&_LagrangeCommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Owner() (common.Address, error) {
	return _LagrangeCommittee.Contract.Owner(&_LagrangeCommittee.CallOpts)
}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCaller) VerifyBlockNumber(opts *bind.CallOpts, comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "verifyBlockNumber", comparisonNumber, rlpData, comparisonBlockHash, chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeSession) VerifyBlockNumber(comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	return _LagrangeCommittee.Contract.VerifyBlockNumber(&_LagrangeCommittee.CallOpts, comparisonNumber, rlpData, comparisonBlockHash, chainID)
}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) VerifyBlockNumber(comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	return _LagrangeCommittee.Contract.VerifyBlockNumber(&_LagrangeCommittee.CallOpts, comparisonNumber, rlpData, comparisonBlockHash, chainID)
}

// Add is a paid mutator transaction binding the contract method 0xf0d60817.
//
// Solidity: function add(uint256 chainID, bytes blsPubKey, uint256 stake, uint32 serveUntilBlock) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) Add(opts *bind.TransactOpts, chainID *big.Int, blsPubKey []byte, stake *big.Int, serveUntilBlock uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "add", chainID, blsPubKey, stake, serveUntilBlock)
}

// Add is a paid mutator transaction binding the contract method 0xf0d60817.
//
// Solidity: function add(uint256 chainID, bytes blsPubKey, uint256 stake, uint32 serveUntilBlock) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) Add(chainID *big.Int, blsPubKey []byte, stake *big.Int, serveUntilBlock uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Add(&_LagrangeCommittee.TransactOpts, chainID, blsPubKey, stake, serveUntilBlock)
}

// Add is a paid mutator transaction binding the contract method 0xf0d60817.
//
// Solidity: function add(uint256 chainID, bytes blsPubKey, uint256 stake, uint32 serveUntilBlock) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) Add(chainID *big.Int, blsPubKey []byte, stake *big.Int, serveUntilBlock uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Add(&_LagrangeCommittee.TransactOpts, chainID, blsPubKey, stake, serveUntilBlock)
}

// AddSequencer is a paid mutator transaction binding the contract method 0x8a336231.
//
// Solidity: function addSequencer(address seqAddr) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) AddSequencer(opts *bind.TransactOpts, seqAddr common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "addSequencer", seqAddr)
}

// AddSequencer is a paid mutator transaction binding the contract method 0x8a336231.
//
// Solidity: function addSequencer(address seqAddr) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) AddSequencer(seqAddr common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.AddSequencer(&_LagrangeCommittee.TransactOpts, seqAddr)
}

// AddSequencer is a paid mutator transaction binding the contract method 0x8a336231.
//
// Solidity: function addSequencer(address seqAddr) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) AddSequencer(seqAddr common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.AddSequencer(&_LagrangeCommittee.TransactOpts, seqAddr)
}

// GetCommittee is a paid mutator transaction binding the contract method 0xcf7aa211.
//
// Solidity: function getCommittee(uint256 chainID, uint256 epochNumber) returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeTransactor) GetCommittee(opts *bind.TransactOpts, chainID *big.Int, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "getCommittee", chainID, epochNumber)
}

// GetCommittee is a paid mutator transaction binding the contract method 0xcf7aa211.
//
// Solidity: function getCommittee(uint256 chainID, uint256 epochNumber) returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetCommittee(chainID *big.Int, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.GetCommittee(&_LagrangeCommittee.TransactOpts, chainID, epochNumber)
}

// GetCommittee is a paid mutator transaction binding the contract method 0xcf7aa211.
//
// Solidity: function getCommittee(uint256 chainID, uint256 epochNumber) returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) GetCommittee(chainID *big.Int, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.GetCommittee(&_LagrangeCommittee.TransactOpts, chainID, epochNumber)
}

// GetServeUntilBlock is a paid mutator transaction binding the contract method 0x44a5c4bf.
//
// Solidity: function getServeUntilBlock(address operator) returns(uint32)
func (_LagrangeCommittee *LagrangeCommitteeTransactor) GetServeUntilBlock(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "getServeUntilBlock", operator)
}

// GetServeUntilBlock is a paid mutator transaction binding the contract method 0x44a5c4bf.
//
// Solidity: function getServeUntilBlock(address operator) returns(uint32)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetServeUntilBlock(operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.GetServeUntilBlock(&_LagrangeCommittee.TransactOpts, operator)
}

// GetServeUntilBlock is a paid mutator transaction binding the contract method 0x44a5c4bf.
//
// Solidity: function getServeUntilBlock(address operator) returns(uint32)
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) GetServeUntilBlock(operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.GetServeUntilBlock(&_LagrangeCommittee.TransactOpts, operator)
}

// GetSlashed is a paid mutator transaction binding the contract method 0x44f5b6b4.
//
// Solidity: function getSlashed(address operator) returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeTransactor) GetSlashed(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "getSlashed", operator)
}

// GetSlashed is a paid mutator transaction binding the contract method 0x44f5b6b4.
//
// Solidity: function getSlashed(address operator) returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetSlashed(operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.GetSlashed(&_LagrangeCommittee.TransactOpts, operator)
}

// GetSlashed is a paid mutator transaction binding the contract method 0x44f5b6b4.
//
// Solidity: function getSlashed(address operator) returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) GetSlashed(operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.GetSlashed(&_LagrangeCommittee.TransactOpts, operator)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xf96ff9f2.
//
// Solidity: function registerChain(uint256 chainID, address[] stakedAddrs, uint256 epochPeriod, uint256 freezeDuration) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RegisterChain(opts *bind.TransactOpts, chainID *big.Int, stakedAddrs []common.Address, epochPeriod *big.Int, freezeDuration *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "registerChain", chainID, stakedAddrs, epochPeriod, freezeDuration)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xf96ff9f2.
//
// Solidity: function registerChain(uint256 chainID, address[] stakedAddrs, uint256 epochPeriod, uint256 freezeDuration) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RegisterChain(chainID *big.Int, stakedAddrs []common.Address, epochPeriod *big.Int, freezeDuration *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RegisterChain(&_LagrangeCommittee.TransactOpts, chainID, stakedAddrs, epochPeriod, freezeDuration)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xf96ff9f2.
//
// Solidity: function registerChain(uint256 chainID, address[] stakedAddrs, uint256 epochPeriod, uint256 freezeDuration) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RegisterChain(chainID *big.Int, stakedAddrs []common.Address, epochPeriod *big.Int, freezeDuration *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RegisterChain(&_LagrangeCommittee.TransactOpts, chainID, stakedAddrs, epochPeriod, freezeDuration)
}

// Remove is a paid mutator transaction binding the contract method 0x9dfa4e94.
//
// Solidity: function remove(uint256 chainID, address addr) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) Remove(opts *bind.TransactOpts, chainID *big.Int, addr common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "remove", chainID, addr)
}

// Remove is a paid mutator transaction binding the contract method 0x9dfa4e94.
//
// Solidity: function remove(uint256 chainID, address addr) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) Remove(chainID *big.Int, addr common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Remove(&_LagrangeCommittee.TransactOpts, chainID, addr)
}

// Remove is a paid mutator transaction binding the contract method 0x9dfa4e94.
//
// Solidity: function remove(uint256 chainID, address addr) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) Remove(chainID *big.Int, addr common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Remove(&_LagrangeCommittee.TransactOpts, chainID, addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RenounceOwnership(&_LagrangeCommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RenounceOwnership(&_LagrangeCommittee.TransactOpts)
}

// SetSlashed is a paid mutator transaction binding the contract method 0xac407e17.
//
// Solidity: function setSlashed(address operator, bool slashed) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) SetSlashed(opts *bind.TransactOpts, operator common.Address, slashed bool) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "setSlashed", operator, slashed)
}

// SetSlashed is a paid mutator transaction binding the contract method 0xac407e17.
//
// Solidity: function setSlashed(address operator, bool slashed) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) SetSlashed(operator common.Address, slashed bool) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.SetSlashed(&_LagrangeCommittee.TransactOpts, operator, slashed)
}

// SetSlashed is a paid mutator transaction binding the contract method 0xac407e17.
//
// Solidity: function setSlashed(address operator, bool slashed) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) SetSlashed(operator common.Address, slashed bool) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.SetSlashed(&_LagrangeCommittee.TransactOpts, operator, slashed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.TransferOwnership(&_LagrangeCommittee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.TransferOwnership(&_LagrangeCommittee.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) Update(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "update", chainID)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) Update(chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Update(&_LagrangeCommittee.TransactOpts, chainID)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) Update(chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Update(&_LagrangeCommittee.TransactOpts, chainID)
}

// LagrangeCommitteeInitCommitteeIterator is returned from FilterInitCommittee and is used to iterate over the raw logs and unpacked data for InitCommittee events raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitCommitteeIterator struct {
	Event *LagrangeCommitteeInitCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeInitCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeInitCommittee)
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
		it.Event = new(LagrangeCommitteeInitCommittee)
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
func (it *LagrangeCommitteeInitCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeInitCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeInitCommittee represents a InitCommittee event raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitCommittee struct {
	ChainID        *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitCommittee is a free log retrieval operation binding the contract event 0xd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc4.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterInitCommittee(opts *bind.FilterOpts) (*LagrangeCommitteeInitCommitteeIterator, error) {

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeInitCommitteeIterator{contract: _LagrangeCommittee.contract, event: "InitCommittee", logs: logs, sub: sub}, nil
}

// WatchInitCommittee is a free log subscription operation binding the contract event 0xd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc4.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchInitCommittee(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeInitCommittee) (event.Subscription, error) {

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeInitCommittee)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
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

// ParseInitCommittee is a log parse operation binding the contract event 0xd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc4.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseInitCommittee(log types.Log) (*LagrangeCommitteeInitCommittee, error) {
	event := new(LagrangeCommitteeInitCommittee)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitializedIterator struct {
	Event *LagrangeCommitteeInitialized // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeInitialized)
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
		it.Event = new(LagrangeCommitteeInitialized)
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
func (it *LagrangeCommitteeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeInitialized represents a Initialized event raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LagrangeCommitteeInitializedIterator, error) {

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeInitializedIterator{contract: _LagrangeCommittee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeInitialized) (event.Subscription, error) {

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeInitialized)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseInitialized(log types.Log) (*LagrangeCommitteeInitialized, error) {
	event := new(LagrangeCommitteeInitialized)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LagrangeCommittee contract.
type LagrangeCommitteeOwnershipTransferredIterator struct {
	Event *LagrangeCommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeOwnershipTransferred)
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
		it.Event = new(LagrangeCommitteeOwnershipTransferred)
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
func (it *LagrangeCommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the LagrangeCommittee contract.
type LagrangeCommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LagrangeCommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeOwnershipTransferredIterator{contract: _LagrangeCommittee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeOwnershipTransferred)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*LagrangeCommitteeOwnershipTransferred, error) {
	event := new(LagrangeCommitteeOwnershipTransferred)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeUpdateCommitteeIterator is returned from FilterUpdateCommittee and is used to iterate over the raw logs and unpacked data for UpdateCommittee events raised by the LagrangeCommittee contract.
type LagrangeCommitteeUpdateCommitteeIterator struct {
	Event *LagrangeCommitteeUpdateCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeUpdateCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeUpdateCommittee)
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
		it.Event = new(LagrangeCommitteeUpdateCommittee)
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
func (it *LagrangeCommitteeUpdateCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeUpdateCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeUpdateCommittee represents a UpdateCommittee event raised by the LagrangeCommittee contract.
type LagrangeCommitteeUpdateCommittee struct {
	ChainID *big.Int
	Current [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateCommittee is a free log retrieval operation binding the contract event 0xc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd908.
//
// Solidity: event UpdateCommittee(uint256 chainID, bytes32 current)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterUpdateCommittee(opts *bind.FilterOpts) (*LagrangeCommitteeUpdateCommitteeIterator, error) {

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "UpdateCommittee")
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeUpdateCommitteeIterator{contract: _LagrangeCommittee.contract, event: "UpdateCommittee", logs: logs, sub: sub}, nil
}

// WatchUpdateCommittee is a free log subscription operation binding the contract event 0xc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd908.
//
// Solidity: event UpdateCommittee(uint256 chainID, bytes32 current)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchUpdateCommittee(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeUpdateCommittee) (event.Subscription, error) {

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "UpdateCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeUpdateCommittee)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "UpdateCommittee", log); err != nil {
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

// ParseUpdateCommittee is a log parse operation binding the contract event 0xc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd908.
//
// Solidity: event UpdateCommittee(uint256 chainID, bytes32 current)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseUpdateCommittee(log types.Log) (*LagrangeCommitteeUpdateCommittee, error) {
	event := new(LagrangeCommitteeUpdateCommittee)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "UpdateCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
