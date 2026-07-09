package main

import (
	"fmt"
	"os"
	"strconv"

	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
)

func main() {

	bc, err := blockchain.LoadFromFile("chain.json")

	if err != nil {
		bc = blockchain.NewBlockchain()
	}

	l := ledger.NewLedger()



	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":

		if len(os.Args) < 5 {
			fmt.Println("Usage: add <sender> <receiver> <amount>")
			return
		}

		amount, err := strconv.ParseFloat(os.Args[4], 64)

		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		tx := ledger.Transaction{
			Sender:   os.Args[2],
			Receiver: os.Args[3],
			Amount:   amount,
		}

		bc.AddTransaction(tx)

		err = bc.SaveToFile("chain.json")

		if err != nil {
			fmt.Println("Error saving blockchain:", err)
			return
		}

		fmt.Println("Transaction added")

	case "mine":

		bc.MinePendingTransactions(l)

		err = bc.SaveToFile("chain.json")

		if err != nil {
			fmt.Println("Error saving blockchain:", err)
			return
		}

		fmt.Println("Mining completed")

	case "print":

		bc.PrintChain()

	case "validate":

		fmt.Println("Blockchain valid:", bc.ValidateChain())

case "balance":
    balances := blockchain.CalculateBalances(bc.Blocks)
    fmt.Println("Balances:", balances)

	default:

		fmt.Println("Commands:")
		fmt.Println("  add <sender> <receiver> <amount>")
		fmt.Println("  mine")
		fmt.Println("  print")
		fmt.Println("  validate")
		fmt.Println("  balance")
	}
}