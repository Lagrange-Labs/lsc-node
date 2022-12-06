package main

import (
	"bytes"
	"testing"

	"crypto/ed25519"

	"github.com/kisom/testio"
)

var alice, bob, carol *Identity

func TestNewIdentity(t *testing.T) {
	t.Logf("TestNewIdentity started")
	var err error
	alice, err = NewIdentity()
	if err != nil {
		t.Fatalf("%v", err)
	}
	bob, err = NewIdentity()
	if err != nil {
		t.Fatalf("%v", err)
	}
	alice.AddPeer(bob.Public())
	bob.AddPeer(alice.Public())
	bob.AddPeer(alice.Public())
	if len(bob.peers) != 1 {
		t.Fatal("duplicate peers added")
	}
	carol, err = NewIdentity()
	if err != nil {
		t.Fatalf("%v", err)
	}
	carol.AddPeer(bob.Public())

	aliceOut := Marshal(alice)
	if _, err = Unmarshal(aliceOut); err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("TestNewIdentity completed successfully")
}

func TestUntrustedIdentity(t *testing.T) {
	t.Logf("TestUntrustedIdentity started")
	conn := testio.NewBufferConn()
	sk, _, err := bob.NewSession()
	if err != nil {
		t.Fatalf("%v", err)
	}

	conn.WritePeer(sk[:])
	_, err = carol.Dial(conn)
	if err != nil {
		t.Fatalf("%v", err)
	}

	var csk [SessionKeySize]byte
	_, err = conn.ReadClient(csk[:])
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, ok := bob.VerifySessionKey(&csk)
	if ok {
		t.Fatal("carol should not be trusted by bob")
	}
	t.Logf("TestUntrustedIdentity completed successfully")
}

func TestPeerLookup(t *testing.T) {
	t.Logf("TestPeerLookup started")
	bob.peerLookup = func(k *[ed25519.PublicKeySize]byte) bool {
		return false
	}

	conn := testio.NewBufferConn()
	sk, _, err := bob.NewSession()
	if err != nil {
		t.Fatalf("%v", err)
	}

	conn.WritePeer(sk[:])
	_, err = carol.Dial(conn)
	if err != nil {
		t.Fatalf("%v", err)
	}

	var csk [SessionKeySize]byte
	_, err = conn.ReadClient(csk[:])
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, ok := bob.VerifySessionKey(&csk)
	if ok {
		t.Fatal("carol should not be trusted by bob")
	}

	bob.peerLookup = func(k *[ed25519.PublicKeySize]byte) bool {
		return true
	}

	conn = testio.NewBufferConn()
	sk, _, err = bob.NewSession()
	if err != nil {
		t.Fatalf("%v", err)
	}

	conn.WritePeer(sk[:])
	_, err = carol.Dial(conn)
	if err != nil {
		t.Fatalf("%v", err)
	}

	Zero(csk[:])
	_, err = conn.ReadClient(csk[:])
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, ok = bob.VerifySessionKey(&csk)
	if !ok {
		t.Fatal("carol should be trusted by bob")
	}

	bob.peerLookup = nil
	t.Logf("TestPeerLookup completed successfully")
}

var m = []byte(`Hello, I am testing!`)

func TestDial(t *testing.T) {
	t.Logf("TestDial started")
	conn := testio.NewBufferConn()
	sk, bs, err := bob.NewSession()
	if err != nil {
		t.Fatalf("%v", err)
	}

	conn.WritePeer(sk[:])
	as, err := alice.Dial(conn)
	if err != nil {
		t.Fatalf("%v", err)
	}

	var ask [SessionKeySize]byte
	_, err = conn.ReadClient(ask[:])
	if err != nil {
		t.Fatalf("%v", err)
	}

	peer, ok := bob.VerifySessionKey(&ask)
	if !ok {
		t.Fatal("alice wasn't trusted by bob")
	}

	bs.ChangeKeys(peer, false)
	buf := &bytes.Buffer{}
	as.Channel = buf
	bs.Channel = as.Channel

	err = as.Send(m)
	if err != nil {
		t.Fatalf("%v", err)
	}

	// TLA intercepted a message.
	first := buf.Bytes()

	rcv, err := bs.Receive()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !bytes.Equal(rcv, m) {
		t.Fatal("bob didn't get the right message")
	}

	for i := 0; i < 5; i++ {
		err = as.Send(m)
		if err != nil {
			t.Fatalf("%v", err)
		}

		rcv, err = bs.Receive()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !bytes.Equal(rcv, m) {
			t.Fatal("bob didn't get the right message")
		}
	}

	// TLA tries to replay message.
	as.Channel.Write(first)
	_, err = bs.Receive()
	if err == nil {
		t.Fatal("TLA wins.")
	}
	// \o/

	bs.Close()
	as.Close()
	t.Logf("TestDial completed successfully")
}

func TestListen(t *testing.T) {
	t.Logf("TestListen started")
	conn := testio.NewBufferConn()
	sk, bs, err := bob.NewSession()
	if err != nil {
		t.Fatalf("%v", err)
	}

	conn.WritePeer(sk[:])
	as, err := alice.Listen(conn)
	if err != nil {
		t.Fatalf("%v", err)
	}

	var ask [SessionKeySize]byte
	_, err = conn.ReadClient(ask[:])
	if err != nil {
		t.Fatalf("%v", err)
	}

	peer, ok := bob.VerifySessionKey(&ask)
	if !ok {
		t.Fatal("alice wasn't trusted by bob")
	}

	bs.ChangeKeys(peer, true)
	buf := &bytes.Buffer{}
	as.Channel = buf
	bs.Channel = as.Channel

	err = as.Send(m)
	if err != nil {
		t.Fatalf("%v", err)
	}

	// TLA intercepted a message.
	first := buf.Bytes()

	rcv, err := bs.Receive()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !bytes.Equal(rcv, m) {
		t.Fatal("bob didn't get the right message")
	}

	for i := 0; i < 5; i++ {
		err = as.Send(m)
		if err != nil {
			t.Fatalf("%v", err)
		}

		rcv, err = bs.Receive()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !bytes.Equal(rcv, m) {
			t.Fatal("bob didn't get the right message")
		}
	}

	// TLA tries to replay message.
	as.Channel.Write(first)
	_, err = bs.Receive()
	if err == nil {
		t.Fatal("TLA wins.")
	}
	// \o/

	bs.Close()
	as.Close()
	t.Logf("TestListen completed successfully")
}
