package models

import (
	"fmt"
	"retro-hamster/assets"
	"time"
)

type Level struct {
	Rounds          map[int][]*Spawn
	Animations      *SceneAnimations
	OnLevelComplete func()
}

type LevelSpawnConstructor = func(index int) *Spawn

var SymbolToSpawnMap = map[string]LevelSpawnConstructor{
	"o": func(index int) *Spawn {
		mod := float64(assets.AnimationSeed.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS - mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		seed := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationSeed,
		})
		seed.Power = 1
		seed.OnCollision = func(ham *Hamster) {
			ham.XP.Current += seed.Power
			seed.IsAlive = false
			fmt.Printf("XP %d\n", ham.XP.Current)
		}
		return seed
	},
	"m": func(index int) *Spawn {
		mod := float64(assets.AnimationBlock.InitialSprite.W / 5)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		blockSpawn := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationBlock,
		})
		blockSpawn.Power = 1
		blockSpawn.IsObstacle = true
		blockSpawn.OnCollision = func(ham *Hamster) {
			if ham.X < blockSpawn.X {
				ham.Blocked = DIRECTION_RIGHT
			} else {
				ham.Blocked = DIRECTION_LEFT
			}

			ham.Momentum.Current = 0
		}
		return blockSpawn
	},
	"|": func(index int) *Spawn {
		mod := float64(assets.AnimationFence.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		fence := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationFence,
		})
		fence.Power = 1
		fence.IsObstacle = true
		fence.LastActivation = time.Now().Add(-time.Minute)
		fence.ActivationCoolDown = time.Second
		fence.OnCollision = func(ham *Hamster) {
			if ham.X < fence.X {
				ham.Blocked = DIRECTION_RIGHT
			} else {
				ham.Blocked = DIRECTION_LEFT
			}

			ham.Momentum.Current = 0

			if time.Since(fence.LastActivation) > fence.ActivationCoolDown {
				ham.Health -= int(fence.Power)
				fence.LastActivation = time.Now()
			}
		}
		return fence
	},

	/* ENEMIES */
	"S": func(index int) *Spawn {
		mod := float64(assets.AnimationSnake.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		snake := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          8,
			CurrentFrame: 0,
			Details:      assets.AnimationSnake,
			X:            0,
			Y:            0,
		})
		snake.Direction = DIRECTION_LEFT
		snake.Power = 50
		snake.OnCollision = func(ham *Hamster) {
			damage := ham.Momentum.Current - snake.Power
			ham.Momentum.Current -= snake.Power

			if damage < 0 {
				ham.Health -= 1
			}

			snake.IsAlive = false

			fmt.Printf("M %d\n", ham.Momentum.Current)
		}
		return snake
	},
	"Q": func(index int) *Spawn {
		mod := float64(assets.AnimationShark.InitialSprite.H)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		shark := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          8,
			CurrentFrame: 0,
			Details:      assets.AnimationShark,
			X:            0,
			Y:            0,
		})
		shark.Direction = DIRECTION_LEFT
		shark.Power = 50
		shark.OnCollision = func(ham *Hamster) {
			damage := ham.Momentum.Current - shark.Power
			ham.Momentum.Current -= shark.Power

			if damage < 0 {
				ham.Health -= 1
			}

			shark.IsAlive = false

			fmt.Printf("M %d\n", ham.Momentum.Current)
		}
		return shark
	},
}
