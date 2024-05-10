package arbitrum

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/nitro/util/arbmath"
)

const (
	L2MessageKind_UnsignedUserTx  = 0
	L2MessageKind_ContractTx      = 1
	L2MessageKind_NonmutatingCall = 2
	L2MessageKind_Batch           = 3
	L2MessageKind_SignedTx        = 4
	// 5 is reserved
	L2MessageKind_Heartbeat          = 6 // deprecated
	L2MessageKind_SignedCompressedTx = 7
	// 8 is reserved for BLS signed batch
)

func parseL2Message(msg []byte, requestId *common.Hash, chainId *big.Int, depth int) (types.Transactions, error) {
	rd := bytes.NewReader(msg)
	var l2KindBuf [1]byte
	if _, err := rd.Read(l2KindBuf[:]); err != nil {
		return nil, err
	}

	switch l2KindBuf[0] {
	case L2MessageKind_UnsignedUserTx:
		logger.Infof("UnsignedUserTx")
		// do nothing
		return nil, nil
	case L2MessageKind_ContractTx:
		logger.Infof("ContractTx")
		// do nothing
		return nil, nil
	case L2MessageKind_NonmutatingCall:
		return nil, errors.New("L2 message kind NonmutatingCall is unimplemented")
	case L2MessageKind_Batch:
		logger.Infof("Batch")
		if depth >= 16 {
			return nil, errors.New("L2 message batches have a max depth of 16")
		}
		segments := make(types.Transactions, 0)
		index := big.NewInt(0)
		for {
			nextMsg, err := BytestringFromReader(rd, MaxL2MessageSize)
			if err != nil {
				// an error here means there are no further messages in the batch
				// nolint:nilerr
				return segments, nil
			}

			var nextRequestId *common.Hash
			if requestId != nil {
				subRequestId := crypto.Keccak256Hash(requestId[:], arbmath.U256Bytes(index))
				nextRequestId = &subRequestId
			}
			nestedSegments, err := parseL2Message(nextMsg, nextRequestId, chainId, depth+1)
			if err != nil {
				return nil, err
			}
			segments = append(segments, nestedSegments...)
			index.Add(index, big.NewInt(1))
		}
	case L2MessageKind_SignedTx:
		newTx := new(types.Transaction)
		// Safe to read in its entirety, as all input readers are limited
		readBytes, err := io.ReadAll(rd)
		if err != nil {
			return nil, err
		}
		if err := newTx.UnmarshalBinary(readBytes); err != nil {
			return nil, err
		}
		return types.Transactions{newTx}, nil
	case L2MessageKind_Heartbeat:
		// do nothing
		return nil, nil
	case L2MessageKind_SignedCompressedTx:
		return nil, errors.New("L2 message kind SignedCompressedTx is unimplemented")
	default:
		// ignore invalid message kind
		return nil, fmt.Errorf("unkown L2 message kind %v", l2KindBuf[0])
	}
}

func Uint64FromReader(rd io.Reader) (uint64, error) {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(rd, buf); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(buf), nil
}

func BytestringFromReader(rd io.Reader, maxBytesToRead uint64) ([]byte, error) {
	size, err := Uint64FromReader(rd)
	if err != nil {
		return nil, err
	}
	if size > maxBytesToRead {
		return nil, errors.New("size too large in ByteStringFromReader")
	}
	buf := make([]byte, size)
	if _, err = io.ReadFull(rd, buf); err != nil {
		return nil, err
	}
	return buf, nil
}
