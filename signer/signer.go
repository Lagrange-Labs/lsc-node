package signer

// Signer is the interface that wraps the basic Sign method.
type Signer interface {
	Sign([]byte) ([]byte, error)
	GetPubKey() ([]byte, error)
}
