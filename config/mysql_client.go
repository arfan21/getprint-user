package config

import (
	"log"
	"strings"

	"github.com/arfan21/getprint-user/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client interface {
	Conn() *gorm.DB
	Close() error
}

func NewMySQLClient(DBURL string) (Client, error) {
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
