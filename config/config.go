package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TelegramTokenBot string
	MyID             int64
	TempLimit        float64
	PiholeHost       string
	PiholeAPIToken   string
}

func ReadConfig(path string) (Config, error) {
	var conf Config

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return Config{}, err
	}

	if conf.TelegramTokenBot == "" {
		return newErr("missing Bot token")
	} else if conf.MyID == 0 {
		return newErr("missing ID")
	} else if conf.TempLimit == 0 {
		return newErr("missing temperature limit value")
	} else if conf.PiholeHost == "" {
		return newErr("missing pihole host")
	} else if conf.PiholeAPIToken == "" {
		return newErr("missing pihole API token")
	}

	return conf, nil
}

func newErr(message string) (Config, error) {
	return Config{}, errors.New(message)
}
