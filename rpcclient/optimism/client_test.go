package optimism

import (
	"sync"
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
	mtx := sync.Mutex{}
	mtx.Lock()
	go func() {
		defer mtx.Unlock()
		time.Sleep(2 * time.Second)
		client.fetcher.StopFetch()
	}()
	mtx.Lock()

	// check if able to restart
	client.SetBeginBlockNumber(50)
	mtx.Unlock()
	mtx.Lock()
	go func() {
		defer mtx.Unlock()
		time.Sleep(2 * time.Second)
		client.fetcher.Stop()
	}()
	mtx.Lock()
	defer mtx.Unlock()
	// check error propagation
	_, err = client.NextBatch()
	require.Error(t, err)
}
