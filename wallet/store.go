package wallet

var Wallets = make(map[string]*Wallet)

func CreateWallet(name string) error {

	w, err := NewWallet()

	if err != nil {
		return err
	}

	Wallets[name] = w

	return SaveWallet(name, w)

}

func GetWallet(name string) (*Wallet, bool) {

	if w, ok := Wallets[name]; ok {
		return w, true
	}

	w, err := LoadWallet(name)
	if err != nil {
		return nil, false
	}

	Wallets[name] = w

	return w, true
}
