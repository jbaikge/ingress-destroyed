package main

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Location struct {
	Lat float64
	Lon float64
}

func e6Location(lat, lon string) (l *Location, err error) {
	var iLat, iLon int64
	if iLat, err = strconv.ParseInt(lat, 10, 64); err != nil {
		return
	}
	if iLon, err = strconv.ParseInt(lon, 10, 64); err != nil {
		return
	}
	l = &Location{
		Lat: float64(iLat) / float64(1e6),
		Lon: float64(iLon) / float64(1e6),
	}
	return
}

func llLocation(s string) (l *Location, err error) {
	l = &Location{}
	latlon := strings.Split(s, ",")
	if l.Lat, err = strconv.ParseFloat(latlon[0], 10); err != nil {
		return
	}
	if l.Lon, err = strconv.ParseFloat(latlon[1], 10); err != nil {
		return
	}
	return
}

func urlLocation(u *url.URL) (l *Location, err error) {
	if ll, ok := u.Query()["ll"]; ok {
		return llLocation(ll[0])
	} else if latE6, ok := u.Query()["latE6"]; ok {
		return e6Location(latE6[0], u.Query()["lngE6"][0])
	}
	return nil, errors.New("No Latitude or Longitude found")
}

func urlLocations(urls []*url.URL) (locs []*Location, err error) {
	locs = make([]*Location, len(urls))
	var l *Location
	for i := range urls {
		if l, err = urlLocation(urls[i]); err != nil {
			break
		}
		locs[i] = l
	}
	return
}
