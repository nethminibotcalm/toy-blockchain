package ledger

type Ledger struct {
	Balances map[string]int
}

func NewLedger(balances map[string]int) *Ledger {
	return &Ledger{
		Balances: balances,
	}
}
func (l *Ledger) ApplyTransaction(t Transaction) {

	l.Balances[t.Sender] -= t.Amount
	l.Balances[t.Receiver] += t.Amount

}
func (l *Ledger) ValidateTransaction(t Transaction) bool {

	if t.Amount <= 0 {
		return false
	}

	senderBalance, exists := l.Balances[t.Sender]

	if !exists {
		return false
	}

	if senderBalance < t.Amount {
		return false
	}

	return true
}
