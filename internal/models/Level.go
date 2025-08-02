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
		angle := float64(index) * SPAWN_SPACING / (wheelRadiusModified)
		// angle -= math.Pi / 2 /* THis will translate to top as starting point */
		// angle := float64(i) * 2 * math.Pi / 5
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
}
