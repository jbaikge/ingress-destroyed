package extract

import (
	"bytes"
	"github.com/jbaikge/ingress-destroyed/damage"
	"net/url"
	"regexp"
	"strconv"
)

var (
	breakSplit = []byte(`<br/><br/>`)
	reURL      = regexp.MustCompile(`href="(https?://(www.)?ingress.com/intel\?[^"]+)"`)
	reBR       = regexp.MustCompile(`\s*<br/>(\s*<br/>)+\s*`)
)

func Damage(b []byte, d *damage.Damage) {
	d.Type = damage.Unknown

	fields := bytes.Fields(b)
	if len(fields) < 2 {
		return
	}
	switch 0 {
	case bytes.Compare(fields[1], []byte(`Resonator(s)`)):
		d.Type = damage.Resonator
		d.Count, _ = strconv.Atoi(string(fields[0]))
	case bytes.Compare(fields[1], []byte(`Mod(s)`)):
		d.Type = damage.Mod
		d.Count, _ = strconv.Atoi(string(fields[0]))
	case bytes.Compare(fields[1], []byte(`Link`)):
		d.Type = damage.Link
		d.Count = 1
	}
}

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

func Lines(b []byte) (l [][]byte) {
	b = reBR.ReplaceAll(b, breakSplit)
	return bytes.Split(b, breakSplit)
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
