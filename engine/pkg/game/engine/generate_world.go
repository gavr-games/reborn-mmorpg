package engine

import (
	"math/rand"
	"log"

	"github.com/KEINOS/go-noise"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/world_maps"
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
	floorMap := world_maps.NewFloorMap(0, int(floorSize)) // hardcoded 0 floor
	// Terrain
	for x := 0.; x < floorSize; x++ {
		for y := 0.; y < floorSize; y++ {
			noise:= n.Eval64(x / smoothness, y / smoothness)
			//log.Println(fmt.Sprintf("%f:%f:%f", x, y, noise))
			surfaceKind := "surface/grass"
			if noise < WaterLevel {
				surfaceKind = "surface/water"
			} else if noise < SandLevel {
				surfaceKind = "surface/sand"
			}
			CreateGameObject(e, surfaceKind, x, y, 0, nil)
			floorMap.Cells <- &world_maps.WorldCell{x, y, surfaceKind}
		}
	}
	floorMap.Finish <- true // save map to img

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
	// bats
	for i := 0; float64(i) < floorSize / 4; i++ {
		x := 1.0 + rand.Float64() * (floorSize - 1.0)
		y := 1.0 + rand.Float64() * (floorSize - 1.0)
		CreateGameObject(e, "mob/bat", x, y, 0, nil)
	}
}