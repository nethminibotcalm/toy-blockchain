package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func Sign(data string, privateKey *ecdsa.PrivateKey) (string, error) {

	hash := sha256.Sum256([]byte(data))

	r, s, err := ecdsa.Sign(
		rand.Reader,
		privateKey,
		hash[:],
	)

	if err != nil {
		return "", err
	}

	signature := append(
		r.Bytes(),
		s.Bytes()...,
	)

	return hex.EncodeToString(signature), nil
}
