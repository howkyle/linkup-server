package validation

import "github.com/go-playground/validator"

type Validator interface {
	ValidateStruct(s interface{}) error
}

type inputValidator struct {
	validator *validator.Validate
}

func (i inputValidator) ValidateStruct(s interface{}) error {
	return i.validator.Struct(s)
}

func NewValidator() Validator {
	return inputValidator{validator.New()}
}
