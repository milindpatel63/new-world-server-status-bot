package command

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"newWorldServerStatusBot/discord/handlers/command/bug"
	"newWorldServerStatusBot/discord/handlers/command/help"
	"newWorldServerStatusBot/discord/handlers/command/notify"
	"newWorldServerStatusBot/discord/handlers/command/status"
)

type CommandRouter struct {
	commands []Command
}

func NewCommandRouter(options []interface{}) *CommandRouter {
	var notifyUsersQueue chan notify.Subscriber
	for _, option := range options {
		switch opt := option.(type) {
		case chan notify.Subscriber:
			notifyUsersQueue = opt
		}
	}
	return &CommandRouter{commands: []Command{
		status.StatusCommand{},
		notify.NewNotifyCommand(notifyUsersQueue),
		bug.BugCommand{},
		help.HelpCommand{},
	}}
}

type Command interface {
	Check(parameters []string) bool
	Handle(author discordgo.User, parameters []string) (string, error)
}

func (r *CommandRouter) ExecuteCommand(author discordgo.User, parameters []string) (string, error) {
	for _, c := range r.commands {
		if c.Check(parameters) {
			response, err := c.Handle(author, parameters)
			if err != nil {
				return "", err
			}
			return response, nil
		}
	}
	return "", errors.New("command not found")
}
