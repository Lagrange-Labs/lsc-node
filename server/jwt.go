package server

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Lagrange-Labs/lsc-node/core"
	"github.com/Lagrange-Labs/lsc-node/core/logger"
)

const (
	// TokenExpireTime is the expire time for the token.
	TokenExpireTime = time.Hour * 24
)

var secretKey = ""

// getSecretKey returns the secret key for JWT. Generate a new one based on the current time
// and the random string.
func getSecretKey(curTime int64) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(curTime))
	r := make([]byte, 24)
	_, err := rand.Read(r)
	if err != nil {
		panic(err)
	}
	b = append(b, r...)
	return fmt.Sprintf("%x", b)
}

// checkExpire checks if the secret key is expired.
func checkExpire(secretKey string) bool {
	b := core.Hex2Bytes(secretKey[:16])
	t := time.Unix(0, int64(binary.LittleEndian.Uint64(b)))

	return time.Since(t) > TokenExpireTime
}

// GenerateToken generates a JWT token.
func GenerateToken(operator string) (string, error) {
	if secretKey == "" || checkExpire(secretKey) {
		secretKey = getSecretKey(time.Now().UnixNano())
		logger.Infof("Generate new secret key %s", secretKey)
	}

	claims := jwt.MapClaims{
		"operator": operator,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates a JWT token.
func ValidateToken(tokenString, operator string) (bool, error) {
	if secretKey == "" || checkExpire(secretKey) {
		secretKey = getSecretKey(time.Now().UnixNano())
		logger.Infof("Generate new secret key %s", secretKey)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["operator"] == operator, nil
	}

	return false, nil
}
