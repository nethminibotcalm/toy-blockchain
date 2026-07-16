package blockchain

import (
	"toy-blockchain/block"
)

func (bc *Blockchain) MineBlockConcurrent() (block.Block, error) {

	newBlock := block.Block{
		Index:        len(bc.Blocks),
		Timestamp:    int64(len(bc.Blocks)),
		Transactions: bc.PendingTransactions,
		PreviousHash: bc.Blocks[len(bc.Blocks)-1].Hash,
		Difficulty:   bc.Difficulty,
	}

	_, _ = MineBlockConcurrent(
		&newBlock,
		bc.Difficulty,
		4,
	)

	return newBlock, nil
}
