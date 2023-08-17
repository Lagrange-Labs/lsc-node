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
	_ = abi.ConvertType
)

// EvidenceVerifierEvidence is an auto generated low-level Go binding around an user-defined struct.
type EvidenceVerifierEvidence struct {
	Operator                    common.Address
	BlockHash                   [32]byte
	CorrectBlockHash            [32]byte
	CurrentCommitteeRoot        [32]byte
	CorrectCurrentCommitteeRoot [32]byte
	NextCommitteeRoot           [32]byte
	CorrectNextCommitteeRoot    [32]byte
	BlockNumber                 *big.Int
	EpochBlockNumber            *big.Int
	BlockSignature              []byte
	CommitSignature             []byte
	ChainID                     uint32
	RawBlockHeader              []byte
}

// RLPReaderRLPItem is an auto generated low-level Go binding around an user-defined struct.
type RLPReaderRLPItem struct {
	Len    *big.Int
	MemPtr *big.Int
}

// LagrangeMetaData contains all meta data concerning the Lagrange contract.
var LagrangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"_committee\",\"type\":\"address\"},{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"UploadEvidence\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_HEADER_EXTRADATA_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_HEADER_NUMBER_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_ARBITRUM_NITRO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_MAINNET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_OPTIMISM_BEDROCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_AMOUNT_CHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_REGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_UNREGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"latestHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"sequence\",\"type\":\"bytes[]\"}],\"name\":\"_verifyRawHeaderSequence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"}],\"name\":\"calculateBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"}],\"name\":\"checkAndDecodeRLP\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPtr\",\"type\":\"uint256\"}],\"internalType\":\"structRLPReader.RLPItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"rawBlockHeader\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"checkCommitSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArbAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"rawBlockHeader\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractArbitrumVerifier\",\"name\":\"_arb\",\"type\":\"address\"}],\"name\":\"setArbAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptimismVerifier\",\"name\":\"_opt\",\"type\":\"address\"}],\"name\":\"setOptAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"rawBlockHeader\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"uploadEvidence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"comparisonNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"verifyBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200210d3803806200210d833981016040819052620000349162000134565b6001600160a01b03808316608052811660a0526200005162000059565b505062000173565b600054610100900460ff1615620000c65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000119576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013157600080fd5b50565b600080604083850312156200014857600080fd5b825162000155816200011b565b602084015190925062000168816200011b565b809150509250929050565b60805160a051611f43620001ca600039600081816101de01528181610484015261110a01526000818161037f015281816103f8015281816108640152818161095d0152818161109101526111d30152611f436000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806375274711116100f9578063cec9892b11610097578063f2fde38b11610071578063f2fde38b146103a1578063f44c5c71146103b4578063f98fe1c4146103cd578063fd793ed5146103d757600080fd5b8063cec9892b14610372578063d742da1a14610372578063d864e7401461037a57600080fd5b8063acc41352116100d3578063acc4135214610319578063aef524b01461032c578063ba42f69e1461034c578063c4d66de81461035f57600080fd5b806375274711146102eb578063873b05a2146102fe5780638da5cb5b1461031157600080fd5b80634fdb9291116101665780636203902211610140578063620390221461029e578063655f61a2146102c1578063715018a6146102d25780637229b160146102da57600080fd5b80634fdb9291146102535780635364f104146102665780635ddbe7f51461029657600080fd5b806314501001146101ae5780631d393c09146101c957806323097e86146101d15780633998fdd3146101d95780633c3059c7146102185780633df379a914610221575b600080fd5b6101b6600881565b6040519081526020015b60405180910390f35b6101b6600381565b6101b6600281565b6102007f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101c0565b6101b66101a481565b61025161022f3660046118b9565b606580546001600160a01b0319166001600160a01b0392909216919091179055565b005b61025161026136600461198b565b6103e1565b6102516102743660046118b9565b606680546001600160a01b0319166001600160a01b0392909216919091179055565b6101b6600c81565b6102b16102ac3660046119ef565b61052e565b60405190151581526020016101c0565b6065546001600160a01b0316610200565b610251610547565b6066546001600160a01b0316610200565b6102b16102f9366004611a46565b61055b565b6101b661030c366004611a46565b6105db565b61020061064b565b6102b1610327366004611a82565b610664565b61033f61033a366004611b01565b61079e565b6040516101c09190611b46565b61025161035a366004611a46565b610858565b61025161036d3660046118b9565b610bdd565b6101b6600181565b6102007f000000000000000000000000000000000000000000000000000000000000000081565b6102516103af3660046118b9565b610cf0565b6101b66103c2366004611b95565b805160209091012090565b6101b662014a3381565b6101b662066eed81565b6040516327067c2360e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634e0cf84690610433903390869088908790600401611bca565b600060405180830381600087803b15801561044d57600080fd5b505af1158015610461573d6000803e3d6000fd5b505060405163175d320560e01b815233600482015263ffffffff841660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063175d32059150604401600060405180830381600087803b1580156104d257600080fd5b505af11580156104e6573d6000803e3d6000fd5b50506040805133815263ffffffff851660208201527f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a935001905060405180910390a1505050565b60008061053d86868686610d69565b9695505050505050565b61054f610db0565b6105596000610e0f565b565b600080610567836105db565b905060006105b78261057d610140870187611c50565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e6192505050565b90506105c660208501856118b9565b6001600160a01b039182169116149392505050565b60006020820135606083013560a084013560e0850135610100860135610605610120880188611c50565b6106176101808a016101608b01611c97565b60405160200161062e989796959493929190611cb4565b604051602081830303815290604052805190602001209050919050565b600061065f6033546001600160a01b031690565b905090565b60008060005b8381101561077f5760006106dd6106d887878581811061068c5761068c611d00565b905060200281019061069e9190611c50565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8592505050565b610eb2565b90506000816000815181106106f4576106f4611d00565b60200260200101519050600061070982610fc8565b9050831580159061071a5750848114155b1561072d57600095505050505050610797565b87878581811061073f5761073f611d00565b90506020028101906107519190611c50565b60405161075f929190611d16565b60405180910390209450505050808061077790611d3c565b91505061066a565b50808514610791576000915050610797565b60019150505b9392505050565b815160208301206060908281146108195760405162461bcd60e51b815260206004820152603460248201527f48617368206f6620524c5020646174612064697665726765732066726f6d20636044820152730dedae0c2e4d2e6dedc40c4d8dec6d640d0c2e6d60631b60648201526084015b60405180910390fd5b600061084f6106d88660408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b95945050505050565b60006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344a5c4bf61089660208501856118b9565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af11580156108dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109009190611d57565b63ffffffff16116109535760405162461bcd60e51b815260206004820152601e60248201527f546865206f70657261746f72206973206e6f74207265676973746572656400006044820152606401610810565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344f5b6b461098f60208401846118b9565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af11580156109d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f99190611d74565b15610a465760405162461bcd60e51b815260206004820152601760248201527f546865206f70657261746f7220697320736c61736865640000000000000000006044820152606401610810565b610a4f8161055b565b610aa75760405162461bcd60e51b815260206004820152602360248201527f54686520636f6d6d6974207369676e6174757265206973206e6f7420636f72726044820152621958dd60ea1b6064820152608401610810565b610ae46040820135602083013560e0840135610ac7610180860186611c50565b610ad961018088016101608901611c97565b63ffffffff16611016565b610afd57610afd610af860208301836118b9565b611072565b610b316080820135606083013560c084013560a0850135610100860135610b2c61018088016101608901611c97565b6111a2565b610b4557610b45610af860208301836118b9565b7fa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f610b7360208301836118b9565b6020830135606084013560a085013560e0860135610100870135610b9b610120890189611c50565b610ba96101408b018b611c50565b610bbb6101808d016101608e01611c97565b604051610bd29b9a99989796959493929190611dbf565b60405180910390a150565b600054610100900460ff1615808015610bfd5750600054600160ff909116105b80610c175750303b158015610c17575060005460ff166001145b610c7a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610810565b6000805460ff191660011790558015610c9d576000805461ff0019166101001790555b610ca682610e0f565b8015610cec576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b610cf8610db0565b6001600160a01b038116610d5d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610810565b610d6681610e0f565b50565b600080610d76858561079e565b9050600081600881518110610d8d57610d8d611d00565b602002602001015190506000610da282610fc8565b909714979650505050505050565b33610db961064b565b6001600160a01b0316146105595760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610810565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806000610e70858561132c565b91509150610e7d8161139c565b509392505050565b60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6060610ebd82611557565b610ec657600080fd5b6000610ed183611590565b905060008167ffffffffffffffff811115610eee57610eee6118e8565b604051908082528060200260200182016040528015610f3357816020015b6040805180820190915260008082526020820152815260200190600190039081610f0c5790505b5090506000610f458560200151611613565b8560200151610f549190611e35565b90506000805b84811015610fbd57610f6b8361168e565b9150604051806040016040528083815260200184815250848281518110610f9457610f94611d00565b6020908102919091010152610fa98284611e35565b925080610fb581611d3c565b915050610f5a565b509195945050505050565b805160009015801590610fdd57508151602110155b610fe657600080fd5b600080610ff284611737565b81519193509150602082101561100e5760208290036101000a90045b949350505050565b600061105c8585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c925087915061052e9050565b801561106757508686145b979650505050505050565b604051633dd3153360e21b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f74c54cc90602401600060405180830381600087803b1580156110d557600080fd5b505af11580156110e9573d6000803e3d6000fd5b5050604051630e323b9960e21b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001692506338c8ee649150602401600060405180830381600087803b15801561115057600080fd5b505af1158015611164573d6000803e3d6000fd5b50506040516001600160a01b03841681527fd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a306892506020019050610bd2565b60405163def9e7d560e01b815263ffffffff8216600482015260248101839052600090819081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063def9e7d5906044016080604051808303816000875af115801561121c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112409190611e4d565b8151919350915089146112ad5760405162461bcd60e51b815260206004820152602f60248201527f5265666572656e63652063757272656e7420636f6d6d697474656520726f6f7460448201526e39903237903737ba1036b0ba31b41760891b6064820152608401610810565b8681146113115760405162461bcd60e51b815260206004820152602c60248201527f5265666572656e6365206e65787420636f6d6d697474656520726f6f7473206460448201526b37903737ba1036b0ba31b41760a11b6064820152608401610810565b888814801561131f57508686145b9998505050505050505050565b6000808251604114156113635760208301516040840151606085015160001a6113578782858561177e565b94509450505050611395565b82516040141561138d576020830151604084015161138286838361186b565b935093505050611395565b506000905060025b9250929050565b60008160048111156113b0576113b0611ebd565b14156113b95750565b60018160048111156113cd576113cd611ebd565b141561141b5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610810565b600281600481111561142f5761142f611ebd565b141561147d5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610810565b600381600481111561149157611491611ebd565b14156114ea5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610810565b60048160048111156114fe576114fe611ebd565b1415610d665760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610810565b805160009061156857506000919050565b6020820151805160001a9060c0821015611586575060009392505050565b5060019392505050565b80516000906115a157506000919050565b6000806115b18460200151611613565b84602001516115c09190611e35565b90506000846000015185602001516115d89190611e35565b90505b8082101561160a576115ec8261168e565b6115f69083611e35565b91508261160281611d3c565b9350506115db565b50909392505050565b8051600090811a608081101561162c5750600092915050565b60b8811080611647575060c08110801590611647575060f881105b156116555750600192915050565b60c08110156116825761166a600160b8611ed3565b6116779060ff1682611ef6565b610797906001611e35565b61166a600160f8611ed3565b80516000908190811a60808110156116a95760019150611730565b60b88110156116cf576116bd608082611ef6565b6116c8906001611e35565b9150611730565b60c08110156116fc5760b78103600185019450806020036101000a85510460018201810193505050611730565b60f8811015611710576116bd60c082611ef6565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b60008060006117498460200151611613565b9050600081856020015161175d9190611e35565b905060008286600001516117719190611ef6565b9196919550909350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156117b55750600090506003611862565b8460ff16601b141580156117cd57508460ff16601c14155b156117de5750600090506004611862565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611832573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661185b57600060019250925050611862565b9150600090505b94509492505050565b6000806001600160ff1b0383168161188860ff86901c601b611e35565b90506118968782888561177e565b935093505050935093915050565b6001600160a01b0381168114610d6657600080fd5b6000602082840312156118cb57600080fd5b8135610797816118a4565b63ffffffff81168114610d6657600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f83011261190f57600080fd5b813567ffffffffffffffff8082111561192a5761192a6118e8565b604051601f8301601f19908116603f01168101908282118183101715611952576119526118e8565b8160405283815286602085880101111561196b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156119a057600080fd5b83356119ab816118d6565b9250602084013567ffffffffffffffff8111156119c757600080fd5b6119d3868287016118fe565b92505060408401356119e4816118d6565b809150509250925092565b60008060008060808587031215611a0557600080fd5b84359350602085013567ffffffffffffffff811115611a2357600080fd5b611a2f878288016118fe565b949794965050505060408301359260600135919050565b600060208284031215611a5857600080fd5b813567ffffffffffffffff811115611a6f57600080fd5b82016101a0818503121561079757600080fd5b600080600060408486031215611a9757600080fd5b83359250602084013567ffffffffffffffff80821115611ab657600080fd5b818601915086601f830112611aca57600080fd5b813581811115611ad957600080fd5b8760208260051b8501011115611aee57600080fd5b6020830194508093505050509250925092565b60008060408385031215611b1457600080fd5b823567ffffffffffffffff811115611b2b57600080fd5b611b37858286016118fe565b95602094909401359450505050565b602080825282518282018190526000919060409081850190868401855b82811015611b8857815180518552860151868501529284019290850190600101611b63565b5091979650505050505050565b600060208284031215611ba757600080fd5b813567ffffffffffffffff811115611bbe57600080fd5b61100e848285016118fe565b60018060a01b038516815260006020608081840152855180608085015260005b81811015611c065787810183015185820160a001528201611bea565b81811115611c1857600060a083870101525b50601f01601f1916830160a0019150611c3b9050604083018563ffffffff169052565b63ffffffff8316606083015295945050505050565b6000808335601e19843603018112611c6757600080fd5b83018035915067ffffffffffffffff821115611c8257600080fd5b60200191503681900382131561139557600080fd5b600060208284031215611ca957600080fd5b8135610797816118d6565b888152876020820152866040820152856060820152846080820152828460a083013760e09190911b6001600160e01b03191660a0919092019081019190915260a4019695505050505050565b634e487b7160e01b600052603260045260246000fd5b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415611d5057611d50611d26565b5060010190565b600060208284031215611d6957600080fd5b8151610797816118d6565b600060208284031215611d8657600080fd5b8151801515811461079757600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061012060018060a01b038e1683528c60208401528b60408401528a60608401528960808401528860a08401528060c0840152611e00818401888a611d96565b905082810360e0840152611e15818688611d96565b91505063ffffffff83166101008301529c9b505050505050505050505050565b60008219821115611e4857611e48611d26565b500190565b6000808284036080811215611e6157600080fd5b6060811215611e6f57600080fd5b506040516060810181811067ffffffffffffffff82111715611e9357611e936118e8565b60409081528451825260208086015190830152848101519082015260609093015192949293505050565b634e487b7160e01b600052602160045260246000fd5b600060ff821660ff841680821015611eed57611eed611d26565b90039392505050565b600082821015611f0857611f08611d26565b50039056fea2646970667358221220f1982804ca22961c7b1e0054f06ed5bfcf9545c46864a93056c176f9c31d478d64736f6c634300080c0033",
}

// LagrangeABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangeMetaData.ABI instead.
var LagrangeABI = LagrangeMetaData.ABI

// LagrangeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LagrangeMetaData.Bin instead.
var LagrangeBin = LagrangeMetaData.Bin

// DeployLagrange deploys a new Ethereum contract, binding an instance of Lagrange to it.
func DeployLagrange(auth *bind.TransactOpts, backend bind.ContractBackend, _committee common.Address, _serviceManager common.Address) (common.Address, *types.Transaction, *Lagrange, error) {
	parsed, err := LagrangeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LagrangeBin), backend, _committee, _serviceManager)
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
	parsed, err := LagrangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// BLOCKHEADEREXTRADATAINDEX is a free data retrieval call binding the contract method 0x5ddbe7f5.
//
// Solidity: function BLOCK_HEADER_EXTRADATA_INDEX() view returns(uint256)
func (_Lagrange *LagrangeCaller) BLOCKHEADEREXTRADATAINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "BLOCK_HEADER_EXTRADATA_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKHEADEREXTRADATAINDEX is a free data retrieval call binding the contract method 0x5ddbe7f5.
//
// Solidity: function BLOCK_HEADER_EXTRADATA_INDEX() view returns(uint256)
func (_Lagrange *LagrangeSession) BLOCKHEADEREXTRADATAINDEX() (*big.Int, error) {
	return _Lagrange.Contract.BLOCKHEADEREXTRADATAINDEX(&_Lagrange.CallOpts)
}

