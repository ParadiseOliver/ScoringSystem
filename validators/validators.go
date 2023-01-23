package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateIsAfter(field validator.FieldLevel) bool {
	date, ok := field.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
