package types

import (
	"testing"

	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
)

func createTestRoundState() (*RoundState, []*bls.SecretKey, *ValidatorSet) {
	chainHeader := &sequencertypes.ChainHeader{
		ChainId:       1,
		BlockNumber:   1,
		BlockHash:     utils.RandomHex(32),
		L1BlockNumber: 1,
		L1TxHash:      utils.RandomHex(32),
	}
	proposerSecKey, proposerPubKey := utils.RandomBlsKey()
	pBlock := &sequencertypes.Block{
		BlockHeader: &sequencertypes.BlockHeader{
			ProposerPubKey:   proposerPubKey,
			TotalVotingPower: 10,
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
		},
		ChainHeader: chainHeader,
	}

	blsSigHash := pBlock.BlsSignature().Hash()
	proposerSigMsg, _ := proposerSecKey.Sign(blsSigHash)
	pBlock.BlockHeader.ProposerSignature = utils.BlsSignatureToHex(proposerSigMsg)

	secKeys := []*bls.SecretKey{}

	nodes := []networktypes.ClientNode{}
	for i := 0; i < 10; i++ {
		secKey, pubKey := utils.RandomBlsKey()
		secKeys = append(secKeys, secKey)
		node := networktypes.ClientNode{
			PublicKey:    pubKey,
			StakeAddress: utils.RandomHex(20),
			VotingPower:  1,
		}
		nodes = append(nodes, node)
	}

	vs := NewValidatorSet(nodes, uint64(len(nodes)))

	rs := NewEmptyRoundState()
	rs.UpdateRoundState(pBlock)

	return rs, secKeys, vs
}

func TestCheckVotingPower(t *testing.T) {
	rs, _, vs := createTestRoundState()

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
	rs, secKeys, vs := createTestRoundState()

	blsSignature := rs.GetCurrentBlock().BlsSignature()
	sigHash := blsSignature.Hash()

	// Test 1: valid case
	for i := 0; i < len(secKeys); i++ {
		blsSign := blsSignature.Clone()
		signature, err := secKeys[i].Sign(sigHash)
		require.NoError(t, err)
		signatureMsg := signature.Serialize()
		blsSign.BlsSignature = common.Bytes2Hex(signatureMsg[:])

		pubKey := new(bls.PublicKey)
		require.NoError(t, pubKey.Deserialize(common.FromHex(vs.validators[i].BlsPubKey)))

		verified, err := signature.VerifyByte(pubKey, sigHash)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(blsSign, vs.validators[i].BlsPubKey, vs.validators[i].StakeAddress)
	}
	err := rs.CheckAggregatedSignature()
	require.NoError(t, err)

	// Test 2: invalid case
	wrongSignature := ""
	rs, secKeys, vs = createTestRoundState()

	blsSignature = rs.GetCurrentBlock().BlsSignature()
	sigHash = blsSignature.Hash()

	for i := 0; i < len(secKeys); i++ {
		blsSign := *blsSignature //nolint:govet
		if i == 8 {
			blsSign.NextCommittee = "0x111" // wrong contents
		}
		signature, err := secKeys[i].Sign(sigHash)
		require.NoError(t, err)
		signatureMsg := signature.Serialize()
		blsSign.BlsSignature = common.Bytes2Hex(signatureMsg[:])
		if i == 7 {
			blsSign.BlsSignature = "0x000" // invalid signature
		} else if i == 8 {
			wrongSignature = blsSign.BlsSignature
		} else if i == 9 {
			blsSign.BlsSignature = wrongSignature // wrong signature
		}
		pubKey := new(bls.PublicKey)
		require.NoError(t, pubKey.Deserialize(common.FromHex(vs.validators[i].BlsPubKey)))

		verified, err := signature.VerifyByte(pubKey, sigHash)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(&blsSign, vs.validators[i].BlsPubKey, vs.validators[i].StakeAddress)
	}
	err = rs.CheckAggregatedSignature()
	require.Error(t, err)
	require.Len(t, rs.evidences, 3)
}
