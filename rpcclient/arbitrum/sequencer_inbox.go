package arbitrum

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/arbinbox"
)

var sequencerBridgeABI *abi.ABI
var batchDeliveredID common.Hash
var addSequencerL2BatchFromOriginCallABI abi.Method
var sequencerBatchDataABI abi.Event

const sequencerBatchDataEvent = "SequencerBatchData"

type batchDataLocation uint8

const (
	batchDataTxInput batchDataLocation = iota
	batchDataSeparateEvent
	batchDataNone
	batchDataBlobHashes
)

func init() {
	var err error
	sequencerBridgeABI, err = arbinbox.ArbinboxMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	batchDeliveredID = sequencerBridgeABI.Events["SequencerBatchDelivered"].ID
	sequencerBatchDataABI = sequencerBridgeABI.Events[sequencerBatchDataEvent]
	addSequencerL2BatchFromOriginCallABI = sequencerBridgeABI.Methods["addSequencerL2BatchFromOrigin0"]
}

// SequencerInbox is the struct to fetch the batch transactions from the sequencer inbox.
type SequencerInbox struct {
	inbox   *arbinbox.Arbinbox
	address common.Address
	client  *ethclient.Client

	chainId *big.Int
}

// SequencerBatch is the struct to represent the sequencer batch.
type SequencerBatch struct {
	BlockHash      common.Hash
	BlockNumber    uint64
	TxIndex        uint
	TxHash         common.Hash
	SequenceNumber uint64

	rawLog       types.Log
	dataLocation batchDataLocation
	serialized   []byte
	segments     [][]byte
	txes         []*types.Transaction
}

// NewSequencerInbox creates a new SequencerInbox instance.
func NewSequencerInbox(inboxAddr common.Address, client *ethclient.Client) (*SequencerInbox, error) {
	inbox, err := arbinbox.NewArbinbox(inboxAddr, client)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &SequencerInbox{
		inbox:   inbox,
		address: inboxAddr,
		client:  client,
		chainId: chainId,
	}, nil
}

// fetchBatchTransactions fetches the batch transactions from the sequencer inbox.
func (s *SequencerInbox) fetchBatchTransactions(ctx context.Context, from, to *big.Int) ([]*SequencerBatch, error) {
	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{s.address},
		Topics:    [][]common.Hash{{batchDeliveredID}},
	}
	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	messages := make([]*SequencerBatch, 0, len(logs))
	var lastSeqNum *uint64
	for _, log := range logs {
		if log.Topics[0] != batchDeliveredID {
			return nil, errors.New("unexpected log selector")
		}
		parsedLog, err := s.inbox.ParseSequencerBatchDelivered(log)
		if err != nil {
			return nil, err
		}
		if !parsedLog.BatchSequenceNumber.IsUint64() {
			return nil, errors.New("sequencer inbox event has non-uint64 sequence number")
		}
		if !parsedLog.AfterDelayedMessagesRead.IsUint64() {
			return nil, errors.New("sequencer inbox event has non-uint64 delayed messages read")
		}

		seqNum := parsedLog.BatchSequenceNumber.Uint64()
		if lastSeqNum != nil {
			if seqNum != *lastSeqNum+1 {
				return nil, fmt.Errorf("sequencer batches out of order; after batch %v got batch %v", lastSeqNum, seqNum)
			}
		}
		lastSeqNum = &seqNum
		batch := &SequencerBatch{
			BlockHash:      log.BlockHash,
			BlockNumber:    log.BlockNumber,
			TxIndex:        log.TxIndex,
			TxHash:         log.TxHash,
			SequenceNumber: seqNum,

			rawLog:       log,
			dataLocation: batchDataLocation(parsedLog.DataLocation),
		}
		if _, err := s.serialize(batch); err != nil {
			return nil, err
		}
		messages = append(messages, batch)
	}
	return messages, nil
}

func (s *SequencerInbox) getLogTransaction(log types.Log) (*types.Transaction, error) {
	tx, err := s.client.TransactionInBlock(context.Background(), log.BlockHash, log.TxIndex)
	if err != nil {
		return nil, err
	}
	if tx.Hash() != log.TxHash {
		return nil, fmt.Errorf("L1 client returned unexpected transaction hash %v when looking up block %v transaction %v with expected hash %v", tx.Hash(), log.BlockHash, log.TxIndex, log.TxHash)
	}
	return tx, nil
}

