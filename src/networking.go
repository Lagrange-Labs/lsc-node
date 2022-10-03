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

type JoinMessage struct {
	PublicKey string
	GenericMessage string
	Timestamp string
	Salt string
	ECDSASignatureTuple string
}

/*
Ethereum Public Key
Generic Message (i.e. “Hello Network”)
Timestamp
Salt
ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the generic message + timestamp + salt
*/

func getGossipSub(node host.Host, roomName string) (*pubsub.PubSub,*pubsub.Topic,*pubsub.Subscription) {
	separator := getSeparator()
	
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
	
	creds := getCredentials()
	
	// ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the generic message + timestamp + salt
	
	genericMessage := "It's always morning in web3."
	timestampStr := strconv.FormatInt(time.Now().UTC().Unix(),10)
	saltStr := genSalt32()
	tuple := genericMessage + separator + timestampStr + separator + saltStr

	tupleHash := keccakHash(tuple)

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
	
	writeMessages(node,topic,creds.address.Hex(),msg)
	
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
