package mysql

import (
	"errors"

	"github.com/arfan21/getprint-user/app/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
)

type MysqlUserRepositoryMock struct {
	Mock mock.Mock
}

func (m *MysqlUserRepositoryMock) Create(user models.User) (*models.User, error) {
	arg := m.Mock.Called(user)
	if arg.Get(0) == nil {
		return nil, errors.New("failed create user")
	} else {
		newUser := arg.Get(0).(models.User)
		newUser.ID = user.ID
		newUser.Identities.ProviderID = user.Identities.ProviderID
		newUser.Identities.UserID = user.Identities.UserID
		newUser.UserLog.UserID = user.UserLog.UserID
		if newUser.Email == "" {
			return nil, errors.New("need email")
		}

		return &newUser, nil
	}
}
func (m *MysqlUserRepositoryMock) GetByID(id uuid.UUID) (*models.User, error) {
	arg := m.Mock.Called(id)
	if arg.Get(0) == nil {
		return nil, errors.New("failed get user")
	} else {
		user := arg.Get(0).(models.User)
		return &user, nil
	}
}

func (m *MysqlUserRepositoryMock) GetByEmail(email string) (*models.User, error) {
	arg := m.Mock.Called(email)
	if arg.Get(0) == nil {
		return nil, errors.New("failed get user")
	} else {
		user := arg.Get(0).(*models.User)
		if user.Email == "" {
			return nil, errors.New("need email")
		}

		if user.Email != email {
			return nil, errors.New("email not found")
		}

		return user, nil
	}
}

func (m *MysqlUserRepositoryMock) GetByProviderID(providerID string) (*models.User, error) {
	arg := m.Mock.Called(providerID)
	if arg.Get(0) == nil {
		return nil, errors.New("failed get user")
	} else {
		user := arg.Get(0).(*models.User)

		if user.Identities.ProviderID != providerID {
			return nil, errors.New("providerID not found")
		}

		return user, nil
	}
}

func (m *MysqlUserRepositoryMock) Update(user *models.User) error {
	arg := m.Mock.Called(user)
	if arg.Get(0) == nil {
		return errors.New("failed update user")
	} else {
		newUser := arg.Get(0).(models.User)
		newUser.ID = user.ID
		if newUser.Email == "" {
			return errors.New("need email")
		}
		user = &newUser
		return nil
	}
}
func (m *MysqlUserRepositoryMock) UpdateUserLog(userLog *models.UserLog) error {
	return nil
}
