package services

import (
	"fmt"
	"strings"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/arfan21/getprint-user/app/repository/mysql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user models.User) (*models.UserResoponse, error)
	GetByID(id string) (*models.UserResoponse, error)
	Update(user models.User) (*models.UserResoponse, error)
	Login(user models.User) (*models.UserResoponse, error)
	LoginUsingLine(dataLine *models.LineVerifyIdTokenResponse) (*models.UserResoponse, error)
}

type services struct {
	userRepo mysql.UserRepository
}

func NewUserServices(userRepo mysql.UserRepository) UserService {
	return &services{userRepo}
}

func (s *services) Create(user models.User) (*models.UserResoponse, error) {
	user.ID = uuid.NewV4()
	user.Identities.UserID = user.ID
	user.UserLog.UserID = user.ID
	if user.Identities.Provider == "" {
		user.Identities.Provider = "getprint"
		user.Identities.ProviderID = user.ID.String()
	}

	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	err := s.userRepo.Create(&user)

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
		ProviderID:    user.Identities.ProviderID,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
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
		ProviderID:    user.Identities.ProviderID,
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
		ProviderID:    user.Identities.ProviderID,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) Login(user models.User) (*models.UserResoponse, error) {
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

func (s *services) LoginUsingLine(dataLine *models.LineVerifyIdTokenResponse) (*models.UserResoponse, error) {
	lineID := dataLine.Sub
	userData, err := s.userRepo.GetByProviderID(lineID)

	fmt.Println("User data :", userData)
	fmt.Println("User err :", err)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			userData := new(models.User)

			userData.Identities.Provider = "line"
			userData.Identities.ProviderID = dataLine.Sub
			userData.Picture.Scan(dataLine.Picture)
			userData.Name = dataLine.Name
			userData.Email = fmt.Sprintf("%s@line.com", dataLine.Name)

			return s.Create(*userData)

		} else {
			return nil, err
		}

	}

	userData.UserLog.UserID = userData.ID

	go func() {
		_ = s.userRepo.UpdateUserLog(&userData.UserLog)
	}()

	return &models.UserResoponse{
		ID:            userData.ID,
		CreatedAt:     userData.CreatedAt,
		Name:          userData.Name,
		Picture:       userData.Picture.String,
		Email:         userData.Email,
		EmailVerified: userData.EmailVerified,
		PhoneNumber:   userData.PhoneNumber.String,
		Address:       userData.Address.String,
		Role:          userData.Role,
		Provider:      userData.Identities.Provider,
		LastLogin:     userData.UserLog.LastLogin.Time,
	}, nil
}
