package healthcheckhttp

import (
	"github.com/NoobMM/golang/app/utils/respfmt"
	"github.com/gin-gonic/gin"
)

func (handler *httpHandler) CheckLiveness(c *gin.Context) {
	respfmt.JSONSuccessResponse(c, nil)
}
