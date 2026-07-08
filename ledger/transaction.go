package ledger

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}