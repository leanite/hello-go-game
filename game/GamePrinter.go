package game

import (
	"fmt"

	"hello.go.game/characters"
	"hello.go.game/spells"
)

func PrintTurnText(group string) {
	fmt.Println("c====[] ", group, "' turn   =====>")
}

func PrintVerticalSpacing() {
	fmt.Println()
}

func PrintEndOfBattleText(winner string) {
	fmt.Println(
		fmt.Sprintf("** %s won the battle! **\n", winner),
	)
}

func PrintSpellUsedText(spell spells.Spell, attacker *characters.Npc, target *characters.Npc) {
	fmt.Println(
		fmt.Sprintf("%s [%s] used %s on %s [%s]!\n", attacker.Name, attacker.CharacterClass, spell.GetName(), target.Name, target.CharacterClass),
	)
}

func PrintOutOfManaText(spell spells.Spell, attacker *characters.Npc) {
	fmt.Println(
		fmt.Sprintf("%s [%s] tried to use %s, but %s was out of mana!\n", attacker.Name, attacker.CharacterClass, spell.GetName(), attacker.Name),
	)
}

func PrintAttackText(attacker *characters.Npc, target *characters.Npc) {
	fmt.Println(
		fmt.Sprintf("%s [%s] attacked %s [%s]!\n", attacker.Name, attacker.CharacterClass, target.Name, target.CharacterClass),
	)
}

func PrintDamageText(target *characters.Npc, damage int) {
	fmt.Println(
		fmt.Sprintf("%s took %d point(s) of damage!\n", target.Name, damage),
	)
	fmt.Println(
		fmt.Sprintf("%s HP: %d\n", target.Name, target.Hp),
	)
}

func PrintDefeatedText(defeated *characters.Npc) {
	fmt.Println(
		fmt.Sprintf("%s [%s] was defeated!\n", defeated.Name, defeated.CharacterClass),
	)
}
