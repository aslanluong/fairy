package utils

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	*discordgo.Message

	session *discordgo.Session
}

func SendMessage(s *discordgo.Session, channelID string, content string) (*Message, error) {
	msg, err := s.ChannelMessageSend(channelID, content)
	return &Message{msg, s}, err
}

func SendMessageReply(s *discordgo.Session, reference *discordgo.MessageReference, content string) (*Message, error) {
	msg, err := s.ChannelMessageSendReply(reference.ChannelID, "Invalid command name, please try again!", reference)
	return &Message{msg, s}, err
}

func (msg *Message) DeleteAfter(d time.Duration) (err error) {
	time.AfterFunc(d, func() {
		err = msg.session.ChannelMessageDelete(msg.ChannelID, msg.ID)
	})
	return
}
