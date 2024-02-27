package network

import (
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

// GenerateToken generates a JWT token.
func GenerateToken(operator string) (string, error) {
	claims := jwt.MapClaims{
		"operator": operator,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates a JWT token.
func ValidateToken(tokenString, operator string) (bool, error) {
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
