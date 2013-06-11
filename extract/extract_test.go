package extract

import (
	"testing"
)

var extractHtml = []byte(`jbaikge,<br/><br/>2 Resonator(s) were destroyed by GuristaPrime at 19:21 hrs. - <a href="http://www.ingress.com/intel?ll=38.751726,-77.473140&pll=38.751726,-77.473140&z=19">View location</a><br/> <br/><br/><br/><br/>Your Link has been destroyed by GuristaPrime at 19:21 hrs. - <a href="http://www.ingress.com/intel?ll=38.751586,-77.473347&pll=38.751586,-77.473347&z=19">View start location</a> - <a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19">View end location</a><br/><br/><a href="http://www.ingress.com/intel?ll=38.751586,-77.473347&pll=38.751586,-77.473347&z=19"><img src="http://lh4.ggpht.com/caAHm4YCEDWcHW_LKsTyBW5wMSfZ4V2cMRvaJM-1FjsaFK0E0ISytG4-4N_WKAmBVQCD-iE02bGOP4iY74ji"/></a>&nbsp;<a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19"><img src="http://lh5.ggpht.com/AiADPDS4o27nZ7nT3Iqpc0Q2sZomkFoFclEbQoWQRJ-qXbRSuie47irq8nvYZTlIlAoRgNACgSf6irIm7uOz"/></a><br/><br/>Your Link has been destroyed by GuristaPrime at 19:21 hrs. - <a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19">View start location</a> - <a href="http://www.ingress.com/intel?ll=38.751782,-77.473600&pll=38.751782,-77.473600&z=19">View end location</a><br/><br/><a href="http://www.ingress.com/intel?ll=38.751757,-77.472814&pll=38.751757,-77.472814&z=19"><img src="http://lh5.ggpht.com/AiADPDS4o27nZ7nT3Iqpc0Q2sZomkFoFclEbQoWQRJ-qXbRSuie47irq8nvYZTlIlAoRgNACgSf6irIm7uOz"/></a>&nbsp;<br/><br/>------------------------------------------<br/><a href="http://www.ingress.com/intel">Dashboard</a>&nbsp;<a href="http://support.google.com/ingress">Contact</a><br/>`)

func TestEnemy(t *testing.T) {
	if e := Enemy(extractHtml); e != "GuristaPrime" {
		t.Errorf("Expected GuristaPrime, got: %s", e)
	}
}

func TestLinks(t *testing.T) {
	urls := Links(extractHtml)
	for i, u := range urls {
		t.Logf("[%d] %s", i, u)
	}
}

func TestName(t *testing.T) {
	if n := Name(extractHtml); n != "jbaikge" {
		t.Errorf("Expected jbaikge; got: %s", n)
	}
}