// BLOCKHEADEREXTRADATAINDEX is a free data retrieval call binding the contract method 0x5ddbe7f5.
//
// Solidity: function BLOCK_HEADER_EXTRADATA_INDEX() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) BLOCKHEADEREXTRADATAINDEX() (*big.Int, error) {
	return _Lagrange.Contract.BLOCKHEADEREXTRADATAINDEX(&_Lagrange.CallOpts)
}

// BLOCKHEADERNUMBERINDEX is a free data retrieval call binding the contract method 0x14501001.
//
// Solidity: function BLOCK_HEADER_NUMBER_INDEX() view returns(uint256)
func (_Lagrange *LagrangeCaller) BLOCKHEADERNUMBERINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "BLOCK_HEADER_NUMBER_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKHEADERNUMBERINDEX is a free data retrieval call binding the contract method 0x14501001.
//
// Solidity: function BLOCK_HEADER_NUMBER_INDEX() view returns(uint256)
func (_Lagrange *LagrangeSession) BLOCKHEADERNUMBERINDEX() (*big.Int, error) {
	return _Lagrange.Contract.BLOCKHEADERNUMBERINDEX(&_Lagrange.CallOpts)
}

// BLOCKHEADERNUMBERINDEX is a free data retrieval call binding the contract method 0x14501001.
//
// Solidity: function BLOCK_HEADER_NUMBER_INDEX() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) BLOCKHEADERNUMBERINDEX() (*big.Int, error) {
	return _Lagrange.Contract.BLOCKHEADERNUMBERINDEX(&_Lagrange.CallOpts)
}

