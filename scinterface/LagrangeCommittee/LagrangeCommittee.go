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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon4Elements\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"InitCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"next1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"next2\",\"type\":\"uint256\"}],\"name\":\"RotateCommittee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNT_CREATION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUTHORISE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_CURRENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"COMMITTEE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"COMMITTEE_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeLeaves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeMapLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERMEZ_NETWORK_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_blsPubKey\",\"type\":\"bytes\"}],\"name\":\"committeeAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epoch2committee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epoch2height\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epoch2startblock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getCommitteeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getCommitteeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getCommitteeStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structLagrangeCommittee.CommitteeLeaf\",\"name\":\"cleaf\",\"type\":\"tuple\"}],\"name\":\"getLeafHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getNext1CommitteeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getNextCommitteeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"hash2Elements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"initCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeCommitteeAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"rotateCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"comparisonNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"verifyBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106102275760003560e01c8063759f5b5a11610130578063c473af33116100b8578063e62f6b921161007c578063e62f6b921461064b578063f1f2fcab1461022c578063f2fde38b14610672578063f5425bd514610685578063fe095d5c1461068d57600080fd5b8063c473af33146105b3578063d4b96c44146105da578063d8b5364c146105ed578063e00aab1a14610618578063e04392f11461064357600080fd5b80639743c7b7116100ff5780639743c7b71461050c5780639e4e73181461052e578063adacd99214610555578063ba1efb5814610568578063c2b841261461059357600080fd5b8063759f5b5a146104b15780637d99c864146104c45780638da5cb5b146104cc578063946a55cf146104ec57600080fd5b80633644e515116101b35780634d008f4d116101825780634d008f4d1461043357806355c1c2bc14610453578063620390221461046657806364a83b4b14610489578063715018a6146104a957600080fd5b80633644e5151461033957806340d810b4146103e257806340dda085146103f55780634204e0551461040857600080fd5b80632b846504116101fa5780632b846504146102d85780632e253dcc146102eb5780632f374006146102fe578063333edc3f1461031e5780633408e4701461033357600080fd5b806304622c2e1461022c5780630c01dc12146102665780631300aff0146102915780631c04e7f1146102b8575b600080fd5b6102537fbe287413178bfeddef8d9753ad4be825ae998706a6dabff23978b59dccaea0ad81565b6040519081526020015b60405180910390f35b61025361027436600461161b565b607060209081526000928352604080842090915290825290205481565b6102537fff946cf82975b1a2b6e6d28c9a76a4b8d7a1fd0592b785cb92771933310f9ee781565b6102536102c636600461163d565b60696020526000908152604090205481565b6102536102e636600461161b565b6106a0565b6102536102f936600461161b565b6106d1565b61025361030c36600461163d565b606f6020526000908152604090205481565b61033161032c36600461163d565b610703565b005b46610253565b61025360007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7fbe287413178bfeddef8d9753ad4be825ae998706a6dabff23978b59dccaea0ad7fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646604080516020810195909552840192909252606083015260808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6103316103f036600461161b565b610932565b6103316104033660046116f9565b610a44565b61025361041636600461161b565b607160209081526000928352604080842090915290825290205481565b61025361044136600461163d565b60009081526068602052604090205490565b610253610461366004611765565b610b69565b610479610474366004611803565b610b9e565b604051901515815260200161025d565b61025361049736600461163d565b60009081526069602052604090205490565b610331610c28565b6103316104bf36600461163d565b610c3c565b610253600181565b6104d4610eac565b6040516001600160a01b03909116815260200161025d565b6102536104fa36600461163d565b60686020526000908152604090205481565b61051f61051a36600461161b565b610ec5565b60405161025d939291906118a7565b6102537fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc681565b61025361056336600461163d565b610f85565b61025361057636600461161b565b606e60209081526000928352604080842090915290825290205481565b6102536105a136600461163d565b606a6020526000908152604090205481565b6102537f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6102536105e836600461161b565b611270565b6102536105fb36600461161b565b607260209081526000928352604080842090915290825290205481565b61025361062636600461161b565b600091825260706020908152604080842092845291905290205490565b610253600281565b6102537fafd642c6a37a2e6887dc4ad5142f84197828a904e53d3204ecb1100329231eaa81565b6103316106803660046118ce565b61128c565b610253600081565b61025361069b36600461161b565b611305565b606d60205281600052604060002081815481106106bc57600080fd5b90600052602060002001600091509150505481565b6000828152607060205260408120816106eb8460016118ff565b81526020810191909152604001600020549392505050565b3360005b6000838152606b602052604090205481101561092d576000838152606b6020526040812080548390811061073d5761073d611917565b6000918252602080832090910154868352606b909152604082208054919350906107699060019061192d565b8154811061077957610779611917565b6000918252602080832090910154878352606c825260408084208685529092529120549091506001600160a01b039081169085161415610918576000858152606c602090815260408083208484529091528082208483529120815481546001600160a01b0319166001600160a01b0390911617815560018083015490820155600280830180549183019161080c90611944565b6108179291906114cd565b509050506108366000806040518060200160405280600081525061133a565b6000868152606c60209081526040808320858452825291829020835181546001600160a01b0319166001600160a01b0390911617815583820151600182015591830151805161088b9260028501920190611558565b5050506000858152606b6020526040902080546108aa9060019061192d565b815481106108ba576108ba611917565b9060005260206000200154606b600087815260200190815260200160002084815481106108e9576108e9611917565b6000918252602080832090910192909255868152606a909152604081208054916109128361197f565b91905055505b5050808061092590611996565b915050610707565b505050565b61093a61137d565b600082815260686020526040902054156109ab5760405162461bcd60e51b815260206004820152602760248201527f436f6d6d69747465652068617320616c7265616479206265656e20696e69746960448201526630b634bd32b21760c91b60648201526084015b60405180910390fd5b6000828152606860209081526040808320439055606990915290208190556109d48160026119b1565b6109de90436118ff565b600083815260716020908152604080832060028452825280832093909355848252606f81528282209190915581518481529081018390527f6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b910160405180910390a15050565b6040805160008082526020808301808552878352606d90915292902090513392610a6e92916115cc565b506000848152606e60209081526040808320600284529091528120819055610a9782858561133a565b90506000610aa482610b69565b6000878152606c60209081526040808320848452825291829020855181546001600160a01b0319166001600160a01b039091161781558582015160018201559185015180519394508593610afe9260028501920190611558565b5050506000868152606b60209081526040808320805460018101825590845282842001849055888352606a9091528120805491610b3a83611996565b90915550506000958652606d602090815260408720805460018101825590885296209095019490945550505050565b6000610b9882600001516001600160a01b031661069b846020015185604001518051906020012060001c611305565b92915050565b604051633101c81160e11b815260009073__$5f60e9b106869b445e169c91071e8d06eb$__90636203902290610bde9088908890889088906004016119d0565b602060405180830381865af4158015610bfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1f91906119fc565b95945050505050565b610c3061137d565b610c3a60006113dc565b565b600081815260696020908152604080832054606890925290912054610c6191906118ff565b4311610cd55760405162461bcd60e51b815260206004820152603b60248201527f426c6f636b206e756d62657220646f6573206e6f742065786365656420656e6460448201527f20626c6f636b206f662063757272656e7420636f6d6d6974746565000000000060648201526084016109a2565b610cde8161142e565b6000818152606860209081526040808320439055606e82528083206001845282528083208054848052828520556002808552828520549182905585855260708452828520606f90945291842054909391610d37916118ff565b8152602080820192909252604090810160009081209390935583835260718252808320606f909252822054909190610d71906001906118ff565b8152602080820192909252604090810160009081205484825260718452828220606f90945291812054919291610da9906002906118ff565b81526020808201929092526040908101600090812093909355838352606a82528083205460728352818420606f90935290832054909290610dec906002906118ff565b81526020808201929092526040908101600090812093909355838352606e8252808320600284528252808320839055838352606f9091528120805491610e3183611996565b90915550506000818152606e6020908152604080832083805282528083205460018452818420546002855293829020548251868152938401919091529082019290925260608101919091527f0256e53795397d2ae94765793b7c5e056ad393405fa2c82f68b70411f635c3199060800160405180910390a150565b6000610ec06033546001600160a01b031690565b905090565b606c6020908152600092835260408084209091529082529020805460018201546002830180546001600160a01b03909316939192610f0290611944565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2e90611944565b8015610f7b5780601f10610f5057610100808354040283529160200191610f7b565b820191906000526020600020905b815481529060010190602001808311610f5e57829003601f168201915b5050505050905083565b6000818152606d6020526040812054610fa357610b98600080611305565b6000828152606d602052604090205460011415610feb576000828152606d602052604081208054909190610fd957610fd9611917565b90600052602060002001549050919050565b6000828152606d602052604081205490611006600283611a1e565b67ffffffffffffffff81111561101e5761101e611656565b604051908082528060200260200182016040528015611047578160200160208202803683370190505b5090506000805b61105960018561192d565b811015611115576000868152606d602052604090206110dc9061107c83856118ff565b8154811061108c5761108c611917565b6000918252602080832090910154898352606d90915260409091206110b184866118ff565b6110bc9060016118ff565b815481106110cc576110cc611917565b9060005260206000200154611305565b836110e8600284611a1e565b815181106110f8576110f8611917565b602090810291909101015261110e6002826118ff565b905061104e565b50611121600284611a1e565b92505b821561123f576000611137600285611a1e565b67ffffffffffffffff81111561114f5761114f611656565b604051908082528060200260200182016040528015611178578160200160208202803683370190505b50905060005b61118960018661192d565b811015611224576111eb8461119e83866118ff565b815181106111ae576111ae611917565b60200260200101518583866111c391906118ff565b6111ce9060016118ff565b815181106111de576111de611917565b6020026020010151611305565b826111f7600284611a1e565b8151811061120757611207611917565b602090810291909101015261121d6002826118ff565b905061117e565b5091506000905081611237600285611a1e565b935050611124565b816001835161124e919061192d565b8151811061125e5761125e611917565b60200260200101519350505050919050565b606b60205281600052604060002081815481106106bc57600080fd5b61129461137d565b6001600160a01b0381166112f95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109a2565b611302816113dc565b50565b600061132460405180604001604052808581526020018481525061145b565b9392505050565b6001600160a01b03163b151590565b6040805160608082018352600080835260208301529181019190915250604080516060810182526001600160a01b03909416845260208401929092529082015290565b33611386610eac565b6001600160a01b031614610c3a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109a2565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061143982610f85565b6000928352606e60209081526040808520600286529091529092209190915550565b6065546040516314d2f97b60e11b81526000916001600160a01b0316906329a5f2f69061148c908590600401611a40565b602060405180830381865afa1580156114a9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b989190611a71565b8280546114d990611944565b90600052602060002090601f0160209004810192826114fb5760008555611548565b82601f1061150c5780548555611548565b8280016001018555821561154857600052602060002091601f016020900482015b8281111561154857825482559160010191906001019061152d565b50611554929150611606565b5090565b82805461156490611944565b90600052602060002090601f0160209004810192826115865760008555611548565b82601f1061159f57805160ff1916838001178555611548565b82800160010185558215611548579182015b828111156115485782518255916020019190600101906115b1565b82805482825590600052602060002090810192821561154857916020028201828111156115485782518255916020019190600101906115b1565b5b808211156115545760008155600101611607565b6000806040838503121561162e57600080fd5b50508035926020909101359150565b60006020828403121561164f57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261167d57600080fd5b813567ffffffffffffffff8082111561169857611698611656565b604051601f8301601f19908116603f011681019082821181831017156116c0576116c0611656565b816040528381528660208588010111156116d957600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561170e57600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561173357600080fd5b61173f8682870161166c565b9150509250925092565b80356001600160a01b038116811461176057600080fd5b919050565b60006020828403121561177757600080fd5b813567ffffffffffffffff8082111561178f57600080fd5b90830190606082860312156117a357600080fd5b6040516060810181811083821117156117be576117be611656565b6040526117ca83611749565b8152602083013560208201526040830135828111156117e857600080fd5b6117f48782860161166c565b60408301525095945050505050565b6000806000806080858703121561181957600080fd5b84359350602085013567ffffffffffffffff81111561183757600080fd5b6118438782880161166c565b949794965050505060408301359260600135919050565b6000815180845260005b8181101561188057602081850181015186830182015201611864565b81811115611892576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b0384168152826020820152606060408201526000610c1f606083018461185a565b6000602082840312156118e057600080fd5b61132482611749565b634e487b7160e01b600052601160045260246000fd5b60008219821115611912576119126118e9565b500190565b634e487b7160e01b600052603260045260246000fd5b60008282101561193f5761193f6118e9565b500390565b600181811c9082168061195857607f821691505b6020821081141561197957634e487b7160e01b600052602260045260246000fd5b50919050565b60008161198e5761198e6118e9565b506000190190565b60006000198214156119aa576119aa6118e9565b5060010190565b60008160001904831182151516156119cb576119cb6118e9565b500290565b8481526080602082015260006119e9608083018661185a565b6040830194909452506060015292915050565b600060208284031215611a0e57600080fd5b8151801515811461132457600080fd5b600082611a3b57634e487b7160e01b600052601260045260246000fd5b500490565b60408101818360005b6002811015611a68578151835260209283019290910190600101611a49565b50505092915050565b600060208284031215611a8357600080fd5b505191905056fea2646970667358221220232792ac846978bf23ee3ea21654f1c34faf33868a45ee6b5d933ae1b3f0d8ca64736f6c634300080c0033",
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

