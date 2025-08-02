package assets

type AssetAnimation struct {
	TotalFrames   int
	InitialSprite AssetSprite
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
	}
	AnimationSnake = AssetAnimation{
		TotalFrames: 6,
		InitialSprite: AssetSprite{
			AssetKey: AssetKey_Snake_PNG,
			X:        0,
			Y:        0,
			W:        256,
			H:        256,
		},
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
	}
)
