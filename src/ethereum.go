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

	//json "encoding/json"

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

func loadEthClient(ethEndpoint string) *ethClient.Client {
	eth, err := ethClient.Dial(ethEndpoint)
	if err != nil {
		panic(err)
	}
	fmt.Println("Endpoint:",ethEndpoint)
	return eth
}

func loadRpcClient(ethEndpoint string) *rpc.Client {
	rpc, err := rpc.DialHTTP(ethEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer rpc.Close()
	return rpc
}

func rpcCall(rpc *rpc.Client,To string,Data string) {
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

func ethTest(eth *ethClient.Client) {
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

func initKeystore(privateKey *ecdsa.PrivateKey) accounts.Account {
	// Generate keystore with private key
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	input := scan("Enter passphrase for new keystore:")
	account,err := ks.ImportECDSA(privateKey,input)
	if(err != nil) { panic(err) }
	fmt.Println("New keystore created for address",account.Address)
	fmt.Println("URL:",account.URL)
	return account
}

func generateKeypair() (string, string) {
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil { log.Fatal(err) }
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))[2:]

	account := initKeystore(privateKey)
	_ = account
	
	// Generate public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok { log.Fatal("error casting public key to ECDSA") }
	publicKeyHex := hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:]

	fmt.Println("Wallet Created:",crypto.PubkeyToAddress(*publicKeyECDSA))
	
	return privateKeyHex, publicKeyHex
}

func handleStakingEvent(vLog *NodestakingStakedNode) {
	node := vLog.Node
	chainId := vLog.ChainId
	amount := vLog.Amount
	claimTime := vLog.ClaimTime

	fmt.Println(node,chainId,amount,claimTime)
}

func listenForStaking(ethWS *ethClient.Client) {
	sc := getStakingContract(ethWS)
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
				handleStakingEvent(vLog)
		}
	}
}

func generateStateRootString(eth *ethClient.Client, block *types.Block) string {
	stateRootSeparator := "::"
	
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

func keccakHash(stateRootStr string) []byte {
	return crypto.Keccak256([]byte(stateRootStr))
}

func keccakHashString(stateRootStr string) string {
	return hexutil.Encode(keccakHash(stateRootStr))
}

type StateRootMessage struct {
	StateRoot string
	Timestamp string
	BlockNumber string
	ShardedEdDSASignatureTuple string
	ECDSASignatureTuple string
	EthereumPublicKey string
}

func listenForBlocks(eth *ethClient.Client,node host.Host, topic *pubsub.Topic, ps *pubsub.PubSub, nick string, subscription *pubsub.Subscription) {
	for {
		block, err := eth.BlockByNumber(context.Background(),nil)
		if(err != nil) { panic(err) }
		txns := block.Transactions()
		
		// concatenate relevant fields
		stateRootStr := generateStateRootString(eth, block)
		
		// generate hash from concatenated fields
		stateHash := keccakHash(stateRootStr)
		fmt.Println("keccak256 hex:",keccakHashString(stateRootStr))
		
		// sign resultant hash
		creds := getCredentials()
		privateKey := creds.privateKeyECDSA
		signature, err := crypto.Sign([]byte(stateHash), privateKey) // we should probably be signing the raw signature, not the hex (TODO)
		if err != nil { panic(err) }
		signatureHex := hexutil.Encode(signature)
		fmt.Println("signature hex:",signatureHex)
		
/*
For gossiping of state roots:

1. State Root
2. Timestamp
3. Block Number
4. Sharded EdDSA Signature Tuple (TBD exact parameters)
5. ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the State root, Timestamp, Block Number and Sharded EdDSA Signature Tuple
6. Ethereum Public Key
*/
		
		msg := "{'block.Number':"+block.Number().String()+",'block.Hash':"+block.Hash().String()+",'txnCount':"+strconv.Itoa(len(txns))+"}"
		writeMessages(node,topic,nick,msg)
		
		time.Sleep(1 * time.Second)
	}
}

func getNonce(client *ethClient.Client, fromAddress common.Address) uint64 {
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

func getCredentials() *LagrangeNodeCredentials {
	privateKeyString := getPrivateKey()
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

func getStakingContract(client *ethClient.Client) *Nodestaking {
	address := common.HexToAddress(NODE_STAKING_ADDRESS)
	instance, err := NewNodestaking(address,client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded contract address",address)
	return instance
}

func getGasPrice(client *ethClient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

func StakeBegin(instance *Nodestaking) *big.Int {
	stake, err := instance.StakeAmount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stake amount:",stake)
	return stake
}

func StakeAdd(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.AddStake(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

func StakeRemoveBegin(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.StartStakeRemoval(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

func StakeRemoveFinish(instance *Nodestaking, auth *bind.TransactOpts) {
	res,err := instance.FinishStakeRemoval(auth,big.NewInt(4))
	if(err != nil) { panic(err) }
	_ = res
}

func getAuth(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	return auth
}

func mineBlocks(rpc *rpc.Client) {
	var hex hexutil.Bytes
	for i := 0; i < 5; i++ {
		rpc.Call(&hex,"evm_mine")
	}
}

func VerifyStake(client *ethClient.Client, instance *Nodestaking, addr common.Address) bool {
	activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},addr,big.NewInt(4))
	if(err != nil) { panic(err) }
	return activeStakes == STAKE_STATUS_OPEN
}

func activeStakesTest(rpc *rpc.Client, client *ethClient.Client) {
	instance := getStakingContract(client)
	_ = instance
//	activeStakes,err := instance.ActiveStakes(&bind.CallOpts{},common.Address(""),big.NewInt(4))
//	if(err != nil) { panic(err) }
//	fmt.Println(activestakes)
}

func ctrIntTest(rpc *rpc.Client, client *ethClient.Client) {
	// Connect to Staking Contract
	instance := getStakingContract(client)

	// Retrieve private key, public key, address
	credentials := getCredentials()
	privateKey := credentials.privateKeyECDSA
	fromAddress := credentials.address

	// Verify Stake
	isStaked := VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Request nonce for transaction	
	nonce := getNonce(client,fromAddress)

	// Request gas price
	gasPrice := getGasPrice(client)

	// Begin Staking Transaction
	stake := StakeBegin(instance)

	auth := getAuth(privateKey)
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
	mineBlocks(rpc)

	// Update Nonce and Val
	nonce = getNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	
	// Begin Stake Removal
	StakeRemoveBegin(instance, auth)

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)

	// Hardhat - Mine More Blocks
	mineBlocks(rpc)
	
	// Update Nonce
	nonce = getNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	
	// Finalize Stake Removal
	StakeRemoveFinish(instance,auth)

	// Verify Stake
	isStaked = VerifyStake(client,instance,fromAddress)
	fmt.Println("Stake Verification:",isStaked)
	
	fmt.Println("*DONE*")
	os.Exit(0)
}
