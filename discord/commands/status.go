package commands

import (
	"errors"
	"fmt"
	"newWorldServerStatusBot/status"
)

type StatusCommand struct {
}

func (c StatusCommand) Check(parameters []string) bool {
	return parameters[0] == "USE" || parameters[0] == "USW" || parameters[0] == "SA" || parameters[0] == "EU" || parameters[0] == "AP" ||
		parameters[0] == "use" || parameters[0] == "usw" || parameters[0] == "sa" || parameters[0] == "eu" || parameters[0] == "ap"
}

func (c StatusCommand) Handle(parameters []string) (string, error) {
	region, err := formatRegionToRegionIndex(parameters[0])
	if err != nil {
		return "", err
	}

	servers, err := status.GetStatuses(region)
	if err != nil {
		return "", err
	}

	if len(parameters) == 1 {
		return reduceServerStatus(region.String(), servers), nil
	} else if len(parameters) == 2 {
		if s, ok := servers[parameters[1]]; ok {
			return fmt.Sprintf("'%s' is %s", parameters[1], s), nil
		} else {
			return "", fmt.Errorf("server '%s' not found in %s", parameters[1], region.String())
		}
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

func formatRegionToRegionIndex(region string) (status.Region, error) {
	if region == "USE" || region == "use" {
		return status.UsEast, nil
	} else if region == "USW" || region == "usw" {
		return status.UsWest, nil
	} else if region == "SA" || region == "sa" {
		return status.SaEast, nil
	} else if region == "EU" || region == "eu" {
		return status.EuCentral, nil
	} else if region == "AP" || region == "ap" {
		return status.ApSoutheast, nil
	} else {
		return -1, errors.New("could not parse region parameter: " + region)
	}
}
