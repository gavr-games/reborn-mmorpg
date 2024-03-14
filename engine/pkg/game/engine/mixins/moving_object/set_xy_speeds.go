package moving_object

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)

// Set x and y speeds depending on the direction
func (mObj *MovingObject) SetXYSpeeds(e entity.IEngine, direction string) {
	obj := mObj.gameObj
	speed := obj.GetProperty("speed").(float64)
	axisSpeed := math.Sqrt(speed * speed / 2)

	validDirection := false
	for _, dir := range constants.GetPossibleDirections() {
		if dir == direction {
			validDirection = true
			break
		}
	}

	if !validDirection {
		//TODO: log error
		return
	}

	switch direction {
		case "move_north":
			obj.SetProperty("speed_x", 0.0)
			obj.SetProperty("speed_y", speed)
		case "move_south":
			obj.SetProperty("speed_x", 0.0)
			obj.SetProperty("speed_y", -speed)
		case "move_east":
			obj.SetProperty("speed_x", speed)
			obj.SetProperty("speed_y", 0.0)
		case "move_west":
			obj.SetProperty("speed_x", -speed)
			obj.SetProperty("speed_y", 0.0)
		case "move_north_east":
			obj.SetProperty("speed_x", axisSpeed)
			obj.SetProperty("speed_y", axisSpeed)
		case "move_north_west":
			obj.SetProperty("speed_x", -axisSpeed)
			obj.SetProperty("speed_y", axisSpeed)
		case "move_south_east":
			obj.SetProperty("speed_x", axisSpeed)
			obj.SetProperty("speed_y", -axisSpeed)
		case "move_south_west":
			obj.SetProperty("speed_x", -axisSpeed)
			obj.SetProperty("speed_y", -axisSpeed)
	}

	// this is required to determine where character/mob looks, when stopped
	// it is important to later hit the target with weapon
	obj.SetRotationByDirection(direction)

	e.SendGameObjectUpdate(obj, "update_object")
}
