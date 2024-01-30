package types

import (
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
)

func createTestRoundState(blsCurve crypto.BLSCurve) (*RoundState, [][]byte, *ValidatorSet) {
	blsScheme := crypto.NewBLSScheme(blsCurve)

	chainHeader := &sequencertypes.ChainHeader{
		ChainId:       1,
		BlockNumber:   1,
		BlockHash:     utils.RandomHex(32),
		L1BlockNumber: 1,
		L1TxHash:      utils.RandomHex(32),
	}
	proposerSecKey, _ := blsScheme.GenerateRandomKey()
	proposerPubKey, _ := blsScheme.GetPublicKey(proposerSecKey)
	pBlock := &sequencertypes.Block{
		BlockHeader: &sequencertypes.BlockHeader{
			ProposerPubKey:   utils.Bytes2Hex(proposerPubKey),
			TotalVotingPower: 10,
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
		},
		ChainHeader: chainHeader,
	}

	blsSigHash := pBlock.BlsSignature().Hash()
	proposerSigMsg, _ := blsScheme.Sign(proposerSecKey, blsSigHash)
	pBlock.BlockHeader.ProposerSignature = utils.Bytes2Hex(proposerSigMsg)

	secKeys := make([][]byte, 0)
	nodes := []networktypes.ClientNode{}
	for i := 0; i < 10; i++ {
		secKey, _ := blsScheme.GenerateRandomKey()
		pubKey, _ := blsScheme.GetPublicKey(secKey)
		secKeys = append(secKeys, secKey)
		node := networktypes.ClientNode{
			PublicKey:    pubKey,
			StakeAddress: utils.RandomHex(20),
			VotingPower:  1,
		}
		nodes = append(nodes, node)
	}

	vs := NewValidatorSet(nodes, uint64(len(nodes)))

	rs := NewEmptyRoundState(blsScheme)
	rs.UpdateRoundState(pBlock)

	return rs, secKeys, vs
}

func TestCheckVotingPower(t *testing.T) {
	rs, _, vs := createTestRoundState(crypto.BN254)

	// Test 1: not enough case
	for i := 0; i < 6; i++ {
		rs.AddCommit(&sequencertypes.BlsSignature{}, vs.validators[i].BlsPubKey, vs.validators[i].StakeAddress)
	}
	require.False(t, rs.CheckEnoughVotingPower(vs))
	// Test 2: enough case
	rs.AddCommit(&sequencertypes.BlsSignature{}, vs.validators[6].BlsPubKey, vs.validators[6].StakeAddress)
	require.True(t, rs.CheckEnoughVotingPower(vs))
}

func TestCheckAggregatedSignature(t *testing.T) {
	rs, secKeys, vs := createTestRoundState(crypto.BN254)
	blsScheme := crypto.NewBLSScheme(crypto.BN254)

	blsSignature := rs.GetCurrentBlock().BlsSignature()
	sigHash := blsSignature.Hash()

	// Test 1: valid case
	for i := 0; i < len(secKeys); i++ {
		blsSign := blsSignature.Clone()
		signature, err := blsScheme.Sign(secKeys[i], sigHash)
		require.NoError(t, err)
		blsSign.BlsSignature = utils.Bytes2Hex(signature)

		verified, err := blsScheme.VerifySignature(vs.validators[i].BlsPubKey, sigHash, signature)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(blsSign, vs.validators[i].BlsPubKey, vs.validators[i].StakeAddress)
	}
	err := rs.CheckAggregatedSignature()
	require.NoError(t, err)

	// Test 2: invalid case
	wrongSignature := ""
	rs, secKeys, vs = createTestRoundState(crypto.BLS12381)
	blsScheme = crypto.NewBLSScheme(crypto.BLS12381)
	blsSignature = rs.GetCurrentBlock().BlsSignature()
	sigHash = blsSignature.Hash()

	for i := 0; i < len(secKeys); i++ {
		blsSign := *blsSignature //nolint:govet
		if i == 8 {
			blsSign.NextCommittee = "0x111" // wrong contents
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

		verified, err := blsScheme.VerifySignature(vs.validators[i].BlsPubKey, sigHash, signature)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(&blsSign, vs.validators[i].BlsPubKey, vs.validators[i].StakeAddress)
	}
	err = rs.CheckAggregatedSignature()
	require.Error(t, err)
	require.Len(t, rs.evidences, 3)
}
