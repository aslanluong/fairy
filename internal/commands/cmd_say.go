package commands

import (
	"errors"
	"flag"
	"strings"
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
	if err = f.Parse(args); err != nil {
		return errors.New("invalid flag(s) provided")
	}

	if !*fHide {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, strings.Join(f.Args(), " "))
		return
	}

	if err = ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID); err == nil {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, strings.Join(f.Args(), " "))
	}
	return
}
