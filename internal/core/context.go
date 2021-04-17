package core

import (
	"github.com/aslanluong/fairy/internal/utils"
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Channel *discordgo.Channel
	Args    ArgumentList
}

func (ctx *Context) SendMessage(content string) {
	utils.SendMessage(ctx.Session, ctx.Channel.ID, content)
}

func (ctx *Context) SendMessageReply(content string) {
	utils.SendMessageReply(ctx.Session, ctx.Message.Reference(), content)
}
