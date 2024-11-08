package controller

import (
	"fmt"
	"gen-SFT-Dataset/internal/adapter"
	"gen-SFT-Dataset/internal/process"
	"github.com/devinyf/dashscopego"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	m "github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"log"
)

// Generate 通过调用厂商接口将用户传入内容 content 生成为 SFT 格式的训练数据集
func Generate(ctx *gin.Context) {
	model, ok := ctx.Get("model")
	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "model param not found"})
		return
	}

	contentList, ok := ctx.Get("contentList")
	if !ok {
		ctx.JSON(400, gin.H{"error": "contentList not found"})
		return
	}

	for _, content := range contentList.([]string) {
		resp, err := adapter.NewRequest(content, model.(string))
		if resp == nil || err != nil {
			log.Println("got no response while request LLM api")
			ctx.JSON(400, gin.H{
				"error": "got no response while request LLM api",
				"info":  err,
			})
			return
		}

		var generateContext string

		// 根据使用的 model 选择不同的发送请求方式
		switch model.(string) {
		case adapter.Contains(adapter.OPENAI_MODELS, model.(string)):
			generateContext = resp.(*openai.ChatCompletionResponse).Choices[0].Message.Content
			break
		case adapter.Contains(adapter.MODEL_SCOPE_MODELS, model.(string)):
			choices := resp.(*dashscopego.TextQwenResponse).GetChoices()
			generateContext = fmt.Sprintf("%s", choices[0].Message.Content)
			break
		case func(model string) string {
			splitString := model[0:3]
			if splitString == "ep-" {
				return model
			} else {
				return ""
			}
		}(model.(string)):
			// 执行火山引擎的接口
			a := resp.(*m.ChatCompletionResponse).Choices[0].Message.Content.StringValue // 返回值 StringValue 类型 *string...
			fmt.Println(*a)
			generateContext = fmt.Sprintf("%s", *resp.(*m.ChatCompletionResponse).Choices[0].Message.Content.StringValue)
		}

		// 转换并写入到数据库中
		err = process.Convert(generateContext, model.(string))
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error})
			return
		}
	}
	ctx.JSON(200, gin.H{
		"message":     "success",
		"data_length": len(contentList.([]string)),
	})
}
