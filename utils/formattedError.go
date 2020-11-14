package utils

import (
	"errors"
	"strings"

	"github.com/arfan21/getprint-user/models"
)

func FormatedErrors(err string) error {
	switch {
	case strings.Contains(err, "email"):
		return errors.New(models.ErrorEmailRegistered)
	case strings.Contains(err, "not found"):
		return errors.New(models.ErrorEmailNotFound)
	default:
		return errors.New(err)
	}
}
