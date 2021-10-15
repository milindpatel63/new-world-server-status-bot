package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"newWorldServerStatusBot/config"
	notify2 "newWorldServerStatusBot/discord/handlers/command/notify"
	"newWorldServerStatusBot/discord/handlers/messageCreate"
	"newWorldServerStatusBot/notify"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = godotenv.Load()

	c := config.New()

	dg, err := discordgo.New("Bot " + c.DiscordToken)
	if err != nil {
		log.Fatalln("error creating Discord session: ", err)
		return
	}

	notifySubQueue := make(chan notify2.Subscriber)
	go notify.NotifyHandler(dg, notifySubQueue)
	dg.AddHandler(messageCreate.MessageCreateHandler([]interface{}{
		notifySubQueue,
	}))

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