// COMMITTEEDURATION is a free data retrieval call binding the contract method 0x1c04e7f1.
//
// Solidity: function COMMITTEE_DURATION(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) COMMITTEEDURATION(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "COMMITTEE_DURATION", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEEDURATION is a free data retrieval call binding the contract method 0x1c04e7f1.
//
// Solidity: function COMMITTEE_DURATION(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) COMMITTEEDURATION(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEEDURATION(&_LagrangeCommittee.CallOpts, arg0)
}

// COMMITTEEDURATION is a free data retrieval call binding the contract method 0x1c04e7f1.
//
// Solidity: function COMMITTEE_DURATION(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) COMMITTEEDURATION(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEEDURATION(&_LagrangeCommittee.CallOpts, arg0)
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

// COMMITTEESTART is a free data retrieval call binding the contract method 0x946a55cf.
//
// Solidity: function COMMITTEE_START(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) COMMITTEESTART(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "COMMITTEE_START", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMITTEESTART is a free data retrieval call binding the contract method 0x946a55cf.
//
// Solidity: function COMMITTEE_START(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) COMMITTEESTART(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEESTART(&_LagrangeCommittee.CallOpts, arg0)
}

// COMMITTEESTART is a free data retrieval call binding the contract method 0x946a55cf.
//
// Solidity: function COMMITTEE_START(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) COMMITTEESTART(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.COMMITTEESTART(&_LagrangeCommittee.CallOpts, arg0)
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

