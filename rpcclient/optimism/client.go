package optimism

import (
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
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
	header, err := c.fetcher.getL2BlockHeader(l2BlockNumber)
	if err != nil {
		if err == types.ErrBlockNotFound {
			// from the sequencer
			if l1TxHash == (common.Hash{}) {
				return nil, types.ErrBlockNotFound
			}
			// from the client
			return c.fetcher.getL2BlockHeaderByTxHash(l2BlockNumber, l1TxHash)
		}
		return nil, err
	}

	return header, nil
}
