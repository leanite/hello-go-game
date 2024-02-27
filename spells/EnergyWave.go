package spells

import "hello.go.game/numbers"

type EnergyWave struct {
	SpellBase
}

func NewEnergyWave() EnergyWave {
	return EnergyWave{SpellBase: SpellBase{Name: "EnergyWave", ManaCost: 15, BaseDamage: 10, SpellModifier: 2.0}}
}

func (ew EnergyWave) GetName() string {
	return ew.Name
}

func (ew EnergyWave) GetBaseDamage() int {
	return ew.BaseDamage
}

func (ew EnergyWave) GetSpellModifier() float32 {
	return ew.SpellModifier
}

func (ew EnergyWave) GetManaCost() int {
	return ew.ManaCost
}

func (ew EnergyWave) Equal(other Spell) bool {
	if otherEnergyWave, ok := other.(EnergyWave); ok {
		return ew.Name == otherEnergyWave.Name &&
			ew.BaseDamage == otherEnergyWave.BaseDamage &&
			ew.ManaCost == otherEnergyWave.ManaCost &&
			ew.SpellModifier == otherEnergyWave.SpellModifier
	}
	return false
}

func (ew EnergyWave) CalculateSpellDamage(caster SpellCaster) int {
	return numbers.Round(float32(caster.GetMagicPower())*ew.SpellModifier) + ew.BaseDamage
}
