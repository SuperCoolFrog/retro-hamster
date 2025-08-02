package states

import (
	"fmt"
	"retro-hamster/internal/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverState struct {
	Game *models.Game
}

func (s *GameOverState) OnTransition() {
	fmt.Println("GAME OVER")
}

func (s *GameOverState) Update() error {
	return nil
}

func (s *GameOverState) Draw(screen *ebiten.Image) {

}
