package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/arfan21/getprint-user/app/constant"
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
	case constants.ErrBadParamInput:
		return http.StatusBadRequest
	case constants.ErrConflict:
		return http.StatusConflict
	case constants.ErrNotFound:
		return http.StatusNotFound
	case constants.ErrEmailConflict:
		return http.StatusConflict
	case constants.ErrPasswordNotMatch:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
