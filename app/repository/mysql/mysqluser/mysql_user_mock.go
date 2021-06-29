package mysqluser

import (
	"errors"

	"github.com/arfan21/getprint-user/app/model/modeluser"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
)

type MysqlUserRepositoryMock struct {
	Mock mock.Mock
}

func (m *MysqlUserRepositoryMock) Create(user modeluser.User) (*modeluser.User, error) {
	arg := m.Mock.Called(user)
	if arg.Get(0) == nil {
		return nil, errors.New("failed create user")
	} else {
		newUser := arg.Get(0).(modeluser.User)
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
func (m *MysqlUserRepositoryMock) GetByID(id uuid.UUID) (*modeluser.User, error) {
	arg := m.Mock.Called(id)
	if arg.Get(0) == nil {
		return nil, errors.New("failed get user")
	} else {
		user := arg.Get(0).(modeluser.User)
		return &user, nil
	}
}

func (m *MysqlUserRepositoryMock) GetByEmail(email string) (*modeluser.User, error) {
	arg := m.Mock.Called(email)
	if arg.Get(0) == nil {
		return nil, errors.New("failed get user")
	} else {
		user := arg.Get(0).(*modeluser.User)
		if user.Email == "" {
			return nil, errors.New("need email")
		}

		if user.Email != email {
			return nil, errors.New("email not found")
		}

		return user, nil
	}
}

func (m *MysqlUserRepositoryMock) GetByProviderID(providerID string) (*modeluser.User, error) {
	arg := m.Mock.Called(providerID)
	if arg.Get(0) == nil {
		return nil, errors.New("failed get user")
	} else {
		user := arg.Get(0).(*modeluser.User)

		if user.Identities.ProviderID != providerID {
			return nil, errors.New("providerID not found")
		}

		return user, nil
	}
}

func (m *MysqlUserRepositoryMock) Update(user *modeluser.User) error {
	arg := m.Mock.Called(user)
	if arg.Get(0) == nil {
		return errors.New("failed update user")
	} else {
		newUser := arg.Get(0).(modeluser.User)
		newUser.ID = user.ID
		if newUser.Email == "" {
			return errors.New("need email")
		}
		user = &newUser
		return nil
	}
}
func (m *MysqlUserRepositoryMock) UpdateUserLog(userLog *modeluser.UserLog) error {
	return nil
}
