package validation

import (
	"github.com/arfan21/getprint-user/app/model/modeluser"
	validator "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func Validate(user modeluser.User) error {
	return validator.Errors{
		"name":     validator.Validate(user.Name, validator.Required),
		"email":    validator.Validate(user.Email, validator.Required, is.Email),
		"password": validator.Validate(user.Password.String, validator.When(user.Password.Valid, validator.Length(8, 20))),
	}.Filter()
}

func ValidateLogin(user modeluser.User) error {
	return validator.Errors{
		"email":    validator.Validate(user.Email, validator.Required, is.Email),
		"password": validator.Validate(user.Password.String, validator.When(user.Password.Valid, validator.Length(8, 20))),
	}.Filter()
}
