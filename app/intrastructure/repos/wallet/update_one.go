package walletrepo

import (
	"context"

	"github.com/NoobMM/golang/app/domain/entities"
	"github.com/NoobMM/golang/app/intrastructure/models"
)


type UpdateOneWalletInput struct {
	WalletEntity *entities.Wallet
}

func (repo *repo) UpdateOneWallet(ctx context.Context, input UpdateOneWalletInput) (*entities.Wallet, error) {
	walletModel, err := new(models.Wallet).FromEntity(input.WalletEntity)
	if err != nil {
		return nil, err
		}
	query := repo.DB
	result := query.Model(&models.Wallet{}).Where("id = ?", walletModel.ID).Select("balance").Updates(walletModel)
	if result.Error != nil {
		return nil, result.Error
		}
		resultEntity, err := walletModel.ToEntity()
	if err != nil {
		return nil, err
		}
	return resultEntity, nil
}