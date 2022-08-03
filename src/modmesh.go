package main

import "fmt"
import "flag"
import "github.com/libp2p/go-libp2p"
import "strconv"

import "context"
//import "github.com/ethereum/go-ethereum/accounts/abi/bind"
import "github.com/ethereum/go-ethereum/common"
//import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/ethclient"

import "log"

func main() {

	// Parse Port
	portPtr := flag.Int("port",8080,"Server listening port")
	flag.Parse()
	fmt.Println("Port:",*portPtr)

	// Create ethclient instance pointing to local Hardhat node
	eth, err := ethclient.Dial("http://0.0.0.0:8545")
	if err != nil {
		panic(err)
	}
	  _ = eth

	// Test data requests from Hardhat node
	ctx := context.Background()
	tx, pending, _ := eth.TransactionByHash(ctx, common.HexToHash("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	if !pending {
		fmt.Println("tx:",tx)
	}

	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := eth.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance:", balance) // 25893180161173005034

	// Loopback Interface - Create listener
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/"+strconv.Itoa(*portPtr)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen addresses:", node.Addrs())

	// Shutdown
	if err := node.Close(); err != nil {
		panic(err)
	}
}
