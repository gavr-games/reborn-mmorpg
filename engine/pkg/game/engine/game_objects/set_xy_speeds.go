package game_objects

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Order is simportant here for finding direction by angle in mobs/mob.go
var PossibleDirections = [...]string {"move_east", "move_north_east", "move_north", "move_north_west", "move_west", "move_south_west", "move_south", "move_south_east"}

// Set x and y speeds depending on the direction
func SetXYSpeeds(obj *entity.GameObject, direction string) {
	speed := obj.Properties["speed"].(float64)
	axisSpeed := math.Sqrt(speed * speed / 2)

	validDirection := false
	for _, dir := range PossibleDirections {
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
			obj.Properties["speed_x"] = 0.0
			obj.Properties["speed_y"] = speed
			// this is required to determine where character/mob looks, when stopped
			// it is important to later hit the target with weapon
			obj.Rotation = math.Pi / 2
		case "move_south":
			obj.Properties["speed_x"] = 0.0
			obj.Properties["speed_y"] = -speed
			obj.Rotation = math.Pi * 3 / 2
		case "move_east":
			obj.Properties["speed_x"] = speed
			obj.Properties["speed_y"] = 0.0
			obj.Rotation = 0
		case "move_west":
			obj.Properties["speed_x"] = -speed
			obj.Properties["speed_y"] = 0.0
			obj.Rotation = math.Pi
		case "move_north_east":
			obj.Properties["speed_x"] = axisSpeed
			obj.Properties["speed_y"] = axisSpeed
			obj.Rotation = math.Pi / 4
		case "move_north_west":
			obj.Properties["speed_x"] = -axisSpeed
			obj.Properties["speed_y"] = axisSpeed
			obj.Rotation = math.Pi * 3 / 4
		case "move_south_east":
			obj.Properties["speed_x"] = axisSpeed
			obj.Properties["speed_y"] = -axisSpeed
			obj.Rotation = math.Pi * 7 / 4
		case "move_south_west":
			obj.Properties["speed_x"] = -axisSpeed
			obj.Properties["speed_y"] = -axisSpeed
			obj.Rotation = math.Pi * 5 / 4
	}
}

// Set rotation depending on the direction
func SetRotation(obj *entity.GameObject, direction string) {
	validDirection := false
	for _, dir := range PossibleDirections {
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
			obj.Rotation = math.Pi / 2
		case "move_south":
			obj.Rotation = math.Pi * 3 / 2
		case "move_east":
			obj.Rotation = 0
		case "move_west":
			obj.Rotation = math.Pi
		case "move_north_east":
			obj.Rotation = math.Pi / 4
		case "move_north_west":
			obj.Rotation = math.Pi * 3 / 4
		case "move_south_east":
			obj.Rotation = math.Pi * 7 / 4
		case "move_south_west":
			obj.Rotation = math.Pi * 5 / 4
	}
}