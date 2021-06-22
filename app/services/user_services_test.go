package services

import (
	"testing"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/arfan21/getprint-user/app/repository/mysql"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserSrvTest struct {
	suite.Suite
	userSrv      UserService
	userRepoMock *mysql.MysqlUserRepositoryMock
}

func TestUserServices(t *testing.T) {
	userRepo := &mysql.MysqlUserRepositoryMock{}
	userSrv := NewUserServices(userRepo)

	userSrvTest := &UserSrvTest{
		userSrv:      userSrv,
		userRepoMock: userRepo,
	}

	suite.Run(t, userSrvTest)
}

func (testSuite *UserSrvTest) TestACreateSuccess() {
	dummyPayload := &models.User{
		Name:          "tesname",
		Email:         "test@test.com",
		EmailVerified: false,
		Role:          "buyer",
		Identities: models.Identities{
			Provider: "getprint",
		},
		UserLog: models.UserLog{},
	}
	testSuite.userRepoMock.Mock.On("Create", mock.Anything).Return(dummyPayload)

	res, err := testSuite.userSrv.Create(*dummyPayload)

	assert.NoError(testSuite.T(), err)
	assert.NotEqual(testSuite.T(), uuid.NullUUID{}.UUID, res.ID)
	assert.Equal(testSuite.T(), dummyPayload.Name, res.Name)
	assert.Equal(testSuite.T(), dummyPayload.Email, res.Email)
}
