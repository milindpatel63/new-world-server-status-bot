package status

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

type Region int

const (
	UsEast Region = iota
	SaEast
	EuCentral
	ApSoutheast
	UsWest
)

func (r Region) String() string {
	switch r {
	case UsEast:
		return "US EAST"
	case SaEast:
		return "SA EAST"
	case EuCentral:
		return "EU CENTRAL"
	case ApSoutheast:
		return "AP SOUTHEAST"
	case UsWest:
		return "US WEST"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

func GetStatuses(region Region) (map[string]string, error) {
	resp, err := http.Get("https://www.newworld.com/en-gb/support/server-status") //TODO: add url to .env
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
