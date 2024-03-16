package utils

type IBounds interface {
	HitBox() Bounds
	IsPoint() bool
	Intersects(Bounds) bool
}

// Bounds - A bounding box with a x,y origin and width and height
type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (b Bounds) HitBox() Bounds {
	return b
}

//IsPoint - Checks if a bounds object is a point or not (has no width or height)
func (b Bounds) IsPoint() bool {

	if b.Width == 0 && b.Height == 0 {
		return true
	}

	return false

}

// Intersects - Checks if a Bounds object intersects with another Bounds
func (b Bounds) Intersects(a Bounds) bool {

	aMaxX := a.X + a.Width
	aMaxY := a.Y + a.Height
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	// a is left of b
	if aMaxX <= b.X {
		return false
	}

	// a is right of b
	if a.X >= bMaxX {
		return false
	}

	// a is above b
	if aMaxY <= b.Y {
		return false
	}

	// a is below b
	if a.Y >= bMaxY {
		return false
	}

	// The two overlap
	return true
}
