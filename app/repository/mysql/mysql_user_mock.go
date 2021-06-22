package mysql

import (
	"errors"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/stretchr/testify/mock"
)

type MysqlUserRepositoryMock struct {
	Mock mock.Mock
}

func (m *MysqlUserRepositoryMock) Create(user *models.User) error {
	arg := m.Mock.Called(user)
	if arg.Get(0) == nil {
		return errors.New("failed create user")
	} else {
		newUser := arg.Get(0).(*models.User)
		user = newUser
		return nil
	}
}
func (m *MysqlUserRepositoryMock) GetByID(id string, user *models.User) error {
	return nil
}
func (m *MysqlUserRepositoryMock) GetByEmail(user *models.User) error {
	return nil
}
func (m *MysqlUserRepositoryMock) GetByProviderID(providerID string) (*models.User, error) {
	return nil, nil
}
func (m *MysqlUserRepositoryMock) Update(user *models.User) error {
	return nil
}
func (m *MysqlUserRepositoryMock) UpdateUserLog(userLog *models.UserLog) error {
	return nil
}
