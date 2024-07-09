package types

import (
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	servertypes "github.com/Lagrange-Labs/lagrange-node/server/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
)

func createTestRoundState(blsCurve crypto.BLSCurve) (*RoundState, [][]byte, []servertypes.ClientNode) {
	blsScheme := crypto.NewBLSScheme(blsCurve)

	proposerSecKey, _ := blsScheme.GenerateRandomKey()
	proposerPubKey, _ := blsScheme.GetPublicKey(proposerSecKey, true)
	pBatch := &sequencerv2types.Batch{
		BatchHeader: &sequencerv2types.BatchHeader{
			ChainId:       1,
			BatchNumber:   1,
			L1BlockNumber: 1,
			L1TxHash:      utils.RandomHex(32),
			L2Blocks: []*sequencerv2types.BlockHeader{
				{
					BlockNumber: 1,
					BlockHash:   utils.RandomHex(32),
				},
			},
		},
		CommitteeHeader: &sequencerv2types.CommitteeHeader{

			TotalVotingPower: 10,
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
		},
		ProposerPubKey: utils.Bytes2Hex(proposerPubKey),
	}

	blsSigHash := pBatch.BlsSignature().Hash()
	proposerSigMsg, _ := blsScheme.Sign(proposerSecKey, blsSigHash)
	pBatch.ProposerSignature = utils.Bytes2Hex(proposerSigMsg)

	secKeys := make([][]byte, 0)
	nodes := []servertypes.ClientNode{}
	for i := 0; i < 10; i++ {
		secKey, _ := blsScheme.GenerateRandomKey()
		pubKey, _ := blsScheme.GetPublicKey(secKey, true)
		secKeys = append(secKeys, secKey)
		addr := utils.RandomHex(20)
		node := servertypes.ClientNode{
			PublicKey:    utils.Bytes2Hex(pubKey),
			StakeAddress: addr,
			SignAddress:  addr,
			VotingPower:  1,
		}
		nodes = append(nodes, node)
	}

	rs := NewEmptyRoundState(blsScheme)
	rs.UpdateRoundState(pBatch)

	return rs, secKeys, nodes
}

func TestCheckVotingPower(t *testing.T) {
	rs, _, validators := createTestRoundState(crypto.BN254)
	vs := NewValidatorSet(validators, uint64(len(validators)))
	// Test 1: not enough case
	for i := 0; i < 6; i++ {
		require.NoError(t, rs.AddCommit(&sequencerv2types.BlsSignature{}, validators[i].PublicKey, validators[i].StakeAddress))
	}
	require.False(t, rs.CheckEnoughVotingPower(vs))
	// Test 2: enough case
	require.NoError(t, rs.AddCommit(&sequencerv2types.BlsSignature{}, validators[6].PublicKey, validators[6].StakeAddress))
	require.True(t, rs.CheckEnoughVotingPower(vs))
}

func TestCheckAggregatedSignature(t *testing.T) {
	rs, secKeys, validators := createTestRoundState(crypto.BN254)
	blsScheme := crypto.NewBLSScheme(crypto.BN254)

	blsSignature := rs.GetCurrentBatch().BlsSignature()
	sigHash := blsSignature.Hash()

	// Test 1: valid case
	for i := 0; i < len(secKeys); i++ {
		blsSign := blsSignature.Clone()
		signature, err := blsScheme.Sign(secKeys[i], sigHash)
		require.NoError(t, err)
		blsSign.BlsSignature = utils.Bytes2Hex(signature)

		verified, err := blsScheme.VerifySignature(utils.Hex2Bytes(validators[i].PublicKey), sigHash, signature)
		require.NoError(t, err)
		require.True(t, verified, i)

		require.NoError(t, rs.AddCommit(blsSign, validators[i].PublicKey, validators[i].StakeAddress))
	}
	err := rs.CheckAggregatedSignature()
	require.NoError(t, err)

	// Test 2: invalid case
	wrongSignature := ""
	rs, secKeys, validators = createTestRoundState(crypto.BLS12381)
	blsScheme = crypto.NewBLSScheme(crypto.BLS12381)
	blsSignature = rs.GetCurrentBatch().BlsSignature()
	sigHash = blsSignature.Hash()

	for i := 0; i < len(secKeys); i++ {
		blsSign := blsSignature.Clone()
		if i == 8 {
			blsSign.CommitteeHeader.NextCommittee = "0x111" // wrong contents
		}
		signature, err := blsScheme.Sign(secKeys[i], sigHash)
		require.NoError(t, err)
		blsSign.BlsSignature = utils.Bytes2Hex(signature)
		if i == 7 {
			blsSign.BlsSignature = "0x000" // invalid signature
		} else if i == 8 {
			wrongSignature = blsSign.BlsSignature
		} else if i == 9 {
			blsSign.BlsSignature = wrongSignature // wrong signature
		}

		verified, err := blsScheme.VerifySignature(utils.Hex2Bytes(validators[i].PublicKey), sigHash, signature)
		require.NoError(t, err)
		require.True(t, verified)

		require.NoError(t, rs.AddCommit(blsSign, validators[i].PublicKey, validators[i].StakeAddress))
	}
	err = rs.CheckAggregatedSignature()
	require.Error(t, err)
	require.Len(t, rs.evidences, 3)
	require.Len(t, rs.commitSignatures, 7)
}
