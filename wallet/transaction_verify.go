package wallet

import (
	"fmt"

	"toy-blockchain/ledger"
)

// VerifyTransaction checks two things, not just one:
//  1. That the signature is cryptographically valid for the public key
//     attached to the transaction (as before).
//  2. That the attached public key actually belongs to the wallet
//     registered under the claimed sender's name.
//
// Without check 2, anyone could generate their own key pair, sign a
// transaction with Sender set to someone else's name, and have it accepted
// -- a valid signature only proves "someone" authorized the transaction,
// not that the claimed sender did. Binding the key to the registered
// identity is what makes the signature actually mean something.
func VerifyTransaction(tx ledger.Transaction) bool {

	senderWallet, exists := GetWallet(tx.Sender)
	if !exists {
		return false
	}

	if tx.PublicKey != senderWallet.GetPublicKey() {
		return false
	}

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
