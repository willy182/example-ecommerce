package rsa

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	privateKeyPath = "configs/rsa/app.rsa"
	publicKeyPath  = "configs/rsa/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// InitPublicKey return *rsa.PublicKey
func InitPublicKey() (*rsa.PublicKey, error) {
	verifyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}

// InitPrivateKey return *rsa.PrivateKey
func InitPrivateKey() (*rsa.PrivateKey, error) {
	signBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	return signKey, nil
}
