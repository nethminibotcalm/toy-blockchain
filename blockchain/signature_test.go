package blockchain

import (
	"testing"

	"toy-blockchain/ledger"
)

func TestFakeSignatureRejected(t *testing.T) {

	bc := NewBlockchain()

	fakeTx := ledger.Transaction{
		Sender:    "Alice",
		Receiver:  "Bob",
		Amount:    20,
		PublicKey: "fake-public-key",
		Signature: "fake-signature",
	}

	result := bc.AddTransaction(fakeTx)

	if result {
		t.Fatal("Fake signature was accepted")
	}
}