package wallethttp

import (
	"github.com/NoobMM/golang/app/constants"
	walletusecase "github.com/NoobMM/golang/app/usecases/wallet"
	"github.com/NoobMM/golang/app/utils/respfmt"
	"github.com/NoobMM/golang/app/utils/xerrors"
	"github.com/gin-gonic/gin"
)


type CreateWalletRequest struct {
	Name *string `json:"name"`
}

func (h *httpHandler) CreateWallet(c *gin.Context) {
	var req CreateWalletRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		respfmt.JSONErrorResponse(c, xerrors.InternalError{
			Code: constants.StatusCodeGenericInternalError,
			Message: constants.ErrorMessageInternalError,
		}.Wrap(err))
		return
	}

	wallet, err := h.WalletUseCase.CreateWalletUseCase(c, walletusecase.CreateWalletUseCaseInput{Name: req.Name})

	if err != nil {
		respfmt.JSONErrorResponse(c, err)
		return
	}

	respfmt.JSONSuccessResponse(c, wallet)
}