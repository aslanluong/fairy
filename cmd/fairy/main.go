package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aslanluong/fairy/internal/config"
	"github.com/aslanluong/fairy/internal/events"
	"github.com/bwmarrin/discordgo"
)

func main() {
	err := config.LoadConfigsToEnv()
	if err != nil {
		fmt.Println("error loading configs to env,", err)
		fmt.Println("Using existing environment variables")
	}

	discord, err := discordgo.New("Bot " + os.Getenv(config.EnvKeys.Bot.Token))
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
