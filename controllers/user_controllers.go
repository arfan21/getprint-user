package controllers

import (
	"fmt"
	"github.com/arfan21/getprint-user/models"
	"github.com/arfan21/getprint-user/repository"
	"github.com/arfan21/getprint-user/services"
	"github.com/arfan21/getprint-user/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type userController struct {
	userService models.UserService
}

func NewUserController(route *echo.Echo, db *gorm.DB) {
	userRepo := repository.NewMysqlUserRepository(db)
	userService := services.NewUserServices(userRepo)
	controllers := &userController{userService}

	route.POST("/user", controllers.Create)
	route.GET("/user/:id", controllers.GetByID)
	route.PUT("/user/:id", controllers.Update)
	route.POST("/login", controllers.Login)
}

func (s *userController) Create(c echo.Context) error {
	user := &models.User{}

	//decoded request body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	//save user into database
	err = s.userService.Create(user)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, user))
}

func (s *userController) GetByID(c echo.Context) error {
	user := &models.User{}

	id := c.Param("id")

	err := s.userService.GetByID(id, user)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success get user", user))
}

func (s *userController) Update(c echo.Context) error {
	user := &models.User{}

	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err = s.userService.Update(user)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success update user", user))
}

func (s *userController) Login(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err := s.userService.Login(user)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	//aud := os.Getenv("AUD")
	//iss := os.Getenv("ISS")
	//accessTokenExp := time.Now().Add(time.Minute * 5).Unix()
	//refreshTokenExp := time.Now().AddDate(0, 0, 7).Unix()
	//response := map[string]interface{}{
	//
	//	"access_token": map[string]interface{}{
	//		"aud":          aud,
	//		"iss":          iss,
	//		"sub":          fmt.Sprint(user.ID),
	//		"user_id_line": user.UserIDLine,
	//		"name":         user.Name,
	//		"email":        user.Email,
	//		"picture":      user.Picture.String,
	//		"role":         user.Role,
	//		"exp":          accessTokenExp,
	//	},
	//	"refresh_token": map[string]interface{}{
	//		"aud":          aud,
	//		"iss":          iss,
	//		"user_id_line": user.UserIDLine,
	//		"name":         user.Name,
	//		"email":        user.Email,
	//		"picture":      user.Picture.String,
	//		"sub":          fmt.Sprint(user.ID),
	//		"exp":          refreshTokenExp,
	//	},
	//	"exp": refreshTokenExp,
	//}

	return c.JSON(http.StatusOK, utils.Response("success", nil, map[string]interface{}{"id" : fmt.Sprint(user.ID)}))
}
