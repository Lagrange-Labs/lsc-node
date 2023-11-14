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
	AttestBlockHeader           []byte
	SigProof                    []byte
	AggProof                    []byte
}

// LagrangeMetaData contains all meta data concerning the Lagrange contract.
var LagrangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"_committee\",\"type\":\"address\"},{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"UploadEvidence\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_AMOUNT_CHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_REGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_UNREGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evidenceVerifier\",\"outputs\":[{\"internalType\":\"contractEvidenceVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractISlashingAggregateVerifierTriage\",\"name\":\"_AggVerify\",\"type\":\"address\"},{\"internalType\":\"contractEvidenceVerifier\",\"name\":\"_evidenceVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"attestBlockHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggProof\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"uploadEvidence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100f55760003560e01c80639aa8080411610097578063d742da1a11610066578063d742da1a146101d5578063d84a602f146101dd578063d864e740146101f0578063f2fde38b1461021757600080fd5b80639aa80804146101945780639e00be26146101a7578063aff5edb1146101ba578063c0c53b8b146101c257600080fd5b80632e94d67b116100d35780632e94d67b146101325780633998fdd314610145578063715018a6146101845780638da5cb5b1461018c57600080fd5b80630512d04c146100fa5780631d393c091461010f57806323097e861461012a575b600080fd5b61010d6101083660046111b9565b61022a565b005b610117600381565b6040519081526020015b60405180910390f35b610117600281565b61010d6101403660046111b9565b6102ad565b61016c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610121565b61010d6102fb565b61016c61030f565b61010d6101a23660046111dd565b610328565b61010d6101b5366004611308565b61070e565b61010d6108c4565b61010d6101d036600461137a565b610a36565b610117600181565b60665461016c906001600160a01b031681565b61016c7f000000000000000000000000000000000000000000000000000000000000000081565b61010d6102253660046113c5565b610b7b565b604051630e9f564b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630e9f564b9061027890339085906004016113e2565b600060405180830381600087803b15801561029257600080fd5b505af11580156102a6573d6000803e3d6000fd5b5050505050565b604051633588e1c760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636b11c38e9061027890339085906004016113e2565b610303610bf4565b61030d6000610c53565b565b60006103236033546001600160a01b031690565b905090565b60006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344a5c4bf61036660208501856113c5565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af11580156103ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d09190611401565b63ffffffff16116104285760405162461bcd60e51b815260206004820152601e60248201527f546865206f70657261746f72206973206e6f742072656769737465726564000060448201526064015b60405180910390fd5b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344f5b6b461046460208401846113c5565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af11580156104aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ce919061142e565b1561051b5760405162461bcd60e51b815260206004820152601760248201527f546865206f70657261746f7220697320736c6173686564000000000000000000604482015260640161041f565b606654604051630863d0eb60e01b81526001600160a01b0390911690630863d0eb9061054b9084906004016114bf565b602060405180830381865afa158015610568573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058c919061142e565b6105e45760405162461bcd60e51b815260206004820152602360248201527f54686520636f6d6d6974207369676e6174757265206973206e6f7420636f72726044820152621958dd60ea1b606482015260840161041f565b6105f56105f082611635565b610ca5565b61060e5761060e61060960208301836113c5565b610f58565b80602001358160400135141561062e5761062e61060960208301836113c5565b6106626080820135606083013560c084013560a085013561010086013561065d610180880161016089016111b9565b61100d565b6106765761067661060960208301836113c5565b7fa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f6106a460208301836113c5565b6020830135606084013560a085013560e08601356101008701356106cc610120890189611799565b6106da6101408b018b611799565b6106ec6101808d016101608e016111b9565b6040516107039b9a999897969594939291906117e0565b60405180910390a150565b81516060146107855760405162461bcd60e51b815260206004820152603d60248201527f4c616772616e6765536572766963653a20496e617070726f7072696174656c7960448201527f20707265666f726d617474656420424c53207075626c6963206b65792e000000606482015260840161041f565b604051634edd246960e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634edd2469906107d5903390869086906004016118ae565b600060405180830381600087803b1580156107ef57600080fd5b505af1158015610803573d6000803e3d6000fd5b505060405163175d320560e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063175d3205915061085590339085906004016113e2565b600060405180830381600087803b15801561086f57600080fd5b505af1158015610883573d6000803e3d6000fd5b505050507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a33826040516108b89291906113e2565b60405180910390a15050565b6040516319a74c5f60e01b815233600482015260009081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906319a74c5f9060240160408051808303816000875af115801561092e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095291906118e8565b91509150816109b25760405162461bcd60e51b815260206004820152602660248201527f546865206f70657261746f72206973206e6f742061626c6520746f206465726560448201526533b4b9ba32b960d11b606482015260840161041f565b6040516307fd5de760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630ffabbce90610a0090339085906004016113e2565b600060405180830381600087803b158015610a1a57600080fd5b505af1158015610a2e573d6000803e3d6000fd5b505050505050565b600054610100900460ff1615808015610a565750600054600160ff909116105b80610a705750303b158015610a70575060005460ff166001145b610ad35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161041f565b6000805460ff191660011790558015610af6576000805461ff0019166101001790555b610aff84610c53565b606580546001600160a01b038086166001600160a01b03199283161790925560668054928516929091169190911790558015610b75576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610b83610bf4565b6001600160a01b038116610be85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161041f565b610bf181610c53565b50565b33610bfd61030f565b6001600160a01b03161461030d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041f565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b61016081015160e082015160405163def9e7d560e01b815263ffffffff9092166004830152602482015260009081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063def9e7d5906044016080604051808303816000875af1158015610d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4b9190611914565b506065546101c0850151606086015160a087015160208089015160e08a01516101608b01519288015160405163ef5ef4c360e01b81529899506001600160a01b039097169763ef5ef4c397610da69796959491600401611984565b6020604051808303816000875af1158015610dc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de9919061142e565b610e415760405162461bcd60e51b815260206004820152602360248201527f4167677265676174652070726f6f6620766572696669636174696f6e206661696044820152621b195960ea1b606482015260840161041f565b8251604051634cce1da560e11b81526001600160a01b0391821660048201526000917f0000000000000000000000000000000000000000000000000000000000000000169063999c3b4a906024016000604051808303816000875af1158015610eae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ed691908101906119ce565b6066546040516302a42f7160e61b81529192506000916001600160a01b039091169063a90bdc4090610f0e9088908690600401611a3c565b602060405180830381865afa158015610f2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4f919061142e565b95945050505050565b604051630e323b9960e21b81526001600160a01b0382811660048301527f000000000000000000000000000000000000000000000000000000000000000016906338c8ee6490602401600060405180830381600087803b158015610fbb57600080fd5b505af1158015610fcf573d6000803e3d6000fd5b50506040516001600160a01b03841681527fd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a306892506020019050610703565b60405163def9e7d560e01b815263ffffffff8216600482015260248101839052600090819081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063def9e7d5906044016080604051808303816000875af1158015611087573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ab9190611914565b8151919350915089146111185760405162461bcd60e51b815260206004820152602f60248201527f5265666572656e63652063757272656e7420636f6d6d697474656520726f6f7460448201526e39903237903737ba1036b0ba31b41760891b606482015260840161041f565b86811461117c5760405162461bcd60e51b815260206004820152602c60248201527f5265666572656e6365206e65787420636f6d6d697474656520726f6f7473206460448201526b37903737ba1036b0ba31b41760a11b606482015260840161041f565b888814801561118a57508686145b9998505050505050505050565b63ffffffff81168114610bf157600080fd5b80356111b481611197565b919050565b6000602082840312156111cb57600080fd5b81356111d681611197565b9392505050565b6000602082840312156111ef57600080fd5b813567ffffffffffffffff81111561120657600080fd5b82016101e081850312156111d657600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516101e0810167ffffffffffffffff8111828210171561125357611253611219565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561128257611282611219565b604052919050565b600067ffffffffffffffff8211156112a4576112a4611219565b50601f01601f191660200190565b600082601f8301126112c357600080fd5b81356112d66112d18261128a565b611259565b8181528460208386010111156112eb57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561131b57600080fd5b823567ffffffffffffffff81111561133257600080fd5b61133e858286016112b2565b925050602083013561134f81611197565b809150509250929050565b6001600160a01b0381168114610bf157600080fd5b80356111b48161135a565b60008060006060848603121561138f57600080fd5b833561139a8161135a565b925060208401356113aa8161135a565b915060408401356113ba8161135a565b809150509250925092565b6000602082840312156113d757600080fd5b81356111d68161135a565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60006020828403121561141357600080fd5b81516111d681611197565b805180151581146111b457600080fd5b60006020828403121561144057600080fd5b6111d68261141e565b6000808335601e1984360301811261146057600080fd5b830160208101925035905067ffffffffffffffff81111561148057600080fd5b80360383131561148f57600080fd5b9250929050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526114e0602082016114d38461136f565b6001600160a01b03169052565b602082013560408201526040820135606082015260608201356080820152608082013560a082015260a082013560c082015260c082013560e0820152600061010060e084013581840152610120818501358185015261154181860186611449565b925090506101e0610140818187015261155f61020087018585611496565b935061156d81880188611449565b93509050601f1961016081888703018189015261158b868685611496565b9550611598818a016111a9565b94505061018091506115b18288018563ffffffff169052565b6115bd82890189611449565b945091506101a08188870301818901526115d8868685611496565b95506115e6818a018a611449565b95509250506101c0818887030181890152611602868685611496565b9550611610818a018a611449565b95509250508087860301838801525061162a848483611496565b979650505050505050565b60006101e0823603121561164857600080fd5b61165061122f565b6116598361136f565b81526020830135602082015260408301356040820152606083013560608201526080830135608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152506101208084013567ffffffffffffffff808211156116ca57600080fd5b6116d6368388016112b2565b838501526101409250828601359150808211156116f257600080fd5b6116fe368388016112b2565b8385015261016092506117128387016111a9565b8385015261018092508286013591508082111561172e57600080fd5b61173a368388016112b2565b838501526101a092508286013591508082111561175657600080fd5b611762368388016112b2565b838501526101c092508286013591508082111561177e57600080fd5b5061178b368287016112b2565b918301919091525092915050565b6000808335601e198436030181126117b057600080fd5b83018035915067ffffffffffffffff8211156117cb57600080fd5b60200191503681900382131561148f57600080fd5b600061012060018060a01b038e1683528c60208401528b60408401528a60608401528960808401528860a08401528060c0840152611821818401888a611496565b905082810360e0840152611836818688611496565b91505063ffffffff83166101008301529c9b505050505050505050505050565b60005b83811015611871578181015183820152602001611859565b83811115610b755750506000910152565b6000815180845261189a816020860160208601611856565b601f01601f19169290920160200192915050565b6001600160a01b03841681526060602082018190526000906118d290830185611882565b905063ffffffff83166040830152949350505050565b600080604083850312156118fb57600080fd5b6119048361141e565b9150602083015190509250929050565b600080828403608081121561192857600080fd5b606081121561193657600080fd5b506040516060810181811067ffffffffffffffff8211171561195a5761195a611219565b60409081528451825260208086015190830152848101519082015260609093015192949293505050565b60e08152600061199760e083018a611882565b60208301989098525060408101959095526060850193909352608084019190915263ffffffff1660a083015260c090910152919050565b6000602082840312156119e057600080fd5b815167ffffffffffffffff8111156119f757600080fd5b8201601f81018413611a0857600080fd5b8051611a166112d18261128a565b818152856020838501011115611a2b57600080fd5b610f4f826020830160208601611856565b60408152611a566040820184516001600160a01b03169052565b6020830151606082015260408301516080820152606083015160a0820152608083015160c082015260a083015160e0820152600060c0840151610100818185015260e08601519150610120828186015281870151925061014091508282860152808701519250506101e06101608181870152611ad6610220870185611882565b9350828801519250603f19610180818887030181890152611af78686611882565b9550828a015194506101a09250611b158389018663ffffffff169052565b808a01519450506101c0818887030181890152611b328686611882565b9550828a01519450818887030184890152611b4d8686611882565b9550808a01519450508087860301610200880152505050611b6e8282611882565b9150508281036020840152610f4f818561188256fea26469706673582212203f6a2494d3912baf0efb0c0fc816d3257fdfa098a2547907a3097b82ca198eec64736f6c634300080c0033",
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

