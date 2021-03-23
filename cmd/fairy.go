package main

import (
	"fairy/internal/events"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

func main() {
	LoadEnvVars()

	discord, err := discordgo.New("Bot " + os.Getenv("bot.token"))
	if err != nil {
		fmt.Println("error creating discord session, ", err)
		return
	}
	defer discord.Close()

	discord.AddHandler(events.MessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection, ", err)
	}

	fmt.Println("Bot is now running...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop
}

func LoadEnvVars() {
	env, err := ioutil.ReadFile("./env.yml")
	if err != nil {
		return
	}

	var envVars map[string]interface{}
	if err := yaml.Unmarshal(env, &envVars); err != nil {
		fmt.Println("fail", err)
	} else {
		var setEnvVars func(rootKey string, values map[string]interface{})
		setEnvVars = func(rootKey string, values map[string]interface{}) {
			for key, value := range values {
				_, ok := value.(map[interface{}]interface{})
				if ok {
					valuesMarshal, _ := yaml.Marshal(value)
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
}
