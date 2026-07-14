package blockchain

import (
	"testing"

	"toy-blockchain/block"
)

func TestConcurrentMining(t *testing.T) {

	b := block.Block{
		Index:        1,
		Timestamp:    1,
		PreviousHash: "abc",
	}

	attempts, _ := MineBlockConcurrent(&b, 3, 4)

	if attempts <= 0 {
		t.Fatal("Mining attempts should be greater than zero")
	}

	if b.Hash == "" {
		t.Fatal("Block hash should not be empty")
	}

	if len(b.Hash) < 3 || b.Hash[:3] != "000" {
		t.Fatal("Hash does not satisfy difficulty")
	}
}
