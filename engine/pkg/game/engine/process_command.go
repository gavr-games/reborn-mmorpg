package engine

import (
	"math"
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/trees"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/rocks"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/items"
)

// Process commands from players
func ProcessCommand(e entity.IEngine, characterId int, command map[string]interface{}) {
	if player, ok := e.Players()[characterId]; ok {
		cmd := command["cmd"]
		params := command["params"]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]
		speed := charGameObj.Properties["speed"].(float64)
		axisSpeed := math.Sqrt(speed * speed / 2)

		// List of commands, which don't interrupt current character action.
		// Like get_character_info does not interrupt choping a tree, but any movement does
		nonCancellingCmds := []string{"get_character_info", "open_container"}
		// Cancel character delayed actions
		if !slices.Contains(nonCancellingCmds, cmd.(string)) {
			delayed_actions.Cancel(e, charGameObj)
		}
 
		// Process Cmd
		switch cmd {
		case "stop":
			charGameObj.Properties["speed_x"] = 0.0
			charGameObj.Properties["speed_y"] = 0.0
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_north":
			charGameObj.Properties["speed_x"] = 0.0
			charGameObj.Properties["speed_y"] = speed
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_south":
			charGameObj.Properties["speed_x"] = 0.0
			charGameObj.Properties["speed_y"] = -speed
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_east":
			charGameObj.Properties["speed_x"] = speed
			charGameObj.Properties["speed_y"] = 0.0
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_west":
			charGameObj.Properties["speed_x"] = -speed
			charGameObj.Properties["speed_y"] = 0.0
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_north_east":
			charGameObj.Properties["speed_x"] = axisSpeed
			charGameObj.Properties["speed_y"] = axisSpeed
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_north_west":
			charGameObj.Properties["speed_x"] = -axisSpeed
			charGameObj.Properties["speed_y"] = axisSpeed
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_south_east":
			charGameObj.Properties["speed_x"] = axisSpeed
			charGameObj.Properties["speed_y"] = -axisSpeed
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_south_west":
			charGameObj.Properties["speed_x"] = -axisSpeed
			charGameObj.Properties["speed_y"] = -axisSpeed
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "get_character_info":
			e.SendResponse("character_info", game_objects.GetInfo(e.GameObjects(), charGameObj), player)
		case "open_container":
			e.SendResponse("container_items", containers.GetItems(e, params.(string)), player)
		case "equip_item":
			items.Equip(e, params.(string), player)
		case "unequip_item":
			items.Unequip(e, params.(string), player)
		case "drop_item":
			items.Drop(e, params.(string), player)
		case "pickup_item":
			items.Pickup(e, params.(string), player)
		case "chop_tree":
			treeId := params.(string)
			if trees.CheckChop(e, player, treeId) {
				delayed_actions.Start(e, charGameObj, "Chop", map[string]interface{}{
					"playerId": float64(player.Id), // this conversion is required, because json unmarshal decodes all numbers to float64
					"treeId": treeId,
				})
			}
		case "chip_rock":
			stoneId := params.(string)
			if rocks.CheckChip(e, player, stoneId) {
				delayed_actions.Start(e, charGameObj, "Chip", map[string]interface{}{
					"playerId": float64(player.Id), // this conversion is required, because json unmarshal decodes all numbers to float64
					"rockId": stoneId,
				})
			}
		}
	}
}
