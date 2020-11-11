package xerrors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ParameterError Invalid parameter
type ParameterError struct {
	Code            uint
	Message         string
	ValidatorErrors *validator.ValidationErrors
}

// Wrap ...
func (e ParameterError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e ParameterError) Error() string {
	return e.Message
}

// Is ...
func (e ParameterError) Is(target error) bool {
	_, ok := target.(*ParameterError)
	if !ok {
		return false
	}
	return true
}

// InsufficientError ...
type InsufficientError struct {
	Code            uint
	Message         string
	ValidatorErrors *validator.ValidationErrors
}

// Wrap ...
func (e InsufficientError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e InsufficientError) Error() string {
	return e.Message
}

// Is ...
func (e InsufficientError) Is(target error) bool {
	_, ok := target.(*ParameterError)
	if !ok {
		return false
	}
	return true
}

// AuthError error and etc.
type AuthError struct {
	Code    uint
	Message string
}

// Wrap ...
func (e AuthError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e AuthError) Error() string {
	return e.Message
}

// Is ...
func (e AuthError) Is(target error) bool {
	_, ok := target.(*InternalError)
	if !ok {
		return false
	}
	return true
}

// UnprocessableEntity Valid parameter but invalid business and etc.
type UnprocessableEntity struct {
	Code    uint
	Message string
}

// Wrap ...
func (e UnprocessableEntity) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e UnprocessableEntity) Error() string {
	return e.Message
}

// Is ...
func (e UnprocessableEntity) Is(target error) bool {
	_, ok := target.(*UnprocessableEntity)
	if !ok {
		return false
	}
	return true
}

// Forbidden Valid parameter but invalid business and etc.
type Forbidden struct {
	Code    uint
	Message string
}

// Wrap ...
func (e Forbidden) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e Forbidden) Error() string {
	return e.Message
}

// Is ...
func (e Forbidden) Is(target error) bool {
	_, ok := target.(*Forbidden)
	if !ok {
		return false
	}
	return true
}

// RecordNotFoundError Cannot find resource.
type RecordNotFoundError struct {
	Code    uint
	Message string
}

// Wrap ...
func (e RecordNotFoundError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e RecordNotFoundError) Error() string {
	return e.Message
}

// Is ...
func (e RecordNotFoundError) Is(target error) bool {
	_, ok := target.(*RecordNotFoundError)
	if !ok {
		return false
	}
	return true
}

// InternalError Database error and etc.
type InternalError struct {
	Code    uint
	Message string
}

// Wrap ...
func (e InternalError) Wrap(err error) error {
	return fmt.Errorf("%w: %v", e, err)
}

// Error ...
func (e InternalError) Error() string {
	return e.Message
}

// Is ...
func (e InternalError) Is(target error) bool {
	_, ok := target.(*InternalError)
	if !ok {
		return false
	}
	return true
}
