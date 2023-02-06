package main

import (
	"os"
	"fmt"
	"testing"
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


