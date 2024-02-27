package spells

import "hello.go.game/numbers"

type Nightmare struct {
	SpellBase
}

func NewNightmare() Nightmare {
	return Nightmare{SpellBase: SpellBase{Name: "Nightmare", ManaCost: 13, BaseDamage: 8, SpellModifier: 1.8}}
}

func (n Nightmare) GetName() string {
	return n.Name
}

func (n Nightmare) GetBaseDamage() int {
	return n.BaseDamage
}

func (n Nightmare) GetSpellModifier() float32 {
	return n.SpellModifier
}

func (n Nightmare) GetManaCost() int {
	return n.ManaCost
}

func (n Nightmare) Equal(other Spell) bool {
	if otherNightmare, ok := other.(Nightmare); ok {
		return n.Name == otherNightmare.Name &&
			n.BaseDamage == otherNightmare.BaseDamage &&
			n.ManaCost == otherNightmare.ManaCost &&
			n.SpellModifier == otherNightmare.SpellModifier
	}
	return false
}

func (n Nightmare) CalculateSpellDamage(caster SpellCaster) int {
	return numbers.Round(float32(caster.GetMagicPower())*n.SpellModifier) + n.BaseDamage
}
