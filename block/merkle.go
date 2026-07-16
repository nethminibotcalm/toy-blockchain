package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"toy-blockchain/ledger"
)

func hashTransaction(tx ledger.Transaction) string {

	data := fmt.Sprintf("%s:%s:%d", tx.Sender, tx.Receiver, tx.Amount)

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}

func CalculateMerkleRoot(
	transactions []ledger.Transaction,
) string {

	if len(transactions) == 0 {
		return ""
	}

	var hashes []string

	// Hash every transaction
	for _, tx := range transactions {

		hashes = append(
			hashes,
			hashTransaction(tx),
		)
	}

	// Build tree
	for len(hashes) > 1 {

		var newLevel []string

		for i := 0; i < len(hashes); i += 2 {

			left := hashes[i]

			right := left

			if i+1 < len(hashes) {
				right = hashes[i+1]
			}

			data := left + right

			hash := sha256.Sum256([]byte(data))

			newLevel = append(
				newLevel,
				hex.EncodeToString(hash[:]),
			)

		}

		hashes = newLevel
	}

	return hashes[0]
}
