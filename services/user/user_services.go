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
	Create(user models.User) (*models.UserResoponse, error)
	Get(users *[]models.User) error
	GetByID(id string) (*models.UserResoponse, error)
	Update(user models.User) (*models.UserResoponse, error)
	Login(user models.User) (*models.UserLoginResponse, error)
}

type services struct {
	userRepo _userRepo.UserRepository
}

func NewUserServices(userRepo _userRepo.UserRepository) UserService {
	return &services{userRepo}
}

func (s *services) Create(user models.User) (*models.UserResoponse, error) {
	err := validation.Validate(user)
	if err != nil {
		return nil, err
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

	err = s.userRepo.Create(&user)

	if err != nil {
		return nil, err
	}

	return &models.UserResoponse{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		Name:          user.Name,
		Picture:       user.Picture.String,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		PhoneNumber:   user.PhoneNumber.String,
		Address:       user.Address.String,
		Role:          user.Role,
		Provider:      user.Identities.Provider,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) Get(users *[]models.User) error {
	err := s.userRepo.Get(users)

	if err != nil {
		return err
	}

	return nil
}

func (s *services) GetByID(id string) (*models.UserResoponse, error) {
	user := new(models.User)
	err := s.userRepo.GetByID(id, user)

	if err != nil {
		return nil, err
	}

	return &models.UserResoponse{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		Name:          user.Name,
		Picture:       user.Picture.String,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		PhoneNumber:   user.PhoneNumber.String,
		Address:       user.Address.String,
		Role:          user.Role,
		Provider:      user.Identities.Provider,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) Update(user models.User) (*models.UserResoponse, error) {
	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	err := s.userRepo.Update(&user)

	if err != nil {
		return nil, err
	}

	return &models.UserResoponse{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		Name:          user.Name,
		Picture:       user.Picture.String,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		PhoneNumber:   user.PhoneNumber.String,
		Address:       user.Address.String,
		Role:          user.Role,
		Provider:      user.Identities.Provider,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) Login(user models.User) (*models.UserLoginResponse, error) {
	password := user.Password

	err := s.userRepo.GetByEmail(&user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password.String))
	if err != nil {
		return nil, err
	}

	user.UserLog.UserID = user.ID

	go func() {
		_ = s.userRepo.UpdateUserLog(&user.UserLog)
	}()

	return &models.UserLoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
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
