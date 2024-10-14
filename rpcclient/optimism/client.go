package optimism

import (
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

var _ types.RpcClient = (*Client)(nil)

// Client is a Optimism client.
type Client struct {
	evmclient.Client

	l1ParallelBlocks int
	fetcher          *Fetcher
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config, isLight bool) (*Client, error) {
	client, err := evmclient.NewClient(cfg.RPCURLs)
	if err != nil {
		return nil, err
	}

	fetcher, err := NewFetcher(cfg, isLight)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:           *client,
		l1ParallelBlocks: cfg.L1ParallelBlocks,
		fetcher:          fetcher,
	}, nil
}

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64) bool {
	lastSyncedL1BlockNumber := c.fetcher.GetFetchedBlockNumber()
	lastPulledL1BlockNumber := c.fetcher.GetPulledBlockNumber()
	if lastSyncedL1BlockNumber+uint64(c.l1ParallelBlocks) > l1BlockNumber && l1BlockNumber >= lastPulledL1BlockNumber {
		// if the begin block number is already synced, return false
		return false
	}

	c.fetcher.StopFetch()
	logger.Infof("last synced L1 block number: %d, begin L1 block number: %d, begin L2 block number: %d", lastSyncedL1BlockNumber, l1BlockNumber, l2BlockNumber)

	c.fetcher.InitFetch()
	// Fetch L1 batch headers
	go func() {
		if err := c.fetcher.Fetch(l1BlockNumber, l2BlockNumber); err != nil {
			logger.Errorf("failed to fetch L1 batch headers: %v", err)
			c.fetcher.Stop()
		}
	}()

	return true
}

// NextBatch returns the next batch header after SetBeginBlockNumber.
func (c *Client) NextBatch() (*sequencerv2types.BatchHeader, error) {
	return c.fetcher.nextBatchHeader()
}

// GetL2BatchHeader returns the L2 batch header by the given L1 block number and transaction hash.
func (c *Client) GetL2BatchHeader(l1BlockNumber uint64, txHash string) (*sequencerv2types.BatchHeader, error) {
	return c.fetcher.GetL2BatchHeader(l1BlockNumber, 0, txHash)
}

// VerifyBatchHeader verifies the batch header with the given L1 block number and L2 block number.
func (c *Client) VerifyBatchHeader(l1BlockNumber, l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	return c.fetcher.GetL2BatchHeader(l1BlockNumber, l2BlockNumber, "")
}
