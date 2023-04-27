package trie

import (
	"os"
	"fmt"
	"testing"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/Lagrange-Labs/lagrange-node/test"
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

func TestTrie(t *testing.T) {
  ptrie := NewTrie()
  
  fmt.Println("hash:",hexutil.Encode(ptrie.Hash()))
  test.expectString(T, "0x", hexutil.Encode(ptrie.Hash()))
  
  ptrie.Add([]byte("alice"))
  fmt.Println("hash:",hexutil.Encode(ptrie.Hash()))
  test.expectString(T, "0x43ade07dc3c905adc2612992de1c5ace7320686d7bc1755b4333e4bf229e2b25", hexutil.Encode(ptrie.Hash()))
  
  ptrie.Add([]byte("bob"))
  fmt.Println("hash:",hexutil.Encode(ptrie.Hash()))
  
  ptrie.Add([]byte("carol"))
  fmt.Println("hash:",hexutil.Encode(ptrie.Hash()))
}