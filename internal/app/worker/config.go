package worker

import (
	"github.com/spf13/viper"
)

type Config struct {
	Worker struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"worker"`

	Netbox struct {
		URL   string `mapstructure:"url"`
		Token string `mapstructure:"token"`
	} `mapstructure:"netbox"`
}

func NewConfigFromPath(configFilePath string) (Config, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
