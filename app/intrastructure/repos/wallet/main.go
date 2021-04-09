package walletrepo

import (
	"context"

	"github.com/NoobMM/golang/app/domain/entities"
	"gorm.io/gorm"
)


type Repo interface {
	CreateOneWallet(ctx context.Context, input CreateOneWalletInput) (*entities.Wallet, error)
	FindOneWallet(ctx context.Context, input FindOneWalletInput) (*entities.Wallet, error)
}

type repo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Repo {
	return &repo{
		DB: db,
	}
}