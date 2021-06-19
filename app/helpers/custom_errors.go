package helpers

import (
	"strings"

	"github.com/arfan21/getprint-user/app/constants"
)

func CustomErrors(err error) error {
	if strings.Contains(err.Error(), "users.email") {
		return constants.ErrEmailConflict
	}

	if strings.Contains(err.Error(), "hashedPassword is not the hash") {
		return constants.ErrPasswordNotMatch
	}

	return err
}
