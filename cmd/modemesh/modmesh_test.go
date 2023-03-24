package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Lagrange-Labs/Lagrange-Node/bcclients"
	"github.com/Lagrange-Labs/Lagrange-Node/node"
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
	CleanupWallets()
	return nil
}

// Cleanup test environment

func tearDown() error {
	return nil
}

func CleanupWallets() {
	files, err := filepath.Glob("./test/wallets/*")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

// Common assertion wrappers

func assert(t *testing.T, cond bool, desc string) {
	if cond == false {
		t.Errorf(desc)
	}
}

func expectString(t *testing.T, a string, b string) {
	assert(t, a == b, "Expected '"+a+"', got '"+b+"'.")
}

// Testing keccak hashing wrapper function
func TestKeccakHashString(t *testing.T) {
	// Remember that hex-encoded Keccak hashes return a '0x' prefix.
	NullKeccakHash := "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
	KeccakHash := bcclients.KeccakHashString("")

	expectString(t, KeccakHash, NullKeccakHash)
}

// Test instantiation of a LagrangeNode object
func testNewLagrangeNode(t *testing.T) *node.LagrangeNode {
	n := node.NewLagrangeNode()
	assert(t, n != nil, "LagrangeNode object is nil")
	return n
}
func TestNewLagrangeNode(t *testing.T) {
	testNewLagrangeNode(t)
}

// Test initialization of LagrangeNode
// func testInitializeLagrangeNode(t *testing.T) *node.LagrangeNode {
// 	n := testNewLagrangeNode(t)

// 	cfg := &node.Config{}
// 	cfg.Port = "8090"
// 	cfg.StakingEndpoint = "http://0.0.0.0:8545"
// 	cfg.AttestEndpoint = "http://0.0.0.0:8545"
// 	cfg.StakingWS = "ws://0.0.0.0:8545"
// 	cfg.LogLevel = 5

// 	n.SetWalletPath("./test/wallets/")
// 	n.GenerateAccountFromPrivateKeyString("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")

// 	go n.Start(cfg)
// 	return n
// }

// TODO: create the integration test for this
// func TestInitializeLagrangeNode(t *testing.T) {
// 	n := testInitializeLagrangeNode(t)
// 	time.Sleep(1 * time.Second)
// 	//	n.Stop()
// 	_ = n
// }

//
