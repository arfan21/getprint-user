package repository

import (
	"github.com/arfan21/getprint-user/models"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) models.UserRepository {
	return &mysqlUserRepository{DB}
}

func (m *mysqlUserRepository) Create(user *models.User) error {
	err := m.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) Get(users *[]models.User) error {
	err := m.DB.Debug().Find(&users).Error

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) GetByID(id uint, user *models.User) error {
	err := m.DB.First(&user, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) GetByEmail(user *models.User) error {
	err := m.DB.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) Update(user *models.User) error {
	err := m.DB.Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}
