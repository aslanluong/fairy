package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/aslanluong/fairy/internal/utils"
)

type CmdChoose struct {
}

func (c *CmdChoose) Invokes() []string {
	return []string{"choose"}
}

func (c *CmdChoose) Description() string {
	return "(Maybe) Help you to get the luckiest choice! (seperate with '|')"
}

func (c *CmdChoose) AdminRequired() bool {
	return false
}

func (c *CmdChoose) Exec(ctx *Context) (err error) {
	args := ctx.Args

	if len(args) == 0 || len(args) == 1 && args[0] == "|" {
		return
	}

	if len(args) == 1 {
		_, err = utils.SendMessage(ctx.Session, ctx.Message.ChannelID, fmt.Sprintf("**%s** is the best one for sure", args[0]))
		return
	}

	validArgs := utils.FilterString(strings.Split(strings.Join(args, " "), "|"), func(val string) bool { return val != "" })

	chose := validArgs[rand.Intn(len(validArgs))]
	_, err = utils.SendMessage(ctx.Session, ctx.Message.ChannelID, fmt.Sprintf("**%s** is probably the best choice", chose))
	return
}
