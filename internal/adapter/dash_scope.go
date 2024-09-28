package adapter

import (
	"context"
	"github.com/devinyf/dashscopego"
	"github.com/devinyf/dashscopego/qwen"
	"log"
)

func DashScope(apiKey, model, text string) *dashscopego.TextQwenResponse {
	cli := dashscopego.NewTongyiClient(model, apiKey)

	content := qwen.TextContent{Text: text}

	input := dashscopego.TextInput{
		Messages: []dashscopego.TextMessage{
			{Role: qwen.RoleUser, Content: &content},
		},
	}

	// callback function:  print stream result
	streamCallbackFn := func(_ context.Context, chunk []byte) error {
		log.Print(string(chunk))
		return nil
	}
	req := &dashscopego.TextRequest{
		Input:       input,
		StreamingFn: streamCallbackFn,
	}

	ctx := context.TODO()

	resp, err := cli.CreateCompletion(ctx, req)
	if err != nil {
		panic(err)
	}
	return resp
}
