package model

type Adapter struct {
	Type        string  `mapstructure:"type" yaml:"type"` // 模型厂商 openai、modelScope
	ApiKey      string  `mapstructure:"api_key" yaml:"api_key"`
	Model       string  `mapstructure:"model" yaml:"model"`
	Host        string  `mapstructure:"host" yaml:"host"`
	Region      string  `mapstructure:"region" yaml:"region"`
	Temperature float32 `mapstructure:"temperature" yaml:"temperature"`
	TopP        float32 `mapstructure:"top_p" yaml:"top_p"`
}
