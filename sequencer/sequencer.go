package sequencer

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

const (
	// SyncInterval is the interval between two block syncs after fully synced.
	SyncInterval = 1 * time.Second
)

// Sequencer is the main component of the lagrange node.
// - It is responsible for fetching block headers from the blockchain.
type Sequencer struct {
	storage         storageInterface
	rpcClient       rpctypes.RpcClient
	chainID         uint32
	lastBlockNumber uint64

	ctx    context.Context
	cancel context.CancelFunc
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, rpcCfg *rpcclient.Config, storage storageInterface) (*Sequencer, error) {
	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
	if err != nil {
		return nil, err
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return nil, err
	}

	blockNumber, err := storage.GetLastBlockNumber(context.Background(), chainID)
	if err != nil {
		return nil, err
	}

	if cfg.FromBlockNumber > blockNumber {
		blockNumber = cfg.FromBlockNumber
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Sequencer{
		storage:         storage,
		rpcClient:       rpcClient,
		lastBlockNumber: blockNumber,
		chainID:         uint32(chainID),
		ctx:             ctx,
		cancel:          cancel,
	}, nil
}

// GetChainID returns the chain ID.
func (s *Sequencer) GetChainID() uint32 {
	return s.chainID
}

// Start starts the sequencer.
func (s *Sequencer) Start() error {
	finalizedBlockNumber, err := s.rpcClient.GetFinalizedBlockNumber()
	if err != nil {
		return err
	}

	logger.Infof("sequencer started from %d", finalizedBlockNumber)

	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			lastBlockNumber := s.lastBlockNumber
			if s.lastBlockNumber > finalizedBlockNumber {
				logger.Infof("sequencer synced to %d", finalizedBlockNumber)
				time.Sleep(SyncInterval)
				finalizedBlockNumber, err = s.rpcClient.GetFinalizedBlockNumber()
				if err != nil {
					return err
				}
				continue
			}
			blockHeader, err := s.rpcClient.GetBlockHeaderByNumber(lastBlockNumber, common.Hash{})
			if err != nil {
				if err == rpctypes.ErrBlockNotFound {
					logger.Infof("block %d not found", lastBlockNumber)
					time.Sleep(SyncInterval)
					continue
				}
				return err
			}

			if err := s.storage.AddBlock(s.ctx, &types.Block{
				ChainHeader: &types.ChainHeader{
					BlockNumber:   lastBlockNumber,
					BlockHash:     blockHeader.L2BlockHash.Hex(),
					ChainId:       s.chainID,
					L1BlockNumber: blockHeader.L1BlockNumber,
				},
				SequencedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
			}); err != nil {
				return err
			}

			s.lastBlockNumber = lastBlockNumber + 1
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// Stop stops the sequencer.
func (s *Sequencer) Stop() {
	if s != nil {
		s.cancel()
	}
}
