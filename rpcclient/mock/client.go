package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

var _ types.RpcClient = (*Client)(nil)

// Client is a Arbitrum client.
type Client struct {
	evmclient.Client

	fromL1BlockNumber uint64

	chainID uint32
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config) (*Client, error) {
	client, err := evmclient.NewClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}
	chainID, err := client.GetChainID()
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:  *client,
		chainID: chainID,
	}, nil
}

// GetBlockHeaderByNumber returns the block header for the given L2 block number.
func (c *Client) GetBlockHeaderByNumber(l2BlockNumber uint64, l1TxHash common.Hash) (*ethtypes.Header, error) {
	rawHeader, err := c.GetRawHeaderByNumber(l2BlockNumber)
	if err != nil {
		return nil, err
	}

	var commonHeader ethtypes.Header
	if err := json.Unmarshal(rawHeader, &commonHeader); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Eth block header: %w rawHeader: %s", err, rawHeader)
	}

	return &commonHeader, nil
}

// GetFinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetFinalizedBlockNumber() (uint64, error) {
	return c.GetCurrentBlockNumber()
}

// SetBeginBlockNumber sets the begin L1 & L2 block number.
func (c *Client) SetBeginBlockNumber(l1BlockNumber uint64) {
	c.fromL1BlockNumber = l1BlockNumber
}

// NextBatch returns the next batch after SetBeginBlockNumber.
func (c *Client) NextBatch() (*sequencerv2types.BatchHeader, error) {
	l2BlockNumber := c.fromL1BlockNumber
	blockHeader, err := c.GetBlockHeaderByNumber(l2BlockNumber, common.Hash{})
	if err != nil {
		if errors.Is(err, types.ErrNoResult) {
			// wait for the block to be available
			for {
				blockHeader, err = c.GetBlockHeaderByNumber(l2BlockNumber, common.Hash{})
				if errors.Is(err, types.ErrNoResult) {
					time.Sleep(1 * time.Second)
					continue
				}
				if err != nil {
					return nil, fmt.Errorf("failed to get the block header for block number: %d  error: %w", l2BlockNumber, err)
				}
				break
			}
		} else {
			return nil, fmt.Errorf("failed to get the block header for block number: %d  error: %w", l2BlockNumber, err)
		}
	}

	c.fromL1BlockNumber++
	return &sequencerv2types.BatchHeader{
		BatchNumber: blockHeader.Number.Uint64(),
		ChainId:     c.chainID,
		L2Blocks: []*sequencerv2types.BlockHeader{
			{
				BlockNumber: l2BlockNumber,
				BlockHash:   blockHeader.Hash().Hex(),
			},
		},
		L1BlockNumber: blockHeader.Number.Uint64(),
		L1TxHash:      blockHeader.Hash().Hex(),
	}, nil
}
