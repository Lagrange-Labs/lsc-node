package crypto

import (
	"bytes"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/utils"
)

func TestNextPowerOfTwo(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{255, 256},
		{256, 256},
		{1234, 2048},
		{65535, 65536},
		{65536, 65536},
		{65537, 131072},
	}

	for _, tc := range testCases {
		if got := nextPowerOfTwo(tc.n); got != tc.expected {
			t.Errorf("nextPowerOfTwo(%d) = %d; expected %d", tc.n, got, tc.expected)
		}
	}
}

func TestNodeHash(t *testing.T) {
	// leafHash
	testCases := []struct {
		data     []byte
		expected string
	}{
		{[]byte{}, "5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2"},
		{[]byte{1, 2, 3}, "d4ff8d9d9a44c2b7b9c6a2defc4735f367b95e877ea7efc5f30970ebd56b6df1"},
	}

	for _, tc := range testCases {
		if got := leafHash(tc.data); !bytes.Equal(got, utils.Hex2Bytes(tc.expected)) {
			t.Errorf("leafHash(%x) = %x; expected %v", tc.data, got, tc.expected)
		}
	}

	// innerHash
	testCases1 := []struct {
		left     string
		right    string
		expected string
	}{
		{"5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2", "d4ff8d9d9a44c2b7b9c6a2defc4735f367b95e877ea7efc5f30970ebd56b6df1", "694a4d98579ba1c75f7dbd052a3a0642727e51a8334490cc39f70726bec69ccc"},
		{"5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2", "5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2", "cadc33d5ed150387ed3609c7e28156a06e28d359df6d93fab89b40f8390aa253"},
	}

	for _, tc := range testCases1 {
		if got := innerHash(utils.Hex2Bytes(tc.left), utils.Hex2Bytes(tc.right)); !bytes.Equal(got, utils.Hex2Bytes(tc.expected)) {
			t.Errorf("innerHash(%x, %x) = %x; expected %v", tc.left, tc.right, got, tc.expected)
		}
	}
}

func TestMerkleRoot(t *testing.T) {
	testCases := []struct {
		data     [][]byte
		expected string
	}{
		{[][]byte{}, "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"},
		{[][]byte{{1, 2, 3}}, "d4ff8d9d9a44c2b7b9c6a2defc4735f367b95e877ea7efc5f30970ebd56b6df1"},
		{[][]byte{{1, 2, 3}, {4, 5, 6}}, "50a5c8c83ab0942462e74dc6940bb54b6fc15a79d1a8da84799937f11f4ba43a"},
		{[][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "f07aa539b07550b33c7eef211520f3a9106b7c37299a7d814bc1500b0fb696c2"},
	}

	for _, tc := range testCases {
		if got := MerkleRoot(tc.data); !bytes.Equal(got, utils.Hex2Bytes(tc.expected)) {
			t.Errorf("MerkleRoot(%x) = %x; expected %v", tc.data, got, tc.expected)
		}
	}
}
