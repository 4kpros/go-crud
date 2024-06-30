package utils

import (
	"crypto/ecdsa"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/config"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

var JwtIss string = "go-api-jwt-iss-0011aazz=="

func NewOthersExpiresDate() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.AppEnv.JwtExpiresOthers))
	return &tempDate
}

func NewExpiresDate(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.AppEnv.JwtExpiresStayConnected))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.AppEnv.JwtExpires))
	return &tempDate
}

// Encrypt JWT
func EncryptJWTToken(value *types.JwtToken, privateKey string) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"iss":    JwtIss,
		"userId": fmt.Sprintf("%d", value.UserId),
		"role":   value.Role,
		"device": value.Device,
		"exp":    jwt.NewNumericDate(value.Expires),
		"iat":    jwt.NewNumericDate(time.Now()),
	})
	var signedKey *ecdsa.PrivateKey
	signedKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return
	}
	tokenStr, err = token.SignedString(signedKey)
	if err != nil {
		return
	}
	return
}

// Decrypt JWT
func DecryptJWTToken(tokenStr string, publicKey string) (*types.JwtToken, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (signedKey interface{}, err error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			message := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("%s", message)
		}
		signedKey, err = jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
		return
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		message := "Invalid token or expired! Please enter valid information."
		return nil, fmt.Errorf("%s", message)
	}
	iss := fmt.Sprintf("%s", claims["iss"])
	if iss != JwtIss {
		message := "Invalid token issuer! Please enter valid information."
		return nil, fmt.Errorf("%s", message)
	}
	userId, _ := strconv.Atoi(fmt.Sprintf("%s", claims["userId"]))
	role := fmt.Sprintf("%s", claims["role"])
	device := fmt.Sprintf("%s", claims["device"])
	return &types.JwtToken{
		UserId: uint(userId),
		Role:   role,
		Device: device,
	}, nil
}

// Verify that JWT token is valid
func VerifyJWTToken(tokenStr string, publicKey string) bool {
	decryptedToken, err := DecryptJWTToken(tokenStr, publicKey)
	if err != nil {
		fmt.Printf("\nDecryption of [%s] error ==> %s\n\n", tokenStr, err)
		return false
	}
	return decryptedToken != nil
}

// Encrypt value using Argon2id
func EncryptWithArgon2id(value string) (hash string, err error) {
	params := &argon2id.Params{
		Memory:      uint32(config.AppEnv.ArgonMemoryLeft * config.AppEnv.ArgonMemoryRight),
		Iterations:  uint32(config.AppEnv.ArgonIterations),
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  uint32(config.AppEnv.ArgonSaltLength),
		KeyLength:   uint32(config.AppEnv.ArgonKeyLength),
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
