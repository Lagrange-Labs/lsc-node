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
	"crypto/rand"
	"crypto/subtle"
	binary "encoding/binary"
	"errors"
	"io"
	"math/big"
	"time"

	"github.com/kisom/sbuf"
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
	// public *[32]byte
	public ed25519.PublicKey
	// private *[64]byte
	private ed25519.PrivateKey

	// list of peers with pointers -- peerid in libp2p
	peers []*[ed25519.PublicKeySize]byte
	// Anonymous function for looking up peer
	peerLookup func(*[ed25519.PublicKeySize]byte) bool
}

// define messages to be encrypted during transmission

// =============================================================

// need to change as required
type EdDSASet struct {
	// public  []byte
	// private []byte
}

// need to change as required
type ECDSASet struct {
	V, R, S *big.Int
}

// define content part
type Content struct {
	// hashes for json encoding/decoding
	StateRoot []byte `json:"stateroot"`

	TimeStamp time.Time `json:"timestamp"`

	BlockNumber uint64 `json:"blocknumber"`

	ChainName string `json:"chainname"`

	// sharded EdDSA signature tuple
	EdDSAPair EdDSASet `json:"edddsaset"`

	// ecdsa signature tuple
	ECDSAPair ECDSASet `json:"ecdsaset"`

	// eth public key
	EthereumPublicKey *[]byte `json:"ethereumpublickey"`
}

type Message struct {
	Content []byte
	// used to track the sequence --- four bytes default
	Number uint32 `json:"number"`
}

// =============================================================

// Marshal serialises a copy of the Identity. It is intended to support
// persistent Identities.
func Marshal(id *Identity) []byte {
	bufSize := ed25519.PrivateKeySize + ed25519.PublicKeySize
	bufSize += len(id.peers) * ed25519.PublicKeySize
	buf := sbuf.NewBuffer(bufSize)
	buf.Write(id.private[:])
	buf.Write(id.public[:])
	for i := range id.peers {
		buf.Write(id.peers[i][:])
	}

	return buf.Bytes()
}

// ErrInvalidIdentity is returned if the Identity being unmarshalled
// is invalid.
var ErrInvalidIdentity = errors.New("sessions: invalid identity")

// Unmarshal parses an identity.
func Unmarshal(in []byte) (*Identity, error) {
	buf := sbuf.NewBufferFrom(in)
	id := &Identity{}

	_, err := io.ReadFull(buf, id.private[:])
	if err != nil {
		return nil, ErrInvalidIdentity
	}

	_, err = io.ReadFull(buf, id.public[:])
	if err != nil {
		return nil, ErrInvalidIdentity
	}

	if (buf.Len() % ed25519.PublicKeySize) != 0 {
		return nil, ErrInvalidIdentity
	}

	for {
		if buf.Len() == 0 {
			break
		}

		peer := new([ed25519.PublicKeySize]byte)
		io.ReadFull(buf, peer[:])
		id.peers = append(id.peers, peer)
	}

	return id, nil
}

func MarshalMessage(m Message) []byte {
	out := make([]byte, 4, len(m.Content)+4)
	binary.BigEndian.PutUint32(out[:4], m.Number)
	return append(out, m.Content...)
}

