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
					"resources": map[string]interface{}{
						"logs": 3,
					},
				},
				"pine_5": map[string]interface{}{
					"type": "tree",
					"kind": "pine_5",
					"width": 1.0,
					"height": 1.0,
					"shape": "circle",
					"collidable": true,
					"resources": map[string]interface{}{
						"logs": 3,
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
					"resources": map[string]interface{}{
						"stone": 3,
					},
				},
		},
		"tool": {
				"axe": map[string]interface{}{
					"type": "tool",
					"kind": "axe",
					"width": 0.5,
					"height": 1.0,
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
		"resources": {
			"stone": map[string]interface{}{
	
			},
			"logs": map[string]interface{}{
	
			},
		},
	}

	return gameObjectsAtlas
}
