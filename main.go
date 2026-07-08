package main

import (
	"fmt"
	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
	
)

func main() {

    l := ledger.NewLedger()

    bc := blockchain.NewBlockchain()

    bc.AddTransaction(ledger.Transaction{
        Sender:   "Alice",
        Receiver: "Bob",
        Amount:   20,
    })

    bc.MinePendingTransactions(l)


    bc.AddTransaction(ledger.Transaction{
        Sender:   "Bob",
        Receiver: "Charlie",
        Amount:   10,
    })

    bc.MinePendingTransactions(l)


    fmt.Println("Balances:", l.Balances)

    fmt.Println("Blockchain valid:", bc.ValidateChain())
    fmt.Println("Number of blocks:", len(bc.Blocks))

    for _, block := range bc.Blocks {
        fmt.Println("----------------------------")
        fmt.Println("Index:", block.Index)
        fmt.Println("Transactions:", block.Transactions)
        fmt.Println("Nonce:", block.Nonce)
        fmt.Println("Previous Hash:", block.PreviousHash)
        fmt.Println("Hash:", block.Hash)
    }
}