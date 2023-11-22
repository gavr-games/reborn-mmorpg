package entity

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	MaxDistance = 0.1
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
	Effects map[string]interface{}
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

func (obj GameObject) Clone() *GameObject {
	clone := &GameObject{
		X: obj.X,
		Y: obj.Y,
		Width: obj.Width,
		Height: obj.Height,
		Id: obj.Id,
		Type: obj.Type,
		Floor: obj.Floor,
		Rotation: obj.Rotation,
		Properties: make(map[string]interface{}),
		Effects: make(map[string]interface{}),
	}
	clone.Properties = utils.CopyMap(obj.Properties)
	clone.Effects = utils.CopyMap(obj.Effects)
	return clone
}

// Get approximate distance between objects. Assuming all of them are rectangles
func (a GameObject) GetDistance(b *GameObject) float64 {
	aXCenter := a.X + a.Width / 2
	aYCenter := a.Y + a.Height / 2
	
	bXCenter := b.X + b.Width / 2
	bYCenter := b.Y + b.Height / 2

	xDistance := math.Abs(aXCenter - bXCenter) - (a.Width / 2 + b.Width / 2)
	if xDistance < 0 {
		xDistance = 0.0
	}

	yDistance := math.Abs(aYCenter - bYCenter) - (a.Height / 2 + b.Height / 2)
	if yDistance < 0 {
		yDistance = 0.0
	}

	return math.Sqrt(math.Pow(xDistance, 2.0) + math.Pow(yDistance, 2.0))
}

// Determines if 2 objects are close enough to each other
func (a GameObject) IsCloseTo(b *GameObject) bool {
	if (a.Floor != b.Floor) {
		return false
	}
	return a.GetDistance(b) < MaxDistance
}
