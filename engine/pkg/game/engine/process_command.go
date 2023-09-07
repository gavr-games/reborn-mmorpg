package engine

import (
	"math"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

// Process commands from players
func ProcessCommand(e IEngine, characterId int, command map[string]interface{}) {
	if player, ok := e.Players()[characterId]; ok {
		cmd := command["cmd"]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		speed := charGameObj.Properties["speed"].(float64)
		axisSpeed := math.Sqrt(speed * speed / 2)
		switch cmd {
		case "stop":
			charGameObj.Properties["speed_x"] = 0.0
			charGameObj.Properties["speed_y"] = 0.0
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_north":
			charGameObj.Properties["speed_x"] = 0.0
			charGameObj.Properties["speed_y"] = speed
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_south":
			charGameObj.Properties["speed_x"] = 0.0
			charGameObj.Properties["speed_y"] = -speed
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_east":
			charGameObj.Properties["speed_x"] = speed
			charGameObj.Properties["speed_y"] = 0.0
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_west":
			charGameObj.Properties["speed_x"] = -speed
			charGameObj.Properties["speed_y"] = 0.0
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_north_east":
			charGameObj.Properties["speed_x"] = axisSpeed
			charGameObj.Properties["speed_y"] = axisSpeed
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_north_west":
			charGameObj.Properties["speed_x"] = -axisSpeed
			charGameObj.Properties["speed_y"] = axisSpeed
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_south_east":
			charGameObj.Properties["speed_x"] = axisSpeed
			charGameObj.Properties["speed_y"] = -axisSpeed
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "move_south_west":
			charGameObj.Properties["speed_x"] = -axisSpeed
			charGameObj.Properties["speed_y"] = -axisSpeed
			SendGameObjectUpdate(e, charGameObj, "update_object")
		case "get_character_info":
			SendResponse(e, "character_info", game_objects.GetInfo(e.GameObjects(), charGameObj), player)
		}
	}
}
