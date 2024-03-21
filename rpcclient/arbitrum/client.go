package arbitrum

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

var _ types.RpcClient = (*Client)(nil)

// L2Header is the L2 block header.
type L2Header struct {
	L1BlockNumber *hexutil.Big `json:"l1BlockNumber"`
}

// Client is a Arbitrum client.
type Client struct {
	evmclient.Client
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config) (*Client, error) {
	client, err := evmclient.NewClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: *client,
	}, nil
}

// SetBeginBlockNumber sets the begin L1 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber uint64) {}

// GetBatchHeaderByNumber returns the batch header for the given L2 block number.
func (c *Client) GetBatchHeaderByNumber(l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	return nil, nil
}
