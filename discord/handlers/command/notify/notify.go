package notify

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"newWorldServerStatusBot/status"
	"strings"
)

type Subscriber struct {
	User     discordgo.User
	ServerIn status.ServerInput
}

type NotifyCommand struct {
	UserInputQueue chan Subscriber
}

func NewNotifyCommand(queue chan Subscriber) *NotifyCommand {
	return &NotifyCommand{
		UserInputQueue: queue,
	}
}

func (n *NotifyCommand) Check(parameters []string) bool {
	return parameters[0] == "notify" || parameters[0] == "Notify" || parameters[0] == "NOTIFY"
}

func (n *NotifyCommand) Handle(author discordgo.User, parameters []string) (string, error) {
	// '!nw notify EU Delphnius'
	if len(parameters) < 3 {
		return "", errors.New("invalid argument count")
	}

	region, err := status.ParseRegion(parameters[1])
	if err != nil {
		return "", err
	}

	name := strings.Join(parameters[2:], " ")

	sub := Subscriber{
		User:     author,
		ServerIn: status.ServerInput{Name: name, Region: region},
	}

	n.UserInputQueue <- sub

	return "", nil
}
