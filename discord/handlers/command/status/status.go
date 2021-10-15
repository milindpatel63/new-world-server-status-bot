package status

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"newWorldServerStatusBot/status"
	"strings"
)

type StatusCommand struct {
}

func (c StatusCommand) Check(parameters []string) bool {
	return parameters[0] == "USE" || parameters[0] == "USW" || parameters[0] == "SA" || parameters[0] == "EU" || parameters[0] == "AP" ||
		parameters[0] == "use" || parameters[0] == "usw" || parameters[0] == "sa" || parameters[0] == "eu" || parameters[0] == "ap"
}

func (c StatusCommand) Handle(author discordgo.User, parameters []string) (string, error) {
	region, err := status.ParseRegion(parameters[0])
	if err != nil {
		return "", err
	}

	if len(parameters) == 1 {
		servers, err := status.GetStatusForRegion(region)
		if err != nil {
			return "", err
		}
		return reduceServerStatus(region.String(), servers), nil
	} else if len(parameters) >= 2 {
		name := strings.Join(parameters[1:], " ")
		serverStatus, err := status.GetServerStatus(name, region)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("'%s' is %s", name, serverStatus), nil
	}
	return "", errors.New("invalid number of arguments")
}

func reduceServerStatus(regionName string, servers map[string]string) string {
	distString := fmt.Sprintf("Server Status distribution for %s\n", regionName)
	dist := make(map[string]float64)
	for _, serverStatus := range servers {
		dist[serverStatus] = dist[serverStatus] + 1
	}
	for s, d := range dist {
		l := float64(len(servers))
		distString = fmt.Sprintf("%s\t'%s': %.0f%%\n", distString, s, d/l*100)
	}
	return distString
}