// CHAINIDARBITRUMNITRO is a free data retrieval call binding the contract method 0xfd793ed5.
//
// Solidity: function CHAIN_ID_ARBITRUM_NITRO() view returns(uint256)
func (_Lagrange *LagrangeCaller) CHAINIDARBITRUMNITRO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "CHAIN_ID_ARBITRUM_NITRO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINIDARBITRUMNITRO is a free data retrieval call binding the contract method 0xfd793ed5.
//
// Solidity: function CHAIN_ID_ARBITRUM_NITRO() view returns(uint256)
func (_Lagrange *LagrangeSession) CHAINIDARBITRUMNITRO() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDARBITRUMNITRO(&_Lagrange.CallOpts)
}

// CHAINIDARBITRUMNITRO is a free data retrieval call binding the contract method 0xfd793ed5.
//
// Solidity: function CHAIN_ID_ARBITRUM_NITRO() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) CHAINIDARBITRUMNITRO() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDARBITRUMNITRO(&_Lagrange.CallOpts)
}

// CHAINIDBASE is a free data retrieval call binding the contract method 0xf98fe1c4.
//
// Solidity: function CHAIN_ID_BASE() view returns(uint256)
func (_Lagrange *LagrangeCaller) CHAINIDBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "CHAIN_ID_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINIDBASE is a free data retrieval call binding the contract method 0xf98fe1c4.
//
// Solidity: function CHAIN_ID_BASE() view returns(uint256)
func (_Lagrange *LagrangeSession) CHAINIDBASE() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDBASE(&_Lagrange.CallOpts)
}

