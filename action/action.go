package action

import (
	"fmt"
	"github.com/jbaikge/ingress-destroyed/damage"
	"github.com/jbaikge/ingress-destroyed/extract"
	"github.com/jbaikge/ingress-destroyed/mail"
	"github.com/jbaikge/ingress-destroyed/point"
	"time"
)

type Action struct {
	Id struct {
		MessageId string
		ActionId  int
	}
	Agent    string
	Enemy    string
	Damage   damage.Damage
	Time     time.Time
	Point    point.Point
	EndPoint point.Point // Only used when Damage.Type == damage.Link
}

func FromMessage(m *mail.Message) (actions []*Action, err error) {
	lines := extract.Lines(m.HTML)
	tpl := Action{
		Agent: extract.Name(lines[0]),
		Enemy: extract.Enemy(m.HTML),
		Time:  m.Date,
	}

	actions = make([]*Action, 0, len(lines[1:]))
	for _, l := range lines {
		a := &Action{}
		*a = tpl
		if err := FromLine(l, a); err != nil {
			continue
		}
		actions = append(actions, a)
	}
	return
}

func FromLine(l []byte, a *Action) (err error) {
	extract.Damage(l, &a.Damage)
	if a.Damage.Type == damage.Unknown {
		return fmt.Errorf("Valid damage not found in %s", string(l))
	}
	urls := extract.Links(l)
	if len(urls) == 0 {
		return fmt.Errorf("No URLs found")
	}
	points, err := point.FromURLs(urls)
	if err != nil {
		return
	}
	switch len(points) {
	case 1:
		a.Point, a.EndPoint = *points[0], point.Point{0.0, 0.0}
	case 2:
		a.Point, a.EndPoint = *points[0], *points[1]
	default:
		err = fmt.Errorf("Wasn't expecting %d points", len(points))
	}
	return
}
