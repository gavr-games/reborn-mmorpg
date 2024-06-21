package generate_world

import (
	"log"
	"math"

	"github.com/KEINOS/go-noise"
	"pgregory.net/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/world_maps"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	WaterLevel = -0.4
	SandLevel = -0.3
	Smoothness = 10
	TreeProbability = 0.02
	PalmProbability = 0.02
	RockProbability = 0.02
	BatProbability = 0.02
	ZombieProbability = 0.01
	EggProbability = 0.02
	GrassProbability = 0.02
	CactusProbability = 0.1
)

func GenerateSurface(e entity.IEngine) {
	seed := utils.MakeTimestamp()
	n, err := noise.New(noise.Perlin, seed) // Perlin noise generator
	if (err != nil) {
		log.Fatal("GenerateWorld: ", err)
	}
	// Create surface
	surface := entity.NewGameArea("surface", 0, 0, constants.SurfaceSize, constants.SurfaceSize)
	surfaceId := surface.Id()
	e.GameAreas().Store(surfaceId, surface)
	storage.GetClient().GameAreasUpdates <- surface
	gaMap := world_maps.NewGameAreaMap(surface)
	// Terrain
	for x := 0.; x < constants.SurfaceSize; x++ {
		for y := 0.; y < constants.SurfaceSize; y++ {
			surfaceKind := "surface/grass"
			// World surface generation
			noise:= n.Eval64(x / Smoothness, y / Smoothness)
			//log.Println(fmt.Sprintf("%f:%f:%f", x, y, noise))

			if noise < WaterLevel {
				surfaceKind = "surface/water"
			} else if noise < SandLevel {
				surfaceKind = "surface/sand"
				createWithProbability(e, "plant/cactus", x, y, rand.Float64() * math.Pi * 2, surfaceId, nil, CactusProbability)
				createWithProbability(e, "tree/palm_3", x, y, rand.Float64() * math.Pi * 2, surfaceId, nil, PalmProbability)
			} else { //grass
				p := rand.Float64()
				switch {
				case p >= 0.0 && p < 0.2:
					createWithProbability(e, "rock/rock_moss", x, y, rand.Float64() * math.Pi * 2, surfaceId, nil, RockProbability)
				case p >= 0.2 && p < 0.4:
					createWithProbability(e, "tree", x, y, rand.Float64() * math.Pi * 2, surfaceId, nil, TreeProbability)
				case p >= 0.4 && p < 0.6:
					createWithProbability(e, "resource/fire_dragon_egg", x, y, rand.Float64() * math.Pi * 2, surfaceId, map[string]interface{}{
						"visible": true,
					}, EggProbability)
				case p >= 0.6 && p < 0.8:
					createWithProbability(e, "plant/grass_plant", x, y, rand.Float64() * math.Pi * 2, surfaceId, map[string]interface{}{
						"visible": true,
					}, GrassProbability)
				case p >= 0.8 && p < 1.0:
					bat := createWithProbability(e, "mob/bat", x, y, 0.0, surfaceId, nil, BatProbability)
					if bat == nil {
						createWithProbability(e, "mob/zombie", x, y, 0.0, surfaceId, nil, ZombieProbability)
					}
				}
			}
			e.CreateGameObject(surfaceKind, x, y, 0.0, surfaceId, nil)
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: y, SurfaceKind: surfaceKind}
		}
	}
	gaMap.Finish <- true // save map to img
}

func createWithProbability(e entity.IEngine, objPath string, x, y, rotation float64, gameAreaId string, additionalProps map[string]interface{}, objProbability float64) entity.IGameObject {
	probability := rand.Float64()
	if probability <= objProbability {
		return e.CreateGameObject(objPath, x, y, rotation, gameAreaId, additionalProps)
	}
	return nil
}
