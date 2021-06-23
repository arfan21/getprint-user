package services

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/arfan21/getprint-user/app/repository/mysql"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type UserSrvTest struct {
	suite.Suite
	userSrv      UserService
	userRepoMock *mysql.MysqlUserRepositoryMock
	dummyPayload *models.User
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

func (testSuite *UserSrvTest) TestACreateFail() {
	dummyPayload := models.User{
		Name:          "tesname",
		EmailVerified: false,
		Role:          "buyer",
		Identities: models.Identities{
			Provider: "getprint",
		},
		UserLog: models.UserLog{},
	}

	testSuite.userRepoMock.Mock.On("Create", mock.AnythingOfType("models.User")).Return(dummyPayload).Once()

	res, err := testSuite.userSrv.Create(dummyPayload)
	assert.Error(testSuite.T(), err)
	assert.Equal(testSuite.T(), true, strings.Contains(err.Error(), "email"))
	assert.Nil(testSuite.T(), res)
}
func (testSuite *UserSrvTest) TestBCreateSuccess() {
	dummyPayload := models.User{
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

	res, err := testSuite.userSrv.Create(dummyPayload)
	assert.NoError(testSuite.T(), err)
	assert.NotEqual(testSuite.T(), uuid.NullUUID{}.UUID, res.ID)
	assert.Equal(testSuite.T(), dummyPayload.Name, res.Name)
	assert.Equal(testSuite.T(), dummyPayload.Email, res.Email)
	assert.NotEqual(testSuite.T(), uuid.NullUUID{}.UUID, res.ProviderID)

	dummyPayload.ID = res.ID
	dummyPayload.CreatedAt = res.CreatedAt
	dummyPayload.Name = res.Name
	dummyPayload.Picture.String = res.Picture
	dummyPayload.Email = res.Email
	dummyPayload.EmailVerified = res.EmailVerified
	dummyPayload.PhoneNumber.String = res.PhoneNumber
	dummyPayload.Address.String = res.Address
	dummyPayload.Role = res.Role
	dummyPayload.Identities.Provider = res.Provider
	dummyPayload.Identities.ProviderID = res.ProviderID
	dummyPayload.UserLog.LastLogin.Time = res.LastLogin

	testSuite.dummyPayload = &dummyPayload
}

func (testSuite *UserSrvTest) TestCGetByIDSuccess() {
	testSuite.userRepoMock.Mock.On("GetByID", testSuite.dummyPayload.ID).Return(*testSuite.dummyPayload)

	res, err := testSuite.userSrv.GetByID(testSuite.dummyPayload.ID.String())
	assert.NoError(testSuite.T(), err)
	assert.NotNil(testSuite.T(), res)
	assert.Equal(testSuite.T(), testSuite.dummyPayload.ID.String(), res.ID.String())
}
func (testSuite *UserSrvTest) TestDGetByIDFail() {
	testSuite.userRepoMock.Mock.On("GetByID", uuid.UUID{}).Return(nil)

	res, err := testSuite.userSrv.GetByID("")
	assert.Error(testSuite.T(), err)
	assert.Nil(testSuite.T(), res)
}

func (testSuite *UserSrvTest) TestEUpdateFail() {
	dummyPayloadByte, _ := json.Marshal(testSuite.dummyPayload)
	dummyPayload := new(models.User)
	_ = json.Unmarshal(dummyPayloadByte, dummyPayload)

	dummyPayload.Email = ""
	dummyPayload.Name = "testUpdateName"

	testSuite.userRepoMock.Mock.On("Update", mock.Anything).Return(*dummyPayload).Once()

	res, err := testSuite.userSrv.Update(*dummyPayload)

	assert.Error(testSuite.T(), err)
	assert.Nil(testSuite.T(), res)
}

func (testSuite *UserSrvTest) TestFUpdateSuccess() {
	dummyPayloadByte, _ := json.Marshal(testSuite.dummyPayload)
	dummyPayload := new(models.User)
	_ = json.Unmarshal(dummyPayloadByte, dummyPayload)

	dummyPayload.Email = "testUpdateEmail@tes.com"
	dummyPayload.Name = "testUpdateName"

	testSuite.userRepoMock.Mock.On("Update", mock.Anything).Return(*dummyPayload).Once()

	res, err := testSuite.userSrv.Update(*dummyPayload)

	assert.NoError(testSuite.T(), err)
	assert.NotNil(testSuite.T(), res)
	assert.Equal(testSuite.T(), dummyPayload.Email, res.Email)
	assert.Equal(testSuite.T(), dummyPayload.Name, res.Name)
}

func (testSuite *UserSrvTest) TestGLoginSuccess() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123qweasd"), bcrypt.DefaultCost)
	testSuite.dummyPayload.Password.Scan(hashedPassword)
	testSuite.userRepoMock.Mock.On("GetByEmail", testSuite.dummyPayload.Email).Return(testSuite.dummyPayload).Once()

	dummyPayloadByte, _ := json.Marshal(testSuite.dummyPayload)
	dummyPayload := new(models.User)
	_ = json.Unmarshal(dummyPayloadByte, dummyPayload)
	dummyPayload.Password.Scan("123qweasd")

	res, err := testSuite.userSrv.Login(*dummyPayload)
	assert.NoError(testSuite.T(), err)
	assert.NotNil(testSuite.T(), res)
	assert.Equal(testSuite.T(), dummyPayload.Email, res.Email)
}
func (testSuite *UserSrvTest) TestHLoginFailed() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123qweasd"), bcrypt.DefaultCost)
	testSuite.dummyPayload.Password.Scan(hashedPassword)
	testSuite.userRepoMock.Mock.On("GetByEmail", testSuite.dummyPayload.Email).Return(testSuite.dummyPayload).Once()

	dummyPayloadByte, _ := json.Marshal(testSuite.dummyPayload)
	dummyPayload := new(models.User)
	_ = json.Unmarshal(dummyPayloadByte, dummyPayload)
	dummyPayload.Password.Scan("qwe123asd")

	res, err := testSuite.userSrv.Login(*dummyPayload)
	assert.Error(testSuite.T(), err)
	assert.Nil(testSuite.T(), res)
}

