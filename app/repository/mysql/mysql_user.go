package mysql

import (
	"time"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/arfan21/getprint-user/config"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	Create(user models.User) (*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByProviderID(providerID string) (*models.User, error)
	Update(user *models.User) error
	UpdateUserLog(userLog *models.UserLog) error
}

type mysqlUserRepository struct {
	DB config.Client
}

func NewMysqlUserRepository(DB config.Client) UserRepository {
	return &mysqlUserRepository{DB}
}

func (repo *mysqlUserRepository) Create(user models.User) (*models.User, error) {
	err := repo.DB.Conn().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repo *mysqlUserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := new(models.User)
	err := repo.DB.Conn().Preload("Identities").Preload("UserLog").Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) GetByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := repo.DB.Conn().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) GetByProviderID(providerID string) (*models.User, error) {
	user := new(models.User)
	err := repo.DB.Conn().Debug().Joins("join identities ON identities.user_id =  users.id").Joins("JOIN user_logs ON user_logs.user_id = users.id").Where("identities.provider_id= ?", providerID).First(user).Scan(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) Update(user *models.User) error {
	oldData, err := repo.GetByID(user.ID)
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
