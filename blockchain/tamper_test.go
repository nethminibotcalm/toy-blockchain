package blockchain

import (
	"testing"
	"toy-blockchain/ledger"
)

func TestTamperDetection(t *testing.T) {

	bc := NewBlockchain()

	bc.AddBlock([]ledger.Transaction{
		{
			Sender:   "Alice",
			Receiver: "Bob",
			Amount:   20,
		},
	})

	// Simulate hacker changing transaction data
	bc.Blocks[1].Transactions[0].Amount = 9999

	if bc.ValidateChain() {
		t.Error("Tampered blockchain was accepted")
	}
}
