package states

import (
	"fmt"
	"retro-hamster/assets"
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
	src := s.Game.ImageAssets[assets.AssetKey_WinLose_PNG]

	if s.Win {
		models.DrawAssetSprite(src.Image, screen, float64(s.Game.ScreenW)/2-float64(assets.Sprite_Win.W)/2, 0, assets.Sprite_Win)
	} else {
		models.DrawAssetSprite(src.Image, screen, float64(s.Game.ScreenW)/2-float64(assets.Sprite_Lose.W)/2, 0, assets.Sprite_Lose)
	}
}
