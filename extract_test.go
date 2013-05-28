package main

import (
	"testing"
)

func TestLinks(t *testing.T) {
	html := []byte(`jbaikge,<br/><br/>1 Resonator(s) were destroyed by DAnonymous at 17:52 hrs. - <a href="http://www.ingress.com/intel?ll=38.749662,-77.473900&pll=38.749662,-77.473900&z=19">View location</a><br/> <br/><a href="http://www.ingress.com/intel?ll=38.749662,-77.473900&pll=38.749662,-77.473900&z=19"><img src="http://lh3.ggpht.com/Fob5NoNdABw8PINL4Q0ZEx1zHRtU3yNkh1yleUXKdnjTXdbJqk1Tom4exEM44XnAsa8sJYhd7gukCuigfndY"/></a><br/><br/><br/><br/>------------------------------------------<br/><a href="http://www.ingress.com/intel">Dashboard</a>&nbsp;<a href="http://support.google.com/ingress">Contact</a><br/>`)
	urls := extractLinks(html)
	for i, u := range urls {
		t.Logf("[%02d] %s", i, u)
	}
}
