package main

import (
	box "golang.org/x/crypto/nacl/box"
	ascon "lukechampine.com/ascon"
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

func asconDec(nonce, key [16]byte, cipher []byte) []byte {
	/*
		:param nonce: the same nonce used for encryption
		:param key: the same key used for encryption
		:param cipher: the encrypted message
	*/
	aead, _ := ascon.New(key)
	plaintext, err := aead.Open(nil, nonce, cipher, nil)
	if err != nil {
		panic(err)
	}
	return plaintext
}

func cha20Dec(nonce, key [32]byte, cipher []byte) []byte {
	aead, _ := ascon.NewX(key[:])
	plaintext, _ := aead.Open(nil, nonce, cipher, nil)

	return plaintext
}
