package main

/* Implement multiple encryption methods to do a benchmark

What has been considered in this implementation:
	- compatibility with libp2p protocol, especially the peer Id and signatures
	- latency

*/

import (
	rand "crypto/rand"
	sha256 "crypto/sha256"
	"errors"
	"io"

	chacha20 "golang.org/x/crypto/chacha20poly1305"
	box "golang.org/x/crypto/nacl/box"
	ascon "lukechampine.com/ascon"
)

type AEAD struct {
	key []byte
}

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: decryption failed")
)

// implement ascon as the symmetric encryption method
func asconEnc(msg []byte, key [16]byte) ([]byte, []byte, error) {
	/*
			KeySize=16
			Nonce=16
		:param msg: the message to encrypt
		:param key: the given 16 bytes key pointer type in order to change
	*/
	// generate random key

	// keysize is 16 bytes
	// key := make([]byte, ascon.KeySize)
	aead, _ := ascon.New(key)
	// noncesize is 16 bytes
	nonce := make([]byte, ascon.NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, nil, err
	}

	ciphertext := aead.Seal(nil, nonce, msg, nil)
	return ciphertext, nonce, nil

}

// implement chacha20 as the symmetric encrytion method
func cha20Enc(msg []byte) ([]byte, error) {
	/*
			KeySize = 32
			NonceSize = 12
		:param msg: the message to encrypt
	*/
	// generate key by hashing with sha256 --- 32 bytes
	key := sha256.Sum256([]byte(msg))
	// create symmertic key
	aead, _ := chacha20.NewX(key[:])
	// make nonce
	nonce := make([]byte, chacha20.NonceSize)
	// encrypt
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	ciphertext := aead.Seal(nil, nonce, msg, nil)

	return ciphertext, nil
}

func curveEnc(msg []byte, senderPrivateKey, receiverPublicKey *[32]byte) ([]byte, error) {
	// implement curve25519 (asymmertic) for encryption

	// create shared key to speed up processing, new defines a pointer type
	sharedEncKey := new([32]byte)
	//Precompute(sharedKey, peersPublicKey, privateKey *[32]byte)
	box.Precompute(sharedEncKey, receiverPublicKey, senderPrivateKey)

	// create new nonce every time
	var nonce [24]byte

	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	// encrypt msg and appends the result to the nonce
	ciphertext := box.SealAfterPrecomputation(nonce[:], msg, &nonce, sharedEncKey)

	return ciphertext.nil
}
