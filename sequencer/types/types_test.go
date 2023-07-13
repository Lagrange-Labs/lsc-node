package types

import (
	"encoding/binary"
	"encoding/json"
	"math/big"
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

func TestBlsSignatureHash(t *testing.T) {
	var blockNumberBuf common.Hash
	chainIDBuf := make([]byte, 4)

	blsSignature := &BlsSignature{
		ChainHeader: &ChainHeader{
			BlockHash:   "0x0257554703067ea1fe00f949e542f4a34074f67c457e00a1427d101b823f77fa",
			BlockNumber: 13202389,
			ChainId:     4321,
		},
		CurrentCommittee: "0x09f582a8133bb26ee103a78a78999466a84455f9a409c46b8599e1aebb95fc8e",
		NextCommittee:    "0x22355f09a8afa99cd6c98e0169af50e87d9b7ec858abb60542d4d0139d9aa496",
	}

	binary.BigEndian.PutUint32(chainIDBuf, uint32(blsSignature.ChainHeader.ChainId))
	big.NewInt(int64(blsSignature.ChainHeader.BlockNumber)).FillBytes(blockNumberBuf[:])
	t.Logf("blockNumberBuf: %s", common.Bytes2Hex(blockNumberBuf[:]))
	t.Logf("chainIDBuf: %s", common.Bytes2Hex(chainIDBuf))
	chainHash := utils.Hash(common.FromHex(blsSignature.ChainHeader.BlockHash), blockNumberBuf[:], chainIDBuf)
	require.Equal(t, common.Bytes2Hex(chainHash), "06c3a68be875459bbc887f758c9d4aab01b7eb997daa25b91c99c9fb1e35f14a")

	blsHash := blsSignature.Hash()
	require.Equal(t, common.Bytes2Hex(blsHash), "0c6dd657c4fa2048bd8fc9d135207485bab7aad9fff6b70c17a2b787a8bcb52e")
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
					BlockHash:   common.Bytes2Hex(utils.Hash(common.Hex2Bytes(utils.RandomHex(32)))),
					BlockNumber: uint64(i + 1),
					ChainId:     5,
				},
				CurrentCommittee: common.Bytes2Hex(utils.PoseidonHash(common.Hex2Bytes(utils.RandomHex(32)))),
				NextCommittee:    common.Bytes2Hex(utils.PoseidonHash(common.Hex2Bytes(utils.RandomHex(32)))),
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
