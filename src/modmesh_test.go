package main

import (
	"os"
	"fmt"
	"testing"
	"time"
)

// Main test running function

func TestMain(m *testing.M) {
  if err := setup(); err != nil {
    os.Exit(1)
  }

  exitCode := m.Run()

  if err := tearDown(); err != nil {
    os.Exit(1)
  }

  fmt.Println("*DONE*")
  os.Exit(exitCode)
}

// Setup test environment

func setup() error {
  return nil
}

// Cleanup test environment

func tearDown() error {
  return nil
}

// Common assertion wrappers

func assert(t *testing.T, cond bool, desc string) {
	if cond == false {
		t.Errorf(desc)
	}
}

func expectString(t *testing.T, a string, b string) {
	assert(t, a == b,"Expected '"+a+"', got '"+b+"'.")
}

// Testing keccak hashing wrapper function
func TestKeccakHashString(t *testing.T) {
	// Remember that hex-encoded Keccak hashes return a '0x' prefix.
	NullKeccakHash := "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
	KeccakHash := KeccakHashString("")
	
	expectString(t, KeccakHash, NullKeccakHash)
}

// Test instantiation of a LagrangeNode object
func testNewLagrangeNode(t *testing.T) *LagrangeNode {
	n := NewLagrangeNode()
	assert(t, n != nil, "LagrangeNode object is nil")
	return n
}
func TestNewLagrangeNode(t *testing.T) {
	testNewLagrangeNode(t)
}

// Test initialization of LagrangeNode
func testInitializeLagrangeNode(t *testing.T) *LagrangeNode {
	n := testNewLagrangeNode(t)
	
	opts := &Opts{}
	opts.port = 8090
	opts.stakingEndpoint = "http://0.0.0.0:8545"
	opts.attestEndpoint = "http://0.0.0.0:8545"
	opts.stakingWS = "ws://0.0.0.0:8545"
	opts.logLevel = 5

	n.GenerateAccountFromPrivateKeyString("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	
	n.SetOpts(opts)
	go n.Start()
	return n
}
func TestInitializeLagrangeNode(t *testing.T) {
	n := testInitializeLagrangeNode(t)
	time.Sleep(1 * time.Second)
//	n.Stop()
	_ = n
}

// 
