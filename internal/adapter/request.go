package adapter

import (
	"gen-SFT-Dataset/config"
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

// NewRequest 根据 model 的厂商发起一个 post 请求
func NewRequest(content string, model string) (resp any, err error) {
	var defaultModel string

	if len(model) > 0 {
		switch model {
		case Contains(OPENAI_MODELS, model):
			// 发起 openai 请求
			resp = OpenAI(
				config.GetInstance().AppConfig.Adapters[0].ApiKey,
				model,
				prompt+content,
				config.GetInstance().AppConfig.Adapters[0].Temperature,
				config.GetInstance().AppConfig.Adapters[0].TopP,
				false,
			)
			break
		case Contains(MODEL_SCOPE_MODELS, model):
			// 发起 model scope 请求
			resp = DashScope(
				config.GetInstance().AppConfig.Adapters[1].ApiKey,
				model,
				prompt+content,
			)
			break
		case func(model string) string {
			splitString := model[0:3]
			if splitString == "ep-" {
				return model
			} else {
				return ""
			}
		}(model):
			// 发起火山引擎接口请求
			resp, err = Volcengine(prompt + content)
			if err != nil {
				log.Println("Volcengine request error:", err)
				return nil, err
			}
			return resp, nil
			break
		default:
			log.Println("不支持的模型:", model)
			break
		}
	} else {
		// 如果用户不传入则读取配置文件 config.yml 的 model 发起请求
		if defaultModel = config.GetInstance().AppConfig.Adapters[0].Model; defaultModel != "" {
			resp = OpenAI(
				config.GetInstance().AppConfig.Adapters[0].ApiKey,
				defaultModel,
				prompt+content,
				config.GetInstance().AppConfig.Adapters[0].Temperature,
				config.GetInstance().AppConfig.Adapters[0].TopP,
				false,
			)
			return resp, nil
		}
	}
	return resp, nil
}

func Contains(slice []string, value string) string {
	for _, v := range slice {
		if v == value {
			return value
		}
	}
	return ""
}
