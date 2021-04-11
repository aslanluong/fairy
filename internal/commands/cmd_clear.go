package commands

import (
	"errors"
	"flag"
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

	f := flag.NewFlagSet("clearflags", flag.ContinueOnError)
	fYou := f.Bool("you", false, "bot messages filtered")
	fMe := f.Bool("me", false, "user messages fitered")
	if err = f.Parse(args); err != nil {
		return errors.New("invalid flag(s) provided")
	}

	limit, err := strconv.Atoi(f.Args()[0])
	if err != nil || limit < 1 || limit > 100 {
		if message, fail := ctx.Session.ChannelMessageSendReply(ctx.Message.ChannelID, "Invalid amount (1-100), please try again!", ctx.Message.Reference()); fail == nil {
			time.AfterFunc(3*time.Second, func() {
				ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, message.ID)
			})
		}
		return
	}

	if *fYou || *fMe {
		messages, fail := ctx.Session.ChannelMessages(ctx.Message.ChannelID, 100, ctx.Message.ID, "", "")
		if fail != nil {
			return
		}
		filtered := make([]string, 0)
		for _, message := range messages {
			if *fYou && len(filtered) < limit && message.Author.ID == ctx.Session.State.User.ID {
				filtered = append(filtered, message.ID)
			}
			if *fMe && len(filtered) < limit && message.Author.ID == ctx.Message.Author.ID {
				filtered = append(filtered, message.ID)
			}
		}
		err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, filtered)
	} else {
		messages, fail := ctx.Session.ChannelMessages(ctx.Message.ChannelID, limit, ctx.Message.ID, "", "")
		if fail != nil {
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
	}

	notiContent := "I've cleared that message."
	if limit > 1 && !*fYou && !*fMe {
		notiContent = fmt.Sprintf("I've cleared %d messages.", limit)
	}
	if limit > 1 && (*fYou || *fMe) {
		notiContent = fmt.Sprintf("I've cleared %d messages with the selected filter.", limit)
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
