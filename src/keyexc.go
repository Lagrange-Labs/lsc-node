package main

/*
	Implementation of authenticated key exchange
	ed25519-signed ECDH key exchange

	key exchange is the process to combine one part's public key and the
	other part's private key to generate a shared key ---- offline, then the
	shared key can be used for symmetric cipher

	session implemented ---- avoid reply attack, stateful communication
	signed session --- authentication for trust purpose

	features:
		- bidirectional secure channel
		- two parties in a session --- dialer and listener
		- a list of peers --- function to look up a peer and determine trust or not

	steps:
		- dialer send two public keys to listener
			- first public key for encryption from dialer to listener
			- second public key for encryption from listener to dialer

*/

import (
	"bytes"
	"crypto/ed25519"
	binary "encoding/binary"
	"io"
	"time"

	util "git.metacircular.net/kyle/gocrypto/util"

	ed25519 "github.com/agl/ed25519"
	box "golang.org/x/crypto/nacl/box"
)

type Channel io.ReadWriter

// track message number for a session
type Session struct {
	// used to track message number -- not replay
	lastSent uint32

	// keep seperate keys for each direction
	// key used to encrypt outgoing messages to the peer
	sendKey *[32]byte

	// track sent and received message separately
	lastRecv uint32

	// key used to decrypt incoming messages from peer
	recvKey *[32]byte

	// private key used in key exchange, wiped after exchanging
	priv *[64]byte

	Channel Channel
}

// identity contains the signature keypair for key change
// include a list of peer and func to look up / verify peer
type Identity struct {
	// key pair
	public  *[32]byte
	private *[64]byte
	// list of peers with pointers -- peerid in libp2p
	peers []*[ed25519.PublickKeySize]byte
	// look up function
	peerLookup func([]*[ed25519.PublickKeySize]byte) bool
}

// need to change as required
type EdDSASet struct {
	public  []byte
	private []byte
}

// need to change as required
type ECDSASet struct {
	V []byte
	R []byte
	S []byte
}

// define content part
type Content struct {
	// hashes for json encoding/decoding
	StateRoot []byte `json:"stateroot"`

	TimeStamp time.Time `json:"timestamp"`

	BlockNumber uint64 `json:"blocknumber"`

	ChainName string `json:"chainname"`

	// sharded EdDSA signature tuple
	EdDSASet `json:"edddsaset'`

	// ecdsa signature tuple
	ECDSASet `json:"ecdsaset"`

	// eth public key
	EthereumPublicKey *[]byte `json:"ethereumpublickey"`
}

type Message struct {
	Content
	// used to track the sequence --- four bytes default
	Number uint32 `json:"number"`
}

// serialise a message into a byte slice
func MarshalMessage(m Message) (*bytes.Buffer, error) {
	/*
		:param m: the struct of message going to be encrypted
	*/

	// implement with binary.BigEndian
	// get the size of content, use len here in order to add
	out := make([]byte, 4, binary.Size(m.Content)+4)
	buffer := bytes.NewBuffer(out)
	binary.Write(buffer, binary.BigEndian, m.Number)
	// interpret it as a byte array pointer then turn it into a slice
	err := binary.Write(buffer, binary.BigEndian, m.Content)
	if err != nil {
		return nil, err
	}

	return buffer, nil
	// implement with unsafe
	// out := make([]byte, 4, unsafe.Sizeof(m.Content) + 4)
	// contentSize := unsafe.Sizeof(m.Content)
	// out = append(out[:4], byte(m.Number))
	// // binary.BigEndian.PutUint32(out[:4], m.Number)

	// // interpret it as a byte array pointer then turn it into a slice
	// return append(out, (*[contentSize]byte)(unsafe.Pointer(&m.Content))[:])
}

// parse a message from byte slice
func UnmarshalMessage(cipher *bytes.Buffer) (Message, error) {
	// decode from binary implemetation
	m := Message{}
	binary.Read(cipher, binary.BigEndian, m.Number)
	err := binary.Read(cipher, binary.BigEndian, m.Content)
	if err != nil {
		return m, err
	}
	return m, nil

	// m := Message{}
	// if len(cipher) < 4 {
	// 	return nil, false
	// }
	// // return the number
	// m.Number = *(*[4]byte)(unsafe.Pointer(cipher[:4]))
	// // return the context
}

// add a new peer key to the Identity's peer list -- handler
func (id *Identity) AddPeer(peerId *[ed25519.PublicKeySize]byte) {
	// check whether new peer has existed in list
	for _, peer := range id.peers {
		if bytes.Equal(peer[:], peerId[:]) {
			return
		}
	}
	id.peers = append(id.peers, peerId)
}

// sessionkey is a long combined key of the identity public key,
// the session key, and the signature size

func (s *Session) LastSent() uint32 {
	return s.lastSent
}

func (s *Session) LastRecv() uint32 {
	return s.lastRecv
}

// count only key size including priv exchange key
const SessionKeySize = ed25519.PublicKeySize + 64 + ed25519.SignatureSize

const blobDataSize = ed25519.PublicKeySize + 64

// create a new session, return with signed key
func (id *Identity) NewSession() (*Session, error) {

	// define the context of a Session
	sess := &Session{
		// define pointer type
		receKey: new([32]byte),
		sendKey: new([32]byte),
		Channel: ch,
	}

	// create key pair
	GenerateKeyPairs()
	// sign private key + message with ed25519

}

// implement ECDH key exchange
func keyExchange(shared *[32]byte, priv, pub []byte) {
	/*
		:param shared:
		:param priv:
		:param pub:
	*/
	// copy private key and wipe it --- no longer needed
	var kexPriv [32]byte
	// des, src
	copy(kexPriv[:], priv)
	util.Zero(priv)

	var kexPub [32]byte
	copy(kexPub[:], pub)
	// the shared key can be used to encrypt the messages when using the
	// same pair of keys repeatedly
	box.Precompute(shared, &kexPub, &kexPriv)
	util.Zero(kexPriv[:])
}

func (s *Session) changeKeys(priv, peer *[64]byte, dialer bool) {
	/*
		param priv: private key
		param peer: hashed public key == peer id
		param dialer: default true, initiate the conversation
	*/
	if dialer {
		keyExchange(s.sendKey, priv[:32], peer[:32])
		keyExchange(s.recvKey, priv[32:], peer[32:])
	} else {
		keyExchange(s.recvKey, priv[:32], peer[:32])
		keyExchange(s.sendKey, priv[32:], peer[32:])
	}

	s.lastSent = 0
	s.lastRecv = 0
}

// define verification process for authentication
func (id *Identity) VerifySessionKey() {

}

// function to set up a new session over a channel
func Dial() {
	/*
		:param ch: channel used to communicate

		return: *Session --- verify the status

	*/

	// sender(identity) create new session

	// write identity related information to channel

	// verify the signed key

	// perform key exchange (after the exchange of public keys)

}

// function to access incoming session and to eastablish a new session
func Listen() {
	// read session key

	// src, dst
	io.ReadFull()

	// verify the identity

	// create new session

	// write session key

	// exchange key

}
