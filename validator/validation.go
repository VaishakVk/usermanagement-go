package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

// Init Func
func Init() {
	validate = validator.New()
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		// r, _ := regexp.Compile("(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}")
		// fmt.Println("Here", r)
		matched, _ := regexp.MatchString("(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}", "Hello123")
		fmt.Println("Here")
		return matched
	})
}

// Validate function
func Validate(payload interface{}) error {
	err := validate.Struct(payload)
	return err
}
