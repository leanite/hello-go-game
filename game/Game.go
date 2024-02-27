package game

type Game struct {
	turn *GameTurn
}

func NewGame() *Game {
	return &Game{
		turn: NewGameTurn(BuildHeroes(), BuildMonsters()),
	}
}

func (game *Game) Loop() {
	for {
		game.battleTurn()
	}
}

func (game *Game) battleTurn() {
	game.turn.HeroesTurn()
	game.turn.MonstersTurn()
	PrintVerticalSpacing()
}
