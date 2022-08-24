package main

import "fmt"

//import "github.com/ethereum/go-ethereum/accounts/abi/bind"
import common "github.com/ethereum/go-ethereum/common"
//import "github.com/ethereum/go-ethereum/crypto"
import ethClient "github.com/ethereum/go-ethereum/ethclient"
import rpc "github.com/ethereum/go-ethereum/rpc"
import log "log"

import context "context"

func loadEthClient(ethEndpoint string) *ethClient.Client {
	eth, err := ethClient.Dial(ethEndpoint)
	if err != nil {
		panic(err)
	}
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
	fmt.Printf("%s\n", owner.Hex()) // 0x281017b4E914b79371d62518b17693B36c7a221e
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

