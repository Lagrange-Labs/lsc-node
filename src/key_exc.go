package main

import (
	"io"
	"util"

	box "golang.org/x/crypto/nacl/box"
)

type Channel io.ReadWriter

type Session struct {
	lastSent uint32
	sendKey  *[32]byte

	lastRecv uint32
	recvKey  *[32]byte

	Channel Channel
}

func (s *Session) LastSent() uint32 {
	return s.lastSent
}

func (s *Session) LastRecv() uint32 {
	return s.lastRecv
}

func NewSession(ch Channel) *Session {
	return &Session{
		// define pointer type
		receKey: new([32]byte),
		sendKey: new([32]byte),
		Channel: ch,
	}
}

func keyExchange(shared *[32]byte, priv, pub []byte) {
	var kexPriv [32]byte
	copy(kexPriv[:], priv)
	util.Zero(priv)

	var kexPub [32]byte
	copy(kexPub[:], pub)

	box.Precompute(shared, &kexPub, &kexPriv)
	util.Zero(kexPriv[:])
}

func (s *Session) keyExchange(priv, peer *[64]byte, dialer bool) {
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