// CommitteeRoot is a free data retrieval call binding the contract method 0xba1efb58.
//
// Solidity: function CommitteeRoot(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeRoot(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "CommitteeRoot", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitteeRoot is a free data retrieval call binding the contract method 0xba1efb58.
//
// Solidity: function CommitteeRoot(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeRoot(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeRoot(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeRoot is a free data retrieval call binding the contract method 0xba1efb58.
//
// Solidity: function CommitteeRoot(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeRoot(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.CommitteeRoot(&_LagrangeCommittee.CallOpts, arg0, arg1)
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

// EpochNumber is a free data retrieval call binding the contract method 0x2f374006.
//
// Solidity: function EpochNumber(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) EpochNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "EpochNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochNumber is a free data retrieval call binding the contract method 0x2f374006.
//
// Solidity: function EpochNumber(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) EpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.EpochNumber(&_LagrangeCommittee.CallOpts, arg0)
}

// EpochNumber is a free data retrieval call binding the contract method 0x2f374006.
//
// Solidity: function EpochNumber(uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) EpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.EpochNumber(&_LagrangeCommittee.CallOpts, arg0)
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

// Epoch2committee is a free data retrieval call binding the contract method 0x0c01dc12.
//
// Solidity: function epoch2committee(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Epoch2committee(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "epoch2committee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch2committee is a free data retrieval call binding the contract method 0x0c01dc12.
//
// Solidity: function epoch2committee(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) Epoch2committee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Epoch2committee(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// Epoch2committee is a free data retrieval call binding the contract method 0x0c01dc12.
//
// Solidity: function epoch2committee(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Epoch2committee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Epoch2committee(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// Epoch2height is a free data retrieval call binding the contract method 0xd8b5364c.
//
// Solidity: function epoch2height(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Epoch2height(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "epoch2height", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch2height is a free data retrieval call binding the contract method 0xd8b5364c.
//
// Solidity: function epoch2height(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) Epoch2height(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Epoch2height(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// Epoch2height is a free data retrieval call binding the contract method 0xd8b5364c.
//
// Solidity: function epoch2height(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Epoch2height(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Epoch2height(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// Epoch2startblock is a free data retrieval call binding the contract method 0x4204e055.
//
// Solidity: function epoch2startblock(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Epoch2startblock(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "epoch2startblock", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch2startblock is a free data retrieval call binding the contract method 0x4204e055.
//
// Solidity: function epoch2startblock(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) Epoch2startblock(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Epoch2startblock(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// Epoch2startblock is a free data retrieval call binding the contract method 0x4204e055.
//
// Solidity: function epoch2startblock(uint256 , uint256 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Epoch2startblock(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.Epoch2startblock(&_LagrangeCommittee.CallOpts, arg0, arg1)
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

