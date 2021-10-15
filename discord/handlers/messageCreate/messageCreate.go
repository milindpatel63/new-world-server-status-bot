package messageCreate

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"newWorldServerStatusBot/discord/handlers/command"
	"strings"
)

func MessageCreateHandler(commandOptions []interface{}) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	router := command.NewCommandRouter(commandOptions)

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.HasPrefix(m.Content, "!nw") {
			commandPrefix := strings.TrimPrefix(m.Content, "!nw ")
			params := strings.Split(commandPrefix, " ")
			response, err := router.ExecuteCommand(*m.Message.Author, params)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "oops! something went wrong: "+err.Error())
				log.Fatalf("command '%v' failed: %s", params, err)
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, response)
			}
		}
	}
}
