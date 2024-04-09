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

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64) {}

// NextBatch returns the next batch after SetBeginBlockNumber.
func (c *Client) NextBatch() (*sequencerv2types.BatchHeader, error) {
	return nil, types.ErrNoResult
}
