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
	Root             [32]byte
	LeafCount        *big.Int
	TotalVotingPower *big.Int
}

// ILagrangeCommitteeUnsubscribedParam is an auto generated low-level Go binding around an user-defined struct.
type ILagrangeCommitteeUnsubscribedParam struct {
	ChainID     uint32
	BlockNumber *big.Int
}

// CommitteeMetaData contains all meta data concerning the Committee contract.
var CommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeService\",\"name\":\"_service\",\"type\":\"address\"},{\"internalType\":\"contractIVoteWeigher\",\"name\":\"_voteWeigher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"InitCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"current\",\"type\":\"bytes32\"}],\"name\":\"UpdateCommittee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COMMITTEE_CURRENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMITTEE_NEXT_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INNER_NODE_PREFIX\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEAF_NODE_PREFIX\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"blsPubKey\",\"type\":\"uint256[2]\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committeeAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"committeeHeights\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"committeeLeavesMap\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committeeNodes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"committeeParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotingPower\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getBlsPubKey\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotingPower\",\"type\":\"uint256\"}],\"internalType\":\"structILagrangeCommittee.CommitteeData\",\"name\":\"currentCommittee\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"nextRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"}],\"name\":\"getOperatorStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structILagrangeCommittee.UnsubscribedParam[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"getOperatorVotingPower\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isUnregisterable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"isUpdatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"quorunNumber\",\"type\":\"uint8\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"service\",\"outputs\":[{\"internalType\":\"contractILagrangeService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"subscribeChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"totalVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribeChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"updateOperatorAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"updatedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteWeigher\",\"outputs\":[{\"internalType\":\"contractIVoteWeigher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101f05760003560e01c806385ab9a7a1161010f578063da16ce83116100a2578063f1d3353e11610071578063f1d3353e1461060c578063f2fde38b1461061f578063f5425bd514610632578063fd39105a1461063a57600080fd5b8063da16ce8314610530578063def9e7d514610576578063e377dcad146105b0578063ef030673146105e557600080fd5b8063a63490a2116100de578063a63490a21461048e578063bf988ab6146104e3578063c4d66de8146104f6578063d598d4c91461050957600080fd5b806385ab9a7a146104065780638da5cb5b146104295780638dfccad81461044e578063999c3b4a1461046e57600080fd5b80635af1a88f116101875780637285645511610156578063728564551461037d57806378d81d08146103d8578063798c7244146103eb5780637d99c864146103fe57600080fd5b80635af1a88f146102f55780636b11c38e14610357578063715018a61461036a578063727cb30f1461037257600080fd5b80633bc72805116101c35780633bc728051461029c5780633bc9c733146102af5780633d6a2679146102c25780634db6f74a146102d557600080fd5b80630e9f564b146101f557806319a74c5f1461020a57806334ae75071461023957806338fa59d114610278575b600080fd5b6102086102033660046120d6565b61065b565b005b61021d610218366004612109565b61081b565b6040805192151583526020830191909152015b60405180910390f35b61026a61024736600461213c565b606860209081526000938452604080852082529284528284209052825290205481565b604051908152602001610230565b610283600160f81b81565b6040516001600160f81b03199091168152602001610230565b6102086102aa366004612178565b6108d6565b6102086102bd3660046121a2565b6109ab565b61021d6102d03660046121e8565b6109d0565b61026a6102e33660046121e8565b606e6020526000908152604090205481565b61033f6103033660046120d6565b6001600160a01b0382166000908152606d6020908152604080832063ffffffff851684526003019091529020546001600160601b031692915050565b6040516001600160601b039091168152602001610230565b6102086103653660046120d6565b610a6e565b610208610dc9565b610283600160f91b81565b6103b561038b3660046121e8565b60656020526000908152604090208054600182015460028301546003909301549192909160ff1684565b6040805194855260208501939093529183015260ff166060820152608001610230565b61026a6103e6366004612178565b610ddd565b6102086103f93660046120d6565b610e23565b61026a600181565b610419610414366004612178565b61102e565b6040519015158152602001610230565b6033546001600160a01b03165b6040516001600160a01b039091168152602001610230565b61026a61045c3660046121e8565b60676020526000908152604090205481565b61048161047c366004612109565b611093565b6040516102309190612203565b6104c861049c366004612178565b606660209081526000928352604080842090915290825290208054600182015460029092015490919083565b60408051938452602084019290925290820152606001610230565b6104366104f1366004612178565b6110e4565b610208610504366004612109565b61111c565b6104367f000000000000000000000000000000000000000000000000000000000000000081565b61056161053e366004612234565b606960209081526000928352604080842090915290825290205463ffffffff1681565b60405163ffffffff9091168152602001610230565b610589610584366004612178565b6112b4565b60408051835181526020808501519082015292810151908301526060820152608001610230565b6105d36105be3660046121e8565b606a6020526000908152604090205460ff1681565b60405160ff9091168152602001610230565b6104367f000000000000000000000000000000000000000000000000000000000000000081565b61020861061a366004612274565b611345565b61020861062d366004612109565b611400565b61026a600081565b61064d610648366004612109565b611479565b604051610230929190612305565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106ac5760405162461bcd60e51b81526004016106a390612366565b60405180910390fd5b6001600160a01b0382166000908152606d6020908152604080832063ffffffff8516845260038101909252909120546001600160601b03166107005760405162461bcd60e51b81526004016106a3906123b3565b60008061070c846109d0565b91509150811561072e5760405162461bcd60e51b81526004016106a3906123f8565b63ffffffff841660009081526003840160209081526040808320546067909252822080546001600160601b0390921692909161076b908490612445565b909155505063ffffffff8481166000818152600386016020908152604080832080546001600160601b03191690558051808201909152928352828101858152600488018054600181810183559185529290932093516002928302909401805463ffffffff1916949095169390931784559151928101929092558401546107f4919060ff1661245c565b60028401805460ff191660ff929092169190911790556108148585611519565b5050505050565b6001600160a01b0381166000908152606d60205260408120600281015482919060ff161561084f5750600093849350915050565b6000805b60048301548110156108c95760008360040182815481106108765761087661247f565b600091825260209182902060408051808201909152600290920201805463ffffffff1682526001015491810182905291508310156108b657806020015192505b50806108c181612495565b915050610853565b5060019590945092505050565b6108e0828261102e565b6109465760405162461bcd60e51b815260206004820152603160248201527f426c6f636b206e756d626572206973207072696f7220746f20636f6d6d69747460448201527032b290333932b2bd32903bb4b73237bb9760791b60648201526084016106a3565b63ffffffff82166000908152606e6020526040902054811161099d5760405162461bcd60e51b815260206004820152601060248201526f20b63932b0b23c903ab83230ba32b21760811b60448201526064016106a3565b6109a7828261181d565b5050565b6109b3611956565b6109bf848484846119b0565b6109ca84600061181d565b50505050565b63ffffffff811660009081526065602052604081206001015481906109fa57506000928392509050565b6000610a068443610ddd565b63ffffffff8516600090815260656020526040812080546001909101549293509091610a3290846124b0565b610a3c91906124cf565b63ffffffff8616600090815260656020526040902060020154909150610a629082612445565b43119590945092505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ab65760405162461bcd60e51b81526004016106a390612366565b6000610ac1826109d0565b5090508015610ae25760405162461bcd60e51b81526004016106a3906123f8565b6001600160a01b0383166000908152606d60205260408120905b6004820154811015610bec576000826004018281548110610b1f57610b1f61247f565b600091825260209182902060408051808201909152600290920201805463ffffffff908116808452600190920154938301939093529092509086161415610bd95760008160200151118015610b78575043816020015110155b15610bd95760405162461bcd60e51b815260206004820152602b60248201527f5468652064656463696174656420636861696e206973207768696c6520756e7360448201526a3ab139b1b934b134b7339760a91b60648201526084016106a3565b5080610be481612495565b915050610afc565b5063ffffffff831660009081526003820160205260409020546001600160601b031615610c6e5760405162461bcd60e51b815260206004820152602a60248201527f5468652064656469636174656420636861696e20697320616c726561647920736044820152693ab139b1b934b132b21760b11b60648201526084016106a3565b63ffffffff831660009081526065602052604090819020600301549051631b48130b60e31b815260ff90911660048201526001600160a01b0385811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063da409858906044016020604051808303816000875af1158015610cfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1e91906124e7565b63ffffffff84166000908152600383016020526040902080546001600160601b0319166001600160601b03929092169190911790556002810154610d669060ff166001612510565b60028201805460ff191660ff9290921691909117905563ffffffff831660009081526003820160209081526040808320546067909252822080546001600160601b03909216929091610db99084906124cf565b909155506109ca90508484611b19565b610dd1611956565b610ddb6000611c5f565b565b63ffffffff82166000908152606560205260408120805460019091015480610e058386612445565b610e0f9190612535565b610e1a9060016124cf565b95945050505050565b6000610e2e826109d0565b5090508015610e4f5760405162461bcd60e51b81526004016106a3906123f8565b6001600160a01b0383166000908152606d6020908152604080832063ffffffff8616845260038101909252909120546001600160601b0316610ea35760405162461bcd60e51b81526004016106a3906123b3565b63ffffffff8316600090815260656020526040808220600301549051631b48130b60e31b815260ff90911660048201526001600160a01b0386811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063da409858906044016020604051808303816000875af1158015610f2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5291906124e7565b63ffffffff851660009081526003840160205260409020549091506001600160601b038083169116146108145763ffffffff841660009081526003830160209081526040808320546067909252822080546001600160601b03909216929091610fbc908490612445565b909155505063ffffffff8416600090815260676020526040812080546001600160601b0384169290610fef9084906124cf565b909155505063ffffffff84166000908152600383016020526040902080546001600160601b0319166001600160601b0383161790556108148585611cb1565b63ffffffff82166000908152606560205260408120805460019091015482919061105890856124b0565b61106291906124cf565b63ffffffff85166000908152606560205260409020600201549091506110888183612445565b431195945050505050565b61109b612035565b6001600160a01b0382166000908152606d6020526040908190208151808301928390529160029082845b8154815260200190600101908083116110c55750505050509050919050565b606b602052816000526040600020818154811061110057600080fd5b6000918252602090912001546001600160a01b03169150829050565b600054610100900460ff161580801561113c5750600054600160ff909116105b806111565750303b158015611156575060005460ff166001145b6111b95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106a3565b6000805460ff1916600117905580156111dc576000805461ff0019166101001790555b60015b60148160ff16116112615761123c606c60006111fc60018561245c565b60ff1660ff16815260200190815260200160002054606c6000600185611222919061245c565b60ff1660ff16815260200190815260200160002054611d1a565b60ff82166000908152606c60205260409020558061125981612557565b9150506111df565b5061126b82611c5f565b80156109a7576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6040805160608101825260008082526020820181905291810182905290806112dc8585610ddd565b905060006112ef866103e68760016124cf565b63ffffffff96909616600090815260666020908152604080832094835284825280832081516060810183528154815260018201548185015260029091015481830152988352939052919091205494959350505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461138d5760405162461bcd60e51b81526004016106a390612366565b6001600160a01b0382166000908152606d602052604090208054156113f45760405162461bcd60e51b815260206004820152601f60248201527f4f70657261746f7220697320616c726561647920726567697374657265642e0060448201526064016106a3565b6109ca81836002612053565b611408611956565b6001600160a01b03811661146d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106a3565b61147681611c5f565b50565b6001600160a01b0381166000908152606d602090815260408083206002810154600482018054845181870281018701909552808552606095939460ff90931693919291839190889084015b828210156115095760008481526020908190206040805180820190915260028502909101805463ffffffff1682526001908101548284015290835290920191016114c4565b5050505090509250925050915091565b63ffffffff80821660008181526069602090815260408083206001600160a01b0388168452825280832054938352606b90915281205491909216919061156190600190612445565b63ffffffff84166000908152606b60205260408120805492935090918390811061158d5761158d61247f565b6000918252602090912001546001600160a01b0316905063ffffffff83168211156116675763ffffffff8481166000818152606860209081526040808320838052825280832087845282528083205494881680845281842095909555928252606b905220805483929081106116045761160461247f565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905563ffffffff8781168352606982526040808420948616845293909152919020805463ffffffff19169185169182179055611667908590611d59565b63ffffffff84166000908152606b6020526040902080548061168b5761168b612577565b60008281526020808220830160001990810180546001600160a01b031916905590920190925563ffffffff86168252606a8152604080832054606883528184208480528352818420868552909252822082905560ff1660015b8160ff168160ff1610156118135782806117015750846001166001145b15611723576001925061171e87611718858461245c565b87611db8565b6117f4565b63ffffffff8716600090815260686020908152604080832060ff85168452909152812081611752600289612535565b815260208101919091526040016000205561176e60028361245c565b60ff168160ff1614156117f45763ffffffff871660009081526068602052604081208161179c60018661245c565b60ff168152602080820192909252604090810160009081208180529092529020556117c860018361245c565b63ffffffff88166000908152606a60205260409020805460ff191660ff92909216919091179055611813565b6117ff600286612535565b94508061180b81612557565b9150506116e4565b5050505050505050565b600061182a6001836124cf565b63ffffffff84166000818152606b602090815260408083205460668352818420868552835281842060010155928252606a9052205490915060ff16156118d65763ffffffff83166000908152606860209081526040808320606a90925282205490919061189c9060019060ff1661245c565b60ff1681526020808201929092526040908101600090812081805283528181205463ffffffff871682526066845282822085835290935220555b63ffffffff8316600081815260676020908152604080832054606683528184208685528084528285206002810192909255858552606e84528285208890559386905292825291548251938452908301527fc6ee71ee195b28e5f3e5f5737bdae699800c460cc899508d730e8cc9eeedd908910160405180910390a1505050565b6033546001600160a01b03163314610ddb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106a3565b63ffffffff841660009081526065602052604090205415611a235760405162461bcd60e51b815260206004820152602760248201527f436f6d6d69747465652068617320616c7265616479206265656e20696e69746960448201526630b634bd32b21760c91b60648201526084016106a3565b6040805160808082018352438252602080830187815283850187815260ff878116606080880182815263ffffffff8e166000818152606589528b81209a518b5596516001808c019190915595516002808c019190915591516003909a01805460ff19169a9095169990991790935588518082018a52858152808701868152818b018781528a8852606689528b88208880528952968b9020915182555194810194909455935192909101919091558551948552918401889052938301869052928201929092527fe4166ce16a6b34a7c665c8a0b0cc0cfac48559a82b8dae718ce478257becfa1a910160405180910390a150505050565b63ffffffff81166000908152606b60209081526040822080546001810182559083529120810180546001600160a01b0319166001600160a01b038516179055611b628383611f99565b63ffffffff8084166000818152606860209081526040808320838052825280832094871680845294825280832095909555918152606982528381206001600160a01b038816825290915291909120805463ffffffff1916821790551580611bf9575063ffffffff82166000908152606a6020526040902054611be99060019060ff1661245c565b60ff166001901b8163ffffffff16145b15611c4a5763ffffffff82166000908152606a6020526040902054611c229060ff166001612510565b63ffffffff83166000908152606a60205260409020805460ff191660ff929092169190911790555b611c5a828263ffffffff16611d59565b505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b63ffffffff80821660009081526069602090815260408083206001600160a01b038716845290915290205416611ce78383611f99565b63ffffffff831660009081526068602090815260408083208380528252808320858452909152902055611c5a8282611d59565b604051600160f91b6020820152602181018390526041810182905260009060610160405160208183030381529060405280519060200120905092915050565b60005b63ffffffff83166000908152606a6020526040902054611d819060019060ff1661245c565b60ff168160ff161015611c5a57611d99838284611db8565b611da4600283612535565b915080611db081612557565b915050611d5c565b6000808260011660011415611e815763ffffffff8516600090815260686020908152604080832060ff88168452909152812090611df6600186612445565b8152602080820192909252604090810160009081205463ffffffff891682526068845282822060ff8916835284528282208783529093522054909250611e4f575060ff83166000908152606c6020526040902054611f33565b5063ffffffff8416600090815260686020908152604080832060ff871684528252808320858452909152902054611f33565b63ffffffff8516600090815260686020908152604080832060ff881680855281845282852088865280855292852054908552925290935090611ec48560016124cf565b8152602081019190915260400160002054611ef2575060ff83166000908152606c6020526040902054611f33565b63ffffffff8516600090815260686020908152604080832060ff88168452909152812090611f218560016124cf565b81526020019081526020016000205490505b611f3d8282611d1a565b63ffffffff8616600090815260686020526040812090611f5e876001612510565b60ff1660ff1681526020019081526020016000206000600286611f819190612535565b81526020810191909152604001600020555050505050565b6001600160a01b0382166000908152606d602090815260408083208054600182015463ffffffff87168652600383018552838620549351600160f81b9581019590955260218501919091526041840152606086901b6001600160601b031916606184015260a09190911b6001600160a01b0319166075830152906081016040516020818303038152906040528051906020012091505092915050565b60405180604001604052806002906020820280368337509192915050565b8260028101928215612081579160200282015b82811115612081578251825591602001919060010190612066565b5061208d929150612091565b5090565b5b8082111561208d5760008155600101612092565b80356001600160a01b03811681146120bd57600080fd5b919050565b803563ffffffff811681146120bd57600080fd5b600080604083850312156120e957600080fd5b6120f2836120a6565b9150612100602084016120c2565b90509250929050565b60006020828403121561211b57600080fd5b612124826120a6565b9392505050565b803560ff811681146120bd57600080fd5b60008060006060848603121561215157600080fd5b61215a846120c2565b92506121686020850161212b565b9150604084013590509250925092565b6000806040838503121561218b57600080fd5b612194836120c2565b946020939093013593505050565b600080600080608085870312156121b857600080fd5b6121c1856120c2565b935060208501359250604085013591506121dd6060860161212b565b905092959194509250565b6000602082840312156121fa57600080fd5b612124826120c2565b60408101818360005b600281101561222b57815183526020928301929091019060010161220c565b50505092915050565b6000806040838503121561224757600080fd5b612250836120c2565b9150612100602084016120a6565b634e487b7160e01b600052604160045260246000fd5b6000806060838503121561228757600080fd5b612290836120a6565b9150602084603f8501126122a357600080fd5b6040516040810181811067ffffffffffffffff821117156122c6576122c661225e565b6040528060608601878111156122db57600080fd5b8387015b818110156122f657803583529184019184016122df565b50505080925050509250929050565b6000604080830160ff8616845260208281860152818651808452606087019150828801935060005b81811015612358578451805163ffffffff16845284015184840152938301939185019160010161232d565b509098975050505050505050565b6020808252602d908201527f4f6e6c79204c616772616e676520736572766963652063616e2063616c6c207460408201526c3434b990333ab731ba34b7b71760991b606082015260800190565b60208082526025908201527f5468652064656469636174656420636861696e206973206e6f742073756273636040820152641c9a58995960da1b606082015260800190565b6020808252601e908201527f5468652064656469636174656420636861696e206973206c6f636b65642e0000604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000828210156124575761245761242f565b500390565b600060ff821660ff8416808210156124765761247661242f565b90039392505050565b634e487b7160e01b600052603260045260246000fd5b60006000198214156124a9576124a961242f565b5060010190565b60008160001904831182151516156124ca576124ca61242f565b500290565b600082198211156124e2576124e261242f565b500190565b6000602082840312156124f957600080fd5b81516001600160601b038116811461212457600080fd5b600060ff821660ff84168060ff0382111561252d5761252d61242f565b019392505050565b60008261255257634e487b7160e01b600052601260045260246000fd5b500490565b600060ff821660ff81141561256e5761256e61242f565b60010192915050565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220707bd43d3b4187a50d4b79e9206c76c0140f10b2b4af7e23157e9821d377117e64736f6c634300080c0033",
}

// CommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitteeMetaData.ABI instead.
var CommitteeABI = CommitteeMetaData.ABI

// CommitteeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitteeMetaData.Bin instead.
var CommitteeBin = CommitteeMetaData.Bin

// DeployCommittee deploys a new Ethereum contract, binding an instance of Committee to it.
func DeployCommittee(auth *bind.TransactOpts, backend bind.ContractBackend, _service common.Address, _voteWeigher common.Address) (common.Address, *types.Transaction, *Committee, error) {
	parsed, err := CommitteeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitteeBin), backend, _service, _voteWeigher)
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

// INNERNODEPREFIX is a free data retrieval call binding the contract method 0x727cb30f.
//
// Solidity: function INNER_NODE_PREFIX() view returns(bytes1)
func (_Committee *CommitteeCaller) INNERNODEPREFIX(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "INNER_NODE_PREFIX")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INNERNODEPREFIX is a free data retrieval call binding the contract method 0x727cb30f.
//
// Solidity: function INNER_NODE_PREFIX() view returns(bytes1)
func (_Committee *CommitteeSession) INNERNODEPREFIX() ([1]byte, error) {
	return _Committee.Contract.INNERNODEPREFIX(&_Committee.CallOpts)
}

// INNERNODEPREFIX is a free data retrieval call binding the contract method 0x727cb30f.
//
// Solidity: function INNER_NODE_PREFIX() view returns(bytes1)
func (_Committee *CommitteeCallerSession) INNERNODEPREFIX() ([1]byte, error) {
	return _Committee.Contract.INNERNODEPREFIX(&_Committee.CallOpts)
}

