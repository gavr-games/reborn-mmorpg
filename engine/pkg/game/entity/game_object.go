package entity

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

type GameObject struct {
	// params for quadtree
	X      float64
	Y      float64
	Width  float64
	Height float64

	// game params
	Id string
	Type string
	Floor int // -1 for does not belong to any floor
	CurrentAction *DelayedAction
	Rotation float64 // from 0 to math.Pi * 2
	Properties map[string]interface{}
}

func (obj GameObject) HitBox() utils.Bounds {
	return utils.Bounds{
		X: obj.X,
		Y: obj.Y,
		Width: obj.Width,
		Height: obj.Height,
	}
}

//IsPoint - Checks if a bounds object is a point or not (has no width or height)
func (obj GameObject) IsPoint() bool {
	if obj.Width == 0 && obj.Height == 0 {
		return true
	}
	return false
}

// Intersects - Checks if a Bounds object intersects with another Bounds
func (a GameObject) Intersects(b utils.Bounds) bool {
	aMaxX := a.X + a.Width
	aMaxY := a.Y + a.Height
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	// a is left of b
	if aMaxX < b.X {
		return false
	}

	// a is right of b
	if a.X > bMaxX {
		return false
	}

	// a is above b
	if aMaxY < b.Y {
		return false
	}

	// a is below b
	if a.Y > bMaxY {
		return false
	}

	// The two overlap
	return true
}
