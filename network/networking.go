package network

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Lagrange-Labs/Lagrange-Node/utils"
	"github.com/libp2p/go-libp2p"

	host "github.com/libp2p/go-libp2p/core/host"
	peerstore "github.com/libp2p/go-libp2p/core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	multiaddr "github.com/multiformats/go-multiaddr"

	context "context"
)

func CreateListener(portPtr string) host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/" + portPtr),
	)
	if err != nil {
		panic(err)
	}
	utils.LogMessage("Listen addresses:"+fmt.Sprintf("%v", node.Addrs()), utils.LOG_INFO)
	return node
}

func GetAddrInfo(node host.Host) peerstore.AddrInfo {
	// print the node's PeerInfo in multiaddr format
	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	utils.LogMessage("libp2p node address: "+fmt.Sprintf("%v", addrs[0]), utils.LOG_INFO)
	utils.LogMessage("Peer Info: "+fmt.Sprintf("%v", peerInfo), utils.LOG_INFO)
	_ = err
	return peerInfo
}

type JoinMessage struct {
	GenericMessage      string
	Timestamp           string
	Salt                string
	ECDSASignatureTuple string
}

/*
Ethereum Public Key
Generic Message (i.e. “Hello Network”)
Timestamp
Salt
ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the generic message + timestamp + salt
*/

func GenerateVerificationTuple() (string, string, string, string) {
	timestampStr := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	genericMessage := "It's always morning in web3."
	saltStr := utils.GenSalt32()
	return GenerateVerificationTupleFromJoinMessage(genericMessage, timestampStr, saltStr), timestampStr, genericMessage, saltStr
}

func GenerateVerificationTupleFromJoinMessage(genericMessage string, timestampStr string, salt string) string {
	separator := utils.GetSeparator()
	return genericMessage + separator + timestampStr + separator + salt
}

func GetGossipSub(node host.Host, roomName string) (*pubsub.PubSub, *pubsub.Topic, *pubsub.Subscription) {
	ps, err := pubsub.NewGossipSub(context.Background(), node)
	if err != nil {
		panic(err)
	}

	// mDNS discovery
	if err := utils.SetupDiscovery(node); err != nil {
		panic(err)
	}

	// Join
	topic, err := ps.Join(TopicName(roomName))
	if err != nil {
		panic(err)
	}
	_ = topic

	// Subscribe
	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}

	utils.LogMessage("Room joined and subscribed: "+TopicName(roomName), utils.LOG_INFO)

	return ps, topic, sub
}

func TopicName(networkName string) string {
	return "modmesh-" + networkName
}

func ConnectRemote(node host.Host, peerAddr string) {
	if len(peerAddr) > 0 {
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
			utils.LogMessage("Connected: "+fmt.Sprintf("%v", *peer), utils.LOG_INFO)
		}
	}
}

func HandleNewPeer() {

}
