package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)

// TODO: optimize the memory by using integers instead of string constants
func GetObjectsAtlas() map[string]map[string]interface{} {
	gameObjectsAtlas := map[string]map[string]interface{}{
		"surface": {
			"grass": map[string]interface{}{
				"type":    "surface",
				"kind":    "grass",
				"width":   1.0,
				"height":  1.0,
				"shape":   "rectangle",
				"visible": true,
			},
			"dirt": map[string]interface{}{
				"type":    "surface",
				"kind":    "dirt",
				"width":   1.0,
				"height":  1.0,
				"shape":   "rectangle",
				"visible": true,
				"current_action": map[string]interface{}{
					"func_name": "GrowGrass",
					"params": map[string]interface{}{
						"game_object_id": nil,
					},
					"time_left": 60000.0,
				},
			},
			"sand": map[string]interface{}{
				"type":    "surface",
				"kind":    "sand",
				"width":   1.0,
				"height":  1.0,
				"shape":   "rectangle",
				"visible": true,
			},
			"water": map[string]interface{}{
				"type":       "surface",
				"kind":       "water",
				"width":      1.0,
				"height":     1.0,
				"collidable": true,
				"shape":      "rectangle",
				"visible":    true,
			},
			"stone": map[string]interface{}{
				"type":       "surface",
				"kind":       "stone",
				"width":      1.0,
				"height":     1.0,
				"collidable": false,
				"shape":      "rectangle",
				"visible":    true,
			},
			"dungeon_floor": map[string]interface{}{
				"type":    "surface",
				"kind":    "dungeon_floor",
				"width":   1.0,
				"height":  1.0,
				"shape":   "rectangle",
				"visible": true,
			},
		},
		"tree": {
			"tree_5": map[string]interface{}{
				"type":       "tree",
				"kind":       "tree_5",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": true,
				"visible":    true,
				"resources": map[string]interface{}{
					"log": 3.0,
				},
				"actions": map[string]interface{}{
					"chop": map[string]interface{}{
						"cmd":    "chop_tree",
						"params": "self", // self - id of current object
					},
				},
			},
			"pine_5": map[string]interface{}{
				"type":       "tree",
				"kind":       "pine_5",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": true,
				"visible":    true,
				"resources": map[string]interface{}{
					"log": 3.0,
				},
				"actions": map[string]interface{}{
					"chop": map[string]interface{}{
						"cmd":    "chop_tree",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"rock": {
			"rock_moss": map[string]interface{}{
				"type":       "rock",
				"kind":       "rock_moss",
				"width":      0.876,
				"height":     1.098,
				"shape":      "rectangle",
				"collidable": true,
				"visible":    true,
				"resources": map[string]interface{}{
					"stone": 3.0,
				},
				"actions": map[string]interface{}{
					"chip": map[string]interface{}{
						"cmd":    "chip_rock",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"plant": {
			"carrot_sprout": map[string]interface{}{
				"type":       "plant",
				"kind":       "carrot_sprout",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"grows_into": "plant/carrot_ripe",
				"resources": map[string]interface{}{
					"carrot_seed": 1.0,
				},
				"current_action": map[string]interface{}{
					"func_name": "GrowPlant",
					"params": map[string]interface{}{
						"game_object_id": nil,
					},
					"time_left": 60000.0,
				},
				"actions": map[string]interface{}{
					"harvest": map[string]interface{}{
						"cmd":    "harvest_plant",
						"params": "self", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"carrot_ripe": map[string]interface{}{
				"type":       "plant",
				"kind":       "carrot_ripe",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"resources": map[string]interface{}{
					"carrot_seed": 2.0,
					"carrot": 1.0,
				},
				"actions": map[string]interface{}{
					"harvest": map[string]interface{}{
						"cmd":    "harvest_plant",
						"params": "self", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"cactus": map[string]interface{}{
				"type":       "plant",
				"kind":       "cactus",
				"width":      0.4,
				"height":     0.4,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"resources": map[string]interface{}{
					"cactus_slice": 3.0,
				},
				"actions": map[string]interface{}{
					"cut": map[string]interface{}{
						"cmd":    "cut_plant",
						"params": "self", // self - id of current object
					},
				},
			},
			"grass_plant": map[string]interface{}{
				"type":       "plant",
				"kind":       "grass_plant",
				"width":      0.5,
				"height":     0.5,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"resources": map[string]interface{}{
					"grass": 3.0,
				},
				"actions": map[string]interface{}{
					"cut": map[string]interface{}{
						"cmd":    "cut_plant",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"axe": { // there might be tools from different materials later
			"axe": map[string]interface{}{
				"type":         "axe",
				"kind":         "axe",
				"width":        0.624,
				"height":       1.575,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
				},
			},
		},
		"pickaxe": {
			"pickaxe": map[string]interface{}{
				"type":         "pickaxe",
				"kind":         "pickaxe",
				"width":        0.632,
				"height":       2.255,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
				},
			},
		},
		"knife": {
			"stone_knife": map[string]interface{}{
				"type":         "knife",
				"kind":         "stone_knife",
				"width":        0.1,
				"height":       1.0,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"hammer": {
			"stone_hammer": map[string]interface{}{
				"type":         "hammer",
				"kind":         "stone_hammer",
				"width":        1.181,
				"height":       2.5,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"saw": {
			"bone_saw": map[string]interface{}{
				"type":         "saw",
				"kind":         "bone_saw",
				"width":        0.9,
				"height":       0.33,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"shovel": {
			"wooden_shovel": map[string]interface{}{
				"type":         "shovel",
				"kind":         "wooden_shovel",
				"width":        0.5,
				"height":       2.0,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"dig": map[string]interface{}{
						"cmd":    "dig_surface",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"needle": {
			"bone_needle": map[string]interface{}{
				"type":         "needle",
				"kind":         "bone_needle",
				"width":        0.5,
				"height":       0.02,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"fishing_rod": {
			"wooden_fishing_rod": map[string]interface{}{
				"type":         "fishing_rod",
				"kind":         "wooden_fishing_rod",
				"width":        2.0,
				"height":       0.1,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"equipable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"catch_fish": map[string]interface{}{
						"cmd":    "catch_fish",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"bonfire": {
			"bonfire": map[string]interface{}{
				"type":       "bonfire",
				"kind":       "bonfire",
				"width":      1.0,
				"height":     1.0,
				"shape":      "rectangle",
				"state":      "extinguished", // extinguished/burning
				"fuel":       0.0,
				"max_fuel":   300000.0, // 5 minutes
				"collidable": true,
				"visible":    true,
				"allowed_fuels": map[string]interface{}{
					"log": 20000.0, // how many seconds burns on fuel
				},
				"actions": map[string]interface{}{
					"burn": map[string]interface{}{
						"cmd":    "burn",
						"params": "self", // self - id of current object
					},
					"extinguish": map[string]interface{}{
						"cmd":    "extinguish",
						"params": "self", // self - id of current object
					},
					"add fuel": map[string]interface{}{
						"cmd":    "add_fuel",
						"params": "self,fuel_id", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"container": {
			"backpack": map[string]interface{}{
				"type":                "container",
				"kind":                "backpack",
				"width":               0.5,
				"height":              0.5,
				"shape":               "rectangle",
				"max_capacity":        16.0,
				"free_capacity":       16.0,
				"size":                4.0,
				"parent_container_id": nil,
				"owner_id":            nil,
				"equipable":           true,
				"visible":             false,
				"target_slots": map[string]interface{}{
					"back": true,
				},
				"actions": map[string]interface{}{
					"open": map[string]interface{}{
						"cmd":    "open_container",
						"params": "self", // self - id of current object
					},
				},
			},
			"small_bag": map[string]interface{}{
				"type":                "container",
				"kind":                "small_bag",
				"width":               0.5,
				"height":              0.48,
				"shape":               "circle",
				"max_capacity":        4.0,
				"free_capacity":       4.0,
				"size":                2.0,
				"parent_container_id": nil,
				"owner_id":            nil,
				"equipable":           true,
				"visible":             false,
				"actions": map[string]interface{}{
					"open": map[string]interface{}{
						"cmd":    "open_container",
						"params": "self", // self - id of current object
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"wooden_chest": map[string]interface{}{
				"type":                "container",
				"kind":                "wooden_chest",
				"width":               2.0,
				"height":              1.66,
				"shape":               "rectangle",
				"max_capacity":        32.0,
				"free_capacity":       32.0,
				"size":                8.0,
				"parent_container_id": nil,
				"owner_id":            nil,
				"visible":             true,
				"collidable":          true,
				"liftable":            true,
				"actions": map[string]interface{}{
					"open": map[string]interface{}{
						"cmd":    "open_container",
						"params": "self", // self - id of current object
					},
					"lift": map[string]interface{}{
						"cmd":    "lift",
						"params": "self",
					},
					"put": map[string]interface{}{
						"cmd":    "put_lifted",
						"params": "self,coordinates,rotation",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"dungeon_chest": map[string]interface{}{
				"type":                "container",
				"kind":                "dungeon_chest",
				"width":               2.0,
				"height":              1.66,
				"shape":               "rectangle",
				"max_capacity":        32.0,
				"free_capacity":       32.0,
				"size":                8.0,
				"parent_container_id": nil,
				"owner_id":            nil,
				"visible":             true,
				"collidable":          true,
				"liftable":            false,
				"actions": map[string]interface{}{
					"open": map[string]interface{}{
						"cmd":    "open_container",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"player": {
			"player": map[string]interface{}{
				"type":        "player",
				"kind":        "player",
				"width":       1.0,
				"height":      1.0,
				"shape":       "circle",
				"speed":       constants.PlayerSpeed,
				"speed_x":     0.0,
				"speed_y":     0.0,
				"health":      100.0,
				"max_health":  100.0,
				"level":       0.0,
				"experience":  0.0,
				"claim_obelisk_id": nil,
				"dragons_ids": []interface{}{},
				"max_dragons": 1.0,
				"max_dungeon_lvl": 1.0,
				"current_dungeon_id": nil,
				"visible":     true,
				"targetable":  true,
				"game_master": false,
				"slots": map[string]interface{}{
					"back":      nil,
					"body":      nil,
					"left_arm":  nil,
					"right_arm": nil,
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd":    "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd":    "deselect_target",
						"params": "self",
					},
				},
			},
			"player_vision_area": map[string]interface{}{
				"type":    "player",
				"kind":    "player_vision_area",
				"width":   constants.PlayerVisionArea,
				"height":  constants.PlayerVisionArea,
				"shape":   "rectangle",
				"visible": false,
				"craft_collidable": false,
			},
		},
		"armor": {
			"golden_armor": map[string]interface{}{
				"type":         "armor",
				"kind":         "golden_armor",
				"width":        1.9,
				"height":       0.8,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"body":  true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"leather_robe": map[string]interface{}{
				"type":         "armor",
				"kind":         "leather_robe",
				"width":        1.33,
				"height":       0.58,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"target_slots": map[string]interface{}{
					"body":  true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"melee_weapon": {
			"stone_spear": map[string]interface{}{
				"type":         "melee_weapon",
				"kind":         "stone_spear",
				"width":        2.0,
				"height":       2.0,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"damage":       10.0,
				"cooldown":     1000.0, //ms
				"hit_radius":   1.5,    // maximum distance to target
				"hit_angle":    90.0,   // degrees
				"target_slots": map[string]interface{}{
					"left_arm":  true,
					"right_arm": true,
				},
				"actions": map[string]interface{}{
					"equip": map[string]interface{}{
						"cmd":    "equip_item",
						"params": "self", // self - id of current object
					},
					"unequip": map[string]interface{}{
						"cmd":    "unequip_item",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"equipment": {
			"anvil": map[string]interface{}{
				"type":       "equipment",
				"kind":       "anvil",
				"width":      2.0,
				"height":     0.56,
				"shape":      "rectangle",
				"collidable": true,
				"visible":    true,
				"actions": map[string]interface{}{
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
			"dragon_altar": map[string]interface{}{
				"type":       "equipment",
				"kind":       "dragon_altar",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"actions": map[string]interface{}{
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"resource": {
			"carrot": map[string]interface{}{
				"type":         "resource",
				"kind":         "carrot",
				"width":        0.24,
				"height":       0.88,
				"shape":        "rectangle",
				"container_id": nil,
				"fullness":     10.0,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"eatable":      true,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"carrot_seed": map[string]interface{}{
				"type":         "resource",
				"kind":         "carrot_seed",
				"width":        0.75,
				"height":       0.5,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"fish": map[string]interface{}{
				"type":         "resource",
				"kind":         "fish",
				"width":        1.0,
				"height":       0.3,
				"shape":        "rectangle",
				"container_id": nil,
				"fullness":     5.0,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"eatable":      true,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"stone": map[string]interface{}{
				"type":         "resource",
				"kind":         "stone",
				"width":        0.629,
				"height":       0.525,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"log": map[string]interface{}{
				"type":         "resource",
				"kind":         "log",
				"width":        0.316,
				"height":       1.335,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"cactus_slice": map[string]interface{}{
				"type":         "resource",
				"kind":         "cactus_slice",
				"width":        0.3,
				"height":       0.3,
				"shape":        "circle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"grass": map[string]interface{}{
				"type":         "resource",
				"kind":         "grass",
				"width":        0.5,
				"height":       0.185,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"rope": map[string]interface{}{
				"type":         "resource",
				"kind":         "rope",
				"width":        0.9,
				"height":       0.75,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"fire_dragon_egg": map[string]interface{}{
				"type":         "resource",
				"kind":         "fire_dragon_egg",
				"width":        1.0,
				"height":       1.0,
				"shape":        "circle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"bone": map[string]interface{}{
				"type":         "resource",
				"kind":         "bone",
				"width":        0.3,
				"height":       1.0,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"animal_skin": map[string]interface{}{
				"type":         "resource",
				"kind":         "animal_skin",
				"width":        0.8,
				"height":       1.0,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"gold": map[string]interface{}{
				"type":         "resource",
				"kind":         "gold",
				"width":        1.0,
				"height":       1.0,
				"shape":        "circle",
				"amount":       0.0,
				"container_id": nil,
				"pickable":     true,
				"stackable":    true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"claim_stone": map[string]interface{}{
				"type":         "resource",
				"kind":         "claim_stone",
				"width":        0.5,
				"height":       0.5,
				"shape":        "circle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"iron_nails": map[string]interface{}{
				"type":         "resource",
				"kind":         "iron_nails",
				"width":        0.7,
				"height":       0.7,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"iron_ingot": map[string]interface{}{
				"type":         "resource",
				"kind":         "iron_ingot",
				"width":        1.0,
				"height":       0.15,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"gold_ingot": map[string]interface{}{
				"type":         "resource",
				"kind":         "gold_ingot",
				"width":        1.0,
				"height":       0.15,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
			"dungeon_key": map[string]interface{}{
				"type":         "resource",
				"kind":         "dungeon_key",
				"width":        0.23,
				"height":       0.5,
				"shape":        "rectangle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"actions": map[string]interface{}{
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self",
					},
				},
			},
		},
		"potion": {
			"healing_balm": map[string]interface{}{
				"type":         "potion",
				"kind":         "healing_balm",
				"width":        0.25,
				"height":       0.25,
				"shape":        "circle",
				"container_id": nil,
				"pickable":     true,
				"droppable":    true,
				"visible":      false,
				"effect": map[string]interface{}{
					"type":             "periodic", // periodic (once per cooldown) or constant (constant value for the defined total_time)
					"attribute":        "health",
					"value":            5.0,
					"cooldown":         2000.0,
					"current_cooldown": 0.0,
					"number":           10.0,
					"cant_go_negative": true,
					"group":            "potion_healing", // this is used to prevent multiple effects from one group
				},
				"actions": map[string]interface{}{
					"use": map[string]interface{}{
						"cmd":    "apply_effect",
						"params": "self",
					},
					"drop": map[string]interface{}{
						"cmd":    "drop_item",
						"params": "self",
					},
					"pickup": map[string]interface{}{
						"cmd":    "pickup_item",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_item",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"wall": {
			"stone_wall": map[string]interface{}{
				"type":       "wall",
				"kind":       "stone_wall",
				"width":      1.0,
				"height":     2.0,
				"shape":      "rectangle",
				"collidable": true,
				"visible":    true,
				"actions": map[string]interface{}{
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
			"wooden_wall": map[string]interface{}{
				"type":       "wall",
				"kind":       "wooden_wall",
				"width":      0.3,
				"height":     3.0,
				"shape":      "rectangle",
				"collidable": true,
				"visible":    true,
				"actions": map[string]interface{}{
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
			"dungeon_exit": map[string]interface{}{
				"type":         "wall",
				"kind":         "dungeon_exit",
				"width":        2.0,
				"height":       0.27,
				"shape":        "rectangle",
				"level":        nil,
				"character_id": nil,
				"collidable":   true,
				"visible":      true,
				"actions": map[string]interface{}{
					"exit": map[string]interface{}{
						"cmd":    "exit_dungeon",
						"params": "self",
					},
				},
			},
			"dungeon_wall": map[string]interface{}{
				"type":         "wall",
				"kind":         "dungeon_wall",
				"width":        2.0,
				"height":       0.4,
				"shape":        "rectangle",
				"collidable":   true,
				"visible":      true,
			},
			"dungeon_column": map[string]interface{}{
				"type":         "wall",
				"kind":         "dungeon_column",
				"width":        0.5,
				"height":       0.5,
				"shape":        "rectangle",
				"collidable":   true,
				"visible":      true,
			},
		},
		"door": {
			"wooden_door": map[string]interface{}{
				"type":       "door",
				"kind":       "wooden_door",
				"width":      2.0,
				"height":     0.3,
				"shape":      "rectangle",
				"state":      "closed", // opened or closed
				"collidable": true,
				"visible":    true,
				"actions": map[string]interface{}{
					"open": map[string]interface{}{
						"cmd":    "open_door",
						"params": "self",
					},
					"close": map[string]interface{}{
						"cmd":    "close_door",
						"params": "self",
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"hatchery": {
			"fire_dragon_hatchery": map[string]interface{}{
				"type":       "hatchery",
				"kind":       "fire_dragon_hatchery",
				"width":      2.0,
				"height":     2.0,
				"shape":      "circle",
				"collidable": true,
				"visible":    true,
				"hatch_mob": "mob/baby_fire_dragon",
				"hatching_resources": map[string]interface{}{
					"log": 2.0,
				},
				"actions": map[string]interface{}{
					"hatch": map[string]interface{}{
						"cmd":    "hatch_fire_dragon",
						"params": "self", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_building",
						"params": "self", // self - id of current object
					},
				},
			},
		},
		"claim": {
			"claim_obelisk": map[string]interface{}{
				"type":        "claim",
				"kind":        "claim_obelisk",
				"width":       1.0,
				"height":      1.0,
				"shape":       "rectangle",
				"collidable":  true,
				"visible":     true,
				"payed_until": nil,
				"claim_area_id": nil,
				"current_action": map[string]interface{}{
					"func_name": "InitClaim",
					"params": map[string]interface{}{
						"game_object_id": nil,
					},
					"time_left": 1.0,
				},
				"actions": map[string]interface{}{
					"pay rent": map[string]interface{}{
						"cmd":    "pay_rent",
						"params": "self", // self - id of current object
					},
					"info": map[string]interface{}{
						"cmd":    "get_item_info",
						"params": "self", // self - id of current object
					},
					"destroy": map[string]interface{}{
						"cmd":    "destroy_claim",
						"params": "self", // self - id of current object
					},
				},
			},
			"claim_area": map[string]interface{}{
				"type":       "claim",
				"kind":       "claim_area",
				"width":      constants.ClaimArea,
				"height":     constants.ClaimArea,
				"shape":      "rectangle",
				"claim_obelisk_id": nil,
				"collidable": false,
				"craft_collidable": false,
				"visible":    true,
			},
		},
		"mob": {
			"baby_fire_dragon": map[string]interface{}{
				"type":        "mob",
				"kind":        "baby_fire_dragon",
				"width":       1.0,
				"height":      1.0,
				"shape":       "circle",
				"speed":       1.5,
				"speed_x":     0.0,
				"speed_y":     0.0,
				"health":      50.0,
				"max_health":  50.0,
				"level":       0.0,
				"experience":  0.0,
				"fullness":       0.0,
				"max_fullness":   50.0,
				"food_to_evolve": 60.0,
				"feedable":    true,
				"collidable":  false,
				"visible":     true,
				"targetable":  true,
				"alive":       true,
				"attack_type": "melee",
				"damage":      4.0,
				"cooldown":    3000.0, //ms
				"hit_radius":  1.5,    // maximum distance to target
				"hit_angle":   80.0,  // degrees
				"owner_id":    nil,
				"evolve_to":   "fire_dragon",
				"effects": map[string]interface{}{
					"hunger": map[string]interface{}{
						"type":             "periodic", // periodic (once per cooldown) or constant (constant value for the defined total_time)
						"attribute":        "fullness",
						"value":            -1.0,
						"cooldown":         6000.0,
						"current_cooldown": 0.0,
						"number":           -1.0, // infinite repeat
						"group":            "hunger", // this is used to prevent multiple effects from one group
					},
				},
				"drop": map[string]interface{}{
					"resource/gold": map[string]interface{}{
						"probability": 1.0,
						"min":         5.0,
						"max":         50.0,
					},
					"resource/bone": map[string]interface{}{
						"probability": 0.4,
					},
					"resource/animal_skin": map[string]interface{}{
						"probability": 0.4,
					},
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd":    "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd":    "deselect_target",
						"params": "self",
					},
					"follow": map[string]interface{}{
						"cmd":    "follow",
						"params": "self",
					},
					"stop": map[string]interface{}{
						"cmd":    "order_to_stop",
						"params": "self",
					},
					"evolve": map[string]interface{}{
						"cmd":    "evolve",
						"params": "self",
					},
					"feed": map[string]interface{}{
						"cmd":    "feed",
						"params": "self,food_id",
					},
					"info": map[string]interface{}{
						"cmd":    "get_item_info",
						"params": "self",
					},
				},
			},
			"fire_dragon": map[string]interface{}{
				"type":        "mob",
				"kind":        "fire_dragon",
				"width":       2.0,
				"height":      2.0,
				"shape":       "circle",
				"speed":       2.0,
				"speed_x":     0.0,
				"speed_y":     0.0,
				"health":      100.0,
				"max_health":  100.0,
				"level":       0.0,
				"experience":  0.0,
				"collidable":  false,
				"visible":     true,
				"targetable":  true,
				"alive":       true,
				"attack_type": "melee",
				"damage":      10.0,
				"cooldown":    2000.0, //ms
				"hit_radius":  2.0,    // maximum distance to target
				"hit_angle":   120.0,  // degrees
				"owner_id":    nil,
				"drop": map[string]interface{}{
					"resource/gold": map[string]interface{}{
						"probability": 1.0,
						"min":         10.0,
						"max":         100.0,
					},
					"resource/bone": map[string]interface{}{
						"probability": 0.7,
					},
					"resource/animal_skin": map[string]interface{}{
						"probability": 0.7,
					},
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd":    "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd":    "deselect_target",
						"params": "self",
					},
					"follow": map[string]interface{}{
						"cmd":    "follow",
						"params": "self",
					},
					"stop": map[string]interface{}{
						"cmd":    "order_to_stop",
						"params": "self",
					},
					"info": map[string]interface{}{
						"cmd":    "get_item_info",
						"params": "self",
					},
				},
			},
			"bat": map[string]interface{}{
				"type":        "mob",
				"kind":        "bat",
				"width":       1.0,
				"height":      1.0,
				"shape":       "circle",
				"speed":       2.0,
				"speed_x":     0.0,
				"speed_y":     0.0,
				"health":      50.0,
				"max_health":  50.0,
				"collidable":  false,
				"visible":     true,
				"targetable":  true,
				"alive":       true,
				"attack_type": "melee",
				"damage":      5.0,
				"cooldown":    1000.0, //ms
				"hit_radius":  1.5,    // maximum distance to target
				"hit_angle":   70.0,   // degrees
				"drop": map[string]interface{}{
					"resource/gold": map[string]interface{}{
						"probability": 1.0,
						"min":         5.0,
						"max":         50.0,
					},
					"resource/bone": map[string]interface{}{
						"probability": 0.7,
					},
					"resource/animal_skin": map[string]interface{}{
						"probability": 0.7,
					},
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd":    "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd":    "deselect_target",
						"params": "self",
					},
				},
			},
			"zombie": map[string]interface{}{
				"type":        "mob",
				"kind":        "zombie",
				"width":       1.0,
				"height":      1.0,
				"shape":       "circle",
				"speed":       2.0,
				"speed_x":     0.0,
				"speed_y":     0.0,
				"health":      60.0,
				"max_health":  60.0,
				"collidable":  false,
				"visible":     true,
				"targetable":  true,
				"alive":       true,
				"agressive":   true,
				"attack_type": "melee",
				"damage":      6.0,
				"cooldown":    2000.0, //ms
				"hit_radius":  1.7,    // maximum distance to target
				"hit_angle":   80.0,   // degrees
				"agro_radius":  5.0,    // if player is closer, mob attacks them
				"drop": map[string]interface{}{
					"resource/gold": map[string]interface{}{
						"probability": 1.0,
						"min":         10.0,
						"max":         60.0,
					},
					"resource/bone": map[string]interface{}{
						"probability": 0.8,
					},
					"resource/animal_skin": map[string]interface{}{
						"probability": 0.3,
					},
				},
				"actions": map[string]interface{}{
					"select as target": map[string]interface{}{
						"cmd":    "select_target",
						"params": "self",
					},
					"deselect target": map[string]interface{}{
						"cmd":    "deselect_target",
						"params": "self",
					},
				},
			},
		},
		"npc": {
			"town_keeper": map[string]interface{}{
				"type":       "npc",
				"kind":       "town_keeper",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"sells": map[string]interface{}{ //what items NPC sells
					"resource/claim_stone": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    10.0,
					},
					"resource/carrot_seed": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    3.0,
					},
					"resource/iron_nails": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    5.0,
					},
					"resource/iron_ingot": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    25.0,
					},
					"resource/gold_ingot": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    50.0,
					},
				},
				"buys": map[string]interface{}{
					"resource/fire_dragon_egg": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    1.0,
					},
					"resource/log": map[string]interface{}{
						"amount":   1.0,
						"resource": "gold",
						"price":    2.0,
					},
					"resource/cactus_slice": map[string]interface{}{
						"amount":   1.0,
						"resource": "fire_dragon_egg",
						"price":    3.0,
					},
				},
				"actions": map[string]interface{}{
					"trade": map[string]interface{}{
						"cmd":    "get_npc_trade_info",
						"params": "self",
					},
				},
			},
			"dungeon_keeper": map[string]interface{}{
				"type":       "npc",
				"kind":       "dungeon_keeper",
				"width":      1.0,
				"height":     1.0,
				"shape":      "circle",
				"collidable": false,
				"visible":    true,
				"actions": map[string]interface{}{
					"dungeons": map[string]interface{}{
						"cmd":    "get_npc_dungeons_info",
						"params": "self",
					},
				},
			},
		},
	}

	return gameObjectsAtlas
}
