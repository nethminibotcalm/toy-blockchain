package blockchain

import (
	"fmt"
	"testing"

	"toy-blockchain/ledger"
	"toy-blockchain/wallet"
)

var testWallets = make(map[string]*wallet.Wallet)


func createSignedTransaction(
	t *testing.T,
	sender string,
	receiver string,
	amount int,
) ledger.Transaction {


	if testWallets[sender] == nil {

		w, err := wallet.NewWallet()

		if err != nil {
			t.Fatal(err)
		}

		testWallets[sender] = w
	}


	w := testWallets[sender]


	tx := ledger.Transaction{
		Sender: sender,
		Receiver: receiver,
		Amount: amount,
	}


	data := fmt.Sprintf(
		"%s:%s:%d",
		tx.Sender,
		tx.Receiver,
		tx.Amount,
	)


	signature, err := wallet.Sign(
		data,
		w.PrivateKey,
	)


	if err != nil {
		t.Fatal(err)
	}


	tx.PublicKey = w.GetPublicKey()
	tx.Signature = signature


	return tx
}