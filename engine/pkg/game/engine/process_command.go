package engine

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
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
			charGameObj.(entity.IMovingObject).SetXYSpeeds(e, cmd.(string))
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
			if npcObj, npcOk := e.GameObjects()[params.(map[string]interface{})["npc_id"].(string)]; npcOk {
				npcObj.(entity.INpcObject).BuyItem(e, charGameObj,
					params.(map[string]interface{})["item_name"].(string),
					params.(map[string]interface{})["amount"].(float64))
			}
		case "get_craft_atlas":
			e.SendResponse("craft_atlas", craft.GetAtlas(), player)
		case "open_container":
			container := e.GameObjects()[params.(string)].(entity.IContainerObject)
			e.SendResponse("container_items", container.GetItems(e), player)
		case "equip_item":
			item := e.GameObjects()[params.(string)].(entity.IEquipableObject)
			item.Equip(e, player)
		case "unequip_item":
			item := e.GameObjects()[params.(string)].(entity.IEquipableObject)
			item.Unequip(e, player)
		case "drop_item":
			item := e.GameObjects()[params.(string)].(entity.IPickableObject)
			item.Drop(e, player)
		case "pickup_item":
			item := e.GameObjects()[params.(string)].(entity.IPickableObject)
			item.Pickup(e, player)
		case "destroy_item":
			item := e.GameObjects()[params.(string)].(entity.IPickableObject)
			item.Destroy(e, player)
		case "destroy_building":
			building := e.GameObjects()[params.(string)].(entity.IBuildingObject)
			building.Destroy(e, player)
		case "destroy_claim":
			claim := e.GameObjects()[params.(string)].(entity.IClaimObeliskObject)
			claim.Destroy(e, player)
		case "put_to_container":
			item := e.GameObjects()[params.(map[string]interface{})["item_id"].(string)].(entity.IPickableObject)
			item.PutToContainer(
				e,
				params.(map[string]interface{})["container_id"].(string),
				int(params.(map[string]interface{})["position"].(float64)),
				player)
		case "apply_effect":
			potion := e.GameObjects()[params.(string)]
			potion.(entity.IPotionObject).ApplyToPlayer(e, player)
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
			cactus := e.GameObjects()[params.(string)]
			if cactus.(entity.ICactusObject).CheckCut(e, charGameObj) {
				delayed_actions.Start(e, charGameObj, "CutCactus", map[string]interface{}{
					"characterId": charGameObj.Id(),
					"cactusId": cactus.Id(),
				}, -1.0)
			}
		case "dig_surface":
			shovel := e.GameObjects()[params.(string)]
			if shovel.(entity.IShovelObject).CheckDig(e, charGameObj) {
				delayed_actions.Start(e, charGameObj, "Dig", map[string]interface{}{
					"characterId": charGameObj.Id(),
					"shovelId": shovel.Id(),
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
			hatchery := e.GameObjects()[params.(string)]
			if hatchery.(entity.IHatcheryObject).CheckHatch(e, charGameObj) {
				delayed_actions.Start(e, hatchery, "Hatch", map[string]interface{}{
					"hatcheryId": hatchery.Id(),
					"mobPath": "mob/fire_dragon",
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
			charGameObj.(entity.ICharacterObject).SelectTarget(e, targetId)
		case "deselect_target":
			charGameObj.(entity.ICharacterObject).DeselectTarget(e)
		case "melee_hit":
			charGameObj.(entity.ICharacterObject).MeleeHit(e)
		case "pay_rent":
			claim := e.GameObjects()[params.(string)].(entity.IClaimObeliskObject)
			claim.ExtendRent(e)
		}
	}
}
