package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Printf("Logged in as %s (%s) - Running on %d servers", e.User.String(), e.User.ID, len(e.Guilds))
}
