package user

import (
	"github.com/labstack/echo/v4"
)

func (ctrl userController) Routes(route *echo.Echo) {
	route.POST("/user", ctrl.Create)
	route.GET("/user/:id", ctrl.GetByID)
	route.PUT("/user/:id", ctrl.Update)
	route.POST("/login", ctrl.Login)
	route.POST("/login-line", ctrl.LoginLine)
}
