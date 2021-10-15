package notify

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"newWorldServerStatusBot/discord/handlers/command/notify"
	"newWorldServerStatusBot/discord/handlers/dm"
	"newWorldServerStatusBot/status"
	"time"
)

func NotifyHandler(s *discordgo.Session, subQueue <-chan notify.Subscriber) {
	prev := make([]status.Server, 0)
	current := make([]status.Server, 0)
	var subs []notify.Subscriber
	pollTicker := time.NewTicker(15 * time.Minute) // TODO: add to config

	for {
		select {
		case sub := <-subQueue: // a user has just typed '!nw notify <region> <server>' and is being added to the sub list
			// TODO: add to db here
			subs = append(subs, sub)
			err := dm.DmUser(s, sub.User, "I will now notify you whenever "+sub.ServerIn.Name+" is online again")
			if err != nil {
				log.Printf("Failed to Dm user: %s\n") // TODO: tell user that they might have dms disabled
			}
		case <-pollTicker.C: // check for servers that just went online and tell users who care
			// update current and previous server statuses
			var subbedServers []status.ServerInput
			for _, s := range subs {
				subbedServers = append(subbedServers, s.ServerIn)
			}
			var err error
			prev = current
			current, err = status.GetServerStatusMany(subbedServers)
			if err != nil {
				log.Fatalf("failed to update server statuses")
			}
			fmt.Printf("polled new server updates\nprev: %v\ncurrent: %v\n\n", prev, current)
			// dm users about their server which just went online
			toBeNotifiedUsers := getUsersWithNewlyOnlineServers(subs, current, prev)
			for u, sn := range toBeNotifiedUsers {
				msg := formatDmMessage(sn)
				if err = dm.DmUser(s, u, msg); err != nil {
					log.Printf("Failed to Dm user: %s\n")
				}
			}
		}
	}
}

func formatDmMessage(servernames []string) string {
	var msg string
	if len(servernames) == 1 {
		msg = fmt.Sprintf("%s has just gone online again!", servernames[0])
	} else if len(servernames) == 2 {
		msg = fmt.Sprintf("%s and %s have just gone online again!", servernames[0], servernames[1])
	} else {
		for i, s := range servernames {
			msg = fmt.Sprintf("%s%s, ", msg, s)
			if i >= len(servernames)-2 {
				break
			}
		}
		msg = fmt.Sprintf("%sand %s have just gone online again!", msg, servernames[len(servernames)-1])
	}
	return msg
}

func getUsersWithNewlyOnlineServers(subs []notify.Subscriber, current []status.Server, prev []status.Server) map[discordgo.User][]string {
	freshServers, changed := getNewlyOnlineServers(prev, current)
	if !changed {
		return nil
	}

	toBeNotifiedUsers := make(map[discordgo.User][]string)
	for _, freshServer := range freshServers {
		for _, s := range subs {
			if s.ServerIn.Name == freshServer {
				toBeNotifiedUsers[s.User] = append(toBeNotifiedUsers[s.User], freshServer)
			}
		}
	}
	return toBeNotifiedUsers
}

func getNewlyOnlineServers(prev []status.Server, current []status.Server) ([]string, bool) {
	var newlyOnline []string
	hasNewServers := false
	for _, currentServer := range current {
		for _, prevServer := range prev {
			if prevServer.Name == currentServer.Name {
				if prevServer.Status != "Online" && currentServer.Status == "Online" {
					newlyOnline = append(newlyOnline, currentServer.Name)
					hasNewServers = true
				}
				break
			}
		}
	}
	return newlyOnline, hasNewServers
}
