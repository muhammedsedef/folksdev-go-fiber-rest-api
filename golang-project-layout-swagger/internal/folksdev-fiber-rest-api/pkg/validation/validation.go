package validation

import "github.com/go-playground/validator/v10"

// custom validation error struct
type CustomValidationError struct {
	HasError bool
	Field    string
	Tag      string
	Param    string
	Value    interface{}
}

type ICustomValidator interface {
	Validate(data interface{}) []CustomValidationError
}

type customValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) ICustomValidator {
	return &customValidator{validator: validator}
}

func (cv *customValidator) Validate(data interface{}) []CustomValidationError {
	var customValidationErrors []CustomValidationError

	if errors := cv.validator.Struct(data); errors != nil {
		for _, fieldError := range errors.(validator.ValidationErrors) {
			customValidationErrors = append(customValidationErrors, CustomValidationError{
				HasError: true,
				Field:    fieldError.Field(),
				Tag:      fieldError.Tag(),
				Param:    fieldError.Param(),
				Value:    fieldError.Value(),
			})
		}
	}
	return customValidationErrors
}
