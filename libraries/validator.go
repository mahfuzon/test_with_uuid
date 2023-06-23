package libraries

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.Validator.RegisterValidation("genderValidation", genderValidation)
	cv.Validator.RegisterValidation("phoneValidation", phoneValidation)
	err := cv.Validator.Struct(i)
	if err != nil {
		return err
	}
	return nil
}

func genderValidation(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value == "male" || value == "female" {
			return true
		}
	}

	return false
}

func phoneValidation(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		regex := `^\+62\d{9,15}$`
		match, err := regexp.MatchString(regex, value)
		if err != nil {
			panic(err.Error())
		}

		if match {
			return true
		}
	}

	return false
}
