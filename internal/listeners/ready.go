package listeners

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type ReadyListener struct {
}

func NewReadyListener() *ReadyListener {
	return &ReadyListener{}
}

func (h *ReadyListener) Handler(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Printf("Logged in as %s (%s) - Running on %d servers", e.User.String(), e.User.ID, len(e.Guilds))
}
