package mysql

import (
	"time"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/arfan21/getprint-user/configs"
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
	DB configs.Client
}

func NewMysqlUserRepository(DB configs.Client) UserRepository {
	return &mysqlUserRepository{DB}
}

func (repo *mysqlUserRepository) Create(user *models.User) error {
	return repo.DB.Conn().Create(&user).Error
}

func (repo *mysqlUserRepository) Get(users *[]models.User) error {
	return repo.DB.Conn().Debug().Find(&users).Error
}

func (repo *mysqlUserRepository) GetByID(id string, user *models.User) error {
	return repo.DB.Conn().Preload("Identities").Preload("UserLog").Where("id=?", id).First(&user).Error
}

func (repo *mysqlUserRepository) GetByEmail(user *models.User) error {
	return repo.DB.Conn().Where("email = ?", user.Email).First(&user).Error
}

func (repo *mysqlUserRepository) GetByLineID(lineID string) (*models.User, error) {
	user := new(models.User)
	err := repo.DB.Conn().Debug().Joins("join identities ON identities.user_id =  users.id").Joins("JOIN user_logs ON user_logs.user_id = users.id").Where("identities.provider_id= ?", lineID).First(user).Scan(user).Error

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
	err = repo.DB.Conn().Model(user).Updates(user).Error
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
	err := repo.DB.Conn().Model(userLog).Where("user_id=?", userLog.UserID).First(userLog).Error
	if err != nil {
		return err
	}

	err = repo.DB.Conn().Model(userLog).Where("user_id=?", userLog.UserID).Update("last_login", time.Now()).Error

	if err != nil {
		return err
	}

	return nil
}
