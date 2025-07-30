package main

import (
	"log"
	"retro-hamster/assets"
	"retro-hamster/internal/models"
	"retro-hamster/internal/states"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game = models.NewGame(1920, 1080)
)

func init() {
	game.LoadGameAssets(assets.Assets)
}

func main() {
	ebiten.SetWindowSize(game.ScreenW, game.ScreenH)
	ebiten.SetWindowTitle("Retro Hamster - The Big Wheel")

	// game.ChangeState(&states.PlaceholderState{})
	game.ChangeState(&states.WheelState{
		Game: game,
	})

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
