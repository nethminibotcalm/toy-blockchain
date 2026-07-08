package blockchain

import "testing"

func TestValidChain(t *testing.T) {

	bc := NewBlockchain()

	bc.AddBlock(nil)

	if !bc.ValidateChain() {
		t.Error("Valid blockchain failed validation")
	}
}