package blockchain

import (
	"testing"
	"toy-blockchain/block"
)
func TestResolveForkAcceptsLongerValidChain(t *testing.T) {

	bc := NewBlockchain()


	candidate := append(
		[]block.Block{},
		bc.Blocks...,
	)


	// create a longer valid chain
	for i := 1; i <= 2; i++ {

		newBlock := block.Block{
			Index:        len(candidate),
			Timestamp:    int64(len(candidate)),
			PreviousHash: candidate[len(candidate)-1].Hash,
			Difficulty:   bc.Difficulty,
		}


		MineBlock(
			&newBlock,
			newBlock.Difficulty,
		)


		candidate = append(candidate, newBlock)
	}


	err := bc.ResolveFork(candidate)


	if err != nil {
		t.Fatalf("expected fork to be accepted: %v", err)
	}


	if len(bc.Blocks) != len(candidate) {

		t.Fatal("blockchain was not replaced")
	}
}
func TestResolveForkRejectsShorterChain(t *testing.T) {

	bc := NewBlockchain()


	shorter := bc.Blocks


	err := bc.ResolveFork(shorter)


	if err == nil {

		t.Fatal(
			"expected shorter chain to be rejected",
		)
	}

}
func TestResolveForkRejectsInvalidChain(t *testing.T) {

	bc := NewBlockchain()


	candidate := append(
		[]block.Block{},
		bc.Blocks...,
	)


	fake := block.Block{
		Index:        1,
		Timestamp:    10,
		PreviousHash: "wrong",
		Difficulty:   bc.Difficulty,
	}


	candidate = append(candidate, fake)


	err := bc.ResolveFork(candidate)


	if err == nil {

		t.Fatal(
			"expected invalid chain to be rejected",
		)
	}

}