// GetCommitteeDuration is a free data retrieval call binding the contract method 0x64a83b4b.
//
// Solidity: function getCommitteeDuration(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetCommitteeDuration(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getCommitteeDuration", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitteeDuration is a free data retrieval call binding the contract method 0x64a83b4b.
//
// Solidity: function getCommitteeDuration(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetCommitteeDuration(chainID *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetCommitteeDuration(&_LagrangeCommittee.CallOpts, chainID)
}

// GetCommitteeDuration is a free data retrieval call binding the contract method 0x64a83b4b.
//
// Solidity: function getCommitteeDuration(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetCommitteeDuration(chainID *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetCommitteeDuration(&_LagrangeCommittee.CallOpts, chainID)
}

// GetCommitteeRoot is a free data retrieval call binding the contract method 0xe00aab1a.
//
// Solidity: function getCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetCommitteeRoot(opts *bind.CallOpts, chainID *big.Int, _epoch *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getCommitteeRoot", chainID, _epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitteeRoot is a free data retrieval call binding the contract method 0xe00aab1a.
//
// Solidity: function getCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetCommitteeRoot(chainID *big.Int, _epoch *big.Int) ([32]byte, error) {
	return _LagrangeCommittee.Contract.GetCommitteeRoot(&_LagrangeCommittee.CallOpts, chainID, _epoch)
}