// CHAINIDBASE is a free data retrieval call binding the contract method 0xf98fe1c4.
//
// Solidity: function CHAIN_ID_BASE() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) CHAINIDBASE() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDBASE(&_Lagrange.CallOpts)
}

// CHAINIDMAINNET is a free data retrieval call binding the contract method 0xcec9892b.
//
// Solidity: function CHAIN_ID_MAINNET() view returns(uint256)
func (_Lagrange *LagrangeCaller) CHAINIDMAINNET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "CHAIN_ID_MAINNET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINIDMAINNET is a free data retrieval call binding the contract method 0xcec9892b.
//
// Solidity: function CHAIN_ID_MAINNET() view returns(uint256)
func (_Lagrange *LagrangeSession) CHAINIDMAINNET() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDMAINNET(&_Lagrange.CallOpts)
}

// CHAINIDMAINNET is a free data retrieval call binding the contract method 0xcec9892b.
//
// Solidity: function CHAIN_ID_MAINNET() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) CHAINIDMAINNET() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDMAINNET(&_Lagrange.CallOpts)
}

// CHAINIDOPTIMISMBEDROCK is a free data retrieval call binding the contract method 0x3c3059c7.
//
// Solidity: function CHAIN_ID_OPTIMISM_BEDROCK() view returns(uint256)
func (_Lagrange *LagrangeCaller) CHAINIDOPTIMISMBEDROCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "CHAIN_ID_OPTIMISM_BEDROCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINIDOPTIMISMBEDROCK is a free data retrieval call binding the contract method 0x3c3059c7.
//
// Solidity: function CHAIN_ID_OPTIMISM_BEDROCK() view returns(uint256)
func (_Lagrange *LagrangeSession) CHAINIDOPTIMISMBEDROCK() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDOPTIMISMBEDROCK(&_Lagrange.CallOpts)
}

// CHAINIDOPTIMISMBEDROCK is a free data retrieval call binding the contract method 0x3c3059c7.
//
// Solidity: function CHAIN_ID_OPTIMISM_BEDROCK() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) CHAINIDOPTIMISMBEDROCK() (*big.Int, error) {
	return _Lagrange.Contract.CHAINIDOPTIMISMBEDROCK(&_Lagrange.CallOpts)
}

// UPDATETYPEAMOUNTCHANGE is a free data retrieval call binding the contract method 0x23097e86.
//
// Solidity: function UPDATE_TYPE_AMOUNT_CHANGE() view returns(uint256)
func (_Lagrange *LagrangeCaller) UPDATETYPEAMOUNTCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "UPDATE_TYPE_AMOUNT_CHANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPDATETYPEAMOUNTCHANGE is a free data retrieval call binding the contract method 0x23097e86.
//
// Solidity: function UPDATE_TYPE_AMOUNT_CHANGE() view returns(uint256)
func (_Lagrange *LagrangeSession) UPDATETYPEAMOUNTCHANGE() (*big.Int, error) {
	return _Lagrange.Contract.UPDATETYPEAMOUNTCHANGE(&_Lagrange.CallOpts)
}

// UPDATETYPEAMOUNTCHANGE is a free data retrieval call binding the contract method 0x23097e86.
//
// Solidity: function UPDATE_TYPE_AMOUNT_CHANGE() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) UPDATETYPEAMOUNTCHANGE() (*big.Int, error) {
	return _Lagrange.Contract.UPDATETYPEAMOUNTCHANGE(&_Lagrange.CallOpts)
}

// UPDATETYPEREGISTER is a free data retrieval call binding the contract method 0xd742da1a.
//
// Solidity: function UPDATE_TYPE_REGISTER() view returns(uint256)
func (_Lagrange *LagrangeCaller) UPDATETYPEREGISTER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "UPDATE_TYPE_REGISTER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPDATETYPEREGISTER is a free data retrieval call binding the contract method 0xd742da1a.
//
// Solidity: function UPDATE_TYPE_REGISTER() view returns(uint256)
func (_Lagrange *LagrangeSession) UPDATETYPEREGISTER() (*big.Int, error) {
	return _Lagrange.Contract.UPDATETYPEREGISTER(&_Lagrange.CallOpts)
}

// UPDATETYPEREGISTER is a free data retrieval call binding the contract method 0xd742da1a.
//
// Solidity: function UPDATE_TYPE_REGISTER() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) UPDATETYPEREGISTER() (*big.Int, error) {
	return _Lagrange.Contract.UPDATETYPEREGISTER(&_Lagrange.CallOpts)
}

// UPDATETYPEUNREGISTER is a free data retrieval call binding the contract method 0x1d393c09.
//
// Solidity: function UPDATE_TYPE_UNREGISTER() view returns(uint256)
func (_Lagrange *LagrangeCaller) UPDATETYPEUNREGISTER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "UPDATE_TYPE_UNREGISTER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPDATETYPEUNREGISTER is a free data retrieval call binding the contract method 0x1d393c09.
//
// Solidity: function UPDATE_TYPE_UNREGISTER() view returns(uint256)
func (_Lagrange *LagrangeSession) UPDATETYPEUNREGISTER() (*big.Int, error) {
	return _Lagrange.Contract.UPDATETYPEUNREGISTER(&_Lagrange.CallOpts)
}

