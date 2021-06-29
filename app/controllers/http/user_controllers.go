package http

import (
	"net/http"

	"github.com/arfan21/getprint-user/app/model/modelresponse"
	"github.com/arfan21/getprint-user/app/model/modeluser"
	"github.com/arfan21/getprint-user/app/service/serviceuser"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"github.com/arfan21/getprint-user/app/helpers"
	"github.com/arfan21/getprint-user/validation"
)

type UserController interface {
	Create(c echo.Context) error
	GetByID(c echo.Context) error
	Update(c echo.Context) error
	Login(c echo.Context) error
	LoginLine(c echo.Context) error
}

type userController struct {
	userService serviceuser.UserService
}

func NewUserController(userService serviceuser.UserService) UserController {
	return &userController{userService}
}

func (s *userController) Create(c echo.Context) error {
	user := new(modeluser.User)

	//decoded request body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.Response("error", err.Error(), nil))
	}

	//validation user
	err = validation.Validate(*user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err, nil))
	}

	//save user into database
	data, err := s.userService.Create(*user)
	if err != nil {
		err = helpers.CustomErrors(err)
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err, nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (s *userController) GetByID(c echo.Context) error {
	id := c.Param("id")

	data, err := s.userService.GetByID(id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (s *userController) Update(c echo.Context) error {
	id := c.Param("id")
	user := new(modeluser.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.Response("error", err.Error(), nil))
	}

	user.ID = uuid.FromStringOrNil(id)

	data, err := s.userService.Update(*user)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (s *userController) Login(c echo.Context) error {

	user := new(modeluser.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.Response("error", err.Error(), nil))
	}

	data, err := s.userService.Login(*user)
	if err != nil {
		err = helpers.CustomErrors(err)
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (s *userController) LoginLine(c echo.Context) error {
	dataLine := new(modelresponse.LineVerifyIdToken)
	if err := c.Bind(dataLine); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.Response("error", err.Error(), nil))
	}

	data, err := s.userService.LoginUsingLine(dataLine)
	if err != nil {
		err = helpers.CustomErrors(err)
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}
