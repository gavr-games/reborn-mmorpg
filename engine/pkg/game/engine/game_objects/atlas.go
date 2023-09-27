package game_objects

const (
	PlayerVisionArea = 70.0
	PlayerSpeed = 2.0
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
						"logs": 3.0,
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
						"logs": 3.0,
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
				"speed": PlayerSpeed,
				"speed_x": 0.0,
				"speed_y": 0.0,
				"visible": true,
				"slots": map[string]interface{}{
					"back": nil,
					"left_arm": nil,
					"right_arm": nil,
				},
			},
			"player_vision_area": map[string]interface{}{
				"type": "player",
				"kind": "player_vision_area",
				"width": PlayerVisionArea,
				"height": PlayerVisionArea,
				"shape": "rectangle",
				"visible": false,
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
				"collidable": false,
				"visible": true,
				"actions": map[string]interface{}{},
			},
		},
	}

	return gameObjectsAtlas
}
