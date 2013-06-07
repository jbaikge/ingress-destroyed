package location

import (
	"time"
)

type Action struct {
	Agent     []byte
	Destroyer []byte
	Time      time.Time
	Count     int
	Location  Point
}
