package extract

import (
	"bytes"
	"net/url"
	"regexp"
)

var reURL = regexp.MustCompile(`href="(https?://(www.)?ingress.com/intel\?[^"]+)"`)

func Enemy(b []byte) (s string) {
	fields := bytes.Fields(b)
	cmp := []byte(`destroyed`)
	for i, f := range fields {
		if bytes.Compare(f, cmp) == 0 {
			s = string(fields[i+2])
			break
		}
	}
	return
}

func Links(b []byte) (urls []*url.URL) {
	for _, rawurl := range reURL.FindAllSubmatch(b, -1) {
		u, err := url.Parse(string(rawurl[1]))
		if err != nil {
			continue
		}
		urls = append(urls, u)
	}
	return
}

func Name(b []byte) string {
	var end int
	stop := byte(',')
	for end = range b {
		if b[end] == stop {
			break
		}
	}
	return string(b[:end])
}
