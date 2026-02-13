package test

import (
	"log"
	"os"
	"template/internal/config"
	"template/internal/database"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var TestConfig *config.Config
var App *gin.Engine
var DB *sqlx.DB

var ClientDetails = struct {
}{}

func TestMain(m *testing.M) {
	TestConfig, err := config.NewConfigTest()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.NewMySQL(TestConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db

	App, err = InitializeServer()
	if err != nil {
		log.Fatal("Failed to Run Injector", err)
	}

	code := m.Run()
	os.Exit(code)
}
