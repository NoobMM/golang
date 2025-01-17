package entities

import (
	"github.com/AlekSi/pointer"
)

type Wallet struct {
	ID      *string
	Name    *string
	Balance *int64
}

func NewWallet(name *string) *Wallet {
	return &Wallet{
		Name:    name,
		Balance: pointer.ToInt64(0),
	}
}

func (w *Wallet) SetBalance(balance *int64) *Wallet {
	w.Balance = balance
	return w
}

func (w *Wallet) AddBalance(amount *int64) *Wallet {
	newBalance := *w.Balance + *amount
	w.Balance = &newBalance
	return w
}
