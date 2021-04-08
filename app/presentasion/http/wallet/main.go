package wallethttp

import (
	walletusecase "github.com/NoobMM/golang/app/usecases/wallet"	
	"github.com/gin-gonic/gin"
)

type HTTPHandler interface {
	CreateWallet(c *gin.Context)
}

type httpHandler struct {
	WalletUseCase walletusecase.UseCase
}

func New(walletusecase walletusecase.UseCase) HTTPHandler {
	return &httpHandler{
		WalletUseCase: walletusecase,
	}
}