package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Monero(20))

		want := Monero(20)

		checkBalanceposttransction(t, wallet, want)
	})
	t.Run("withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Monero(20)}

		wallet.WithDraw(Monero(10))

		want := Monero(10)
		checkBalanceposttransction(t, wallet, want)

	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Monero(20)
		wallet := Wallet{startingBalance}
		err := wallet.WithDraw(Monero(100))

		checktransactionalAnomaly(t, err, "cannot withdraw, insufficient funds")
		checkBalanceposttransction(t, wallet, startingBalance)
	})
}

func checkBalanceposttransction(t testing.TB, wallet Wallet, want Monero) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func checktransactionalAnomaly(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but expected one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func checkifnoerror(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("got an err but didn't want one")
	}
}
