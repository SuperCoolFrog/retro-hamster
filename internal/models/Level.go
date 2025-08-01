package models

import (
	"fmt"
	"retro-hamster/assets"
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
	"S": func(index int) *Spawn {
		mod := float64(assets.AnimationSnake.InitialSprite.W / 4)
		wheelRadiusModified := WHEEL_RADIUS + mod
		angle := float64(index+1) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		snake := NewSpawn(angle, wheelRadiusModified, &Animation{
			FPS:          12,
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
}
