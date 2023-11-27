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
	SigProof                    []byte
	AggProof                    []byte
}

// LagrangeMetaData contains all meta data concerning the Lagrange contract.
var LagrangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"_committee\",\"type\":\"address\"},{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"UploadEvidence\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_AMOUNT_CHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_REGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_TYPE_UNREGISTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evidenceVerifier\",\"outputs\":[{\"internalType\":\"contractEvidenceVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractEvidenceVerifier\",\"name\":\"_evidenceVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctCurrentCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"correctNextCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"commitSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sigProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggProof\",\"type\":\"bytes\"}],\"internalType\":\"structEvidenceVerifier.Evidence\",\"name\":\"evidence\",\"type\":\"tuple\"}],\"name\":\"uploadEvidence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100f55760003560e01c80638da5cb5b11610097578063d742da1a11610066578063d742da1a146101d5578063d84a602f146101dd578063d864e740146101f0578063f2fde38b1461021757600080fd5b80638da5cb5b1461019f5780639e00be26146101a7578063aff5edb1146101ba578063bd496305146101c257600080fd5b80632e94d67b116100d35780632e94d67b146101325780633998fdd314610145578063485cc95514610184578063715018a61461019757600080fd5b80630512d04c146100fa5780631d393c091461010f57806323097e861461012a575b600080fd5b61010d610108366004611180565b61022a565b005b610117600381565b6040519081526020015b60405180910390f35b610117600281565b61010d610140366004611180565b6102ad565b61016c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610121565b61010d6101923660046111c4565b6102fb565b61010d61042f565b61016c610443565b61010d6101b53660046112ec565b61045c565b61010d610612565b61010d6101d0366004611333565b610784565b610117600181565b60655461016c906001600160a01b031681565b61016c7f000000000000000000000000000000000000000000000000000000000000000081565b61010d61022536600461136f565b610b65565b604051630e9f564b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630e9f564b90610278903390859060040161138c565b600060405180830381600087803b15801561029257600080fd5b505af11580156102a6573d6000803e3d6000fd5b5050505050565b604051633588e1c760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636b11c38e90610278903390859060040161138c565b600054610100900460ff161580801561031b5750600054600160ff909116105b806103355750303b158015610335575060005460ff166001145b61039d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156103c0576000805461ff0019166101001790555b6103c983610bde565b606580546001600160a01b0319166001600160a01b038416179055801561042a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610437610c30565b6104416000610bde565b565b60006104576033546001600160a01b031690565b905090565b81516060146104d35760405162461bcd60e51b815260206004820152603d60248201527f4c616772616e6765536572766963653a20496e617070726f7072696174656c7960448201527f20707265666f726d617474656420424c53207075626c6963206b65792e0000006064820152608401610394565b604051634edd246960e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634edd24699061052390339086908690600401611407565b600060405180830381600087803b15801561053d57600080fd5b505af1158015610551573d6000803e3d6000fd5b505060405163175d320560e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063175d320591506105a3903390859060040161138c565b600060405180830381600087803b1580156105bd57600080fd5b505af11580156105d1573d6000803e3d6000fd5b505050507f3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a338260405161060692919061138c565b60405180910390a15050565b6040516319a74c5f60e01b815233600482015260009081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906319a74c5f9060240160408051808303816000875af115801561067c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a09190611451565b91509150816107005760405162461bcd60e51b815260206004820152602660248201527f546865206f70657261746f72206973206e6f742061626c6520746f206465726560448201526533b4b9ba32b960d11b6064820152608401610394565b6040516307fd5de760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630ffabbce9061074e903390859060040161138c565b600060405180830381600087803b15801561076857600080fd5b505af115801561077c573d6000803e3d6000fd5b505050505050565b60006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344a5c4bf6107c2602085018561136f565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af1158015610808573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082c919061147d565b63ffffffff161161087f5760405162461bcd60e51b815260206004820152601e60248201527f546865206f70657261746f72206973206e6f74207265676973746572656400006044820152606401610394565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166344f5b6b46108bb602084018461136f565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016020604051808303816000875af1158015610901573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610925919061149a565b156109725760405162461bcd60e51b815260206004820152601760248201527f546865206f70657261746f7220697320736c61736865640000000000000000006044820152606401610394565b606554604051632d0c72c560e11b81526001600160a01b0390911690635a18e58a906109a290849060040161152b565b602060405180830381865afa1580156109bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e3919061149a565b610a3b5760405162461bcd60e51b815260206004820152602360248201527f54686520636f6d6d6974207369676e6174757265206973206e6f7420636f72726044820152621958dd60ea1b6064820152608401610394565b610a4c610a4782611677565b610c8f565b610a6557610a65610a60602083018361136f565b610f1f565b806020013581604001351415610a8557610a85610a60602083018361136f565b610ab96080820135606083013560c084013560a0850135610100860135610ab461018088016101608901611180565b610fd4565b610acd57610acd610a60602083018361136f565b7fa3df44f3e14b2d57c4eed4929c8cd401795e6739ea5b89dd902f25a05fea132f610afb602083018361136f565b6020830135606084013560a085013560e0860135610100870135610b236101208901896117b3565b610b316101408b018b6117b3565b610b436101808d016101608e01611180565b604051610b5a9b9a999897969594939291906117fa565b60405180910390a150565b610b6d610c30565b6001600160a01b038116610bd25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610394565b610bdb81610bde565b50565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b33610c39610443565b6001600160a01b0316146104415760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610394565b61016081015160e082015160405163def9e7d560e01b815263ffffffff9092166004830152602482015260009081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063def9e7d5906044016080604051808303816000875af1158015610d11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d359190611870565b5060655460208201516040516328e4f4b960e01b81529293506001600160a01b03909116916328e4f4b991610d6f918791906004016119d4565b602060405180830381865afa158015610d8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db0919061149a565b610e085760405162461bcd60e51b815260206004820152602360248201527f4167677265676174652070726f6f6620766572696669636174696f6e206661696044820152621b195960ea1b6064820152608401610394565b8251604051634cce1da560e11b81526001600160a01b0391821660048201526000917f0000000000000000000000000000000000000000000000000000000000000000169063999c3b4a906024016000604051808303816000875af1158015610e75573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e9d91908101906119f6565b606554604051633dc26dcd60e11b81529192506000916001600160a01b0390911690637b84db9a90610ed59088908690600401611a64565b602060405180830381865afa158015610ef2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f16919061149a565b95945050505050565b604051630e323b9960e21b81526001600160a01b0382811660048301527f000000000000000000000000000000000000000000000000000000000000000016906338c8ee6490602401600060405180830381600087803b158015610f8257600080fd5b505af1158015610f96573d6000803e3d6000fd5b50506040516001600160a01b03841681527fd8f676e084105f4a403cee55f7a0c0aae9a015ce7a743ff68cd4e422fd4a306892506020019050610b5a565b60405163def9e7d560e01b815263ffffffff8216600482015260248101839052600090819081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063def9e7d5906044016080604051808303816000875af115801561104e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110729190611870565b8151919350915089146110df5760405162461bcd60e51b815260206004820152602f60248201527f5265666572656e63652063757272656e7420636f6d6d697474656520726f6f7460448201526e39903237903737ba1036b0ba31b41760891b6064820152608401610394565b8681146111435760405162461bcd60e51b815260206004820152602c60248201527f5265666572656e6365206e65787420636f6d6d697474656520726f6f7473206460448201526b37903737ba1036b0ba31b41760a11b6064820152608401610394565b888814801561115157508686145b9998505050505050505050565b63ffffffff81168114610bdb57600080fd5b803561117b8161115e565b919050565b60006020828403121561119257600080fd5b813561119d8161115e565b9392505050565b6001600160a01b0381168114610bdb57600080fd5b803561117b816111a4565b600080604083850312156111d757600080fd5b82356111e2816111a4565b915060208301356111f2816111a4565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040516101c0810167ffffffffffffffff81118282101715611237576112376111fd565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611266576112666111fd565b604052919050565b600067ffffffffffffffff821115611288576112886111fd565b50601f01601f191660200190565b600082601f8301126112a757600080fd5b81356112ba6112b58261126e565b61123d565b8181528460208386010111156112cf57600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156112ff57600080fd5b823567ffffffffffffffff81111561131657600080fd5b61132285828601611296565b92505060208301356111f28161115e565b60006020828403121561134557600080fd5b813567ffffffffffffffff81111561135c57600080fd5b82016101c0818503121561119d57600080fd5b60006020828403121561138157600080fd5b813561119d816111a4565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60005b838110156113c65781810151838201526020016113ae565b838111156113d5576000848401525b50505050565b600081518084526113f38160208601602086016113ab565b601f01601f19169290920160200192915050565b6001600160a01b038416815260606020820181905260009061142b908301856113db565b905063ffffffff83166040830152949350505050565b8051801515811461117b57600080fd5b6000806040838503121561146457600080fd5b61146d83611441565b9150602083015190509250929050565b60006020828403121561148f57600080fd5b815161119d8161115e565b6000602082840312156114ac57600080fd5b61119d82611441565b6000808335601e198436030181126114cc57600080fd5b830160208101925035905067ffffffffffffffff8111156114ec57600080fd5b8036038313156114fb57600080fd5b9250929050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815261154c6020820161153f846111b9565b6001600160a01b03169052565b602082013560408201526040820135606082015260608201356080820152608082013560a082015260a082013560c082015260c082013560e0820152600061010060e08401358184015261012081850135818501526115ad818601866114b5565b925090506101c061014081818701526115cb6101e087018585611502565b93506115d9818801886114b5565b93509050601f196101608188870301818901526115f7868685611502565b9550611604818a01611170565b945050610180915061161d8288018563ffffffff169052565b611629828901896114b5565b945091506101a0818887030181890152611644868685611502565b9550611652818a018a6114b5565b95509250508087860301838801525061166c848483611502565b979650505050505050565b60006101c0823603121561168a57600080fd5b611692611213565b61169b836111b9565b81526020830135602082015260408301356040820152606083013560608201526080830135608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152506101208084013567ffffffffffffffff8082111561170c57600080fd5b61171836838801611296565b8385015261014092508286013591508082111561173457600080fd5b61174036838801611296565b838501526101609250611754838701611170565b8385015261018092508286013591508082111561177057600080fd5b61177c36838801611296565b838501526101a092508286013591508082111561179857600080fd5b506117a536828701611296565b918301919091525092915050565b6000808335601e198436030181126117ca57600080fd5b83018035915067ffffffffffffffff8211156117e557600080fd5b6020019150368190038213156114fb57600080fd5b600061012060018060a01b038e1683528c60208401528b60408401528a60608401528960808401528860a08401528060c084015261183b818401888a611502565b905082810360e0840152611850818688611502565b91505063ffffffff83166101008301529c9b505050505050505050505050565b600080828403608081121561188457600080fd5b606081121561189257600080fd5b506040516060810181811067ffffffffffffffff821117156118b6576118b66111fd565b60409081528451825260208086015190830152848101519082015260609093015192949293505050565b80516001600160a01b0316825260006101c06020830151602085015260408301516040850152606083015160608501526080830151608085015260a083015160a085015260c083015160c085015260e083015160e085015261010080840151818601525061012080840151828287015261195c838701826113db565b92505050610140808401518583038287015261197883826113db565b92505050610160808401516119948287018263ffffffff169052565b505061018080840151858303828701526119ae83826113db565b925050506101a080840151858303828701526119ca83826113db565b9695505050505050565b6040815260006119e760408301856118e0565b90508260208301529392505050565b600060208284031215611a0857600080fd5b815167ffffffffffffffff811115611a1f57600080fd5b8201601f81018413611a3057600080fd5b8051611a3e6112b58261126e565b818152856020838501011115611a5357600080fd5b610f168260208301602086016113ab565b604081526000611a7760408301856118e0565b8281036020840152610f1681856113db56fea2646970667358221220e440b9b498cb5b6645f88d377441085d8e6ce33b89c911a000e0b7b08d19d13164736f6c634300080c0033",
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _evidenceVerifier) returns()
func (_Lagrange *LagrangeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _evidenceVerifier common.Address) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "initialize", initialOwner, _evidenceVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _evidenceVerifier) returns()
func (_Lagrange *LagrangeSession) Initialize(initialOwner common.Address, _evidenceVerifier common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner, _evidenceVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _evidenceVerifier) returns()
func (_Lagrange *LagrangeTransactorSession) Initialize(initialOwner common.Address, _evidenceVerifier common.Address) (*types.Transaction, error) {
	return _Lagrange.Contract.Initialize(&_Lagrange.TransactOpts, initialOwner, _evidenceVerifier)
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

// UploadEvidence is a paid mutator transaction binding the contract method 0xbd496305.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes,bytes) evidence) returns()
func (_Lagrange *LagrangeTransactor) UploadEvidence(opts *bind.TransactOpts, evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.contract.Transact(opts, "uploadEvidence", evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0xbd496305.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes,bytes) evidence) returns()
func (_Lagrange *LagrangeSession) UploadEvidence(evidence EvidenceVerifierEvidence) (*types.Transaction, error) {
	return _Lagrange.Contract.UploadEvidence(&_Lagrange.TransactOpts, evidence)
}

// UploadEvidence is a paid mutator transaction binding the contract method 0xbd496305.
//
// Solidity: function uploadEvidence((address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bytes,bytes,uint32,bytes,bytes) evidence) returns()
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
