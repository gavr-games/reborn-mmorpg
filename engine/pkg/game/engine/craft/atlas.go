package craft

func GetAtlas() map[string]interface{} {
	craftAtlas:= map[string]interface{}{
		"anvil": map[string]interface{}{
			"skill": "blacksmithing",
			"resources": map[string]interface{}{
				"stone": 3.0,
				"iron_ingot": 3.0,
			},
			"title": "Anvil",
			"description": "Used to create some armor, tools and weapons from ingots.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
				"pickaxe",
			}, //tools equipped required to craft something
			"surfaces": []string{
				"grass",
				"stone",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 7000.0, // ms
			"width": 2.0,
			"height": 0.56,
		},
		"bone_needle": map[string]interface{}{
			"skill": "survival",
			"resources": map[string]interface{}{
				"bone": 1.0,
			},
			"title": "Bone needle",
			"description": "A needle to create some bags or leather armor.",
			"inputs": []string{},
			"tools": []string{
				"knife",
			},
			"place_in_real_world": false,
			"duration": 4000.0,
		},
		"bone_saw": map[string]interface{}{
			"skill": "lumberjacking",
			"resources": map[string]interface{}{
				"bone": 2.0,
				"log": 1.0,
			},
			"title": "Bone saw",
			"description": "A solid saw for beginners.",
			"inputs": []string{},
			"tools": []string{
				"knife",
			},
			"place_in_real_world": false,
			"duration": 4000.0,
		},
		"carrot_sprout": map[string]interface{}{
			"skill": "farming",
			"resources": map[string]interface{}{
				"carrot_seed": 1.0,
			},
			"title": "Carrot Sprout",
			"description": "Grows into a ripe carrot ready to harvest.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"shovel",
			}, //tools equipped required to craft something
			"surfaces": []string{
				"dirt",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 500.0, // ms
			"width": 1.0,
			"height": 1.0,
		},
		"claim_obelisk": map[string]interface{}{
			"skill": "householding",
			"resources": map[string]interface{}{
				"stone": 2.0,
				"claim_stone": 1.0,
			},
			"title": "Claim Obelisk",
			"description": "A first step to build your own home",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
			}, //tools equipped required to craft something
			"surfaces": []string{
				"grass",
				"stone",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 20000.0, // ms
			"width": 1.0,
			"height": 1.0,
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
			"surfaces": []string{
				"grass",
				"stone",
				"sand",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 10000.0, // ms
			"width": 2.0,
			"height": 2.0,
		},
		"golden_armor": map[string]interface{}{
			"skill": "blacksmithing",
			"resources": map[string]interface{}{
				"gold_ingot": 4.0,
				"animal_skin": 2.0,
			},
			"title": "Golden Armor",
			"description": "Should protect you a little bit in the fight.",
			"inputs": []string{},
			"tools": []string{
				"hammer",
			},
			"equipment": []string{
				"anvil",
			},
			"place_in_real_world": false,
			"duration": 20000.0,
		},
		"healing_balm": map[string]interface{}{
			"skill": "herbalism",
			"resources": map[string]interface{}{
				"cactus_slice": 2.0,
			},
			"title": "Healing Balm",
			"description": "Useful to heal small wounds.",
			"inputs": []string{},
			"tools": []string{
				"knife",
			},
			"place_in_real_world": false,
			"duration": 7000.0,
		},
		"leather_robe": map[string]interface{}{
			"skill": "tailoring",
			"resources": map[string]interface{}{
				"rope": 2.0,
				"animal_skin": 4.0,
			},
			"title": "Leather Robe",
			"description": "Looks very warm and should have basic protection from spells.",
			"inputs": []string{},
			"tools": []string{
				"knife",
				"needle",
			},
			"place_in_real_world": false,
			"duration": 20000.0,
		},
		"rope": map[string]interface{}{
			"skill": "survival",
			"resources": map[string]interface{}{
				"grass": 4.0,
			},
			"title": "A rope",
			"description": "Very useful rope to craft bags and everything else.",
			"inputs": []string{},
			"tools": []string{
				"knife",
			},
			"place_in_real_world": false,
			"duration": 6000.0,
		},
		"small_bag": map[string]interface{}{
			"skill": "Leatherworking",
			"resources": map[string]interface{}{
				"rope": 1.0,
				"animal_skin": 4.0,
			},
			"title": "Small Bag",
			"description": "Useful for carrying more items.",
			"inputs": []string{},
			"tools": []string{
				"knife",
				"needle",
			},
			"place_in_real_world": false,
			"duration": 15000.0,
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
		"stone_knife": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 1.0,
				"log": 1.0,
			},
			"title": "Stone Knife",
			"description": "Useful to cut something like cactus.",
			"inputs": []string{},
			"tools": []string{
				"axe",
			},
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
			"surfaces": []string{
				"grass",
				"stone",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 5000.0, // ms
			"width": 1.0,
			"height": 2.0,
		},
		"wooden_chest": map[string]interface{}{
			"skill": "lumberjacking",
			"resources": map[string]interface{}{ // TODO: add nails
				"log": 3.0,
				"iron_nails": 2.0,
			},
			"title": "Wooden chest",
			"description": "Great to store some more items. Put it on Claim for more safety.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"saw", "axe",
			}, //tools equipped required to craft something
			"surfaces": []string{
				"grass",
				"stone",
				"sand",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 10000.0, // ms
			"width": 2.0,
			"height": 1.66,
		},
		"wooden_fishing_rod": map[string]interface{}{
			"skill": "fishing",
			"resources": map[string]interface{}{
				"log": 1.0,
				"bone": 1.0,
				"rope": 1.0,
			},
			"title": "Wooden Fishing Rod",
			"description": "Useful to catch some fish.",
			"inputs": []string{},
			"tools": []string{
				"knife",
			},
			"place_in_real_world": false,
			"duration": 10000.0,
		},
		"wooden_shovel": map[string]interface{}{
			"skill": "lumberjacking",
			"resources": map[string]interface{}{
				"log": 2.0,
			},
			"title": "Wooden Shovel",
			"description": "Basic shovel to dig fields for your crops.",
			"inputs": []string{},
			"tools": []string{},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"wooden_wall": map[string]interface{}{
			"skill": "lumberjacking",
			"resources": map[string]interface{}{
				"log": 3.0,
			},
			"title": "Wooden Wall",
			"description": "Protects from strangers and keeps your animals safe.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer", "axe",
			}, //tools equipped required to craft something
			"surfaces": []string{
				"grass",
				"stone",
			}, //allowed surfaces to craft this item on
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 6000.0, // ms
			"width": 0.3,
			"height": 3.0,
		},
	}

	return craftAtlas
}
