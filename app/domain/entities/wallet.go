package entities

import "github.com/AlekSi/pointer"

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