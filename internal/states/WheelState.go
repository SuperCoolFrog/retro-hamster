package states

import (
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/models"
	"retro-hamster/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

const HAMSTER_DIRECTION_RIGHT = 1
const HAMSTER_DIRECTION_LEFT = -1

type WheelState struct {
	Game             *models.Game
	angle            float64
	hamsterIsRunning bool
	hamsterDirection int
	animations       *models.SceneAnimations
}

func (s *WheelState) OnTransition() {
	if s.animations == nil {
		s.animations = models.NewSceneAnimations()
	}
	s.hamsterDirection = HAMSTER_DIRECTION_RIGHT

	// For Testing
	// scale := 2.0
	// wheelW := float64(assets.Sprite_Wheel.W) * scale
	// wheelH := float64(assets.Sprite_Wheel.H) * scale
	hamsterAnim := &models.Animation{
		FPS:          12,
		CurrentFrame: 0,
		Details:      assets.AnimationHamsterRun,
		X:            s.Game.ScreenW/2 - assets.AnimationHamsterRun.InitialSprite.W/2,
		Y:            s.Game.ScreenH/2 - int(float64(assets.AnimationHamsterRun.InitialSprite.H)/1.25), //+wheelH/2,
	}
	s.animations.AddSceneAnimation(hamsterAnim)
}

func (s *WheelState) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.angle += 2 * math.Pi / 180
		s.hamsterDirection = HAMSTER_DIRECTION_LEFT
		s.hamsterIsRunning = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.angle -= 2 * math.Pi / 180
		s.hamsterDirection = HAMSTER_DIRECTION_RIGHT
		s.hamsterIsRunning = true
	} else {
		s.hamsterIsRunning = false
	}

	s.animations.Update()

	return nil
}

func (s *WheelState) Draw(screen *ebiten.Image) {

	/* #region wheel */
	wheelPng := s.Game.ImageAssets[assets.AssetKey_Wheel_PNG]
	op := ebiten.DrawImageOptions{}

	scale := 2.0
	wheelW := float64(assets.Sprite_Wheel.W) * scale
	wheelH := float64(assets.Sprite_Wheel.H) * scale

	// Move the origin to the center of the image before rotating
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(-wheelW/2, -wheelH/2)
	op.GeoM.Rotate(s.angle)
	op.Filter = ebiten.FilterLinear
	// Move it to the screen center after rotation
	op.GeoM.Translate(float64(s.Game.ScreenW)/2, float64(s.Game.ScreenH)/2+wheelH/2)

	scenes.DrawAssetSpriteWithOptions(wheelPng.Image, screen, assets.Sprite_Wheel, op)
	/* #endregion wheel */

	/* #region Animations */
	animationSprites := s.animations.GetAllCurrentSprites()
	for i := range animationSprites {
		img := animationSprites[i]
		if animSs, animSsExists := s.Game.ImageAssets[img.AssetKey]; animSsExists {
			scenes.DrawSprite(animSs.Image, screen, img.TargetX, img.TargetY, img.X, img.Y, img.W, img.H)
		}
	}
	/* #endregion Animations */
}
