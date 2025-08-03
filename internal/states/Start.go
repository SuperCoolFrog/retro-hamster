package states

import (
	"fmt"
	"retro-hamster/assets"
	"retro-hamster/internal/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StartState struct {
	Game      *models.Game
	menu      []*models.GameAction
	startItem *models.GameAction
	exitItem  *models.GameAction
}

func (s *StartState) OnTransition() {

	s.menu = make([]*models.GameAction, 0)

	s.startItem = &models.GameAction{
		Id:       1,
		Text:     "Start",
		Selected: false,
		X:        float64(s.Game.ScreenW)/2 - float64(assets.Sprite_Button.W)/2,
		Y:        (float64(s.Game.ScreenH) * 2 / 3),
		W:        float64(assets.Sprite_Button.W),
		H:        float64(assets.Sprite_Button.H),
	}
	s.menu = append(s.menu, s.startItem)

	s.exitItem = &models.GameAction{
		Id:       2,
		Text:     "Exit",
		Selected: false,
		X:        float64(s.Game.ScreenW)/2 - float64(assets.Sprite_Button.W)/2,
		Y:        float64(s.Game.ScreenH)*2/3 + float64(assets.Sprite_Button.H),
		W:        float64(assets.Sprite_Button.W),
		H:        float64(assets.Sprite_Button.H),
	}
	s.menu = append(s.menu, s.exitItem)

	s.Game.CURSOR_X = -1
	s.Game.CURSOR_Y = -1
}

func (s *StartState) Update() error {

	s.Game.OnMouseMoved(func(x, y int) {
		s.startItem.Focused = s.startItem.Contains(float64(x), float64(y))
		s.exitItem.Focused = s.exitItem.Contains(float64(x), float64(y))
	})

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if s.startItem.Focused {
			s.Game.ChangeState(&WheelState{
				Game:         s.Game,
				CurrentLevel: -1,
				CurrentRound: -1,
			})
			// s.Game.ChangeState(&BossPhase1State{
			// 	Game: s.Game,
			// })
			return nil
		} else if s.exitItem.Focused {
			s.Game.UTIL_EXIT = true
			return fmt.Errorf("clean exit")
		}
	}

	return nil
}

func (s *StartState) Draw(screen *ebiten.Image) {
	/* Background */
	bg := s.Game.ImageAssets[assets.AssetKey_Start_PNG]
	models.DrawAssetSprite(bg.Image, screen, 0, 0, assets.Sprite_Start)

	models.DrawButton(s.Game, screen, s.startItem)
	models.DrawButton(s.Game, screen, s.exitItem)
}
