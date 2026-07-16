package blockchain

import (
	"fmt"
	"strings"

	"toy-blockchain/block"
)

func (bc *Blockchain) ValidateChain() error {
	// Check blockchain is not empty
	if len(bc.Blocks) == 0 {
		return fmt.Errorf("blockchain is empty")
	}

	// Check genesis block
	genesis := bc.Blocks[0]

	// Genesis block must have fixed properties
	if genesis.Index != 0 {
		return fmt.Errorf("invalid genesis index")
	}

	expectedPreviousHash := "0000000000000000000000000000000000000000000000000000000000000000"

	if genesis.PreviousHash != expectedPreviousHash {
		return fmt.Errorf("invalid genesis previous hash")
	}

	// Validate genesis Merkle root
	calculatedRoot := block.CalculateMerkleRoot(genesis.Transactions)
	if calculatedRoot != genesis.MerkleRoot {
		return fmt.Errorf("block 0: invalid Merkle root")
	}

	// Validate genesis hash
	if block.CalculateHash(genesis) != genesis.Hash {
		return fmt.Errorf("block 0: invalid genesis hash")
	}

	// Check remaining blocks
	for i := 1; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		// Check Merkle root
		calculatedRoot := block.CalculateMerkleRoot(current.Transactions)
		if calculatedRoot != current.MerkleRoot {
			return fmt.Errorf("block %d: invalid Merkle root", i)
		}

		// Check block hash
		calculatedHash := block.CalculateHash(current)
		if calculatedHash != current.Hash {
			return fmt.Errorf("block %d: invalid hash", i)
		}

		// Check previous hash link
		if current.PreviousHash != previous.Hash {
			return fmt.Errorf("block %d: invalid previous hash link", i)
		}

		// Check block order
		if current.Index != previous.Index+1 {
			return fmt.Errorf("block %d: invalid index", i)
		}

		// Check timestamp order
		if current.Timestamp < previous.Timestamp {
			return fmt.Errorf("block %d: invalid timestamp order", i)
		}

		// Check proof of work
		prefix := strings.Repeat("0", current.Difficulty)
		if !strings.HasPrefix(current.Hash, prefix) {
			return fmt.Errorf("block %d: invalid proof of work", i)
		}
	}

	// Check transaction balances
	if !bc.ValidateBalances() {
		return fmt.Errorf("invalid balances: negative balance or invalid transaction")
	}

	return nil
}
