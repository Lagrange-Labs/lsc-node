package client

import (
	"bytes"
	"fmt"
	"math/big"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	committeeCacheSize = 10
)

type Adapter interface {
	GetBatchHeader(l1BlockNumber, l2BlockNumber uint64, l1TxIndex uint32) (*sequencerv2types.BatchHeader, error)
	GetPrevBatchL1Number(l1BlockNumber uint64, l1TxIndex uint32) (uint64, error)
}

// Verifier is the struct to verify the batch from the sequencer.
type Verifier struct {
	adapter   Adapter
	blsScheme crypto.BLSScheme

	committeeSC        *committee.Committee
	committeeCache     *lru.Cache[uint64, *committee.ILagrangeCommitteeCommitteeData]
	genesisBlockNumber uint64
	chainID            uint32
}

// newVerifier creates a new verifier.
func newVerifier(cfg *Config, adapter Adapter, chainID uint32) (*Verifier, error) {
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
	defer telemetry.MeasureSince(ti, "client", "get_committee")

	committeeData, err := v.committeeSC.GetCommittee(nil, v.chainID, big.NewInt(int64(blockNumber)))
	if err != nil || committeeData.LeafCount == 0 {
		return nil, fmt.Errorf("failed to get the committee data for block number %d: %v", blockNumber, err)
	}
	v.committeeCache.Add(blockNumber, &committeeData)

	return &committeeData, nil
}

// VerifyPrevBatch verifies the previous batch.
func (v *Verifier) VerifyPrevBatch(l1BlockNumber, l2BlockNumber uint64) error {
	batchHeader, err := v.adapter.GetBatchHeader(l1BlockNumber, l2BlockNumber, 0)
	if err != nil {
		return fmt.Errorf("failed to get the previous batch header for L1 block number %d, L2 block number %d: %v", l1BlockNumber, l2BlockNumber, err)
	}

	if batchHeader == nil {
		return fmt.Errorf("the batch header is not found for L1 block number %d, L2 block number %d", l1BlockNumber, l2BlockNumber)
	}

	return nil
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
	// verify if the batch hash is correct
	l1BlockNumber := batch.L1BlockNumber()
	batchHeader, err := v.adapter.GetBatchHeader(l1BlockNumber, batch.BatchHeader.FromBlockNumber(), batch.BatchHeader.L1TxIndex)
	if err != nil || batchHeader == nil {
		logger.Errorf("failed to get the batch header for L1BlockNumber %d, L2FromBlockNumber %d, L1TxIndex %d: %v", l1BlockNumber, batch.BatchHeader.FromBlockNumber(), batch.BatchHeader.L1TxIndex, err)
		return ErrBatchNotFound
	}
	if l1BlockNumber != batchHeader.L1BlockNumber {
		return fmt.Errorf("the batch L1 block number %d is not equal to the rpc L1 block number %d", batch.L1BlockNumber(), batchHeader.L1BlockNumber)
	}
	batchHash := batch.BatchHeader.Hash()
	bhHash := batchHeader.Hash()
	if !bytes.Equal(batchHash, bhHash) {
		return fmt.Errorf("the batch hash %s is not equal to the batch header hash %s", batchHash, utils.Bytes2Hex(bhHash))
	}

	// verify the sequencer signature
	if len(batch.ProposerPubKey) == 0 {
		return fmt.Errorf("the block %d proposer key is empty", batch.BatchNumber())
	}
	blsSigHash := batch.BlsSignature().Hash()
	verified, err := v.blsScheme.VerifySignature(common.FromHex(batch.ProposerPubKey), blsSigHash, common.FromHex(batch.ProposerSignature))
	if err != nil || !verified {
		return fmt.Errorf("failed to verify the proposer signature: %v", err)
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
		if !bytes.Equal(utils.Hex2Bytes(batch.CurrentCommittee()), prevCommitteeData.Root[:]) {
			return fmt.Errorf("the current batch committee root %s is not equal to the previous batch next committee root %s", batch.CurrentCommittee(), utils.Bytes2Hex(prevCommitteeData.Root[:]))
		}
	}

	// verify the current batch's next committee root
	curCommitteeData, err := v.getCommitteeRoot(blockNumber)
	if err != nil {
		return fmt.Errorf("failed to get the current committee root: %v", err)
	}
	if !bytes.Equal(utils.Hex2Bytes(batch.NextCommittee()), curCommitteeData.Root[:]) {
		return fmt.Errorf("the current batch next committee root %s is not equal to the current committee root %s", batch.NextCommittee(), utils.Bytes2Hex(curCommitteeData.Root[:]))
	}

	return nil
}
