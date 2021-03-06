package commands

type CmdPing struct {
}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdPing) Description() string {
	return "Pong!"
}

func (c *CmdPing) AdminRequired() bool {
	return false
}

func (c *CmdPing) Exec(ctx *Context) (err error) {
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!")
	return
}
