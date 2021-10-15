package status

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

type ServerInput struct {
	Name   string
	Region Region
}

type Server struct {
	Name   string
	Region Region
	Status string
}

func GetStatusForRegion(region Region) (map[string]string, error) {
	resp, err := http.Get("http://localhost:3000/mock") //TODO: add url to .env
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("got non 2xx response: %d %s\n", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return parseServerListForRegion(doc, region), nil
}

func GetServerStatus(name string, region Region) (string, error) {
	resp, err := http.Get("http://localhost:3000/mock") //TODO: add url to .env
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("got non 2xx response: %d %s\n", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	status := parseServerStatus(doc, name, region)
	if status == "" {
		return "", fmt.Errorf("server status for '%s' not found", name)
	}
	return status, nil
}

func GetServerStatusMany(serverInputs []ServerInput) ([]Server, error) {
	var ServerStatuses []Server
	// decide which regions we need to parse
	regionMap := make(map[Region]bool)
	for _, serverInput := range serverInputs {
		regionMap[serverInput.Region] = true
	}

	// https://www.newworld.com/en-gb/support/server-status
	resp, err := http.Get("http://localhost:3000/mock") //TODO: add url to .env
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("got non 2xx response: %d %s\n", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	for region := range regionMap {
		statuses := parseServerListForRegion(doc, region)
		for _, serverInput := range serverInputs {
			ServerStatuses = append(ServerStatuses, Server{
				Name:   serverInput.Name,
				Region: region,
				Status: statuses[serverInput.Name],
			})
		}
	}
	return ServerStatuses, nil
}
