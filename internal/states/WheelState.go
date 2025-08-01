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

var (
	wheelScale  = 2.0
	wheelW      = float64(assets.Sprite_Wheel.W) * wheelScale
	wheelH      = float64(assets.Sprite_Wheel.H) * wheelScale
	wheelRadius = 1100.0
)

type WheelState struct {
	Game               *models.Game
	angle              float64
	hamsterIsRunning   bool
	hamsterDirection   int
	hamsterAnimationId int
	animations         *models.SceneAnimations
	Spawns             []*models.Spawn
}

func (s *WheelState) OnTransition() {
	if s.animations == nil {
		s.animations = models.NewSceneAnimations()
	}

	if s.Spawns == nil {
		s.Spawns = make([]*models.Spawn, 0)
	}

	s.hamsterDirection = HAMSTER_DIRECTION_RIGHT
	s.hamsterAnimationId = -1

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
		s.hamsterDirection = HAMSTER_DIRECTION_LEFT
		s.hamsterIsRunning = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.angle -= 2 * math.Pi / 180
		s.hamsterDirection = HAMSTER_DIRECTION_RIGHT
		s.hamsterIsRunning = true
	} else {
		s.hamsterIsRunning = false
	}

	if s.hamsterIsRunning {
		if s.hamsterAnimationId == -1 {
			s.addHamsterRunAnimation()
		}
	} else {
		if s.hamsterAnimationId != -1 {
			s.removeHamsterRunAnimation()
		}
	}

	s.animations.Update()

	for _, spawn := range s.Spawns {
		x := float64(s.Game.ScreenW) / 2.75
		y := float64(s.Game.ScreenH/2) + wheelH/2.1
		spawn.Update(x, y, s.angle)
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

	if !s.hamsterIsRunning {
		s.drawStaticHamster(screen)
	}

	/* #region Animations */
	animationSprites := s.animations.GetAllCurrentSprites()
	for i := range animationSprites {
		img := animationSprites[i]
		if animSs, animSsExists := s.Game.ImageAssets[img.AssetKey]; animSsExists {
			if img.AssetKey == assets.AssetKey_Hamster_Run_PNG && s.hamsterDirection == HAMSTER_DIRECTION_LEFT {
				hOpts := ebiten.DrawImageOptions{}
				hOpts.GeoM.Scale(float64(s.hamsterDirection), 1)
				hOpts.GeoM.Translate(float64(img.TargetX), float64(img.TargetY))
				scenes.DrawAssetSpriteWithOptionsWithBoundsCorrect(animSs.Image, screen, img.AssetSprite, hOpts)
			} else {
				scenes.DrawSprite(animSs.Image, screen, img.TargetX, img.TargetY, img.X, img.Y, img.W, img.H)
			}
		}
	}
	/* #endregion Animations */

	for _, spawn := range s.Spawns {
		spawn.Draw(s.Game, screen)
	}
}

func (s *WheelState) getHamsterPosition() (x, y float64) {
	x = float64(s.Game.ScreenW/2 - assets.AnimationHamsterRun.InitialSprite.W/2)
	y = float64(s.Game.ScreenH/2 - int(float64(assets.AnimationHamsterRun.InitialSprite.H)/1.25))
	return x, y
}

func (s *WheelState) drawStaticHamster(screen *ebiten.Image) {
	x, y := s.getHamsterPosition()
	src := s.Game.ImageAssets[assets.AssetKey_Static_PNG]

	if s.hamsterDirection == HAMSTER_DIRECTION_RIGHT {
		scenes.DrawAssetSprite(src.Image, screen, x, y, assets.Sprite_Hamster)
	} else {
		hOpts := ebiten.DrawImageOptions{}
		hOpts.GeoM.Scale(float64(s.hamsterDirection), 1)
		hOpts.GeoM.Translate(float64(x), float64(y))
		scenes.DrawAssetSpriteWithOptionsWithBoundsCorrect(src.Image, screen, assets.Sprite_Hamster, hOpts)
	}
}

func (s *WheelState) addHamsterRunAnimation() {
	x, y := s.getHamsterPosition()
	hamsterAnim := &models.Animation{
		FPS:          12,
		CurrentFrame: 0,
		Details:      assets.AnimationHamsterRun,
		X:            x,
		Y:            y,
	}
	s.hamsterAnimationId = s.animations.AddSceneAnimation(hamsterAnim)
}

func (s *WheelState) removeHamsterRunAnimation() {
	s.animations.RemoveAnimation(s.hamsterAnimationId)
	s.hamsterAnimationId = -1
}
