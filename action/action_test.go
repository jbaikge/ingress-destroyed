package action

import (
	"github.com/jbaikge/ingress-destroyed/damage"
	"github.com/jbaikge/ingress-destroyed/extract"
	"github.com/jbaikge/ingress-destroyed/point"
	"testing"
)

var extractHtml = []byte(`jbaikge,<br/><br/>2 Resonator(s) were destroyed by GuristaPrime at 19:21 hrs. - <a href="http://www.ingress.com/intel?ll=38.751726,-77.473140&pll=38.751726,-77.473140&z=19">View location</a><br/> <br/><br/><br/><br/>Your Link has been destroyed by GuristaPrime at 19:21 hrs. - <a href="http://www.ingress.com/intel?ll=38.751586,-77.473347&pll=38.751586,-77.473347&z=19">View start location</a> - <a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19">View end location</a><br/><br/><a href="http://www.ingress.com/intel?ll=38.751586,-77.473347&pll=38.751586,-77.473347&z=19"><img src="http://lh4.ggpht.com/caAHm4YCEDWcHW_LKsTyBW5wMSfZ4V2cMRvaJM-1FjsaFK0E0ISytG4-4N_WKAmBVQCD-iE02bGOP4iY74ji"/></a>&nbsp;<a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19"><img src="http://lh5.ggpht.com/AiADPDS4o27nZ7nT3Iqpc0Q2sZomkFoFclEbQoWQRJ-qXbRSuie47irq8nvYZTlIlAoRgNACgSf6irIm7uOz"/></a><br/><br/>Your Link has been destroyed by GuristaPrime at 19:21 hrs. - <a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19">View start location</a> - <a href="http://www.ingress.com/intel?ll=38.751782,-77.473600&pll=38.751782,-77.473600&z=19">View end location</a><br/><br/><a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19"><img src="http://lh5.ggpht.com/AiADPDS4o27nZ7nT3Iqpc0Q2sZomkFoFclEbQoWQRJ-qXbRSuie47irq8nvYZTlIlAoRgNACgSf6irIm7uOz"/></a>&nbsp;<br/><br/>------------------------------------------<br/><a href="http://www.ingress.com/intel">Dashboard</a>&nbsp;<a href="http://support.google.com/ingress">Contact</a><br/>`)

func TestLine(t *testing.T) {
	lines := extract.Lines(extractHtml)
	lineActions := map[int]Action{
		1: Action{
			Damage: &damage.Damage{
				Count: 2,
				Type:  damage.Resonator,
			},
			Point: &point.Point{38.751726, -77.473140},
		},
		3: Action{
			Damage: &damage.Damage{
				Count: 1,
				Type:  damage.Link,
			},
			Point:    &point.Point{38.751586, -77.473347},
			EndPoint: &point.Point{38.751757, -77.472814},
		},
		5: Action{
			Damage: &damage.Damage{
				Count: 1,
				Type:  damage.Link,
			},
			Point:    &point.Point{38.751757, -77.472814},
			EndPoint: &point.Point{38.751782, -77.473600},
		},
	}
	for i, A := range lineActions {
		line := lines[i]
		a := &Action{}
		if err := FromLine(line, a); err != nil {
			t.Fatal(err)
		}
		switch {
		case A.Damage.Type != a.Damage.Type:
			t.Errorf("[%d] Invalid damage type: %s != %s", i, A.Damage.Type, a.Damage.Type)
		case A.Damage.Count != a.Damage.Count:
			t.Errorf("[%d] Invalid damage count: %d != %d", i, A.Damage.Count, a.Damage.Count)
		case A.Point.Lat != a.Point.Lat:
			t.Errorf("[%d] Invalid point latitude: %d != %d", i, A.Point.Lat, a.Point.Lat)
		case A.Point.Lon != a.Point.Lon:
			t.Errorf("[%d] Invalid point longitude: %d != %d", i, A.Point.Lon, a.Point.Lon)
		case A.Damage.Type == damage.Link && A.EndPoint.Lat != a.EndPoint.Lat:
			t.Errorf("[%d] Invalid endpoint latitude: %d != %d", i, A.EndPoint.Lat, a.EndPoint.Lat)
		case A.Damage.Type == damage.Link && A.EndPoint.Lon != a.EndPoint.Lon:
			t.Errorf("[%d] Invalid endpoint longitude: %d != %d", i, A.EndPoint.Lon, a.EndPoint.Lon)
		}
	}
}
