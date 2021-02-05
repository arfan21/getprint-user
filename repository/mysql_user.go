package repository

import (
	"github.com/arfan21/getprint-user/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) models.UserRepository {
	return &mysqlUserRepository{DB}
}

func (repo *mysqlUserRepository) Create(user *models.User) error {
	user.ID = uuid.NewV4()
	err := repo.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) Get(users *[]models.User) error {
	err := repo.DB.Debug().Find(&users).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) GetByID(id string, user *models.User) error {
	err := repo.DB.First(&user, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) GetByEmail(user *models.User) error {
	err := repo.DB.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) GetByLineID(user *models.User) error{
	err := repo.DB.Where("user_id_provider=?", user.Identities.UserIDProvider).First(user.Identities).Error

	if err != nil{
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) Update(user *models.User) error {
	err := repo.DB.Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) UpdateUserLog(userLog *models.UserLog) error{
	err := repo.DB.Model(userLog).Where("user_id=?", userLog.UserID).Update("last_login", userLog.LastLogin).Error

	if err != nil{
		return err
	}

	return nil
}
