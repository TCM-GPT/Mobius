package mysql

import (
	"fmt"
	"gen-SFT-Dataset/config"
	"gen-SFT-Dataset/internal/model"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"

var db *gorm.DB

func GetInstance() *gorm.DB {
	return db
}

func InitMysqlSetup() {
	DB, err := gorm.Open(mysql.New(mysql.Config{DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetInstance().AppConfig.Mysql.User,
		config.GetInstance().AppConfig.Mysql.Password,
		config.GetInstance().AppConfig.Mysql.Host,
		config.GetInstance().AppConfig.Mysql.DbName),
		DisableDatetimePrecision:  true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		log.Println("mysql connect failed")
		panic(err)
		return
	}

	err = DB.AutoMigrate(&model.Dataset{})
	if err != nil {
		log.Println("mysql migrate failed")
		return
	}

	db = DB
}
