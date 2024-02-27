package game

import (
	"hello.go.game/characters"
	"hello.go.game/items"
	"hello.go.game/spells"
)

func BuildHeroes() []*characters.Npc {
	return []*characters.Npc{
		buildWarrior(),
		buildMage(),
	}
}

func BuildMonsters() []*characters.Npc {
	return []*characters.Npc{
		buildSkeleton(),
		buildSlime(),
		buildBlackMage(),
	}
}

func buildWarrior() *characters.Npc {
	return (&characters.NpcBuilder{}).
		CreateWarrior().
		Name("Glenn").
		Hp(120).
		Mana(40).
		AttackPower(18).
		MagicPower(3).
		Defense(13).
		MagicDefense(6).
		Weapon(items.NewWeapon("Great Sword", 15)).
		Spells(spells.CreateWarriorSpells()).
		Build()
}

func buildMage() *characters.Npc {
	return (&characters.NpcBuilder{}).
		CreateMage().
		Name("Azalea").
		Hp(90).
		Mana(110).
		AttackPower(5).
		MagicPower(22).
		Defense(6).
		MagicDefense(19).
		Weapon(items.NewWeapon("Rod", 7)).
		Spells(spells.CreateMageSpells()).
		Build()
}

func buildSkeleton() *characters.Npc {
	return (&characters.NpcBuilder{}).
		CreateSkeleton().
		Name("Skeleton").
		Hp(80).
		Mana(0).
		AttackPower(12).
		MagicPower(0).
		Defense(9).
		MagicDefense(8).
		Weapon(items.NewWeapon("Bone Sword", 12)).
		Build()
}

func buildSlime() *characters.Npc {
	return (&characters.NpcBuilder{}).
		CreateSlime().
		Name("Slime").
		Hp(60).
		Mana(0).
		AttackPower(10).
		MagicPower(0).
		Defense(8).
		MagicDefense(5).
		Build()
}

func buildBlackMage() *characters.Npc {
	return (&characters.NpcBuilder{}).
		CreateBlackMage().
		Name("Black Mage").
		Hp(150).
		Mana(120).
		AttackPower(5).
		MagicPower(20).
		Defense(8).
		MagicDefense(15).
		Weapon(items.NewWeapon("Dark Rod", 8)).
		Spells(spells.CreateBlackMageSpells()).
		Build()
}
