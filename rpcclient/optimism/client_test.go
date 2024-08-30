package optimism

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestErrorHandling(t *testing.T) {
	cfg := &Config{
		RPCURLs:            []string{"http://localhost:8545"},
		L1RPCURLs:          []string{"http://localhost:8545"},
		BeaconURL:          "http://localhost:8545",
		BatchInbox:         common.Address{}.Hex(),
		BatchSender:        common.Address{}.Hex(),
		ConcurrentFetchers: 4,
	}
	client, err := NewClient(cfg, true)
	require.NoError(t, err)

	client.SetBeginBlockNumber(10, 10)
	time.Sleep(1 * time.Second)
	client.fetcher.StopFetch()

	// check if able to restart
	client.SetBeginBlockNumber(50, 50)
	time.Sleep(1 * time.Second)
	client.fetcher.Stop()
	// check error propagation
	_, err = client.NextBatch()
	require.Error(t, err)
}
