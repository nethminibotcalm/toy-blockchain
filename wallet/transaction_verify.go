package wallet

import (
	"fmt"

	"toy-blockchain/ledger"
)

func VerifyTransaction(tx ledger.Transaction) bool {

	data := fmt.Sprintf(
		"%s:%s:%d",
		tx.Sender,
		tx.Receiver,
		tx.Amount,
	)

	return Verify(
		data,
		tx.Signature,
		tx.PublicKey,
	)
}
