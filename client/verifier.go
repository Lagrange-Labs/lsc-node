package client

import (
	"bytes"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/lsc-node/core"
	"github.com/Lagrange-Labs/lsc-node/core/crypto"
	"github.com/Lagrange-Labs/lsc-node/core/logger"
	"github.com/Lagrange-Labs/lsc-node/core/telemetry"
	"github.com/Lagrange-Labs/lsc-node/scinterface/committee"
	sequencerv2types "github.com/Lagrange-Labs/lsc-node/sequencer/types/v2"
)

const (
	committeeCacheSize = 10
	numWorkers         = 4
)

// AdapterCaller is the interface to get the batch header from the rpc client.
type AdapterCaller interface {
	GetBatchHeader(l1BlockNumber uint64, txHash string, l1TxIndex uint32) (*sequencerv2types.BatchHeader, error)
	GetPrevBatchL1Number(l1BlockNumber uint64, l1TxIndex uint32) (uint64, error)
	VerifyBatchHeader(l1BlockNumber, l2BlockNumber uint64) error
	GetBlockHash(rlpHeader []byte) (common.Hash, common.Hash, error)
}

// CommitteeBackend is the interface to get the committee data from the smart contract.
type CommitteeBackend interface {
	GetCommittee(opts *bind.CallOpts, chainID uint32, blockNumber *big.Int) (committee.ILagrangeCommitteeCommitteeData, error)
}

var _ VerifierCaller = (*Verifier)(nil)

// Verifier is the struct to verify the batch from the sequencer.
type Verifier struct {
	adapter   AdapterCaller
	blsScheme crypto.BLSScheme

	committeeSC        CommitteeBackend
	committeeCache     *lru.Cache[uint64, *committee.ILagrangeCommitteeCommitteeData]
	genesisBlockNumber uint64
	chainID            uint32
}

// newVerifier creates a new verifier.
func newVerifier(cfg *Config, adapter AdapterCaller, chainID uint32) (*Verifier, error) {
	etherClient, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create the ethereum client: %v", err)
	}
	committeeSC, err := committee.NewCommitteeCaller(common.HexToAddress(cfg.CommitteeSCAddress), etherClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create the committee smart contract: %v", err)
	}
	params, err := committeeSC.CommitteeParams(nil, chainID)
	if err != nil {
		logger.Fatalf("failed to get the committee params: %v", err)
	}
	logger.Infof("committee params: %+v", params)
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))

	return &Verifier{
		adapter:            adapter,
		committeeSC:        committeeSC,
		blsScheme:          blsScheme,
		committeeCache:     lru.NewCache[uint64, *committee.ILagrangeCommitteeCommitteeData](committeeCacheSize),
		chainID:            chainID,
		genesisBlockNumber: uint64(params.GenesisBlock.Int64() - params.L1Bias.Int64()),
	}, nil
}

func (v *Verifier) getCommitteeRoot(blockNumber uint64) (*committee.ILagrangeCommitteeCommitteeData, error) {
	if committeeData, ok := v.committeeCache.Get(blockNumber); ok {
		return committeeData, nil
	}

	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_committee_root")

	committeeData, err := v.committeeSC.GetCommittee(nil, v.chainID, big.NewInt(int64(blockNumber)))
	if err != nil || committeeData.LeafCount == 0 {
		return nil, fmt.Errorf("failed to get the committee data for block number %d: %v", blockNumber, err)
	}
	v.committeeCache.Add(blockNumber, &committeeData)

	return &committeeData, nil
}

// VerifyPrevBatch verifies the previous batch.
func (v *Verifier) VerifyPrevBatch(l1BlockNumber, l2BlockNumber uint64) error {
	return v.adapter.VerifyBatchHeader(l1BlockNumber, l2BlockNumber)
}

