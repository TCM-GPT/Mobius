package cmd

import (
	"gen-SFT-Dataset/config"
	"gen-SFT-Dataset/internal/route"
	"gen-SFT-Dataset/mysql"
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
	mysql.InitMysqlSetup()
	s.InitGin()
}

func (s *server) InitGin() {
	s.Engine = gin.Default()
	route.Register(s.Engine)
	s.Engine.Run(":8080")
}
