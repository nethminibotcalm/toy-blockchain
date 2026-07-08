package ledger

import "testing"

func TestRejectOverspendingTransaction(t *testing.T) {

	l := NewLedger()

	tx := Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   150,
	}

	result := l.ValidateTransaction(tx)

	if result {
		t.Error("Overspending transaction was accepted")
	}

	if l.Balances["Alice"] != 100 {
		t.Error("Alice balance changed after rejected transaction")
	}
}