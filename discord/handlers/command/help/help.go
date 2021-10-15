package help

import "github.com/bwmarrin/discordgo"

type HelpCommand struct {
}

func (h HelpCommand) Check(parameters []string) bool {
	return parameters[0] == "help" || parameters[0] == "Help" || parameters[0] == "HELP" || parameters[0] == "?"
}

func (h HelpCommand) Handle(author discordgo.User, parameters []string) (string, error) {
	return "**!nw <region>** - shows the distribution of statuses of all servers in a region\n" +
		"**!nw <region> <servername>** - gets the status of a server\n" +
		"**!nw bug** - shows instructions on how to report a bug or issue with the bot\n" +
		"**!nw help** - shows this help message", nil
}
