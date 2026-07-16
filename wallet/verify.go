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

	// P-256 coordinates are always 32 bytes
	if len(pubBytes) != 64 {
		return false
	}

	x := new(big.Int).SetBytes(pubBytes[:32])
	y := new(big.Int).SetBytes(pubBytes[32:])

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// r and s are also 32 bytes each
	if len(sigBytes) != 64 {
		return false
	}

	r := new(big.Int).SetBytes(sigBytes[:32])
	s := new(big.Int).SetBytes(sigBytes[32:])

	hash := sha256.Sum256([]byte(data))

	return ecdsa.Verify(
		&pub,
		hash[:],
		r,
		s,
	)
}
