package main

import "fmt"
import "flag"
import "strconv"

import "os"
import "os/signal"
import "syscall"

import "github.com/libp2p/go-libp2p"
import peer "github.com/libp2p/go-libp2p-core/peer"
import host "github.com/libp2p/go-libp2p-core/host"
//import protocol "github.com/libp2p/go-libp2p-core/protocol"
//import network "github.com/libp2p/go-libp2p-core/network"
import ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"


import context "context"
//import "github.com/ethereum/go-ethereum/accounts/abi/bind"
import common "github.com/ethereum/go-ethereum/common"
//import "github.com/ethereum/go-ethereum/crypto"
import ethClient "github.com/ethereum/go-ethereum/ethclient"

import log "log"

import pubsub "github.com/libp2p/go-libp2p-pubsub"

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

func getGossipSub(node host.Host, roomName string) (*pubsub.PubSub,*pubsub.Topic,*pubsub.Subscription) {
	ps, err := pubsub.NewGossipSub(context.Background(), node)
	if err != nil {
		panic(err)
	}
	
	// mDNS discovery
	if err := setupDiscovery(node); err != nil {
		panic(err)
	}
	
	// Join
	topic, err := ps.Join(topicName(roomName))
	if err != nil {
		panic(err)
	}
	_ = topic
	
	// Subscribe
	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Room joined and subscribed:",topicName(roomName))
	
	return ps,topic,sub
}

func topicName(networkName string) string {
	return "modmesh-" + networkName
}

func main() {
	// Parse Port
	portPtr := flag.Int("port",8081,"Server listening port")
	// Parse Nickname
	nickPtr := flag.String("nick","","Nickname - CLI flag, blank by default, consider addresses or protocol TLDs later.")
	// Parse Room
	roomPtr := flag.String("room","rinkeby","Room / Network")
	flag.Parse()

	port := *portPtr
	nick := *nickPtr
	room := *roomPtr

	fmt.Println("Port:",port)

	// Create ethclient instance pointing to local Hardhat node
	eth := loadEthClient()
	_ = eth
	
	// Test data requests from Hardhat node
//	ethTest(eth)
	
	// Create listener
	node := createListener(port)

	if(len(nick) == 0) {
		nick = fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(node.ID()))
	}
	fmt.Println("Nickname:",nick)

	// Get P2P Address Info
	peerInfo := getAddrInfo(node);
	_ = peerInfo

	// Ping test - please determine an approach to finding peers, rather than self-pinging	
	ch := ping.Ping(context.Background(), node, peerInfo.ID)
	for i := 0; i < 5; i++ {
		res := <-ch
		fmt.Println("Got ping response.", "Latency:", res.RTT)
	}
	
	ps, topic, subscription := getGossipSub(node,room)

	go handleMessaging(node,topic,ps,nick,subscription)

        // SIGINT | SIGTERM Signal Handling - End
        termHandler(node)
}
