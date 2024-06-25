package utils

import (
	"fmt"
	"time"

	"github.com/4kpros/go-crud/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/sha3"
)

// Encrypt map value using JWT HS256
func EncryptJWTToken(value interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"username": value,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenStr, _ := token.SignedString(config.EnvConfig.JwtTokenSecret)

	return tokenStr
}

// Decrypt JWT HS256 value
func DecryptJWTToken(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.EnvConfig.JwtTokenSecret), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Verify that JWT token is valid
func VerifyJWTToken(tokenStr string) bool {
	decryptedToken, err := DecryptJWTToken(tokenStr)
	if err != nil {
		return false
	}
	return decryptedToken != nil
}

// Encrypt value using SHA3-512
func EncryptValue(value string) string {
	h := sha3.New512()
	h.Write([]byte(value))
	sum := h.Sum(nil)
	return string(sum)
}

// Decrypt SHA3-512 value
func DecryptValue(value string) string {
	h := sha3.New512()
	h.Write([]byte(value))
	sum := h.Sum(nil)
	return string(sum)
}
