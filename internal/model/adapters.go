package model

type Adapter struct {
	Type        string  `mapstructure:"type" yaml:"type"` // 模型厂商 openai、modelScope
	ApiKey      string  `mapstructure:"api_key" yaml:"api_key"`
	Model       string  `mapstructure:"model" yaml:"model"`
	Temperature float32 `mapstructure:"temperature" yaml:"temperature"`
	TopP        float32 `mapstructure:"top_p" yaml:"top_p"`
}
