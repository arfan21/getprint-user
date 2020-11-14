package services

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arfan21/getprint-user/models"
	"github.com/arfan21/getprint-user/repository"
	"github.com/arfan21/getprint-user/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func loadEnv() {
	rootPath, err := os.Getwd()

	err = godotenv.Load(os.ExpandEnv(filepath.Dir(rootPath) + "/.env"))

	if err != nil {
		log.Fatalf("can't load env file : %v", err)
	}
}

func TestCreate(t *testing.T) {
	loadEnv()

	db, err := utils.Connect(os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	reqBody := models.User{
		Name:  "arfan",
		Email: "arfan2@email.com",
	}
	reqBody.UserIDLine.Scan("U82kakhihqwken281lma9i21221")
	reqBody.Picture.Scan("https://github.com/gm.jpg")
	reqBody.Password.Scan("password")
	reqBody.PhoneNumber.Scan(62821363121)
	reqBody.Address.Scan("jakarta, indonesia")

	reqBodyJSON, _ := json.Marshal(reqBody)

	args := models.User{}

	_ = json.Unmarshal(reqBodyJSON, &args)

	userRepo := repository.NewMysqlUserRepository(db)
	userService := NewUserServices(userRepo)

	err = userService.Create(&args)
	assert.NoError(t, err)
	assert.NotZero(t, args.ID, "ID is zero")
	assert.Equal(t, "buyer", args.Role)
	assert.Equal(t, reqBody.Name, args.Name)
	assert.Equal(t, reqBody.Email, args.Email)
}

func TestGet(t *testing.T) {
	loadEnv()

	db, err := utils.Connect(os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	userRepo := repository.NewMysqlUserRepository(db)
	userService := NewUserServices(userRepo)

	args := &[]models.User{}

	err = userService.Get(args)

	if err != nil {
		t.Fatalf("%v", err)
	}

	log.Println(args)
}

func TestGetByID(t *testing.T) {
	loadEnv()

	db, err := utils.Connect(os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	args := &models.User{}

	userRepo := repository.NewMysqlUserRepository(db)
	userService := NewUserServices(userRepo)

	err = userService.GetByID(1, args)

	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.NotZero(t, args.ID, "ID is zero")
	assert.Equal(t, "", args.Password.String, "password leaked")
	assert.NotNil(t, args.Name)
	assert.NotNil(t, args.Email)
}

func TestUpdate(t *testing.T) {
	loadEnv()

	db, err := utils.Connect(os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	userRepo := repository.NewMysqlUserRepository(db)
	userService := NewUserServices(userRepo)

	reqBody := &models.User{}
	reqBody.Password.Scan("password1")

	args := &models.User{}
	err = userService.GetByID(1, args)

	if err != nil {
		t.Fatalf("%v", err)
	}

	args.Picture.Scan("https://github.com/gmgm2.jpg")
	err = userService.Update(args)

	assert.Equal(t, "https://github.com/gmgm2.jpg", args.Picture.String, "not equal")
}