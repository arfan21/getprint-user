package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_userCtrl "github.com/arfan21/getprint-user/controllers/http/user"
	"github.com/arfan21/getprint-user/utils"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	db, err := utils.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	route := echo.New()

	route.Use(middleware.Recover())
	route.Use(middleware.Logger())

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world, thanks")
	})

	route.Static("/shared", "shared")

	userCtrl := _userCtrl.NewUserController(db)
	userCtrl.Routes(route)
	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
