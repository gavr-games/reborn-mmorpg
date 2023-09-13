package craft

func GetAtlas() map[string]map[string]interface{} {
	craftAtlas:= map[string]map[string]interface{}{
		//"lumberjacking":map[string]interface{}{
		//	"wooden_wall": map[string]interface{}{
		//		"resources": map[string]interface{}{
		//			"log": 5.0,
		//		},
		//		"duration": 5000.0, // ms
		//	},
		//},
		"stoneworking":map[string]interface{}{
			"stone_wall": map[string]interface{}{
				"resources": map[string]interface{}{
					"stone": 5.0,
				},
				"title": "Stone Wall",
				"description": "Protects from strangers and keeps your animals safe.",
				"inputs": []string{
					"coordinates",
					"rotation",
				},
				"place_in_real_world": true, //place item in real world or put into container
				"duration": 5000.0, // ms
			},
		},
	}

	return craftAtlas
}

func GetSerializableAtlas() map[string]interface{} {
	serializableAtlas := make(map[string]interface{})
	for key, obj := range GetAtlas() {
		serializableAtlas[key] = obj
	}
	return serializableAtlas
}
