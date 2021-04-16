package wallethttp

import (
	"github.com/NoobMM/golang/app/constants"
	walletusecase "github.com/NoobMM/golang/app/usecases/wallet"
	"github.com/NoobMM/golang/app/utils/respfmt"
	"github.com/NoobMM/golang/app/utils/xerrors"
	"github.com/gin-gonic/gin"
)


type UpdateWalletRequest struct {
	Name   *string `json:"name"`
	Amount *int64 `json:"amount"`
}

func (h *httpHandler) UpdateWallet(c *gin.Context) {
	var req UpdateWalletRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		respfmt.JSONErrorResponse(c, xerrors.InternalError{
			Code:    constants.StatusCodeGenericInternalError,
			Message: constants.ErrorMessageInternalError,
		}.Wrap(err))
		return
	}

	wallet, err := h.WalletUseCase.UpdateWalletUseCase(c, walletusecase.UpdateWalletUseCaseInput{Name: req.Name, Amount: req.Amount})

	if err != nil {
		respfmt.JSONErrorResponse(c, err)
		return
	}

	respfmt.JSONSuccessResponse(c, wallet)
}