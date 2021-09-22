package app

import (
	"fmt"
	"log"

	"github.com/imufiid/goabsen/app/core"
	"github.com/imufiid/goabsen/app/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

func initialize(appConfig *AppConfig, dbConfig *DBConfig) {
	// connecting db
	db := initializeDB(dbConfig)

	repository := core.InitRepository(db)
	service := core.InitService(repository)
	handler := core.InitHandler(service)
	
	// initialize router
	addrs := fmt.Sprintf(":%s", appConfig.Port)
	router := NewRouter(addrs, handler)
	router.Router()

	fmt.Println("Initialize on port ", appConfig.Port)
}

func initializeDB(dbConfig *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func Run() {
	appConfig := &AppConfig{
		Port: helper.GetENV("APP_PORT", "8080"),
	}

	dbConfig := &DBConfig{
		Host:     helper.GetENV("DB_HOST", "localhost"),
		Name:     helper.GetENV("DB_NAME", "goabsen"),
		User:     helper.GetENV("DB_USER", "postgres"),
		Password: helper.GetENV("DB_PASSWORD", "root"),
		Port:     helper.GetENV("DB_PORT", "5432"),
	}

	initialize(appConfig, dbConfig)
}
