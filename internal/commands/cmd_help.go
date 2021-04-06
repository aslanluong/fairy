package commands

import (
	"time"
)

type CmdHelp struct {
}

func (c *CmdHelp) Invokes() []string {
	return []string{"help"}
}

func (c *CmdHelp) Description() string {
	return "Get help about command."
}

func (c *CmdHelp) AdminRequired() bool {
	return false
}

func (c *CmdHelp) Exec(ctx *Context) (err error) {
	args := ctx.Args

	cmd := ctx.Handler.cmdMap[args[0]]
	if cmd == nil {
		if message, fail := ctx.Session.ChannelMessageSendReply(ctx.Message.ChannelID, "Invalid command name, please try again!", ctx.Message.Reference()); fail == nil {
			time.AfterFunc(3*time.Second, func() {
				ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, message.ID)
			})
		}
		return
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, cmd.Description())
	return
}
