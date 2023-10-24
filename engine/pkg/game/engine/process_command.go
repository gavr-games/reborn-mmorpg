package engine

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/trees"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/rocks"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/hatcheries"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/targets"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/items"
)

// Process commands from players
func ProcessCommand(e entity.IEngine, characterId int, command map[string]interface{}) {
	if player, ok := e.Players()[characterId]; ok {
		cmd := command["cmd"]
		params := command["params"]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]

		// List of commands, which don't interrupt current character action.
		// Like get_character_info does not interrupt choping a tree, but any movement does
		nonCancellingCmds := []string{"get_character_info", "open_container", "get_craft_atlas"}
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
		case "move_north", "move_south", "move_east", "move_west",
				"move_north_east", "move_north_west", "move_south_east", "move_south_west":
			game_objects.SetXYSpeeds(charGameObj, cmd.(string))
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "get_character_info":
			e.SendResponse("character_info", game_objects.GetInfo(e.GameObjects(), charGameObj), player)
		case "get_craft_atlas":
			e.SendResponse("craft_atlas", craft.GetAtlas(), player)
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
		case "destroy_item":
			items.Destroy(e, params.(string), player)
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
		case "craft":
			if craft.Check(e, player, params.(map[string]interface{})) {
				params.(map[string]interface{})["playerId"] = float64(player.Id)
				delayed_actions.StartCraft(e, charGameObj, "Craft", params.(map[string]interface{}))
			}
		case "hatch_fire_dragon":
			hatcheryId := params.(string)
			if hatcheries.CheckHatch(e, player, hatcheryId) {
				delayed_actions.Start(e, e.GameObjects()[hatcheryId], "HatchFireDragon", map[string]interface{}{
					"hatcheryId": hatcheryId,
				})
			}
		case "follow":
			mobId := params.(string)
			mob, ok := e.Mobs()[mobId]
			if ok {
				//TODO: Check commands  can be executed only close enough to the mob
				mob.Follow(charGameObj.Id)
			}
		case "unfollow":
			mobId := params.(string)
			mob, ok := e.Mobs()[mobId]
			if ok {
				//TODO: Check commands  can be executed only close enough to the mob
				mob.Unfollow()
			}
		case "select_target":
			targetId := params.(string)
			targets.Select(e, charGameObj, targetId)
		case "deselect_target":
			targets.Deselect(e, charGameObj)
		}
	}
}
