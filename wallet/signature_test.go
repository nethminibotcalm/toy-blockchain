package wallet

import "testing"

func TestSignatureWorks(t *testing.T) {

	w, err := NewWallet()

	if err != nil {
		t.Fatal(err)
	}

	data := "Alice:Bob:20"

	signature, err := Sign(data, w.PrivateKey)

	if err != nil {
		t.Fatal(err)
	}

	valid := Verify(
		data,
		signature,
		w.GetPublicKey(),
	)

	if !valid {
		t.Fatal("signature verification failed")
	}
}