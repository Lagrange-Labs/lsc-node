package optimism

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestErrorHandling(t *testing.T) {
	cfg := &Config{
		RPCURL:             "http://localhost:8545",
		L1RPCURL:           "http://localhost:8545",
		BeaconURL:          "http://localhost:8545",
		BatchInbox:         common.Address{}.Hex(),
		BatchSender:        common.Address{}.Hex(),
		ConcurrentFetchers: 4,
	}
	client, err := NewClient(cfg)
	require.NoError(t, err)

	client.SetBeginBlockNumber(10)
	go func() {
		time.Sleep(2 * time.Second)
		client.fetcher.Stop()
	}()
	_, err = client.NextBatch()
	require.Error(t, err)
	// check if able to restart
	client.SetBeginBlockNumber(10)
}