// VerifyBatch verifies the proposed batch.
func (v *Verifier) VerifyBatch(batch *sequencerv2types.Batch) error {
	// verify the batch header
	if err := v.verifyBatchHeader(batch); err != nil {
		return fmt.Errorf("failed to verify the batch header: %v", err)
	}

	// verify the committee root
	if err := v.verifyCommitteeRoot(batch); err != nil {
		return fmt.Errorf("failed to verify the committee root: %v", err)
	}

	return nil
}

// verifyBatchHeader verifies the batch header with the source chain one.
func (v *Verifier) verifyBatchHeader(batch *sequencerv2types.Batch) error {
	l1BlockNumber := batch.L1BlockNumber()
	batchHeader, err := v.adapter.GetBatchHeader(l1BlockNumber, batch.BatchHeader.L1TxHash, batch.BatchHeader.L1TxIndex)
	if err != nil || batchHeader == nil {
		logger.Errorf("failed to get the batch header for L1BlockNumber %d, L2FromBlockNumber %d, L1TxIndex %d: %v", l1BlockNumber, batch.BatchHeader.FromBlockNumber(), batch.BatchHeader.L1TxIndex, err)
		return ErrBatchNotFound
	}
	// verify the L2 blocks
	if err := v.verifyL2Blocks(batch, batchHeader); err != nil {
		return fmt.Errorf("failed to verify the L2 blocks: %v", err)
	}
	if l1BlockNumber != batchHeader.L1BlockNumber {
		return fmt.Errorf("the batch L1 block number %d is not equal to the rpc L1 block number %d", batch.L1BlockNumber(), batchHeader.L1BlockNumber)
	}

	// verify the sequencer signature
	if len(batch.ProposerPubKey) == 0 {
		return fmt.Errorf("the block %d proposer key is empty", batch.BatchNumber())
	}
	blsSigHash := batch.BlsSignature().Hash()
	verified, err := v.blsScheme.VerifySignature(core.Hex2Bytes(batch.ProposerPubKey), blsSigHash, core.Hex2Bytes(batch.ProposerSignature), true)
	if err != nil || !verified {
		return fmt.Errorf("failed to verify the proposer signature: %v", err)
	}

	return nil
}

