package models

import (
	"github.com/NoobMM/golang/app/domain/entities"
	"github.com/jinzhu/copier"
)

type Wallet struct {
	ID      string `gorm:"primaryKey;default:uuid_generate_v4()"`
	Name    string
	Balance int64
}

func (w *Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) FromEntity(e *entities.Wallet) (*Wallet, error) {
	var model Wallet
	err := copier.Copy(&model, &e)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (w *Wallet) ToEntity() (*entities.Wallet, error) {
	var e entities.Wallet
	err := copier.Copy(&e, &w)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
