package controllers

import (
	"net/http"
	"strconv"

	"github.com/arfan21/getprint-user/models"
	"github.com/arfan21/getprint-user/repository"
	"github.com/arfan21/getprint-user/services"
	"github.com/arfan21/getprint-user/utils"
	"github.com/arfan21/getprint-user/validation"
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

	route.POST("/user", controllers.Create)
	route.GET("/user/:id", controllers.GetByID)
	route.PUT("/user/:id", controllers.Update)
}

func (s *userController) Create(c echo.Context) error {
	user := &models.User{}

	//decoded request body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	//validate user value
	err = validation.Validate(*user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "error validating data", err))
	}

	//save user into database
	err = s.userService.Create(user)
	if err != nil {
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
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	err = c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	err = validation.Validate(*user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "error validating data", err))
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

	if err := validation.ValidateLogin(*user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", "error validating data", err))
	}

	err := s.userService.Login(user)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "success login", user))
}
