package v2

import (
	"bytes"
	"testing"

	"github.com/Lagrange-Labs/lsc-node/core"
)

func TestBatchHeaderHash(t *testing.T) {
	testCases := []struct {
		batchHeader *BatchHeader
		expected    string
	}{
		{
			&BatchHeader{
				ChainId:       1,
				L1BlockNumber: 1,
				L1TxHash:      "0x123",
				L2Blocks:      []*BlockHeader{},
			},
			"241a85723c1478bebd111469b3282335149bbc72ef42445d97de235c9e4957de",
		},
		{
			&BatchHeader{
				ChainId:       10,
				L1BlockNumber: 12345,
				L1TxHash:      "7a6f632997d9346b4c51b347f1e34c57283825cbedc2df253d79cf963ef6efc5",
				L2Blocks: []*BlockHeader{
					{
						BlockNumber: 1,
						BlockHash:   "68d891948c3ef3128440bee226475e090e21e6d03a846f7e1b2cb30226f75f52",
					},
					{
						BlockNumber: 2,
						BlockHash:   "f7e1b2cb30226f75f5268d891948c3ef3128440bee226475e090e21e6d03a846",
					},
				},
			},
			"e32f1b5b40547285278235d6a015ddaf94802f39a178ae947f4ae565056b70dd",
		},
		{
			&BatchHeader{
				ChainId:       42161,
				L1BlockNumber: 1234567,
				L1TxHash:      "3330a380ddc29466a604b8636ecf463e83d3ac884fe05aa2a443fbe5c28d35e5",
				L2Blocks: []*BlockHeader{
					{
						BlockNumber: 4567890,
						BlockHash:   "f9341a456b8e42557ba3540ad93a7f36141a94692fd5020cd7f57e631aba0c7d",
					},
				},
			},
			"7e690cf00737001ed131d1bb98ba03b8ae1308c0d8f18ed32305ab3e5ae746d0",
		},
	}

	for _, tc := range testCases {
		if got := tc.batchHeader.Hash(); !bytes.Equal(got, core.Hex2Bytes(tc.expected)) {
			t.Errorf("batchHeader.Hash() = %x; expected %s", got, tc.expected)
		}
	}
}

func TestBatchHeaderMerkleHash(t *testing.T) {
	testCases := []struct {
		batchHeader *BatchHeader
		expected    string
	}{
		{
			&BatchHeader{
				ChainId:       1,
				L1BlockNumber: 1,
				L1TxHash:      "0x123",
				L2Blocks:      []*BlockHeader{},
			},
			"ba8098a7067139daf67f8aa1fd87c0e06dda1acd41228bacd606623283dedc35",
		},
		{
			&BatchHeader{
				ChainId:       10,
				L1BlockNumber: 12345,
				L1TxHash:      "7a6f632997d9346b4c51b347f1e34c57283825cbedc2df253d79cf963ef6efc5",
				L2Blocks: []*BlockHeader{
					{
						BlockNumber: 1,
						BlockHash:   "68d891948c3ef3128440bee226475e090e21e6d03a846f7e1b2cb30226f75f52",
					},
					{
						BlockNumber: 2,
						BlockHash:   "f7e1b2cb30226f75f5268d891948c3ef3128440bee226475e090e21e6d03a846",
					},
				},
			},
			"3c4b00ebbca32053f1d25f603b7585779db5a8199ab15b3e9a81fe4ca05fdb2f",
		},
		{
			&BatchHeader{
				ChainId:       42161,
				L1BlockNumber: 1234567,
				L1TxHash:      "3330a380ddc29466a604b8636ecf463e83d3ac884fe05aa2a443fbe5c28d35e5",
				L2Blocks: []*BlockHeader{
					{
						BlockNumber: 4567890,
						BlockHash:   "f9341a456b8e42557ba3540ad93a7f36141a94692fd5020cd7f57e631aba0c7d",
					},
				},
			},
			"0aca845d6950dc312d4a835d700c91db78d244e36a5420ed42c3f87a0785995a",
		},
	}

	for _, tc := range testCases {
		if got := tc.batchHeader.MerkleHash(); !bytes.Equal(got, core.Hex2Bytes(tc.expected)) {
			t.Errorf("batchHeader.Hash() = %x; expected %s", got, tc.expected)
		}
	}
}
