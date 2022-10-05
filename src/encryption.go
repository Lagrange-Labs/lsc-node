package main

/* Implement multiple encryption methods to do a benchmark

What has been considered in this implementation:
	- compatibility with libp2p protocol, especially the peer Id and signatures
	- latency

*/

import (
	rand "crypto/rand"
	"io"

	box "golang.org/x/crypto/nacl/box"
	ascon "lukechampine.com/ascon"
)

type AEAD struct {
	key []byte
}

// implement ascon as the symmetric encryption method
func asconEnc(msg []byte) ([]byte, *AEAD) {
	key := make([]byte, ascon.KeySize)
	aead, _ := ascon.New(key)
	nonce := make([]byte, ascon.NonceSize)
	ciphertext := aead.Seal(nil, nonce, msg, nil)
	return ciphertext, &aead

}

// implement cha20 as the symmetric encrytion method
func cha20Enc() {

}

func curveEnc(msg []byte, senderPrivateKey, receiverPublicKey *[32]byte) []byte {
	// implement curve25519 (asymmertic) for encryption

	// create pairs for sender
	// senderPublicKey, senderPrivateKey, err := box.GenerateKey(rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	// // create pairs for receiver
	// receiverPublicKey, receiverPrivateKey, err := box.GenerateKey(rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	// create shared key to speed up processing, new defines a pointer type
	sharedEncKey := new([32]byte)
	//Precompute(sharedKey, peersPublicKey, privateKey *[32]byte)
	box.Precompute(sharedEncKey, receiverPublicKey, senderPrivateKey)

	// create new nonce every time
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	// encrypt msg and appends the result to the nonce
	ciphertext := box.SealAfterPrecomputation(nonce[:], msg, &nonce, sharedEncKey)

	return ciphertext
}
