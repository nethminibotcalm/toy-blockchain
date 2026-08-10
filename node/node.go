package node

import (
	"sync"

	"toy-blockchain/blockchain"
)

type Node struct {
	Config     Config
	Blockchain *blockchain.Blockchain
	mu         sync.RWMutex
}

func NewNode(
	config Config,
	chain *blockchain.Blockchain,
) *Node {
	return &Node{
		Config:     config,
		Blockchain: chain,
	}
}
