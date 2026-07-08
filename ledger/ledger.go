package ledger

type Ledger struct{
	Balances map[string]float64
}
func NewLedger() *Ledger {
    return &Ledger{
        Balances: map[string]float64{
            "Alice": 100,
            "Bob": 100,
            "Charlie": 100,
        },
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
