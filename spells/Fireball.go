package spells

import (
	"hello.go.game/numbers"
)

type Fireball struct {
	SpellBase
}

func NewFireball() Fireball {
	return Fireball{SpellBase: SpellBase{Name: "Fireball", ManaCost: 10, BaseDamage: 6, SpellModifier: 1.8}}
}

func (f Fireball) GetName() string {
	return f.Name
}

func (f Fireball) GetBaseDamage() int {
	return f.BaseDamage
}

func (f Fireball) GetSpellModifier() float32 {
	return f.SpellModifier
}

func (f Fireball) GetManaCost() int {
	return f.ManaCost
}

func (f Fireball) Equal(other Spell) bool {
	if otherFireball, ok := other.(Fireball); ok {
		return f.Name == otherFireball.Name &&
			f.BaseDamage == otherFireball.BaseDamage &&
			f.ManaCost == otherFireball.ManaCost &&
			f.SpellModifier == otherFireball.SpellModifier
	}
	return false
}

func (f Fireball) CalculateSpellDamage(caster SpellCaster) int {
	return numbers.Round(float32(caster.GetMagicPower())*f.SpellModifier) + numbers.RandomWithLowerOne(f.BaseDamage)
}
