package adapter

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
)

func OpenAI(apiKey, model, content string, temperature, topP float32, stream bool) *openai.ChatCompletionResponse {
	cli := openai.NewClient(apiKey)
	resp, err := cli.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
			Stream:      stream,
			Temperature: temperature,
			TopP:        topP,
		},
	)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return nil
	}
	return &resp
}
