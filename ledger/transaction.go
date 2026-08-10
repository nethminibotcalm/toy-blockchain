package ledger

type Transaction struct {
	Sender        string
	SenderAddress string
	Receiver      string
	Amount        int
	Nonce         int
	Signature     string
	PublicKey     string
}