package controllers

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func loadEnv() {
	rootPath, err := os.Getwd()

	err = godotenv.Load(os.ExpandEnv(filepath.Dir(rootPath) + "/.env"))

	if err != nil {
		log.Fatalf("can't load env file : %v", err)
	}
}

func TestCreate(t *testing.T) {

}
