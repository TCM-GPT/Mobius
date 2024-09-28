package middleware

import (
	"gen-SFT-Dataset/internal/process"
	"github.com/gin-gonic/gin"
	"net/http"
)

type input struct {
	SplitPrefix string `form:"split_prefix" json:"split_prefix" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Restrict    uint64 `form:"restrict" json:"restrict" binding:"required"`
	Model       string `form:"model" json:"model" binding:"required"`
}

// Split 处理文本内容拆分
func Split() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// bind input data
		var in input
		if err := ctx.ShouldBind(&in); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"input_param_error": err.Error(),
			})
			return
		}

		// process content
		contentList := process.Split(in.Content, in.SplitPrefix, in.Restrict)
		ctx.Set("contentList", contentList)
		ctx.Set("model", in.Model)

		ctx.Next()
	}
}
