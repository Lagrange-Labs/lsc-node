package types

import (
	"testing"

	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	synctypes "github.com/Lagrange-Labs/lagrange-node/synchronizer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/protobuf/proto"
)

func createTestRoundState() (*RoundState, []*bls.SecretKey) {
	chainHeader := &synctypes.ChainHeader{
		Chain:       "test",
		BlockNumber: 1,
		StateRoot:   "0xbb345e208bda953c908027a45aa443d6cab6b8d2fd64e83ec52f1008ddeafa58",
	}
	_, proposerPubKey := utils.RandomBlsKey()
	pBlock := &sequencertypes.Block{
		Header: &sequencertypes.BlockHeader{
			BlockNumber:       1,
			ProposerPubKey:    proposerPubKey,
			ParentHash:        "0x000000000",
			BlockHash:         "0x000000000",
			ProposerSignature: "0x000000000",
			CurrentCommittee:  "0x0000000",
			NextCommittee:     "0x0000000",
		},
		ChainHeader: chainHeader,
		Proof:       "0x0000000",
	}
	proposer := &Validator{
		PublicKey:   proposerPubKey,
		VotingPower: 0,
	}

	secKeys := []*bls.SecretKey{}

	nodes := []*sequencertypes.ClientNode{}
	for i := 0; i < 10; i++ {
		secKey, pubKey := utils.RandomBlsKey()
		secKeys = append(secKeys, secKey)
		node := &sequencertypes.ClientNode{
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

	chainHeader := &synctypes.ChainHeader{
		Chain:       "test",
		BlockNumber: 1,
		StateRoot:   "0xbb345e208bda953c908027a45aa443d6cab6b8d2fd64e83ec52f1008ddeafa58",
	}
	blsSignature := &sequencertypes.Signature{
		ChainHeader:      chainHeader,
		CurrentCommittee: "0x0000000",
		NextCommittee:    "0x0000000",
	}
	signMsg, err := proto.Marshal(blsSignature)
	require.NoError(t, err)

	// Test 1: valid case
	for i := 0; i < 10; i++ {
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
			BlockNumber:  1,
			PubKey:       rs.Validators.Validators[i].PublicKey,
			BlsSignature: &blsSign,
		})
	}
	_, _, err = rs.CheckAggregatedSignature()
	require.NoError(t, err)

	// Test 2: invalid case
	wrongSignature := ""
	rs, secKeys = createTestRoundState()
	for i := 0; i < 10; i++ {
		blsSign := *blsSignature //nolint:govet
		if i == 8 {
			blsSign.NextCommittee = "0x111" // wrong contents
		}
		signature, err := secKeys[i].Sign(signMsg)
		require.NoError(t, err)
		signatureMsg := signature.Serialize()
		blsSign.Signature = common.Bytes2Hex(signatureMsg[:])
		if i == 8 {
			wrongSignature = blsSign.Signature
		}
		if i == 7 {
			blsSign.Signature = "0x000" // invalid signature
		} else if i == 9 {
			blsSign.Signature = wrongSignature // wrong signature
		}
		pubKey := new(bls.PublicKey)
		require.NoError(t, pubKey.Deserialize(common.FromHex(rs.Validators.Validators[i].PublicKey)))

		verified, err := signature.VerifyByte(pubKey, signMsg)
		require.NoError(t, err)
		require.True(t, verified)

		rs.AddCommit(&networktypes.CommitBlockRequest{
			BlockNumber:  1,
			PubKey:       rs.Validators.Validators[i].PublicKey,
			BlsSignature: &blsSign,
		})
	}
	_, _, err = rs.CheckAggregatedSignature()
	require.Error(t, err)
	require.Len(t, rs.GetEvidences(), 3)
}
