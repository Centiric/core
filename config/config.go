package config

import "github.com/spf13/viper"

type Config struct {
	Grpc struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"grpc"`
	Services struct {
		Media struct {
			Address string `mapstructure:"address"`
		} `mapstructure:"media"`
	} `mapstructure:"services"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
