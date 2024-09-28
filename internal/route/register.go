package route

import (
	"gen-SFT-Dataset/internal/controller"
	"gen-SFT-Dataset/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.POST("/api/v1/generate", middleware.Split(), controller.Generate)
}
