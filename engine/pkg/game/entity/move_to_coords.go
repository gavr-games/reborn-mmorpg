package entity

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	MoveToExactCoords = 1
	MoveCloseToBounds = 2
)

// Used in GameObject to represent coords, where to go.
// This is used when user clicks character to go somewhere with mouse or craft something in distance.
type MoveToCoords struct {
	Mode int
	Bounds utils.Bounds
	DirectionChangeTime float64 // how often to change movement direction
	TimeUntilDirectionChange float64 // how much time left until next direction change
	Callback func() // execute some code, when moveTo iis finished
}
