package blockchain

import (
	"fmt"
	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

func (bc *Blockchain) ResolveFork(candidate []block.Block) error {

	// Candidate chain must be longer
	if len(candidate) <= len(bc.Blocks) {
		return fmt.Errorf("candidate chain is not longer")
	}


	// Create temporary blockchain for validation
	temp := Blockchain{
		Blocks:              candidate,
		InitialBalances:     bc.InitialBalances,
		PendingTransactions: []ledger.Transaction{},
		Difficulty:          bc.Difficulty,
	}


	// Validate candidate chain
	if err := temp.ValidateChain(); err != nil {
		return fmt.Errorf("invalid candidate chain: %v", err)
	}


	// Accept candidate chain
	bc.Blocks = candidate


	return nil
}