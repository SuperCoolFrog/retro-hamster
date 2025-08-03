package assets

type AnimationDirection int8

const (
	ANIMATION_DIRECTION_NONE  = 0
	ANIMATION_DIRECTION_LEFT  = -1
	ANIMATION_DIRECTION_RIGHT = 1
)

type AssetAnimation struct {
	TotalFrames   int
	InitialSprite AssetSprite
	Direction     AnimationDirection
}

var (
	AnimationHamsterRun = AssetAnimation{
		TotalFrames: 4,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Hamster_Run_PNG,
			X:        0,
			Y:        0,
			W:        256,
			H:        256,
		},
		Direction: ANIMATION_DIRECTION_RIGHT,
	}
	AnimationSnake = AssetAnimation{
		TotalFrames: 4,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Snake_PNG,
			X:        0,
			Y:        0,
			W:        256,
			H:        256,
		},
		Direction: ANIMATION_DIRECTION_LEFT,
	}
	AnimationShark = AssetAnimation{
		TotalFrames: 4,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Shark_PNG,
			X:        0,
			Y:        0,
			W:        256,
			H:        256,
		},
		Direction: ANIMATION_DIRECTION_LEFT,
	}
	AnimationHedgeHog = AssetAnimation{
		TotalFrames: 1,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_HedgeHog_PNG,
			X:        0,
			Y:        0,
			W:        256,
			H:        256,
		},
		Direction: ANIMATION_DIRECTION_RIGHT,
	}
	AnimationSeed = AssetAnimation{
		TotalFrames: 1,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Static_PNG,
			X:        Sprite_Seed.X,
			Y:        Sprite_Seed.Y,
			W:        Sprite_Seed.W,
			H:        Sprite_Seed.H,
		},
		Direction: ANIMATION_DIRECTION_NONE,
	}
	AnimationBlock = AssetAnimation{
		TotalFrames: 1,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Static_PNG,
			X:        Sprite_Block.X,
			Y:        Sprite_Block.Y,
			W:        Sprite_Block.W,
			H:        Sprite_Block.H,
		},
		Direction: ANIMATION_DIRECTION_NONE,
	}
	AnimationFence = AssetAnimation{
		TotalFrames: 1,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Static_PNG,
			X:        Sprite_Fence.X,
			Y:        Sprite_Fence.Y,
			W:        Sprite_Fence.W,
			H:        Sprite_Fence.H,
		},
		Direction: ANIMATION_DIRECTION_LEFT,
	}
)
