package server

import (
	"github.com/arfan21/getprint-user/app/controllers/http"
	"github.com/arfan21/getprint-user/app/repository/mysql"
	"github.com/arfan21/getprint-user/app/services"
	"github.com/arfan21/getprint-user/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(mysqlClient configs.Client) *echo.Echo{
	route := echo.New()
	route.Use(middleware.Recover())
	route.Use(middleware.Logger())
	apiV1 := route.Group("/v1")

	// routing order
	userRepo := mysql.NewMysqlUserRepository(mysqlClient)
	userSRv := services.NewUserServices(userRepo)
	userCtrl := http.NewUserController(userSRv)

	apiUser := apiV1.Group("/user")
	apiUser.POST("", userCtrl.Create)
	apiUser.GET("/:id", userCtrl.GetByID)
	apiUser.PUT("/:id", userCtrl.Update)
	apiUser.POST("/login", userCtrl.Login)
	apiUser.POST("/login-line", userCtrl.LoginLine)

	return route
}