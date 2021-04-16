package walletrepo

import (
	"context"
	"errors"

	"github.com/NoobMM/golang/app/domain/entities"
	"github.com/NoobMM/golang/app/intrastructure/models"
	"github.com/jinzhu/gorm"
)


type FindOneWalletInput struct {
	ID *string
	Name *string
}

func (repo *repo) FindOneWallet(ctx context.Context, input FindOneWalletInput) (*entities.Wallet, error) {
	var wallet models.Wallet
	query := repo.DB
	if input.ID != nil {
		query = query.Where(`id = ?`,*input.ID)
	}
	result := query.First(&wallet)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, result.Error
	}
	resultEntity, err := wallet.ToEntity()
	if err != nil {
		return nil, err
	}
	return resultEntity, nil
}