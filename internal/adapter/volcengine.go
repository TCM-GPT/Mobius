package adapter

import (
	"context"
	"gen-SFT-Dataset/config"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"log"
)

// Volcengine TODO waiting to finish this function.
func Volcengine(content, models string) (*model.ChatCompletionResponse, error) {
	cli := arkruntime.NewClientWithApiKey(config.GetInstance().AppConfig.Adapters[VOLCENGINE].ApiKey)

	ctx := context.Background()
	req := model.ChatCompletionRequest{
		Model: models,
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("你是豆包，是由字节跳动开发的 AI 人工智能助手"),
				},
			},
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(content),
				},
			},
		},
	}
	resp, err := cli.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Printf("vloc api request error: %v", err)
		return nil, err
	}

	return &resp, nil
}
