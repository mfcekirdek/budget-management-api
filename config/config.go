package config

import (
	"fmt"
	"github.com/k0kubun/pp"
)
import "github.com/spf13/viper"

type Config struct {
	IsDebug   bool
	AppName   string
	Server    ServerConfig
	Couchbase CouchbaseConfig
}

func New(configPath, filename string) (*Config, error) {
	v, err := readConfig(configPath, filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	config := &Config{}
	err = v.Unmarshal(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return config, nil
}

func readConfig(configPath, filename string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(filename)
	err := v.ReadInConfig()
	return v, err
}

func (c *Config) Print() {
	_, _ = pp.Println(c)
}
