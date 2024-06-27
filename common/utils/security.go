package utils

import (
	"fmt"
	"runtime"
	"time"

	"github.com/4kpros/go-api/config"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

// Encrypt map value using JWT HS256
func EncryptJWTToken(value interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"username": value,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenStr, _ := token.SignedString(config.AppEnvConfig.JwtTokenSecret)

	return tokenStr
}

// Decrypt JWT HS256 value
func DecryptJWTToken(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			message := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("%s", message)
		}

		return []byte(config.AppEnvConfig.JwtTokenSecret), nil
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

// Encrypt value using Argon2id
func EncryptWithArgon2id(value string) (hash string, err error) {
	params := &argon2id.Params{
		Memory:      uint32(config.CryptoEnvConfig.ArgonMemoryLeft * config.CryptoEnvConfig.ArgonMemoryRight),
		Iterations:  uint32(config.CryptoEnvConfig.ArgonIterations),
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  uint32(config.CryptoEnvConfig.ArgonSaltLength),
		KeyLength:   uint32(config.CryptoEnvConfig.ArgonKeyLength),
	}
	hash, err = argon2id.CreateHash(value, params)
	return
}

// Verify if Argon2id value matches string
func CompareToArgon2id(value string, encrypted string) (match bool, err error) {
	fmt.Printf("COMPARE ARGON2ID\n Value: %s \n Encrypted: %s", value, encrypted)
	match, err = argon2id.ComparePasswordAndHash(value, encrypted)
	return
}
