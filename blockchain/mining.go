package blockchain
import(
	"toy-blockchain/block"
	"strings"
)
func MineBlock(b *block.Block, difficulty int) {

	target := strings.Repeat("0", difficulty)

	for {

		hash := block.CalculateHash(*b)

		if strings.HasPrefix(hash, target) {

			b.Hash = hash
			return
		}

		b.Nonce++
	}
}