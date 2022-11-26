package validate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func Email(email string) error {
	return validation.Validate(email, validation.Required, is.EmailFormat)
}
