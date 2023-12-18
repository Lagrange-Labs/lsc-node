package arbitrum

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

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

	ethClient    *ethclient.Client
	batchStorage common.Address // Address of the L1BatchStorage contract
}

// NewClient creates a new Client instance.
func NewClient(rpcURL, l1RpcURL string, batchStorageAddr string) (*Client, error) {
	client, err := evmclient.NewClient(rpcURL)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(l1RpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:       *client,
		ethClient:    ethClient,
		batchStorage: common.HexToAddress(batchStorageAddr),
	}, nil
}

// GetL1BlockNumber returns the current L1 block number for the given L2 block number.
func (c *Client) GetL1BlockNumber(l2BlockNumber uint64) (uint64, error) {
	rawHeader, err := c.GetRawHeaderByNumber(l2BlockNumber)
	if err != nil {
		return 0, fmt.Errorf("failed to get L2 block header: %w", err)
	}

	var header L2Header
	if err := json.Unmarshal(rawHeader, &header); err != nil {
		return 0, fmt.Errorf("failed to unmarshal L2 block header: %w rawHeader: %s", err, rawHeader)
	}

	if header.L1BlockNumber == nil {
		return 0, fmt.Errorf("L1 block number is nil")
	}

	return header.L1BlockNumber.ToInt().Uint64(), nil
}
