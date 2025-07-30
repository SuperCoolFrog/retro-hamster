package models

import "github.com/hajimehoshi/ebiten/v2"

type IGameState interface {
	OnTransition()
	Update() error
	Draw(screen *ebiten.Image)
}
