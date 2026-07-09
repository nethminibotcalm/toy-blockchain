package blockchain

import (
	"strings"
	"testing"
	"toy-blockchain/ledger"
)

func TestMiningDifficulty(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   50,
	}

	bc.AddTransaction(tx)

	l := ledger.NewLedger()

	bc.MinePendingTransactions(l)

	if len(bc.Blocks) < 2 {
		t.Fatal("expected a mined block to be added")
	}

	minedBlock := bc.Blocks[len(bc.Blocks)-1]

	difficulty := 4

	prefix := strings.Repeat("0", difficulty)

	if !strings.HasPrefix(minedBlock.Hash, prefix) {
		t.Fatalf(
			"hash %s does not satisfy difficulty %d",
			minedBlock.Hash,
			difficulty,
		)
	}
}