package walletusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/NoobMM/golang/app/constants"
	"github.com/NoobMM/golang/app/domain/entities"
	walletrepo "github.com/NoobMM/golang/app/intrastructure/repos/wallet"
	"github.com/NoobMM/golang/app/utils/xerrors"
)


type UpdateWalletUseCaseInput struct {
	Name *string
	Amount *int64
}

func (uc *useCase) UpdateWalletUseCase(ctx context.Context, input UpdateWalletUseCaseInput) (*entities.Wallet, error) {
	if input.Name == nil || *input.Name == "" {
		return nil, xerrors.ParameterError{
			Code:    constants.StatusCodeMissingRequiredParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtRequired, "name"),
		}.Wrap(errors.New("name is required"))
	}
	if input.Amount == nil || *input.Amount == 0 {
		return nil, xerrors.ParameterError{
			Code:    constants.StatusCodeMissingRequiredParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtRequired, "amount"),
		}.Wrap(errors.New("amount is required"))
	}
	foundWallet, err := uc.WalletRepo.FindOneWallet(ctx, walletrepo.FindOneWalletInput{
		Name: input.Name,
	})
	if err != nil {
		return nil, xerrors.InternalError{
			Code: constants.StatusCodeDatabaseError,
			Message: constants.ErrorMessageDatabaseError,
		}.Wrap(err)
	}
	foundWallet = foundWallet.AddBalance(input.Amount)
	updated, err := uc.WalletRepo.UpdateOneWallet(ctx, walletrepo.UpdateOneWalletInput{
		WalletEntity: foundWallet,
	})
	if err != nil {
		return nil, xerrors.InternalError{
			Code:    constants.StatusCodeDatabaseError,
			Message: constants.ErrorMessageDatabaseError,
		}.Wrap(err)
	}
	return updated, nil
}