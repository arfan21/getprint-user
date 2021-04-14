package user

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/arfan21/getprint-user/models"
	_userRepo "github.com/arfan21/getprint-user/repository/mysql/user"
	"github.com/arfan21/getprint-user/validation"
)

type UserService interface {
	Create(user *models.User) error
	Get(users *[]models.User) error
	GetByID(id string, user *models.User) error
	Update(user *models.User) error
	Login(user *models.User) error
}

type services struct {
	userRepo _userRepo.UserRepository
}

func NewUserServices(userRepo _userRepo.UserRepository) UserService {
	return &services{userRepo}
}

func (s *services) Create(user *models.User) error {
	err := validation.Validate(*user)
	if err != nil {
		return err
	}

	user.ID = uuid.NewV4()
	user.Identities.UserID = user.ID
	user.UserLog.UserID = user.ID
	if user.Identities.Provider == "" {
		user.Identities.Provider = "getprint"
		user.Identities.UserIDProvider = user.ID.String()
	}

	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	err = s.userRepo.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *services) Get(users *[]models.User) error {
	err := s.userRepo.Get(users)

	if err != nil {
		return err
	}

	return nil
}

func (s *services) GetByID(id string, user *models.User) error {
	err := s.userRepo.GetByID(id, user)

	user.Password.Scan("")
	if err != nil {
		return err
	}

	return nil
}

func (s *services) Update(user *models.User) error {
	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	err := s.userRepo.Update(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *services) Login(user *models.User) error {
	password := user.Password

	err := s.userRepo.GetByEmail(user)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password.String))

	if err != nil {
		return err
	}

	return nil
}

func (s *services) LoginUsingLine(user *models.User) error {
	lineID := user.Identities.UserIDProvider
	fmt.Println(lineID)
	err := s.userRepo.GetByLineID(user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {

		}
		return err
	}
	return nil
}
