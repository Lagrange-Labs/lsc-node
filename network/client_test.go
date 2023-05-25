package network

import (
	"strings"
	"testing"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestECDSASignVerify(t *testing.T) {
	// sign the CommitBlockRequest
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix("0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429", "0x"))
	require.NoError(t, err)

	req := &types.CommitBlockRequest{
		BlsSignature: &sequencertypes.BlsSignature{
			ChainHeader: &sequencertypes.ChainHeader{
				ChainId:     1,
				BlockHash:   utils.RandomHex(32),
				BlockNumber: 1,
			},
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
			Signature:        utils.RandomHex(32),
		},
		EpochNumber: 1,
		PubKey:      utils.RandomHex(32),
	}
	reqMsg := contypes.GetCommitRequestHash(req)
	sig, err := crypto.Sign(reqMsg, privateKey)
	require.NoError(t, err)
	// verify the signature
	isVerified, addr, err := utils.VerifyECDSASignature(reqMsg, sig)
	require.NoError(t, err)
	require.True(t, isVerified)
	require.Equal(t, addr, "516d6c27c23ced21bf7930e2a01f0bca9a141a0d")
}
