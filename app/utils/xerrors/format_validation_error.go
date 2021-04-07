package xerrors

import (
	"fmt"
	"strings"

	"github.com/NoobMM/golang/app/constants"
	"github.com/go-playground/validator/v10"
	"github.com/stoewer/go-strcase"
)

// FormatValidationError ...
func FormatValidationError(e validator.FieldError) string {
	field := e.Field()
	field = strcase.LowerCamelCase(field)

	switch e.Tag() {
	case "required":
		return fmt.Sprintf(constants.ErrorMessageFmtRequired, field)
	case "notblank":
		return fmt.Sprintf(constants.ErrorMessageFmtRequired, field)
	default:
		return fmt.Sprintf(constants.ErrorMessageFmtInvalidFormat, field)
	}
}

// FormatValidationCode is a util function for mapping validation tag to an internal status code
func FormatValidationCode(e validator.FieldError) uint {
	field := e.Field()
	field = strcase.LowerCamelCase(field)

	switch e.Tag() {
	case "required":
		return constants.StatusCodeMissingRequiredParameters
	case "notblank":
		return constants.StatusCodeMissingRequiredParameters
	default:
		return constants.StatusCodeInvalidParameters
	}
}

// FormatInvalidParameter ...
func FormatInvalidParameter(e validator.ValidationErrors) string {
	var errFields []string
	for _, fe := range e {
		errFields = append(errFields, strcase.LowerCamelCase(fe.Field()))
	}
	return fmt.Sprintf(constants.ErrorMessageFmtInvalidFormat, strings.Join(errFields, ", "))
}
