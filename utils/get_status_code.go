package utils

import (
	"net/http"
	"strings"

	"github.com/arfan21/getprint-user/models"
)

func GetStatusCode(err error) int {
	if strings.Contains(err.Error(), "Duplicate") {
		return http.StatusConflict
	}
	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound
	}

	if err.Error() == models.ErrEmailConflict.Error() {
		return http.StatusConflict
	}

	switch err {
	case models.ErrBadParamInput:
		return http.StatusBadRequest
	case models.ErrConflict:
		return http.StatusConflict
	case models.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
