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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"_committee\",\"type\":\"address\"},{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"UploadEvidence\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_HEADER_EXTRADATA_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_HEADER_NUMBER_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_ARBITRUM_NITRO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_MAINNET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_OPTIMISM_BEDROCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_AMOUNT_CHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_REGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_UNREGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"latestHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"sequence\",\"type\":\"bytes[]\"}],\"name\":\"_verifyRawHeaderSequence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"}],\"name\":\"calculateBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"}],\"name\":\"checkAndDecodeRLP\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPtr\",\"type\":\"uint256\"}],\"internalType\":\"structRLPReader.RLPItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"rawBlockHeader\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"checkCommitSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArbAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"rawBlockHeader\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractArbitrumVerifier\",\"name\":\"_arb\",\"type\":\"address\"}],\"name\":\"setArbAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptimismVerifier\",\"name\":\"_opt\",\"type\":\"address\"}],\"name\":\"setOptAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"rawBlockHeader\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"uploadEvidence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"comparisonNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"comparisonBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"verifyBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101da5760003560e01c8063873b05a211610104578063c4d66de8116100a2578063f2fde38b11610071578063f2fde38b14610400578063f44c5c7114610413578063f98fe1c41461042c578063fd793ed51461043657600080fd5b8063c4d66de8146103be578063cec9892b146103d1578063d742da1a146103d1578063d864e740146103d957600080fd5b8063acc41352116100de578063acc4135214610370578063aef524b014610383578063aff5edb1146103a3578063ba42f69e146103ab57600080fd5b8063873b05a2146103425780638da5cb5b146103555780639e00be261461035d57600080fd5b80633df379a91161017c578063655f61a21161014b578063655f61a214610305578063715018a6146103165780637229b1601461031e578063752747111461032f57600080fd5b80633df379a91461027a5780635364f104146102aa5780635ddbe7f5146102da57806362039022146102e257600080fd5b806323097e86116101b857806323097e86146102175780632e94d67b1461021f5780633998fdd3146102325780633c3059c71461027157600080fd5b80630512d04c146101df57806314501001146101f45780631d393c091461020f575b600080fd5b6101f26101ed366004611b42565b610440565b005b6101fc600881565b6040519081526020015b60405180910390f35b6101fc600381565b6101fc600281565b6101f261022d366004611b42565b6104c3565b6102597f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610206565b6101fc6101a481565b6101f2610288366004611b74565b606580546001600160a01b0319166001600160a01b0392909216919091179055565b6101f26102b8366004611b74565b606680546001600160a01b0319166001600160a01b0392909216919091179055565b6101fc600c81565b6102f56102f0366004611c34565b610511565b6040519015158152602001610206565b6065546001600160a01b0316610259565b6101f261052a565b6066546001600160a01b0316610259565b6102f561033d366004611c8b565b61053e565b6101fc610350366004611c8b565b6105be565b61025961062e565b6101f261036b366004611cc7565b610647565b6102f561037e366004611d19565b610802565b610396610391366004611d98565b61093c565b6040516102069190611ddd565b6101f26109f1565b6101f26103b9366004611c8b565b610b63565b6101f26103cc366004611b74565b610ee8565b6101fc600181565b6102597f000000000000000000000000000000000000000000000000000000000000000081565b6101f261040e366004611b74565b610ff7565b6101fc610421366004611e2c565b805160209091012090565b6101fc62014a3381565b6101fc62066eed81565b604051630e9f564b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630e9f564b9061048e9033908590600401611e61565b600060405180830381600087803b1580156104a857600080fd5b505af11580156104bc573d6000803e3d6000fd5b5050505050565b604051633588e1c760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636b11c38e9061048e9033908590600401611e61565b60008061052086868686611070565b9695505050505050565b6105326110b7565b61053c6000611116565b565b60008061054a836105be565b9050600061059a82610560610140870187611e80565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061116892505050565b90506105a96020850185611b74565b6001600160a01b039182169116149392505050565b60006020820135606083013560a084013560e08501356101008601356105e8610120880188611e80565b6105fa6101808a016101608b01611b42565b604051602001610611989796959493929190611ec7565b604051602081830303815290604052805190602001209050919050565b60006106426033546001600160a01b031690565b905090565b81516060146106c35760405162461bcd60e51b815260206004820152603d60248201527f4c616772616e6765536572766963653a20496e617070726f7072696174656c7960448201527f20707265666f726d617474656420424c53207075626c6963206b65792e00000060648201526084015b60405180910390fd5b604051634edd246960e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634edd24699061071390339086908690600401611f13565b600060405180830381600087803b15801561072d57600080fd5b505af1158015610741573d6000803e3d6000fd5b505060405163175d320560e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063175d320591506107939033908590600401611e61565b600060405180830381600087803b1580156107ad57600080fd5b505af11580156107c1573d6000803e3d6000fd5b505050507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a33826040516107f6929190611e61565b60405180910390a15050565b60008060005b8381101561091d57600061087b61087687878581811061082a5761082a611f86565b905060200281019061083c9190611e80565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061118c92505050565b6111b9565b905060008160008151811061089257610892611f86565b6020026020010151905060006108a7826112cf565b905083158015906108b85750848114155b156108cb57600095505050505050610935565b8787858181106108dd576108dd611f86565b90506020028101906108ef9190611e80565b6040516108fd929190611f9c565b60405180910390209450505050808061091590611fc2565b915050610808565b5080851461092f576000915050610935565b60019150505b9392505050565b815160208301206060908281146109b25760405162461bcd60e51b815260206004820152603460248201527f48617368206f6620524c5020646174612064697665726765732066726f6d20636044820152730dedae0c2e4d2e6dedc40c4d8dec6d640d0c2e6d60631b60648201526084016106ba565b60006109e86108768660408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b95945050505050565b6040516319a74c5f60e01b815233600482015260009081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906319a74c5f9060240160408051808303816000875af1158015610a5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7f9190611ff2565b9150915081610adf5760405162461bcd60e51b815260206004820152602660248201527f546865206f70657261746f72206973206e6f742061626c6520746f206465726560448201526533b4b9ba32b960d11b60648201526084016106ba565b6040516307fd5de760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630ffabbce90610b2d9033908590600401611e61565b600060405180830381600087803b158015610b4757600080fd5b505af1158015610b5b573d6000803e3d6000fd5b505050505050565b60006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344a5c4bf610ba16020850185611b74565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af1158015610be7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0b919061201e565b63ffffffff1611610c5e5760405162461bcd60e51b815260206004820152601e60248201527f546865206f70657261746f72206973206e6f742072656769737465726564000060448201526064016106ba565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344f5b6b4610c9a6020840184611b74565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af1158015610ce0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d04919061203b565b15610d515760405162461bcd60e51b815260206004820152601760248201527f546865206f70657261746f7220697320736c617368656400000000000000000060448201526064016106ba565b610d5a8161053e565b610db25760405162461bcd60e51b815260206004820152602360248201527f54686520636f6d6d6974207369676e6174757265206973206e6f7420636f72726044820152621958dd60ea1b60648201526084016106ba565b610def6040820135602083013560e0840135610dd2610180860186611e80565b610de461018088016101608901611b42565b63ffffffff1661131d565b610e0857610e08610e036020830183611b74565b611379565b610e3c6080820135606083013560c084013560a0850135610100860135610e3761018088016101608901611b42565b61142e565b610e5057610e50610e036020830183611b74565b7fa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f610e7e6020830183611b74565b6020830135606084013560a085013560e0860135610100870135610ea6610120890189611e80565b610eb46101408b018b611e80565b610ec66101808d016101608e01611b42565b604051610edd9b9a9998979695949392919061207f565b60405180910390a150565b600054610100900460ff1615808015610f085750600054600160ff909116105b80610f225750303b158015610f22575060005460ff166001145b610f855760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106ba565b6000805460ff191660011790558015610fa8576000805461ff0019166101001790555b610fb182611116565b8015610ff3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016107f6565b5050565b610fff6110b7565b6001600160a01b0381166110645760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106ba565b61106d81611116565b50565b60008061107d858561093c565b905060008160088151811061109457611094611f86565b6020026020010151905060006110a9826112cf565b909714979650505050505050565b336110c061062e565b6001600160a01b03161461053c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106ba565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080600061117785856115b8565b9150915061118481611628565b509392505050565b60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60606111c4826117e3565b6111cd57600080fd5b60006111d88361181c565b905060008167ffffffffffffffff8111156111f5576111f5611b91565b60405190808252806020026020018201604052801561123a57816020015b60408051808201909152600080825260208201528152602001906001900390816112135790505b509050600061124c856020015161189f565b856020015161125b91906120f5565b90506000805b848110156112c4576112728361191a565b915060405180604001604052808381526020018481525084828151811061129b5761129b611f86565b60209081029190910101526112b082846120f5565b9250806112bc81611fc2565b915050611261565b509195945050505050565b8051600090158015906112e457508151602110155b6112ed57600080fd5b6000806112f9846119c3565b8151919350915060208210156113155760208290036101000a90045b949350505050565b60006113638585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c92508791506105119050565b801561136e57508686145b979650505050505050565b604051630e323b9960e21b81526001600160a01b0382811660048301527f000000000000000000000000000000000000000000000000000000000000000016906338c8ee6490602401600060405180830381600087803b1580156113dc57600080fd5b505af11580156113f0573d6000803e3d6000fd5b50506040516001600160a01b03841681527fd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a306892506020019050610edd565b60405163def9e7d560e01b815263ffffffff8216600482015260248101839052600090819081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063def9e7d5906044016080604051808303816000875af11580156114a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114cc919061210d565b8151919350915089146115395760405162461bcd60e51b815260206004820152602f60248201527f5265666572656e63652063757272656e7420636f6d6d697474656520726f6f7460448201526e39903237903737ba1036b0ba31b41760891b60648201526084016106ba565b86811461159d5760405162461bcd60e51b815260206004820152602c60248201527f5265666572656e6365206e65787420636f6d6d697474656520726f6f7473206460448201526b37903737ba1036b0ba31b41760a11b60648201526084016106ba565b88881480156115ab57508686145b9998505050505050505050565b6000808251604114156115ef5760208301516040840151606085015160001a6115e387828585611a0a565b94509450505050611621565b825160401415611619576020830151604084015161160e868383611af7565b935093505050611621565b506000905060025b9250929050565b600081600481111561163c5761163c61217d565b14156116455750565b60018160048111156116595761165961217d565b14156116a75760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016106ba565b60028160048111156116bb576116bb61217d565b14156117095760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016106ba565b600381600481111561171d5761171d61217d565b14156117765760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016106ba565b600481600481111561178a5761178a61217d565b141561106d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016106ba565b80516000906117f457506000919050565b6020820151805160001a9060c0821015611812575060009392505050565b5060019392505050565b805160009061182d57506000919050565b60008061183d846020015161189f565b846020015161184c91906120f5565b905060008460000151856020015161186491906120f5565b90505b80821015611896576118788261191a565b61188290836120f5565b91508261188e81611fc2565b935050611867565b50909392505050565b8051600090811a60808110156118b85750600092915050565b60b88110806118d3575060c081108015906118d3575060f881105b156118e15750600192915050565b60c081101561190e576118f6600160b8612193565b6119039060ff16826121b6565b6109359060016120f5565b6118f6600160f8612193565b80516000908190811a608081101561193557600191506119bc565b60b881101561195b576119496080826121b6565b6119549060016120f5565b91506119bc565b60c08110156119885760b78103600185019450806020036101000a855104600182018101935050506119bc565b60f881101561199c5761194960c0826121b6565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b60008060006119d5846020015161189f565b905060008185602001516119e991906120f5565b905060008286600001516119fd91906121b6565b9196919550909350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611a415750600090506003611aee565b8460ff16601b14158015611a5957508460ff16601c14155b15611a6a5750600090506004611aee565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611abe573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611ae757600060019250925050611aee565b9150600090505b94509492505050565b6000806001600160ff1b03831681611b1460ff86901c601b6120f5565b9050611b2287828885611a0a565b935093505050935093915050565b63ffffffff8116811461106d57600080fd5b600060208284031215611b5457600080fd5b813561093581611b30565b6001600160a01b038116811461106d57600080fd5b600060208284031215611b8657600080fd5b813561093581611b5f565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611bb857600080fd5b813567ffffffffffffffff80821115611bd357611bd3611b91565b604051601f8301601f19908116603f01168101908282118183101715611bfb57611bfb611b91565b81604052838152866020858801011115611c1457600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060808587031215611c4a57600080fd5b84359350602085013567ffffffffffffffff811115611c6857600080fd5b611c7487828801611ba7565b949794965050505060408301359260600135919050565b600060208284031215611c9d57600080fd5b813567ffffffffffffffff811115611cb457600080fd5b82016101a0818503121561093557600080fd5b60008060408385031215611cda57600080fd5b823567ffffffffffffffff811115611cf157600080fd5b611cfd85828601611ba7565b9250506020830135611d0e81611b30565b809150509250929050565b600080600060408486031215611d2e57600080fd5b83359250602084013567ffffffffffffffff80821115611d4d57600080fd5b818601915086601f830112611d6157600080fd5b813581811115611d7057600080fd5b8760208260051b8501011115611d8557600080fd5b6020830194508093505050509250925092565b60008060408385031215611dab57600080fd5b823567ffffffffffffffff811115611dc257600080fd5b611dce85828601611ba7565b95602094909401359450505050565b602080825282518282018190526000919060409081850190868401855b82811015611e1f57815180518552860151868501529284019290850190600101611dfa565b5091979650505050505050565b600060208284031215611e3e57600080fd5b813567ffffffffffffffff811115611e5557600080fd5b61131584828501611ba7565b6001600160a01b0392909216825263ffffffff16602082015260400190565b6000808335601e19843603018112611e9757600080fd5b83018035915067ffffffffffffffff821115611eb257600080fd5b60200191503681900382131561162157600080fd5b888152876020820152866040820152856060820152846080820152828460a083013760e09190911b6001600160e01b03191660a0919092019081019190915260a4019695505050505050565b60018060a01b038416815260006020606081840152845180606085015260005b81811015611f4f57868101830151858201608001528201611f33565b81811115611f61576000608083870101525b5063ffffffff9490941660408401525050601f91909101601f19160160800192915050565b634e487b7160e01b600052603260045260246000fd5b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415611fd657611fd6611fac565b5060010190565b80518015158114611fed57600080fd5b919050565b6000806040838503121561200557600080fd5b61200e83611fdd565b9150602083015190509250929050565b60006020828403121561203057600080fd5b815161093581611b30565b60006020828403121561204d57600080fd5b61093582611fdd565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061012060018060a01b038e1683528c60208401528b60408401528a60608401528960808401528860a08401528060c08401526120c0818401888a612056565b905082810360e08401526120d5818688612056565b91505063ffffffff83166101008301529c9b505050505050505050505050565b6000821982111561210857612108611fac565b500190565b600080828403608081121561212157600080fd5b606081121561212f57600080fd5b506040516060810181811067ffffffffffffffff8211171561215357612153611b91565b60409081528451825260208086015190830152848101519082015260609093015192949293505050565b634e487b7160e01b600052602160045260246000fd5b600060ff821660ff8416808210156121ad576121ad611fac565b90039392505050565b6000828210156121c8576121c8611fac565b50039056fea26469706673582212201650d34ceb56a0394a70526f8430c74e80858cfdad6787aa4297edddcff394c664736f6c634300080c0033",
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

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Lagrange *LagrangeTransactor) Deregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "deregister")
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Lagrange *LagrangeSession) Deregister() (*types.Transaction, error) {
	return _Lagrange.Contract.Deregister(&_Lagrange.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Lagrange *LagrangeTransactorSession) Deregister() (*types.Transaction, error) {
	return _Lagrange.Contract.Deregister(&_Lagrange.TransactOpts)
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

// Register is a paid mutator transaction binding the contract method 0x9e00be26.
//
// Solidity: function register(bytes _blsPubKey, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactor) Register(opts *bind.TransactOpts, _blsPubKey []byte, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "register", _blsPubKey, serveUntilBlock)
}

// Register is a paid mutator transaction binding the contract method 0x9e00be26.
//
// Solidity: function register(bytes _blsPubKey, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeSession) Register(_blsPubKey []byte, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, _blsPubKey, serveUntilBlock)
}

// Register is a paid mutator transaction binding the contract method 0x9e00be26.
//
// Solidity: function register(bytes _blsPubKey, uint32 serveUntilBlock) returns()
func (_Lagrange *LagrangeTransactorSession) Register(_blsPubKey []byte, serveUntilBlock uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Register(&_Lagrange.TransactOpts, _blsPubKey, serveUntilBlock)
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

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactor) Subscribe(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "subscribe", chainID)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeSession) Subscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Subscribe(&_Lagrange.TransactOpts, chainID)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactorSession) Subscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Subscribe(&_Lagrange.TransactOpts, chainID)
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

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactor) Unsubscribe(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "unsubscribe", chainID)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeSession) Unsubscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Unsubscribe(&_Lagrange.TransactOpts, chainID)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_Lagrange *LagrangeTransactorSession) Unsubscribe(chainID uint32) (*types.Transaction, error) {
	return _Lagrange.Contract.Unsubscribe(&_Lagrange.TransactOpts, chainID)
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
