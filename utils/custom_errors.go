package utils

import (
	"strings"

	"github.com/arfan21/getprint-user/models"
)

func CustomErrors(err error) error {
	if strings.Contains(err.Error(), "users.email") {
		return models.ErrEmailConflict
	}

	if strings.Contains(err.Error(), "hashedPassword is not the hash") {
		return models.ErrPasswordNotMatch
	}

	return err
}
