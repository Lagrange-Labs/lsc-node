package optimism

import (
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/ethereum/go-ethereum/common"
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

// GetBlockHeaderByNumber returns the L2 block header for the given L2 block number.
func (c *Client) GetBlockHeaderByNumber(l2BlockNumber uint64, l1TxHash common.Hash) (*types.L2BlockHeader, error) {
	return nil, nil
}

// GetBatchHeaderByNumber returns the batch header for the given L2 block number.
func (c *Client) GetBatchHeaderByNumber(l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	return c.fetcher.getL2BatchData(l2BlockNumber)
}

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64) {
	go func() {
		if err := c.fetcher.Fetch(l1BlockNumber); err != nil {
			logger.Errorf("failed to fetch L2 block headers: %v", err)
		}
	}()
}