// GetCommitteeRoot is a free data retrieval call binding the contract method 0xe00aab1a.
//
// Solidity: function getCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetCommitteeRoot(chainID *big.Int, _epoch *big.Int) ([32]byte, error) {
	return _LagrangeCommittee.Contract.GetCommitteeRoot(&_LagrangeCommittee.CallOpts, chainID, _epoch)
}

// GetCommitteeStart is a free data retrieval call binding the contract method 0x4d008f4d.
//
// Solidity: function getCommitteeStart(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetCommitteeStart(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getCommitteeStart", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitteeStart is a free data retrieval call binding the contract method 0x4d008f4d.
//
// Solidity: function getCommitteeStart(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetCommitteeStart(chainID *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetCommitteeStart(&_LagrangeCommittee.CallOpts, chainID)
}

// GetCommitteeStart is a free data retrieval call binding the contract method 0x4d008f4d.
//
// Solidity: function getCommitteeStart(uint256 chainID) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetCommitteeStart(chainID *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetCommitteeStart(&_LagrangeCommittee.CallOpts, chainID)
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

// GetNextCommitteeRoot is a free data retrieval call binding the contract method 0x2e253dcc.
//
// Solidity: function getNextCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetNextCommitteeRoot(opts *bind.CallOpts, chainID *big.Int, _epoch *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getNextCommitteeRoot", chainID, _epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextCommitteeRoot is a free data retrieval call binding the contract method 0x2e253dcc.
//
// Solidity: function getNextCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetNextCommitteeRoot(chainID *big.Int, _epoch *big.Int) ([32]byte, error) {
	return _LagrangeCommittee.Contract.GetNextCommitteeRoot(&_LagrangeCommittee.CallOpts, chainID, _epoch)
}

