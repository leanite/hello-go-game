package game

import (
	"fmt"
	"os"

	"hello.go.game/characters"
	"hello.go.game/lists"
)

type GameAction struct {
	heroes   *[]*characters.Npc
	monsters *[]*characters.Npc
}

func NewGameAction(heroes []*characters.Npc, monsters []*characters.Npc) *GameAction {
	return &GameAction{
		heroes:   &heroes,
		monsters: &monsters,
	}
}

func (action *GameAction) HeroesAction() {
	action.groupAction(action.heroes, action.monsters)
}

func (action *GameAction) MonstersAction() {
	action.groupAction(action.monsters, action.heroes)
}

func (action *GameAction) groupAction(attackerGroup *[]*characters.Npc, targetGroup *[]*characters.Npc) {
	for _, attacker := range *attackerGroup {
		action.attackAction(attacker, targetGroup)
	}
}

func (action *GameAction) attackAction(attacker *characters.Npc, targetGroup *[]*characters.Npc) {
	target := GetRandomTargetFrom(*targetGroup)

	action.decideAttackerAction(attacker, target)
	action.updateTargetGroupStatus(target, targetGroup)
}

func (action *GameAction) decideAttackerAction(attacker *characters.Npc, target *characters.Npc) {
	if IsAttackerUsingSpell(attacker) {
		action.useSpell(attacker, target)
	} else {
		action.normalAttack(attacker, target)
	}
}

func (action *GameAction) useSpell(attacker *characters.Npc, target *characters.Npc) {
	spell := GetRandomSpellBy(attacker)
	damage, outOfMana := attacker.UseSpell(spell, target)

	if outOfMana != nil {
		PrintOutOfManaText(spell, attacker)
	} else {
		PrintSpellUsedText(spell, attacker, target)
		PrintDamageText(target, damage)
	}
}

func (action *GameAction) normalAttack(attacker *characters.Npc, target *characters.Npc) {
	damage := attacker.Attack(target)

	PrintAttackText(attacker, target)
	PrintDamageText(target, damage)
}

func (action *GameAction) updateTargetGroupStatus(targetAttacked *characters.Npc, targetGroup *[]*characters.Npc) {
	if !targetAttacked.Alive {
		action.defeatTarget(targetAttacked, targetGroup)
	}

	action.checkIfTargetGroupIsAlive()
}

func (action *GameAction) defeatTarget(target *characters.Npc, targetGroup *[]*characters.Npc) {
	PrintDefeatedText(target)
	lists.RemoveNpcFromList(target, targetGroup) //erro aqui: não tá atualizando a referencia pq é passagem por valor e não ponteiro para
}

func (action *GameAction) checkIfTargetGroupIsAlive() {
	if len(*action.heroes) == 0 {
		action.endBattle("Monsters")
	} else if len(*action.monsters) == 0 {
		action.endBattle("Heroes")
	}
}

func (action *GameAction) endBattle(winner string) {
	PrintVerticalSpacing()
	PrintEndOfBattleText(winner)
	os.Exit(0)
}

func (action *GameAction) PrintLists() {
	fmt.Println("Heroes List:")
	for _, hero := range *action.heroes {
		fmt.Println(hero)
	}

	fmt.Println("Monsters List:")
	for _, monster := range *action.monsters {
		fmt.Println(monster)
	}
}
