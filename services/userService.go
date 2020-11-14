package services

import (
	"github.com/arfan21/getprint-user/models"
	"golang.org/x/crypto/bcrypt"
)

type services struct {
	userRepo models.UserRepository
}

func NewUserServices(userRepo models.UserRepository) models.UserService {
	return &services{userRepo}
}

func (s *services) Create(user *models.User) error {
	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	err := s.userRepo.Create(user)

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

func (s *services) GetByID(id uint, user *models.User) error {
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
