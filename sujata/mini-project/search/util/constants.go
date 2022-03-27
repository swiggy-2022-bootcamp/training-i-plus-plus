package util

import "crypto/rsa"

const (
	PubKeyPath = "keys/app.rsa.pub"
)

// verify key
var (
	VerifyKey *rsa.PublicKey
)
