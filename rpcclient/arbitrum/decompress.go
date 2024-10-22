package arbitrum

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/andybalholm/brotli"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/offchainlabs/nitro/zeroheavy"

	"github.com/Lagrange-Labs/lsc-node/core/logger"
)

const (
	BlobHashesHeaderFlag       byte = 0x10
	ZeroheavyMessageHeaderFlag byte = 0x20
	BrotliMessageHeaderByte    byte = 0x00

	MaxDecompressedLen             int = 1024 * 1024 * 16 // 16 MiB
	maxZeroheavyDecompressedLen        = 101*MaxDecompressedLen/100 + 64
	MaxSegmentsPerSequencerMessage     = 100 * 1024
)

var brotliReader = brotli.NewReader(nil)

func decompress(data []byte) ([][]byte, error) {
	var segments [][]byte
	if data[0] == ZeroheavyMessageHeaderFlag {
		pl, err := io.ReadAll(io.LimitReader(zeroheavy.NewZeroheavyDecoder(bytes.NewReader(data[1:])), int64(maxZeroheavyDecompressedLen)))
		if err != nil {
			logger.Warnf("error reading from zeroheavy decoder, error: %v", err)
			return nil, err
		}
		data = pl
	}

	if data[0] == BrotliMessageHeaderByte {
		if err := brotliReader.Reset(bytes.NewReader(data[1:])); err != nil {
			logger.Warnf("error resetting brotli reader, error: %v", err)
			return nil, err
		}
		stream := rlp.NewStream(brotliReader, uint64(MaxDecompressedLen))
		for {
			var segment []byte
			err := stream.Decode(&segment)
			if err != nil {
				if !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
					logger.Warnf("error parsing sequencer message segment, error: %v", err)
				}
				break
			}
			if len(segments) >= MaxSegmentsPerSequencerMessage {
				logger.Warnf("too many segments in sequence batch")
				break
			}
			segments = append(segments, segment)
		}
	}
	return segments, nil
}

func decompressBytes(data []byte) ([]byte, error) {
	if err := brotliReader.Reset(bytes.NewReader(data)); err != nil {
		logger.Warnf("error resetting brotli reader, error: %v", err)
		return nil, err
	}
	return io.ReadAll(brotliReader)
}

// The number of bits in a BLS scalar that aren't part of a whole byte.
const spareBlobBits = 6 // = math.floor(math.log2(BLS_MODULUS)) % 8

// decodeBlobs decodes blobs into the batch data encoded in them.
func decodeBlobs(blobs []*eth.Blob) ([]byte, error) {
	var rlpData []byte
	for _, blob := range blobs {
		for fieldIndex := 0; fieldIndex < params.BlobTxFieldElementsPerBlob; fieldIndex++ {
			rlpData = append(rlpData, blob[fieldIndex*32+1:(fieldIndex+1)*32]...)
		}
		var acc uint16
		accBits := 0
		for fieldIndex := 0; fieldIndex < params.BlobTxFieldElementsPerBlob; fieldIndex++ {
			acc |= uint16(blob[fieldIndex*32]) << accBits
			accBits += spareBlobBits
			if accBits >= 8 {
				rlpData = append(rlpData, uint8(acc))
				acc >>= 8
				accBits -= 8
			}
		}
		if accBits != 0 {
			return nil, fmt.Errorf("somehow ended up with %v spare accBits", accBits)
		}
	}
	var outputData []byte
	err := rlp.Decode(bytes.NewReader(rlpData), &outputData)
	return outputData, err
}
