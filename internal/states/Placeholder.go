package states

import (
	"retro-hamster/internal/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlaceholderState struct {
	Game *models.Game
}

func (s *PlaceholderState) OnTransition() {}

func (s *PlaceholderState) Update() error {
	return nil
}

func (s *PlaceholderState) Draw(screen *ebiten.Image) {

}
