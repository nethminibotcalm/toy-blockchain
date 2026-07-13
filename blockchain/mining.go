package blockchain

import (
	"strings"
	"time"

	"toy-blockchain/block"
)

func MineBlock(b *block.Block, difficulty int) (int, time.Duration) {
    start := time.Now()
attempts := 0
	target := strings.Repeat("0", difficulty)

	for {

		attempts++

hash := block.CalculateHash(*b)

		if strings.HasPrefix(hash, target) {

			b.Hash = hash
			return attempts, time.Since(start)
		}

		b.Nonce++
	}
}
