package blockchain

func (bc *Blockchain) ValidateBalances() bool {

	balances := make(map[string]int)

	for name, amount := range bc.InitialBalances {
		balances[name] = amount
	}

	for _, b := range bc.Blocks {

		for _, tx := range b.Transactions {

			if tx.Amount <= 0 {
				return false
			}

			senderBalance, exists := balances[tx.Sender]

			if !exists {
				return false
			}

			if senderBalance < tx.Amount {
				return false
			}

			balances[tx.Sender] -= tx.Amount
			balances[tx.Receiver] += tx.Amount
		}
	}

	return true
}
