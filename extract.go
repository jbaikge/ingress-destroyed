package main

import (
	"net/url"
	"regexp"
)

var reURL = regexp.MustCompile(`href="(https?://(www.)?ingress.com/intel\?[^"]+)"`)

func extractLinks(b []byte) (urls []*url.URL) {
	for _, rawurl := range reURL.FindAllSubmatch(b, -1) {
		u, err := url.Parse(string(rawurl[1]))
		if err != nil {
			continue
		}
		urls = append(urls, u)
	}
	return
}
