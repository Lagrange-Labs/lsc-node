package v2

import (
	"bytes"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/utils"
)

/*
	"batch_header": {
	    "batch_number": 32000,
	    "chain_id": 11155420,
	    "l1_block_number": 6464398,
	    "l1_tx_hash": "0xfe403cda312e22f5d820814ebc6871c70092a2d28189b19cc28266ce440f1482",
	    "l1_tx_index": 70000,

...

	}
*/

func TestBatchHeaderHash(t *testing.T) {
	testCases := []struct {
		batchHeader *BatchHeader
		expected    string
	}{
		// gupeng
		{
			&BatchHeader{
				ChainId:       11155420,
				L1BlockNumber: 6464398,
				L1TxHash:      "0xfe403cda312e22f5d820814ebc6871c70092a2d28189b19cc28266ce440f1482",
				L2Blocks: []*BlockHeader{
					{
						BlockHash:   "0x48120c19ed719091789e06f6cae05b2eb682a86be9d13d2bead0862e5c9c43f0",
						BlockNumber: 15685020,
					},
					{
						BlockHash:   "0x86b71deeeb8fe28002509c12e07b79559cb59152271b4fe9aa57ea15cf104283",
						BlockNumber: 15685021,
					},
					{
						BlockHash:   "0xc4ddd47e9131968e8de0c09d94354b41456345048944da677ac4dd4d5ef2ab99",
						BlockNumber: 15685022,
					},
					{
						BlockHash:   "0x0a39cc25e333929649a6b40ba544a587f4482d85e3ae8ecdca13678ebeff5a32",
						BlockNumber: 15685023,
					},
					{
						BlockHash:   "0xac45b83fe85c7d21632e07448a973c5f83122033b20a35c6a2abdd58b3d2e91b",
						BlockNumber: 15685024,
					},
					{
						BlockHash:   "0x0444a0545e7cf52e1f99a838a3d8c4b70749ec2a0b727f02ccdee515e62ce12d",
						BlockNumber: 15685025,
					},
					{
						BlockHash:   "0x37e83795c14b3f089a7c3a06b16d78526c55e079aa68f7c3afaec583dcb3f8e4",
						BlockNumber: 15685026,
					},
					{
						BlockHash:   "0x1a8b84350efcb5ed3a636883054aad754c6f64998591d82f4152226aa17cd786",
						BlockNumber: 15685027,
					},
					{
						BlockHash:   "0xddc0764bd382cdaa95382d67ca09356f4c2da8e2cd1e07e75068d550fb5c6f18",
						BlockNumber: 15685028,
					},
					{
						BlockHash:   "0x65781311d4151d57b659c9bc893bc76b8cc107f9d4301696a9c4b5af6c9f2575",
						BlockNumber: 15685029,
					},
					{
						BlockHash:   "0xcee6d0239ac18840f8f479359d1a3c2a8d79acda0da39b26b457d70c6fd5a750",
						BlockNumber: 15685030,
					},
					{
						BlockHash:   "0xded03555c95e1ba89a285116ed7c3434757eb6a0fdf5a26fd4cc6572c3b799ca",
						BlockNumber: 15685031,
					},
					{
						BlockHash:   "0xe5fec57ca3e05e2b1ed7ff1ce72193a9f76a30e2f191032f77292d25852415d1",
						BlockNumber: 15685032,
					},
					{
						BlockHash:   "0x68942a0509582686e70f5e018f36bba8df25b19a94d4af69e9e8864a4bca9715",
						BlockNumber: 15685033,
					},
					{
						BlockHash:   "0x82a7a36cfa62d54b7dc8b1f733adbfe8837c4601ea86274334a0238a75e5ae2e",
						BlockNumber: 15685034,
					},
					{
						BlockHash:   "0xf6dd7f81a42e5b61b82931dec1253ba63ab158a64deeb43a2f5114523d202b0b",
						BlockNumber: 15685035,
					},
					{
						BlockHash:   "0x38f4cbb805c36c10be88928af7a5986c4a8a172ed590364b48dcb075bb0f9bf8",
						BlockNumber: 15685036,
					},
					{
						BlockHash:   "0xb45098c6fbd9d21bd29fc2b4f611d5470e6400d96a0596f1646bcd9c87d76e88",
						BlockNumber: 15685037,
					},
					{
						BlockHash:   "0x276fce711e2571f9a448247d61b8bea15445546c3acdc2bc1700fe0271ed3ace",
						BlockNumber: 15685038,
					},
					{
						BlockHash:   "0x9fd79e89d55073f5e976f473b74d5efb911bf24dd3a579102bf8de6a21630811",
						BlockNumber: 15685039,
					},
					{
						BlockHash:   "0x902433bfe2b435d7860a81c1fee2b393f893e6ea3257cdf6416e89eacd85247d",
						BlockNumber: 15685040,
					},
					{
						BlockHash:   "0x613ea16c12afee5b918f86709805d48ff757c2b1bad62754bc5e0e3f8b9da622",
						BlockNumber: 15685041,
					},
					{
						BlockHash:   "0x8c2c89b198e60a00ed7d09c6d15c1681fbeca14e3e876152c1dcafc22cba42ac",
						BlockNumber: 15685042,
					},
					{
						BlockHash:   "0x8b3e773fbdb2ba04d63691c8ecaa49e7fb82ea7931fb5eda3c36ad623e6f5947",
						BlockNumber: 15685043,
					},
					{
						BlockHash:   "0x138f4ef55e97193a371b09c42418da2913ebe32f6b68742df99c24a8d8a7a836",
						BlockNumber: 15685044,
					},
					{
						BlockHash:   "0x88da92dd002db674f8973250bfba7760b5808c86f033515fd43f5ff5c3c8dea7",
						BlockNumber: 15685045,
					},
					{
						BlockHash:   "0xd0efea4d90ff218c2665538cac289748b4260bb4b630a054f80f07cf0408f603",
						BlockNumber: 15685046,
					},
				},
			},
			"241a85723c1478bebd111469b3282335149bbc72ef42445d97de235c9e4957de",
		},
		/*
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
		*/
	}

	for _, tc := range testCases {
		// gupeng
		if got := tc.batchHeader.Hash(); !bytes.Equal(got, utils.Hex2Bytes(tc.expected)) {
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
		if got := tc.batchHeader.MerkleHash(); !bytes.Equal(got, utils.Hex2Bytes(tc.expected)) {
			t.Errorf("batchHeader.Hash() = %x; expected %s", got, tc.expected)
		}
	}
}
