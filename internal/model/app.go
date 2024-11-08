package model

type App struct {
	AppConfig AppConfig `mapstructure:"app_config" yaml:"app_config"`
}

type AppConfig struct {
	Name        string    `mapstructure:"name" yaml:"name"`
	Version     string    `mapstructure:"version" yaml:"version"`
	Description string    `mapstructure:"description" yaml:"description"`
	Adapters    []Adapter `mapstructure:"adapters" yaml:"adapters"`
	UseDb       string    `mapstructure:"useDb" yaml:"useDb"`
	Mysql       Mysql     `mapstructure:"mysql" yaml:"mysql"`
	Postgresql  Postgres  `mapstructure:"postgresql" yaml:"postgresql"`
}
