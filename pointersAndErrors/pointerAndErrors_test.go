package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		depositAmt := Bitcoin(10)
		wallet.Deposit(depositAmt)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(15))

		withdrawAmt := Bitcoin(5)
		err := wallet.Withdraw(withdrawAmt)

		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBal := Bitcoin(20)
		wallet := Wallet{balance: startingBal}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBal)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("did not get an error but wanted one!")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("didn't want an error but got one")
	}
}
