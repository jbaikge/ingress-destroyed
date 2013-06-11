package point

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Point struct {
	Lat float64
	Lon float64
}

func FromURLs(urls []*url.URL) (points []*Point, err error) {
	points = make([]*Point, len(urls))
	var l *Point
	for i := range urls {
		if l, err = urlPoint(urls[i]); err != nil {
			break
		}
		points[i] = l
	}
	return
}

func e6Point(lat, lon string) (l *Point, err error) {
	var iLat, iLon int64
	if iLat, err = strconv.ParseInt(lat, 10, 64); err != nil {
		return
	}
	if iLon, err = strconv.ParseInt(lon, 10, 64); err != nil {
		return
	}
	l = &Point{
		Lat: float64(iLat) / float64(1e6),
		Lon: float64(iLon) / float64(1e6),
	}
	return
}

func llPoint(s string) (l *Point, err error) {
	l = &Point{}
	latlon := strings.Split(s, ",")
	if l.Lat, err = strconv.ParseFloat(latlon[0], 10); err != nil {
		return
	}
	if l.Lon, err = strconv.ParseFloat(latlon[1], 10); err != nil {
		return
	}
	return
}

func urlPoint(u *url.URL) (l *Point, err error) {
	if ll, ok := u.Query()["ll"]; ok {
		return llPoint(ll[0])
	} else if latE6, ok := u.Query()["latE6"]; ok {
		return e6Point(latE6[0], u.Query()["lngE6"][0])
	}
	return nil, errors.New("No Latitude or Longitude found")
}
