package walletusecase

import (
	"context"
	"fmt"

	"github.com/NoobMM/golang/app/constants"
	"github.com/NoobMM/golang/app/domain/entities"
	walletrepo "github.com/NoobMM/golang/app/intrastructure/repos/wallet"
	"github.com/NoobMM/golang/app/utils/xerrors"
)


type CreateWalletUseCaseInput struct {
	Name *string
}

func (uc *useCase) CreateWalletUseCase(ctx context.Context, input CreateWalletUseCaseInput) (*entities.Wallet, error) {
	if input.Name == nil || *input.Name == "" {
		return nil, xerrors.ParameterError{
			Code: constants.StatusCodeMissingRequiredParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtRequired, "name"),
		}
	}
	newWallet := entities.NewWallet(input.Name)
	created, err := uc.WalletRepo.CreateOneWallet(ctx, walletrepo.CreateOneWalletInput{
		WalletEntity: newWallet,
	})
	if err != nil {
		return nil, xerrors.InternalError{
			Code: constants.StatusCodeDatabaseError,
			Message: constants.ErrorMessageDatabaseError,
		}
	}
	return created, nil
}