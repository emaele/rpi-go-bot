package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// Config is the bot configuration representation, read
// from a configuration file.
type Config struct {
	TelegramTokenBot string
	MyID             int64
	TempLimit        float64
	Pihole           bool
	PiholeHost       string
	PiholeAPIToken   string
}

// ReadConfig loads the values from the config file
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
	} else if conf.Pihole {
		if conf.PiholeHost == "" {
			return newErr("missing pihole host")
		} else if conf.PiholeAPIToken == "" {
			return newErr("missing pihole API token")
		}
	}
	return conf, nil
}

func newErr(message string) (Config, error) {
	return Config{}, errors.New(message)
}