func (s *SequencerInbox) serialize(batch *SequencerBatch) ([]byte, error) {
	if batch.serialized != nil {
		return batch.serialized, nil
	}

	getSequencerBatchData := func() ([]byte, error) {
		switch batch.dataLocation {
		case batchDataTxInput:
			tx, err := s.getLogTransaction(batch.rawLog)
			if err != nil {
				return nil, err
			}
			args := make(map[string]interface{})
			err = addSequencerL2BatchFromOriginCallABI.Inputs.UnpackIntoMap(args, tx.Data()[4:])
			if err != nil {
				return nil, err
			}
			return args["data"].([]byte), nil
		case batchDataSeparateEvent:
			var numberAsHash common.Hash
			binary.BigEndian.PutUint64(numberAsHash[(32-8):], batch.SequenceNumber)
			query := ethereum.FilterQuery{
				BlockHash: &batch.BlockHash,
				Addresses: []common.Address{s.address},
				Topics:    [][]common.Hash{{sequencerBatchDataABI.ID}, {numberAsHash}},
			}
			logs, err := s.client.FilterLogs(context.Background(), query)
			if err != nil {
				return nil, err
			}
			if len(logs) == 0 {
				return nil, errors.New("expected to find sequencer batch data")
			}
			if len(logs) > 1 {
				return nil, errors.New("expected to find only one matching sequencer batch data")
			}
			event := new(arbinbox.ArbinboxSequencerBatchData)
			err = sequencerBridgeABI.UnpackIntoInterface(event, sequencerBatchDataEvent, logs[0].Data)
			if err != nil {
				return nil, err
			}
			return event.Data, nil
		case batchDataNone:
			// No data when in a force inclusion batch
			return nil, nil
		case batchDataBlobHashes:
			tx, err := s.getLogTransaction(batch.rawLog)
			if err != nil {
				return nil, err
			}
			if len(tx.BlobHashes()) == 0 {
				return nil, fmt.Errorf("blob batch transaction %v has no blobs", tx.Hash())
			}
			return []byte{BlobHashesHeaderFlag}, nil
		default:
			return nil, fmt.Errorf("batch has invalid data location %v", batch.dataLocation)
		}
	}

	fullData, err := getSequencerBatchData()
	if err != nil {
		return nil, err
	}

	batch.serialized = fullData
	return fullData, nil
}

type L1IncomingMessageHeader struct {
	Kind        uint8          `json:"kind"`
	Poster      common.Address `json:"sender"`
	BlockNumber uint64         `json:"blockNumber"`
	Timestamp   uint64         `json:"timestamp"`
	RequestId   *common.Hash   `json:"requestId" rlp:"nilList"`
	L1BaseFee   *big.Int       `json:"baseFeeL1"`
}

type L1IncomingMessage struct {
	Header *L1IncomingMessageHeader `json:"header"`
	L2msg  []byte                   `json:"l2Msg"`

	// Only used for `L1MessageType_BatchPostingReport`
	BatchGasCost *uint64 `json:"batchGasCost,omitempty" rlp:"optional"`
}

const BatchSegmentKindL2Message uint8 = 0
const BatchSegmentKindL2MessageBrotli uint8 = 1
const BatchSegmentKindDelayedMessages uint8 = 2
const BatchSegmentKindAdvanceTimestamp uint8 = 3
const BatchSegmentKindAdvanceL1BlockNumber uint8 = 4

const (
	L1MessageType_L2Message             = 3
	L1MessageType_EndOfBlock            = 6
	L1MessageType_L2FundedByL1          = 7
	L1MessageType_RollupEvent           = 8
	L1MessageType_SubmitRetryable       = 9
	L1MessageType_BatchForGasEstimation = 10 // probably won't use this in practice
	L1MessageType_Initialize            = 11
	L1MessageType_EthDeposit            = 12
	L1MessageType_BatchPostingReport    = 13
	L1MessageType_Invalid               = 0xFF
)

const MaxL2MessageSize = 256 * 1024

// parseL2Transactions parses L1IncomingMessage from the sequencer batch.
func (s *SequencerInbox) parseL2Transactions(batch *SequencerBatch) ([]*types.Transaction, error) {
	if batch.segments == nil {
		return nil, errors.New("batch segments not yet decompressed")
	}

	batch.txes = make([]*types.Transaction, 0)
	logger.Warnf("batch block number: %v tx hash: %v", batch.BlockNumber, batch.TxHash)
	for p, segment := range batch.segments {
		logger.Warnf("segment: %v", p)
		kind := segment[0]
		segment = segment[1:]
		if kind == BatchSegmentKindL2Message || kind == BatchSegmentKindL2MessageBrotli {
			if kind == BatchSegmentKindL2MessageBrotli {
				decompressed, err := decompressBytes(segment)
				if err != nil {
					return nil, nil
				}
				segment = decompressed
			}
			txes, err := parseL2Message(segment, nil, s.chainId, 0)
			if err != nil {
				return nil, err
			}
			batch.txes = append(batch.txes, txes...)
		}
	}

	return batch.txes, nil
}
