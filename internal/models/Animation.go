package models

import "retro-hamster/assets"

const normalFPS = 60

type Animation struct {
	FPS          int
	CurrentFrame int
	Details      assets.AssetAnimation
	OnComplete   func()
	counter      int
	X            float64
	Y            float64
}

type AnimationFrame struct {
	assets.AssetSprite
	TargetX float64
	TargetY float64
}

func (a *Animation) GetCurrentFrame() AnimationFrame {
	return AnimationFrame{
		AssetSprite: assets.AssetSprite{
			AssetKey: a.Details.InitialSprite.AssetKey,
			X:        a.Details.InitialSprite.X + a.Details.InitialSprite.W*a.CurrentFrame,
			Y:        a.Details.InitialSprite.Y,
			W:        a.Details.InitialSprite.W,
			H:        a.Details.InitialSprite.H,
		},
		TargetX: a.X,
		TargetY: a.Y,
	}
}

func (a *Animation) getSkipCount() int {
	return normalFPS / a.FPS
}

func (a *Animation) AdvanceFrame() {

	if a.FPS == 0 {
		return
	}

	a.counter++
	if a.counter < a.getSkipCount() {
		return
	}

	a.counter = 0

	if a.CurrentFrame+1 >= a.Details.TotalFrames {
		a.CurrentFrame = 0
		if a.OnComplete != nil {
			a.OnComplete()
		}
	} else {
		a.CurrentFrame++
	}
}
