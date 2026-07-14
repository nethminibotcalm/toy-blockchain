package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func Verify(
	data string,
	signature string,
	publicKey string,
) bool {

	pubBytes, err := hex.DecodeString(publicKey)

	if err != nil {
		return false
	}

	sigBytes, err := hex.DecodeString(signature)

	if err != nil {
		return false
	}

	// Split public key
	x := new(big.Int).SetBytes(pubBytes[:len(pubBytes)/2])
	y := new(big.Int).SetBytes(pubBytes[len(pubBytes)/2:])

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// Split signature
	r := new(big.Int).SetBytes(sigBytes[:len(sigBytes)/2])
	s := new(big.Int).SetBytes(sigBytes[len(sigBytes)/2:])

	hash := sha256.Sum256([]byte(data))

	return ecdsa.Verify(
		&pub,
		hash[:],
		r,
		s,
	)
}
