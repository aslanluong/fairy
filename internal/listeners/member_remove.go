package listeners

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MemberRemoveListener struct {
}

func NewMemberRemoveListener() *MemberRemoveListener {
	return &MemberRemoveListener{}
}

func (l *MemberRemoveListener) Handler(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		fmt.Println("error getting guild object,", err)
		return
	}

	fmt.Printf("Member %s left the guild %s\n", e.Member.User.String(), guild.Name)
}
