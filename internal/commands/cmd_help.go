package commands

import (
	"time"

	"github.com/aslanluong/fairy/internal/utils"
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

	cmd, ok := ctx.Handler.cmdMap[args[0]]
	if !ok {
		if msg, fail := utils.SendMessageReply(ctx.Session, ctx.Message.Reference(), "Invalid command name, please try again!"); fail == nil {
			msg.DeleteAfter(3 * time.Second)
		}
		return
	}

	_, err = utils.SendMessage(ctx.Session, ctx.Message.ChannelID, cmd.Description())
	return
}
