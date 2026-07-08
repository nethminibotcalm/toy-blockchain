package blockchain

import (
	"time"
	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

type Blockchain struct {
	Blocks []block.Block
}
func CreateGenesisBlock() block.Block{
	genesis:=block.Block{
		Index: 0,
		Timestamp: time.Now().Unix(),
		Transactions: []ledger.Transaction{},
		PreviousHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Nonce: 0,
	
	
	}
	//Calculate and assign the hash.
	genesis.Hash=block.CalculateHash(genesis)
	return genesis
}
// NewBlockchain creates a blockchain with the Genesis Block.
func NewBlockchain() *Blockchain {

	genesis := CreateGenesisBlock()

	return &Blockchain{
		Blocks: []block.Block{genesis},
	}
}
func (bc *Blockchain) AddBlock(transactions []ledger.Transaction) {

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := block.Block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PreviousHash: lastBlock.Hash,
		Nonce:        0,
	}

	MineBlock(&newBlock, 4)

	bc.Blocks = append(bc.Blocks, newBlock)
}