package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arfan21/getprint-user/app/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client interface {
	Conn() *gorm.DB
	Close() error
}

func NewClient(DBURL string) (Client, error) {
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		if strings.Contains(err.Error(), "Unknown") {
			DBURLSplitted := strings.Split(DBURL, "/")
			newSplit := strings.Split(DBURLSplitted[1], "?")
			dbName := newSplit[0]
			newDBURL := DBURLSplitted[0] + "/?" + newSplit[1]

			db, err = gorm.Open(mysql.Open(newDBURL), &gorm.Config{})
			if err != nil {
				return nil, err
			}

			db.Debug().Exec("CREATE DATABASE " + dbName)
			db.Debug().Exec("USE " + dbName)
		} else {
			return nil, err
		}

	}

	err = db.AutoMigrate(&models.User{}, &models.Identities{}, &models.UserLog{})
	if err != nil {

		return nil, err
	}
	log.Println("MySql Connected")

	return &client{db}, nil
}

type client struct {
	db *gorm.DB
}

func (c *client) Conn() *gorm.DB {
	return c.db
}

func (c *client) Close() error {
	sqlDB, err := c.db.DB()
	if err == nil {
		return err
	}
	return sqlDB.Close()
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func NewConfig() *MySQLConfig {
	dbConfig := MySQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func NewConfigForTest() *MySQLConfig {
	dbConfig := MySQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME_TEST"),
	}
	return &dbConfig
}

func (dbConfig *MySQLConfig) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
