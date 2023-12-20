package arbitrum

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
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

// GetBlockHeaderByNumber returns the L2 block header for the given L2 block number.
func (c *Client) GetBlockHeaderByNumber(l2BlockNumber uint64, l1TxHash common.Hash) (*types.L2BlockHeader, error) {
	rawHeader, err := c.GetRawHeaderByNumber(l2BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block header: %w", err)
	}

	var header L2Header
	if err := json.Unmarshal(rawHeader, &header); err != nil {
		return nil, fmt.Errorf("failed to unmarshal L2 block header: %w rawHeader: %s", err, rawHeader)
	}

	var commonHeader ethtypes.Header
	if err := json.Unmarshal(rawHeader, &commonHeader); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Eth block header: %w rawHeader: %s", err, rawHeader)
	}
	if header.L1BlockNumber == nil {
		return nil, fmt.Errorf("L1 block number is nil")
	}

	return &types.L2BlockHeader{
		L1BlockNumber: header.L1BlockNumber.ToInt().Uint64(),
		L2BlockHash:   commonHeader.Hash(),
	}, nil
}
