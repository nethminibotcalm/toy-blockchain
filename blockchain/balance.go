package blockchain

import "toy-blockchain/block"

func CalculateBalances(blocks []block.Block, initialBalances map[string]int) map[string]int {
	balances := make(map[string]int, len(initialBalances))

	for name, balance := range initialBalances {
		balances[name] = int(balance)
	}

	for _, b := range blocks {

		for _, tx := range b.Transactions {

			balances[tx.Sender] -= tx.Amount
			balances[tx.Receiver] += tx.Amount
		}
	}

	return balances
}
