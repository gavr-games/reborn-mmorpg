package moving_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

const (
	ExactCoordsDistance = 0.15
)

// Perform object moving to the selected coords
func (mObj *MovingObject) PerformMoveTo(e entity.IEngine, tickDelta int64) {
	obj := mObj.gameObj
	moveTo := obj.MoveToCoords()
	if moveTo != nil {
		if mObj.needToStop() {
			mObj.Stop(e)
		}
		moveTo.TimeUntilDirectionChange = moveTo.TimeUntilDirectionChange - float64(tickDelta)
		if (moveTo.TimeUntilDirectionChange <= 0) {
			moveTo.TimeUntilDirectionChange = moveTo.DirectionChangeTime
			mObj.SetXYSpeeds(e, obj.GetDirectionToXY(moveTo.Bounds.X, moveTo.Bounds.Y))
		}
	}
}

func (mObj *MovingObject) needToStop() bool {
	obj := mObj.gameObj
	moveTo := obj.MoveToCoords()
	// is close to exact coords
	if (moveTo.Mode == entity.MoveToExactCoords && obj.GetDistanceToXY(moveTo.Bounds.X, moveTo.Bounds.Y) < ExactCoordsDistance) {
		return true
	}
	// is close to bounds
	if (moveTo.Mode == entity.MoveCloseToBounds) {
		tempGameObj := &entity.GameObject{}
		tempGameObj.SetProperties(make(map[string]interface{}))
		tempGameObj.SetX(moveTo.Bounds.X)
		tempGameObj.SetY(moveTo.Bounds.Y)
		tempGameObj.SetWidth(moveTo.Bounds.Width)
		tempGameObj.SetHeight(moveTo.Bounds.Height)
		return obj.IsCloseTo(tempGameObj)
	}
	return false
}
