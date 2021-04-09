package walletusecase

import (
	"context"

	"github.com/NoobMM/golang/app/domain/entities"
	walletrepo "github.com/NoobMM/golang/app/intrastructure/repos/wallet"
)


type UseCase interface {
	CreateWalletUseCase(ctx context.Context, input CreateWalletUseCaseInput) (*entities.Wallet, error)
	FindWalletUseCase(ctx context.Context, input FindWalletUseCaseInput) (*entities.Wallet, error)
}

type useCase struct {
	WalletRepo walletrepo.Repo
}

func New(walletrepo walletrepo.Repo) UseCase {
	return &useCase{
		WalletRepo: walletrepo,
	}
}