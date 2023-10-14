package engine

import (
	"math/rand"
	"log"

	"github.com/KEINOS/go-noise"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

const (
	WaterLevel = -0.4
	SandLevel = -0.3
	smoothness = 10
)

func GenerateWorld(e entity.IEngine, floorSize float64) {
	seed := utils.MakeTimestamp()
	n, err := noise.New(noise.Perlin, seed) // Perlin noise generator
	if (err != nil) {
		log.Fatal("GenerateWorld: ", err)
	}
	// Terrain
	for x := 0.; x < floorSize; x++ {
		for y := 0.; y < floorSize; y++ {
			noise:= n.Eval64(x / smoothness, y / smoothness)
			//log.Println(fmt.Sprintf("%f:%f:%f", x, y, noise))
			if noise < WaterLevel {
				CreateGameObject(e, "surface/water", x, y, 0, nil)
			} else if noise < SandLevel {
				CreateGameObject(e, "surface/sand", x, y, 0, nil)
			} else {
				CreateGameObject(e, "surface/grass", x, y, 0, nil)
			}
		}
	}
	// rocks
	for i := 0; float64(i) < floorSize / 4; i++ {
		x := 1.0 + rand.Float64() * (floorSize - 1.0)
		y := 1.0 + rand.Float64() * (floorSize - 1.0)
		CreateGameObject(e, "rock/rock_moss", x, y, 0, nil)
	}
	// trees
	for i := 0; float64(i) < floorSize / 4; i++ {
		x := 1.0 + rand.Float64() * (floorSize - 1.0)
		y := 1.0 + rand.Float64() * (floorSize - 1.0)
		CreateGameObject(e, "tree", x, y, 0, nil)
	}
	// eggs
	for i := 0; float64(i) < floorSize / 4; i++ {
		x := 1.0 + rand.Float64() * (floorSize - 1.0)
		y := 1.0 + rand.Float64() * (floorSize - 1.0)
		CreateGameObject(e, "resource/fire_dragon_egg", x, y, 0, map[string]interface{}{
			"visible": true,
		})
	}
}