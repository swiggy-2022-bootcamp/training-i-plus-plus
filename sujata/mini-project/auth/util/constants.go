package util

import "crypto/rsa"

const (
	PrivKeyPath = "keys/app.rsa"
	PubKeyPath  = "keys/app.rsa.pub"
)

// verify key and sign key
var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)
