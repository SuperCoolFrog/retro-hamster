package states

import (
	"fmt"
	"retro-hamster/internal/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverState struct {
	Game *models.Game
	Win  bool
}

func (s *GameOverState) OnTransition() {
	if s.Win {
		fmt.Println("YOU WIN")
	} else {
		fmt.Println("YOU LOSE")
	}
}

func (s *GameOverState) Update() error {
	return nil
}

func (s *GameOverState) Draw(screen *ebiten.Image) {

}
