package block

import (
	"testing"
	"toy-blockchain/ledger"
)

func TestHashDeterminism(t *testing.T) {

	b := Block{
		Index: 1,
		Timestamp: 12345,
		Transactions: []ledger.Transaction{
			{
				Sender: "Alice",
				Receiver: "Bob",
				Amount: 20,
			},
		},
		PreviousHash: "0000",
		Nonce: 10,
	}

	hash1 := CalculateHash(b)
	hash2 := CalculateHash(b)

	if hash1 != hash2 {
		t.Error("Hash is not deterministic")
	}
}