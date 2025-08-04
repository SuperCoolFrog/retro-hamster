package models

import (
	"fmt"
	"math"
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
		}
		return seed
	},
	"m": func(index int) *Spawn {
		wheelRadiusModified := WHEEL_RADIUS
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
		mod := float64(WHEEL_RADIUS / 25)
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
		snake.ModHitBox = .5
		snake.OnCollision = func(ham *Hamster) {
			damage := ham.Momentum.Current - snake.Power
			ham.Momentum.Current -= snake.Power

			if damage < 0 {
				ham.Health -= 1
			}

			snake.IsAlive = false
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
		shark.ModHitBox = .5
		shark.OnCollision = func(ham *Hamster) {
			damage := ham.Momentum.Current - shark.Power
			ham.Momentum.Current -= shark.Power

			if damage < 0 {
				ham.Health -= 1
			}

			shark.IsAlive = false
		}
		return shark
	},
	"M": func(index int) *Spawn {
		mod := float64(assets.AnimationShark.InitialSprite.H / 5)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		hedgehog := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          8,
			CurrentFrame: 0,
			Details:      assets.AnimationHedgeHog,
			X:            0,
			Y:            0,
		})
		hedgehog.ModHitBox = .5
		hedgehog.Direction = DIRECTION_LEFT
		hedgehog.Power = 50
		hedgehog.OnCollision = func(ham *Hamster) {
			damage := ham.Momentum.Current - hedgehog.Power
			ham.Momentum.Current -= hedgehog.Power

			if damage < 0 {
				ham.Health -= 1
			}

			hedgehog.IsAlive = false
		}
		return hedgehog
	},
	"B": func(index int) *Spawn {
		mod := float64(assets.AnimationBossPhase1.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		boss := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationBossPhase1,
		})
		boss.Power = 100
		boss.IsAlive = true
		boss.LastActivation = time.Now().Add(-time.Minute)
		boss.ActivationCoolDown = time.Second
		boss.OnCollision = func(ham *Hamster) {
			if time.Since(boss.LastActivation) > boss.ActivationCoolDown {
				damage := ham.Momentum.Current - boss.Power
				ham.Momentum.Current -= boss.Power

				if damage < 0 {
					ham.Health -= 1
				} else {
					boss.Health -= 1
					BOSS_HEALTH = boss.Health
					fmt.Printf("Boss H %f\n", boss.Health)
				}

				if boss.StartingHealth-boss.Health >= 2 {
					boss.IsAlive = false
				}

				boss.LastActivation = time.Now()
			}

			if ham.X < boss.X {
				ham.Blocked = DIRECTION_RIGHT
			} else {
				ham.Blocked = DIRECTION_LEFT
			}

			ham.Momentum.Current = 0
		}

		boss.Init = func() {
			BOSS_HAS_INIT = true
			boss.StartingHealth = BOSS_HEALTH
			boss.Health = BOSS_HEALTH
		}

		boss.SkewAngle = math.Pi / 12
		return boss
	},
	"2": func(index int) *Spawn {
		mod := float64(assets.AnimationBossPhase1.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		boss := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationBossPhase2,
		})
		boss.Power = 100
		boss.IsAlive = true
		boss.LastActivation = time.Now().Add(-time.Minute)
		boss.ActivationCoolDown = time.Second
		boss.OnCollision = func(ham *Hamster) {
			if time.Since(boss.LastActivation) > boss.ActivationCoolDown {
				damage := ham.Momentum.Current - boss.Power
				ham.Momentum.Current -= boss.Power

				if damage < 0 {
					ham.Health -= 1
				} else {
					boss.Health -= 1
					BOSS_HEALTH = boss.Health
					fmt.Printf("Boss H %f\n", boss.Health)
				}

				if boss.StartingHealth-boss.Health >= 2 {
					boss.IsAlive = false
				}

				boss.LastActivation = time.Now()
			}

			if ham.X < boss.X {
				ham.Blocked = DIRECTION_RIGHT
			} else {
				ham.Blocked = DIRECTION_LEFT
			}

			ham.Momentum.Current = 0
		}

		boss.Init = func() {
			BOSS_HAS_INIT = true
			boss.StartingHealth = BOSS_HEALTH
			boss.Health = BOSS_HEALTH
		}

		boss.SkewAngle = math.Pi / 12
		return boss
	},
	"3": func(index int) *Spawn {
		mod := float64(assets.AnimationBossPhase3.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		boss := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          0,
			CurrentFrame: 0,
			Details:      assets.AnimationBossPhase3,
		})
		boss.Power = 100
		boss.IsAlive = true
		boss.OnCollision = func(ham *Hamster) {

			if ham.X < boss.X {
				ham.Blocked = DIRECTION_RIGHT
			} else {
				ham.Blocked = DIRECTION_LEFT
			}

			ham.Momentum.Current = 0
		}

		boss.Init = func() {
			BOSS_HAS_INIT = true
			boss.StartingHealth = BOSS_HEALTH
			boss.Health = BOSS_HEALTH
			boss.LastActivation = time.Now()
			boss.ActivationCoolDown = time.Second * 10
		}

		boss.OtherUpdate = func() {
			if time.Since(boss.LastActivation) > boss.ActivationCoolDown {
				BOSS_HEALTH -= 2
				fmt.Printf("Boss H %f\n", boss.Health)

				boss.IsAlive = false

				boss.LastActivation = time.Now()
			}
		}

		boss.SkewAngle = math.Pi / 12
		return boss
	},
}
