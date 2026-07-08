package blockchain
import "toy-blockchain/block"

func (bc *Blockchain) ValidateChain() bool {
	for i:=1;i<len(bc.Blocks);i++{
		current :=bc.Blocks[i]
		previous := bc.Blocks[i-1]
		calculatedHash := block.CalculateHash(current)
		if calculatedHash != current.Hash {
			return false
		}
		// Check if blocks are properly connected
		if current.PreviousHash != previous.Hash {
			return false
		}
		}
		return true
		
	}
	