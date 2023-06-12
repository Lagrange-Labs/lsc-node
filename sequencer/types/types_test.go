package types

import (
	"encoding/json"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
)

type BlsSignatureVector struct {
	BlsSignature *BlsSignature `json:"bls_signature"`
	Hash         string        `json:"hash"`
	BlsSecKey    string        `json:"bls_sec_key"`
	BlsPubKey    string        `json:"bls_pub_key"`
	BlsSigMsg    string        `json:"bls_sig"`
}

func TestGenerateVector(t *testing.T) {
	sec, pub := utils.RandomBlsKey()
	pubkey := new(bls.PublicKey)
	require.NoError(t, pubkey.Deserialize(common.FromHex(pub)))

	sigs := make([]*BlsSignatureVector, 10)
	for i := 0; i < 10; i++ {
		b := &BlsSignatureVector{
			BlsSignature: &BlsSignature{
				ChainHeader: &ChainHeader{
					BlockHash:   utils.RandomHex(32),
					BlockNumber: uint64(i + 1),
					ChainId:     5,
				},
				CurrentCommittee: utils.RandomHex(32),
				NextCommittee:    utils.RandomHex(32),
				TotalVotingPower: 1000000,
			},
			BlsSecKey: utils.BlsPrivKeyToHex(sec),
			BlsPubKey: utils.BlsPubKeyToHex(pubkey),
		}
		b.Hash = common.Bytes2Hex(b.BlsSignature.Hash())
		sig, err := sec.Sign(common.FromHex(b.Hash))
		require.NoError(t, err)
		b.BlsSigMsg = utils.BlsSignatureToHex(sig)

		isVerified, err := utils.VerifySignature(common.FromHex(b.BlsPubKey), common.FromHex(b.Hash), common.FromHex(b.BlsSigMsg))
		require.NoError(t, err)
		require.True(t, isVerified)

		sigs[i] = b
	}
	msg, err := json.Marshal(sigs)
	require.NoError(t, err)
	t.Logf("%s", msg)
}
