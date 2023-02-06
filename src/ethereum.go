package main

import (
	"fmt"
	"time"
	"strings"
	"math/big"
	common "github.com/ethereum/go-ethereum/common"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
	log "log"
	context "context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	accounts "github.com/ethereum/go-ethereum/accounts"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
)

// Loads Ethereum client provided an endpoint URL.
func LoadEthClient(ethEndpoint string) *ethClient.Client {
	eth, err := ethClient.Dial(ethEndpoint)
	if err != nil {
		panic(err)
	}
	LogMessage("Endpoint Loaded: "+ethEndpoint,LOG_INFO)
	return eth
}

//
func LoadEthClientMulti(ethEndpoint string) []*ethClient.Client {
	ethAttestAddrs := strings.Split(ethEndpoint,",")
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
func RpcCall(rpc *rpc.Client,To string,Data string) {
	type request struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}

	var result string

	req := request{To,Data}
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
		fmt.Println("tx:",tx)
	}

	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := eth.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance:", balance) // 25893180161173005034
}

// Generates and returns keystore from private key.
func InitKeystore(privateKey *ecdsa.PrivateKey) accounts.Account {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	input := Scan("Enter passphrase for new keystore:")
	account,err := ks.ImportECDSA(privateKey,input)
	if(err != nil) { panic(err) }
	fmt.Println("New keystore created for address",account.Address)
	fmt.Println("URL:",account.URL)
	return account
}

// Generates and returns new ECDSA keypair.
func GenerateKeypair() (string, string) {
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil { log.Fatal(err) }
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))[2:]

	account := InitKeystore(privateKey)
	_ = account
	
	// Generate public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok { log.Fatal("error casting public key to ECDSA") }
	publicKeyHex := hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:]

	fmt.Println("Wallet Created:",crypto.PubkeyToAddress(*publicKeyECDSA))
	
	return privateKeyHex, publicKeyHex
}

// Returns Keccak hash of string as bytes.
func KeccakHash(stateRootStr string) []byte {
	return crypto.Keccak256([]byte(stateRootStr))
}

// Returns Keccak hash of string as hex string with '0x' prefix.
func KeccakHashString(stateRootStr string) string {
	return hexutil.Encode(KeccakHash(stateRootStr))
}

// Simple pop function for discarding offline/malfunctioning endpoints
func ethClientsShift(ethClients []*ethClient.Client,recycle bool) (*ethClient.Client, []*ethClient.Client) {
	eth := ethClients[0]
	ethClients = ethClients[1:]
	if(recycle) {
		ethClients = append(ethClients,eth)
	}
	return eth, ethClients
}

func GetNonce(client *ethClient.Client, fromAddress common.Address) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("Nonce:",nonce)
	return nonce;
}

type LagrangeNodeCredentials struct {
	privateKey string
	publicKey interface{} // crypto.PublicKey
	address common.Address
	
	privateKeyECDSA *ecdsa.PrivateKey
	publicKeyECDSA *ecdsa.PublicKey
}

func GetCredentials() *LagrangeNodeCredentials {
	privateKeyString := GetPrivateKey()
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}
//	fmt.Println("Private key loaded.");

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
//	fmt.Println("Public key loaded.")

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	fmt.Println("Address isolated.")
	
	res := LagrangeNodeCredentials {
		privateKey: privateKeyString,
		publicKey: publicKey,
		address: fromAddress,
		privateKeyECDSA: privateKey,
		publicKeyECDSA: publicKeyECDSA }
	
	return &res
}

// Requests and returns network gas price.
func GetGasPrice(client *ethClient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

func GetAuth(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	return auth
}

// Test function - mines a new block each second.
func MineTest(rpc *rpc.Client) {
	for {
		MineBlocks(rpc,1)
		time.Sleep(1 * time.Second)
	}
}

// Test function - mines n blocks on Hardhat node
func MineBlocks(rpc *rpc.Client, num int) {
	var hex hexutil.Bytes
	for i := 0; i < num; i++ {
		rpc.Call(&hex,"evm_mine")
	}
}
