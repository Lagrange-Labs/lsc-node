package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

const batchBlockCount = 5

var _ types.RpcClient = (*Client)(nil)

// Client is a Arbitrum client.
type Client struct {
	evmclient.Client

	isLight           bool
	mtx               sync.Mutex
	fromL1BlockNumber uint64
	chainID           uint32
}

// NewClient creates a new Client instance.
func NewClient(cfg *Config, isLight bool) (*Client, error) {
	client, err := evmclient.NewClient(cfg.RPCURLs)
	if err != nil {
		return nil, err
	}
	chainID, err := client.GetChainID()
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:  *client,
		isLight: isLight,
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
func (c *Client) SetBeginBlockNumber(l1BlockNumber, _ uint64) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.fromL1BlockNumber = l1BlockNumber
	return true
}

// NextBatch returns the next batch after SetBeginBlockNumber.
func (c *Client) NextBatch() (*sequencerv2types.BatchHeader, error) {
	c.mtx.Lock()
	l1BlockNumber := c.fromL1BlockNumber
	c.mtx.Unlock()
	l2BlockNumber := l1BlockNumber + batchBlockCount - 1
	_, err := c.GetBlockHeaderByNumber(l2BlockNumber, common.Hash{})
	if err != nil {
		if errors.Is(err, types.ErrNoResult) {
			// wait for the block to be available
			for {
				_, err = c.GetBlockHeaderByNumber(l2BlockNumber, common.Hash{})
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

	l2Blocks := make([]*sequencerv2types.BlockHeader, 0, batchBlockCount)
	if c.isLight {
		blockHeader, err := c.GetBlockHeaderByNumber(l2BlockNumber-batchBlockCount+1, common.Hash{})
		if err != nil {
			return nil, fmt.Errorf("failed to get the block header for block number: %d  error: %w", l2BlockNumber-batchBlockCount, err)
		}
		l2Blocks = append(l2Blocks, &sequencerv2types.BlockHeader{
			BlockNumber: l2BlockNumber - batchBlockCount + 1,
			BlockHash:   blockHeader.Hash().Hex(),
		})
	} else {
		for i := uint64(0); i < batchBlockCount; i++ {
			blockHeader, err := c.GetBlockHeaderByNumber(l2BlockNumber-batchBlockCount+i+1, common.Hash{})
			if err != nil {
				return nil, fmt.Errorf("failed to get the block header for block number: %d  error: %w", l2BlockNumber-batchBlockCount+i, err)
			}
			var buffer bytes.Buffer
			if err := rlp.Encode(&buffer, blockHeader); err != nil {
				return nil, fmt.Errorf("failed to encode the block header: %v", err)
			}
			l2Blocks = append(l2Blocks, &sequencerv2types.BlockHeader{
				BlockNumber: l2BlockNumber - batchBlockCount + i + 1,
				BlockHash:   blockHeader.Hash().Hex(),
				BlockRlp:    core.Bytes2Hex(buffer.Bytes()),
			})
		}
	}

	c.mtx.Lock()
	c.fromL1BlockNumber += batchBlockCount
	c.mtx.Unlock()

	return &sequencerv2types.BatchHeader{
		BatchNumber:       l1BlockNumber,
		ChainId:           c.chainID,
		L2Blocks:          l2Blocks,
		L1BlockNumber:     l1BlockNumber,
		L2FromBlockNumber: l2BlockNumber - batchBlockCount + 1,
		L2ToBlockNumber:   l2BlockNumber,
		L1TxHash:          l2Blocks[0].BlockHash,
	}, nil
}

// GetL2BatchHeader returns the L2 batch header by the given L1 block number and transaction hash.
func (c *Client) GetL2BatchHeader(l1BlockNumber uint64, txHash string) (*sequencerv2types.BatchHeader, error) {
	c.SetBeginBlockNumber(l1BlockNumber, 0)
	return c.NextBatch()
}

// VerifyBatchHeader verifies the batch header with the given L1 block number and L2 block number.
func (c *Client) VerifyBatchHeader(l1BlockNumber, l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	c.SetBeginBlockNumber(l1BlockNumber, l2BlockNumber)
	return c.NextBatch()
}
