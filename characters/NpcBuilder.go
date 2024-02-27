package characters

import (
	"hello.go.game/items"
	"hello.go.game/spells"
)

type NpcBuilder struct {
	npc *Npc
}

func (builder *NpcBuilder) Create() *NpcBuilder {
	builder.npc = &Npc{}
	builder.npc.Alive = true

	return builder
}

func (builder *NpcBuilder) CreateWarrior() *NpcBuilder {
	builder.npc = &Npc{}
	builder.npc.CharacterClass = "Warrior"
	builder.npc.AttackModifier = 1.7
	builder.npc.DefenseModifier = 1.5
	builder.npc.MagicDefenseModifier = 0.85
	builder.npc.Alive = true

	return builder
}

func (builder *NpcBuilder) CreateMage() *NpcBuilder {
	builder.npc = &Npc{}
	builder.npc.CharacterClass = "Mage"
	builder.npc.AttackModifier = 1.0
	builder.npc.DefenseModifier = 0.7
	builder.npc.MagicDefenseModifier = 2.0
	builder.npc.Alive = true

	return builder
}

func (builder *NpcBuilder) CreateSkeleton() *NpcBuilder {
	builder.npc = &Npc{}
	builder.npc.CharacterClass = "Monster"
	builder.npc.AttackModifier = 1.5
	builder.npc.DefenseModifier = 1.1
	builder.npc.MagicDefenseModifier = 1.1
	builder.npc.Alive = true

	return builder
}

func (builder *NpcBuilder) CreateSlime() *NpcBuilder {
	builder.npc = &Npc{}
	builder.npc.CharacterClass = "Monster"
	builder.npc.AttackModifier = 1.0
	builder.npc.DefenseModifier = 0.9
	builder.npc.MagicDefenseModifier = 0.9
	builder.npc.Alive = true

	return builder
}

func (builder *NpcBuilder) CreateBlackMage() *NpcBuilder {
	builder.npc = &Npc{}
	builder.npc.CharacterClass = "Monster"
	builder.npc.AttackModifier = 1.0
	builder.npc.DefenseModifier = 0.95
	builder.npc.MagicDefenseModifier = 1.85
	builder.npc.Alive = true

	return builder
}

func (builder *NpcBuilder) Name(name string) *NpcBuilder {
	builder.npc.Name = name

	return builder
}

func (builder *NpcBuilder) Hp(hp int) *NpcBuilder {
	builder.npc.Hp = hp

	return builder
}

func (builder *NpcBuilder) Mana(mana int) *NpcBuilder {
	builder.npc.Mana = mana

	return builder
}

func (builder *NpcBuilder) AttackPower(attackPower int) *NpcBuilder {
	builder.npc.AttackPower = attackPower

	return builder
}

func (builder *NpcBuilder) MagicPower(magicPower int) *NpcBuilder {
	builder.npc.MagicPower = magicPower

	return builder
}

func (builder *NpcBuilder) Defense(defense int) *NpcBuilder {
	builder.npc.Defense = defense

	return builder
}

func (builder *NpcBuilder) MagicDefense(magicDefense int) *NpcBuilder {
	builder.npc.MagicDefense = magicDefense

	return builder
}

func (builder *NpcBuilder) Weapon(weapon *items.Weapon) *NpcBuilder {
	builder.npc.Weapon = weapon

	return builder
}

func (builder *NpcBuilder) Spells(spells []spells.Spell) *NpcBuilder {
	builder.npc.Spells = spells

	return builder
}

func (builder *NpcBuilder) Build() *Npc {
	return builder.npc
}
