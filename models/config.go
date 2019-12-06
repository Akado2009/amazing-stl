package models

import (
	"github.com/BurntSushi/toml"
	"os"
)

var (
	configPath = os.Getenv("CONFIG_TOML")

	//AppConfig global instance
	AppConfig Config
)

//Config struct
type Config struct {
	Application Application `toml:"Application"`
	Email Email `toml:"Email"`
}

func LoadConfig(c *Config) error {
	_, err := toml.DecodeFile(configPath, c)
	return err
}
