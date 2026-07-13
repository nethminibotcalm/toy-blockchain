package blockchain

import "testing"

func TestValidChain(t *testing.T) {

	bc := NewBlockchain()

	bc.AddBlock(nil)

	if err := bc.ValidateChain(); err != nil {
		t.Fatal(err)
	}
}
