package mysqluser

import (
	"time"

	"github.com/arfan21/getprint-user/app/model/modeluser"
	"github.com/arfan21/getprint-user/config/database/mysql"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	Create(user modeluser.User) (*modeluser.User, error)
	GetByID(id uuid.UUID) (*modeluser.User, error)
	GetByEmail(email string) (*modeluser.User, error)
	GetByProviderID(providerID string) (*modeluser.User, error)
	Update(user *modeluser.User) error
	UpdateUserLog(userLog *modeluser.UserLog) error
}

type mysqlUserRepository struct {
	DB mysql.Client
}

func New(DB mysql.Client) UserRepository {
	return &mysqlUserRepository{DB}
}

func (repo *mysqlUserRepository) Create(user modeluser.User) (*modeluser.User, error) {
	err := repo.DB.Conn().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repo *mysqlUserRepository) GetByID(id uuid.UUID) (*modeluser.User, error) {
	user := new(modeluser.User)
	err := repo.DB.Conn().Preload("Identities").Preload("UserLog").Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) GetByEmail(email string) (*modeluser.User, error) {
	user := new(modeluser.User)
	err := repo.DB.Conn().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) GetByProviderID(providerID string) (*modeluser.User, error) {
	user := new(modeluser.User)
	err := repo.DB.Conn().Debug().Joins("join identities ON identities.user_id =  users.id").Joins("JOIN user_logs ON user_logs.user_id = users.id").Where("identities.provider_id= ?", providerID).First(user).Scan(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *mysqlUserRepository) Update(user *modeluser.User) error {
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

func (repo *mysqlUserRepository) UpdateUserLog(userLog *modeluser.UserLog) error {
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
