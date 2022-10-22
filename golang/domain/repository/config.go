package repository

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName  string `default:"ami-back"`
	Port     string `default:"8000"`
	Use      string `default:"postgres"`
	Enabled  bool   `default:"true"`
	Host     string `mapstructure:"DATABASE_HOST"`
	DB_Port  string `mapstructure:"DATABASE_PORT"`
	UserName string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Database string `mapstructure:"POSTGRES_DATABASE_NAME"`
}

func (c *Config) NewConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
