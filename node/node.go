package node

import (
	"sync"

	"toy-blockchain/blockchain"
)

type Node struct {
	Config           Config
	Blockchain       *blockchain.Blockchain
	mu               sync.RWMutex
	seenTransactions map[string]bool
}

func NewNode(
	config Config,
	chain *blockchain.Blockchain,
) *Node {
	seenTransactions := make(map[string]bool)

	// Remember transactions already stored in mined blocks.
	for _, currentBlock := range chain.Blocks {
		for _, tx := range currentBlock.Transactions {
			if tx.ID != "" {
				seenTransactions[tx.ID] = true
			}
		}
	}

	// Remember transactions currently waiting in the pending pool.
	for _, tx := range chain.PendingTransactions {
		if tx.ID != "" {
			seenTransactions[tx.ID] = true
		}
	}

	return &Node{
		Config:           config,
		Blockchain:       chain,
		seenTransactions: seenTransactions,
	}
}
