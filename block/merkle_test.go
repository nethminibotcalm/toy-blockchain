package block_test

import (
	"testing"

	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

func TestMerkleRoot(t *testing.T) {

	transactions := []ledger.Transaction{

		{
			Sender:   "Alice",
			Receiver: "Bob",
			Amount:   10,
		},

		{
			Sender:   "Bob",
			Receiver: "Charlie",
			Amount:   5,
		},
	}

	root := block.CalculateMerkleRoot(
		transactions,
	)

	if root == "" {

		t.Fatal(
			"Merkle root should not be empty",
		)

	}

	t.Log(
		"Merkle Root:",
		root,
	)

}
