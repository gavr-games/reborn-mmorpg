package engine

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/plants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/hatcheries"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/targets"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/items"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/effects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/npcs"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
)

// Process commands from players
func ProcessCommand(e entity.IEngine, characterId int, command map[string]interface{}) {
	if player, ok := e.Players()[characterId]; ok {
		cmd := command["cmd"]
		params := command["params"]
		charGameObj := e.GameObjects()[player.CharacterGameObjectId]

		// List of commands, which don't interrupt current character action.
		// Like get_character_info does not interrupt choping a tree, but any movement does
		nonCancellingCmds := []string{"get_character_info", "open_container", "get_craft_atlas", "npc_trade_info", "get_item_info"}
		// Cancel character delayed actions
		if !slices.Contains(nonCancellingCmds, cmd.(string)) {
			delayed_actions.Cancel(e, charGameObj)
		}

		// Process Cmd
		switch cmd {
		case "stop":
			charGameObj.Properties()["speed_x"] = 0.0
			charGameObj.Properties()["speed_y"] = 0.0
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_north", "move_south", "move_east", "move_west",
				"move_north_east", "move_north_west", "move_south_east", "move_south_west":
			game_objects.SetXYSpeeds(charGameObj, cmd.(string))
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "get_character_info":
			e.SendResponse("character_info", serializers.GetInfo(e.GameObjects(), charGameObj), player)
		case "get_item_info":
			itemId := params.(string)
			e.SendResponse("item_info", serializers.GetInfo(e.GameObjects(), e.GameObjects()[itemId]), player)
		case "npc_trade_info":
			if npcObj, npcOk := e.GameObjects()[params.(string)]; npcOk {
				e.SendResponse("npc_trade_info", serializers.GetInfo(e.GameObjects(), npcObj), player)
			}
		case "npc_buy_item":
			npcs.BuyItem(e, charGameObj,
				params.(map[string]interface{})["npc_id"].(string),
				params.(map[string]interface{})["item_name"].(string),
				params.(map[string]interface{})["amount"].(float64))
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
		case "put_to_container":
			items.PutToContainer(
				e,
				params.(map[string]interface{})["container_id"].(string),
				int(params.(map[string]interface{})["position"].(float64)),
				params.(map[string]interface{})["item_id"].(string),
				player)
		case "apply_effect":
			effects.ApplyPlayer(e, params.(string), player)
		case "chop_tree":
			tree := e.GameObjects()[params.(string)]
			if tree.(entity.ITreeObject).CheckChop(e, charGameObj) {
				delayed_actions.Start(e, charGameObj, "Chop", map[string]interface{}{
					"characterId": charGameObj.Id(),
					"treeId": tree.Id(),
				}, -1.0)
			}
		case "chip_rock":
			rock := e.GameObjects()[params.(string)]
			if rock.(entity.IRockObject).CheckChip(e, charGameObj) {
				delayed_actions.Start(e, charGameObj, "Chip", map[string]interface{}{
					"characterId": charGameObj.Id(),
					"rockId": rock.Id(),
				}, -1.0)
			}
		case "cut_cactus":
			cactusId := params.(string)
			if plants.CheckCutCactus(e, player, cactusId) {
				delayed_actions.Start(e, charGameObj, "CutCactus", map[string]interface{}{
					"playerId": float64(player.Id), // this conversion is required, because json unmarshal decodes all numbers to float64
					"cactusId": cactusId,
				}, -1.0)
			}
		case "craft":
			if craft.Check(e, player, params.(map[string]interface{})) {
				params.(map[string]interface{})["playerId"] = float64(player.Id)
				craftItem := params.(map[string]interface{})["item_name"].(string)
				delayed_actions.Start(
					e, charGameObj, "Craft",
					params.(map[string]interface{}),
					craft.GetAtlas()[craftItem].(map[string]interface{})["duration"].(float64))
			}
		case "hatch_fire_dragon":
			hatcheryId := params.(string)
			if hatcheries.CheckHatch(e, player, hatcheryId) {
				delayed_actions.Start(e, e.GameObjects()[hatcheryId], "HatchFireDragon", map[string]interface{}{
					"hatcheryId": hatcheryId,
				}, -1.0)
			}
		case "town_teleport":
			delayed_actions.Start(e, charGameObj, "TownTeleport", map[string]interface{}{
				"playerId": float64(player.Id),
			}, -1.0)
		case "claim_teleport":
			delayed_actions.Start(e, charGameObj, "ClaimTeleport", map[string]interface{}{
				"playerId": float64(player.Id),
			}, -1.0)
		case "follow":
			mobId := params.(string)
			mob, ok := e.Mobs()[mobId]
			if ok {
				//TODO: Check commands  can be executed only close enough to the mob
				mob.Follow(charGameObj.Id())
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
		case "melee_hit":
			characters.MeleeHit(e, charGameObj, player)
		case "pay_rent":
			obeliskId := params.(string)
			claims.ExtendRent(e, obeliskId)
		}
	}
}
