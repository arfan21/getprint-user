package mysql

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/arfan21/getprint-user/app/models"
	"github.com/arfan21/getprint-user/config"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func InitializeDatabase() (config.Client, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal(err)
	}

	mysqlConfig := config.NewConfigForTest()
	mysqlClient, err := config.NewClient(mysqlConfig.String())
	if err != nil {
		return nil, err
	}

	return mysqlClient, nil
}

type MySQLUserTest struct {
	suite.Suite
	mysqlClient config.Client
	userRepo    UserRepository
	dataUser    *models.User
}

func TestMySQLUserTest(t *testing.T) {
	db, err := InitializeDatabase()
	if err != nil {
		t.Fatal(err)
	}

	db.Conn().Unscoped().Where("1 = 1").Delete(&models.User{})
	userRepo := NewMysqlUserRepository(db)

	mySQLUserTest := &MySQLUserTest{
		mysqlClient: db,
		userRepo:    userRepo,
	}

	suite.Run(t, mySQLUserTest)
}

func (testSuit *MySQLUserTest) TearDownSuite() {
	log.Println("Test All Done!!!")
	testSuit.mysqlClient.Conn().Debug().Exec("DROP DATABASE " + os.Getenv("DB_NAME_TEST"))
	defer testSuit.mysqlClient.Close()
}

func (testSuite *MySQLUserTest) TestAcreateSuccess() {

	newUUID := uuid.NewV4()
	dummyPayload := &models.User{
		ID:            newUUID,
		Name:          "tesname",
		Email:         "test@test.com",
		EmailVerified: false,
		Role:          "buyer",
		Identities: models.Identities{
			UserID:     newUUID,
			Provider:   "getprint",
			ProviderID: newUUID.String(),
		},
		UserLog: models.UserLog{
			UserID: newUUID,
		},
	}

	err := testSuite.userRepo.Create(dummyPayload)

	assert.NoError(testSuite.T(), err)
	assert.NotZero(testSuite.T(), dummyPayload.Identities.ID)
	assert.NotZero(testSuite.T(), dummyPayload.UserLog.ID)
	testSuite.dataUser = dummyPayload
}
func (testSuite *MySQLUserTest) TestBcreateFailDuplicateEmail() {
	newUUID := uuid.NewV4()
	dummyPayload := &models.User{
		ID:            newUUID,
		Name:          "tesname",
		Email:         "test@test.com",
		EmailVerified: false,
		Role:          "buyer",
		Identities: models.Identities{
			UserID:     newUUID,
			Provider:   "getprint",
			ProviderID: newUUID.String(),
		},
		UserLog: models.UserLog{
			UserID: newUUID,
		},
	}

	err := testSuite.userRepo.Create(dummyPayload)

	assert.Error(testSuite.T(), err)
	assert.Equal(testSuite.T(), true, strings.Contains(err.Error(), "Duplicate"))
}

func (testSuite *MySQLUserTest) TestCgetUserByIDSuccess() {
	response := new(models.User)

	err := testSuite.userRepo.GetByID(testSuite.dataUser.ID.String(), response)

	assert.NoError(testSuite.T(), err)
	assert.Equal(testSuite.T(), testSuite.dataUser.ID.String(), response.ID.String())
	assert.Equal(testSuite.T(), testSuite.dataUser.Email, response.Email)
}

func (testSuite *MySQLUserTest) TestDgetUserByIDNotFound() {
	response := new(models.User)

	err := testSuite.userRepo.GetByID("", response)

	assert.Error(testSuite.T(), err)
	assert.Equal(testSuite.T(), true, strings.Contains(err.Error(), "not found"))
	assert.Equal(testSuite.T(), uuid.NullUUID{}.UUID, response.ID)
}

func (testSuite *MySQLUserTest) TestEgetUserByEmailSuccess() {
	response := new(models.User)
	response.Email = testSuite.dataUser.Email
	err := testSuite.userRepo.GetByEmail(response)

	assert.NoError(testSuite.T(), err)
	assert.Equal(testSuite.T(), testSuite.dataUser.ID.String(), response.ID.String())
	assert.Equal(testSuite.T(), testSuite.dataUser.Email, response.Email)
}

func (testSuite *MySQLUserTest) TestFgetUserByEmailNotFound() {
	response := new(models.User)
	err := testSuite.userRepo.GetByEmail(response)

	assert.Error(testSuite.T(), err)
	assert.Equal(testSuite.T(), true, strings.Contains(err.Error(), "not found"))
	assert.Equal(testSuite.T(), uuid.NullUUID{}.UUID, response.ID)
}

func (testSuite *MySQLUserTest) TestGgetUserByProviderIDSuccess() {
	response, err := testSuite.userRepo.GetByProviderID(testSuite.dataUser.Identities.ProviderID)

	assert.NoError(testSuite.T(), err)
	assert.Equal(testSuite.T(), testSuite.dataUser.ID.String(), response.ID.String())
	assert.Equal(testSuite.T(), testSuite.dataUser.Email, response.Email)
}

func (testSuite *MySQLUserTest) TestHgetUserByProviderIDNotFound() {
	response, err := testSuite.userRepo.GetByProviderID("tes")

	assert.Error(testSuite.T(), err)
	assert.Equal(testSuite.T(), true, strings.Contains(err.Error(), "not found"))
	assert.Nil(testSuite.T(), response)
}

func (testSuite *MySQLUserTest) TestIupdateUserLog() {
	userLog := testSuite.dataUser.UserLog
	err := testSuite.userRepo.UpdateUserLog(&userLog)

	assert.NoError(testSuite.T(), err)
	assert.NotEqual(testSuite.T(), testSuite.dataUser.UserLog.LastLogin.Time, userLog.LastLogin)
}
