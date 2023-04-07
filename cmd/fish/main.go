package main

import (
	"fish/internal/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH = 640
	HEIGHT = 480
)

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("FISH!")

	game, err := game.New(WIDTH, HEIGHT)
	if err != nil {
		log.Fatal(err)
	}

	if err = ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
