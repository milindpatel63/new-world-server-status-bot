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
	var bugCommand commands.BugCommand
	var helpCommand commands.HelpCommand

	if statusCommand.Check(parameters) {
		resp, err := statusCommand.Handle(parameters)
		if err != nil {
			return "", err
		}
		return resp, nil
	} else if bugCommand.Check(parameters) {
		resp, _ := bugCommand.Handle(parameters)
		return resp, nil
	} else if helpCommand.Check(parameters) {
		resp, _ := helpCommand.Handle(parameters)
		return resp, nil
	}
	return "", errors.New("command not found")
}
