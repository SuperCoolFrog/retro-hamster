package states

import (
	"fmt"
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/levels"
	"retro-hamster/internal/models"
	"retro-hamster/internal/scenes"
	"strings"

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

	Levels       []*models.Level
	CurrentLevel int
	CurrentRound int
}

func (s *WheelState) OnTransition() {
	if s.animations == nil {
		s.animations = models.NewSceneAnimations()
	}

	if s.Spawns == nil {
		s.Spawns = make([]*models.Spawn, 0)
	}

	s.ham = models.NewHamster(s.Game)

	// snake := &models.Animation{
	// 	FPS:          12,
	// 	CurrentFrame: 0,
	// 	Details:      assets.AnimationSnake,
	// 	X:            0,
	// 	Y:            0,
	// }

	// snakeSpawn := models.NewSpawn(5, wheelRadius+float64(assets.AnimationSnake.InitialSprite.W)/4, snake)
	// snakeSpawn.Direction = models.DIRECTION_LEFT
	// snakeSpawn.Power = 50
	// snakeSpawn.OnCollision = func(hamster *models.Hamster) {
	// 	damage := hamster.Momentum.Current - snakeSpawn.Power
	// 	hamster.Momentum.Current -= snakeSpawn.Power

	// 	if damage < 0 {
	// 		hamster.Health -= 1
	// 	}

	// 	snakeSpawn.IsAlive = false
	// }
	// s.Spawns = append(s.Spawns, snakeSpawn)

	// mod := float64(assets.AnimationSeed.InitialSprite.W / 4)
	// wheelRadiusModified := wheelRadius - mod
	// space := 512.0
	// for i := range 5 {
	// 	angle := float64(i) * space / (wheelRadiusModified)
	// 	// angle -= math.Pi / 2 /* THis will translate to top as starting point */
	// 	// angle := float64(i) * 2 * math.Pi / 5
	// 	seed := models.NewSpawn(angle, wheelRadiusModified, &models.Animation{
	// 		FPS:          0,
	// 		CurrentFrame: 0,
	// 		Details:      assets.AnimationSeed,
	// 	})
	// 	seed.Power = 5
	// 	seed.OnCollision = func(ham *models.Hamster) {
	// 		s.ham.XP.Current += seed.Power
	// 		seed.IsAlive = false
	// 	}
	// 	s.Spawns = append(s.Spawns, seed)
	// }

	s.setupAllLevels()

	if s.CurrentLevel == -1 {
		s.loadLevel(0, 0)
	}
}

func (s *WheelState) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		// s.angle += 2 * math.Pi / 180
		s.angle += 1 * math.Pi / 180
		s.ham.Direction = models.DIRECTION_LEFT
		s.ham.IsRunning = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		// s.angle -= 2 * math.Pi / 180
		s.angle -= 1 * math.Pi / 180
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
		x := (float64(s.Game.ScreenW) / 2) - spawn.W
		y := float64(s.Game.ScreenH/2) + wheelH/2
		spawn.Update(x, y, s.angle)

		if spawn.IsAlive && s.ham.GetCollisionRect().Intersects(spawn.GetCollisionRect()) {
			spawn.OnCollision(s.ham)
		}
	}

	// In Place Remove dead
	writeIndex := 0
	for _, spawn := range s.Spawns {
		if spawn.IsAlive { // Condition: keep even numbers
			s.Spawns[writeIndex] = spawn
			writeIndex++
		}
	}
	s.Spawns = s.Spawns[:writeIndex]

	s.checkLevel()

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

func (s *WheelState) setupAllLevels() {
	if s.Levels != nil {
		return
	}

	for i := range levels.ALL_LEVEL_CHARTS {
		level := &models.Level{
			Rounds: map[int][]*models.Spawn{},
		}

		chart := levels.ALL_LEVEL_CHARTS[i]
		fmt.Printf("%s", chart)

		roundsChart := strings.Split(chart, "\n")[1:] // skip first line

		for roundIdx := range roundsChart {
			round := roundsChart[roundIdx]

			spawns := make([]*models.Spawn, 0)

			for idx, spawnSymbol := range round {
				if string(spawnSymbol) != " " {
					spawns = append(spawns, models.SymbolToSpawnMap[string(spawnSymbol)](idx))
				}
			}

			level.Rounds[roundIdx] = spawns
		}

		s.Levels = append(s.Levels, level)
	}
}

func (s *WheelState) checkLevel() {
	canAdvance := true

	for _, spawn := range s.Spawns {
		canAdvance = canAdvance && (spawn == nil || (!spawn.IsAlive || spawn.IsObstacle))
	}

	if canAdvance {
		s.nextRoundOrLevel()
	}
}

func (s *WheelState) loadLevel(levelIdx, roundIdx int) {
	level := s.Levels[levelIdx]
	roundSpawns := level.Rounds[roundIdx]
	s.Spawns = roundSpawns
	s.CurrentLevel = levelIdx
	s.CurrentRound = roundIdx
}

func (s *WheelState) nextRoundOrLevel() {
	nextLevel := s.CurrentLevel
	level := s.Levels[nextLevel]
	nextRound := s.CurrentRound + 1

	if nextRound >= len(level.Rounds) {

		nextLevel = s.CurrentLevel + 1

		if nextLevel >= len(s.Levels) {
			s.Game.ChangeState(&GameOverState{
				Game: s.Game,
			})
			return
		}

		nextRound = 0
	}

	s.loadLevel(nextLevel, nextRound)
}
