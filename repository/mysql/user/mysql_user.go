package user

import (
	"time"

	"gorm.io/gorm"

	"github.com/arfan21/getprint-user/models"
)

type UserRepository interface {
	Create(user *models.User) error
	Get(users *[]models.User) error
	GetByID(id string, user *models.User) error
	GetByEmail(user *models.User) error
	GetByLineID(lineID string) (*models.User, error)
	Update(user *models.User) error
	UpdateUserLog(userLog *models.UserLog) error
}

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) UserRepository {
	return &mysqlUserRepository{DB}
}

func (repo *mysqlUserRepository) Create(user *models.User) error {
	return repo.DB.Create(&user).Error
}

func (repo *mysqlUserRepository) Get(users *[]models.User) error {
	return repo.DB.Debug().Find(&users).Error
}

func (repo *mysqlUserRepository) GetByID(id string, user *models.User) error {
	return repo.DB.Preload("Identities").Preload("UserLog").Where("id=?", id).First(&user).Error
}

func (repo *mysqlUserRepository) GetByEmail(user *models.User) error {
	return repo.DB.Where("email = ?", user.Email).First(&user).Error
}

func (repo *mysqlUserRepository) GetByLineID(lineID string) (*models.User, error) {
	user := new(models.User)
	err := repo.DB.Debug().Joins("join identities ON identities.user_id =  users.id").Joins("JOIN user_logs ON user_logs.user_id = users.id").Where("identities.provider_id= ?", lineID).First(user).Scan(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) Update(user *models.User) error {
	oldData := new(models.User)
	err := repo.GetByID(user.ID.String(), oldData)
	if err != nil {
		return err
	}
	err = repo.DB.Model(user).Updates(user).Error
	if err != nil {
		return err
	}

	user.CreatedAt = oldData.CreatedAt
	user.Role = oldData.Role
	user.Identities.Provider = oldData.Identities.Provider
	user.UserLog.LastLogin.Scan(oldData.UserLog.LastLogin.Time)

	return nil
}

func (repo *mysqlUserRepository) UpdateUserLog(userLog *models.UserLog) error {
	err := repo.DB.Model(userLog).Where("user_id=?", userLog.UserID).First(userLog).Error
	if err != nil {
		return err
	}

	err = repo.DB.Model(userLog).Where("user_id=?", userLog.UserID).Update("last_login", time.Now()).Error

	if err != nil {
		return err
	}

	return nil
}
