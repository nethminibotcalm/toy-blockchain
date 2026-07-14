package wallet

import (
	"crypto/elliptic"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
)

type walletFile struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// Save a wallet to wallets/<name>.json
func SaveWallet(name string, w *Wallet) error {

	err := os.MkdirAll("wallets", 0755)
	if err != nil {
		return err
	}

	privateBytes, err := x509.MarshalECPrivateKey(w.PrivateKey)
	if err != nil {
		return err
	}

	data := walletFile{
		PrivateKey: hex.EncodeToString(privateBytes),
		PublicKey:  w.GetPublicKey(),
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join("wallets", name+".json")

	return os.WriteFile(path, bytes, 0644)
}

// Load wallets/<name>.json
func LoadWallet(name string) (*Wallet, error) {

	path := filepath.Join("wallets", name+".json")

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data walletFile

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	privateBytes, err := hex.DecodeString(data.PrivateKey)
	if err != nil {
		return nil, err
	}

	privateKey, err := x509.ParseECPrivateKey(privateBytes)
	if err != nil {
		return nil, err
	}

	publicKey := append(
		privateKey.PublicKey.X.Bytes(),
		privateKey.PublicKey.Y.Bytes()...,
	)

	privateKey.PublicKey.Curve = elliptic.P256()

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}
