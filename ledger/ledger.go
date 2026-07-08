package ledger

type Ledger struct{
	Balances map[string]float64
}
func NewLedger() *Ledger {
	return &Ledger{
		Balances: make(map[string]float64),
	}
}
func (l *Ledger) ApplyTransaction(t Transaction) bool {

    if t.Amount <= 0 {
        return false
    }

    senderBalance := l.Balances[t.Sender]

    if senderBalance < t.Amount {
        return false
    }

    l.Balances[t.Sender] -= t.Amount
    l.Balances[t.Receiver] += t.Amount

    return true
}