package models

type SceneAnimations struct {
	animations    map[int]*Animation
	availableKeys []int
	endKey        int
}

func NewSceneAnimations() *SceneAnimations {
	return &SceneAnimations{
		animations:    map[int]*Animation{},
		availableKeys: make([]int, 0),
		endKey:        -1,
	}
}

func (a *SceneAnimations) Update() {
	for _, anim := range a.animations {
		if anim != nil {
			anim.AdvanceFrame()
		}
	}
}

func (a *SceneAnimations) GetAllCurrentSprites() []AnimationFrame {
	sprites := make([]AnimationFrame, 0)
	for _, anim := range a.animations {
		if anim != nil {
			sprites = append(sprites, anim.GetCurrentFrame())
		}
	}

	return sprites
}

func (a *SceneAnimations) getNextId() (id int) {
	if len(a.availableKeys) < 1 {
		a.endKey += 1
		return a.endKey
	}

	pop, others := a.availableKeys[0], a.availableKeys[1:]

	a.availableKeys = others

	return pop
}

func (a *SceneAnimations) AddSceneAnimation(animation *Animation) (id int) {
	id = a.getNextId()

	a.animations[id] = animation

	return id
}

func (a *SceneAnimations) AddOneTimeSceneAnimation(animation *Animation) (id int) {
	id = a.getNextId()

	oldOnComplete := animation.OnComplete
	nuOnComplete := func() {
		if oldOnComplete != nil {
			oldOnComplete()
		}

		a.RemoveAnimation(id)
	}
	animation.OnComplete = nuOnComplete
	a.animations[id] = animation

	return id
}

func (a *SceneAnimations) RemoveAnimation(id int) {
	if _, exists := a.animations[id]; exists {
		a.availableKeys = append(a.availableKeys, id)
		a.animations[id] = nil
	}
}
