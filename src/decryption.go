package main

import (
	chacha20 "golang.org/x/crypto/chacha20poly1305"
	box "golang.org/x/crypto/nacl/box"
	ascon "lukechampine.com/ascon"
)

func curveDec(ciphertext []byte, senderPublicKey, receiverPrivateKey *[32]byte) ([]byte, error) {
	// implement curve25519 (asymmertic) for decryption
	// check the length of ciphertext --- valid encryption message
	if len(ciphertext) < (24 + box.Overhead) {
		return nil, ErrDecrypt
	}

	// create shared key for decryption
	var sharedDecKey [32]byte
	box.Precompute(&sharedDecKey, senderPublicKey, receiverPrivateKey)

	// use the same nonce for decryption
	var decryptNonce [24]byte
	copy(decryptNonce[:], ciphertext[:24])
	// OpenAfterPrecomputation(out, box []byte, nonce *[24]byte, sharedKey *[32]byte) ([]byte, bool)
	decrypted, ok := box.OpenAfterPrecomputation(nil, ciphertext[24:], &decryptNonce, &sharedDecKey)
	if !ok {
		return nil, ErrDecrypt
	}
	return decrypted, nil
}

func asconDec(nonce, key [16]byte, cipher []byte) ([]byte, error) {
	/*
		:param nonce: the same nonce used for encryption
		:param key: the same key used for encryption
		:param cipher: the encrypted message
	*/
	// check the length of ciphertext --- valid encryption message
	if len(cipher) < (16 + box.Overhead) {
		return nil, ErrDecrypt
	}
	aead, _ := ascon.New(key)
	plaintext, err := aead.Open(nil, nonce, cipher, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func cha20Dec(nonce, key [32]byte, cipher []byte) ([]byte, error) {
	/*
		:param nonce: the same nonce as in encryption
		:param key: the same key used to encrypt"
		:param cipher: encrypted message
	*/
	aead, _ := chacha20.NewX(key[:])

	if len(cipher) < aead.NonceSize() {
		return nil, ErrDecrypt
	}

	plaintext, _ := aead.Open(nil, nonce[:], cipher, nil)

	return plaintext, nil
}
