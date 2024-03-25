package sequencer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	v2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
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
	logger.Infof("Creating sequencer with config: %+v", cfg)

	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
	if err != nil {
		logger.Errorf("failed to create rpc client: %v", err)
		return nil, err
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		logger.Errorf("failed to get chain ID: %v", err)
		return nil, err
	}

	lastBlockNumber := uint64(0)
	batchNumber, err := storage.GetLastBatchNumber(context.Background(), chainID)
	if err != nil {
		if errors.Is(err, storetypes.ErrBatchNotFound) {
			logger.Infof("no batch found")
			rpcClient.SetBeginBlockNumber(cfg.FromL1BlockNumber)
		} else {
			logger.Errorf("failed to get last batch number: %v", err)
			return nil, err
		}
	} else {
		batch, err := storage.GetBatch(context.Background(), chainID, batchNumber)
		if err != nil {
			logger.Errorf("failed to get batch for batch number: %d error : %v", batchNumber, err)
			return nil, err
		}
		lastBlockNumber = batch.BatchHeader.ToBlockNumber()
		rpcClient.SetBeginBlockNumber(batch.L1BlockNumber())
	}
	if cfg.FromL2BlockNumber > lastBlockNumber {
		lastBlockNumber = cfg.FromL2BlockNumber - 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Sequencer{
		storage:         storage,
		rpcClient:       rpcClient,
		lastBlockNumber: lastBlockNumber,
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
	logger.Infof("Sequencer started from %d", s.lastBlockNumber+1)

	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			nextBlockNumber := s.lastBlockNumber + 1
			batchHeader, err := s.rpcClient.GetBatchHeaderByNumber(nextBlockNumber)
			if err != nil {
				if errors.Is(err, rpctypes.ErrBatchNotFound) {
					logger.Infof("block %d is not ready", nextBlockNumber)
					time.Sleep(SyncInterval)
					continue
				}
				logger.Errorf("failed to get batch header: %v", err)
				return err
			}

			if err := s.storage.AddBatch(context.Background(), &v2types.Batch{
				BatchHeader:   batchHeader,
				SequencedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
			}); err != nil {
				logger.Errorf("failed to add batch: %v", err)
				return err
			}

			s.lastBlockNumber = batchHeader.ToBlockNumber()
			logger.Infof("batch block sequenced up to %d", s.lastBlockNumber)
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