// UPDATETYPEUNREGISTER is a free data retrieval call binding the contract method 0x1d393c09.
//
// Solidity: function UPDATE_TYPE_UNREGISTER() view returns(uint256)
func (_Lagrange *LagrangeCallerSession) UPDATETYPEUNREGISTER() (*big.Int, error) {
	return _Lagrange.Contract.UPDATETYPEUNREGISTER(&_Lagrange.CallOpts)
}

// VerifyRawHeaderSequence is a free data retrieval call binding the contract method 0xacc41352.
//
// Solidity: function _verifyRawHeaderSequence(bytes32 latestHash, bytes[] sequence) view returns(bool)
func (_Lagrange *LagrangeCaller) VerifyRawHeaderSequence(opts *bind.CallOpts, latestHash [32]byte, sequence [][]byte) (bool, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "_verifyRawHeaderSequence", latestHash, sequence)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRawHeaderSequence is a free data retrieval call binding the contract method 0xacc41352.
//
// Solidity: function _verifyRawHeaderSequence(bytes32 latestHash, bytes[] sequence) view returns(bool)
func (_Lagrange *LagrangeSession) VerifyRawHeaderSequence(latestHash [32]byte, sequence [][]byte) (bool, error) {
	return _Lagrange.Contract.VerifyRawHeaderSequence(&_Lagrange.CallOpts, latestHash, sequence)
}

// VerifyRawHeaderSequence is a free data retrieval call binding the contract method 0xacc41352.
//
// Solidity: function _verifyRawHeaderSequence(bytes32 latestHash, bytes[] sequence) view returns(bool)
func (_Lagrange *LagrangeCallerSession) VerifyRawHeaderSequence(latestHash [32]byte, sequence [][]byte) (bool, error) {
	return _Lagrange.Contract.VerifyRawHeaderSequence(&_Lagrange.CallOpts, latestHash, sequence)
}

// CalculateBlockHash is a free data retrieval call binding the contract method 0xf44c5c71.
//
// Solidity: function calculateBlockHash(bytes rlpData) pure returns(bytes32)
func (_Lagrange *LagrangeCaller) CalculateBlockHash(opts *bind.CallOpts, rlpData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "calculateBlockHash", rlpData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateBlockHash is a free data retrieval call binding the contract method 0xf44c5c71.
//
// Solidity: function calculateBlockHash(bytes rlpData) pure returns(bytes32)
func (_Lagrange *LagrangeSession) CalculateBlockHash(rlpData []byte) ([32]byte, error) {
	return _Lagrange.Contract.CalculateBlockHash(&_Lagrange.CallOpts, rlpData)
}

// CalculateBlockHash is a free data retrieval call binding the contract method 0xf44c5c71.
//
// Solidity: function calculateBlockHash(bytes rlpData) pure returns(bytes32)
func (_Lagrange *LagrangeCallerSession) CalculateBlockHash(rlpData []byte) ([32]byte, error) {
	return _Lagrange.Contract.CalculateBlockHash(&_Lagrange.CallOpts, rlpData)
}

// CheckAndDecodeRLP is a free data retrieval call binding the contract method 0xaef524b0.
//
// Solidity: function checkAndDecodeRLP(bytes rlpData, bytes32 comparisonBlockHash) pure returns((uint256,uint256)[])
func (_Lagrange *LagrangeCaller) CheckAndDecodeRLP(opts *bind.CallOpts, rlpData []byte, comparisonBlockHash [32]byte) ([]RLPReaderRLPItem, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "checkAndDecodeRLP", rlpData, comparisonBlockHash)

	if err != nil {
		return *new([]RLPReaderRLPItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]RLPReaderRLPItem)).(*[]RLPReaderRLPItem)

	return out0, err

}

// CheckAndDecodeRLP is a free data retrieval call binding the contract method 0xaef524b0.
//
// Solidity: function checkAndDecodeRLP(bytes rlpData, bytes32 comparisonBlockHash) pure returns((uint256,uint256)[])
func (_Lagrange *LagrangeSession) CheckAndDecodeRLP(rlpData []byte, comparisonBlockHash [32]byte) ([]RLPReaderRLPItem, error) {
	return _Lagrange.Contract.CheckAndDecodeRLP(&_Lagrange.CallOpts, rlpData, comparisonBlockHash)
}

// CheckAndDecodeRLP is a free data retrieval call binding the contract method 0xaef524b0.
//
// Solidity: function checkAndDecodeRLP(bytes rlpData, bytes32 comparisonBlockHash) pure returns((uint256,uint256)[])
func (_Lagrange *LagrangeCallerSession) CheckAndDecodeRLP(rlpData []byte, comparisonBlockHash [32]byte) ([]RLPReaderRLPItem, error) {
	return _Lagrange.Contract.CheckAndDecodeRLP(&_Lagrange.CallOpts, rlpData, comparisonBlockHash)
}

