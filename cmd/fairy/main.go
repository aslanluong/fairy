package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aslanluong/fairy/internal/commands"
	"github.com/aslanluong/fairy/internal/config"
	"github.com/aslanluong/fairy/internal/listeners"
	"github.com/bwmarrin/discordgo"
)

func main() {
	err := config.LoadConfigsToEnv()
	if err != nil {
		fmt.Println("error loading configs to env,", err)
		fmt.Println("Using existing environment variables")
	}

	session, err := discordgo.New("Bot " + os.Getenv(config.EnvKeys.Bot.Token))
	if err != nil {
		fmt.Println("error creating discord session, ", err)
		return
	}
	defer session.Close()

	registerEvents(session)
	registerCommands(session, os.Getenv(config.EnvKeys.Bot.Prefix))

	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection, ", err)
	}

	fmt.Println("Bot is now running...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop
}

func registerEvents(s *discordgo.Session) {
	s.AddHandler(listeners.NewReadyListener().Handler)
	s.AddHandler(listeners.NewMemberAddListener().Handler)
	s.AddHandler(listeners.NewMemberRemoveListener().Handler)
}

func registerCommands(s *discordgo.Session, prefix string) {
	cmdHandler := commands.NewCommandHandler(prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		fmt.Printf("Command execution failed: %s\n", err.Error())
	}

	cmdHandler.RegisterCommand(&commands.CmdHelp{})
	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterCommand(&commands.CmdSay{})
	cmdHandler.RegisterCommand(&commands.CmdClear{})

	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
