package models

import (
	"retro-hamster/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParallaxDetails struct {
	PerTick    float64
	X, Y, W, H float64
	Sprite     assets.AssetSprite
	p2         *ParallaxDetails
	Source     GameAssetImg
}

type Parallaxer struct {
	Details []*ParallaxDetails
}

func NewParallaxer() *Parallaxer {
	return &Parallaxer{
		Details: make([]*ParallaxDetails, 0),
	}
}

func (p *Parallaxer) AddDetails(PerTick, X, Y, W, H float64, Sprite assets.AssetSprite, Source GameAssetImg) {
	pd := &ParallaxDetails{
		PerTick, X, Y, W, H, Sprite, nil, Source,
	}
	pd2 := &ParallaxDetails{
		PerTick, X + W, Y, W, H, Sprite, nil, Source,
	}

	pd.p2 = pd2
	p.Details = append(p.Details, pd)
}

func (p *Parallaxer) Update(direction DIRECTION) {
	for _, d := range p.Details {
		d.X += (d.PerTick * float64(direction))
		d.p2.X += (d.PerTick * float64(direction))

		if d.X < -d.W && d.p2.X > d.X {
			d.X = d.W + d.p2.X
		}

		if d.X > d.W && d.p2.X > 0 {
			d.X = d.p2.X - d.W
		}

		if d.p2.X < -d.W && d.X > d.p2.X {
			d.p2.X = d.W + d.X
		}

		if d.p2.X > d.W && d.X > 0 {
			d.p2.X = d.X - d.W
		}
	}
}

func (p *Parallaxer) Draw(screen *ebiten.Image) {
	for _, d := range p.Details {
		DrawAssetSprite(d.Source.Image, screen, d.X, d.Y, d.Sprite)
		DrawAssetSprite(d.Source.Image, screen, d.p2.X, d.p2.Y, d.Sprite)
	}
}
