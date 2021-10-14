package status

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

const (
	UsEast = iota
	SaEast
	EuCentral
	ApSoutheast
	UsWest
)

func GetStatuses() map[string]string {
	resp, err := http.Get("https://www.newworld.com/en-gb/support/server-status") //TODO: add url to .env
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		log.Fatalf("got non 2xx response: %d %s\n", resp.StatusCode, resp.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return parseServerListForRegion(doc, EuCentral)
}
