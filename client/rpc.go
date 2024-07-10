package client

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/store/goleveldb"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/protobuf/proto"
)

const (
	committeeCacheSize = 10
	clientDBPath       = ".lagrange/db/"
	pruningBlocks      = 1000
)

type rpcAdapter struct {
	client         rpctypes.RpcClient
	committeeSC    *committee.Committee
	committeeCache *lru.Cache[uint64, *committee.ILagrangeCommitteeCommitteeData]
	db             storetypes.KVStorage

	isSetBeginBlockNumber atomic.Bool
	openL1BlockNumber     atomic.Uint64
	chainID               uint32
	genesisBlockNumber    uint64
}

// newRpcAdapter creates a new rpc adapter.
func newRpcAdapter(rpcCfg *rpcclient.Config, cfg *Config, pubkey []byte) (*rpcAdapter, error) {
	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create the rpc client: %v, please check the chain name, the chain name should look like 'optimism', 'base'", err)
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return nil, fmt.Errorf("failed to get the chain ID: %v", err)
	}

	etherClient, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create the ethereum client: %v", err)
	}
	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), etherClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create the committee smart contract: %v", err)
	}
	params, err := committeeSC.CommitteeParams(nil, chainID)
	if err != nil {
		logger.Fatalf("failed to get the committee params: %v", err)
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		logger.Fatalf("failed to get the home directory: %v", err)
	}
	dbPath := filepath.Clean(filepath.Join(homePath, clientDBPath))
	if err := os.MkdirAll(dbPath, os.ModePerm); err != nil {
		logger.Fatalf("failed to create the database directory: %v", err)
	}
	dbPath = filepath.Join(dbPath, fmt.Sprintf("client_%d_%x.db", chainID, pubkey))
	db, err := goleveldb.NewDB(dbPath)
	if err != nil {
		logger.Fatalf("failed to create the database: %v", err)
	}

	return &rpcAdapter{
		client:             rpcClient,
		committeeSC:        committeeSC,
		committeeCache:     lru.NewCache[uint64, *committee.ILagrangeCommitteeCommitteeData](committeeCacheSize),
		db:                 db,
		chainID:            chainID,
		genesisBlockNumber: uint64(params.GenesisBlock.Int64() - params.L1Bias.Int64()),
	}, nil
}

