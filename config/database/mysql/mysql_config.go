package mysql

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func NewMySQLConfig() *MySQLConfig {
	dbConfig := MySQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func NewMySQLConfigForTest() *MySQLConfig {
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
