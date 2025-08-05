package models

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fader struct {
	IsFading       bool
	OnFadeComplete func()
	w, h           float64
	on             bool
	fadeAlpha      float64
	fadingOut      bool
	fadeSpeed      float64
}

func NewFader(out bool, w, h float64) *Fader {
	fader := &Fader{
		w:              w,
		h:              h,
		fadeSpeed:      0.05,
		IsFading:       false,
		on:             false,
		fadingOut:      out,
		OnFadeComplete: func() {},
	}

	if out {
		fader.fadeAlpha = 0.0
	} else {
		fader.fadeAlpha = 1.0
	}

	return fader
}

func (f *Fader) Start() {
	f.IsFading = true
	f.on = true
}

func (f *Fader) Stop() {
	f.on = false
	f.IsFading = false
}

func (f *Fader) Update() {
	if !f.on {
		return
	}
	if f.fadingOut {
		f.fadeAlpha += f.fadeSpeed
		if f.fadeAlpha >= 1.0 {
			f.IsFading = false
			f.on = false
			f.OnFadeComplete()
		}
	} else {
		f.fadeAlpha -= f.fadeSpeed
		if f.fadeAlpha <= 0.0 {
			f.IsFading = false
			f.on = false
			f.OnFadeComplete()
		}
	}
}

func (f *Fader) Draw(screen *ebiten.Image) {
	if !f.on {
		return
	}
	// Overlay fade rectangle
	// if f.fadeAlpha > 0 {
	// Create a black overlay
	overlay := ebiten.NewImage(int(f.w), int(f.h))
	overlay.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(float32(math.Min(f.fadeAlpha, 1.0)))
	screen.DrawImage(overlay, op)
	// }
}
