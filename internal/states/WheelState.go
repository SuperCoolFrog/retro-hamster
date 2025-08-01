package states

import (
	"fmt"
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/models"
	"retro-hamster/internal/scenes"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// const HAMSTER_DIRECTION_RIGHT = 1
// const HAMSTER_DIRECTION_LEFT = -1

var (
	wheelScale  = 2.0
	wheelW      = float64(assets.Sprite_Wheel.W) * wheelScale
	wheelH      = float64(assets.Sprite_Wheel.H) * wheelScale
	wheelRadius = 1100.0
)

type WheelState struct {
	Game       *models.Game
	angle      float64
	ham        *models.Hamster
	animations *models.SceneAnimations
	Spawns     []*models.Spawn
}

func (s *WheelState) OnTransition() {
	if s.animations == nil {
		s.animations = models.NewSceneAnimations()
	}

	if s.Spawns == nil {
		s.Spawns = make([]*models.Spawn, 0)
	}

	s.ham = models.NewHamster(s.Game)

	snake := &models.Animation{
		FPS:          12,
		CurrentFrame: 0,
		Details:      assets.AnimationSnake,
		X:            0,
		Y:            0,
	}

	// s.animations.AddSceneAnimation(snake)

	snakeSpawn := models.NewSpawn(5, wheelRadius, snake)
	snakeSpawn.Direction = models.DIRECTION_LEFT
	s.Spawns = append(s.Spawns, snakeSpawn)

}

func (s *WheelState) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.angle += 2 * math.Pi / 180
		s.ham.Direction = models.DIRECTION_LEFT
		s.ham.IsRunning = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.angle -= 2 * math.Pi / 180
		s.ham.Direction = models.DIRECTION_RIGHT
		s.ham.IsRunning = true
	} else {
		s.ham.IsRunning = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.ham.InitJump()
	}

	s.ham.Update()
	s.animations.Update()

	for _, spawn := range s.Spawns {
		x := float64(s.Game.ScreenW) / 2.75
		y := float64(s.Game.ScreenH/2) + wheelH/2.1
		spawn.Update(x, y, s.angle)

		if s.ham.GetCollisionRect().Intersects(spawn.GetCollisionRect()) {
			fmt.Printf("Collided: %d\n", time.Now().Nanosecond())
			s.ham.Momentum.Current += 1
		}
	}

	return nil
}

func (s *WheelState) Draw(screen *ebiten.Image) {

	/* #region wheel */
	wheelPng := s.Game.ImageAssets[assets.AssetKey_Wheel_PNG]
	op := ebiten.DrawImageOptions{}

	op.GeoM.Scale(wheelScale, wheelScale)
	// Move the origin to the center of the image before rotating
	op.GeoM.Translate(-wheelW/2, -wheelH/2)
	op.GeoM.Rotate(s.angle)
	op.Filter = ebiten.FilterLinear
	// Move it to the screen center after rotation
	op.GeoM.Translate(float64(s.Game.ScreenW)/2, float64(s.Game.ScreenH)/2+wheelH/2)

	scenes.DrawAssetSpriteWithOptions(wheelPng.Image, screen, assets.Sprite_Wheel, op)
	/* #endregion wheel */

	/* #region Animations */
	// animationSprites := s.animations.GetAllCurrentSprites()
	// for i := range animationSprites {
	// 	img := animationSprites[i]
	// if animSs, animSsExists := s.Game.ImageAssets[img.AssetKey]; animSsExists {
	// if img.AssetKey == assets.AssetKey_Hamster_Run_PNG && s.ham.Direction == models.DIRECTION_LEFT {
	// 	hOpts := ebiten.DrawImageOptions{}
	// 	hOpts.GeoM.Scale(float64(s.ham.Direction), 1)
	// 	hOpts.GeoM.Translate(float64(img.TargetX), float64(img.TargetY))
	// 	scenes.DrawAssetSpriteWithOptionsWithBoundsCorrect(animSs.Image, screen, img.AssetSprite, hOpts)
	// } else {
	// 	scenes.DrawSprite(animSs.Image, screen, img.TargetX, img.TargetY, img.X, img.Y, img.W, img.H)
	// }
	// }
	// }
	/* #endregion Animations */

	s.ham.Draw(screen)

	for _, spawn := range s.Spawns {
		spawn.Draw(s.Game, screen)
	}
}
