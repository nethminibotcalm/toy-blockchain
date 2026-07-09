package blockchain

import "toy-blockchain/block"

func CalculateBalances(blocks []block.Block) map[string]float64 {

	balances := map[string]float64{
		"Alice":   100,
		"Bob":     100,
		"Charlie": 100,
	}

	for _, b := range blocks {

		for _, tx := range b.Transactions {

			balances[tx.Sender] -= tx.Amount
			balances[tx.Receiver] += tx.Amount
		}
	}

	return balances
}