package commands

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CmdSay struct {
}

func (c *CmdSay) Invokes() []string {
	return []string{"say", "s"}
}

func (c *CmdSay) Description() string {
	return "Say something!"
}

func (c *CmdSay) AdminRequired() bool {
	return false
}

func (c *CmdSay) Exec(ctx *Context) (err error) {
	args := ctx.Args

	f := flag.NewFlagSet("sayflags", flag.ContinueOnError)
	fHide := f.Bool("hide", false, "hide original message")
	fEmbed := f.Bool("embed", false, "send message with embed")
	if err = f.Parse(args); err != nil {
		return errors.New("invalid flag(s) provided")
	}

	if !*fHide {
		if *fEmbed {
			_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID,
				&discordgo.MessageEmbed{Title: fmt.Sprintf("%s said:", ctx.Message.Author), Description: strings.Join(f.Args(), " ")})
			return
		}
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, strings.Join(f.Args(), " "))
		return
	}

	errs := make(chan error, 1)

	if *fEmbed {
		go func() {
			_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID,
				&discordgo.MessageEmbed{Title: fmt.Sprintf("%s said:", ctx.Message.Author), Description: strings.Join(f.Args(), " ")})

			errs <- err
		}()
		go func() {
			err = ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
			errs <- err
		}()
		<-errs
		return
	}

	go func() {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, strings.Join(f.Args(), " "))
		errs <- err
	}()
	go func() {
		err = ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
		errs <- err
	}()
	<-errs
	return
}
