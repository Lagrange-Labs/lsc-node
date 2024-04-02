package optimism

import (
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ types.RpcClient = (*Client)(nil)

// Client is a Optimism client.
type Client struct {
	evmclient.Client

	ethClient *ethclient.Client
	fetcher   *Fetcher
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config) (*Client, error) {
	logger.Infof("creating rpc client for confg: %+v", cfg)

	client, err := evmclient.NewClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}

	fetcher, err := NewFetcher(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:    *client,
		ethClient: ethClient,
		fetcher:   fetcher,
	}, nil
}

// GetBatchHeaderByNumber returns the batch header for the given L2 block number.
func (c *Client) GetBatchHeaderByNumber(l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	return c.fetcher.getL2BatchData(l2BlockNumber)
}

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64) {
	lastSyncedL1BlockNumber := c.fetcher.GetFetchedBlockNumber()
	if lastSyncedL1BlockNumber+ParallelBlocks > l1BlockNumber {
		return
	}
	logger.Infof("last synced L1 block number: %d, begin L1 block number: %d", lastSyncedL1BlockNumber, l1BlockNumber)
	c.fetcher.Stop()

	c.fetcher.InitFetch()
	// Fetch L1 batch headers
	go func() {
		if err := c.fetcher.Fetch(l1BlockNumber); err != nil {
			logger.Errorf("failed to fetch L1 batch headers: %v", err)
		}
	}()
	// Fetch L2 block headers
	go func() {
		if err := c.fetcher.FetchL2Blocks(l2BlockNumber); err != nil {
			logger.Errorf("failed to fetch L2 block headers: %v", err)
		}
	}()
}
