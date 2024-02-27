package game

import (
	"hello.go.game/characters"
)

type GameTurn struct {
	gameAction *GameAction
}

func NewGameTurn(heroes []*characters.Npc, monsters []*characters.Npc) *GameTurn {
	return &GameTurn{
		gameAction: NewGameAction(heroes, monsters),
	}
}

func (turn *GameTurn) HeroesTurn() {
	PrintTurnText("Heroes")
	turn.gameAction.HeroesAction()
}

func (turn *GameTurn) MonstersTurn() {
	PrintTurnText("Monsters")
	turn.gameAction.MonstersAction()
}