// startBatchFetching starts the batch fetching loop.
func (r *rpcAdapter) startBatchFetching(chErr chan<- error) {
	for {
		batch, err := r.client.NextBatch()
		if err != nil {
			logger.Errorf("failed to get the next batch: %v", err)
			chErr <- err
			return
		}
		telemetry.SetGauge(float64(batch.L1BlockNumber), "client", "fetch_batch_l1_block_number")
		logger.Infof("fetch the batch with L1 block number %d", batch.L1BlockNumber)

		// block the writeBatchHeader if the batch is too far from the current block
		for openBlockNumber := r.openL1BlockNumber.Load(); openBlockNumber > 0 && openBlockNumber+pruningBlocks/4 < batch.L1BlockNumber; openBlockNumber = r.openL1BlockNumber.Load() {
			if r.isSetBeginBlockNumber.Load() {
				break
			}
			time.Sleep(1 * time.Second)
		}
		openL1BlockNumber := r.openL1BlockNumber.Load()
		if openL1BlockNumber > 0 && openL1BlockNumber+pruningBlocks/4 < batch.L1BlockNumber {
			logger.Infof("Rolling back the batch fetching to the block number %d", r.openL1BlockNumber.Load())
		} else if openL1BlockNumber > 0 && openL1BlockNumber+pruningBlocks/2 < batch.L1BlockNumber {
			logger.Warnf("The batch %d fetching is too far from the current block number %d", batch.L1BlockNumber, openL1BlockNumber)
			continue
		} else if openL1BlockNumber > 0 && openL1BlockNumber+pruningBlocks/4 > batch.L1BlockNumber {
			r.isSetBeginBlockNumber.Store(false)
		}

		if err := r.writeBatchHeader(batch); err != nil {
			logger.Errorf("failed to write the batch header: %v", err)
			chErr <- err
			return
		}
		if openL1BlockNumber > pruningBlocks {
			prunedBlockNumber := openL1BlockNumber - pruningBlocks
			prefix := make([]byte, 12)
			binary.BigEndian.PutUint64(prefix, prunedBlockNumber)
			if err := r.db.Prune(prefix); err != nil {
				logger.Errorf("failed to prune the database: %v", err)
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (r *rpcAdapter) getCommitteeRoot(blockNumber uint64) (*committee.ILagrangeCommitteeCommitteeData, error) {
	if committeeData, ok := r.committeeCache.Get(blockNumber); ok {
		return committeeData, nil
	}

	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_committee")

	committeeData, err := r.committeeSC.GetCommittee(nil, r.chainID, big.NewInt(int64(blockNumber)))
	if err != nil || committeeData.LeafCount == 0 {
		return nil, fmt.Errorf("failed to get the committee data %+v: %v", committeeData, err)
	}
	r.committeeCache.Add(blockNumber, &committeeData)

	return &committeeData, nil
}

// initBeginBlockNumber initializes the begin block number for the RPC client.
func (r *rpcAdapter) initBeginBlockNumber(blockNumber uint64) error {
	// set the open block number
	r.openL1BlockNumber.Store(blockNumber)

	lastStoredBlockNumber := uint64(0)
	// get the last stored block number
	key := make([]byte, 12)
	binary.BigEndian.PutUint64(key, math.MaxUint64)
	binary.BigEndian.PutUint32(key[8:], math.MaxUint32)
	pKey, _, err := r.db.Prev(key)
	if err != nil {
		return fmt.Errorf("failed to get the previous key: %v", err)
	}
	if pKey != nil {
		lastStoredBlockNumber = binary.BigEndian.Uint64(pKey)
	}
	if lastStoredBlockNumber > blockNumber {
		// check if the block number exists in the database
		prefix := make([]byte, 12)
		storedBlockNumber := uint64(0)
		binary.BigEndian.PutUint64(prefix, blockNumber+1)
		pKey, _, err := r.db.Prev(prefix)
		if err != nil {
			return fmt.Errorf("failed to get the previous key: %v", err)
		}
		if pKey != nil {
			storedBlockNumber = binary.BigEndian.Uint64(pKey)
		}
		if storedBlockNumber == blockNumber {
			blockNumber = lastStoredBlockNumber
		}
	}

	r.isSetBeginBlockNumber.Store(r.client.SetBeginBlockNumber(blockNumber))

	return nil
}

// writeBatchHeader writes the batch header to the database.
func (r *rpcAdapter) writeBatchHeader(batchHeader *sequencerv2types.BatchHeader) error {
	key := make([]byte, 12)
	binary.BigEndian.PutUint64(key, batchHeader.L1BlockNumber)
	binary.BigEndian.PutUint32(key[8:], batchHeader.L1TxIndex)
	value, err := proto.Marshal(batchHeader)
	if err != nil {
		return fmt.Errorf("failed to marshal the batch header: %v", err)
	}

	return r.db.Put(key, value)
}

// getPrevBatchL1Number gets the previous batch L1 number from the database.
func (r *rpcAdapter) getPrevBatchL1Number(l1BlockNumber uint64, l1TxIndex uint32) (uint64, error) {
	key := make([]byte, 12)
	binary.BigEndian.PutUint64(key, l1BlockNumber)
	binary.BigEndian.PutUint32(key[8:], l1TxIndex)

	prevKey, _, err := r.db.Prev(key)
	if err != nil {
		return 0, fmt.Errorf("failed to get the previous key: %v", err)
	}
	var prevL1BlockNumber uint64
	if prevKey != nil {
		prevL1BlockNumber = binary.BigEndian.Uint64(prevKey[:8])
	}

	return prevL1BlockNumber, nil
}

// getBatchHeader gets the batch header from the database.
func (r *rpcAdapter) getBatchHeader(l1BlockNumber, l2BlockNumber uint64, l1TxIndex uint32) (*sequencerv2types.BatchHeader, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_batch_header")

	if l1TxIndex > 0 {
		key := make([]byte, 12)
		binary.BigEndian.PutUint64(key, l1BlockNumber)
		binary.BigEndian.PutUint32(key[8:], l1TxIndex)
		value, err := r.db.Get(key)
		if err != nil {
			return nil, err
		}
		var batchHeader sequencerv2types.BatchHeader
		if err := proto.Unmarshal(value, &batchHeader); err != nil {
			return nil, fmt.Errorf("failed to unmarshal the batch header: %v", err)
		}
		return &batchHeader, nil
	}

	var res *sequencerv2types.BatchHeader

	prefix := make([]byte, 8)
	binary.BigEndian.PutUint64(prefix, l1BlockNumber)

	errFound := fmt.Errorf("found")
	if err := r.db.Iterate(prefix, func(key, value []byte) error {
		var batchHeader sequencerv2types.BatchHeader
		if err := proto.Unmarshal(value, &batchHeader); err != nil {
			return fmt.Errorf("failed to unmarshal the batch header: %v", err)
		}
		if batchHeader.FromBlockNumber() == l2BlockNumber {
			res = &batchHeader
			return errFound
		}
		return nil
	}); errors.Is(errFound, err) {
		return res, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to iterate the database: %v", err)
	} else {
		return nil, fmt.Errorf("the batch header is not found for L1 block number %d, L2 block number %d", l1BlockNumber, l2BlockNumber)
	}
}

// setOpenL1BlockNumber sets the open L1 block number.
func (r *rpcAdapter) setOpenL1BlockNumber(blockNumber uint64) {
	r.openL1BlockNumber.Store(blockNumber)
}
