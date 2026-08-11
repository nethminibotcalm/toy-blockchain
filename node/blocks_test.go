package node

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
	"toy-blockchain/wallet"
)

func TestMineAndGossipBlockBetweenTwoNodes(t *testing.T) {
	chainB := blockchain.NewBlockchain()

	nodeB := NewNode(
		Config{
			Address: "node-b",
			Peers:   []string{"node-a"},
		},
		chainB,
	)

	serverB := httptest.NewServer(nodeB.Handler())
	defer serverB.Close()

	chainA := blockchain.NewBlockchain()

	nodeA := NewNode(
		Config{
			Address: "node-a",
			Peers:   []string{serverB.URL},
		},
		chainA,
	)
	w, err := wallet.NewWallet()

	if err != nil {
		t.Fatal(err)
	}

	tx, err := wallet.SignTransaction(
		ledger.Transaction{
			Sender:   "Alice",
			Receiver: "Bob",
			Amount:   20,
			Nonce:    1,
		},
		w,
	)

	if err != nil {
		t.Fatal(err)
	}

	transactionJSON, err := json.Marshal(tx)

	if err != nil {
		t.Fatal(err)
	}

	transactionRequest := httptest.NewRequest(
		http.MethodPost,
		"/transactions",
		bytes.NewReader(transactionJSON),
	)

	transactionResponse := httptest.NewRecorder()

	nodeA.Handler().ServeHTTP(
		transactionResponse,
		transactionRequest,
	)

	if transactionResponse.Code != http.StatusAccepted {
		t.Fatalf(
			"expected transaction status 202, got %d: %s",
			transactionResponse.Code,
			transactionResponse.Body.String(),
		)
	}
	mineRequest := httptest.NewRequest(
		http.MethodPost,
		"/mine",
		nil,
	)

	mineResponse := httptest.NewRecorder()

	nodeA.Handler().ServeHTTP(
		mineResponse,
		mineRequest,
	)

	if mineResponse.Code != http.StatusCreated {
		t.Fatalf(
			"expected mining status 201, got %d: %s",
			mineResponse.Code,
			mineResponse.Body.String(),
		)
	}
	if len(chainA.Blocks) != 2 {
		t.Fatalf(
			"expected Node A to have 2 blocks, got %d",
			len(chainA.Blocks),
		)
	}

	if len(chainB.Blocks) != 2 {
		t.Fatalf(
			"expected Node B to have 2 blocks, got %d",
			len(chainB.Blocks),
		)
	}

	if chainA.Blocks[1].Hash != chainB.Blocks[1].Hash {
		t.Fatal("expected both nodes to have the same block")
	}

	if len(chainA.PendingTransactions) != 0 {
		t.Fatal("expected Node A pending pool to be empty")
	}

	if len(chainB.PendingTransactions) != 0 {
		t.Fatal("expected Node B pending pool to be empty")
	}
}
