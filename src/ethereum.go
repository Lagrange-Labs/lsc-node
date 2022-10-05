package main

import (
	"fmt"
	"time"
	"os"

	"math/big"
	
	common "github.com/ethereum/go-ethereum/common"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	rpc "github.com/ethereum/go-ethereum/rpc"
	log "log"

	context "context"

	json "encoding/json"

	host "github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
//	sha3 "github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	accounts "github.com/ethereum/go-ethereum/accounts"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	STAKE_STATUS_CLOSED = 0
	STAKE_STATUS_PENDING = 1
	STAKE_STATUS_OPEN = 2
)

func LoadEthClient(ethEndpoint string) *ethClient.Client {
	eth, err := ethClient.Dial(ethEndpoint)
	if err != nil {
		panic(err)
	}
	fmt.Println("Endpoint:",ethEndpoint)
	return eth
}

func LoadRpcClient(ethEndpoint string) *rpc.Client {
	rpc, err := rpc.DialHTTP(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer rpc.Close()
	return rpc
}

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

func InitKeystore(privateKey *ecdsa.PrivateKey) accounts.Account {
	// Generate keystore with private key
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	input := Scan("Enter passphrase for new keystore:")
	account,err := ks.ImportECDSA(privateKey,input)
	if(err != nil) { panic(err) }
	fmt.Println("New keystore created for address",account.Address)
	fmt.Println("URL:",account.URL)
	return account
}

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

func HandleStakingEvent(vLog *NodestakingStakedNode) {
	node := vLog.Node
	chainId := vLog.ChainId
	amount := vLog.Amount
	claimTime := vLog.ClaimTime

	fmt.Println(node,chainId,amount,claimTime)
}

func ListenForStaking(ethWS *ethClient.Client) {
	sc := GetStakingContract(ethWS)
	_ = sc
	logs := make(chan *NodestakingStakedNode)

	block := uint64(100)
	
	sub,err := sc.WatchStakedNode(&bind.WatchOpts{&block,context.Background()},logs)
	if err != nil { panic(err) }
	
	for {
		fmt.Println("Listening to contract", NODE_STAKING_ADDRESS)
		select {
			case err := <-sub.Err():
				log.Fatal(err)
				continue
			case vLog := <-logs:
				HandleStakingEvent(vLog)
		}
	}
}

func GetSeparator() string { return "::" }

func GenerateStateRootString(eth *ethClient.Client, block *types.Block) string {
	//5. ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the State root, Timestamp, Block Number and Sharded EdDSA Signature Tuple
	stateRootSeparator := GetSeparator()
	
	blockRoot := block.Root().String()
	blockTime := strconv.FormatUint(block.Time(),10)
	blockNumber := block.Number().String()
	chain,err := eth.ChainID(context.Background())
	if err != nil { panic(err) }
	chainID := chain.String()		
	
	stateRootStr := blockRoot + stateRootSeparator + blockTime + stateRootSeparator + blockNumber + stateRootSeparator + chainID
	fmt.Println("State Root String:",stateRootStr)

	return stateRootStr
}

func KeccakHash(stateRootStr string) []byte {
	return crypto.Keccak256([]byte(stateRootStr))
}

func KeccakHashString(stateRootStr string) string {
	return hexutil.Encode(KeccakHash(stateRootStr))
}

type StateRootMessage struct {
	StateRoot string
	Timestamp string
	BlockNumber string
	ShardedEdDSASignatureTuple string
	ECDSASignatureTuple string
	EthereumPublicKey string
}

/*
For gossiping of state roots:

1. State Root
2. Timestamp
3. Block Number
4. Sharded EdDSA Signature Tuple (TBD exact parameters)
5. ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the State root, Timestamp, Block Number and Sharded EdDSA Signature Tuple
6. Ethereum Public Key
*/
func ListenForBlocks(eth *ethClient.Client, node host.Host, topic *pubsub.Topic, ps *pubsub.PubSub, nick string, subscription *pubsub.Subscription) {
	stateRootSeparator := GetSeparator()

	for {
		block, err := eth.BlockByNumber(context.Background(),nil)
		if(err != nil) { panic(err) }
		
		// concatenate relevant fields
		stateRootStr := GenerateStateRootString(eth, block)
		
		//ShardedEdDSASignatureTuple - TBD
		shardedSignatureTuple := ""
		
		stateRootStrWithShardedSignatureTuple := stateRootStr + stateRootSeparator + shardedSignatureTuple
		
		// generate hash from concatenated fields
		stateHash := KeccakHash(stateRootStrWithShardedSignatureTuple)
		
		// sign resultant hash
		creds := GetCredentials()
		privateKey := creds.privateKeyECDSA
		signature, err := crypto.Sign([]byte(stateHash), privateKey) // we should probably be signing the raw signature, not the hex (TODO)
		if err != nil { panic(err) }
		ecdsaSignatureHex := hexutil.Encode(signature)

		//timestamp
		timestamp := time.Now().UTC().Unix()
		
		//public key
		publicKeyECDSA := creds.publicKeyECDSA
		
		stateRootMessage := StateRootMessage {
			StateRoot: stateRootStr,
			Timestamp: strconv.FormatInt(timestamp,10),
			BlockNumber: block.Number().String(),
			ShardedEdDSASignatureTuple: shardedSignatureTuple,
			ECDSASignatureTuple: ecdsaSignatureHex,
			EthereumPublicKey: hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA)) }
		
		json,err := json.Marshal(stateRootMessage)
		if err != nil { panic(err) }
		bytes := []byte(json)
		msg := string(bytes)
		
		WriteMessages(node,topic,creds.address.Hex(),msg)
		
		time.Sleep(1 * time.Second)
	}
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

