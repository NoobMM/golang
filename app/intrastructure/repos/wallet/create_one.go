package walletrepo

import (
	"context"

	"github.com/NoobMM/golang/app/domain/entities"
	"github.com/NoobMM/golang/app/intrastructure/models"
)

type CreateOneWalletInput struct {
	WalletEntity *entities.Wallet
}

func (repo *repo) CreateOneWallet(ctx context.Context, input CreateOneWalletInput) (*entities.Wallet, error) {
	walletModel, err := new(models.Wallet).FormEntity(input.WalletEntity)
	if err != nil {
		return nil, err
	}
	query := repo.DB
	result := query.Create(walletModel)
	if result.Error != nil {
		return nil, result.Error
	}
	resultEntity, err := walletModel.ToEntity()
	if err != nil {
		return nil, err
	}
	return resultEntity, nil
}

