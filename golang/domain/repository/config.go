package repository

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"SERVER_PORT"`
	Host     string `mapstructure:"DATABASE_HOST"`
	DB_Port  string `mapstructure:"DATABASE_PORT"`
	UserName string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Database string `mapstructure:"POSTGRES_DATABASE_NAME"`
}

var cfg *Config

func (c *Config) NewConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) GetConfig() (*Config, error) {

	return cfg, nil
}
