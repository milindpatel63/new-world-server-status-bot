package dm

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func DmUser(s *discordgo.Session, u discordgo.User, msg string) error {
	channel, err := s.UserChannelCreate(u.ID)
	if err != nil {
		return fmt.Errorf("error creating channel: %s", err)
	}

	_, err = s.ChannelMessageSend(channel.ID, msg)
	if err != nil {
		return fmt.Errorf("error sending DM message (possibly disabled dms: %s", err)
	}
	return nil
}
