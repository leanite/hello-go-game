package game

import (
	"hello.go.game/characters"
	"hello.go.game/numbers"
	"hello.go.game/spells"
)

func GetRandomTargetFrom(targets []*characters.Npc) *characters.Npc {
	targetIndex := numbers.RandomWithLowerZero(len(targets))

	return targets[targetIndex]
}

func IsAttackerUsingSpell(attacker *characters.Npc) bool {
	return attacker.HasSpells() && IsRandomActionSpellAction()
}

func IsRandomActionSpellAction() bool {
	return numbers.RandomWithLowerOne(10) <= 3
}

func GetRandomSpellBy(attacker *characters.Npc) spells.Spell {
	spells := attacker.Spells
	spellIndex := numbers.RandomWithLowerZero(len(spells))

	return spells[spellIndex]
}
