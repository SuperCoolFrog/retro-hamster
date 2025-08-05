package states

import (
	"fmt"
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/levels"
	"retro-hamster/internal/models"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	NumOfEnemies int

	parallaxer *models.Parallaxer

	fader *models.Fader
}

func (s *WheelState) OnTransition() {
	if s.animations == nil {
		s.animations = models.NewSceneAnimations()
	}

	if s.Spawns == nil {
		s.Spawns = make([]*models.Spawn, 0)
	}

	s.ham = models.NewHamster(s.Game)

	s.setupAllLevels()

	if s.CurrentLevel == -1 {
		s.loadLevel(0, 0)
	}

	s.parallaxer = models.NewParallaxer()
	s.parallaxer.AddDetails(0, 0, 0, float64(assets.Sprite_Background.W), float64(assets.Sprite_Background.H), assets.Sprite_Background, s.Game.ImageAssets[assets.AssetKey_Background_1_PNG])
	s.parallaxer.AddDetails(.25, 0, 0, float64(assets.Sprite_Background.W), float64(assets.Sprite_Background.H), assets.Sprite_Background, s.Game.ImageAssets[assets.AssetKey_Background_2_PNG])
	s.parallaxer.AddDetails(.75, 0, 0, float64(assets.Sprite_Background.W), float64(assets.Sprite_Background.H), assets.Sprite_Background, s.Game.ImageAssets[assets.AssetKey_Background_3_PNG])
	s.parallaxer.AddDetails(1, 0, 0, float64(assets.Sprite_Background.W), float64(assets.Sprite_Background.H), assets.Sprite_Background, s.Game.ImageAssets[assets.AssetKey_Background_4_PNG])

	s.fader = models.NewFader(false, float64(s.Game.ScreenW), float64(s.Game.ScreenH))
	s.Update()
	s.fader.Start()
}

