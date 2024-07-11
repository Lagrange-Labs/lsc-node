package v2

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	servertypes "github.com/Lagrange-Labs/lagrange-node/server/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
)

// CommitteeRoot is the root of the committee.
type CommitteeRoot struct {
	ChainID               uint32                   `json:"chain_id" bson:"chain_id"`
	CurrentCommitteeRoot  string                   `json:"current_committee_root" bson:"current_committee_root"`
	EpochNumber           uint64                   `json:"epoch_number" bson:"epoch_number"`
	EpochStartBlockNumber uint64                   `json:"epoch_start_block_number" bson:"epoch_start_block_number"`
	TotalVotingPower      uint64                   `json:"total_voting_power" bson:"total_voting_power"`
	Operators             []servertypes.ClientNode `json:"operators" bson:"operators"`
}

// GetLeafHash returns the leaf hash of the operator info.
func GetLeafHash(addr, pubKey []byte, votingPower uint64) []byte {
	res := make([]byte, 0, 32+32+20+12)

	res = append(res, pubKey...)
	res = append(res, addr...)
	res = append(res, common.LeftPadBytes(big.NewInt(int64(votingPower)).Bytes(), 12)...)

	return res
}

// Verify verifies the committee root.
func (c *CommitteeRoot) Verify() error {
	leaves := make([][]byte, len(c.Operators))
	for i, op := range c.Operators {
		leaves[i] = GetLeafHash(utils.Hex2Bytes(op.StakeAddress), utils.Hex2Bytes(op.PublicKey), op.VotingPower)
	}
	root := crypto.MerkleRoot(leaves)
	if !bytes.Equal(utils.Hex2Bytes(c.CurrentCommitteeRoot), root) {
		return fmt.Errorf("invalid committee root %s, expected %s", c.CurrentCommitteeRoot, utils.Bytes2Hex(root))
	}

	return nil
}
