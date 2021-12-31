package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got.String(), want.String())
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.WithDraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Fatal(want)
		}
	}

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.WithDraw(Bitcoin(100))

		assertBalance(t, wallet.balance, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
