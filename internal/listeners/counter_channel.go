package listeners

import (
	"strconv"
	"time"

	"github.com/aslanluong/fairy/internal/utils"
	"github.com/bwmarrin/discordgo"
)

type CounterChannelListener struct {
	enableChannelIDs  map[string]bool
	invalidMessageIDs map[string]bool
}

var listener = &CounterChannelListener{
	enableChannelIDs:  make(map[string]bool),
	invalidMessageIDs: make(map[string]bool),
}

func CounterChannel() *CounterChannelListener {
	return &CounterChannelListener{}
}

func (l *CounterChannelListener) EnableListener(channelID string) {
	listener.enableChannelIDs[channelID] = true
}

func (l *CounterChannelListener) DisableListener(channelID string) {
	delete(listener.enableChannelIDs, channelID)
}

func (l *CounterChannelListener) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot || !listener.enableChannelIDs[e.ChannelID] {
		return
	}

	validCount, err := strconv.Atoi(e.Content)
	if err != nil {
		listener.invalidMessageIDs[e.ID] = true
		msg, err := utils.SendMessageReply(s, e.Message.Reference(), "Invalid count number!")
		if err == nil {
			msg.WaitAndDelete(3 * time.Second)
			utils.DeleteMessage(s, e.ChannelID, e.ID)
		}
		delete(listener.invalidMessageIDs, e.ID)
		return
	}

	prevMessages, err := s.ChannelMessages(e.ChannelID, 100, e.ID, "", "")
	if err != nil {
		utils.DeleteMessage(s, e.ChannelID, e.ID)
		return
	}

	if len(prevMessages) == 0 {
		if validCount != 0 {
			listener.invalidMessageIDs[e.ID] = true
			msg, err := utils.SendMessageReply(s, e.Message.Reference(), "Invalid count number!")
			if err == nil {
				msg.WaitAndDelete(3 * time.Second)
				utils.DeleteMessage(s, e.ChannelID, e.ID)
			}
			delete(listener.invalidMessageIDs, e.ID)
		}
		return
	}

	for _, prevMessage := range prevMessages {
		if !prevMessage.Author.Bot && !listener.invalidMessageIDs[prevMessage.ID] {
			if prevCountNumber, err := strconv.Atoi(prevMessage.Content); (err != nil && validCount != 0) || (err == nil && prevCountNumber+1 != validCount) {
				listener.invalidMessageIDs[e.ID] = true
				msg, err := utils.SendMessageReply(s, e.Message.Reference(), "Invalid count number!")
				if err == nil {
					msg.WaitAndDelete(3 * time.Second)
					utils.DeleteMessage(s, e.ChannelID, e.ID)
				}
				delete(listener.invalidMessageIDs, e.ID)
			}
			break
		}
	}
}
