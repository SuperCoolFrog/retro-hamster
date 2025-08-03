package states

import (
	"math"
	"retro-hamster/assets"
	"retro-hamster/internal/models"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type BossPhase1State struct {
	Game       *models.Game
	WheelState *WheelState
	Boss       *models.Spawn
}

func (s *BossPhase1State) OnTransition() {
	s.WheelState = &WheelState{}
	s.WheelState.Game = s.Game
	s.WheelState.CurrentLevel = 0
	s.WheelState.Levels = make([]*models.Level, 0)

	s.WheelState.Spawns = make([]*models.Spawn, 0)

	if s.Boss == nil {
		mod := float64(assets.AnimationBossPhase1.InitialSprite.H / 5)
		wheelRadiusModified := models.WHEEL_RADIUS + mod
		// angle := math.Pi
		// angle := math.Pi / 2
		// angle := (math.Pi / 2) / (wheelRadiusModified)
		angle := 512.0 / (wheelRadiusModified)
		s.Boss = models.NewSpawn(angle, wheelRadiusModified, &models.Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationBossPhase1,
		})
		s.Boss.IsObstacle = true
		s.Boss.IsAlive = true
		s.Boss.Power = 100
		s.Boss.Health = 10
		s.Boss.Direction = models.DIRECTION_NONE
		s.Boss.SkewAngle = math.Pi / 12
		s.Boss.ModHitBox = 1.25
	}
	s.Boss.ActivationCoolDown = time.Second * 5
	s.Boss.LastActivation = time.Now().Add(-time.Minute)
	s.Boss.ActivationCoolDown = time.Second
	s.Boss.OnCollision = func(ham *models.Hamster) {
		if ham.Direction == models.DIRECTION_RIGHT {
			ham.Blocked = models.DIRECTION_RIGHT
		} else {
			ham.Blocked = models.DIRECTION_LEFT
		}
		if time.Since(s.Boss.LastActivation) > s.Boss.ActivationCoolDown {
			damage := ham.Momentum.Current - s.Boss.Power
			ham.Momentum.Current -= s.Boss.Power

			if damage < 0 {
				ham.Health -= 1
			}

			ham.Health -= 1
			s.Boss.LastActivation = time.Now()
		}
	}
	s.WheelState.Spawns = append(s.WheelState.Spawns, s.Boss)

	s.WheelState.OnTransition()

	// s.WheelState.updateSpawns(s.WheelState.Spawns)
	s.Boss.SetHamsterRelativeOffset(s.WheelState.ham.LogicalAngle)
}

func (s *BossPhase1State) Update() error {
	if res := s.WheelState.Update(); res != nil {
		return res
	}

	return nil
}

func (s *BossPhase1State) Draw(screen *ebiten.Image) {
	s.WheelState.Draw(screen)
}
