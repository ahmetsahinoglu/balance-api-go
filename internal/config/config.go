package config

import (
	"github.com/spf13/viper"
	"github.com/yudai/pp"
)

type (
	Config struct {
		AppName  string
		Server   Server
		LogLevel string
		Clients  Clients
	}

	Clients struct {
		Eth string
	}

	Server struct {
		Port string
	}
)

func New(configPath, configName string) (Config, error) {
	viperConfig, err := readConfig(configPath, configName)
	if err != nil {
		return Config{}, err
	}
	var config Config

	if err := viperConfig.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func readConfig(configPath, configName string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	err := v.ReadInConfig()

	return v, err
}

func (c Config) Print() {
	_, _ = pp.Println(c)
}
