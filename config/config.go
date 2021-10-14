package config

import "os"

type config struct {
	DiscordToken string
}

func New() *config {
	return &config{DiscordToken: os.Getenv("DISCORD_TOKEN")}
}
