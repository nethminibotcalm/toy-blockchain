package ledger

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    int
	Nonce     int
	Signature string
	PublicKey string
}
