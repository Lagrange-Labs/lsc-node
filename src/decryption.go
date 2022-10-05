package main

import (
	box "golang.org/x/crypto/nacl/box"
)

type AEAD struct {
	key []byte
}

func curveDec(ciphertext []byte, senderPublicKey, receiverPrivateKey *[32]byte) []byte {
	// implement curve25519 (asymmertic) for decryption

	// create shared key for decryption
	var sharedDecKey [32]byte
	box.Precompute(&sharedDecKey, senderPublicKey, receiverPrivateKey)

	// use the same nonce for decryption
	var decryptNonce [24]byte
	copy(decryptNonce[:], ciphertext[:24])
	// OpenAfterPrecomputation(out, box []byte, nonce *[24]byte, sharedKey *[32]byte) ([]byte, bool)
	decrypted, ok := box.OpenAfterPrecomputation(nil, ciphertext[24:], &decryptNonce, &sharedDecKey)
	if !ok {
		panic("decryption error")
	}
	return decrypted
}

func asconDec(cipher, nonce []byte, aead *AEAD) []byte {
	plaintext, err := &aead.Open(nil, nonce, cipher, nil)
	if err != nil {
		panic(err)
	}
	return plaintext
}