// GetNextCommitteeRoot is a free data retrieval call binding the contract method 0x2e253dcc.
//
// Solidity: function getNextCommitteeRoot(uint256 chainID, uint256 _epoch) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetNextCommitteeRoot(chainID *big.Int, _epoch *big.Int) ([32]byte, error) {
	return _LagrangeCommittee.Contract.GetNextCommitteeRoot(&_LagrangeCommittee.CallOpts, chainID, _epoch)
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

// CommitteeAdd is a paid mutator transaction binding the contract method 0x40dda085.
//
// Solidity: function committeeAdd(uint256 chainID, uint256 stake, bytes _blsPubKey) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) CommitteeAdd(opts *bind.TransactOpts, chainID *big.Int, stake *big.Int, _blsPubKey []byte) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "committeeAdd", chainID, stake, _blsPubKey)
}

// CommitteeAdd is a paid mutator transaction binding the contract method 0x40dda085.
//
// Solidity: function committeeAdd(uint256 chainID, uint256 stake, bytes _blsPubKey) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeAdd(chainID *big.Int, stake *big.Int, _blsPubKey []byte) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.CommitteeAdd(&_LagrangeCommittee.TransactOpts, chainID, stake, _blsPubKey)
}

// CommitteeAdd is a paid mutator transaction binding the contract method 0x40dda085.
//
// Solidity: function committeeAdd(uint256 chainID, uint256 stake, bytes _blsPubKey) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) CommitteeAdd(chainID *big.Int, stake *big.Int, _blsPubKey []byte) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.CommitteeAdd(&_LagrangeCommittee.TransactOpts, chainID, stake, _blsPubKey)
}

// InitCommittee is a paid mutator transaction binding the contract method 0x40d810b4.
//
// Solidity: function initCommittee(uint256 _chainID, uint256 _duration) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) InitCommittee(opts *bind.TransactOpts, _chainID *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "initCommittee", _chainID, _duration)
}

// InitCommittee is a paid mutator transaction binding the contract method 0x40d810b4.
//
// Solidity: function initCommittee(uint256 _chainID, uint256 _duration) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) InitCommittee(_chainID *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.InitCommittee(&_LagrangeCommittee.TransactOpts, _chainID, _duration)
}

// InitCommittee is a paid mutator transaction binding the contract method 0x40d810b4.
//
// Solidity: function initCommittee(uint256 _chainID, uint256 _duration) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) InitCommittee(_chainID *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.InitCommittee(&_LagrangeCommittee.TransactOpts, _chainID, _duration)
}

// RemoveCommitteeAddr is a paid mutator transaction binding the contract method 0x333edc3f.
//
// Solidity: function removeCommitteeAddr(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RemoveCommitteeAddr(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "removeCommitteeAddr", chainID)
}

