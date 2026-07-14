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
		if _, statErr := os.Stat("chain.json"); statErr == nil {
			fmt.Println("Warning: chain.json exists but could not be loaded:", err)
			fmt.Println("Starting a new blockchain instead.")
		}
		bc = blockchain.NewBlockchain()
	}

	balances := blockchain.CalculateBalances(bc.Blocks, bc.InitialBalances)

	l := ledger.NewLedger(balances)

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

		amount, err := strconv.Atoi(os.Args[4])

		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		tx := ledger.Transaction{
			Sender:   os.Args[2],
			Receiver: os.Args[3],
			Amount:   amount,
		}

		if !bc.AddTransaction(tx) {
			fmt.Println("Transaction rejected")
			return
		}

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

		err := bc.ValidateChain()

		if err != nil {
			fmt.Println("Blockchain is invalid:", err)
		} else {
			fmt.Println("Blockchain is valid")
		}

	case "balance":
		balances := blockchain.CalculateBalances(bc.Blocks, bc.InitialBalances)
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
