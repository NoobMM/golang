package wallethttp

import (
	"errors"
	"fmt"

	"github.com/NoobMM/golang/app/constants"
	walletusecase "github.com/NoobMM/golang/app/usecases/wallet"
	"github.com/NoobMM/golang/app/utils/respfmt"
	"github.com/NoobMM/golang/app/utils/xerrors"
	"github.com/gin-gonic/gin"
)

func (h *httpHandler) FindWallet(c *gin.Context) {
	walletID := c.Param("walletID")
	if walletID == "" {
		respfmt.JSONErrorResponse(c, xerrors.ParameterError{
			Code: constants.StatusCodeMissingRequiredParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtRequired, "walletID"),
		}.Wrap(errors.New("Missing Wallet ID")))
		return
	}

	wallet, err := h.WalletUseCase.FindWalletUseCase(c, walletusecase.FindWalletUseCaseInput{ID: &walletID})

	if err != nil {
		respfmt.JSONErrorResponse(c, err)
		return
	}

	respfmt.JSONSuccessResponse(c, wallet)
}