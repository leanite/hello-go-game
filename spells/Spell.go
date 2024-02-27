package spells

type SpellBase struct {
	Name          string
	ManaCost      int
	BaseDamage    int
	SpellModifier float32
}

type Spell interface {
	GetName() string
	GetManaCost() int
	GetBaseDamage() int
	GetSpellModifier() float32
	Equal(other Spell) bool
	CalculateSpellDamage(caster SpellCaster) int
}

func CreateWarriorSpells() []Spell {
	return []Spell{NewDoubleAttack()}
}

func CreateMageSpells() []Spell {
	return []Spell{NewFireball(), NewEnergyWave()}
}

func CreateBlackMageSpells() []Spell {
	return []Spell{NewNightmare()}
}
