package controllers

import (
	"net/http"
	"strconv"

	"github.com/arfan21/getprint-user/models"
	"github.com/arfan21/getprint-user/repository"
	"github.com/arfan21/getprint-user/services"
	"github.com/arfan21/getprint-user/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	userService models.UserService
}

func NewUserController(route *echo.Echo, db *gorm.DB) {
	userRepo := repository.NewMysqlUserRepository(db)
	userService := services.NewUserServices(userRepo)
	controllers := &userController{userService}

	route.POST("/users", controllers.Create)
	route.GET("/users/:id", controllers.GetByID)
	route.PUT("/users/:id", controllers.Update)
}

type Test struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *userController) Create(c echo.Context) error {
	user := &models.User{}

	//decoded request body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	//validate user value
	err = user.Validate()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "error validating data", err))
	}

	//save user into database
	err = s.userService.Create(user)
	if err != nil {
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "Success create user", user))
}

func (s *userController) GetByID(c echo.Context) error {
	user := &models.User{}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "invalid id", nil))
	}

	err = s.userService.GetByID(uint(id), user)
	if err != nil {
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success get user", user))
}

func (s *userController) Update(c echo.Context) error {
	user := &models.User{}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "invalid id", nil))
	}

	err = s.userService.GetByID(uint(id), user)
	if err != nil {
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	err = c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err = user.Validate()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "error validating data", err))
	}

	err = s.userService.Update(user)
	if err != nil {
		err = utils.FormatedErrors(err.Error())
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success update user", user))
}
