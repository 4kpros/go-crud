package config

import (
	"fmt"
	"os"
)

type Pem struct {
	JwtPrivateKey string
	JwtPublicKey  string
}

var AppPem = &Pem{}

func LoadPem() (err error) {
	AppPem.JwtPrivateKey, err = ReadFileToString("jwt_private.pem")
	AppPem.JwtPublicKey, err = ReadFileToString("jwt_public.pem")
	return
}

func ReadFileToString(path string) (contentStr string, err error) {
	content, errRead := os.ReadFile(path)
	if errRead != nil {
		fmt.Printf("\nERROR ReadFileToString ==> %s", err)
		err = errRead
		return
	}
	contentStr = string(content)
	return
}
