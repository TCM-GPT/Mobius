package model

type Mysql struct {
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Host     string `mapstructure:"host" yaml:"host"`
	DbName   string `mapstructure:"db_name" yaml:"db_name"`
}
