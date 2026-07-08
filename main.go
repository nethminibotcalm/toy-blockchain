package main

import (
	"fmt"
	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
	
)

func main() {

	bc := blockchain.NewBlockchain()

bc.AddBlock([]ledger.Transaction{
	{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   20,
	},
})

bc.AddBlock([]ledger.Transaction{
	{
		Sender:   "Bob",
		Receiver: "Charlie",
		Amount:   10,
	},
})
		// Uncomment this line to simulate a hacker changing Block 1
	 //bc.Blocks[1].Transactions[0].Amount = 9999
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