package blockchain

import (
	"fmt"
	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

func (bc *Blockchain) ResolveFork(candidate []block.Block) error {
	candidateCopy := append(
		[]block.Block(nil),
		candidate...,
	)

	// Candidate chain must be longer

	// Candidate must share the same genesis block
	if len(candidateCopy) == 0 || len(bc.Blocks) == 0 {
		return fmt.Errorf("empty blockchain")
	}
	localWork := CalculateChainWork(bc.Blocks)
	candidateWork := CalculateChainWork(candidateCopy)

	if candidateWork.Cmp(localWork) <= 0 {
		return fmt.Errorf(
			"candidate chain does not have more cumulative work",
		)
	}

	if candidateCopy[0].Hash != bc.Blocks[0].Hash {
		return fmt.Errorf("candidate chain has different genesis block")
	}

	// Create temporary blockchain for validation
	temp := Blockchain{
		Blocks:              candidateCopy,
		InitialBalances:     bc.InitialBalances,
		PendingTransactions: []ledger.Transaction{},
		Difficulty:          bc.Difficulty,
	}

	// Validate candidate chain
	if err := temp.ValidateChain(); err != nil {
		return fmt.Errorf("invalid candidate chain: %v", err)
	}
	forkPoint := FindForkPoint(
		bc.Blocks,
		candidateCopy,
	)

	if forkPoint < 0 {
		return fmt.Errorf(
			"candidate chain has no common block",
		)
	}

	confirmedIDs := ConfirmedTransactionIDs(
		candidateCopy,
	)

	orphanedTransactions := OrphanedTransactions(
		bc.Blocks,
		forkPoint,
		confirmedIDs,
	)

	// Accept candidate chain
	bc.Blocks = candidateCopy
	oldPending := append(
		[]ledger.Transaction(nil),
		bc.PendingTransactions...,
	)

	bc.PendingTransactions = []ledger.Transaction{}
	transactionsToReconsider := append(
		oldPending,
		orphanedTransactions...,
	)
	seenPending := make(map[string]bool)

	for _, tx := range transactionsToReconsider {
		if tx.ID == "" {
			continue
		}

		if confirmedIDs[tx.ID] {
			continue
		}

		if seenPending[tx.ID] {
			continue
		}

		if bc.AddTransaction(tx) {
			seenPending[tx.ID] = true
		}
	}

	return nil
}
