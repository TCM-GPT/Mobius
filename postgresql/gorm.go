package postgresql

import (
	"fmt"
	"gen-SFT-Dataset/config"
	"gen-SFT-Dataset/internal/model"
	"gorm.io/gorm"
	"log"
)

import "gorm.io/driver/postgres"

var db *gorm.DB

func GetInstance() *gorm.DB {
	return db
}

func InitPostgresSetup() {
	DB, err := gorm.Open(postgres.New(postgres.Config{DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.GetInstance().AppConfig.Postgresql.Host,
		config.GetInstance().AppConfig.Postgresql.User,
		config.GetInstance().AppConfig.Postgresql.Password,
		config.GetInstance().AppConfig.Postgresql.DbName,
		config.GetInstance().AppConfig.Postgresql.DbPort),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Println("postgresql connect failed")
		panic(err)
		return
	}

	err = DB.AutoMigrate(&model.Dataset{})
	if err != nil {
		log.Println("postgresql migrate failed")
		return
	}

	db = DB
}
