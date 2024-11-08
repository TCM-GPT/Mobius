package model

type Postgres struct {
	Host     string `mapstructure:"host" yaml:"host"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	DbName   string `mapstructure:"db_name" yaml:"db_name"`
	DbPort   string `mapstructure:"db_port" yaml:"db_port"`
}
