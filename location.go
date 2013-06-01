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
	l = &Location{}
	if l.Lat, err = strconv.ParseFloat(lat, 10); err != nil {
		return
	}
	if l.Lon, err = strconv.ParseFloat(lon, 10); err != nil {
		return
	}
	l.Lat = l.Lat / float64(10e6)
	l.Lon = l.Lon / float64(10e6)
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
