package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Init() {
	Validate = validator.New()

	Validate.RegisterValidation("cyrillic", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		matched, _ := regexp.MatchString(`^[а-яА-ЯёЁ\s]+$`, value)
		return matched
	})

	Validate.RegisterValidation("latin", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		matched, _ := regexp.MatchString(`^[a-zA-Z\s]+$`, value)
		return matched
	})
}
