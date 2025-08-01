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
)
