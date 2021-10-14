package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"newWorldServerStatusBot/config"
	"newWorldServerStatusBot/status"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		return
	}

	c := config.New()

	dg, err := discordgo.New("Bot " + c.DiscordToken)
	if err != nil {
		log.Fatalln("error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalln("error opening connection: ", err)
		return
	}

	log.Println("Running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = dg.Close()
	if err != nil {
		log.Fatalln("error closing Discord sesssion: ", err)
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!nw") {
		server := strings.TrimLeft(m.Content, "!nw ")
		servers := status.GetStatuses()
		s.ChannelMessageSend(m.ChannelID, "status for "+server+": "+servers[server])
	}
}
