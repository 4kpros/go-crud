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

var JwtIssuerSession = "session"
var JwtIssuerActivate = "activate"
var JwtIssuerResetCode = "resetCode"
var JwtIssuerResetNewPassword = "resetNewPassword"

func NewExpiresDateDefault() *time.Time {
	tempDate := time.Now().Add(time.Minute * time.Duration(config.AppEnv.JwtExpiresDefault))
	return &tempDate
}

func NewExpiresDateSignIn(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * time.Duration(24*config.AppEnv.JwtExpiresSignInStayConnected))
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * time.Duration(config.AppEnv.JwtExpiresSignIn))
	return &tempDate
}

func GetCachedKey(jwtToken *types.JwtToken) string {
	return fmt.Sprintf("%d%s%s", jwtToken.UserId, jwtToken.Issuer, jwtToken.Device)
}

func SameCachedKey(jwtToken1 *types.JwtToken, jwtToken2 *types.JwtToken) bool {
	if jwtToken1 == nil || jwtToken2 == nil {
		return false
	}
	if GetCachedKey(jwtToken1) != GetCachedKey(jwtToken2) {
		return false
	}

	return true
}

// Encrypt JWT
func EncryptJWTToken(jwtToken *types.JwtToken, privateKey string, loadCached bool) (newJwt *types.JwtToken, tokenStr string, err error) {
	// Check if there is some cached token
	if loadCached {
		tokenStr, err = config.GetMemcacheVal(GetCachedKey(jwtToken))
		if err == nil && len(tokenStr) > 0 {
			jwtDecrypted, errDecrypted := DecryptJWTToken(tokenStr, config.AppPem.JwtPublicKey)
			if errDecrypted == nil && SameCachedKey(jwtToken, jwtDecrypted) {
				newJwt = jwtDecrypted
				return
			}
		}
	}

	// Otherwise generate new one
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"iss":    jwtToken.Issuer,
		"userId": fmt.Sprintf("%d", jwtToken.UserId),
		"role":   jwtToken.Role,
		"device": jwtToken.Device,
		"exp":    jwt.NewNumericDate(jwtToken.Expires),
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

	// Cache new token
	config.SetMemcacheVal(GetCachedKey(jwtToken), tokenStr)
	newJwt = jwtToken
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
	userId, _ := strconv.Atoi(fmt.Sprintf("%s", claims["userId"]))
	role := fmt.Sprintf("%s", claims["role"])
	codeStr := fmt.Sprintf("%s", claims["code"])
	code, _ := strconv.Atoi(codeStr)
	device := fmt.Sprintf("%s", claims["device"])
	return &types.JwtToken{
		UserId: uint(userId),
		Role:   role,
		Code:   code,
		Device: device,
		Issuer: iss,
	}, nil
}

// Verify that JWT token is valid
func VerifyJWTToken(tokenStr string, publicKey string) bool {
	jwtDecrypted, errDecrypted := DecryptJWTToken(tokenStr, publicKey)
	if errDecrypted != nil || jwtDecrypted == nil {
		return false
	}
	tokenStrCached, errCached := config.GetMemcacheVal(GetCachedKey(jwtDecrypted))
	if errCached != nil || len(tokenStrCached) <= 0 {
		return false
	}
	if tokenStr != tokenStrCached {
		return false
	}
	return true
}

// Encrypt jwtToken using Argon2id
func EncryptWithArgon2id(jwtToken string) (hash string, err error) {
	params := &argon2id.Params{
		Memory:      uint32(config.AppEnv.ArgonMemoryLeft * config.AppEnv.ArgonMemoryRight),
		Iterations:  uint32(config.AppEnv.ArgonIterations),
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  uint32(config.AppEnv.ArgonSaltLength),
		KeyLength:   uint32(config.AppEnv.ArgonKeyLength),
	}
	hash, err = argon2id.CreateHash(jwtToken, params)
	return
}

// Verify if Argon2id jwtToken matches string
func CompareToArgon2id(jwtToken string, encrypted string) (match bool, err error) {
	fmt.Printf("COMPARE ARGON2ID\n Value: %s \n Encrypted: %s", jwtToken, encrypted)
	match, err = argon2id.ComparePasswordAndHash(jwtToken, encrypted)
	return
}
