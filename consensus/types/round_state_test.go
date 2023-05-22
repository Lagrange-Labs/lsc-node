package types

import (
	"testing"

	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/protobuf/proto"
)

func createTestRoundState() (*RoundState, []*bls.SecretKey) {
	chainHeader := &sequencertypes.ChainHeader{
		ChainId:     1,
		BlockNumber: 1,
		BlockHash:   utils.RandomHex(32),
	}
	proposerSecKey, proposerPubKey := utils.RandomBlsKey()
	pBlock := &sequencertypes.Block{
		BlockHeader: &sequencertypes.BlockHeader{
			ProposerPubKey:   proposerPubKey,
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
		},
		ChainHeader: chainHeader,
	}

	blsSigMsg, _ := proto.Marshal(pBlock.BlsSignature())
	proposerSigMsg, _ := proposerSecKey.Sign(blsSigMsg)
	pBlock.BlockHeader.ProposerSignature = utils.BlsSignatureToHex(proposerSigMsg)

	proposer := &Validator{
		PublicKey:   proposerPubKey,
		VotingPower: 0,
	}

	secKeys := []*bls.SecretKey{}

	nodes := []networktypes.ClientNode{}
	for i := 0; i < 10; i++ {
		secKey, pubKey := utils.RandomBlsKey()
		secKeys = append(secKeys, secKey)
		node := networktypes.ClientNode{
			PublicKey:   pubKey,
			VotingPower: 1,
		}
		nodes = append(nodes, node)
	}

	validatorSet := NewValidatorSet(proposer, nodes)
	rs := NewEmptyRoundState()
	rs.UpdateRoundState(validatorSet, pBlock)

	return rs, secKeys
}

func TestCheckVotingPower(t *testing.T) {
	rs, _ := createTestRoundState()

	// Test 1: not enough case
	for i := 0; i < 6; i++ {
		rs.AddCommit(&networktypes.CommitBlockRequest{
			PubKey: rs.Validators.Validators[i].PublicKey,
		})
	}
	require.False(t, rs.CheckEnoughVotingPower())
	// Test 2: enough case
	rs.AddCommit(&networktypes.CommitBlockRequest{
		PubKey: rs.Validators.Validators[6].PublicKey,
	})
	require.True(t, rs.CheckEnoughVotingPower())
}

func TestCheckAggregatedSignature(t *testing.T) {
	rs, secKeys := createTestRoundState()

	blsSignature := rs.ProposalBlock.BlsSignature()
	signMsg, err := proto.Marshal(blsSignature)
	require.NoError(t, err)

	// Test 1: valid case
	for i := 0; i < len(secKeys); i++ {
		blsSign := *blsSignature //nolint:govet
		signature, err := secKeys[i].Sign(signMsg)
		require.NoError(t, err)
		signatureMsg := signature.Serialize()
		blsSign.Signature = common.Bytes2Hex(signatureMsg[:])

		pubKey := new(bls.PublicKey)
		require.NoError(t, pubKey.Deserialize(common.FromHex(rs.Validators.Validators[i].PublicKey)))

		verified, err := signature.VerifyByte(pubKey, signMsg)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(&networktypes.CommitBlockRequest{
			PubKey:       rs.Validators.Validators[i].PublicKey,
			BlsSignature: &blsSign,
		})
	}
	_, _, err = rs.CheckAggregatedSignature()
	require.NoError(t, err)

	// Test 2: invalid case
	wrongSignature := ""
	rs, secKeys = createTestRoundState()

	blsSignature = rs.ProposalBlock.BlsSignature()
	signMsg, err = proto.Marshal(blsSignature)
	require.NoError(t, err)

	for i := 0; i < len(secKeys); i++ {
		blsSign := *blsSignature //nolint:govet
		if i == 8 {
			blsSign.NextCommittee = "0x111" // wrong contents
		}
		signature, err := secKeys[i].Sign(signMsg)
		require.NoError(t, err)
		signatureMsg := signature.Serialize()
		blsSign.Signature = common.Bytes2Hex(signatureMsg[:])
		if i == 7 {
			blsSign.Signature = "0x000" // invalid signature
		} else if i == 8 {
			wrongSignature = blsSign.Signature
		} else if i == 9 {
			blsSign.Signature = wrongSignature // wrong signature
		}
		pubKey := new(bls.PublicKey)
		require.NoError(t, pubKey.Deserialize(common.FromHex(rs.Validators.Validators[i].PublicKey)))

		verified, err := signature.VerifyByte(pubKey, signMsg)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(&networktypes.CommitBlockRequest{
			PubKey:       rs.Validators.Validators[i].PublicKey,
			BlsSignature: &blsSign,
		})
	}
	_, _, err = rs.CheckAggregatedSignature()
	require.Error(t, err)
	require.Len(t, rs.evidences, 3)
}
