package network

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTokenValidation(t *testing.T) {
	token, err := GenerateToken("operator")
	require.NoError(t, err)

	valid, err := ValidateToken(token, "operator")
	require.NoError(t, err)
	require.True(t, valid)

	valid, err = ValidateToken(token, "operator2")
	require.NoError(t, err)
	require.False(t, valid)
}

func TestTokenValidationExpired(t *testing.T) {
	tm := time.Now().UnixNano() - int64(time.Hour*25)
	sk := getSecretKey(tm)
	expired := checkExpire(sk)
	require.True(t, expired)

	secretKey = sk
	token, err := GenerateToken("operator")
	require.NoError(t, err)
	require.NotEqual(t, secretKey, sk)

	valid, err := ValidateToken(token, "operator")
	require.NoError(t, err)
	require.True(t, valid)
}
