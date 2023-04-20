package utils

import (
	"testing"

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
