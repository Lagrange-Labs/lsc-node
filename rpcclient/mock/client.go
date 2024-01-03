package mock

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
)

var _ types.RpcClient = (*Client)(nil)

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

// GetBlockHeaderByNumber returns the L2 block header for the given L2 block number.
func (c *Client) GetBlockHeaderByNumber(l2BlockNumber uint64, l1TxHash common.Hash) (*types.L2BlockHeader, error) {
	rawHeader, err := c.GetRawHeaderByNumber(l2BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block header: %w", err)
	}

	var commonHeader ethtypes.Header
	if err := json.Unmarshal(rawHeader, &commonHeader); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Eth block header: %w rawHeader: %s", err, rawHeader)
	}

	return &types.L2BlockHeader{
		L1BlockNumber: l2BlockNumber,
		L2BlockHash:   commonHeader.Hash(),
	}, nil
}

// GetFinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetFinalizedBlockNumber() (uint64, error) {
	return c.GetCurrentBlockNumber()
}
