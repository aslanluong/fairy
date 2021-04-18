package core

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
