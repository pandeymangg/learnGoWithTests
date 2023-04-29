package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amt Bitcoin) {

	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
