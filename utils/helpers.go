package utils

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// TimeDuration is a wrapper around time.Duration that allows us to unmarshal in TOML.
type TimeDuration time.Duration

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *TimeDuration) UnmarshalText(text []byte) error {
	parsedDuration, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = TimeDuration(parsedDuration)
	return nil
}

// RandomHex generates a random hex string of length n.
func RandomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
