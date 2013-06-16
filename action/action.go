package action

import (
	"errors"
	"fmt"
	"github.com/jbaikge/ingress-destroyed/damage"
	"github.com/jbaikge/ingress-destroyed/extract"
	"github.com/jbaikge/ingress-destroyed/message"
	"github.com/jbaikge/ingress-destroyed/point"
	"time"
)

type Action struct {
	Agent    string
	Enemy    string
	Damage   damage.Damage
	Time     time.Time
	Point    point.Point
	EndPoint point.Point // Only used when Damage.Type == damage.Link
}

type Actions []Action

func FromMessage(m *message.Message) (actions *Actions, err error) {
	// lines := extract.Lines(m.HTML)

	// tpl := Action{
	// 	Agent: extract.Name(lines[0]),
	// 	Enemy: extract.Enemy(m.HTML),
	// }
	// tpl.Time, _ = m.Message.Header.Date()

	// actions = make(&Actions, 0, len(lines[1:]))
	// for _, l := range lines {
	// 	a := tpl
	// 	if err = FromLine(l, &a); err != nil {
	// 		continue
	// 	}
	// 	*actions = append(*actions, a)
	// }
	return
}

func FromLine(l []byte, a *Action) (err error) {
	urls := extract.Links(l)
	if len(urls) == 0 {
		return errors.New("No URLs found")
	}
	points, err := point.FromURLs(urls)
	if err != nil {
		return
	}
	extract.Damage(l, &a.Damage)
	switch len(points) {
	case 1:
		a.Point = *points[0]
	case 2:
		a.Point, a.EndPoint = *points[0], *points[1]
	default:
		errors.New(fmt.Sprintf("Wasn't expecting %d points", len(points)))
	}
	return
}

func (a *Action) GMap() string {
	return fmt.Sprintf("\t\t{location: new google.maps.LatLng(%0.6f, %0.6f), weight: %d},", a.Point.Lat, a.Point.Lon, a.Damage.Count)
}
