package cmd

import (
	"gen-SFT-Dataset/config"
	"gen-SFT-Dataset/internal/route"
	"gen-SFT-Dataset/mysql"
	"gen-SFT-Dataset/postgresql"
	"github.com/gin-gonic/gin"
)

type server struct {
	Engine *gin.Engine
}

func NewServer() *server {
	s := server{}
	s.initAll()
	return &s
}

func (s *server) initAll() {
	config.InitConfigFile()
	switch config.GetInstance().AppConfig.UseDb {
	case "mysql":
		mysql.InitMysqlSetup()
	case "postgresql":
		postgresql.InitPostgresSetup()
	default:
		postgresql.InitPostgresSetup()
	}
	s.InitGin()
}

func (s *server) InitGin() {
	s.Engine = gin.Default()
	route.Register(s.Engine)
	s.Engine.Run(":8080")
}
