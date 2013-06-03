package main

import (
	"bytes"
	"net/url"
	"regexp"
)

type Carnage struct {
	Type  CarnageType
	Count int
}

type CarnageType string

const (
	Link      = CarnageType("Link")
	Mod       = CarnageType("Mod")
	Resonator = CarnageType("Resonator")
)

var reURL = regexp.MustCompile(`href="(https?://(www.)?ingress.com/intel\?[^"]+)"`)

func ExtractCarnage(b []byte) (c []*Carnage) {
	return
}

func ExtractDestroyer(b []byte) (n []byte) {
	fields := bytes.Fields(b)
	cmp := []byte(`destroyed`)
	for i, f := range fields {
		if bytes.Compare(f, cmp) == 0 {
			n = fields[i+2]
			break
		}
	}
	return
}

func ExtractLinks(b []byte) (urls []*url.URL) {
	for _, rawurl := range reURL.FindAllSubmatch(b, -1) {
		u, err := url.Parse(string(rawurl[1]))
		if err != nil {
			continue
		}
		urls = append(urls, u)
	}
	return
}

func ExtractName(b []byte) (n []byte) {
	var end int
	stop := byte(',')
	for end = range b {
		if b[end] == stop {
			break
		}
	}
	return b[:end]
}
