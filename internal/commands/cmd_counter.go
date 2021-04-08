package commands

import (
	"time"

	"github.com/aslanluong/fairy/internal/listeners"
	"github.com/aslanluong/fairy/internal/utils"
)

type CmdCounter struct {
}

func (c *CmdCounter) Invokes() []string {
	return []string{"counter"}
}

func (c *CmdCounter) Description() string {
	return "Enable/disable counter listener"
}

func (c *CmdCounter) AdminRequired() bool {
	return true
}

func (c *CmdCounter) Exec(ctx *Context) (err error) {
	args := ctx.Args

	if args[0] != "enable" && args[0] != "disable" {
		if msg, fail := utils.SendMessageReply(ctx.Session, ctx.Message.Reference(), "Invalid state (enable/disable)"); fail == nil {
			msg.WaitAndDelete(3 * time.Second)
		}
		return
	}

	if args[0] == "enable" {
		listeners.CounterChannel().EnableListener(ctx.Message.ChannelID)
		if err = utils.DeleteMessage(ctx.Session, ctx.Message.ChannelID, ctx.Message.ID); err == nil {
			if msg, fail := utils.SendMessage(ctx.Session, ctx.Message.ChannelID, "Counter enabled!"); fail == nil {
				msg.WaitAndDelete(3 * time.Second)
			}
		}
	}
	if args[0] == "disable" {
		listeners.CounterChannel().DisableListener(ctx.Message.ChannelID)
		if err = utils.DeleteMessage(ctx.Session, ctx.Message.ChannelID, ctx.Message.ID); err == nil {
			if msg, fail := utils.SendMessage(ctx.Session, ctx.Message.ChannelID, "Counter disabled!"); fail == nil {
				msg.WaitAndDelete(3 * time.Second)
			}
		}
	}
	return
}
