// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package committee

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

// ILagrangeCommitteeCommitteeData is an auto generated low-level Go binding around an user-defined struct.
type ILagrangeCommitteeCommitteeData struct {
	Root             *big.Int
	Height           *big.Int
	TotalVotingPower *big.Int
}

// ILagrangeCommitteeCommitteeLeaf is an auto generated low-level Go binding around an user-defined struct.
type ILagrangeCommitteeCommitteeLeaf struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}

// CommitteeMetaData contains all meta data concerning the Committee contract.
var CommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeService\",\"name\":\"_service\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"}],\"name\":\"InitCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"current\",\"type\":\"bytes32\"}],\"name\":\"UpdateCommittee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNT_CREATION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUTHORISE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_CURRENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeLeaves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Committees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotingPower\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERMEZ_NETWORK_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addedAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotingPower\",\"type\":\"uint256\"}],\"internalType\":\"structILagrangeCommittee.CommitteeData\",\"name\":\"currentCommittee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nextRoot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structILagrangeCommittee.CommitteeLeaf\",\"name\":\"cleaf\",\"type\":\"tuple\"}],\"name\":\"getLeafHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getNext1CommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getServeUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon4Elements\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"isUpdatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removedAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"service\",\"outputs\":[{\"internalType\":\"contractILagrangeService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"name\":\"setSlashed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updatedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200230738038062002307833981016040819052620000349162000114565b6001600160a01b0381166080526200004b62000052565b5062000146565b600054610100900460ff1615620000bf5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000112576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012757600080fd5b81516001600160a01b03811681146200013f57600080fd5b9392505050565b6080516121976200017060003960008181610674015281816108840152610d1701526121976000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80638da5cb5b11610125578063cd174cde116100ad578063e62f6b921161007c578063e62f6b9214610696578063f1f2fcab14610221578063f2fde38b146106bd578063f5425bd5146106d0578063f8c8765e146106d857600080fd5b8063cd174cde1461060f578063cf7aa21114610622578063d4b96c441461065c578063d598d4c91461066f57600080fd5b8063b40cfbe4116100f4578063b40cfbe41461058f578063c0050642146105a2578063c2b84126146105b5578063c364091e146105d5578063c473af33146105e857600080fd5b80638da5cb5b146105225780639743c7b7146105335780639e4e731814610555578063adacd9921461057c57600080fd5b80633d22485b116101a857806355c1c2bc1161017757806355c1c2bc146104d75780636241171f146104ea578063715018a6146104ff5780637d99c8641461050757806382ab890a1461050f57600080fd5b80633d22485b146103fc57806344a5c4bf1461040f57806344f5b6b41461045657806353701cf41461049d57600080fd5b80631fb63e4d116101ef5780631fb63e4d146102c5578063244cef9c1461030f5780632b8465041461033a5780633408e4701461034d5780633644e5151461035357600080fd5b806304622c2e146102215780631300aff01461025b57806313e7c9d8146102825780631492d9af146102a5575b600080fd5b6102487fbe287413178bfeddef8d9753ad4be825ae998706a6dabff23978b59dccaea0ad81565b6040519081526020015b60405180910390f35b6102487fff946cf82975b1a2b6e6d28c9a76a4b8d7a1fd0592b785cb92771933310f9ee781565b610295610290366004611c15565b6106eb565b6040516102529493929190611c84565b6102486102b3366004611cbb565b60716020526000908152604090205481565b6102f46102d3366004611cbb565b606a6020526000908152604090208054600182015460029092015490919083565b60408051938452602084019290925290820152606001610252565b61032261031d366004611cd4565b6107aa565b6040516001600160a01b039091168152602001610252565b610248610348366004611cd4565b6107e2565b46610248565b61024860007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7fbe287413178bfeddef8d9753ad4be825ae998706a6dabff23978b59dccaea0ad7fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646604080516020810195909552840192909252606083015260808201523060a082015260c00160405160208183030381529060405280519060200120905090565b61032261040a366004611cd4565b610813565b61044161041d366004611c15565b6001600160a01b031660009081526070602052604090206002015463ffffffff1690565b60405163ffffffff9091168152602001610252565b61048d610464366004611c15565b6001600160a01b0316600090815260706020526040902060020154640100000000900460ff1690565b6040519015158152602001610252565b6102f46104ab366004611cd4565b606b60209081526000928352604080842090915290825290208054600182015460029092015490919083565b6102486104e5366004611d99565b61082f565b6104fd6104f8366004611e37565b610879565b005b6104fd610933565b610248600181565b6104fd61051d366004611cbb565b610947565b6033546001600160a01b0316610322565b610546610541366004611cd4565b610a23565b60405161025293929190611e7c565b6102487fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc681565b61024861058a366004611cbb565b610ae3565b6104fd61059d366004611ea3565b610c62565b61048d6105b0366004611cd4565b610c85565b6102486105c3366004611cbb565b606c6020526000908152604090205481565b6102486105e3366004611cd4565b610ccd565b6102487f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6104fd61061d366004611ecf565b610d0c565b610635610630366004611cd4565b610e19565b60408051835181526020808501519082015292810151908301526060820152608001610252565b61024861066a366004611cd4565b610eac565b6103227f000000000000000000000000000000000000000000000000000000000000000081565b6102487fafd642c6a37a2e6887dc4ad5142f84197828a904e53d3204ecb1100329231eaa81565b6104fd6106cb366004611c15565b610ec8565b610248600081565b6104fd6106e6366004611f4e565b610f41565b6070602052600090815260409020805460018201805491929161070d90611fa2565b80601f016020809104026020016040519081016040528092919081815260200182805461073990611fa2565b80156107865780601f1061075b57610100808354040283529160200191610786565b820191906000526020600020905b81548152906001019060200180831161076957829003601f168201915b5050506002909301549192505063ffffffff81169060ff6401000000009091041684565b606960205281600052604060002081815481106107c657600080fd5b6000918252602090912001546001600160a01b03169150829050565b606f60205281600052604060002081815481106107fe57600080fd5b90600052602060002001600091509150505481565b606860205281600052604060002081815481106107c657600080fd5b6000610873604051806060016040528084600001516001600160a01b031681526020018460200151815260200184604001518051906020012060001c815250611094565b92915050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108ca5760405162461bcd60e51b81526004016108c190611fdd565b60405180910390fd5b6001600160a01b03909216600081815260706020908152604080832060020180549615156401000000000264ff000000001990971696909617909555928152606983529283208054600181018255908452919092200180546001600160a01b0319169091179055565b61093b611106565b6109456000611160565b565b60006109538243610ccd565b905061095f8183610c85565b6109c55760405162461bcd60e51b815260206004820152603160248201527f426c6f636b206e756d626572206973207072696f7220746f20636f6d6d69747460448201527032b290333932b2bd32903bb4b73237bb9760791b60648201526084016108c1565b6000828152607160205260409020548111610a155760405162461bcd60e51b815260206004820152601060248201526f20b63932b0b23c903ab83230ba32b21760811b60448201526064016108c1565b610a1f82826111b2565b5050565b606e6020908152600092835260408084209091529082529020805460018201546002830180546001600160a01b03909316939192610a6090611fa2565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8c90611fa2565b8015610ad95780601f10610aae57610100808354040283529160200191610ad9565b820191906000526020600020905b815481529060010190602001808311610abc57829003601f168201915b5050505050905083565b6000818152606f6020526040812054610b1557610873604051806040016040528060008152602001600081525061142d565b6000828152606f602052604090205460011415610b5d576000828152606f602052604081208054909190610b4b57610b4b612021565b90600052602060002001549050919050565b60025b6000838152606f6020526040902054811015610b8857610b8160028261204d565b9050610b60565b6000838152606f6020526040812080548290610ba657610ba6612021565b600091825260209091200154905060015b82811015610c5a576000858152606f6020526040902054811015610c2757610c206040518060400160405280848152602001606f60008981526020019081526020016000208481548110610c0d57610c0d612021565b906000526020600020015481525061142d565b9150610c48565b610c456040518060400160405280848152602001600081525061142d565b91505b80610c528161206c565b915050610bb7565b509392505050565b610c6a611106565b610c7583838361145e565b610c808360006111b2565b505050565b6000818152606a60205260408120600101548190610ca39085612087565b6000848152606a6020526040902060020154909150610cc2818361209f565b431195945050505050565b6000828152606a60205260408120805460019091015480610cee838661209f565b610cf891906120b6565b610d03906001612087565b95945050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d545760405162461bcd60e51b81526004016108c190611fdd565b60008481526068602090815260408083208054600180820183559185528385200180546001600160a01b0319166001600160a01b038b16908117909155825160808101845287815280850189815263ffffffff8816828601526060820187905291865260708552929094208251815593518051929493610dda9392850192910190611a8d565b5060408201516002909101805460609093015115156401000000000264ffffffffff1990931663ffffffff909216919091179190911790555050505050565b610e3d60405180606001604052806000815260200160008152602001600081525090565b600080610e4a8585610ccd565b90506000610e5d866105e3876001612087565b6000968752606b60209081526040808920948952848252808920815160608101835281548152600182015481850152600290910154818301529289529390529190952054909590945092505050565b606d60205281600052604060002081815481106107fe57600080fd5b610ed0611106565b6001600160a01b038116610f355760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016108c1565b610f3e81611160565b50565b600054610100900460ff1615808015610f615750600054600160ff909116105b80610f7b5750303b158015610f7b575060005460ff166001145b610fde5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016108c1565b6000805460ff191660011790558015611001576000805461ff0019166101001790555b606580546001600160a01b038681166001600160a01b03199283161790925560668054868416908316179055606780549285169290911691909117905561104785611160565b801561108d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6066546040516304b98e1d60e31b81526000916001600160a01b0316906325cc70e8906110c59085906004016120d8565b602060405180830381865afa1580156110e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108739190612109565b6033546001600160a01b031633146109455760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108c1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60005b6000838152606860205260409020548110156113105760008381526068602052604081208054839081106111eb576111eb612021565b60009182526020808320909101546001600160a01b03168083526070825260408084208151608081019092528054825260018101805493965091939092908401919061123690611fa2565b80601f016020809104026020016040519081016040528092919081815260200182805461126290611fa2565b80156112af5780601f10611284576101008083540402835291602001916112af565b820191906000526020600020905b81548152906001019060200180831161129257829003601f168201915b50505091835250506002919091015463ffffffff811660208084019190915264010000000090910460ff1615156040909201919091528151908201519192506112fb91879185916115e9565b505080806113089061206c565b9150506111b5565b5060005b60008381526069602052604090205481101561137c576000838152606960205260409020805461136a9185918490811061135057611350612021565b6000918252602090912001546001600160a01b031661175d565b806113748161206c565b915050611314565b50611387828261198c565b6000828152607160209081526040808320849055606890915281206113ab91611b11565b60008281526069602052604081206113c291611b11565b6000828152606b602052604081207fc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd90891849190611400600186612087565b81526020808201929092526040908101600020548151938452918301919091520160405180910390a15050565b6065546040516314d2f97b60e11b81526000916001600160a01b0316906329a5f2f6906110c5908590600401612122565b6000838152606a6020526040902054156114ca5760405162461bcd60e51b815260206004820152602760248201527f436f6d6d69747465652068617320616c7265616479206265656e20696e69746960448201526630b634bd32b21760c91b60648201526084016108c1565b604080516060808201835243825260208083018681528385018681526000898152606a84528681209551865591516001808701919091559051600295860155855193840186528184528383018281528487018381528a8452606b85528784208480528552878420955186559051918501919091555192909301919091558251828152808201808552878452606d90925292909120915161156b929190611b2f565b506000838152606c602090815260408083208390558051838152808301808352878552606f909352922091516115a2929190611b2f565b5060408051848152602081018490529081018290527fd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc49060600160405180910390a1505050565b6000848152606a602052604090205461166a5760405162461bcd60e51b815260206004820152603760248201527f4120636f6d6d697474656520666f72207468697320636861696e20494420686160448201527f73206e6f74206265656e20696e697469616c697a65642e00000000000000000060648201526084016108c1565b604080516060810182526001600160a01b03851681526020810184905290810182905260006116988261082f565b6000878152606e60209081526040808320848452825291829020855181546001600160a01b0319166001600160a01b0390911617815585820151600182015591850151805193945085936116f29260028501920190611a8d565b5050506000868152606d60209081526040808320805460018101825590845282842001849055888352606c909152812080549161172e8361206c565b90915550506000958652606f602090815260408720805460018101825590885296209095019490945550505050565b60005b6000838152606d6020526040902054811015610c80576000838152606d6020526040812080548390811061179657611796612021565b6000918252602080832090910154868352606d909152604082208054919350906117c29060019061209f565b815481106117d2576117d2612021565b6000918252602080832090910154878352606e825260408084208685529092529120549091506001600160a01b039081169085161415611977576000858152606e602090815260408083208484529091528082208483529120815481546001600160a01b0319166001600160a01b0390911617815560018083015490820155600280830180549183019161186590611fa2565b611870929190611b69565b50506040805160608101825260008082526020808301828152845180830186528381528486019081528b8452606e8352858420888552835294909220835181546001600160a01b0319166001600160a01b03909116178155915160018301559251805192945090926118ea92600285019290910190611a8d565b5050506000858152606d6020526040902080546119099060019061209f565b8154811061191957611919612021565b9060005260206000200154606d6000878152602001908152602001600020848154811061194857611948612021565b6000918252602080832090910192909255868152606c909152604081208054916119718361214a565b91905055505b505080806119849061206c565b915050611760565b600061199783610ae3565b905060006119a6600184612087565b6000858152606c6020908152604080832054606b8352818420858552909252909120600181019190915583905590506119de84611a02565b6000948552606b602090815260408087209387529290529320600201929092555050565b600080805b6000848152606c6020526040902054811015611a86576000848152606e60209081526040808320606d9092528220805491929184908110611a4a57611a4a612021565b906000526020600020015481526020019081526020016000206001015482611a729190612087565b915080611a7e8161206c565b915050611a07565b5092915050565b828054611a9990611fa2565b90600052602060002090601f016020900481019282611abb5760008555611b01565b82601f10611ad457805160ff1916838001178555611b01565b82800160010185558215611b01579182015b82811115611b01578251825591602001919060010190611ae6565b50611b0d929150611be4565b5090565b5080546000825590600052602060002090810190610f3e9190611be4565b828054828255906000526020600020908101928215611b015791602002820182811115611b01578251825591602001919060010190611ae6565b828054611b7590611fa2565b90600052602060002090601f016020900481019282611b975760008555611b01565b82601f10611ba85780548555611b01565b82800160010185558215611b0157600052602060002091601f016020900482015b82811115611b01578254825591600101919060010190611bc9565b5b80821115611b0d5760008155600101611be5565b80356001600160a01b0381168114611c1057600080fd5b919050565b600060208284031215611c2757600080fd5b611c3082611bf9565b9392505050565b6000815180845260005b81811015611c5d57602081850181015186830182015201611c41565b81811115611c6f576000602083870101525b50601f01601f19169290920160200192915050565b848152608060208201526000611c9d6080830186611c37565b63ffffffff9490941660408301525090151560609091015292915050565b600060208284031215611ccd57600080fd5b5035919050565b60008060408385031215611ce757600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611d1d57600080fd5b813567ffffffffffffffff80821115611d3857611d38611cf6565b604051601f8301601f19908116603f01168101908282118183101715611d6057611d60611cf6565b81604052838152866020858801011115611d7957600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215611dab57600080fd5b813567ffffffffffffffff80821115611dc357600080fd5b9083019060608286031215611dd757600080fd5b604051606081018181108382111715611df257611df2611cf6565b604052611dfe83611bf9565b815260208301356020820152604083013582811115611e1c57600080fd5b611e2887828601611d0c565b60408301525095945050505050565b600080600060608486031215611e4c57600080fd5b611e5584611bf9565b92506020840135915060408401358015158114611e7157600080fd5b809150509250925092565b60018060a01b0384168152826020820152606060408201526000610d036060830184611c37565b600080600060608486031215611eb857600080fd5b505081359360208301359350604090920135919050565b600080600080600060a08688031215611ee757600080fd5b611ef086611bf9565b945060208601359350604086013567ffffffffffffffff811115611f1357600080fd5b611f1f88828901611d0c565b93505060608601359150608086013563ffffffff81168114611f4057600080fd5b809150509295509295909350565b60008060008060808587031215611f6457600080fd5b611f6d85611bf9565b9350611f7b60208601611bf9565b9250611f8960408601611bf9565b9150611f9760608601611bf9565b905092959194509250565b600181811c90821680611fb657607f821691505b60208210811415611fd757634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526024908201527f4f6e6c7920736572766963652063616e2063616c6c20746869732066756e637460408201526334b7b71760e11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561206757612067612037565b500290565b600060001982141561208057612080612037565b5060010190565b6000821982111561209a5761209a612037565b500190565b6000828210156120b1576120b1612037565b500390565b6000826120d357634e487b7160e01b600052601260045260246000fd5b500490565b60608101818360005b60038110156121005781518352602092830192909101906001016120e1565b50505092915050565b60006020828403121561211b57600080fd5b5051919050565b60408101818360005b600281101561210057815183526020928301929091019060010161212b565b60008161215957612159612037565b50600019019056fea2646970667358221220cccdaa9437122ca496d81a505ad43da56308e92d531ec54bf5c3cd0ede36f56864736f6c634300080c0033",
}

// CommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitteeMetaData.ABI instead.
var CommitteeABI = CommitteeMetaData.ABI

// CommitteeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitteeMetaData.Bin instead.
var CommitteeBin = CommitteeMetaData.Bin

// DeployCommittee deploys a new Ethereum contract, binding an instance of Committee to it.
func DeployCommittee(auth *bind.TransactOpts, backend bind.ContractBackend, _service common.Address) (common.Address, *types.Transaction, *Committee, error) {
	parsed, err := CommitteeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitteeBin), backend, _service)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Committee{CommitteeCaller: CommitteeCaller{contract: contract}, CommitteeTransactor: CommitteeTransactor{contract: contract}, CommitteeFilterer: CommitteeFilterer{contract: contract}}, nil
}

// Committee is an auto generated Go binding around an Ethereum contract.
type Committee struct {
	CommitteeCaller     // Read-only binding to the contract
	CommitteeTransactor // Write-only binding to the contract
	CommitteeFilterer   // Log filterer for contract events
}

// CommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitteeSession struct {
	Contract     *Committee        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitteeCallerSession struct {
	Contract *CommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitteeTransactorSession struct {
	Contract     *CommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitteeRaw struct {
	Contract *Committee // Generic contract binding to access the raw methods on
}

// CommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitteeCallerRaw struct {
	Contract *CommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// CommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitteeTransactorRaw struct {
	Contract *CommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommittee creates a new instance of Committee, bound to a specific deployed contract.
func NewCommittee(address common.Address, backend bind.ContractBackend) (*Committee, error) {
	contract, err := bindCommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Committee{CommitteeCaller: CommitteeCaller{contract: contract}, CommitteeTransactor: CommitteeTransactor{contract: contract}, CommitteeFilterer: CommitteeFilterer{contract: contract}}, nil
}

// NewCommitteeCaller creates a new read-only instance of Committee, bound to a specific deployed contract.
func NewCommitteeCaller(address common.Address, caller bind.ContractCaller) (*CommitteeCaller, error) {
	contract, err := bindCommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeCaller{contract: contract}, nil
}

// NewCommitteeTransactor creates a new write-only instance of Committee, bound to a specific deployed contract.
func NewCommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitteeTransactor, error) {
	contract, err := bindCommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeTransactor{contract: contract}, nil
}

// NewCommitteeFilterer creates a new log filterer instance of Committee, bound to a specific deployed contract.
func NewCommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitteeFilterer, error) {
	contract, err := bindCommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitteeFilterer{contract: contract}, nil
}

