package location

import (
	"fmt"
	"time"
)

type Action struct {
	Agent     []byte
	Destroyer []byte
	Time      time.Time
	Count     int
	Location  Point
}

func (a *Action) GMap() string {
	return fmt.Sprintf("\t\t{location: new google.maps.LatLng(%0.6f, %0.6f), weight: %d},", a.Location.Lat, a.Location.Lon, a.Count)
}