// EvidenceVerifier is a free data retrieval call binding the contract method 0xd84a602f.
//
// Solidity: function evidenceVerifier() view returns(address)
func (_Lagrange *LagrangeCaller) EvidenceVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lagrange.contract.Call(opts, &out, "evidenceVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EvidenceVerifier is a free data retrieval call binding the contract method 0xd84a602f.
//
// Solidity: function evidenceVerifier() view returns(address)
func (_Lagrange *LagrangeSession) EvidenceVerifier() (common.Address, error) {
	return _Lagrange.Contract.EvidenceVerifier(&_Lagrange.CallOpts)
}

// EvidenceVerifier is a free data retrieval call binding the contract method 0xd84a602f.
//
// Solidity: function evidenceVerifier() view returns(address)
func (_Lagrange *LagrangeCallerSession) EvidenceVerifier() (common.Address, error) {
	return _Lagrange.Contract.EvidenceVerifier(&_Lagrange.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _AggVerify, address _evidenceVerifier) returns()
func (_Lagrange *LagrangeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _AggVerify common.Address, _evidenceVerifier common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "initialize", initialOwner, _AggVerify, _evidenceVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _AggVerify, address _evidenceVerifier) returns()
func (_Lagrange *LagrangeSession) Initialize(initialOwner common.Address, _AggVerify common.Address, _evidenceVerifier common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner, _AggVerify, _evidenceVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _AggVerify, address _evidenceVerifier) returns()
func (_Lagrange *LagrangeTransactorSession) Initialize(initialOwner common.Address, _AggVerify common.Address, _evidenceVerifier common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner, _AggVerify, _evidenceVerifier)
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

// UploadEvidence is a paid mutator transaction binding the contract method 0x9aa80804.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes,bytes,bytes) evidence) returns()
func (_Lagrange *LagrangeTransactor) UploadEvidence(opts *bind.TransactOpts, evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "uploadEvidence", evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0x9aa80804.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes,bytes,bytes) evidence) returns()
func (_Lagrange *LagrangeSession) UploadEvidence(evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.Contract.UploadEvidence(&_Lagrange.TransactOpts, evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0x9aa80804.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes,bytes,bytes) evidence) returns()
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
