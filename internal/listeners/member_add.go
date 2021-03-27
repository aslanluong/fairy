package listeners

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MemberAddListener struct {
}

func NewMemberAddListener() *MemberAddListener {
	return &MemberAddListener{}
}

func (l *MemberAddListener) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		fmt.Println("error getting guild object,", err)
		return
	}

	fmt.Printf("Member %s joined the guild %s\n", e.Member.User.String(), guild.Name)
}
