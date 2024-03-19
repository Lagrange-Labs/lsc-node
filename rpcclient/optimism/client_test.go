package optimism

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSetBeginBlockNumber(t *testing.T) {
	cfg := &Config{
		L1RPCURL:    "http://localhost:8545",
		RPCURL:      "http://localhost:8545",
		BatchInbox:  "0xFf00000000000000000000000000000000008453",
		BatchSender: "0x5050f69a9786f081509234f1a7f4684b5e5b76c9",
	}

	c, err := NewClient(cfg)
	require.NoError(t, err)

	beginBlockNumber := uint64(10)
	b := uint64(0)
	c.SetBeginBlockNumber(beginBlockNumber)
	for i := 0; i < 100; i++ {
		b = c.fetcher.GetFetchedBlockNumber()
		if b > beginBlockNumber {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	c.SetBeginBlockNumber(b + ParallelBlocks + 10)
	for i := 0; i < 100; i++ {
		b2 := c.fetcher.GetFetchedBlockNumber()
		if b2 > 0 {
			require.Greater(t, b2, b+1)
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
}