func (s *WheelState) Update() error {
	s.fader.Update()
	if s.fader.IsFading {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && s.ham.Blocked != models.DIRECTION_LEFT {
		s.angle += .75 * math.Pi / 180
		// s.angle += (math.Max(s.ham.Momentum.Ratio(), .25) * .75) * math.Pi / 180
		s.ham.Direction = models.DIRECTION_LEFT
		s.ham.IsRunning = true
		s.ham.Blocked = models.DIRECTION_NONE
		s.parallaxer.Update(models.DIRECTION_RIGHT)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) && s.ham.Blocked != models.DIRECTION_RIGHT {
		s.angle -= .75 * math.Pi / 180
		// s.angle -= (math.Max(s.ham.Momentum.Ratio(), .25) * .75) * math.Pi / 180
		s.ham.Direction = models.DIRECTION_RIGHT
		s.ham.IsRunning = true
		s.ham.Blocked = models.DIRECTION_NONE
		s.parallaxer.Update(models.DIRECTION_LEFT)
	} else {
		s.ham.IsRunning = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.ham.InitJump()
	}

	s.checkLevel()

	s.ham.Update(s.angle)
	s.animations.Update()
	s.updateSpawns(s.Spawns)

	// for _, spawn := range s.Spawns {
	// 	x := (float64(s.Game.ScreenW) / 2) - spawn.W
	// 	y := float64(s.Game.ScreenH/2) + wheelH/2
	// 	spawn.Update(x, y, s.angle)

	// 	if spawn.IsAlive && s.ham.GetCollisionRect().Intersects(spawn.GetCollisionRect()) {
	// 		spawn.OnCollision(s.ham)
	// 	}
	// }

	// In Place Remove dead
	writeIndex := 0
	for _, spawn := range s.Spawns {
		if spawn.IsAlive {
			s.Spawns[writeIndex] = spawn
			writeIndex++
		}
	}
	s.Spawns = s.Spawns[:writeIndex]

	if models.BOSS_HAS_INIT && models.BOSS_HEALTH < 1 {
		s.Game.ChangeState(&GameOverState{
			Win:  true,
			Game: s.Game,
		})
	} else if s.ham.Health < 1 {
		s.Game.ChangeState(&GameOverState{
			Win:  false,
			Game: s.Game,
		})
	}

	return nil
}

func (s *WheelState) updateSpawns(spawns []*models.Spawn) {
	for _, spawn := range spawns {
		x := (float64(s.Game.ScreenW) / 2) // - spawn.W
		y := float64(s.Game.ScreenH/2) + wheelH/2
		spawn.Update(x, y, s.angle)

		if spawn.IsAlive && s.ham.GetCollisionRect().IntersectsPolygon(spawn.GetHitBox()) {
			spawn.OnCollision(s.ham)
		}
	}
}

func (s *WheelState) Draw(screen *ebiten.Image) {
	/* Background */
	s.parallaxer.Draw(screen)

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

	models.DrawAssetSpriteWithOptions(wheelPng.Image, screen, assets.Sprite_Wheel, op)
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

	lvlTextOpts := &text.DrawOptions{}
	lvlTextOpts.GeoM.Translate(float64(s.Game.ScreenH)/20, float64(s.Game.ScreenH)/20)

	models.WriteToScene(fmt.Sprintf("Level: %d - %d", s.CurrentLevel+1, s.CurrentRound+1), float64(s.Game.ScreenH)/20, s.Game, screen, lvlTextOpts)

	s.fader.Draw(screen)
}

func (s *WheelState) setupAllLevels() {
	if s.Levels != nil {
		return
	}

	allAndBoss := make([]string, 0)
	allAndBoss = append(allAndBoss, levels.ALL_LEVEL_CHARTS...)
	allAndBoss = append(allAndBoss, levels.BOSS_LEVELS...)

	for i := range allAndBoss {
		level := &models.Level{
			Rounds: map[int][]*models.Spawn{},
		}

		chart := allAndBoss[i]

		roundsChart := strings.Split(chart, "\n")[1:] // skip first line

		for roundIdx := range roundsChart {
			round := roundsChart[roundIdx]

			spawns := make([]*models.Spawn, 0)

			for idx, spawnSymbol := range round {
				if string(spawnSymbol) != " " && string(spawnSymbol) != "_" {
					spawns = append(spawns, models.SymbolToSpawnMap[string(spawnSymbol)](idx))
				}
			}

			level.Rounds[roundIdx] = spawns
		}

		s.Levels = append(s.Levels, level)
	}
}

func (s *WheelState) appendLevels(lvls []string) {
	for i := range lvls {
		level := &models.Level{
			Rounds: map[int][]*models.Spawn{},
		}

		chart := lvls[i]

		roundsChart := strings.Split(chart, "\n")[1:] // skip first line

		for roundIdx := range roundsChart {
			round := roundsChart[roundIdx]

			spawns := make([]*models.Spawn, 0)

			for idx, spawnSymbol := range round {
				if string(spawnSymbol) != " " && string(spawnSymbol) != "_" {
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
		s.ham.Blocked = models.DIRECTION_NONE
	}
}

func (s *WheelState) loadLevel(levelIdx, roundIdx int) {
	level := s.Levels[levelIdx]
	roundSpawns := level.Rounds[roundIdx]

	for _, spawn := range roundSpawns {
		spawn.Init()
	}

	s.updateSpawns(roundSpawns)

	for _, spawn := range roundSpawns {
		spawn.SetHamsterRelativeOffset(s.ham.LogicalAngle)
	}

	s.Spawns = roundSpawns
	s.CurrentLevel = levelIdx
	s.CurrentRound = roundIdx
}

func (s *WheelState) nextRoundOrLevel() {
	if len(s.Levels) == 0 {
		return
	}

	nextLevel := s.CurrentLevel
	level := s.Levels[nextLevel]
	nextRound := s.CurrentRound + 1

	if nextRound >= len(level.Rounds) {

		nextLevel = s.CurrentLevel + 1

		if nextLevel >= len(s.Levels) {

			if models.BOSS_HAS_INIT && models.BOSS_HEALTH > 0 {
				s.appendLevels(levels.BOSS_LEVELS)
			} else {
				s.Game.ChangeState(&GameOverState{
					Win:  true,
					Game: s.Game,
				})
			}
			return
		}

		nextRound = 0
	}

	s.loadLevel(nextLevel, nextRound)
}