// bindCommittee binds a generic wrapper to an already deployed contract.
func bindCommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Committee *CommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Committee.Contract.CommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Committee *CommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Committee.Contract.CommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Committee *CommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Committee.Contract.CommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Committee *CommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Committee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Committee *CommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Committee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Committee *CommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Committee.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Committee *CommitteeCaller) ACCOUNTCREATIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "ACCOUNT_CREATION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Committee *CommitteeSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _Committee.Contract.ACCOUNTCREATIONHASH(&_Committee.CallOpts)
}

// ACCOUNTCREATIONHASH is a free data retrieval call binding the contract method 0x1300aff0.
//
// Solidity: function ACCOUNT_CREATION_HASH() view returns(bytes32)
func (_Committee *CommitteeCallerSession) ACCOUNTCREATIONHASH() ([32]byte, error) {
	return _Committee.Contract.ACCOUNTCREATIONHASH(&_Committee.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Committee *CommitteeCaller) AUTHORISETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "AUTHORISE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Committee *CommitteeSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _Committee.Contract.AUTHORISETYPEHASH(&_Committee.CallOpts)
}

// AUTHORISETYPEHASH is a free data retrieval call binding the contract method 0xe62f6b92.
//
// Solidity: function AUTHORISE_TYPEHASH() view returns(bytes32)
func (_Committee *CommitteeCallerSession) AUTHORISETYPEHASH() ([32]byte, error) {
	return _Committee.Contract.AUTHORISETYPEHASH(&_Committee.CallOpts)
}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_Committee *CommitteeCaller) COMMITTEECURRENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "COMMITTEE_CURRENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_Committee *CommitteeSession) COMMITTEECURRENT() (*big.Int, error) {
	return _Committee.Contract.COMMITTEECURRENT(&_Committee.CallOpts)
}

