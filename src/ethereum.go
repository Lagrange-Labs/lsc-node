package main

import "fmt"
import "time"

import "math/big"

import common "github.com/ethereum/go-ethereum/common"
import ethClient "github.com/ethereum/go-ethereum/ethclient"
import rpc "github.com/ethereum/go-ethereum/rpc"
import log "log"

import context "context"

//import json "encoding/json"

import host "github.com/libp2p/go-libp2p-core/host"
import pubsub "github.com/libp2p/go-libp2p-pubsub"
import "strconv"

import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/accounts/abi/bind"

import	"crypto/ecdsa"
import	"github.com/ethereum/go-ethereum/common/hexutil"

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

func listenForBlocks(eth *ethClient.Client,node host.Host, topic *pubsub.Topic, ps *pubsub.PubSub, nick string, subscription *pubsub.Subscription) {
	for {
		block, err := eth.BlockByNumber(context.Background(),nil)
		if(err != nil) { panic(err) }
		txns := block.Transactions()
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
	fmt.Println("Private key loaded.");

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fmt.Println("Public key loaded.")

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Address isolated.")
	
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

func ctrIntTest(rpc *rpc.Client, client *ethClient.Client) {
	// Connect to Staking Contract
	instance := getStakingContract(client)

	// Retrieve private key, public key, address
	credentials := getCredentials()
	privateKey := credentials.privateKeyECDSA
	fromAddress := credentials.address

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

	// Add Stake
	StakeAdd(instance,auth)

	// Hardhat - Mine Blocks
	mineBlocks(rpc)

	// Update Nonce and Val
	nonce = getNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	
	// Begin Stake Removal
	StakeRemoveBegin(instance, auth)

	// Hardhat - Mine More Blocks
	mineBlocks(rpc)
	
	// Update Nonce
	nonce = getNonce(client,fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	
	// Finalize Stake Removal
	StakeRemoveFinish(instance,auth)
}
