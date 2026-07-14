package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

// Create a new wallet
func NewWallet() (*Wallet, error) {

	privateKey, err := ecdsa.GenerateKey(
		elliptic.P256(),
		rand.Reader,
	)

	if err != nil {
		return nil, err
	}

	publicKey := append(
		privateKey.PublicKey.X.Bytes(),
		privateKey.PublicKey.Y.Bytes()...,
	)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// Return public key as string
func (w *Wallet) GetPublicKey() string {

	return hex.EncodeToString(w.PublicKey)

}
