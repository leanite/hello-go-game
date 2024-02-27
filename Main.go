package main

import (
	"hello.go.game/game"
)

func main() {
	game := game.NewGame()
	game.Loop()
}
