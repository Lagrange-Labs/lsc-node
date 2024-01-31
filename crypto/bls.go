package crypto

type BLSCurve string

const (
	BLS12381 BLSCurve = "BLS12-381"
	BN254    BLSCurve = "BN254"
)

// BLSScheme is the interface for the BLS signature operations.
type BLSScheme interface {
	VerifySignature(pubKey, message, signature []byte) (bool, error)
	AggregateSignatures(signatures [][]byte) ([]byte, error)
	AggregatePublicKeys(pubKeys [][]byte) ([]byte, error)
	VerifyAggregatedSignature(pubKeys [][]byte, message, signature []byte) (bool, error)
	Sign(privKey, message []byte) ([]byte, error)
	GenerateRandomKey() ([]byte, error)
	GetPublicKey(privKey []byte) ([]byte, error)
}

// NewBLSScheme returns a new BLS scheme implementation.
func NewBLSScheme(curve BLSCurve) BLSScheme {
	switch curve {
	case BLS12381:
		return &BLS12381Scheme{}
	case BN254:
		return &BN254Scheme{}
	default:
		panic("invalid curve: " + curve)
	}
}
