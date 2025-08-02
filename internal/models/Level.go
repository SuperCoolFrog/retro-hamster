package models

type Level struct {
	Rount           int
	Spawns          []*Spawn
	Animations      *SceneAnimations
	OnLevelComplete func()
}
