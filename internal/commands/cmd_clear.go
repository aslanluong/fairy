package commands

import (
	"fmt"
	"strconv"
	"time"
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
	if err != nil || limit < 1 || limit > 100 {
		if message, fail := ctx.Session.ChannelMessageSendReply(ctx.Message.ChannelID, "Invalid amount (1-100), please try again!", ctx.Message.Reference()); fail == nil {
			time.AfterFunc(3*time.Second, func() {
				ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, message.ID)
			})
		}
		return
	}

	messages, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, limit, ctx.Message.ID, "", "")
	if err != nil {
		return
	}

	msgIds := make([]string, 0)
	for _, message := range messages {
		msgIds = append(msgIds, message.ID)
	}
	err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, msgIds)
	if err != nil {
		return
	}

	notiContent := "I've cleared that message."
	if limit > 1 {
		notiContent = fmt.Sprintf("I've cleared %d messages.", limit)
	}
	noti, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, notiContent)
	if err != nil {
		return
	}

	time.AfterFunc(3*time.Second, func() {
		err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, []string{ctx.Message.ID, noti.ID})
	})
	return
}