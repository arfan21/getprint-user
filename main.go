package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arfan21/getprint-user/controllers"
	"github.com/arfan21/getprint-user/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world, thanks")
	})

	route.Static("/shared", "shared")

	controllers.NewUserController(route, db)

	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
