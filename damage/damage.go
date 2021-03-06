package damage

type Damage struct {
	Type  DamageType
	Count int
}

type DamageType string

const (
	Link      = DamageType("Link")
	Mod       = DamageType("Mod")
	Resonator = DamageType("Resonator")
	Unknown   = DamageType("Unknown")
)

func (t DamageType) String() string {
	return string(t)
}
