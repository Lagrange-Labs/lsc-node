package main

import "fmt"
import "flag"

import _ "reflect"

import "os"
import "os/signal"
import "syscall"

import "github.com/libp2p/go-libp2p"
import peer "github.com/libp2p/go-libp2p-core/peer"
import host "github.com/libp2p/go-libp2p-core/host"
import ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
//import "github.com/libp2p/go-libp2p-peerstore/addr"
import "strconv"

import context "context"
//import "github.com/ethereum/go-ethereum/accounts/abi/bind"
import common "github.com/ethereum/go-ethereum/common"
//import "github.com/ethereum/go-ethereum/crypto"
import ethClient "github.com/ethereum/go-ethereum/ethclient"

import log "log"

//import pubsub "github.com/libp2p/go-libp2p-pubsub"
//import pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"

func getPort() int {
	portPtr := flag.Int("port",8080,"Server listening port")
	flag.Parse()
	fmt.Println("Port:",*portPtr)
	return *portPtr	
}

func loadEthClient() *ethClient.Client {
	eth, err := ethClient.Dial("http://0.0.0.0:8545")
	if err != nil {
		panic(err)
	}
	return eth
}

func createListener(portPtr int) host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/"+strconv.Itoa(portPtr)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen addresses:", node.Addrs())
	return node
}

func termHandler(node host.Host) {
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
        <-ch
        fmt.Println("Received signal, shutting down...")

        // shut the node down
        if err := node.Close(); err != nil {
                panic(err)
        }
}

func getAddrInfo(node host.Host) peer.AddrInfo {
	// print the node's PeerInfo in multiaddr format
	peerInfo := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	fmt.Println("libp2p node address:", addrs[0])
	_ = err
	return peerInfo
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

/*
func getGossipSub(node host.Host) {
	ps, err := pubsub.NewGossipSub(context.Background(), node)
	if err != nil {
		panic(err)
	}
	return ps
}
*/

func main() {
	// Parse Port
	portPtr := getPort()

	// Create ethclient instance pointing to local Hardhat node
	eth := loadEthClient()
	_ = eth
	
	// Test data requests from Hardhat node
//	ethTest(eth)
	
	// Create listener
	node := createListener(portPtr)

	// Get P2P Address Info
	peerInfo := getAddrInfo(node);
	_ = peerInfo

	// Ping test - please determine an approach to finding peers, rather than self-pinging	
	ch := ping.Ping(context.Background(), node, peerInfo.ID)
	for i := 0; i < 5; i++ {
		res := <-ch
		fmt.Println("Got ping response.", "Latency:", res.RTT)
	}
	
//	getGossipSub(node)

        // SIGINT | SIGTERM Signal Handling - End
        termHandler(node)
}

