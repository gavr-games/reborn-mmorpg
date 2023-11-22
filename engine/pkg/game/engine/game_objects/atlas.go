package game_objects

import(
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)

//TODO: optimize the memory by using integers instead of string constants
func GetObjectsAtlas() map[string]map[string]interface{} {
	gameObjectsAtlas:= map[string]map[string]interface{}{
		"surface": {
			"grass": map[string]interface{}{
				"type": "surface",
				"kind": "grass",
				"width": 1.0,
				"height": 1.0,
				"shape": "rectangle",
				"visible": true,
			},
			"sand": map[string]interface{}{
				"type": "surface",
				"kind": "sand",
				"width": 1.0,
				"height": 1.0,
				"shape": "rectangle",
				"visible": true,
			},
			"water": map[string]interface{}{
				"type": "surface",
				"kind": "water",
				"width": 1.0,
				"height": 1.0,
				"collidable": true,
				"shape": "rectangle",
				"visible": true,
			},
			"stone": map[string]interface{}{
				"type": "surface",
				"kind": "stone",
				"width": 1.0,
				"height": 1.0,
				"collidable": false,
				"shape": "rectangle",
				"visible": true,
			},
		},
		"tree": {
				"tree_5": map[string]interface{}{
					"type": "tree",
					"kind": "tree_5",
					"width": 1.0,
					"height": 1.0,
					"shape": "circle",
					"collidable": true,
					"visible": true,
					"resources": map[string]interface{}{
						"log": 3.0,
					},
					"actions": map[string]interface{}{
						"chop": map[string]interface{}{
							"cmd": "chop_tree",
							"params": "self", // self - id of current object
						},
					},
				},
				"pine_5": map[string]interface{}{
					"type": "tree",
					"kind": "pine_5",
					"width": 1.0,
					"height": 1.0,
					"shape": "circle",
					"collidable": true,
					"visible": true,
					"resources": map[string]interface{}{
						"log": 3.0,
					},
					"actions": map[string]interface{}{
						"chop": map[string]interface{}{
							"cmd": "chop_tree",
							"params": "self", // self - id of current object
						},
					},
				},
		},
		"rock": {
				"rock_moss": map[string]interface{}{
					"type": "rock",
					"kind": "rock_moss",
					"width": 0.876,
					"height": 1.098,
					"shape": "rectangle",
					"collidable": true,
					"visible": true,
					"resources": map[string]interface{}{
						"stone": 3.0,
					},
					"actions": map[string]interface{}{
						"chip": map[string]interface{}{
							"cmd": "chip_rock",
							"params": "self", // self - id of current object
						},
					},
				},
		},
		"plant": {
			"cactus": map[string]interface{}{
				"type": "plant",
				"kind": "cactus",
				"width": 0.4,
				"height": 0.4,
				"shape": "circle",
				"collidable": false,
				"visible": true,
				"resources": map[string]interface{}{
					"cactus_slice": 3.0,
				},
				"actions": map[string]interface{}{
					"cut": map[string]interface{}{
						"cmd": "cut_cactus",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"axe": { // there might be tools from different materials later
				"axe": map[string]interface{}{
					"type": "axe",
					"kind": "axe",
					"width": 0.624,
					"height": 1.575,
					"shape": "rectangle",
					"container_id": nil,
					"pickable": true,
					"droppable": true,
					"equipable": true,
					"visible": false,
					"target_slots": map[string]interface{}{
						"left_arm": true, 
						"right_arm": true,
					},
					"actions": map[string]interface{}{
						"equip": map[string]interface{}{
							"cmd": "equip_item",
							"params": "self", // self - id of current object
						},
						"unequip": map[string]interface{}{
							"cmd": "unequip_item",
							"params": "self",
						},
						"drop": map[string]interface{}{
							"cmd": "drop_item",
							"params": "self",
						},
						"pickup": map[string]interface{}{
							"cmd": "pickup_item",
							"params": "self",
						},
					},
				},
		},
		"pickaxe": {
				"pickaxe": map[string]interface{}{
					"type": "pickaxe",
					"kind": "pickaxe",
					"width": 0.632,
					"height": 2.255,
					"shape": "rectangle",
					"container_id": nil,
					"pickable": true,
					"droppable": true,
					"equipable": true,
					"visible": false,
					"target_slots": map[string]interface{}{
						"left_arm": true, 
						"right_arm": true,
					},
					"actions": map[string]interface{}{
						"equip": map[string]interface{}{
							"cmd": "equip_item",
							"params": "self", // self - id of current object
						},
						"unequip": map[string]interface{}{
							"cmd": "unequip_item",
							"params": "self",
						},
						"drop": map[string]interface{}{
							"cmd": "drop_item",
							"params": "self",
						},
						"pickup": map[string]interface{}{
							"cmd": "pickup_item",
							"params": "self",
						},
					},
				},
		},
		"knife": {
			"stone_knife": map[string]interface{}{
				"type": "knife",
				"kind": "stone_knife",
				"width": 0.06,
				"height": 0.5,
				"shape": "rectangle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"equipable": true,
				"visible": false,
				"target_slots": map[string]interface{}{
					"left_arm": true, 
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd": "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd": "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"hammer": {
			"stone_hammer": map[string]interface{}{
				"type": "hammer",
				"kind": "stone_hammer",
				"width": 1.181,
				"height": 2.5,
				"shape": "rectangle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"equipable": true,
				"visible": false,
				"target_slots": map[string]interface{}{
					"left_arm": true, 
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd": "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd": "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"container": {
			"backpack": map[string]interface{}{
				"type": "container",
				"kind": "backpack",
				"width": 0.5,
				"height": 0.5,
				"shape": "rectangle",
				"max_capacity": 16,
				"free_capacity": 16.0,
				"size": 4,
				"parent_container_id": nil,
				"owner_id": nil,
				"equipable": true,
				"visible": false,
				"target_slots": map[string]interface{}{
					"back": true,
				},
				"actions": map[string]interface{}{
					"open": map[string]interface{}{
						"cmd": "open_container",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"player": {
			"player": map[string]interface{}{
				"type": "player",
				"kind": "player",
				"width": 1.0,
				"height": 1.0,
				"shape": "circle",
				"speed": constants.PlayerSpeed,
				"speed_x": 0.0,
				"speed_y": 0.0,
				"health": 100.0,
				"max_health": 100.0,
				"visible": true,
				"targetable": true,
				"slots": map[string]interface{}{
					"back": nil,
					"left_arm": nil,
					"right_arm": nil,
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd": "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd": "deselect_target",
						"params": "self",
					},
				},
			},
			"player_vision_area": map[string]interface{}{
				"type": "player",
				"kind": "player_vision_area",
				"width": constants.PlayerVisionArea,
				"height": constants.PlayerVisionArea,
				"shape": "rectangle",
				"visible": false,
			},
		},
		"melee_weapon": {
			"stone_spear": map[string]interface{}{
				"type": "melee_weapon",
				"kind": "stone_spear",
				"width": 2.0,
				"height": 2.0,
				"shape": "rectangle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"damage": 10.0,
				"cooldown": 1000.0, //ms
				"hit_radius": 1.5, // maximum distance to target
				"hit_angle": 90.0, // degrees
				"target_slots": map[string]interface{}{
					"left_arm": true, 
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd": "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd": "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"resource": {
			"stone": map[string]interface{}{
				"type": "resource",
				"kind": "stone",
				"width": 0.629,
				"height": 0.525,
				"shape": "rectangle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"log": map[string]interface{}{
				"type": "resource",
				"kind": "log",
				"width": 0.316,
				"height": 1.335,
				"shape": "rectangle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"cactus_slice": map[string]interface{}{
				"type": "resource",
				"kind": "cactus_slice",
				"width": 0.3,
				"height": 0.3,
				"shape": "circle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"fire_dragon_egg": map[string]interface{}{
				"type": "resource",
				"kind": "fire_dragon_egg",
				"width": 1.0,
				"height": 1.0,
				"shape": "circle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"gold": map[string]interface{}{
				"type": "resource",
				"kind": "gold",
				"width": 1.0,
				"height": 1.0,
				"shape": "circle",
				"amount": 0.0,
				"container_id": nil,
				"pickable": true,
				"stackable": true,
				"droppable": true,
				"visible": false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"claim_stone": map[string]interface{}{
				"type": "resource",
				"kind": "claim_stone",
				"width": 0.5,
				"height": 0.5,
				"shape": "circle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"potion":{
			"healing_balm": map[string]interface{}{
				"type": "potion",
				"kind": "healing_balm",
				"width": 0.25,
				"height": 0.25,
				"shape": "circle",
				"container_id": nil,
				"pickable": true,
				"droppable": true,
				"visible": false,
				"effect": map[string]interface{}{
					"type": "periodic", // periodic (once per cooldown) or constant (constant value for the defined total_time)
					"attribute": "health",
					"value": 5.0,
					"cooldown": 2000.0,
					"current_cooldown": 0.0,
					"number": 10.0,
					"group": "potion_healing", // this is used to prevent multiple effects from one group
				},
				"actions": map[string]interface{}{
					"use": map[string]interface{}{
						"cmd": "apply_effect",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd": "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd": "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"wall": {
			"stone_wall": map[string]interface{}{
				"type": "wall",
				"kind": "stone_wall",
				"width": 1.0,
				"height": 2.0,
				"shape": "rectangle",
				"collidable": true,
				"visible": true,
				"actions": map[string]interface{}{
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"wooden_wall": map[string]interface{}{
				"type": "wall",
				"kind": "wooden_wall",
				"width": 0.3,
				"height": 3.0,
				"shape": "rectangle",
				"collidable": true,
				"visible": true,
				"actions": map[string]interface{}{
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"hatchery": {
			"fire_dragon_hatchery": map[string]interface{}{
				"type": "hatchery",
				"kind": "fire_dragon_hatchery",
				"width": 2.0,
				"height": 2.0,
				"shape": "circle",
				"collidable": true,
				"visible": true,
				"hatching_resources": map[string]interface{}{
					"log": 2.0,
				},
				"actions": map[string]interface{}{
					"hatch": map[string]interface{}{
						"cmd": "hatch_fire_dragon",
						"params": "self", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"claim": {
			"claim_obelisk": map[string]interface{}{
				"type": "claim",
				"kind": "claim_obelisk",
				"width": 1.0,
				"height": 1.0,
				"shape": "rectangle",
				"collidable": true,
				"visible": true,
				"payed_until": nil,
				"current_action": map[string]interface{}{
					"func_name": "InitClaim",
					"params": map[string]interface{}{
						"game_object_id": nil,
					},
					"time_left": 1.0,
				},
				"actions": map[string]interface{}{
					"pay rent": map[string]interface{}{
						"cmd": "pay_rent",
						"params": "self", // self - id of current object
					},
					"info": map[string]interface{}{
						"cmd": "get_item_info",
						"params": "self", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd": "destroy_claim",
						"params": "self", // self - id of current object
					},
				},
			},
			"claim_area": map[string]interface{}{
				"type": "claim",
				"kind": "claim_area",
				"width": constants.ClaimArea,
				"height": constants.ClaimArea,
				"shape": "rectangle",
				"collidable": false,
				"visible": true,
			},
		},
		"mob": {
			"fire_dragon": map[string]interface{}{
				"type": "mob",
				"kind": "fire_dragon",
				"width": 2.0,
				"height": 2.0,
				"shape": "circle",
				"speed": 2.0,
				"speed_x": 0.0,
				"speed_y": 0.0,
				"health": 100.0,
				"max_health": 100.0,
				"collidable": false,
				"visible": true,
				"targetable": true,
				"attack_type": "melee",
				"damage": 10.0,
				"cooldown": 2000.0, //ms
				"hit_radius": 2.0, // maximum distance to target
				"hit_angle": 120.0, // degrees
				"drop": map[string]interface{}{
					"resource/gold": map[string]interface{}{
						"probability": 1.0,
						"min": 10.0,
						"max": 100.0,
					},
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd": "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd": "deselect_target",
						"params": "self",
					},
					"follow": map[string]interface{}{
						"cmd": "follow",
						"params": "self",
					},
					"unfollow": map[string]interface{}{
						"cmd": "unfollow",
						"params": "self",
					},
				},
			},
			"bat": map[string]interface{}{
				"type": "mob",
				"kind": "bat",
				"width": 1.0,
				"height": 1.0,
				"shape": "circle",
				"speed": 2.0,
				"speed_x": 0.0,
				"speed_y": 0.0,
				"health": 50.0,
				"max_health": 50.0,
				"collidable": false,
				"visible": true,
				"targetable": true,
				"attack_type": "melee",
				"damage": 5.0,
				"cooldown": 1000.0, //ms
				"hit_radius": 1.5, // maximum distance to target
				"hit_angle": 70.0, // degrees
				"drop": map[string]interface{}{
					"resource/gold": map[string]interface{}{
						"probability": 1.0,
						"min": 5.0,
						"max": 50.0,
					},
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd": "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd": "deselect_target",
						"params": "self",
					},
				},
			},
		},
		"npc": {
			"town_keeper": map[string]interface{}{
				"type": "npc",
				"kind": "town_keeper",
				"width": 1.0,
				"height": 1.0,
				"shape": "circle",
				"collidable": false,
				"visible": true,
				"sells": map[string]interface{}{ //what items NPC sells
					"resource/claim_stone": map[string]interface{}{
						"amount": 1.0, // how many items will be give
						"resource": "gold", // for price. allows to request price to be payed in different resources
						"price": 10.0, 
					},
				},
				"actions": map[string]interface{}{
					"trade": map[string]interface{}{
						"cmd": "npc_trade_info",
						"params": "self",
					},
				},
			},
		},
	}

	return gameObjectsAtlas
}
