package wallet

import (
	"testing"
)

func TestSignatureManyRounds(t *testing.T) {

	for i := 0; i < 3000; i++ {

		w, err := NewWallet()

		if err != nil {
			t.Fatal(err)
		}

		data := "Alice:Bob:20"

		signature, err := Sign(
			data,
			w.PrivateKey,
		)

		if err != nil {
			t.Fatal(err)
		}


		valid := Verify(
			data,
			signature,
			w.GetPublicKey(),
		)


		if !valid {
			t.Fatalf(
				"signature verification failed at iteration %d",
				i,
			)
		}
	}
}