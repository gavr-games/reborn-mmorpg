package engine

import (
	"slices"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/delayed_actions"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/dungeons"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/gm"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Process commands from players
// TODO: move commands processing to funcs
// TODO: add incoming params validation
// TODO: reorder alphabetically
func ProcessCommand(e entity.IEngine, characterId int, command map[string]interface{}) bool {
	if player, ok := e.Players().Load(characterId); ok {
		var (
			charGameObj entity.IGameObject
			charOk bool
		)
		cmd := command["cmd"]
		params := command["params"]
		if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
			return false
		}

		// List of commands, which don't interrupt current character action.
		// Like get_character_info does not interrupt choping a tree, but any movement does
		nonCancellingCmds := []string{"get_ping", "get_character_info", "get_dragons_info", "get_npc_dungeons_info", "open_container", "get_craft_atlas", "get_npc_trade_info", "get_item_info"}
		// Cancel character delayed actions and auto moving
		if !slices.Contains(nonCancellingCmds, cmd.(string)) {
			delayed_actions.Cancel(e, charGameObj)
			charGameObj.SetMoveToCoords(nil)
		}

		// Process Cmd
		switch cmd {
		case "get_ping":
			e.SendResponse("ping_info",map[string]interface{}{
				"start_time":  params.(map[string]interface{})["start_time"],
				"server_time": utils.MakeTimestamp(),
			}, player)
		case "stop":
			charGameObj.SetProperty("speed_x", 0.0)
			charGameObj.SetProperty("speed_y", 0.0)
			e.SendGameObjectUpdate(charGameObj, "update_object")
		case "move_north", "move_south", "move_east", "move_west",
			"move_north_east", "move_north_west", "move_south_east", "move_south_west":
			charGameObj.(entity.IMovingObject).SetXYSpeeds(e, cmd.(string))
		case "move_xy":
			charGameObj.SetMoveToCoordsByXY(params.(map[string]interface{})["x"].(float64), params.(map[string]interface{})["y"].(float64))
		case "get_character_info":
			e.SendResponse("character_info", serializers.GetInfo(e, charGameObj), player)
		case "get_item_info":
			itemId := params.(string)
			if item, itemOk := e.GameObjects().Load(itemId); itemOk {
				e.SendResponse("item_info", serializers.GetInfo(e, item), player)
			}
		case "get_npc_trade_info":
			if npcObj, npcOk := e.GameObjects().Load(params.(string)); npcOk {
				e.SendResponse("npc_trade_info", serializers.GetInfo(e, npcObj), player)
			}
		case "get_npc_dungeons_info":
			if npcObj, npcOk := e.GameObjects().Load(params.(string)); npcOk {
				if dungeonsInfo, diErr := npcObj.(entity.INpcObject).GetDungeonsInfo(e, charGameObj); diErr == nil {
					e.SendResponse("dungeons_info", dungeonsInfo, player)
				}
			}
		case "npc_buy_item":
			if npcObj, npcOk := e.GameObjects().Load(params.(map[string]interface{})["npc_id"].(string)); npcOk {
				npcObj.(entity.INpcObject).BuyItem(e, charGameObj,
					params.(map[string]interface{})["item_name"].(string),
					params.(map[string]interface{})["amount"].(float64))
			}
		case "npc_sell_item":
			if npcObj, npcOk := e.GameObjects().Load(params.(map[string]interface{})["npc_id"].(string)); npcOk {
				npcObj.(entity.INpcObject).SellItem(e, charGameObj,
					params.(map[string]interface{})["item_name"].(string),
					params.(map[string]interface{})["amount"].(float64))
			}
		case "get_craft_atlas":
			e.SendResponse("craft_atlas", craft.GetAtlas(), player)
		case "get_dragons_info":
			e.SendResponse("dragons_info", charGameObj.(entity.ICharacterObject).GetDragonsInfo(e), player)
		case "open_container":
			if cont, contOk := e.GameObjects().Load(params.(string)); contOk {
				container := cont.(entity.IContainerObject)
				if container.CheckAccess(e, player) {
					e.SendResponse("container_items", container.GetItems(e), player)
				} else {
					e.SendSystemMessage("You don't have access to this container", player)
				}
			}
		case "equip_item":
			if item, itemOk := e.GameObjects().Load(params.(string)); itemOk {
				item.(entity.IEquipableObject).Equip(e, player)
			}
		case "unequip_item":
			if item, itemOk := e.GameObjects().Load(params.(string)); itemOk {
				item.(entity.IEquipableObject).Unequip(e, player)
			}
		case "drop_item":
			if item, itemOk := e.GameObjects().Load(params.(string)); itemOk {
				item.(entity.IPickableObject).Drop(e, player)
			}
		case "pickup_item":
			if item, itemOk := e.GameObjects().Load(params.(string)); itemOk {
				item.(entity.IPickableObject).Pickup(e, player)
			}
		case "destroy_item":
			if item, itemOk := e.GameObjects().Load(params.(string)); itemOk {
				item.(entity.IDestroyableObject).Destroy(e, player)
			}
		case "destroy_building":
			if building, buildingOk := e.GameObjects().Load(params.(string)); buildingOk {
				building.(entity.IBuildingObject).Destroy(e, player)
			}
		case "destroy_claim":
			if claim, claimOk := e.GameObjects().Load(params.(string)); claimOk {
				claim.(entity.IClaimObeliskObject).Destroy(e, player)
			}
		case "catch_fish":
			if rod, rodOk := e.GameObjects().Load(params.(string)); rodOk {
				if ok, _ := rod.(entity.IFishingRodObject).CheckCatch(e, charGameObj); ok {
					delayed_actions.Start(e, charGameObj, "CatchFish", map[string]interface{}{
						"characterId": charGameObj.Id(),
						"rodId":    rod.Id(),
					}, -1.0)
				}
			}
		case "close_door":
			if door, doorOk := e.GameObjects().Load(params.(string)); doorOk {
				door.(entity.IDoorObject).Close(e, player)
			}
		case "open_door":
			if door, doorOk := e.GameObjects().Load(params.(string)); doorOk {
				door.(entity.IDoorObject).Open(e, player)
			}
		case "put_to_container":
			if item, itemOk := e.GameObjects().Load(params.(map[string]interface{})["item_id"].(string)); itemOk {
				item.(entity.IPickableObject).PutToContainer(
					e,
					params.(map[string]interface{})["container_id"].(string),
					int(params.(map[string]interface{})["position"].(float64)),
					player)
			}
		case "lift":
			if item, itemOk := e.GameObjects().Load(params.(string)); itemOk {
				charGameObj.SetMoveToCoordsByObject(item)
				delayed_actions.Start(e, charGameObj, "Lift", map[string]interface{}{
					"characterId": charGameObj.Id(),
					"itemId":      item.Id(),
				}, -1.0)
			}
		case "put_lifted":
			// TODO: move to func
			if item, itemOk := e.GameObjects().Load(params.(map[string]interface{})["item_id"].(string)); itemOk {
				x := params.(map[string]interface{})["x"].(float64)
				y := params.(map[string]interface{})["y"].(float64)
				rotation := params.(map[string]interface{})["rotation"].(float64)
				clone := item.Clone()
				clone.SetX(x)
				clone.SetY(y)
				clone.Rotate(rotation)
				charGameObj.SetMoveToCoordsByObject(clone)
				delayed_actions.Start(e, charGameObj, "PutLifted", map[string]interface{}{
					"characterId": charGameObj.Id(),
					"x":           x,
					"y":           y,
					"rotation":    rotation,
					"itemId":      item.Id(),
				}, -1.0)
			}
		case "apply_effect":
			if potion, potionOk := e.GameObjects().Load(params.(string)); potionOk {
				potion.(entity.IPotionObject).ApplyToPlayer(e, player)
			}
		case "chop_tree":
			if tree, treeOk := e.GameObjects().Load(params.(string)); treeOk {
				if tree.(entity.ITreeObject).CheckChop(e, charGameObj) {
					charGameObj.SetMoveToCoordsByObject(tree)
					delayed_actions.Start(e, charGameObj, "Chop", map[string]interface{}{
						"characterId": charGameObj.Id(),
						"treeId":      tree.Id(),
					}, -1.0)
				}
			}
		case "chip_rock":
			if rock, rockOk := e.GameObjects().Load(params.(string)); rockOk {
				if rock.(entity.IRockObject).CheckChip(e, charGameObj) {
					charGameObj.SetMoveToCoordsByObject(rock)
					delayed_actions.Start(e, charGameObj, "Chip", map[string]interface{}{
						"characterId": charGameObj.Id(),
						"rockId":      rock.Id(),
					}, -1.0)
				}
			}
		case "cut_plant":
			if plant, plantOk := e.GameObjects().Load(params.(string)); plantOk {
				if plant.(entity.IPlantObject).CheckCut(e, charGameObj) {
					charGameObj.SetMoveToCoordsByObject(plant)
					delayed_actions.Start(e, charGameObj, "CutPlant", map[string]interface{}{
						"characterId": charGameObj.Id(),
						"plantId":     plant.Id(),
					}, -1.0)
				}
			}
		case "harvest_plant":
			if plant, plantOk := e.GameObjects().Load(params.(string)); plantOk {
				if plant.(entity.IPlantObject).CheckHarvest(e, charGameObj) {
					charGameObj.SetMoveToCoordsByObject(plant)
					delayed_actions.Start(e, charGameObj, "HarvestPlant", map[string]interface{}{
						"characterId": charGameObj.Id(),
						"plantId":     plant.Id(),
					}, -1.0)
				}
			}
		case "dig_surface":
			if shovel, shovelOk := e.GameObjects().Load(params.(string)); shovelOk {
				if shovel.(entity.IShovelObject).CheckDig(e, charGameObj) {
					delayed_actions.Start(e, charGameObj, "Dig", map[string]interface{}{
						"characterId": charGameObj.Id(),
						"shovelId":    shovel.Id(),
					}, -1.0)
				}
			}
		case "craft":
			if craft.Check(e, player, params.(map[string]interface{}), false) {
				params.(map[string]interface{})["playerId"] = float64(player.Id)
				craftItem := params.(map[string]interface{})["item_name"].(string)
				delayed_actions.Start(
					e, charGameObj, "Craft",
					params.(map[string]interface{}),
					craft.GetAtlas()[craftItem].(map[string]interface{})["duration"].(float64))
			}
		case "hatch_fire_dragon":
			if hatchery, hatcheryOk := e.GameObjects().Load(params.(string)); hatcheryOk {
				if hatchery.(entity.IHatcheryObject).CheckHatch(e, charGameObj) {
					delayed_actions.Start(e, hatchery, "Hatch", map[string]interface{}{
						"hatcheryId": hatchery.Id(),
						"mobPath":    hatchery.GetProperty("hatch_mob"),
					}, -1.0)
				}
			}
		case "town_teleport":
			delayed_actions.Start(e, charGameObj, "TownTeleport", map[string]interface{}{
				"playerId": float64(player.Id),
			}, -1.0)
		case "claim_teleport":
			delayed_actions.Start(e, charGameObj, "ClaimTeleport", map[string]interface{}{
				"playerId": float64(player.Id),
			}, -1.0)
		case "teleport_to":
			if teleport, teleportOk := e.GameObjects().Load(params.(string)); teleportOk {
				teleport.(entity.ITeleportObject).TeleportTo(e, charGameObj)
			}
		case "follow":
			mobId := params.(string)
			mob, ok := e.Mobs().Load(mobId)
			if ok {
				mob.OrderToFollow(charGameObj)
			}
		case "attack_my_target":
			mobId := params.(string)
			if mob, mobOk := e.Mobs().Load(mobId); mobOk {
				mob.OrderToAttack(charGameObj)
			}
		case "order_to_stop":
			mobId := params.(string)
			if mob, mobOk := e.Mobs().Load(mobId); mobOk {
				mob.OrderToStop(charGameObj)
			}
		case "select_target":
			targetId := params.(string)
			charGameObj.(entity.ICharacterObject).SelectTarget(e, targetId)
		case "deselect_target":
			charGameObj.(entity.ICharacterObject).DeselectTarget(e)
		case "melee_hit":
			charGameObj.(entity.ICharacterObject).MeleeHit(e)
		case "pay_rent":
			if claim, claimOk := e.GameObjects().Load(params.(string)); claimOk {
				claim.(entity.IClaimObeliskObject).ExtendRent(e)
			}
		case "teleport_dragon_to_owner":
			if dragon, dragonOk := e.GameObjects().Load(params.(string)); dragonOk {
				dragon.(entity.IDragonObject).TeleportToOwner(charGameObj)
			}
		case "release_dragon":
			if dragon, dragonOk := e.GameObjects().Load(params.(string)); dragonOk {
				dragon.(entity.IDragonObject).Release(charGameObj)
			}
		case "resurrect_dragon":
			if dragon, dragonOk := e.GameObjects().Load(params.(string)); dragonOk {
				dragon.(entity.IDragonObject).Resurrect(charGameObj)
			}
		case "evolve":
			if obj, objOk := e.GameObjects().Load(params.(string)); objOk {
				obj.(entity.IEvolvableObject).Evolve(e, player)
			}
		case "feed":
			id := params.(map[string]interface{})["id"].(string)
			foodId := params.(map[string]interface{})["food_id"].(string)
			if obj, objOk := e.GameObjects().Load(id); objOk {
				obj.(entity.IFeedableObject).Feed(e, foodId, player)
			}
		case "add_fuel":
			id := params.(map[string]interface{})["id"].(string)
			fuelId := params.(map[string]interface{})["fuel_id"].(string)
			if obj, objOk := e.GameObjects().Load(id); objOk {
				obj.(entity.IBurningObject).AddFuel(e, fuelId, player)
			}
		case "burn":
			if burningObj, bOk := e.GameObjects().Load(params.(string)); bOk {
				burningObj.(entity.IBurningObject).Burn(e, player)
			}
		case "extinguish":
			if burningObj, bOk := e.GameObjects().Load(params.(string)); bOk {
				burningObj.(entity.IBurningObject).Extinguish(e, player)
			}
		case "gm_create_object":
			gm.CreateObject(e, charGameObj, params.(map[string]interface{}))
		case "gm_update_properties":
			gm.UpdateProperties(e, charGameObj, params.(map[string]interface{}))
		case "go_to_dungeon":
			level := params.(map[string]interface{})["level"].(float64)
			dragonIds := params.(map[string]interface{})["dragonIds"].([]interface{})
			go dungeons.GoToDungeon(e, charGameObj, level, dragonIds)
		case "exit_dungeon":
			if dungeonExit, deOk := e.GameObjects().Load(params.(string)); deOk {
				dungeons.Exit(e, charGameObj, dungeonExit)
			}
		}
	}

	return true
}
