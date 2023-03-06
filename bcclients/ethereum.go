package bcclients

import (
	context "context"
	"fmt"
	log "log"
	"math/big"
	"strings"
	"time"

	"github.com/Lagrange-Labs/Lagrange-Node/utils"
	common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
)

// Loads Ethereum client provided an endpoint URL.
func LoadEthClient(ethEndpoint string) *ethClient.Client {
	eth, err := ethClient.Dial(ethEndpoint)
	if err != nil {
		panic(err)
	}
	utils.LogMessage("Endpoint Loaded: "+ethEndpoint, utils.LOG_INFO)
	return eth
}

func LoadEthClientMulti(ethEndpoint string) []*ethClient.Client {
	ethAttestAddrs := strings.Split(ethEndpoint, ",")
	var ethAttestClients []*ethClient.Client
	_ = ethAttestClients
	var ethAttest *ethClient.Client
	_ = ethAttest
	for i, addr := range ethAttestAddrs {
		_ = i
		ethAttest := LoadEthClient(addr)
		ethAttestClients = append(ethAttestClients, ethAttest)
	}
	return ethAttestClients
}

// Loads RPC client provided an endpoint URL.
func LoadRpcClient(ethEndpoint string) *rpc.Client {
	rpc, err := rpc.DialHTTP(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer rpc.Close()
	return rpc
}

// Reference function testing a raw RPC call.
func RpcCall(rpc *rpc.Client, To string, Data string) {
	type request struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}

	var result string

	req := request{To, Data}
	if err := rpc.Call(&result, "eth_call", req, "latest"); err != nil {
		log.Fatal(err)
	}

	owner := common.HexToAddress(result)
	fmt.Printf("RPC Result: %s\n", owner.Hex()) // 0x281017b4E914b79371d62518b17693B36c7a221e
}

// Reference function testing retrieval of Ethereum transaction and balance.
func EthTest(eth *ethClient.Client) {
	ctx := context.Background()
	tx, pending, _ := eth.TransactionByHash(ctx, common.HexToHash("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	if !pending {
		fmt.Println("tx:", tx)
	}

	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := eth.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance:", balance) // 25893180161173005034
}

// Returns Keccak hash of string as bytes.
func KeccakHash(stateRootStr string) []byte {
	return crypto.Keccak256([]byte(stateRootStr))
}

// Returns Keccak hash of string as hex string with '0x' prefix.
func KeccakHashString(stateRootStr string) string {
	return hexutil.Encode(KeccakHash(stateRootStr))
}

func GetNonce(client *ethClient.Client, fromAddress common.Address) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("Nonce:", nonce)
	return nonce
}

// Requests and returns network gas price.
func GetGasPrice(client *ethClient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

// Test function - mines a new block each second.
func MineTest(rpc *rpc.Client) {
	for {
		MineBlocks(rpc, 1)
		time.Sleep(1 * time.Second)
	}
}

// Test function - mines n blocks on Hardhat node
func MineBlocks(rpc *rpc.Client, num int) {
	var hex hexutil.Bytes
	for i := 0; i < num; i++ {
		rpc.Call(&hex, "evm_mine")
	}
}
