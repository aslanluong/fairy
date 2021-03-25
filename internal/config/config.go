package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Keys struct {
	Bot     Bot
	Mongodb Mongodb
}

type Bot struct {
	Token  string
	Prefix string
}

type Mongodb struct {
	Uri string
}

var EnvKeys = Keys{
	Bot: Bot{
		Token:  "bot.token",
		Prefix: "bot.prefix",
	},
	Mongodb: Mongodb{
		Uri: "mongodb.uri",
	},
}

func LoadConfigsToEnv() (err error) {
	env, err := os.ReadFile("./env.yml")
	if err != nil {
		return
	}

	var envVars map[string]interface{}
	if err := yaml.Unmarshal(env, &envVars); err != nil {
		fmt.Println("error unmarshalling env vars,", err)
	} else {
		var setEnvVars func(rootKey string, values map[string]interface{})
		setEnvVars = func(rootKey string, values map[string]interface{}) {
			for key, value := range values {
				_, ok := value.(map[interface{}]interface{})
				if ok {
					valuesMarshal, _ := yaml.Marshal(&value)
					var valuesUnmarshal map[string]interface{}
					yaml.Unmarshal(valuesMarshal, &valuesUnmarshal)
					setEnvVars(rootKey+key+".", valuesUnmarshal)
				} else {
					if valueStr, ok := value.(string); ok {
						os.Setenv(rootKey+key, valueStr)
					}
					if valueInt, ok := value.(int); ok {
						os.Setenv(rootKey+key, strconv.Itoa(valueInt))
					}
				}
			}
		}
		setEnvVars("", envVars)
	}
	return
}
