package discord

import (
	"errors"
	"newWorldServerStatusBot/discord/commands"
)

type Command interface {
	Check(parameters []string) bool
	Handle(parameters []string) (string, error)
}

func ExecuteCommand(parameters []string) (string, error) {
	var statusCommand commands.StatusCommand

	if statusCommand.Check(parameters) {
		resp, err := statusCommand.Handle(parameters)
		if err != nil {
			return "", err
		}
		return resp, nil
	}
	return "", errors.New("command not found")
}
