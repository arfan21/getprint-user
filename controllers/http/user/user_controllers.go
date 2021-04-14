package user

import (
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/arfan21/getprint-user/models"
	_userRepo "github.com/arfan21/getprint-user/repository/mysql/user"
	_userSrv "github.com/arfan21/getprint-user/services/user"
	"github.com/arfan21/getprint-user/utils"
)

type UserController interface {
	Routes(route *echo.Echo)
}

type userController struct {
	userService _userSrv.UserService
}

func NewUserController(db *gorm.DB) UserController {
	userRepo := _userRepo.NewMysqlUserRepository(db)
	userService := _userSrv.NewUserServices(userRepo)

	return &userController{userService}
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
		err = utils.CustomErrors(err)
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

	return c.JSON(http.StatusOK, utils.Response("success", nil, map[string]interface{}{"id": fmt.Sprint(user.ID), "email": user.Email, "role": user.Role}))
}
