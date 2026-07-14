package wallet

import (
	"fmt"

	"toy-blockchain/ledger"
)

func SignTransaction(
	tx ledger.Transaction,
	w *Wallet,
) (ledger.Transaction, error) {

	data := fmt.Sprintf(
		"%s:%s:%d",
		tx.Sender,
		tx.Receiver,
		tx.Amount,
	)

	signature, err := Sign(
		data,
		w.PrivateKey,
	)

	if err != nil {
		return tx, err
	}

	tx.PublicKey = w.GetPublicKey()
	tx.Signature = signature

	return tx, nil
}
