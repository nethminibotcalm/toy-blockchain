package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// calculation generated the SHA-256 hash of a block
func CalculateHash(b Block) string {
	//combine all block fields (expect hash) into one string
	data := fmt.Sprintf("Index:%d|Timestamp:%d|Transactions:%v|PreviousHash:%s|Nonce:%d", //spritf-combine the block fields into one strin
		b.Index,
		b.Timestamp,
		b.Transactions,
		b.PreviousHash,
		b.Nonce)
	//Generate SHA-256 hash(convert into bytes).
	hash := sha256.Sum256([]byte(data))
	//convert the hash bytes into a hexadecimal string.
	return hex.EncodeToString(hash[:])
}
