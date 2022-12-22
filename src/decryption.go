package main

import (
	sha256 "crypto/sha256"

	chacha20 "golang.org/x/crypto/chacha20poly1305"
	box "golang.org/x/crypto/nacl/box"
	ascon "lukechampine.com/ascon"
)

func CurveDec(ciphertext []byte, senderPublicKey, receiverPrivateKey *[32]byte) ([]byte, error) {
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

func AsconDec(nonce, key []byte, cipher []byte) ([]byte, error) {
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

func Cha20Dec(passphrase *[32]byte, cipher []byte) ([]byte, error) {
	/*
		:param nonce: the same nonce as in encryption
		:param key: the same key used to encrypt"
		:param cipher: encrypted message
	*/
	key := sha256.Sum256(passphrase[:])
	aead, _ := chacha20.NewX(key[:])
	// make nonce
	nonce, ciphertext := cipher[:aead.NonceSize()], cipher[aead.NonceSize():]

	if len(ciphertext) < aead.NonceSize() {
		return nil, ErrDecrypt
	}

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
