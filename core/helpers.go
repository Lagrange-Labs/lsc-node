package core

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"reflect"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
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

// Uint64ToBytes converts a uint64 to bytes.
func Uint64ToBytes(num uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, num)
	return b
}

// Hex2Bytes converts a hex string to bytes.
func Hex2Bytes(hex string) []byte {
	return common.FromHex(hex)
}

// Bytes2Hex converts bytes to a hex string.
func Bytes2Hex(bytes []byte) string {
	return common.Bytes2Hex(bytes)
}

// GetValidAddress returns a valid address from the given hex string.
func GetValidAddress(hex string) string {
	return common.HexToAddress(hex).Hex()
}

// GetPrivateField returns the private field of the struct instance.
func GetPrivateField(instance interface{}, fieldName string) unsafe.Pointer {
	// Get the reflect.Value of the struct instance
	value := reflect.ValueOf(instance)

	// Get the reflect.Value of the private field by name
	privateFieldValue := value.Elem().FieldByName(fieldName)

	// Return the interface{} value of the private field
	return unsafe.Pointer(privateFieldValue.UnsafeAddr())
}
