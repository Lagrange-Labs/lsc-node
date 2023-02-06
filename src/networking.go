package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/libp2p/go-libp2p"

	host "github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	multiaddr "github.com/multiformats/go-multiaddr"

	context "context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	json "encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
)

func CreateListener(portPtr int) host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/"+strconv.Itoa(portPtr)),
	)
	if err != nil {
		panic(err)
	}
	LogMessage("Listen addresses:"+ fmt.Sprintf("%v",node.Addrs()),LOG_INFO)
	return node
}

func GetAddrInfo(node host.Host) peerstore.AddrInfo {
	// print the node's PeerInfo in multiaddr format
	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	LogMessage("libp2p node address: " + fmt.Sprintf("%v",addrs[0]),LOG_INFO)
	LogMessage("Peer Info: "+fmt.Sprintf("%v",peerInfo),LOG_INFO)
	_ = err
	return peerInfo
}

type JoinMessage struct {
	PublicKey string
	GenericMessage string
	Timestamp string
	Salt string
	ECDSASignatureTuple string
}

func pingPeer(node host.Host) {
	// Get P2P Address Info
	localInfo := GetAddrInfo(node);
	_ = localInfo
	// Ping test
	ch := ping.Ping(context.Background(), node, localInfo.ID)
	for i := 0; i < 3; i++ {
		res := <-ch
		LogMessage("Got ping response. Latency: "+fmt.Sprintf("%v",res.RTT), LOG_DEBUG)
	}
}

/*
Ethereum Public Key
Generic Message (i.e. “Hello Network”)
Timestamp
Salt
ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the generic message + timestamp + salt
*/

func GenerateVerificationTuple() (string, string, string, string) {
	timestampStr := strconv.FormatInt(time.Now().UTC().Unix(),10)
	genericMessage := "It's always morning in web3."
	saltStr := GenSalt32()
	return GenerateVerificationTupleFromJoinMessage(genericMessage, timestampStr, saltStr), timestampStr, genericMessage, saltStr
}

func GenerateVerificationTupleFromJoinMessage(genericMessage string, timestampStr string, salt string) string {
	separator := GetSeparator()
	return genericMessage + separator + timestampStr + separator + salt
}

func SendVerificationMessage(node host.Host,topic *pubsub.Topic) {

	creds := GetCredentials()
	
	// ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the generic message + timestamp + salt
	
	tuple, timestampStr, genericMessage, saltStr := GenerateVerificationTuple()

	tupleHash := KeccakHash(tuple)

	signatureTuple,err := crypto.Sign(tupleHash, creds.privateKeyECDSA)
	if err != nil { panic(err) }
	
	signatureHex := hexutil.Encode(signatureTuple)
	
	joinMessage := JoinMessage {
		PublicKey: hexutil.Encode(crypto.FromECDSAPub(creds.publicKeyECDSA)),
		GenericMessage: genericMessage,
		Timestamp: timestampStr,
		Salt: saltStr,
		ECDSASignatureTuple: signatureHex }

	json,err := json.Marshal(joinMessage)
	if err != nil { panic(err) }
	bytes := []byte(json)
	msg := string(bytes)
	
	WriteMessages(node,topic,creds.address.Hex(),msg,"JoinMessage")
}

func GetGossipSub(node host.Host, roomName string) (*pubsub.PubSub,*pubsub.Topic,*pubsub.Subscription) {
	ps, err := pubsub.NewGossipSub(context.Background(), node)
	if err != nil {
		panic(err)
	}
	
	// mDNS discovery
	if err := SetupDiscovery(node); err != nil {
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
	
	LogMessage("Room joined and subscribed: "+TopicName(roomName),LOG_INFO)
		
	return ps,topic,sub
}

func TopicName(networkName string) string {
	return "modmesh-" + networkName
}

func ConnectRemote(node host.Host, peerAddr string) {
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
			LogMessage("Connected: "+fmt.Sprintf("%v",*peer),LOG_INFO)
		}
	}
}

func HandleNewPeer() {
	
}
