package craft

func GetAtlas() map[string]interface{} {
	craftAtlas:= map[string]interface{}{
		"stone_wall": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 2.0,
			},
			"title": "Stone Wall",
			"description": "Protects from strangers and keeps your animals safe.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
			}, //tools equipped required to craft something
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 5000.0, // ms
			"width": 1.0,
			"height": 2.0,
		},
		"stone_hammer": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 1.0,
				"log": 1.0,
			},
			"title": "Stone Hammer",
			"description": "Basic hammer used to craft things.",
			"inputs": []string{},
			"tools": []string{},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"stone_spear": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 1.0,
				"log": 1.0,
			},
			"title": "Stone Spear",
			"description": "Basic weapon to defend yourself.",
			"inputs": []string{},
			"tools": []string{
				"axe",
			},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"fire_dragon_hatchery": map[string]interface{}{
			"skill": "taming",
			"resources": map[string]interface{}{
				"stone": 2.0,
				"log": 2.0,
				"fire_dragon_egg": 1.0,
			},
			"title": "Fire Dragon Hatchery",
			"description": "Want a fire dragon? Hatching time is one minute.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
				"axe",
			}, //tools equipped required to craft something
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 10000.0, // ms
			"width": 2.0,
			"height": 2.0,
		},
	}

	return craftAtlas
}
