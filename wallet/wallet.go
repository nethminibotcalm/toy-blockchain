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

func NewWallet() (*Wallet, error) {

	privateKey, err := ecdsa.GenerateKey(
		elliptic.P256(),
		rand.Reader,
	)

	if err != nil {
		return nil, err
	}

	xBytes := make([]byte, 32)
	yBytes := make([]byte, 32)

	privateKey.PublicKey.X.FillBytes(xBytes)
	privateKey.PublicKey.Y.FillBytes(yBytes)

	publicKey := append(xBytes, yBytes...)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func (w *Wallet) GetPublicKey() string {

	return hex.EncodeToString(w.PublicKey)

}
