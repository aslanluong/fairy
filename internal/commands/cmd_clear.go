package commands

import (
	"strconv"
)

type CmdClear struct {
}

func (c *CmdClear) Invokes() []string {
	return []string{"clear"}
}

func (c *CmdClear) Description() string {
	return "Clear messages"
}

func (c *CmdClear) AdminRequired() bool {
	return true
}

func (c *CmdClear) Exec(ctx *Context) (err error) {
	args := ctx.Args

	limit, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	messages, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, limit, "", "", "")
	if err != nil {
		return
	}

	msgIds := make([]string, 0)
	for _, message := range messages {
		msgIds = append(msgIds, message.ID)
	}
	err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, msgIds)

	return
}