// CheckCommitSignature is a free data retrieval call binding the contract method 0x75274711.
//
// Solidity: function checkCommitSignature((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) pure returns(bool)
func (_Lagrange *LagrangeCaller) CheckCommitSignature(opts *bind.CallOpts, evidence EvidenceVerifierEvidence) (bool, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "checkCommitSignature", evidence)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCommitSignature is a free data retrieval call binding the contract method 0x75274711.
//
// Solidity: function checkCommitSignature((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) pure returns(bool)
func (_Lagrange *LagrangeSession) CheckCommitSignature(evidence EvidenceVerifierEvidence) (bool, error) {
	return _Lagrange.Contract.CheckCommitSignature(&_Lagrange.CallOpts, evidence)
}

// CheckCommitSignature is a free data retrieval call binding the contract method 0x75274711.
//
// Solidity: function checkCommitSignature((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) pure returns(bool)
func (_Lagrange *LagrangeCallerSession) CheckCommitSignature(evidence EvidenceVerifierEvidence) (bool, error) {
	return _Lagrange.Contract.CheckCommitSignature(&_Lagrange.CallOpts, evidence)
}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_Lagrange *LagrangeCaller) Committee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "committee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_Lagrange *LagrangeSession) Committee() (common.Address, error) {
	return _Lagrange.Contract.Committee(&_Lagrange.CallOpts)
}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_Lagrange *LagrangeCallerSession) Committee() (common.Address, error) {
	return _Lagrange.Contract.Committee(&_Lagrange.CallOpts)
}

// GetArbAddr is a free data retrieval call binding the contract method 0x7229b160.
//
// Solidity: function getArbAddr() view returns(address)
func (_Lagrange *LagrangeCaller) GetArbAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "getArbAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetArbAddr is a free data retrieval call binding the contract method 0x7229b160.
//
// Solidity: function getArbAddr() view returns(address)
func (_Lagrange *LagrangeSession) GetArbAddr() (common.Address, error) {
	return _Lagrange.Contract.GetArbAddr(&_Lagrange.CallOpts)
}

// GetArbAddr is a free data retrieval call binding the contract method 0x7229b160.
//
// Solidity: function getArbAddr() view returns(address)
func (_Lagrange *LagrangeCallerSession) GetArbAddr() (common.Address, error) {
	return _Lagrange.Contract.GetArbAddr(&_Lagrange.CallOpts)
}

// GetCommitHash is a free data retrieval call binding the contract method 0x873b05a2.
//
// Solidity: function getCommitHash((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) pure returns(bytes32)
func (_Lagrange *LagrangeCaller) GetCommitHash(opts *bind.CallOpts, evidence EvidenceVerifierEvidence) ([32]byte, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "getCommitHash", evidence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitHash is a free data retrieval call binding the contract method 0x873b05a2.
//
// Solidity: function getCommitHash((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) pure returns(bytes32)
func (_Lagrange *LagrangeSession) GetCommitHash(evidence EvidenceVerifierEvidence) ([32]byte, error) {
	return _Lagrange.Contract.GetCommitHash(&_Lagrange.CallOpts, evidence)
}

// GetCommitHash is a free data retrieval call binding the contract method 0x873b05a2.
//
// Solidity: function getCommitHash((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) pure returns(bytes32)
func (_Lagrange *LagrangeCallerSession) GetCommitHash(evidence EvidenceVerifierEvidence) ([32]byte, error) {
	return _Lagrange.Contract.GetCommitHash(&_Lagrange.CallOpts, evidence)
}

// GetOptAddr is a free data retrieval call binding the contract method 0x655f61a2.
//
// Solidity: function getOptAddr() view returns(address)
func (_Lagrange *LagrangeCaller) GetOptAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "getOptAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOptAddr is a free data retrieval call binding the contract method 0x655f61a2.
//
// Solidity: function getOptAddr() view returns(address)
func (_Lagrange *LagrangeSession) GetOptAddr() (common.Address, error) {
	return _Lagrange.Contract.GetOptAddr(&_Lagrange.CallOpts)
}

// GetOptAddr is a free data retrieval call binding the contract method 0x655f61a2.
//
// Solidity: function getOptAddr() view returns(address)
func (_Lagrange *LagrangeCallerSession) GetOptAddr() (common.Address, error) {
	return _Lagrange.Contract.GetOptAddr(&_Lagrange.CallOpts)
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

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_Lagrange *LagrangeCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_Lagrange *LagrangeSession) ServiceManager() (common.Address, error) {
	return _Lagrange.Contract.ServiceManager(&_Lagrange.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_Lagrange *LagrangeCallerSession) ServiceManager() (common.Address, error) {
	return _Lagrange.Contract.ServiceManager(&_Lagrange.CallOpts)
}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) pure returns(bool)
func (_Lagrange *LagrangeCaller) VerifyBlockNumber(opts *bind.CallOpts, comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "verifyBlockNumber", comparisonNumber, rlpData, comparisonBlockHash, chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) pure returns(bool)
func (_Lagrange *LagrangeSession) VerifyBlockNumber(comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	return _Lagrange.Contract.VerifyBlockNumber(&_Lagrange.CallOpts, comparisonNumber, rlpData, comparisonBlockHash, chainID)
}

// VerifyBlockNumber is a free data retrieval call binding the contract method 0x62039022.
//
// Solidity: function verifyBlockNumber(uint256 comparisonNumber, bytes rlpData, bytes32 comparisonBlockHash, uint256 chainID) pure returns(bool)
func (_Lagrange *LagrangeCallerSession) VerifyBlockNumber(comparisonNumber *big.Int, rlpData []byte, comparisonBlockHash [32]byte, chainID *big.Int) (bool, error) {
	return _Lagrange.Contract.VerifyBlockNumber(&_Lagrange.CallOpts, comparisonNumber, rlpData, comparisonBlockHash, chainID)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Lagrange *LagrangeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Lagrange *LagrangeSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Lagrange *LagrangeTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner)
}

