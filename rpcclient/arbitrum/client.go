package arbitrum

import (
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

var _ types.RpcClient = (*Client)(nil)

// Client is a Arbitrum client.
type Client struct {
	evmclient.Client

	fetcher *Fetcher
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config) (*Client, error) {
	client, err := evmclient.NewClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	fetcher, err := NewFetcher(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:  *client,
		fetcher: fetcher,
	}, nil
}

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64) {
	lastSyncedL1BlockNumber := c.fetcher.GetFetchedBlockNumber()
	if lastSyncedL1BlockNumber > 0 && lastSyncedL1BlockNumber+ParallelBlocks > l1BlockNumber {
		return
	}
	c.fetcher.Stop()
	logger.Infof("last synced L1 block number: %d, begin L1 block number: %d", lastSyncedL1BlockNumber, l1BlockNumber)

	c.fetcher.InitFetch()
	// Fetch L1 batch headers
	go func() {
		if err := c.fetcher.Fetch(l1BlockNumber); err != nil {
			logger.Errorf("failed to fetch L1 batch headers: %v", err)
			c.fetcher.Stop()
		}
	}()
}

// NextBatch returns the next batch header after SetBeginBlockNumber.
func (c *Client) NextBatch() (*sequencerv2types.BatchHeader, error) {
	return c.fetcher.nextBatchHeader()
}
