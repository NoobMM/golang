package respfmt

import (
	"errors"
	"net/http"
	"strings"

	"github.com/deuanz/golang-with-heroku/app/constants"
	"github.com/deuanz/golang-with-heroku/app/utils/xerrors"
	"github.com/gin-gonic/gin"
)

// Message constant
const (
	MessageOK string = "OK"
)

// BaseStatus ...
type BaseStatus struct {
	Code     uint     `json:"code"`
	Messages []string `json:"messages"`
}

type baseSuccessResponse struct {
	Status BaseStatus  `json:"status"`
	Data   interface{} `json:"data"`
}

// JSONSuccessResponse ...
func JSONSuccessResponse(c *gin.Context, data interface{}) {
	r := new(baseSuccessResponse)
	r.Status.Code = constants.StatusCodeSuccess
	r.Status.Messages = append(r.Status.Messages, MessageOK)
	r.Data = data
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

type baseSuccessResponseWithPagination struct {
	Status     BaseStatus  `json:"status"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// JSONSuccessResponseWithPagination is a function format the response to a standard format
func JSONSuccessResponseWithPagination(c *gin.Context, pagination *Pagination, data interface{}) {
	r := new(baseSuccessResponseWithPagination)
	r.Status.Code = constants.StatusCodeSuccess
	r.Status.Messages = append(r.Status.Messages, MessageOK)
	r.Pagination = *pagination
	r.Data = data
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

type baseErrorResponse struct {
	Status BaseStatus `json:"status"`
}

// JSONErrorResponse ...
func JSONErrorResponse(c *gin.Context, err error) {
	status := BaseStatus{
		Code:     constants.StatusCodeGenericInternalError,
		Messages: []string{},
	}
	httpStatusCode := http.StatusInternalServerError

	unwrappedErr := errors.Unwrap(err)
	switch unwrappedErr.(type) {
	case xerrors.InternalError:
		xerr, _ := unwrappedErr.(xerrors.InternalError)
		httpStatusCode = http.StatusInternalServerError
		status.Code = xerr.Code
		status.Messages = append(status.Messages, xerr.Message)
	case xerrors.ParameterError:
		xerr, _ := unwrappedErr.(xerrors.ParameterError)
		httpStatusCode = http.StatusBadRequest
		status.Code = xerr.Code
		if xerr.Message != "" {
			status.Messages = append(status.Messages, xerr.Message)
		}
		if xerr.ValidatorErrors != nil {
			status.Code = xerrors.FormatValidationCode((*xerr.ValidatorErrors)[0])
			for _, err := range *xerr.ValidatorErrors {
				status.Messages = append(status.Messages, xerrors.FormatValidationError(err))
			}
		}
	case xerrors.RecordNotFoundError:
		xerr, _ := unwrappedErr.(xerrors.RecordNotFoundError)
		httpStatusCode = http.StatusNotFound
		status.Code = xerr.Code
		status.Messages = append(status.Messages, xerr.Message)
	case xerrors.AuthError:
		xerr, _ := unwrappedErr.(xerrors.AuthError)
		httpStatusCode = http.StatusUnauthorized
		status.Code = xerr.Code
		status.Messages = append([]string{}, xerr.Message)
	case xerrors.UnprocessableEntity:
		xerr, _ := unwrappedErr.(xerrors.UnprocessableEntity)
		httpStatusCode = http.StatusUnprocessableEntity
		status.Code = xerr.Code
		status.Messages = append(status.Messages, xerr.Message)
	case xerrors.Forbidden:
		xerr, _ := unwrappedErr.(xerrors.Forbidden)
		httpStatusCode = http.StatusForbidden
		status.Code = xerr.Code
		status.Messages = append(status.Messages, xerr.Message)
	default:
		status.Messages = append(status.Messages, err.Error())
	}

	errorResponse := baseErrorResponse{Status: status}

	c.AbortWithStatusJSON(httpStatusCode, errorResponse)
}

// ErrorResponseWithLogOnFullError Will use message as an error message and log the error stack traces
func ErrorResponseWithLogOnFullError(err error) string {
	var messages []string
	unwrappedErr := errors.Unwrap(err)
	switch unwrappedErr.(type) {
	case xerrors.InternalError:
		xerr, _ := unwrappedErr.(xerrors.InternalError)
		messages = append(messages, xerr.Message)
	case xerrors.ParameterError:
		xerr, _ := unwrappedErr.(xerrors.ParameterError)
		messages = append(messages, xerr.Message)
		if xerr.ValidatorErrors != nil {
			for _, err := range *xerr.ValidatorErrors {
				messages = append(messages, xerrors.FormatValidationError(err))
			}
		}
	case xerrors.RecordNotFoundError:
		xerr, _ := unwrappedErr.(xerrors.RecordNotFoundError)
		messages = append(messages, xerr.Message)
	case xerrors.AuthError:
		xerr, _ := unwrappedErr.(xerrors.AuthError)
		messages = append([]string{}, xerr.Message)
	case xerrors.UnprocessableEntity:
		xerr, _ := unwrappedErr.(xerrors.UnprocessableEntity)
		messages = append(messages, xerr.Message)
	default:
		messages = append(messages, err.Error())
	}

	return strings.Join(messages, ", ")
}
