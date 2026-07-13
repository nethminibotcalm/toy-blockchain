package blockchain

import (
	"testing"

	"toy-blockchain/ledger"
)

func TestDoubleSpendPrevention(t *testing.T) {

	bc := NewBlockchain()

	balances := CalculateBalances(bc.Blocks, bc.InitialBalances)

	l := ledger.NewLedger(balances)

	tx1 := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   80,
	}

	tx2 := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Charlie",
		Amount:   80,
	}

	bc.AddTransaction(tx1)
	bc.AddTransaction(tx2)

	bc.MinePendingTransactions(l)

	minedBlock := bc.Blocks[len(bc.Blocks)-1]

	if len(minedBlock.Transactions) != 1 {
		t.Fatalf(
			"expected only 1 transaction to be mined, got %d",
			len(minedBlock.Transactions),
		)
	}
}
