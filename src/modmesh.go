package main

import "fmt"
import "flag"
import "time"
import "strconv"
import json "encoding/json"
//import bufio "bufio"

import _ "reflect"

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
//import pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"

import "github.com/libp2p/go-libp2p/p2p/discovery/mdns"

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
	_ = ps
	_ = topic
	_ = subscription

//	node.SetStreamHandler(protocol.ID(room), handleMessaging)

	handleMessaging()

	message := "test"
	go writeMessages(node,topic,nick,message)
	go readMessages(node,topic,subscription)

        // SIGINT | SIGTERM Signal Handling - End
        termHandler(node)
}

func handleMessaging() {
}

/*
func handleMessaging(stream network.Stream) {
	fmt.Println("(New Stream)")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	_ = rw
	
//	go readMessages(rw)
//	go writeMessages(rw)

	// 'stream' will stay open until you close it (or the other side closes it).
}
*/

// Converted to/from JSON and sent in the body of pubsub messages.
type GossipMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

const bufferSize = 4096

type MsgParams struct {
	ps *pubsub.PubSub
	topic *pubsub.Topic
	subscription *pubsub.Subscription
	node host.Host
	nick string
	message string
}

func readMessages(node host.Host, topic *pubsub.Topic, subscription *pubsub.Subscription) {
//rw *bufio.ReadWriter
	messages :=  make(chan *GossipMessage, bufferSize)
	
	for {
		msg, err := subscription.Next(context.Background())
		if err != nil {
			close(messages)
			return
		}
		// only forward messages delivered by others
		if msg.ReceivedFrom == node.ID() {
			continue
		}
		fmt.Println(msg.Data)
		cm := new(GossipMessage)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		// send valid messages onto the Messages channel
		messages <- cm
	}
}

func writeMessages(node host.Host, topic *pubsub.Topic, nick string, message string) error {
//rw *bufio.ReadWriter
	m := GossipMessage{
		Message:    message,
		SenderID:   node.ID().Pretty(),
		SenderNick: nick,
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return topic.Publish(context.Background(), msgBytes)
}

/* pubsub example helpers */

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "modmesh"

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(h host.Host) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{h: h})
	return s.Start()
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}
