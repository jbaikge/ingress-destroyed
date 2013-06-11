package damage

import (
	"testing"
)

func TestString(t *testing.T) {
	cmp := []struct {
		D DamageType
		S string
	}{
		{Link, "Link"},
		{Mod, "Mod"},
		{Resonator, "Resonator"},
	}
	for _, ds := range cmp {
		if s := ds.D.String(); s != ds.S {
			t.Errorf("%s != %s", s, ds.S)
		}
	}
}
