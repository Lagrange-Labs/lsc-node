package crypto

type BLSCurve string

const (
	BLS12381 BLSCurve = "BLS12-381"
	BN254    BLSCurve = "BN254"
)

// BLSScheme is the interface for the BLS signature operations.
type BLSScheme interface {
	VerifySignature(pubKey, message, signature []byte, isG1 bool) (bool, error)
	AggregateSignatures(signatures [][]byte, isG1 bool) ([]byte, error)
	AggregatePublicKeys(pubKeys [][]byte, isG1 bool) ([]byte, error)
	VerifyAggregatedSignature(pubKeys [][]byte, message, signature []byte, isG1 bool) (bool, error)
	Sign(privKey, message []byte, isG1 bool) ([]byte, error)
	GenerateRandomKey() ([]byte, error)
	GetPublicKey(privKey []byte, isCompressed, isG1 bool) ([]byte, error)
	ConvertPublicKey(pubKey []byte, isCompressed, isG1 bool) ([]byte, error)
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
