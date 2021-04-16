package healthcheckhttp

import (
	"github.com/NoobMM/golang/app/utils/respfmt"
	"github.com/gin-gonic/gin"
)

func (handler *httpHandler) CheckReadiness(c *gin.Context) {
	err := handler.HealthCheckUseCase.CheckReadiness(c)
	if err != nil {
		respfmt.JSONErrorResponse(c, err)
		return
	}

	respfmt.JSONSuccessResponse(c, nil)
}
