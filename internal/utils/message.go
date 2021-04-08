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
	msg, err := s.ChannelMessageSendReply(reference.ChannelID, content, reference)
	return &Message{msg, s}, err
}

func DeleteMessage(s *discordgo.Session, channelID string, messageID string) (err error) {
	err = s.ChannelMessageDelete(channelID, messageID)
	return
}

func DeleteMessageAfter(s *discordgo.Session, channelID string, messageID string, d time.Duration, callback func()) (err error) {
	time.AfterFunc(d, func() {
		err = s.ChannelMessageDelete(channelID, messageID)
		if callback != nil {
			callback()
		}
	})
	return
}

func (msg *Message) WaitAndDelete(d time.Duration) (err error) {
	time.Sleep(d)
	err = msg.session.ChannelMessageDelete(msg.ChannelID, msg.ID)
	return
}
