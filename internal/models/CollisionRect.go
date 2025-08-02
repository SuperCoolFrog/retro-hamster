package models

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CollisionRect struct {
	X, Y float64
	W, H float64
}

func (r CollisionRect) Intersects(other CollisionRect) bool {
	return r.X < other.X+other.W &&
		r.X+r.W > other.X &&
		r.Y < other.Y+other.H &&
		r.Y+r.H > other.Y
}

func (r CollisionRect) toVector4() [4]Vector {
	corners := [4]Vector{
		{r.X, r.Y},             // top-left
		{r.X + r.W, r.Y},       // top-right
		{r.X, r.Y + r.H},       // bottom-right
		{r.X + r.W, r.Y + r.H}, // bottom-left
	}

	return corners
}

func (r CollisionRect) IntersectsPolygon(other [4]Vector) bool {
	a := r.toVector4()
	b := other

	axes := getAxes(a)
	axes = append(axes, getAxes(b)...)

	for _, axis := range axes {
		minA, maxA := projectPolygon(axis, a)
		minB, maxB := projectPolygon(axis, b)

		if maxA < minB || maxB < minA {
			return false // gap found — no collision
		}
	}

	return true // no separating axis — they intersect
}

func DrawCollisionRect(screen *ebiten.Image, r CollisionRect, col color.Color) {
	screen.Set(int(r.X), int(r.Y), col)
	ebitenutil.DrawRect(screen, r.X, r.Y, r.W, r.H, col)
}

// Get perpendicular axes from polygon edges
func getAxes(poly [4]Vector) []Vector {
	var axes []Vector
	for i := 0; i < 4; i++ {
		p1 := poly[i]
		p2 := poly[(i+1)%4]

		edge := Vector{X: p2.X - p1.X, Y: p2.Y - p1.Y}
		normal := Vector{X: -edge.Y, Y: edge.X} // perpendicular
		length := math.Hypot(normal.X, normal.Y)
		if length > 0 {
			normal.X /= length
			normal.Y /= length
		}
		axes = append(axes, normal)
	}
	return axes
}

// Project polygon onto an axis
func projectPolygon(axis Vector, poly [4]Vector) (min, max float64) {
	min = dot(axis, poly[0])
	max = min
	for i := 1; i < 4; i++ {
		proj := dot(axis, poly[i])
		if proj < min {
			min = proj
		}
		if proj > max {
			max = proj
		}
	}
	return
}

func dot(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}
