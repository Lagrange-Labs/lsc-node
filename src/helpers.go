package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	
	"math/rand"

	context "context"

	peer "github.com/libp2p/go-libp2p-core/peer"
	host "github.com/libp2p/go-libp2p-core/host"

	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)
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
	fmt.Println("peer.AddrInfo:",pi);
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

const (
        InfoColor    = "\033[1;34m%s\033[0m"
        NoticeColor  = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)


func scan(prompt string) string {
	if(prompt != "") {
		fmt.Println(prompt)
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func genSalt32() string {
	res := ""
	
	rand.Seed(time.Now().UnixNano())

	// String
	charset := "0123456789abcdef"
	
	for i := 0; i < 32; i++ {
		// Getting random character
		c := charset[rand.Intn(len(charset))]
		res += string(c)
	}
	
	return res
}

func signMessageWithPrivateKey(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	messageHash := accounts.TextHash([]byte(message))

	signature, err := crypto.Sign(messageHash, privateKey)
	if err != nil {
		return "", err
	}

	signature[crypto.RecoveryIDOffset] += 27

	return hexutil.Encode(signature), nil
}
