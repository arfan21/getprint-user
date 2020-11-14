package utils

import (
	"net/http"

	"github.com/arfan21/getprint-user/models"
)

func GetStatusCode(err error) int {
	switch err.Error() {
	case models.ErrorEmailRegistered:
		return http.StatusConflict
	case models.ErrorEmailNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
