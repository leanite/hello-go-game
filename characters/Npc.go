package characters

import (
	"hello.go.game/errors"
	"hello.go.game/etc"
	"hello.go.game/items"
	"hello.go.game/numbers"
	"hello.go.game/spells"
)

type Npc struct {
	Name                 string
	CharacterClass       string
	Alive                bool
	Hp                   int
	Mana                 int
	AttackPower          int
	AttackModifier       float32
	MagicPower           int
	Defense              int
	DefenseModifier      float32
	MagicDefense         int
	MagicDefenseModifier float32
	Weapon               *items.Weapon
	Spells               []spells.Spell
}

func (npc *Npc) UseSpell(spell spells.Spell, target *Npc) (int, error) {
	if npc.canUseSpell(spell) {
		npc.Mana -= spell.GetManaCost()
	} else {
		return 0, &errors.OutOfManaError{Message: "Out of mana!"}
	}

	return target.takeDamage(spell.CalculateSpellDamage(npc), etc.MAGIC), nil
}

func (npc Npc) canUseSpell(spell spells.Spell) bool {
	return npc.Mana >= spell.GetManaCost()
}

func (npc Npc) GetMagicPower() int {
	return npc.MagicPower
}

func (npc *Npc) Attack(target *Npc) int {
	return target.takeDamage(npc.calculateAttackDamage(), etc.PHYSICAL)
}

func (npc Npc) calculateAttackDamage() int {
	var baseDamage int = npc.getBaseDamage()
	var randomDamage int = npc.generateRandomAttackDamage()

	return baseDamage + randomDamage
}

func (npc Npc) getBaseDamage() int {
	return numbers.Round(float32(npc.AttackPower) * npc.AttackModifier)
}

func (npc Npc) generateRandomAttackDamage() int {
	var weaponRandomDamage int = 0

	if npc.HasWeapon() {
		weaponRandomDamage = numbers.RandomWithLowerOne(npc.Weapon.AttackPower())
	}

	return weaponRandomDamage
}

func (npc *Npc) takeDamage(damage int, damageType etc.DamageType) int {
	var damageTaken int = npc.calculateDamageTaken(damage, damageType)

	npc.loseHitPoints(damageTaken)
	npc.checkAliveStatus()

	return damageTaken
}

func (npc Npc) calculateDamageTaken(damage int, damageType etc.DamageType) int {
	var trueDamage int = npc.calculateTrueDamageByDamageType(damage, damageType)

	return checkNegativeDamage(trueDamage)
}

func checkNegativeDamage(damage int) int {
	if damage <= 0 {
		return 1
	} else {
		return damage
	}
}

func (npc Npc) calculateTrueDamageByDamageType(damage int, damageType etc.DamageType) int {
	switch damageType {
	case etc.PHYSICAL:
		return (damage - npc.calculateDefense(npc.Defense, npc.DefenseModifier))
	case etc.MAGIC:
		return (damage - npc.calculateDefense(npc.MagicDefense, npc.MagicDefenseModifier))
	default:
		return 1
	}
}

func (npc Npc) calculateDefense(attribute int, modifier float32) int {
	return numbers.Round(float32(attribute) * modifier)
}

func (npc *Npc) loseHitPoints(trueDamage int) {
	npc.Hp -= trueDamage
}

func (npc *Npc) checkAliveStatus() {
	if npc.Hp <= 0 {
		npc.Alive = false
	}
}

func (npc Npc) HasWeapon() bool {
	return npc.Weapon != nil
}

func (npc Npc) HasSpells() bool {
	return len(npc.Spells) != 0
}

func (npc *Npc) Equal(other *Npc) bool {
	return npc.Name == other.Name
}
