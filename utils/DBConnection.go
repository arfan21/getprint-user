package utils

import (
	"fmt"
	"log"

	"github.com/arfan21/getprint-user/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME string) (*gorm.DB, error) {
	var DBURL string

	if DB_PASSWORD == "" {
		DBURL = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_HOST, DB_PORT, DB_NAME)
	} else {
		DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	}
	var err error
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	log.Println("MySql Connected")
	return db, nil
}
