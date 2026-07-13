package blockchain

import (
"toy-blockchain/block"
"strings"
)

func (bc *Blockchain) ValidateChain() bool {
	genesis := bc.Blocks[0]

if block.CalculateHash(genesis) != genesis.Hash {
    return false
}
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]
		calculatedHash := block.CalculateHash(current)
		if calculatedHash != current.Hash {
			return false
		}
		difficulty := 4
prefix := strings.Repeat("0", difficulty)

if !strings.HasPrefix(current.Hash, prefix) {
	return false
}
		// Check if blocks are properly connected
		if current.PreviousHash != previous.Hash {
			return false
		}
		if current.Timestamp < previous.Timestamp {
	return false
}
if current.Index != previous.Index+1 {
	return false
}
	}
	if !bc.ValidateBalances() {
	return false
}
	
	return true

}
