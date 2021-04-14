package user

import (
	"gorm.io/gorm"

	"github.com/arfan21/getprint-user/models"
)

type UserRepository interface{
	Create(user *models.User) error
	Get(users *[]models.User) error
	GetByID(id string, user *models.User) error
	GetByEmail(user *models.User) error
	GetByLineID(user *models.User) error
	Update(user *models.User) error
}

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) UserRepository{
	return &mysqlUserRepository{DB}
}

func (repo *mysqlUserRepository) Create(user *models.User) error {
	return repo.DB.Create(&user).Error
}

func (repo *mysqlUserRepository) Get(users *[]models.User) error {
	return repo.DB.Debug().Find(&users).Error
}

func (repo *mysqlUserRepository) GetByID(id string, user *models.User) error {
	return  repo.DB.First(&user, id).Error
}

func (repo *mysqlUserRepository) GetByEmail(user *models.User) error {
	return repo.DB.Where("email = ?", user.Email).First(&user).Error
}

func (repo *mysqlUserRepository) GetByLineID(user *models.User) error{
	return repo.DB.Where("user_id_provider=?", user.Identities.UserIDProvider).First(user.Identities).Error
}

func (repo *mysqlUserRepository) Update(user *models.User) error {
	return repo.DB.Save(&user).Error
}

func (repo *mysqlUserRepository) UpdateUserLog(userLog *models.UserLog) error{
	err := repo.DB.Model(userLog).Where("user_id=?", userLog.UserID).First(userLog).Error
	if err != nil{
		return err
	}

	err = repo.DB.Model(userLog).Where("user_id=?", userLog.UserID).Update("last_login", userLog.LastLogin).Error

	if err != nil{
		return err
	}

	return nil
}