// COMMITTEECURRENT is a free data retrieval call binding the contract method 0xf5425bd5.
//
// Solidity: function COMMITTEE_CURRENT() view returns(uint256)
func (_Committee *CommitteeCallerSession) COMMITTEECURRENT() (*big.Int, error) {
	return _Committee.Contract.COMMITTEECURRENT(&_Committee.CallOpts)
}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_Committee *CommitteeCaller) COMMITTEENEXT1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "COMMITTEE_NEXT_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_Committee *CommitteeSession) COMMITTEENEXT1() (*big.Int, error) {
	return _Committee.Contract.COMMITTEENEXT1(&_Committee.CallOpts)
}

// COMMITTEENEXT1 is a free data retrieval call binding the contract method 0x7d99c864.
//
// Solidity: function COMMITTEE_NEXT_1() view returns(uint256)
func (_Committee *CommitteeCallerSession) COMMITTEENEXT1() (*big.Int, error) {
	return _Committee.Contract.COMMITTEENEXT1(&_Committee.CallOpts)
}

// CommitteeLeaves is a free data retrieval call binding the contract method 0x2b846504.
//
// Solidity: function CommitteeLeaves(uint256 , uint256 ) view returns(uint256)
func (_Committee *CommitteeCaller) CommitteeLeaves(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "CommitteeLeaves", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeLeaves is a free data retrieval call binding the contract method 0x2b846504.
//
// Solidity: function CommitteeLeaves(uint256 , uint256 ) view returns(uint256)
func (_Committee *CommitteeSession) CommitteeLeaves(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Committee.Contract.CommitteeLeaves(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeLeaves is a free data retrieval call binding the contract method 0x2b846504.
//
// Solidity: function CommitteeLeaves(uint256 , uint256 ) view returns(uint256)
func (_Committee *CommitteeCallerSession) CommitteeLeaves(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Committee.Contract.CommitteeLeaves(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_Committee *CommitteeCaller) CommitteeMap(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "CommitteeMap", arg0, arg1)

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
func (_Committee *CommitteeSession) CommitteeMap(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	return _Committee.Contract.CommitteeMap(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeMap is a free data retrieval call binding the contract method 0x9743c7b7.
//
// Solidity: function CommitteeMap(uint256 , uint256 ) view returns(address addr, uint256 stake, bytes blsPubKey)
func (_Committee *CommitteeCallerSession) CommitteeMap(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr      common.Address
	Stake     *big.Int
	BlsPubKey []byte
}, error) {
	return _Committee.Contract.CommitteeMap(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_Committee *CommitteeCaller) CommitteeMapKeys(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "CommitteeMapKeys", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_Committee *CommitteeSession) CommitteeMapKeys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Committee.Contract.CommitteeMapKeys(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeMapKeys is a free data retrieval call binding the contract method 0xd4b96c44.
//
// Solidity: function CommitteeMapKeys(uint256 , uint256 ) view returns(uint256)
func (_Committee *CommitteeCallerSession) CommitteeMapKeys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Committee.Contract.CommitteeMapKeys(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_Committee *CommitteeCaller) CommitteeMapLength(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "CommitteeMapLength", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_Committee *CommitteeSession) CommitteeMapLength(arg0 *big.Int) (*big.Int, error) {
	return _Committee.Contract.CommitteeMapLength(&_Committee.CallOpts, arg0)
}

// CommitteeMapLength is a free data retrieval call binding the contract method 0xc2b84126.
//
// Solidity: function CommitteeMapLength(uint256 ) view returns(uint256)
func (_Committee *CommitteeCallerSession) CommitteeMapLength(arg0 *big.Int) (*big.Int, error) {
	return _Committee.Contract.CommitteeMapLength(&_Committee.CallOpts, arg0)
}

// CommitteeParams is a free data retrieval call binding the contract method 0x1fb63e4d.
//
// Solidity: function CommitteeParams(uint256 ) view returns(uint256 startBlock, uint256 duration, uint256 freezeDuration)
func (_Committee *CommitteeCaller) CommitteeParams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock     *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "CommitteeParams", arg0)

	outstruct := new(struct {
		StartBlock     *big.Int
		Duration       *big.Int
		FreezeDuration *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FreezeDuration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CommitteeParams is a free data retrieval call binding the contract method 0x1fb63e4d.
//
// Solidity: function CommitteeParams(uint256 ) view returns(uint256 startBlock, uint256 duration, uint256 freezeDuration)
func (_Committee *CommitteeSession) CommitteeParams(arg0 *big.Int) (struct {
	StartBlock     *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
}, error) {
	return _Committee.Contract.CommitteeParams(&_Committee.CallOpts, arg0)
}

// CommitteeParams is a free data retrieval call binding the contract method 0x1fb63e4d.
//
// Solidity: function CommitteeParams(uint256 ) view returns(uint256 startBlock, uint256 duration, uint256 freezeDuration)
func (_Committee *CommitteeCallerSession) CommitteeParams(arg0 *big.Int) (struct {
	StartBlock     *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
}, error) {
	return _Committee.Contract.CommitteeParams(&_Committee.CallOpts, arg0)
}

// Committees is a free data retrieval call binding the contract method 0x53701cf4.
//
// Solidity: function Committees(uint256 , uint256 ) view returns(uint256 root, uint256 height, uint256 totalVotingPower)
func (_Committee *CommitteeCaller) Committees(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Root             *big.Int
	Height           *big.Int
	TotalVotingPower *big.Int
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "Committees", arg0, arg1)

	outstruct := new(struct {
		Root             *big.Int
		Height           *big.Int
		TotalVotingPower *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Height = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalVotingPower = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Committees is a free data retrieval call binding the contract method 0x53701cf4.
//
// Solidity: function Committees(uint256 , uint256 ) view returns(uint256 root, uint256 height, uint256 totalVotingPower)
func (_Committee *CommitteeSession) Committees(arg0 *big.Int, arg1 *big.Int) (struct {
	Root             *big.Int
	Height           *big.Int
	TotalVotingPower *big.Int
}, error) {
	return _Committee.Contract.Committees(&_Committee.CallOpts, arg0, arg1)
}

// Committees is a free data retrieval call binding the contract method 0x53701cf4.
//
// Solidity: function Committees(uint256 , uint256 ) view returns(uint256 root, uint256 height, uint256 totalVotingPower)
func (_Committee *CommitteeCallerSession) Committees(arg0 *big.Int, arg1 *big.Int) (struct {
	Root             *big.Int
	Height           *big.Int
	TotalVotingPower *big.Int
}, error) {
	return _Committee.Contract.Committees(&_Committee.CallOpts, arg0, arg1)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Committee *CommitteeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Committee *CommitteeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Committee.Contract.DOMAINSEPARATOR(&_Committee.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator)
func (_Committee *CommitteeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Committee.Contract.DOMAINSEPARATOR(&_Committee.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Committee *CommitteeCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "EIP712DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Committee *CommitteeSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Committee.Contract.EIP712DOMAINHASH(&_Committee.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xc473af33.
//
// Solidity: function EIP712DOMAIN_HASH() view returns(bytes32)
func (_Committee *CommitteeCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Committee.Contract.EIP712DOMAINHASH(&_Committee.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Committee *CommitteeCaller) HERMEZNETWORKHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "HERMEZ_NETWORK_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Committee *CommitteeSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _Committee.Contract.HERMEZNETWORKHASH(&_Committee.CallOpts)
}

// HERMEZNETWORKHASH is a free data retrieval call binding the contract method 0xf1f2fcab.
//
// Solidity: function HERMEZ_NETWORK_HASH() view returns(bytes32)
func (_Committee *CommitteeCallerSession) HERMEZNETWORKHASH() ([32]byte, error) {
	return _Committee.Contract.HERMEZNETWORKHASH(&_Committee.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Committee *CommitteeCaller) NAMEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "NAME_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Committee *CommitteeSession) NAMEHASH() ([32]byte, error) {
	return _Committee.Contract.NAMEHASH(&_Committee.CallOpts)
}

// NAMEHASH is a free data retrieval call binding the contract method 0x04622c2e.
//
// Solidity: function NAME_HASH() view returns(bytes32)
func (_Committee *CommitteeCallerSession) NAMEHASH() ([32]byte, error) {
	return _Committee.Contract.NAMEHASH(&_Committee.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Committee *CommitteeCaller) VERSIONHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "VERSION_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Committee *CommitteeSession) VERSIONHASH() ([32]byte, error) {
	return _Committee.Contract.VERSIONHASH(&_Committee.CallOpts)
}

// VERSIONHASH is a free data retrieval call binding the contract method 0x9e4e7318.
//
// Solidity: function VERSION_HASH() view returns(bytes32)
func (_Committee *CommitteeCallerSession) VERSIONHASH() ([32]byte, error) {
	return _Committee.Contract.VERSIONHASH(&_Committee.CallOpts)
}

// AddedAddrs is a free data retrieval call binding the contract method 0x3d22485b.
//
// Solidity: function addedAddrs(uint256 , uint256 ) view returns(address)
func (_Committee *CommitteeCaller) AddedAddrs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "addedAddrs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddedAddrs is a free data retrieval call binding the contract method 0x3d22485b.
//
// Solidity: function addedAddrs(uint256 , uint256 ) view returns(address)
func (_Committee *CommitteeSession) AddedAddrs(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Committee.Contract.AddedAddrs(&_Committee.CallOpts, arg0, arg1)
}

// AddedAddrs is a free data retrieval call binding the contract method 0x3d22485b.
//
// Solidity: function addedAddrs(uint256 , uint256 ) view returns(address)
func (_Committee *CommitteeCallerSession) AddedAddrs(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Committee.Contract.AddedAddrs(&_Committee.CallOpts, arg0, arg1)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Committee *CommitteeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Committee *CommitteeSession) GetChainId() (*big.Int, error) {
	return _Committee.Contract.GetChainId(&_Committee.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Committee *CommitteeCallerSession) GetChainId() (*big.Int, error) {
	return _Committee.Contract.GetChainId(&_Committee.CallOpts)
}

// GetCommittee is a free data retrieval call binding the contract method 0xcf7aa211.
//
// Solidity: function getCommittee(uint256 chainID, uint256 blockNumber) view returns((uint256,uint256,uint256) currentCommittee, uint256 nextRoot)
func (_Committee *CommitteeCaller) GetCommittee(opts *bind.CallOpts, chainID *big.Int, blockNumber *big.Int) (struct {
	CurrentCommittee ILagrangeCommitteeCommitteeData
	NextRoot         *big.Int
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getCommittee", chainID, blockNumber)

	outstruct := new(struct {
		CurrentCommittee ILagrangeCommitteeCommitteeData
		NextRoot         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentCommittee = *abi.ConvertType(out[0], new(ILagrangeCommitteeCommitteeData)).(*ILagrangeCommitteeCommitteeData)
	outstruct.NextRoot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCommittee is a free data retrieval call binding the contract method 0xcf7aa211.
//
// Solidity: function getCommittee(uint256 chainID, uint256 blockNumber) view returns((uint256,uint256,uint256) currentCommittee, uint256 nextRoot)
func (_Committee *CommitteeSession) GetCommittee(chainID *big.Int, blockNumber *big.Int) (struct {
	CurrentCommittee ILagrangeCommitteeCommitteeData
	NextRoot         *big.Int
}, error) {
	return _Committee.Contract.GetCommittee(&_Committee.CallOpts, chainID, blockNumber)
}

// GetCommittee is a free data retrieval call binding the contract method 0xcf7aa211.
//
// Solidity: function getCommittee(uint256 chainID, uint256 blockNumber) view returns((uint256,uint256,uint256) currentCommittee, uint256 nextRoot)
func (_Committee *CommitteeCallerSession) GetCommittee(chainID *big.Int, blockNumber *big.Int) (struct {
	CurrentCommittee ILagrangeCommitteeCommitteeData
	NextRoot         *big.Int
}, error) {
	return _Committee.Contract.GetCommittee(&_Committee.CallOpts, chainID, blockNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0xc364091e.
//
// Solidity: function getEpochNumber(uint256 chainID, uint256 blockNumber) view returns(uint256)
func (_Committee *CommitteeCaller) GetEpochNumber(opts *bind.CallOpts, chainID *big.Int, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getEpochNumber", chainID, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0xc364091e.
//
// Solidity: function getEpochNumber(uint256 chainID, uint256 blockNumber) view returns(uint256)
func (_Committee *CommitteeSession) GetEpochNumber(chainID *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _Committee.Contract.GetEpochNumber(&_Committee.CallOpts, chainID, blockNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0xc364091e.
//
// Solidity: function getEpochNumber(uint256 chainID, uint256 blockNumber) view returns(uint256)
func (_Committee *CommitteeCallerSession) GetEpochNumber(chainID *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _Committee.Contract.GetEpochNumber(&_Committee.CallOpts, chainID, blockNumber)
}

// GetLeafHash is a free data retrieval call binding the contract method 0x55c1c2bc.
//
// Solidity: function getLeafHash((address,uint256,bytes) cleaf) view returns(uint256)
func (_Committee *CommitteeCaller) GetLeafHash(opts *bind.CallOpts, cleaf ILagrangeCommitteeCommitteeLeaf) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getLeafHash", cleaf)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLeafHash is a free data retrieval call binding the contract method 0x55c1c2bc.
//
// Solidity: function getLeafHash((address,uint256,bytes) cleaf) view returns(uint256)
func (_Committee *CommitteeSession) GetLeafHash(cleaf ILagrangeCommitteeCommitteeLeaf) (*big.Int, error) {
	return _Committee.Contract.GetLeafHash(&_Committee.CallOpts, cleaf)
}

// GetLeafHash is a free data retrieval call binding the contract method 0x55c1c2bc.
//
// Solidity: function getLeafHash((address,uint256,bytes) cleaf) view returns(uint256)
func (_Committee *CommitteeCallerSession) GetLeafHash(cleaf ILagrangeCommitteeCommitteeLeaf) (*big.Int, error) {
	return _Committee.Contract.GetLeafHash(&_Committee.CallOpts, cleaf)
}

// GetNext1CommitteeRoot is a free data retrieval call binding the contract method 0xadacd992.
//
// Solidity: function getNext1CommitteeRoot(uint256 chainID) view returns(uint256)
func (_Committee *CommitteeCaller) GetNext1CommitteeRoot(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getNext1CommitteeRoot", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNext1CommitteeRoot is a free data retrieval call binding the contract method 0xadacd992.
//
// Solidity: function getNext1CommitteeRoot(uint256 chainID) view returns(uint256)
func (_Committee *CommitteeSession) GetNext1CommitteeRoot(chainID *big.Int) (*big.Int, error) {
	return _Committee.Contract.GetNext1CommitteeRoot(&_Committee.CallOpts, chainID)
}

// GetNext1CommitteeRoot is a free data retrieval call binding the contract method 0xadacd992.
//
// Solidity: function getNext1CommitteeRoot(uint256 chainID) view returns(uint256)
func (_Committee *CommitteeCallerSession) GetNext1CommitteeRoot(chainID *big.Int) (*big.Int, error) {
	return _Committee.Contract.GetNext1CommitteeRoot(&_Committee.CallOpts, chainID)
}

// GetServeUntilBlock is a free data retrieval call binding the contract method 0x44a5c4bf.
//
// Solidity: function getServeUntilBlock(address operator) view returns(uint32)
func (_Committee *CommitteeCaller) GetServeUntilBlock(opts *bind.CallOpts, operator common.Address) (uint32, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getServeUntilBlock", operator)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetServeUntilBlock is a free data retrieval call binding the contract method 0x44a5c4bf.
//
// Solidity: function getServeUntilBlock(address operator) view returns(uint32)
func (_Committee *CommitteeSession) GetServeUntilBlock(operator common.Address) (uint32, error) {
	return _Committee.Contract.GetServeUntilBlock(&_Committee.CallOpts, operator)
}

// GetServeUntilBlock is a free data retrieval call binding the contract method 0x44a5c4bf.
//
// Solidity: function getServeUntilBlock(address operator) view returns(uint32)
func (_Committee *CommitteeCallerSession) GetServeUntilBlock(operator common.Address) (uint32, error) {
	return _Committee.Contract.GetServeUntilBlock(&_Committee.CallOpts, operator)
}

// GetSlashed is a free data retrieval call binding the contract method 0x44f5b6b4.
//
// Solidity: function getSlashed(address operator) view returns(bool)
func (_Committee *CommitteeCaller) GetSlashed(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getSlashed", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSlashed is a free data retrieval call binding the contract method 0x44f5b6b4.
//
// Solidity: function getSlashed(address operator) view returns(bool)
func (_Committee *CommitteeSession) GetSlashed(operator common.Address) (bool, error) {
	return _Committee.Contract.GetSlashed(&_Committee.CallOpts, operator)
}

// GetSlashed is a free data retrieval call binding the contract method 0x44f5b6b4.
//
// Solidity: function getSlashed(address operator) view returns(bool)
func (_Committee *CommitteeCallerSession) GetSlashed(operator common.Address) (bool, error) {
	return _Committee.Contract.GetSlashed(&_Committee.CallOpts, operator)
}

// IsUpdatable is a free data retrieval call binding the contract method 0xc0050642.
//
// Solidity: function isUpdatable(uint256 epochNumber, uint256 chainID) view returns(bool)
func (_Committee *CommitteeCaller) IsUpdatable(opts *bind.CallOpts, epochNumber *big.Int, chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "isUpdatable", epochNumber, chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdatable is a free data retrieval call binding the contract method 0xc0050642.
//
// Solidity: function isUpdatable(uint256 epochNumber, uint256 chainID) view returns(bool)
func (_Committee *CommitteeSession) IsUpdatable(epochNumber *big.Int, chainID *big.Int) (bool, error) {
	return _Committee.Contract.IsUpdatable(&_Committee.CallOpts, epochNumber, chainID)
}

// IsUpdatable is a free data retrieval call binding the contract method 0xc0050642.
//
// Solidity: function isUpdatable(uint256 epochNumber, uint256 chainID) view returns(bool)
func (_Committee *CommitteeCallerSession) IsUpdatable(epochNumber *big.Int, chainID *big.Int) (bool, error) {
	return _Committee.Contract.IsUpdatable(&_Committee.CallOpts, epochNumber, chainID)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, bytes blsPubKey, uint32 serveUntilBlock, bool slashed)
func (_Committee *CommitteeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount          *big.Int
	BlsPubKey       []byte
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		Amount          *big.Int
		BlsPubKey       []byte
		ServeUntilBlock uint32
		Slashed         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlsPubKey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ServeUntilBlock = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Slashed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, bytes blsPubKey, uint32 serveUntilBlock, bool slashed)
func (_Committee *CommitteeSession) Operators(arg0 common.Address) (struct {
	Amount          *big.Int
	BlsPubKey       []byte
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	return _Committee.Contract.Operators(&_Committee.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 amount, bytes blsPubKey, uint32 serveUntilBlock, bool slashed)
func (_Committee *CommitteeCallerSession) Operators(arg0 common.Address) (struct {
	Amount          *big.Int
	BlsPubKey       []byte
	ServeUntilBlock uint32
	Slashed         bool
}, error) {
	return _Committee.Contract.Operators(&_Committee.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Committee *CommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Committee *CommitteeSession) Owner() (common.Address, error) {
	return _Committee.Contract.Owner(&_Committee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Committee *CommitteeCallerSession) Owner() (common.Address, error) {
	return _Committee.Contract.Owner(&_Committee.CallOpts)
}

// RemovedAddrs is a free data retrieval call binding the contract method 0x244cef9c.
//
// Solidity: function removedAddrs(uint256 , uint256 ) view returns(address)
func (_Committee *CommitteeCaller) RemovedAddrs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "removedAddrs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemovedAddrs is a free data retrieval call binding the contract method 0x244cef9c.
//
// Solidity: function removedAddrs(uint256 , uint256 ) view returns(address)
func (_Committee *CommitteeSession) RemovedAddrs(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Committee.Contract.RemovedAddrs(&_Committee.CallOpts, arg0, arg1)
}

// RemovedAddrs is a free data retrieval call binding the contract method 0x244cef9c.
//
// Solidity: function removedAddrs(uint256 , uint256 ) view returns(address)
func (_Committee *CommitteeCallerSession) RemovedAddrs(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Committee.Contract.RemovedAddrs(&_Committee.CallOpts, arg0, arg1)
}

// Service is a free data retrieval call binding the contract method 0xd598d4c9.
//
// Solidity: function service() view returns(address)
func (_Committee *CommitteeCaller) Service(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "service")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Service is a free data retrieval call binding the contract method 0xd598d4c9.
//
// Solidity: function service() view returns(address)
func (_Committee *CommitteeSession) Service() (common.Address, error) {
	return _Committee.Contract.Service(&_Committee.CallOpts)
}

// Service is a free data retrieval call binding the contract method 0xd598d4c9.
//
// Solidity: function service() view returns(address)
func (_Committee *CommitteeCallerSession) Service() (common.Address, error) {
	return _Committee.Contract.Service(&_Committee.CallOpts)
}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x1492d9af.
//
// Solidity: function updatedEpoch(uint256 ) view returns(uint256)
func (_Committee *CommitteeCaller) UpdatedEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "updatedEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x1492d9af.
//
// Solidity: function updatedEpoch(uint256 ) view returns(uint256)
func (_Committee *CommitteeSession) UpdatedEpoch(arg0 *big.Int) (*big.Int, error) {
	return _Committee.Contract.UpdatedEpoch(&_Committee.CallOpts, arg0)
}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x1492d9af.
//
// Solidity: function updatedEpoch(uint256 ) view returns(uint256)
func (_Committee *CommitteeCallerSession) UpdatedEpoch(arg0 *big.Int) (*big.Int, error) {
	return _Committee.Contract.UpdatedEpoch(&_Committee.CallOpts, arg0)
}

// AddOperator is a paid mutator transaction binding the contract method 0xcd174cde.
//
// Solidity: function addOperator(address operator, uint256 chainID, bytes blsPubKey, uint256 stake, uint32 serveUntilBlock) returns()
func (_Committee *CommitteeTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address, chainID *big.Int, blsPubKey []byte, stake *big.Int, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "addOperator", operator, chainID, blsPubKey, stake, serveUntilBlock)
}

// AddOperator is a paid mutator transaction binding the contract method 0xcd174cde.
//
// Solidity: function addOperator(address operator, uint256 chainID, bytes blsPubKey, uint256 stake, uint32 serveUntilBlock) returns()
func (_Committee *CommitteeSession) AddOperator(operator common.Address, chainID *big.Int, blsPubKey []byte, stake *big.Int, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Committee.Contract.AddOperator(&_Committee.TransactOpts, operator, chainID, blsPubKey, stake, serveUntilBlock)
}

// AddOperator is a paid mutator transaction binding the contract method 0xcd174cde.
//
// Solidity: function addOperator(address operator, uint256 chainID, bytes blsPubKey, uint256 stake, uint32 serveUntilBlock) returns()
func (_Committee *CommitteeTransactorSession) AddOperator(operator common.Address, chainID *big.Int, blsPubKey []byte, stake *big.Int, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Committee.Contract.AddOperator(&_Committee.TransactOpts, operator, chainID, blsPubKey, stake, serveUntilBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Committee *CommitteeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "initialize", initialOwner, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Committee *CommitteeSession) Initialize(initialOwner common.Address, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Committee.Contract.Initialize(&_Committee.TransactOpts, initialOwner, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Committee *CommitteeTransactorSession) Initialize(initialOwner common.Address, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Committee.Contract.Initialize(&_Committee.TransactOpts, initialOwner, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb40cfbe4.
//
// Solidity: function registerChain(uint256 chainID, uint256 epochPeriod, uint256 freezeDuration) returns()
func (_Committee *CommitteeTransactor) RegisterChain(opts *bind.TransactOpts, chainID *big.Int, epochPeriod *big.Int, freezeDuration *big.Int) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "registerChain", chainID, epochPeriod, freezeDuration)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb40cfbe4.
//
// Solidity: function registerChain(uint256 chainID, uint256 epochPeriod, uint256 freezeDuration) returns()
func (_Committee *CommitteeSession) RegisterChain(chainID *big.Int, epochPeriod *big.Int, freezeDuration *big.Int) (*types.Transaction, error) {
	return _Committee.Contract.RegisterChain(&_Committee.TransactOpts, chainID, epochPeriod, freezeDuration)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb40cfbe4.
//
// Solidity: function registerChain(uint256 chainID, uint256 epochPeriod, uint256 freezeDuration) returns()
func (_Committee *CommitteeTransactorSession) RegisterChain(chainID *big.Int, epochPeriod *big.Int, freezeDuration *big.Int) (*types.Transaction, error) {
	return _Committee.Contract.RegisterChain(&_Committee.TransactOpts, chainID, epochPeriod, freezeDuration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Committee *CommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Committee *CommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Committee.Contract.RenounceOwnership(&_Committee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Committee *CommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Committee.Contract.RenounceOwnership(&_Committee.TransactOpts)
}

// SetSlashed is a paid mutator transaction binding the contract method 0x6241171f.
//
// Solidity: function setSlashed(address operator, uint256 chainID, bool slashed) returns()
func (_Committee *CommitteeTransactor) SetSlashed(opts *bind.TransactOpts, operator common.Address, chainID *big.Int, slashed bool) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "setSlashed", operator, chainID, slashed)
}

// SetSlashed is a paid mutator transaction binding the contract method 0x6241171f.
//
// Solidity: function setSlashed(address operator, uint256 chainID, bool slashed) returns()
func (_Committee *CommitteeSession) SetSlashed(operator common.Address, chainID *big.Int, slashed bool) (*types.Transaction, error) {
	return _Committee.Contract.SetSlashed(&_Committee.TransactOpts, operator, chainID, slashed)
}

// SetSlashed is a paid mutator transaction binding the contract method 0x6241171f.
//
// Solidity: function setSlashed(address operator, uint256 chainID, bool slashed) returns()
func (_Committee *CommitteeTransactorSession) SetSlashed(operator common.Address, chainID *big.Int, slashed bool) (*types.Transaction, error) {
	return _Committee.Contract.SetSlashed(&_Committee.TransactOpts, operator, chainID, slashed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Committee *CommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Committee *CommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Committee.Contract.TransferOwnership(&_Committee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Committee *CommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Committee.Contract.TransferOwnership(&_Committee.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 chainID) returns()
func (_Committee *CommitteeTransactor) Update(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "update", chainID)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 chainID) returns()
func (_Committee *CommitteeSession) Update(chainID *big.Int) (*types.Transaction, error) {
	return _Committee.Contract.Update(&_Committee.TransactOpts, chainID)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 chainID) returns()
func (_Committee *CommitteeTransactorSession) Update(chainID *big.Int) (*types.Transaction, error) {
	return _Committee.Contract.Update(&_Committee.TransactOpts, chainID)
}

// CommitteeInitCommitteeIterator is returned from FilterInitCommittee and is used to iterate over the raw logs and unpacked data for InitCommittee events raised by the Committee contract.
type CommitteeInitCommitteeIterator struct {
	Event *CommitteeInitCommittee // Event containing the contract specifics and raw log

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
func (it *CommitteeInitCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeInitCommittee)
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
		it.Event = new(CommitteeInitCommittee)
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
func (it *CommitteeInitCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitteeInitCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitteeInitCommittee represents a InitCommittee event raised by the Committee contract.
type CommitteeInitCommittee struct {
	ChainID        *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitCommittee is a free log retrieval operation binding the contract event 0xd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc4.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration)
func (_Committee *CommitteeFilterer) FilterInitCommittee(opts *bind.FilterOpts) (*CommitteeInitCommitteeIterator, error) {

	logs, sub, err := _Committee.contract.FilterLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return &CommitteeInitCommitteeIterator{contract: _Committee.contract, event: "InitCommittee", logs: logs, sub: sub}, nil
}

// WatchInitCommittee is a free log subscription operation binding the contract event 0xd07f5f940c054019c6c46eed514ed7d35417d411b6f94c49ada89240be6c7fc4.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration)
func (_Committee *CommitteeFilterer) WatchInitCommittee(opts *bind.WatchOpts, sink chan<- *CommitteeInitCommittee) (event.Subscription, error) {

	logs, sub, err := _Committee.contract.WatchLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitteeInitCommittee)
				if err := _Committee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
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
func (_Committee *CommitteeFilterer) ParseInitCommittee(log types.Log) (*CommitteeInitCommittee, error) {
	event := new(CommitteeInitCommittee)
	if err := _Committee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitteeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Committee contract.
type CommitteeInitializedIterator struct {
	Event *CommitteeInitialized // Event containing the contract specifics and raw log

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
func (it *CommitteeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeInitialized)
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
		it.Event = new(CommitteeInitialized)
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
func (it *CommitteeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitteeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitteeInitialized represents a Initialized event raised by the Committee contract.
type CommitteeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Committee *CommitteeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CommitteeInitializedIterator, error) {

	logs, sub, err := _Committee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CommitteeInitializedIterator{contract: _Committee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Committee *CommitteeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CommitteeInitialized) (event.Subscription, error) {

	logs, sub, err := _Committee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitteeInitialized)
				if err := _Committee.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Committee *CommitteeFilterer) ParseInitialized(log types.Log) (*CommitteeInitialized, error) {
	event := new(CommitteeInitialized)
	if err := _Committee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Committee contract.
type CommitteeOwnershipTransferredIterator struct {
	Event *CommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeOwnershipTransferred)
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
		it.Event = new(CommitteeOwnershipTransferred)
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
func (it *CommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the Committee contract.
type CommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Committee *CommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Committee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeOwnershipTransferredIterator{contract: _Committee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Committee *CommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Committee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitteeOwnershipTransferred)
				if err := _Committee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Committee *CommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*CommitteeOwnershipTransferred, error) {
	event := new(CommitteeOwnershipTransferred)
	if err := _Committee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitteeUpdateCommitteeIterator is returned from FilterUpdateCommittee and is used to iterate over the raw logs and unpacked data for UpdateCommittee events raised by the Committee contract.
type CommitteeUpdateCommitteeIterator struct {
	Event *CommitteeUpdateCommittee // Event containing the contract specifics and raw log

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
func (it *CommitteeUpdateCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeUpdateCommittee)
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
		it.Event = new(CommitteeUpdateCommittee)
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
func (it *CommitteeUpdateCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitteeUpdateCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitteeUpdateCommittee represents a UpdateCommittee event raised by the Committee contract.
type CommitteeUpdateCommittee struct {
	ChainID *big.Int
	Current [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateCommittee is a free log retrieval operation binding the contract event 0xc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd908.
//
// Solidity: event UpdateCommittee(uint256 chainID, bytes32 current)
func (_Committee *CommitteeFilterer) FilterUpdateCommittee(opts *bind.FilterOpts) (*CommitteeUpdateCommitteeIterator, error) {

	logs, sub, err := _Committee.contract.FilterLogs(opts, "UpdateCommittee")
	if err != nil {
		return nil, err
	}
	return &CommitteeUpdateCommitteeIterator{contract: _Committee.contract, event: "UpdateCommittee", logs: logs, sub: sub}, nil
}

// WatchUpdateCommittee is a free log subscription operation binding the contract event 0xc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd908.
//
// Solidity: event UpdateCommittee(uint256 chainID, bytes32 current)
func (_Committee *CommitteeFilterer) WatchUpdateCommittee(opts *bind.WatchOpts, sink chan<- *CommitteeUpdateCommittee) (event.Subscription, error) {

	logs, sub, err := _Committee.contract.WatchLogs(opts, "UpdateCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitteeUpdateCommittee)
				if err := _Committee.contract.UnpackLog(event, "UpdateCommittee", log); err != nil {
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
func (_Committee *CommitteeFilterer) ParseUpdateCommittee(log types.Log) (*CommitteeUpdateCommittee, error) {
	event := new(CommitteeUpdateCommittee)
	if err := _Committee.contract.UnpackLog(event, "UpdateCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
