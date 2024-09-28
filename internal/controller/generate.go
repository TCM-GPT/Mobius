package controller

import (
	"gen-SFT-Dataset/config"
	"gen-SFT-Dataset/internal/adapter"
	"gen-SFT-Dataset/internal/process"
	"github.com/gin-gonic/gin"
	"log"
)

var prompt = `
你是一个能根据提供的文本内容生成QA对的机器人。以下是你的任务要求：
1. 生成尽可能多的QA对，答案尽可能的详细但是保证不偏离问题。
2. 每个QA对包含一个问题和一个简洁的答案。
3. 答案必须用简体中文。
4. 生成的QA对不能重复。
5. 使用json格式将QA对包裹起来，问题用"instruction"表示，"input" (可选) 是对instruction的补充，提供了完整的上下文和指导，用于生成输出。答案用"output"表示。

示例格式：
[
	{
		"instruction": "...",
		"input": "...",
		"output": "..."
	},
	{
		"instruction": "...",
		"input": "...",
		"output": "..."
	}
]
以下是给定的文本内容：
`

// Generate 通过调用 openai Api 将文档生成为 SFT 格式的训练数据集
func Generate(ctx *gin.Context) {
	contentList, ok := ctx.Get("contentList")
	if !ok {
		ctx.JSON(400, gin.H{"error": "contentList not found"})
		return
	}

	for _, content := range contentList.([]string) {
		resp := adapter.OpenAI(
			config.GetInstance().AppConfig.Adapters[0].ApiKey,
			config.GetInstance().AppConfig.Adapters[0].Model,
			prompt+content,
			config.GetInstance().AppConfig.Adapters[0].Temperature,
			config.GetInstance().AppConfig.Adapters[0].TopP,
			false,
		)
		if resp == nil {
			log.Println("resp is nil")
			ctx.JSON(400, gin.H{"error": "resp is nil"})
			return
		}

		generateContext := resp.Choices[0].Message.Content
		err := process.Convert(generateContext)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message":     "success",
			"data_length": len(contentList.([]string)),
		})
	}
}
