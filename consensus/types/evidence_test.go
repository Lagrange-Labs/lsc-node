package types

import (
	"testing"

	sequencertypes "github.com/Lagrange-Labs/lsc-node/sequencer/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCommitHash(t *testing.T) {
	signature := &sequencertypes.BlsSignature{
		ChainHeader: &sequencertypes.ChainHeader{
			ChainId:       1337,
			BlockHash:     "0xafe58890693444d9116c940a5ff4418723e7f75869b30c9d8e4528e147cb4b7f",
			BlockNumber:   3,
			L1BlockNumber: 1,
			L1TxHash:      "0xafe58890693444d9116c940a5ff4418723e7f75869b30c9d8e4528e147cb4b7f",
		},
		CurrentCommittee: "0x9c11dac30afc6d443066d31976ece1015527da8d1c6f5e540ce649970f2e9129",
		NextCommittee:    "0x0538f196c8c36715f077e40f62b62795d83a4d82fddff30511375c9f6917a26b",
		BlsSignature:     "0xb3ad75be8554f25871e395268a2aec2d1d65003e70d4cd5b1560f37a85c7917fb82d66e22829c333043b4d6c3434151b13fb6b60d06f150132390f177c7891e97213c34cc843937f5e372035dcbb8be32ba6bf61a1545bdc2aafabd0fb60c5a4",
		EcdsaSignature:   "0x92d7f640e9b492e561046e4761438ce13ccc7bc4aa5d0d92ffc570cef559245a0c938002a0675a82eeb139cd60ddc7d2f2be232903d1d85a31f78b9c1230c62501",
	}

	reqHash := GetCommitRequestHash(signature)
	require.Equal(t, common.Bytes2Hex(reqHash), "b63341673d94ef9a7e86926be02601f40b1dc500be8a2b96bcc5b36d6c92690d")
}
