package main

import (
	"net/url"
	"testing"
)

func TestE6URL(t *testing.T) {
	u, _ := url.Parse("http://www.ingress.com/intel?latE6=38752729&lngE6=-77469455&z=19")
	l, err := urlLocation(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%0.6f %0.6f", l.Lat, l.Lon)
}

func TestLLURL(t *testing.T) {
	u, _ := url.Parse("http://www.ingress.com/intel?ll=38.749662,-77.473900&pll=38.749662,-77.473900&z=19")
	l, err := urlLocation(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%0.6f %0.6f", l.Lat, l.Lon)
}

func TestURLs(t *testing.T) {
	u := make([]*url.URL, 2)
	u[0], _ = url.Parse("http://www.ingress.com/intel?latE6=38752729&lngE6=-77469455&z=19")
	u[1], _ = url.Parse("http://www.ingress.com/intel?ll=38.749662,-77.473900&pll=38.749662,-77.473900&z=19")
	locs, err := urlLocations(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(locs[0])
	t.Log(locs[1])
}
