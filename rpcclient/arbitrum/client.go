package arbitrum

import (
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ types.RpcClient = (*Client)(nil)

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
	// TODO: This is a temporary workaround for testing.
	return l2BlockNumber, nil
}