// Register is a paid mutator transaction binding the contract method 0x4fdb9291.
//
// Solidity: function register(uint32 chainID, bytes _blsPubKey, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactor) Register(opts *bind.TransactOpts, chainID uint32, _blsPubKey []byte, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "register", chainID, _blsPubKey, serveUntilBlock)
}

// Register is a paid mutator transaction binding the contract method 0x4fdb9291.
//
// Solidity: function register(uint32 chainID, bytes _blsPubKey, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeSession) Register(chainID uint32, _blsPubKey []byte, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, chainID, _blsPubKey, serveUntilBlock)
}

// Register is a paid mutator transaction binding the contract method 0x4fdb9291.
//
// Solidity: function register(uint32 chainID, bytes _blsPubKey, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactorSession) Register(chainID uint32, _blsPubKey []byte, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, chainID, _blsPubKey, serveUntilBlock)
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

// SetArbAddr is a paid mutator transaction binding the contract method 0x5364f104.
//
// Solidity: function setArbAddr(address _arb) returns()
func (_Lagrange *LagrangeTransactor) SetArbAddr(opts *bind.TransactOpts, _arb common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "setArbAddr", _arb)
}

// SetArbAddr is a paid mutator transaction binding the contract method 0x5364f104.
//
// Solidity: function setArbAddr(address _arb) returns()
func (_Lagrange *LagrangeSession) SetArbAddr(_arb common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.SetArbAddr(&_Lagrange.TransactOpts, _arb)
}

// SetArbAddr is a paid mutator transaction binding the contract method 0x5364f104.
//
// Solidity: function setArbAddr(address _arb) returns()
func (_Lagrange *LagrangeTransactorSession) SetArbAddr(_arb common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.SetArbAddr(&_Lagrange.TransactOpts, _arb)
}

// SetOptAddr is a paid mutator transaction binding the contract method 0x3df379a9.
//
// Solidity: function setOptAddr(address _opt) returns()
func (_Lagrange *LagrangeTransactor) SetOptAddr(opts *bind.TransactOpts, _opt common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "setOptAddr", _opt)
}

// SetOptAddr is a paid mutator transaction binding the contract method 0x3df379a9.
//
// Solidity: function setOptAddr(address _opt) returns()
func (_Lagrange *LagrangeSession) SetOptAddr(_opt common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.SetOptAddr(&_Lagrange.TransactOpts, _opt)
}

// SetOptAddr is a paid mutator transaction binding the contract method 0x3df379a9.
//
// Solidity: function setOptAddr(address _opt) returns()
func (_Lagrange *LagrangeTransactorSession) SetOptAddr(_opt common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.SetOptAddr(&_Lagrange.TransactOpts, _opt)
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

// UploadEvidence is a paid mutator transaction binding the contract method 0xba42f69e.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) returns()
func (_Lagrange *LagrangeTransactor) UploadEvidence(opts *bind.TransactOpts, evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "uploadEvidence", evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0xba42f69e.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) returns()
func (_Lagrange *LagrangeSession) UploadEvidence(evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.Contract.UploadEvidence(&_Lagrange.TransactOpts, evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0xba42f69e.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes) evidence) returns()
func (_Lagrange *LagrangeTransactorSession) UploadEvidence(evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.Contract.UploadEvidence(&_Lagrange.TransactOpts, evidence)
}

// LagrangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lagrange contract.
type LagrangeInitializedIterator struct {
	Event *LagrangeInitialized // Event containing the contract specifics and raw log

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
func (it *LagrangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeInitialized)
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
		it.Event = new(LagrangeInitialized)
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
func (it *LagrangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeInitialized represents a Initialized event raised by the Lagrange contract.
type LagrangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrange *LagrangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LagrangeInitializedIterator, error) {

	logs, sub, err := _Lagrange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LagrangeInitializedIterator{contract: _Lagrange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lagrange *LagrangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LagrangeInitialized) (event.Subscription, error) {

	logs, sub, err := _Lagrange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeInitialized)
				if err := _Lagrange.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Lagrange *LagrangeFilterer) ParseInitialized(log types.Log) (*LagrangeInitialized, error) {
	event := new(LagrangeInitialized)
	if err := _Lagrange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
