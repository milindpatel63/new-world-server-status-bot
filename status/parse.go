package status

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

func parseServerListForRegion(doc *goquery.Document, region Region) map[string]string {
	servers := make(map[string]string)
	doc.Find(".ags-ServerStatus-content-responses-response").Each(func(i int, s *goquery.Selection) {
		val, exist := s.Attr("data-index")
		if exist {
			regionIndex, err := strconv.Atoi(val)
			if err == nil && Region(regionIndex) == region {
				s = s.Children().Filter(".ags-ServerStatus-content-responses-response-server")
				s.Each(func(i int, s *goquery.Selection) {
					serverDiv := s.Children()
					name := strings.TrimSpace(serverDiv.Last().Text())
					wrapper := serverDiv.First().Children()
					status, exists := wrapper.First().Attr("title")
					if !exists {
						log.Fatalf("something went wrong trying to find the status of server \"%s\"\n", name)
					} else {
						servers[name] = status
					}
				})
			}
		}
	})
	return servers
}
