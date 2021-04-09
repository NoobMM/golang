package walletusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/NoobMM/golang/app/constants"
	"github.com/NoobMM/golang/app/domain/entities"
	"github.com/NoobMM/golang/app/utils/xerrors"
	walletrepo "github.com/NoobMM/golang/app/intrastructure/repos/wallet"
)


type FindWalletUseCaseInput struct {
	ID *string
}

func (uc *useCase) FindWalletUseCase(ctx context.Context, input FindWalletUseCaseInput) (*entities.Wallet, error) {
	if input.ID == nil || *input.ID == "" {
		return nil, xerrors.ParameterError{
			Code: constants.StatusCodeMissingRequiredParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtRequired, "id"),
		}.Wrap(errors.New("ID is Required"))
	}
	findWallet, err := uc.WalletRepo.FindOneWallet(ctx, walletrepo.FindOneWalletInput{
		ID: input.ID,
	})
	if err != nil {
		return nil, xerrors.InternalError{
			Code: constants.StatusCodeDatabaseError,
			Message: constants.ErrorMessageDatabaseError,
		}.Wrap(err)
	}
	return findWallet, nil
}