// RemoveCommitteeAddr is a paid mutator transaction binding the contract method 0x333edc3f.
//
// Solidity: function removeCommitteeAddr(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RemoveCommitteeAddr(chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RemoveCommitteeAddr(&_LagrangeCommittee.TransactOpts, chainID)
}

// RemoveCommitteeAddr is a paid mutator transaction binding the contract method 0x333edc3f.
//
// Solidity: function removeCommitteeAddr(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RemoveCommitteeAddr(chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RemoveCommitteeAddr(&_LagrangeCommittee.TransactOpts, chainID)
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

// RotateCommittee is a paid mutator transaction binding the contract method 0x759f5b5a.
//
// Solidity: function rotateCommittee(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RotateCommittee(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "rotateCommittee", chainID)
}

// RotateCommittee is a paid mutator transaction binding the contract method 0x759f5b5a.
//
// Solidity: function rotateCommittee(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RotateCommittee(chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RotateCommittee(&_LagrangeCommittee.TransactOpts, chainID)
}

// RotateCommittee is a paid mutator transaction binding the contract method 0x759f5b5a.
//
// Solidity: function rotateCommittee(uint256 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RotateCommittee(chainID *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RotateCommittee(&_LagrangeCommittee.TransactOpts, chainID)
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
	ChainID  *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitCommittee is a free log retrieval operation binding the contract event 0x6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterInitCommittee(opts *bind.FilterOpts) (*LagrangeCommitteeInitCommitteeIterator, error) {

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeInitCommitteeIterator{contract: _LagrangeCommittee.contract, event: "InitCommittee", logs: logs, sub: sub}, nil
}

// WatchInitCommittee is a free log subscription operation binding the contract event 0x6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration)
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

// ParseInitCommittee is a log parse operation binding the contract event 0x6daa941c9959a81de0793d0665491a251ab5993b3868c26ab2f8c2a0c644ac0b.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration)
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

// LagrangeCommitteeRotateCommitteeIterator is returned from FilterRotateCommittee and is used to iterate over the raw logs and unpacked data for RotateCommittee events raised by the LagrangeCommittee contract.
type LagrangeCommitteeRotateCommitteeIterator struct {
	Event *LagrangeCommitteeRotateCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeRotateCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeRotateCommittee)
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
		it.Event = new(LagrangeCommitteeRotateCommittee)
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
func (it *LagrangeCommitteeRotateCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeRotateCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeRotateCommittee represents a RotateCommittee event raised by the LagrangeCommittee contract.
type LagrangeCommitteeRotateCommittee struct {
	ChainID *big.Int
	Current *big.Int
	Next1   *big.Int
	Next2   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRotateCommittee is a free log retrieval operation binding the contract event 0x0256e53795397d2ae94765793b7c5e056ad393405fa2c82f68b70411f635c319.
//
// Solidity: event RotateCommittee(uint256 chainID, uint256 current, uint256 next1, uint256 next2)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterRotateCommittee(opts *bind.FilterOpts) (*LagrangeCommitteeRotateCommitteeIterator, error) {

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "RotateCommittee")
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeRotateCommitteeIterator{contract: _LagrangeCommittee.contract, event: "RotateCommittee", logs: logs, sub: sub}, nil
}

// WatchRotateCommittee is a free log subscription operation binding the contract event 0x0256e53795397d2ae94765793b7c5e056ad393405fa2c82f68b70411f635c319.
//
// Solidity: event RotateCommittee(uint256 chainID, uint256 current, uint256 next1, uint256 next2)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchRotateCommittee(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeRotateCommittee) (event.Subscription, error) {

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "RotateCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeRotateCommittee)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "RotateCommittee", log); err != nil {
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

// ParseRotateCommittee is a log parse operation binding the contract event 0x0256e53795397d2ae94765793b7c5e056ad393405fa2c82f68b70411f635c319.
//
// Solidity: event RotateCommittee(uint256 chainID, uint256 current, uint256 next1, uint256 next2)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseRotateCommittee(log types.Log) (*LagrangeCommitteeRotateCommittee, error) {
	event := new(LagrangeCommitteeRotateCommittee)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "RotateCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
