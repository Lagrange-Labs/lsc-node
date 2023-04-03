package utils

import (
	"bufio"
	"os"
	"time"

	"crypto/rand"

	context "context"

	"github.com/Lagrange-Labs/Lagrange-Node/logger"
	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"

	"github.com/ethereum/go-ethereum/crypto"

	"crypto/ecdsa"

	accounts "github.com/ethereum/go-ethereum/accounts"
)

/* pubsub example helpers */

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "modmesh"

// DiscoveryNotifee gets notified when we find a new peer via mDNS discovery
type DiscoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	logger.Log.Infof("discovered new peer %s", pi.ID.Pretty())
	logger.Log.Infof("peer.AddrInfo: %v", pi)
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		logger.Log.Errorf("error connecting to peer %s: %s", pi.ID.Pretty(), err)
	}
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func SetupDiscovery(h host.Host) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &DiscoveryNotifee{h: h})
	return s.Start()
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func ShortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}

func Scan(prompt string) string {
	if prompt != "" {
		logger.Log.Info(prompt)
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func GenSalt32() string {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	res := hexutil.Encode(salt)
	return res
}

func SignMessageWithPrivateKey(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	messageHash := accounts.TextHash([]byte(message))

	signature, err := crypto.Sign(messageHash, privateKey)
	if err != nil {
		return "", err
	}

	signature[crypto.RecoveryIDOffset] += 27

	return hexutil.Encode(signature), nil
}

func GetUnixTimestamp() int64 {
	return time.Now().Unix()
}

// Returns standard delimiter for strings that are hashed and signed.
func GetSeparator() string { return "::" }

// TimeDuration is a wrapper around time.Duration that allows us to unmarshal in TOML.
type TimeDuration time.Duration

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *TimeDuration) UnmarshalText(text []byte) error {
	parsedDuration, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = TimeDuration(parsedDuration)
	return nil
}

// Hash calculates  the keccak hash of elements.
func Hash(data ...[]byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	for _, d := range data {
		hash.Write(d[:]) //nolint:errcheck,gosec
	}
	return hash.Sum(nil)
}