// LEAFNODEPREFIX is a free data retrieval call binding the contract method 0x38fa59d1.
//
// Solidity: function LEAF_NODE_PREFIX() view returns(bytes1)
func (_Committee *CommitteeCaller) LEAFNODEPREFIX(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "LEAF_NODE_PREFIX")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// LEAFNODEPREFIX is a free data retrieval call binding the contract method 0x38fa59d1.
//
// Solidity: function LEAF_NODE_PREFIX() view returns(bytes1)
func (_Committee *CommitteeSession) LEAFNODEPREFIX() ([1]byte, error) {
	return _Committee.Contract.LEAFNODEPREFIX(&_Committee.CallOpts)
}

// LEAFNODEPREFIX is a free data retrieval call binding the contract method 0x38fa59d1.
//
// Solidity: function LEAF_NODE_PREFIX() view returns(bytes1)
func (_Committee *CommitteeCallerSession) LEAFNODEPREFIX() ([1]byte, error) {
	return _Committee.Contract.LEAFNODEPREFIX(&_Committee.CallOpts)
}

// CommitteeAddrs is a free data retrieval call binding the contract method 0xbf988ab6.
//
// Solidity: function committeeAddrs(uint32 , uint256 ) view returns(address)
func (_Committee *CommitteeCaller) CommitteeAddrs(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "committeeAddrs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitteeAddrs is a free data retrieval call binding the contract method 0xbf988ab6.
//
// Solidity: function committeeAddrs(uint32 , uint256 ) view returns(address)
func (_Committee *CommitteeSession) CommitteeAddrs(arg0 uint32, arg1 *big.Int) (common.Address, error) {
	return _Committee.Contract.CommitteeAddrs(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeAddrs is a free data retrieval call binding the contract method 0xbf988ab6.
//
// Solidity: function committeeAddrs(uint32 , uint256 ) view returns(address)
func (_Committee *CommitteeCallerSession) CommitteeAddrs(arg0 uint32, arg1 *big.Int) (common.Address, error) {
	return _Committee.Contract.CommitteeAddrs(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeHeights is a free data retrieval call binding the contract method 0xe377dcad.
//
// Solidity: function committeeHeights(uint32 ) view returns(uint8)
func (_Committee *CommitteeCaller) CommitteeHeights(opts *bind.CallOpts, arg0 uint32) (uint8, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "committeeHeights", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CommitteeHeights is a free data retrieval call binding the contract method 0xe377dcad.
//
// Solidity: function committeeHeights(uint32 ) view returns(uint8)
func (_Committee *CommitteeSession) CommitteeHeights(arg0 uint32) (uint8, error) {
	return _Committee.Contract.CommitteeHeights(&_Committee.CallOpts, arg0)
}

// CommitteeHeights is a free data retrieval call binding the contract method 0xe377dcad.
//
// Solidity: function committeeHeights(uint32 ) view returns(uint8)
func (_Committee *CommitteeCallerSession) CommitteeHeights(arg0 uint32) (uint8, error) {
	return _Committee.Contract.CommitteeHeights(&_Committee.CallOpts, arg0)
}

// CommitteeLeavesMap is a free data retrieval call binding the contract method 0xda16ce83.
//
// Solidity: function committeeLeavesMap(uint32 , address ) view returns(uint32)
func (_Committee *CommitteeCaller) CommitteeLeavesMap(opts *bind.CallOpts, arg0 uint32, arg1 common.Address) (uint32, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "committeeLeavesMap", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CommitteeLeavesMap is a free data retrieval call binding the contract method 0xda16ce83.
//
// Solidity: function committeeLeavesMap(uint32 , address ) view returns(uint32)
func (_Committee *CommitteeSession) CommitteeLeavesMap(arg0 uint32, arg1 common.Address) (uint32, error) {
	return _Committee.Contract.CommitteeLeavesMap(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeLeavesMap is a free data retrieval call binding the contract method 0xda16ce83.
//
// Solidity: function committeeLeavesMap(uint32 , address ) view returns(uint32)
func (_Committee *CommitteeCallerSession) CommitteeLeavesMap(arg0 uint32, arg1 common.Address) (uint32, error) {
	return _Committee.Contract.CommitteeLeavesMap(&_Committee.CallOpts, arg0, arg1)
}

// CommitteeNodes is a free data retrieval call binding the contract method 0x34ae7507.
//
// Solidity: function committeeNodes(uint32 , uint8 , uint256 ) view returns(bytes32)
func (_Committee *CommitteeCaller) CommitteeNodes(opts *bind.CallOpts, arg0 uint32, arg1 uint8, arg2 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "committeeNodes", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitteeNodes is a free data retrieval call binding the contract method 0x34ae7507.
//
// Solidity: function committeeNodes(uint32 , uint8 , uint256 ) view returns(bytes32)
func (_Committee *CommitteeSession) CommitteeNodes(arg0 uint32, arg1 uint8, arg2 *big.Int) ([32]byte, error) {
	return _Committee.Contract.CommitteeNodes(&_Committee.CallOpts, arg0, arg1, arg2)
}

// CommitteeNodes is a free data retrieval call binding the contract method 0x34ae7507.
//
// Solidity: function committeeNodes(uint32 , uint8 , uint256 ) view returns(bytes32)
func (_Committee *CommitteeCallerSession) CommitteeNodes(arg0 uint32, arg1 uint8, arg2 *big.Int) ([32]byte, error) {
	return _Committee.Contract.CommitteeNodes(&_Committee.CallOpts, arg0, arg1, arg2)
}

// CommitteeParams is a free data retrieval call binding the contract method 0x72856455.
//
// Solidity: function committeeParams(uint32 ) view returns(uint256 startBlock, uint256 duration, uint256 freezeDuration, uint8 quorumNumber)
func (_Committee *CommitteeCaller) CommitteeParams(opts *bind.CallOpts, arg0 uint32) (struct {
	StartBlock     *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	QuorumNumber   uint8
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "committeeParams", arg0)

	outstruct := new(struct {
		StartBlock     *big.Int
		Duration       *big.Int
		FreezeDuration *big.Int
		QuorumNumber   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FreezeDuration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.QuorumNumber = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// CommitteeParams is a free data retrieval call binding the contract method 0x72856455.
//
// Solidity: function committeeParams(uint32 ) view returns(uint256 startBlock, uint256 duration, uint256 freezeDuration, uint8 quorumNumber)
func (_Committee *CommitteeSession) CommitteeParams(arg0 uint32) (struct {
	StartBlock     *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	QuorumNumber   uint8
}, error) {
	return _Committee.Contract.CommitteeParams(&_Committee.CallOpts, arg0)
}

// CommitteeParams is a free data retrieval call binding the contract method 0x72856455.
//
// Solidity: function committeeParams(uint32 ) view returns(uint256 startBlock, uint256 duration, uint256 freezeDuration, uint8 quorumNumber)
func (_Committee *CommitteeCallerSession) CommitteeParams(arg0 uint32) (struct {
	StartBlock     *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	QuorumNumber   uint8
}, error) {
	return _Committee.Contract.CommitteeParams(&_Committee.CallOpts, arg0)
}

// Committees is a free data retrieval call binding the contract method 0xa63490a2.
//
// Solidity: function committees(uint32 , uint256 ) view returns(bytes32 root, uint256 leafCount, uint256 totalVotingPower)
func (_Committee *CommitteeCaller) Committees(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (struct {
	Root             [32]byte
	LeafCount        *big.Int
	TotalVotingPower *big.Int
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "committees", arg0, arg1)

	outstruct := new(struct {
		Root             [32]byte
		LeafCount        *big.Int
		TotalVotingPower *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.LeafCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalVotingPower = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Committees is a free data retrieval call binding the contract method 0xa63490a2.
//
// Solidity: function committees(uint32 , uint256 ) view returns(bytes32 root, uint256 leafCount, uint256 totalVotingPower)
func (_Committee *CommitteeSession) Committees(arg0 uint32, arg1 *big.Int) (struct {
	Root             [32]byte
	LeafCount        *big.Int
	TotalVotingPower *big.Int
}, error) {
	return _Committee.Contract.Committees(&_Committee.CallOpts, arg0, arg1)
}

// Committees is a free data retrieval call binding the contract method 0xa63490a2.
//
// Solidity: function committees(uint32 , uint256 ) view returns(bytes32 root, uint256 leafCount, uint256 totalVotingPower)
func (_Committee *CommitteeCallerSession) Committees(arg0 uint32, arg1 *big.Int) (struct {
	Root             [32]byte
	LeafCount        *big.Int
	TotalVotingPower *big.Int
}, error) {
	return _Committee.Contract.Committees(&_Committee.CallOpts, arg0, arg1)
}

// GetBlsPubKey is a free data retrieval call binding the contract method 0x999c3b4a.
//
// Solidity: function getBlsPubKey(address operator) view returns(uint256[2])
func (_Committee *CommitteeCaller) GetBlsPubKey(opts *bind.CallOpts, operator common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getBlsPubKey", operator)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBlsPubKey is a free data retrieval call binding the contract method 0x999c3b4a.
//
// Solidity: function getBlsPubKey(address operator) view returns(uint256[2])
func (_Committee *CommitteeSession) GetBlsPubKey(operator common.Address) ([2]*big.Int, error) {
	return _Committee.Contract.GetBlsPubKey(&_Committee.CallOpts, operator)
}

// GetBlsPubKey is a free data retrieval call binding the contract method 0x999c3b4a.
//
// Solidity: function getBlsPubKey(address operator) view returns(uint256[2])
func (_Committee *CommitteeCallerSession) GetBlsPubKey(operator common.Address) ([2]*big.Int, error) {
	return _Committee.Contract.GetBlsPubKey(&_Committee.CallOpts, operator)
}

// GetCommittee is a free data retrieval call binding the contract method 0xdef9e7d5.
//
// Solidity: function getCommittee(uint32 chainID, uint256 blockNumber) view returns((bytes32,uint256,uint256) currentCommittee, bytes32 nextRoot)
func (_Committee *CommitteeCaller) GetCommittee(opts *bind.CallOpts, chainID uint32, blockNumber *big.Int) (struct {
	CurrentCommittee ILagrangeCommitteeCommitteeData
	NextRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getCommittee", chainID, blockNumber)

	outstruct := new(struct {
		CurrentCommittee ILagrangeCommitteeCommitteeData
		NextRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentCommittee = *abi.ConvertType(out[0], new(ILagrangeCommitteeCommitteeData)).(*ILagrangeCommitteeCommitteeData)
	outstruct.NextRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetCommittee is a free data retrieval call binding the contract method 0xdef9e7d5.
//
// Solidity: function getCommittee(uint32 chainID, uint256 blockNumber) view returns((bytes32,uint256,uint256) currentCommittee, bytes32 nextRoot)
func (_Committee *CommitteeSession) GetCommittee(chainID uint32, blockNumber *big.Int) (struct {
	CurrentCommittee ILagrangeCommitteeCommitteeData
	NextRoot         [32]byte
}, error) {
	return _Committee.Contract.GetCommittee(&_Committee.CallOpts, chainID, blockNumber)
}

// GetCommittee is a free data retrieval call binding the contract method 0xdef9e7d5.
//
// Solidity: function getCommittee(uint32 chainID, uint256 blockNumber) view returns((bytes32,uint256,uint256) currentCommittee, bytes32 nextRoot)
func (_Committee *CommitteeCallerSession) GetCommittee(chainID uint32, blockNumber *big.Int) (struct {
	CurrentCommittee ILagrangeCommitteeCommitteeData
	NextRoot         [32]byte
}, error) {
	return _Committee.Contract.GetCommittee(&_Committee.CallOpts, chainID, blockNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x78d81d08.
//
// Solidity: function getEpochNumber(uint32 chainID, uint256 blockNumber) view returns(uint256)
func (_Committee *CommitteeCaller) GetEpochNumber(opts *bind.CallOpts, chainID uint32, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getEpochNumber", chainID, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x78d81d08.
//
// Solidity: function getEpochNumber(uint32 chainID, uint256 blockNumber) view returns(uint256)
func (_Committee *CommitteeSession) GetEpochNumber(chainID uint32, blockNumber *big.Int) (*big.Int, error) {
	return _Committee.Contract.GetEpochNumber(&_Committee.CallOpts, chainID, blockNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x78d81d08.
//
// Solidity: function getEpochNumber(uint32 chainID, uint256 blockNumber) view returns(uint256)
func (_Committee *CommitteeCallerSession) GetEpochNumber(chainID uint32, blockNumber *big.Int) (*big.Int, error) {
	return _Committee.Contract.GetEpochNumber(&_Committee.CallOpts, chainID, blockNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address opAddr) view returns(uint8, (uint32,uint256)[])
func (_Committee *CommitteeCaller) GetOperatorStatus(opts *bind.CallOpts, opAddr common.Address) (uint8, []ILagrangeCommitteeUnsubscribedParam, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getOperatorStatus", opAddr)

	if err != nil {
		return *new(uint8), *new([]ILagrangeCommitteeUnsubscribedParam), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new([]ILagrangeCommitteeUnsubscribedParam)).(*[]ILagrangeCommitteeUnsubscribedParam)

	return out0, out1, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address opAddr) view returns(uint8, (uint32,uint256)[])
func (_Committee *CommitteeSession) GetOperatorStatus(opAddr common.Address) (uint8, []ILagrangeCommitteeUnsubscribedParam, error) {
	return _Committee.Contract.GetOperatorStatus(&_Committee.CallOpts, opAddr)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address opAddr) view returns(uint8, (uint32,uint256)[])
func (_Committee *CommitteeCallerSession) GetOperatorStatus(opAddr common.Address) (uint8, []ILagrangeCommitteeUnsubscribedParam, error) {
	return _Committee.Contract.GetOperatorStatus(&_Committee.CallOpts, opAddr)
}

// GetOperatorVotingPower is a free data retrieval call binding the contract method 0x5af1a88f.
//
// Solidity: function getOperatorVotingPower(address opAddr, uint32 chainID) view returns(uint96)
func (_Committee *CommitteeCaller) GetOperatorVotingPower(opts *bind.CallOpts, opAddr common.Address, chainID uint32) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "getOperatorVotingPower", opAddr, chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorVotingPower is a free data retrieval call binding the contract method 0x5af1a88f.
//
// Solidity: function getOperatorVotingPower(address opAddr, uint32 chainID) view returns(uint96)
func (_Committee *CommitteeSession) GetOperatorVotingPower(opAddr common.Address, chainID uint32) (*big.Int, error) {
	return _Committee.Contract.GetOperatorVotingPower(&_Committee.CallOpts, opAddr, chainID)
}

// GetOperatorVotingPower is a free data retrieval call binding the contract method 0x5af1a88f.
//
// Solidity: function getOperatorVotingPower(address opAddr, uint32 chainID) view returns(uint96)
func (_Committee *CommitteeCallerSession) GetOperatorVotingPower(opAddr common.Address, chainID uint32) (*big.Int, error) {
	return _Committee.Contract.GetOperatorVotingPower(&_Committee.CallOpts, opAddr, chainID)
}

// IsLocked is a free data retrieval call binding the contract method 0x3d6a2679.
//
// Solidity: function isLocked(uint32 chainID) view returns(bool, uint256)
func (_Committee *CommitteeCaller) IsLocked(opts *bind.CallOpts, chainID uint32) (bool, *big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "isLocked", chainID)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// IsLocked is a free data retrieval call binding the contract method 0x3d6a2679.
//
// Solidity: function isLocked(uint32 chainID) view returns(bool, uint256)
func (_Committee *CommitteeSession) IsLocked(chainID uint32) (bool, *big.Int, error) {
	return _Committee.Contract.IsLocked(&_Committee.CallOpts, chainID)
}

// IsLocked is a free data retrieval call binding the contract method 0x3d6a2679.
//
// Solidity: function isLocked(uint32 chainID) view returns(bool, uint256)
func (_Committee *CommitteeCallerSession) IsLocked(chainID uint32) (bool, *big.Int, error) {
	return _Committee.Contract.IsLocked(&_Committee.CallOpts, chainID)
}

// IsUnregisterable is a free data retrieval call binding the contract method 0x19a74c5f.
//
// Solidity: function isUnregisterable(address operator) view returns(bool, uint256)
func (_Committee *CommitteeCaller) IsUnregisterable(opts *bind.CallOpts, operator common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "isUnregisterable", operator)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// IsUnregisterable is a free data retrieval call binding the contract method 0x19a74c5f.
//
// Solidity: function isUnregisterable(address operator) view returns(bool, uint256)
func (_Committee *CommitteeSession) IsUnregisterable(operator common.Address) (bool, *big.Int, error) {
	return _Committee.Contract.IsUnregisterable(&_Committee.CallOpts, operator)
}

// IsUnregisterable is a free data retrieval call binding the contract method 0x19a74c5f.
//
// Solidity: function isUnregisterable(address operator) view returns(bool, uint256)
func (_Committee *CommitteeCallerSession) IsUnregisterable(operator common.Address) (bool, *big.Int, error) {
	return _Committee.Contract.IsUnregisterable(&_Committee.CallOpts, operator)
}

// IsUpdatable is a free data retrieval call binding the contract method 0x85ab9a7a.
//
// Solidity: function isUpdatable(uint32 chainID, uint256 epochNumber) view returns(bool)
func (_Committee *CommitteeCaller) IsUpdatable(opts *bind.CallOpts, chainID uint32, epochNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "isUpdatable", chainID, epochNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdatable is a free data retrieval call binding the contract method 0x85ab9a7a.
//
// Solidity: function isUpdatable(uint32 chainID, uint256 epochNumber) view returns(bool)
func (_Committee *CommitteeSession) IsUpdatable(chainID uint32, epochNumber *big.Int) (bool, error) {
	return _Committee.Contract.IsUpdatable(&_Committee.CallOpts, chainID, epochNumber)
}

// IsUpdatable is a free data retrieval call binding the contract method 0x85ab9a7a.
//
// Solidity: function isUpdatable(uint32 chainID, uint256 epochNumber) view returns(bool)
func (_Committee *CommitteeCallerSession) IsUpdatable(chainID uint32, epochNumber *big.Int) (bool, error) {
	return _Committee.Contract.IsUpdatable(&_Committee.CallOpts, chainID, epochNumber)
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

// TotalVotingPower is a free data retrieval call binding the contract method 0x8dfccad8.
//
// Solidity: function totalVotingPower(uint32 ) view returns(uint256)
func (_Committee *CommitteeCaller) TotalVotingPower(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "totalVotingPower", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotingPower is a free data retrieval call binding the contract method 0x8dfccad8.
//
// Solidity: function totalVotingPower(uint32 ) view returns(uint256)
func (_Committee *CommitteeSession) TotalVotingPower(arg0 uint32) (*big.Int, error) {
	return _Committee.Contract.TotalVotingPower(&_Committee.CallOpts, arg0)
}

// TotalVotingPower is a free data retrieval call binding the contract method 0x8dfccad8.
//
// Solidity: function totalVotingPower(uint32 ) view returns(uint256)
func (_Committee *CommitteeCallerSession) TotalVotingPower(arg0 uint32) (*big.Int, error) {
	return _Committee.Contract.TotalVotingPower(&_Committee.CallOpts, arg0)
}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x4db6f74a.
//
// Solidity: function updatedEpoch(uint32 ) view returns(uint256)
func (_Committee *CommitteeCaller) UpdatedEpoch(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "updatedEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x4db6f74a.
//
// Solidity: function updatedEpoch(uint32 ) view returns(uint256)
func (_Committee *CommitteeSession) UpdatedEpoch(arg0 uint32) (*big.Int, error) {
	return _Committee.Contract.UpdatedEpoch(&_Committee.CallOpts, arg0)
}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x4db6f74a.
//
// Solidity: function updatedEpoch(uint32 ) view returns(uint256)
func (_Committee *CommitteeCallerSession) UpdatedEpoch(arg0 uint32) (*big.Int, error) {
	return _Committee.Contract.UpdatedEpoch(&_Committee.CallOpts, arg0)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_Committee *CommitteeCaller) VoteWeigher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Committee.contract.Call(opts, &out, "voteWeigher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_Committee *CommitteeSession) VoteWeigher() (common.Address, error) {
	return _Committee.Contract.VoteWeigher(&_Committee.CallOpts)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_Committee *CommitteeCallerSession) VoteWeigher() (common.Address, error) {
	return _Committee.Contract.VoteWeigher(&_Committee.CallOpts)
}

// AddOperator is a paid mutator transaction binding the contract method 0xf1d3353e.
//
// Solidity: function addOperator(address operator, uint256[2] blsPubKey) returns()
func (_Committee *CommitteeTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "addOperator", operator, blsPubKey)
}

// AddOperator is a paid mutator transaction binding the contract method 0xf1d3353e.
//
// Solidity: function addOperator(address operator, uint256[2] blsPubKey) returns()
func (_Committee *CommitteeSession) AddOperator(operator common.Address, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _Committee.Contract.AddOperator(&_Committee.TransactOpts, operator, blsPubKey)
}

// AddOperator is a paid mutator transaction binding the contract method 0xf1d3353e.
//
// Solidity: function addOperator(address operator, uint256[2] blsPubKey) returns()
func (_Committee *CommitteeTransactorSession) AddOperator(operator common.Address, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _Committee.Contract.AddOperator(&_Committee.TransactOpts, operator, blsPubKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Committee *CommitteeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Committee *CommitteeSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Committee.Contract.Initialize(&_Committee.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Committee *CommitteeTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Committee.Contract.Initialize(&_Committee.TransactOpts, initialOwner)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x3bc9c733.
//
// Solidity: function registerChain(uint32 chainID, uint256 epochPeriod, uint256 freezeDuration, uint8 quorunNumber) returns()
func (_Committee *CommitteeTransactor) RegisterChain(opts *bind.TransactOpts, chainID uint32, epochPeriod *big.Int, freezeDuration *big.Int, quorunNumber uint8) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "registerChain", chainID, epochPeriod, freezeDuration, quorunNumber)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x3bc9c733.
//
// Solidity: function registerChain(uint32 chainID, uint256 epochPeriod, uint256 freezeDuration, uint8 quorunNumber) returns()
func (_Committee *CommitteeSession) RegisterChain(chainID uint32, epochPeriod *big.Int, freezeDuration *big.Int, quorunNumber uint8) (*types.Transaction, error) {
	return _Committee.Contract.RegisterChain(&_Committee.TransactOpts, chainID, epochPeriod, freezeDuration, quorunNumber)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x3bc9c733.
//
// Solidity: function registerChain(uint32 chainID, uint256 epochPeriod, uint256 freezeDuration, uint8 quorunNumber) returns()
func (_Committee *CommitteeTransactorSession) RegisterChain(chainID uint32, epochPeriod *big.Int, freezeDuration *big.Int, quorunNumber uint8) (*types.Transaction, error) {
	return _Committee.Contract.RegisterChain(&_Committee.TransactOpts, chainID, epochPeriod, freezeDuration, quorunNumber)
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

// SubscribeChain is a paid mutator transaction binding the contract method 0x6b11c38e.
//
// Solidity: function subscribeChain(address operator, uint32 chainID) returns()
func (_Committee *CommitteeTransactor) SubscribeChain(opts *bind.TransactOpts, operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "subscribeChain", operator, chainID)
}

// SubscribeChain is a paid mutator transaction binding the contract method 0x6b11c38e.
//
// Solidity: function subscribeChain(address operator, uint32 chainID) returns()
func (_Committee *CommitteeSession) SubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.Contract.SubscribeChain(&_Committee.TransactOpts, operator, chainID)
}

// SubscribeChain is a paid mutator transaction binding the contract method 0x6b11c38e.
//
// Solidity: function subscribeChain(address operator, uint32 chainID) returns()
func (_Committee *CommitteeTransactorSession) SubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.Contract.SubscribeChain(&_Committee.TransactOpts, operator, chainID)
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

// UnsubscribeChain is a paid mutator transaction binding the contract method 0x0e9f564b.
//
// Solidity: function unsubscribeChain(address operator, uint32 chainID) returns()
func (_Committee *CommitteeTransactor) UnsubscribeChain(opts *bind.TransactOpts, operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "unsubscribeChain", operator, chainID)
}

// UnsubscribeChain is a paid mutator transaction binding the contract method 0x0e9f564b.
//
// Solidity: function unsubscribeChain(address operator, uint32 chainID) returns()
func (_Committee *CommitteeSession) UnsubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.Contract.UnsubscribeChain(&_Committee.TransactOpts, operator, chainID)
}

// UnsubscribeChain is a paid mutator transaction binding the contract method 0x0e9f564b.
//
// Solidity: function unsubscribeChain(address operator, uint32 chainID) returns()
func (_Committee *CommitteeTransactorSession) UnsubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.Contract.UnsubscribeChain(&_Committee.TransactOpts, operator, chainID)
}

// Update is a paid mutator transaction binding the contract method 0x3bc72805.
//
// Solidity: function update(uint32 chainID, uint256 epochNumber) returns()
func (_Committee *CommitteeTransactor) Update(opts *bind.TransactOpts, chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "update", chainID, epochNumber)
}

// Update is a paid mutator transaction binding the contract method 0x3bc72805.
//
// Solidity: function update(uint32 chainID, uint256 epochNumber) returns()
func (_Committee *CommitteeSession) Update(chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _Committee.Contract.Update(&_Committee.TransactOpts, chainID, epochNumber)
}

// Update is a paid mutator transaction binding the contract method 0x3bc72805.
//
// Solidity: function update(uint32 chainID, uint256 epochNumber) returns()
func (_Committee *CommitteeTransactorSession) Update(chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _Committee.Contract.Update(&_Committee.TransactOpts, chainID, epochNumber)
}

// UpdateOperatorAmount is a paid mutator transaction binding the contract method 0x798c7244.
//
// Solidity: function updateOperatorAmount(address operator, uint32 chainID) returns()
func (_Committee *CommitteeTransactor) UpdateOperatorAmount(opts *bind.TransactOpts, operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.contract.Transact(opts, "updateOperatorAmount", operator, chainID)
}

// UpdateOperatorAmount is a paid mutator transaction binding the contract method 0x798c7244.
//
// Solidity: function updateOperatorAmount(address operator, uint32 chainID) returns()
func (_Committee *CommitteeSession) UpdateOperatorAmount(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.Contract.UpdateOperatorAmount(&_Committee.TransactOpts, operator, chainID)
}

// UpdateOperatorAmount is a paid mutator transaction binding the contract method 0x798c7244.
//
// Solidity: function updateOperatorAmount(address operator, uint32 chainID) returns()
func (_Committee *CommitteeTransactorSession) UpdateOperatorAmount(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _Committee.Contract.UpdateOperatorAmount(&_Committee.TransactOpts, operator, chainID)
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
	QuorumNumber   uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitCommittee is a free log retrieval operation binding the contract event 0xe4166ce16a6b34a7c665c8a0b0cc0cfac48559a82b8dae718ce478257becfa1a.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration, uint8 quorumNumber)
func (_Committee *CommitteeFilterer) FilterInitCommittee(opts *bind.FilterOpts) (*CommitteeInitCommitteeIterator, error) {

	logs, sub, err := _Committee.contract.FilterLogs(opts, "InitCommittee")
	if err != nil {
		return nil, err
	}
	return &CommitteeInitCommitteeIterator{contract: _Committee.contract, event: "InitCommittee", logs: logs, sub: sub}, nil
}

// WatchInitCommittee is a free log subscription operation binding the contract event 0xe4166ce16a6b34a7c665c8a0b0cc0cfac48559a82b8dae718ce478257becfa1a.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration, uint8 quorumNumber)
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

// ParseInitCommittee is a log parse operation binding the contract event 0xe4166ce16a6b34a7c665c8a0b0cc0cfac48559a82b8dae718ce478257becfa1a.
//
// Solidity: event InitCommittee(uint256 chainID, uint256 duration, uint256 freezeDuration, uint8 quorumNumber)
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
