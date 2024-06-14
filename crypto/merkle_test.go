package crypto

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
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
		{[][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "a83a6fba6dc1ab15186f11d8dda138913b22635dc3b5e7a0a559c570e4a5323b"},
	}

	for _, tc := range testCases {
		if got := MerkleRoot(tc.data); !bytes.Equal(got, utils.Hex2Bytes(tc.expected)) {
			t.Errorf("MerkleRoot(%x) = %x; expected %v", tc.data, got, tc.expected)
		}
	}
}

func TestRootWithCommittee(t *testing.T) {
	committee := []struct {
		operator    string
		votingPower uint64
		publicKey   string
	}{
		{"0xFc4AC3204E0458967b758c902BcBBB88A3B7582f", 100, "2d6c60e91f6f4026d3637018d1b66636090a16ace9595d52e01381649f5a8bfb086ce92c7caece1b8efad25e5e23332771080c87c38ebe2805cd57d008980693"},
		{"0x03fbBE0CAe10ea5fe1D9E9a37b57F719574B48D4", 100, "0dae9aed87b0ff66b31048db9e82093c0288c1c20f197501ba1fb53b60aba57c1bb9d89335453c4516dc6c347949c9cfcb0cc1158ffb43204cd0912e3443dd76"},
		{"0x4736933A5B78C6fEba7635Bc57117f758c34d424", 100, "2b67946b6e46b37d20e0c5ce176510853fcc1a4c1f07764b151c83cc0da17ac92dbaf767b87e784aecec6eb6f381f92bf4e644c811d2d7bf9f1de5a605cf5e77"},
	}

	leaves := make([][]byte, 0, len(committee))
	for _, c := range committee {
		res := make([]byte, 0, 32+32+20+12)
		res = append(res, utils.Hex2Bytes(c.publicKey)...)
		res = append(res, utils.Hex2Bytes(c.operator)...)
		res = append(res, common.LeftPadBytes(big.NewInt(int64(c.votingPower)).Bytes(), 12)...)
		leaves = append(leaves, res)
	}
	root := MerkleRoot(leaves)
	expected := "2d64f451e549526f76e0e6b9cf724c229298998b2d9fdfa0d96efbf862685915"
	if !bytes.Equal(root, utils.Hex2Bytes(expected)) {
		t.Errorf("MerkleRoot(%v) = %x; expected %v", leaves, root, expected)
	}
}
