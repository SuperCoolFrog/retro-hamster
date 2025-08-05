package states

import (
	"retro-hamster/internal/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type WheelUpgradePauseState struct {
	Game       *models.Game
	WheelState *WheelState
}

func (s *WheelUpgradePauseState) OnTransition() {}

func (s *WheelUpgradePauseState) Update() error {
	return nil
}

func (s *WheelUpgradePauseState) Draw(screen *ebiten.Image) {

	s.WheelState.Draw(screen)

}
