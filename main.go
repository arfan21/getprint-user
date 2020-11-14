package main

import (
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
	port := ":" + os.Getenv("PORT")
	db, err := utils.Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatal(err.Error())
	}

	route := echo.New()

	route.Use(middleware.Recover())

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!! live-reload using air and using docker for development, thanks")
	})

	controllers.NewUserController(route, db)

	route.Logger.Fatal(route.Start(port))
}
