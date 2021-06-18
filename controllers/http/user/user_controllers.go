package user

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/arfan21/getprint-user/models"
	_userRepo "github.com/arfan21/getprint-user/repository/mysql/user"
	_userSrv "github.com/arfan21/getprint-user/services/user"
	"github.com/arfan21/getprint-user/utils"
	"github.com/arfan21/getprint-user/validation"
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
	user := new(models.User)

	//decoded request body
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	//validation user
	err = validation.Validate(*user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err, nil))
	}

	//save user into database
	data, err := s.userService.Create(*user)
	if err != nil {
		err = utils.CustomErrors(err)
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, data))
}

func (s *userController) GetByID(c echo.Context) error {
	id := c.Param("id")

	data, err := s.userService.GetByID(id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, data))
}

func (s *userController) Update(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	user.ID = uuid.FromStringOrNil(id)

	data, err := s.userService.Update(*user)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, data))
}

func (s *userController) Login(c echo.Context) error {

	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	data, err := s.userService.Login(*user)
	if err != nil {
		err = utils.CustomErrors(err)
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, data))
}

func (s *userController) LoginLine(c echo.Context) error {
	dataLine := new(models.LineVerifyIdTokenResponse)
	if err := c.Bind(dataLine); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Response("error", err.Error(), nil))
	}

	data, err := s.userService.LoginUsingLine(dataLine)
	if err != nil {
		err = utils.CustomErrors(err)
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, data))
}
