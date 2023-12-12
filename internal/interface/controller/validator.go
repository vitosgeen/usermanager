package controller

import (
	"usermanager/internal/apperrors"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return apperrors.ValidatorCustomValidatorValidate.AppendMessage(err)
	}
	return nil
}