// verifyL2Blocks verifies the L2 blocks through the recursive way.
func (v *Verifier) verifyL2Blocks(batch *sequencerv2types.Batch, lightBatchHeader *sequencerv2types.BatchHeader) error {
	// verify the L1 tx hash and L1 block number
	if !bytes.Equal(core.Hex2Bytes(batch.BatchHeader.L1TxHash), core.Hex2Bytes(lightBatchHeader.L1TxHash)) {
		return fmt.Errorf("the light batch L1 tx hash %s is not equal to the batch L1 tx hash %s", lightBatchHeader.L1TxHash, batch.BatchHeader.L1TxHash)
	}
	if batch.BatchHeader.L1BlockNumber != lightBatchHeader.L1BlockNumber {
		return fmt.Errorf("the light batch L1 block number %d is not equal to the batch L1 block number %d", lightBatchHeader.L1BlockNumber, batch.BatchHeader.L1BlockNumber)
	}
	// verify the from and to block number
	if lightBatchHeader.FromBlockNumber() != batch.BatchHeader.FromBlockNumber() {
		return fmt.Errorf("the light batch from block number %d is not equal to the batch from block number %d", lightBatchHeader.FromBlockNumber(), batch.BatchHeader.FromBlockNumber())
	}
	if lightBatchHeader.ToBlockNumber() != batch.BatchHeader.ToBlockNumber() {
		return fmt.Errorf("the light batch to block number %d is not equal to the batch to block number %d", lightBatchHeader.ToBlockNumber(), batch.BatchHeader.ToBlockNumber())
	}
	// compare the last L2 block hash
	if !bytes.Equal(core.Hex2Bytes(lightBatchHeader.L2Blocks[0].BlockHash), core.Hex2Bytes(batch.BatchHeader.L2Blocks[len(batch.BatchHeader.L2Blocks)-1].BlockHash)) {
		return fmt.Errorf("the light batch last block hash %s is not equal to the batch last block hash %s", lightBatchHeader.L2Blocks[0].BlockHash, batch.BatchHeader.L2Blocks[0].BlockHash)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, numWorkers)

	// verify the subsequent blocks recursively
	verifyBlocks := func(start, end int) {
		defer wg.Done()
		for i := start; i <= end; i++ {
			curHash, parentHash, err := v.adapter.GetBlockHash(core.Hex2Bytes(batch.BatchHeader.L2Blocks[i].BlockRlp))
			if err != nil {
				errCh <- fmt.Errorf("failed to decode the block header: %v", err)
			}
			if !bytes.Equal(curHash[:], core.Hex2Bytes(batch.BatchHeader.L2Blocks[i].BlockHash)) {
				errCh <- fmt.Errorf("the current hash %s is not equal to the block hash %s", core.Bytes2Hex(curHash[:]), batch.BatchHeader.L2Blocks[i].BlockHash)
			}
			if i > start {
				if !bytes.Equal(parentHash[:], core.Hex2Bytes(batch.BatchHeader.L2Blocks[i-1].BlockHash)) {
					errCh <- fmt.Errorf("the parent hash %s is not equal to the previous block hash %s", core.Bytes2Hex(parentHash[:]), batch.BatchHeader.L2Blocks[i-1].BlockHash)
				}
			}
		}
		errCh <- nil
	}

	blockRange := int(lightBatchHeader.ToBlockNumber() - lightBatchHeader.FromBlockNumber())
	batchSize := (blockRange + numWorkers - 1) / numWorkers // ceil

	for i := 0; i < numWorkers; i++ {
		start := i * batchSize
		end := min(start+batchSize, blockRange)
		wg.Add(1)
		go verifyBlocks(start, end)
	}

	// wait for all the workers to finish
	wg.Wait()
	close(errCh)

	// check if there is any error
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

// verifyCommitteeRoot verifies the committee root.
func (v *Verifier) verifyCommitteeRoot(batch *sequencerv2types.Batch) error {
	blockNumber := batch.L1BlockNumber()
	// verify the previous batch's next committee root
	if v.genesisBlockNumber == blockNumber { // the genesis block
		if batch.CurrentCommittee() != batch.NextCommittee() {
			return fmt.Errorf("the genesis block current committee root %s is not equal to the next committee root %s", batch.CurrentCommittee(), batch.NextCommittee())
		}
	} else {
		var err error
		prevBatchL1Number, err := v.adapter.GetPrevBatchL1Number(batch.L1BlockNumber(), batch.BatchHeader.L1TxIndex)
		if err != nil {
			return fmt.Errorf("failed to get the previous batch L1 number: %v", err)
		}
		if prevBatchL1Number == 0 {
			return ErrBatchNotFound
		}

		prevCommitteeData, err := v.getCommitteeRoot(prevBatchL1Number)
		if err != nil {
			return fmt.Errorf("failed to get the previous committee root: %v", err)
		}
		if !bytes.Equal(core.Hex2Bytes(batch.CurrentCommittee()), prevCommitteeData.Root[:]) {
			return fmt.Errorf("the current batch committee root %s is not equal to the previous batch next committee root %s", batch.CurrentCommittee(), core.Bytes2Hex(prevCommitteeData.Root[:]))
		}
	}

	// verify the current batch's next committee root
	curCommitteeData, err := v.getCommitteeRoot(blockNumber)
	if err != nil {
		return fmt.Errorf("failed to get the current committee root: %v", err)
	}
	if !bytes.Equal(core.Hex2Bytes(batch.NextCommittee()), curCommitteeData.Root[:]) {
		return fmt.Errorf("the current batch next committee root %s is not equal to the current committee root %s", batch.NextCommittee(), core.Bytes2Hex(curCommitteeData.Root[:]))
	}

	return nil
}
