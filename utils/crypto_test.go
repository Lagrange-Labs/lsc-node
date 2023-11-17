package utils

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/go-eth-consensus/bls"
)

func createTestRandomKeys(b *testing.B, count int) (secs []*bls.SecretKey, pubs []*bls.PublicKey) {
	for i := 0; i < count; i++ {
		sec, pub := RandomBlsKey()
		pubkey := new(bls.PublicKey)
		require.NoError(b, pubkey.Deserialize(common.FromHex(pub)))
		pubs = append(pubs, pubkey)
		secs = append(secs, sec)
	}
	return
}

func TestBLSSignature(t *testing.T) {
	sec, pub := RandomBlsKey()
	pubkey := new(bls.PublicKey)
	require.NoError(t, pubkey.Deserialize(common.FromHex(pub)))

	msg := common.Hex2Bytes("0x24796fc538ee62cea9791079ec6f54a292d05ac40e4fa00fb1f894325fe46067")
	sig, err := sec.Sign(msg)
	require.NoError(t, err)

	t.Logf("PubKey: %s, Signature: %s", pub, BlsSignatureToHex(sig))
	isVerified, err := VerifySignature(common.Hex2Bytes(pub), msg, common.Hex2Bytes(BlsSignatureToHex(sig)))
	require.NoError(t, err)
	require.True(t, isVerified)
}

func BenchmarkSign(b *testing.B) {
	sec, _ := RandomBlsKey()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sec.Sign([]byte("test"))
		require.NoError(b, err)
	}
}

func BenchmarkAggregatedSignature(b *testing.B) {
	benchmarks := []struct {
		testName string
		count    int
	}{
		{"small-100 testing", 100},
		{"medium-1000 testing", 1000},
		{"large-10000 testing", 10000},
	}

	testMsg := []byte("test")
	for _, bm := range benchmarks {
		secs, pubs := createTestRandomKeys(b, bm.count)
		sigs := make([]*bls.Signature, bm.count)
		for i, sec := range secs {
			sig, err := sec.Sign(testMsg)
			require.NoError(b, err)
			sigs[i] = sig
		}

		b.Run(bm.testName+" aggregated sign", func(sub *testing.B) {
			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				bls.AggregateSignatures(sigs)
			}
		})

		aggSig := bls.AggregateSignatures(sigs)
		b.Run(bm.testName+" aggregated signature verification", func(sub *testing.B) {
			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				isVerified, err := aggSig.FastAggregateVerify(pubs, testMsg)
				require.NoError(sub, err)
				require.True(sub, isVerified)
			}
		})
	}

}

func TestSignatureSplit(t *testing.T) {
	now := time.Now()
	res := GetSignatureAffine("9399a04fd3d10ca1354bf7de5d26161e4e8a44ecdfc0ba3791f3639e9e145cbffc35e824eff1799be7fb6fc90fecb54b03fca2a026b0ee2b4f35401972efe41942ff839924ae4c99f263c3afb1ffc919666a7b1ea78efecb6aa353ea3ce8abdb")
	t.Logf("GetSignatureAffine took %s", time.Since(now))
	require.Equal(t, common.Bytes2Hex(res), "03fca2a026b0ee2b4f35401972efe41942ff839924ae4c99f263c3afb1ffc919666a7b1ea78efecb6aa353ea3ce8abdb1399a04fd3d10ca1354bf7de5d26161e4e8a44ecdfc0ba3791f3639e9e145cbffc35e824eff1799be7fb6fc90fecb54b110859cf03915b530d8c93fbe9650daa613fed0e320fba9493d41acaaab87658e8f6742c3b5e090262166ce19da4b81d04ead4b67a375cf848d5c2dacfe5e002ac5cd78077ca370494e2a0c3435531b27e3a30350d3eb0dad59c716f3f86a587")
}
