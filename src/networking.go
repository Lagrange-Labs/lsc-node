package main

import (
	"fmt"
	"strconv"

	"github.com/libp2p/go-libp2p"

	host "github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	multiaddr "github.com/multiformats/go-multiaddr"

	context "context"
)

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

func getAddrInfo(node host.Host) peerstore.AddrInfo {
	// print the node's PeerInfo in multiaddr format
	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
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
	fmt.Println("{address: '0x...'}")
	
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

		// assuming that your tracer runs in x.x.x.x and has a peer ID of QmTracer
		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			panic(err)
		}

		if err := node.Connect(context.Background(), *peer); err != nil {
			panic(err)
		} else {
			fmt.Println("Connected:", *peer)
		}
	}
}

func handleNewPeer() {
	
}
