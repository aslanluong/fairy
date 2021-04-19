package core

import (
	"io"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Bot     *Bot
	Mongodb *Mongodb
}

type Bot struct {
	Token  string
	Prefix string
}

type Mongodb struct {
	Uri string
}

type YAMLConfigParser struct{}

func (y *YAMLConfigParser) Encode(r io.Reader) (cfg *Config, err error) {
	decoder := yaml.NewDecoder(r)
	cfg = new(Config)
	err = decoder.Decode(cfg)
	return
}

func DefaultConfig() *Config {
	return &Config{
		Bot: &Bot{
			Token:  "",
			Prefix: "!",
		},
		Mongodb: &Mongodb{
			Uri: "",
		},
	}
}
