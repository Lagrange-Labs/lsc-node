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

import peerstore "github.com/libp2p/go-libp2p-core/peer"
import multiaddr "github.com/multiformats/go-multiaddr"

import pubsub "github.com/libp2p/go-libp2p-pubsub"
import context "context"


func createListener(portPtr int) host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/"+strconv.Itoa(portPtr)),
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
	fmt.Println(peerInfo)
	_ = err
	return peerInfo
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

func connectRemote(node host.Host, peerAddr string) {
	if (len(peerAddr) > 0) {
		addr, err := multiaddr.NewMultiaddr(peerAddr)
		if err != nil {
			panic(err)
		}

		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			panic(err)
		}
		
		if err := node.Connect(context.Background(), *peer); err != nil {
			panic(err)
		}

		ch := ping.Ping(context.Background(), node, peer.ID)
		for i := 0; i < 5; i++ {
			res := <-ch
			fmt.Println("Got ping response.", "Latency:", res.RTT)
		}
	}
}

func main() {
	// Parse Port
	portPtr := flag.Int("port",8081,"Server listening port")
	// Parse Nickname
	nickPtr := flag.String("nick","","Nickname - CLI flag, blank by default, consider addresses or protocol TLDs later.")
	// Parse Room
	roomPtr := flag.String("room","rinkeby","Room / Network")
	// Parse Remote Peer
	peerAddrPtr := flag.String("peerAddr","","Remote Peer Address")
	// Parse Remote Peer Port
	peerPortPtr := flag.String("peerPort","8081","Remote Peer Port")

	flag.Parse()

	port := *portPtr
	nick := *nickPtr
	room := *roomPtr
	peerAddr := *peerAddrPtr
	peerPort := *peerPortPtr
	
	_ = peerPort

	fmt.Println("Port:",port)

	rpcCall("0xcc13fc627effd6e35d2d2706ea3c4d7396c610ea","0x8da5cb5b")
	
//	os.Exit(0)
	
	// Create listener
	node := createListener(port)

	if(len(nick) == 0) {
		nick = fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(node.ID()))
	}
	fmt.Println("Nickname:",nick)

	// Get P2P Address Info
	localInfo := getAddrInfo(node);
	_ = localInfo

	// Ping test - please determine an approach to finding peers, rather than self-pinging	
	ch := ping.Ping(context.Background(), node, localInfo.ID)
	for i := 0; i < 5; i++ {
		res := <-ch
		fmt.Println("Got ping response.", "Latency:", res.RTT)
	}
	
	// Connect to Remote Peer
	connectRemote(node,peerAddr)
	
	ps, topic, subscription := getGossipSub(node,room)

	go handleMessaging(node,topic,ps,nick,subscription)

        // SIGINT | SIGTERM Signal Handling - End
        termHandler(node)
}
