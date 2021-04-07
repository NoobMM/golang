package walletusecase

import (
	"github.com/NoobMM/golang/app/domain/entities"
	walletrepo "github.com/NoobMM/golang/app/intrastructure/repos/wallet")


type UseCase interface {
	CreateWallet(ctx context.Context, input CreateWalletUseCaseInput) (*entities.Wallet, error)
}

type useCase struct {
	WalletRepo walletrepo.Repo
}

func New(walletrepo walletrepo.Repo) UseCase {
	return &useCase{
		WalletRepo: walletrepo,
	}
}