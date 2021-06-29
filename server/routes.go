package server

import (
	"github.com/arfan21/getprint-user/app/controllers/http"
	"github.com/arfan21/getprint-user/app/repository/mysql/mysqluser"
	"github.com/arfan21/getprint-user/app/service/serviceuser"
	"github.com/arfan21/getprint-user/config/database/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(mysqlClient mysql.Client) *echo.Echo {
	route := echo.New()
	route.Use(middleware.Recover())
	route.Use(middleware.Logger())
	apiV1 := route.Group("/v1")

	// routing order
	userRepo := mysqluser.New(mysqlClient)
	userSRv := serviceuser.New(userRepo)
	userCtrl := http.NewUserController(userSRv)

	apiUser := apiV1.Group("/user")
	apiUser.POST("", userCtrl.Create)
	apiUser.GET("/:id", userCtrl.GetByID)
	apiUser.PUT("/:id", userCtrl.Update)
	apiUser.POST("/login", userCtrl.Login)
	apiUser.POST("/login-line", userCtrl.LoginLine)

	return route
}
