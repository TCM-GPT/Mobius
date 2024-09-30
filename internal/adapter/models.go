package adapter

const (
	OPENAI = iota
	MODEL_SCOPE
	VOLCENGINE
)

// 偷个懒扒过来
const (
	O1Mini                = "o1-mini"
	O1Mini20240912        = "o1-mini-2024-09-12"
	O1Preview             = "o1-preview"
	O1Preview20240912     = "o1-preview-2024-09-12"
	GPT432K0613           = "gpt-4-32k-0613"
	GPT432K0314           = "gpt-4-32k-0314"
	GPT432K               = "gpt-4-32k"
	GPT40613              = "gpt-4-0613"
	GPT40314              = "gpt-4-0314"
	GPT4o                 = "gpt-4o"
	GPT4o20240513         = "gpt-4o-2024-05-13"
	GPT4o20240806         = "gpt-4o-2024-08-06"
	GPT4oLatest           = "chatgpt-4o-latest"
	GPT4oMini             = "gpt-4o-mini"
	GPT4oMini20240718     = "gpt-4o-mini-2024-07-18"
	GPT4Turbo             = "gpt-4-turbo"
	GPT4Turbo20240409     = "gpt-4-turbo-2024-04-09"
	GPT4Turbo0125         = "gpt-4-0125-preview"
	GPT4Turbo1106         = "gpt-4-1106-preview"
	GPT4TurboPreview      = "gpt-4-turbo-preview"
	GPT4VisionPreview     = "gpt-4-vision-preview"
	GPT4                  = "gpt-4"
	GPT3Dot5Turbo0125     = "gpt-3.5-turbo-0125"
	GPT3Dot5Turbo1106     = "gpt-3.5-turbo-1106"
	GPT3Dot5Turbo0613     = "gpt-3.5-turbo-0613"
	GPT3Dot5Turbo0301     = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo16K      = "gpt-3.5-turbo-16k"
	GPT3Dot5Turbo16K0613  = "gpt-3.5-turbo-16k-0613"
	GPT3Dot5Turbo         = "gpt-3.5-turbo"
	GPT3Dot5TurboInstruct = "gpt-3.5-turbo-instruct"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci003 = "text-davinci-003"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci002 = "text-davinci-002"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextCurie001 = "text-curie-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextBabbage001 = "text-babbage-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextAda001 = "text-ada-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci001 = "text-davinci-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	// Deprecated: Model is shutdown. Use davinci-002 instead.
	GPT3Davinci    = "davinci"
	GPT3Davinci002 = "davinci-002"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3CurieInstructBeta = "curie-instruct-beta"
	GPT3Curie             = "curie"
	GPT3Curie002          = "curie-002"
	// Deprecated: Model is shutdown. Use babbage-002 instead.
	GPT3Ada    = "ada"
	GPT3Ada002 = "ada-002"
	// Deprecated: Model is shutdown. Use babbage-002 instead.
	GPT3Babbage    = "babbage"
	GPT3Babbage002 = "babbage-002"
)

var OPENAI_MODELS = []string{
	O1Mini,
	O1Mini20240912,
	O1Preview,
	O1Preview20240912,
	GPT432K0613,
	GPT432K0314,
	GPT432K,
	GPT40613,
	GPT40314,
	GPT4o,
	GPT4o20240513,
	GPT4o20240806,
	GPT4oLatest,
	GPT4oMini,
	GPT4oMini20240718,
	GPT4Turbo,
	GPT4Turbo20240409,
	GPT4Turbo0125,
	GPT4Turbo1106,
	GPT4TurboPreview,
	GPT4VisionPreview,
	GPT4,
	GPT3Dot5Turbo0125,
	GPT3Dot5Turbo1106,
	GPT3Dot5Turbo0613,
	GPT3Dot5Turbo0301,
	GPT3Dot5Turbo16K,
	GPT3Dot5Turbo16K0613,
	GPT3Dot5Turbo,
	GPT3Dot5TurboInstruct,
	GPT3Curie,
	GPT3Curie002,
	GPT3Curie,
	GPT3Ada002,
	GPT3Babbage002,
}

// ModelScope 模型

var MODEL_SCOPE_MODELS = []string{
	// 通义千问 系列
	"qwen-long",
	"qwen-turbo",
	"qwen-turo-0624",
	"qwen-turo-0206",
	"qwen-plus",
	"qwen-plus-0806",
	"qwen-plus-0723",
	"qwen-plus-0624",
	"qwen-plus-0206",
	"qwen-max",
	"qwen-max-0428",
	"qwen-max-0403",
	"qwen-max-0107",
	"qwen-max-longcontext",

	// 通义千问VL系列
	"qwen-vl-max-0809",
	"qwen-vl-max-0201",
	"qwen-vl-max",
	"qwen-vl-plus",
	"qwen-vl-v1",
	"qwen-vl-chat-v1",

	// 通义千问开源系列
	"qwen2-math-72b-instruct",
	"qwen2-math-7b-instruct",
	"qwen2-math-1.5b-instruct",
	"qwen2-57b-a14b-instruct",
	"qwen2-72b-instruct",
	"qwen2-7b-instruct",
	"qwen2-1.5b-instruct",
	"qwen2-0.5b-instruct",
	"qwen1.5-110b-chat",
	"qwen1.5-72b-chat",
	"qwen1.5-32b-chat",
	"qwen1.5-14b-chat",
	"qwen1.5-7b-chat",
	"qwen1.5-1.8b-chat",
	"qwen1.5-0.5b-chat",
	"codeqwen1.5-7b-chat",
	"qwen-72b-chat",
	"qwen-14b-chat",
	"qwen-7b-chat",
	"qwen-1.8b-longcontext-chat",
	"qwen-1.8b-chat",
}
