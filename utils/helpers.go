package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"crypto/rand"

	context "context"

	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"

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
	LogMessage(fmt.Sprintf("discovered new peer %s", pi.ID.Pretty()), LOG_INFO)
	LogMessage("peer.AddrInfo: "+fmt.Sprintf("%v", pi), LOG_INFO)
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		LogMessage(fmt.Sprintf("error connecting to peer %s: %s", pi.ID.Pretty(), err), LOG_ERROR)
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

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	LOG_INFO    = 1
	LOG_NOTICE  = 2
	LOG_WARNING = 3
	LOG_ERROR   = 4
	LOG_DEBUG   = 5
)

func getTimestamp() string {
	time := time.Now()
	return time.Format("2006-01-02 15:04:05")
}

func LogMessage(message string, level int) {
	// if LOG_LEVEL < level {
	// 	return
	// }

	var color string
	var cat string
	switch level {
	case LOG_INFO:
		color = InfoColor
		cat = "INFO"
	case LOG_NOTICE:
		color = NoticeColor
		cat = "NOTICE"
	case LOG_WARNING:
		color = WarningColor
		cat = "WARN"
	case LOG_ERROR:
		color = ErrorColor
		cat = "ERROR"
	case LOG_DEBUG:
		color = DebugColor
		cat = "DEBUG"
	}
	fmt.Printf("%s %s [ %v ] %v\n", color, cat, getTimestamp(), message)
}

func Scan(prompt string) string {
	if prompt != "" {
		fmt.Println(prompt)
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