// Returns a NodeStaking struct instance.
func GetStakingContract(client *ethClient.Client) *Nodestaking {
	address := common.HexToAddress(NODE_STAKING_ADDRESS)
	instance, err := NewNodestaking(address,client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded contract address",address)
	return instance
}

// Requests and returns network gas price.
func GetGasPrice(client *ethClient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

// Initializes staking transaction with NodeStaking smart contract.
func StakeBegin(instance *Nodestaking) *big.Int {
	stake, err := instance.StakeAmount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stake amount:",stake)
	return stake
}

// Adds stake to NodeStaking smart contract.
// Staking transaction must first be initialized with StakeBegin().
func StakeAdd(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.AddStake(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

// Begins removal of stake from NodeStaking smart contract.
func StakeRemoveBegin(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.StartStakeRemoval(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

// Finalizes removal of stake from NodeStaking smart contract.
// Staking removal transaction must first be initialzied with StakeRemoveBegin().
func StakeRemoveFinish(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.FinishStakeRemoval(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

func GetAuth(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	return auth
}

func MineTest(rpc *rpc.Client) {
	for {
		MineBlocks(rpc,1)
		time.Sleep(1 * time.Second)
	}
}
func MineBlocks(rpc *rpc.Client, num int) {
	var hex hexutil.Bytes
	for i := 0; i < num; i++ {
		rpc.Call(&hex,"evm_mine")
	}
}

func VerifyStake(client *ethClient.Client, instance *Nodestaking, addr common.Address) bool {
	activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},addr,big.NewInt(4))
	if(err != nil) { panic(err) }
	return activeStakes == STAKE_STATUS_OPEN
}

func ActiveStakesTest(rpc *rpc.Client, client *ethClient.Client) {
	instance := GetStakingContract(client)
	_ = instance
//	activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},common.Address(""),big.NewInt(4))
//	if(err != nil) { panic(err) }
//	fmt.Println(activestakes)
}

func CtrIntTest(rpc *rpc.Client, client *ethClient.Client) {
	// Connect to Staking Contract
	instance := GetStakingContract(client)

	// Retrieve private key, public key, address
	credentials := GetCredentials()
	privateKey := credentials.privateKeyECDSA
	fromAddress := credentials.address

	// Verify Stake
	isStaked := VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Request nonce for transaction	
	nonce := GetNonce(client,fromAddress)

	// Request gas price
	gasPrice := GetGasPrice(client)

	// Begin Staking Transaction
	stake := StakeBegin(instance)

	auth := GetAuth(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = stake
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Add Stake
	StakeAdd(instance,auth)
	
	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Hardhat - Mine Blocks
	MineBlocks(rpc,5)

	// Update Nonce and Val
	nonce = GetNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	
	// Begin Stake Removal
	StakeRemoveBegin(instance, auth)

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Hardhat - Mine More Blocks
	MineBlocks(rpc,5)
	
	// Update Nonce
	nonce = GetNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	
	// Finalize Stake Removal
	StakeRemoveFinish(instance,auth)

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)
	
	fmt.Println("*DONE*")
	os.Exit(0)
}
