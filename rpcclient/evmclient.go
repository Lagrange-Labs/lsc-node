package rpcclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmClient struct {
	ethClient *ethclient.Client
}

var _ RpcClient = (*EvmClient)(nil)

// CreateRPCClient creates a new rpc client.
func CreateRPCClient(chain, rpcURL string) (RpcClient, error) {
	switch chain {
	case "arbitrum":
		return NewEvmClient(rpcURL)
	case "optimism":
		return NewEvmClient(rpcURL)
	default:
		return nil, nil
	}
}

// NewEvmClient creates a new EvmClient instance.
func NewEvmClient(rpcURL string) (*EvmClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &EvmClient{
		ethClient: client,
	}, nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *EvmClient) GetBlockHashByNumber(blockNumber uint64) (string, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err == ethereum.NotFound {
		return "", ErrBlockNotFound
	}

	return header.Hash().Hex(), err
}

// GetChainID returns the chain ID.
func (c *EvmClient) GetChainID() (int32, error) {
	chainID, err := c.ethClient.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	return int32(chainID.Int64()), err
}
