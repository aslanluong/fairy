package commands

import (
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

	if args[len(args)-1] != "-hide" {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, strings.Join(args, " "))
		return
	}

	if err = ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID); err == nil {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, strings.Join(args[:len(args)-1], " "))
	}
	return
}
