package spells

import "hello.go.game/numbers"

type DoubleAttack struct {
	SpellBase
}

func NewDoubleAttack() DoubleAttack {
	return DoubleAttack{SpellBase: SpellBase{Name: "DoubleAttack", ManaCost: 12, BaseDamage: 0, SpellModifier: 2.5}}
}

func (da DoubleAttack) GetName() string {
	return da.Name
}

func (da DoubleAttack) GetBaseDamage() int {
	return da.BaseDamage
}

func (da DoubleAttack) GetSpellModifier() float32 {
	return da.SpellModifier
}

func (da DoubleAttack) GetManaCost() int {
	return da.ManaCost
}

func (da DoubleAttack) Equal(other Spell) bool {
	if otherDoubleAttack, ok := other.(DoubleAttack); ok {
		return da.Name == otherDoubleAttack.Name &&
			da.BaseDamage == otherDoubleAttack.BaseDamage &&
			da.ManaCost == otherDoubleAttack.ManaCost &&
			da.SpellModifier == otherDoubleAttack.SpellModifier
	}
	return false
}

func (da DoubleAttack) CalculateSpellDamage(caster SpellCaster) int {
	return numbers.Round(float32(caster.GetMagicPower()) * da.SpellModifier)
}