func (testSuite *UserSrvTest) TestILoginLineAlreadyRegisteredSuccess() {
	dummyPayload := &models.LineVerifyIdTokenResponse{
		Sub: testSuite.dummyPayload.ID.String(),
	}
	testSuite.userRepoMock.Mock.On("GetByProviderID", dummyPayload.Sub).Return(testSuite.dummyPayload).Once()

	res, err := testSuite.userSrv.LoginUsingLine(dummyPayload)
	assert.NoError(testSuite.T(), err)
	assert.NotNil(testSuite.T(), res)
	assert.Equal(testSuite.T(), testSuite.dummyPayload.Email, res.Email)
}

func (testSuite *UserSrvTest) TestJLoginLineNotRegisteredSuccess() {
	newProviderID := "eyhhheeeasdawawdmmma822lascm"
	dummyPayload := &models.LineVerifyIdTokenResponse{
		Sub:     newProviderID,
		Name:    "testname",
		Picture: "https://image.com/image.jpg",
	}
	testSuite.userRepoMock.Mock.On("GetByProviderID", dummyPayload.Sub).Return(testSuite.dummyPayload).Once()

	dummyPayloadRegistered := models.User{
		Name:  "testname",
		Email: "testname@line.com",
		Identities: models.Identities{
			Provider:   "line",
			ProviderID: newProviderID,
		},
		UserLog: models.UserLog{},
	}
	dummyPayloadRegistered.Picture.Scan(dummyPayload.Picture)
	testSuite.userRepoMock.Mock.On("Create", mock.Anything).Return(dummyPayloadRegistered).Once()

	res, err := testSuite.userSrv.LoginUsingLine(dummyPayload)
	assert.NoError(testSuite.T(), err)
	assert.NotNil(testSuite.T(), res)
	assert.Equal(testSuite.T(), testSuite.dummyPayload.Email, res.Email)
	assert.NotEqual(testSuite.T(), "", res.ProviderID)
}
