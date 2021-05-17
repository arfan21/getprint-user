package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/arfan21/getprint-user/models"
)

func GetStatusCode(err error) int {
	fmt.Println(err.Error())
	if strings.Contains(err.Error(), "Duplicate") {
		return http.StatusConflict
	}
	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound
	}

	switch err {
	case models.ErrBadParamInput:
		return http.StatusBadRequest
	case models.ErrConflict:
		return http.StatusConflict
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrEmailConflict:
		return http.StatusConflict
	case models.ErrPasswordNotMatch:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
