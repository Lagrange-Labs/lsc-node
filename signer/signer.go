package signer

// Signer is the interface that wraps the basic Sign method.
type Signer interface {
	Sign([]byte, bool) ([]byte, error)
	GetPubKey(bool) ([]byte, error)
}
