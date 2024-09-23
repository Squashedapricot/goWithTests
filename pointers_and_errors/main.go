package main

import (
	"errors"
	"fmt"
)

type Monero int

type Wallet struct {
	balance Monero
}

func (w *Wallet) Deposit(amount Monero) {
	w.balance += amount
}

func (w *Wallet) Balance() Monero {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) WithDraw(withdrawal Monero) error {
	if withdrawal > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= withdrawal
	return nil
}

type Stringer interface {
	String() Stringer
}

func (m Monero) String() string {
	return fmt.Sprintf("%d XMR", m)
}
