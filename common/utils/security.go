package utils

import (
	"crypto/ecdsa"
	"fmt"
	"runtime"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/config"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

func NewOthersExpiresDate() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.AppEnv.JwtExpiresOthers))
	return &tempDate
}

func NewExpiresDate(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.AppEnv.JwtExpiresOthers))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.AppEnv.JwtExpiresOthers))
	return &tempDate
}

// Encrypt JWT
func EncryptJWTToken(value *types.JwtToken) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"iss":    "go-api",
		"id":     value.UserId,
		"role":   value.Role,
		"device": value.Device,
		"exp":    value.Expires.Unix(),
		"iat":    time.Now().Unix(),
	})
	fmt.Printf("\nUnsigned key ==> %s \n", config.AppPem.JwtPrivateKey)
	var signedKey *ecdsa.PrivateKey
	signedKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(config.AppPem.JwtPrivateKey))
	if err != nil {
		fmt.Printf("\nEncrypt token error 1 ==> %s \n", err)
		return
	}
	fmt.Printf("\nSigned key ==> %s \n", signedKey)
	tokenStr, err = token.SignedString(signedKey)
	if err != nil {
		fmt.Printf("\nEncrypt token error 2 ==> %s \n", err)
		return
	}
	fmt.Printf("\nEncrypt token success ==> %s \n", tokenStr)
	return
}

// Decrypt JWT
func DecryptJWTToken(tokenStr string) (*types.JwtToken, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			message := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			fmt.Printf("\nERROR ==> %s \n\n", message)
			return nil, fmt.Errorf("%s", message)
		}
		publicKey, errParsing := jwt.ParseECPublicKeyFromPEM([]byte(config.AppPem.JwtPrivateKey))
		return publicKey, errParsing
	})

	if err != nil {
		fmt.Printf("\nERROR ==> %s \n\n", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		message := "Invalid token claim!"
		return nil, fmt.Errorf("%s", message)
	}
	fmt.Printf("\nJWT claim ==> %s \n\n", claims["exp"])
	fmt.Printf("\nJWT claim ==> %s \n\n", claims["id"])
	fmt.Printf("\nJWT claim ==> %s \n\n", claims["device"])
	fmt.Printf("\nJWT claim ==> %s \n\n", claims["sub"])
	fmt.Printf("\nJWT claim ==> %s \n\n", claims)
	return nil, nil
}

// Verify that JWT token is valid
func VerifyJWTToken(tokenStr string) bool {
	decryptedToken, err := DecryptJWTToken(tokenStr)
	if err != nil {
		fmt.Printf("\nDecryption of [%s] error ==> %s\n\n", tokenStr, err)
		return false
	}
	fmt.Printf("\nDecrypted token ==> %s\n\n", decryptedToken)
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
