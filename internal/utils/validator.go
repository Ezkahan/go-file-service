package utils

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type ValidationErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidationError(err error) interface{} {
	var errors []*ValidationErr
	for _, err := range err.(validator.ValidationErrors) {
		var (
			item  ValidationErr
			field []rune = []rune(err.Field())
		)

		switch err.Tag() {
		case "customDate":
			return fmt.Sprintf("Invalid date format for '%s'. Expected format: yyyy-mm-dd.", err.Field())
		}

		field[0] = unicode.ToLower(field[0])
		item.Field = string(field)
		item.Message = err.Field() + " " + err.Tag()
		errors = append(errors, &item)
	}
	return errors
}