// parse a message from byte slice in order to decrypt
func UnmarshalMessage(in []byte) (Message, bool) {
	m := Message{}
	if len(in) <= 4 {
		return m, false
	}

	m.Number = binary.BigEndian.Uint32(in[:4])
	m.Content = in[4:]
	return m, true
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

// Public exports a copy of the Identity's public key.
func (id *Identity) Public() *[ed25519.PublicKeySize]byte {
	pub := new([ed25519.PublicKeySize]byte)
	copy(pub[:], id.public[:])
	return pub
}

func (s *Session) LastSent() uint32 {
	return s.lastSent
}

func (s *Session) LastRecv() uint32 {
	return s.lastRecv
}

/* sessionkey is a long combined key of the identity public key,
the session key, and the signature size */

// count only key size including priv exchange key
const SessionKeySize = ed25519.PublicKeySize + 64 + ed25519.SignatureSize

// information needs to be signed
const blobDataSize = ed25519.PublicKeySize + 64

// NewIdentity generates a new identity.
func NewIdentity() (*Identity, error) {
	var err error
	id := &Identity{}
	id.public, id.private, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// create a new session, return with signed key
func (id *Identity) NewSession() (*[SessionKeySize]byte, *Session, error) {

	// define the context of a Session
	sess := &Session{
		// define pointer type
		recvKey: new([32]byte),
		sendKey: new([32]byte),
	}

	// create key pair
	pub, priv, err := GenerateKeyPairs()
	if err != nil {
		return nil, nil, err
	}
	// store the key until the session is complete (one key for one session)
	sess.priv = priv

	// build the messsage that will be encrypted
	blobtosign := new([SessionKeySize]byte)
	copy(blobtosign[:], id.public[:])
	copy(blobtosign[ed25519.PublicKeySize:], pub[:])

	// sign private key + message with ed25519
	sig := ed25519.Sign(id.private, blobtosign[:blobDataSize])
	copy(blobtosign[blobDataSize:], sig[:])

	return blobtosign, sess, nil
}

// implement ECDH key exchange
func KeyExchange(shared *[32]byte, priv, pub []byte) {
	/*
		:param shared:
		:param priv:
		:param pub:
	*/
	// copy private key and wipe it --- no longer needed
	var kexPriv [32]byte
	// des, src
	copy(kexPriv[:], priv)
	Zero(priv)

	var kexPub [32]byte
	copy(kexPub[:], pub)
	// the shared key can be used to encrypt the messages when using the
	// same pair of keys repeatedly
	box.Precompute(shared, &kexPub, &kexPriv)
	Zero(kexPriv[:])
}

func (s *Session) ChangeKeys(peer *[64]byte, dialer bool) {
	/*
		param priv: private key
		param peer: hashed public key == peer id
		param dialer: default true, initiate the conversation
	*/
	if dialer {
		KeyExchange(s.sendKey, s.priv[:32], peer[:32])
		KeyExchange(s.recvKey, s.priv[32:], peer[32:])
	} else {
		KeyExchange(s.recvKey, s.priv[:32], peer[:32])
		KeyExchange(s.sendKey, s.priv[32:], peer[32:])
	}

	s.lastSent = 0
	s.lastRecv = 0
}

// define verification process for authentication
func (id *Identity) VerifySessionKey(sk *[SessionKeySize]byte) (*[64]byte, bool) {

	// verify the peerid --- publickey
	peerID := new([ed25519.PublicKeySize]byte)
	keyData := new([64]byte)
	signature := new([ed25519.SignatureSize]byte)
	copy(peerID[:], sk[:ed25519.PublicKeySize])
	copy(keyData[:], sk[ed25519.PublicKeySize:ed25519.PublicKeySize+64])
	copy(signature[:], sk[blobDataSize:])

	var found bool
	for i := range id.peers {
		if subtle.ConstantTimeCompare(id.peers[i][:], peerID[:]) == 1 {
			found = true
		}
	}
	// process not found
	if !found {
		if id.peerLookup != nil {
			if !id.peerLookup(peerID) {
				return nil, false
			}
		} else {
			return nil, false
		}
	}

	// verify the signature
	if !ed25519.Verify(peerID[:], sk[:blobDataSize], signature[:]) {
		return nil, false
	}

	return keyData, true
}

// ErrVerification is returned when the session key for a peer could
// not be authenticated.
var ErrVerification = errors.New("sessions: could not authenticate peer")

// function to set up a new session over a channel
func (id *Identity) Dial(ch Channel) (*Session, error) {
	/*
		:param ch: channel used to communicate

		return: *Session --- verify the status

	*/
	// sender(identity) create new session
	sk, s, err := id.NewSession()
	if err != nil {
		return nil, err
	}

	// write identity related information to channel
	if _, err = ch.Write(sk[:]); err != nil {
		return nil, err
	}

	sk = new([SessionKeySize]byte)
	if _, err = io.ReadFull(ch, sk[:]); err != nil {
		return nil, err
	}

	// verify the signed key
	peer, ok := id.VerifySessionKey(sk)
	if !ok {
		return nil, ErrVerification
	}

	// perform key exchange (after the exchange of public keys)
	s.ChangeKeys(peer, true)
	s.Channel = ch
	return s, nil
}

// function to access incoming session and to eastablish a new session
func (id *Identity) Listen(ch Channel) (*Session, error) {
	// read session key
	sk := new([SessionKeySize]byte)
	if _, err := io.ReadFull(ch, sk[:]); err != nil {
		return nil, err
	}

	// verify the identity
	peer, ok := id.VerifySessionKey(sk)
	if !ok {
		return nil, ErrVerification
	}

	// create new session
	sk, s, err := id.NewSession()
	if err != nil {
		return nil, err
	}

	// write session key
	if _, err = ch.Write(sk[:]); err != nil {
		return nil, err
	}

	// exchange key
	s.ChangeKeys(peer, false)
	s.Channel = ch
	return s, nil
}

// Encrypt adds a message number to the session and secures it with a
// symmetric ciphersuite. The message cannot be empty.
func (s *Session) Encrypt(message []byte) ([]byte, error) {
	if len(message) == 0 {
		return nil, ErrEncrypt
	}
	s.lastSent++
	m := MarshalMessage(Message{message, s.lastSent})
	return Cha20Enc(s.sendKey, m)
}

// Send encrypts the message and sends it out over the channel.
func (s *Session) Send(message []byte) error {
	m, err := s.Encrypt(message)
	if err != nil {
		return err
	}

	err = binary.Write(s.Channel, binary.BigEndian, uint32(len(m)))
	if err != nil {
		return err
	}

	_, err = s.Channel.Write(m)
	return err
}

// Decrypt recovers the message and verifies the message number. If the
// message number hasn't incremented, it's considered a decryption
// failure.
func (s *Session) Decrypt(message []byte) ([]byte, error) {
	out, err := Cha20Dec(s.recvKey, message)
	if err != nil {
		return nil, err
	}

	m, ok := UnmarshalMessage(out)
	if !ok {
		return nil, ErrDecrypt
	}

	if m.Number <= s.lastRecv {
		return nil, ErrDecrypt
	}

	s.lastRecv = m.Number

	return m.Content, nil
}

// Receive listens for a new message on the channel.
func (s *Session) Receive() ([]byte, error) {
	var mlen uint32
	err := binary.Read(s.Channel, binary.BigEndian, &mlen)
	if err != nil {
		return nil, err
	}
	message := make([]byte, int(mlen))
	_, err = io.ReadFull(s.Channel, message)
	if err != nil {
		return nil, err
	}

	return s.Decrypt(message)
}

// Close zeroises the keys in the session. Once a session is closed,
// the traffic that was sent over the channel can no longer be decrypted
// and any attempts at sending or receiving messages over the channel
// will fail.
func (s *Session) Close() error {
	Zero(s.sendKey[:])
	Zero(s.recvKey[:])
	return nil
}

// RandBytes attempts to read the selected number of bytes from the
// operating system PRNG.
func RandBytes(n int) ([]byte, error) {
	r := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Zero attempts to zeroise a byte slice.
func Zero(in []byte) {
	for i := range in {
		in[i] = 0
	}
}
