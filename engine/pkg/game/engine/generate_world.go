package engine

import (
	"math"
	"math/rand"
	"log"

	"github.com/KEINOS/go-noise"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/world_maps"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	WaterLevel = -0.4
	SandLevel = -0.3
	Smoothness = 10
	TreeProbability = 0.02
	RockProbability = 0.02
	BatProbability = 0.02
	ZombieProbability = 0.01
	EggProbability = 0.02
	GrassProbability = 0.02
	CactusProbability = 0.1
)

func GenerateWorld(e entity.IEngine) {
	seed := utils.MakeTimestamp()
	n, err := noise.New(noise.Perlin, seed) // Perlin noise generator
	if (err != nil) {
		log.Fatal("GenerateWorld: ", err)
	}
	gameArea := entity.NewGameArea("surface", 0, 0, constants.SurfaceSize, constants.SurfaceSize)
	e.GameAreas().Store(gameArea.Id(), gameArea)
	storage.GetClient().GameAreasUpdates <- gameArea
	gaMap := world_maps.NewGameAreaMap(gameArea)
	// Terrain
	for x := 0.; x < constants.SurfaceSize; x++ {
		for y := 0.; y < constants.SurfaceSize; y++ {
			surfaceKind := "surface/grass"
			// Town surface generation
			if (x >= constants.SurfaceSize / 2.0  - constants.TownSize / 2.0 && x < constants.SurfaceSize / 2.0  + constants.TownSize / 2.0) &&
				(y >= constants.SurfaceSize / 2.0  - constants.TownSize / 2.0 && y < constants.SurfaceSize / 2.0  + constants.TownSize / 2.0) {
					surfaceKind = "surface/stone"
				} else { // World surface generation
					noise:= n.Eval64(x / Smoothness, y / Smoothness)
					//log.Println(fmt.Sprintf("%f:%f:%f", x, y, noise))
					
					if noise < WaterLevel {
						surfaceKind = "surface/water"
					} else if noise < SandLevel {
						surfaceKind = "surface/sand"
						createWithProbability(e, "plant/cactus", x, y, gameArea.Id(), nil, CactusProbability)
					} else { //grass
						p := rand.Float64()
						switch {
						case p >= 0.0 && p < 0.2:
							createWithProbability(e, "rock/rock_moss", x, y, gameArea.Id(), nil, RockProbability)
						case p >= 0.2 && p < 0.4:
							createWithProbability(e, "tree", x, y, gameArea.Id(), nil, TreeProbability)
						case p >= 0.4 && p < 0.6:
							createWithProbability(e, "resource/fire_dragon_egg", x, y, gameArea.Id(), map[string]interface{}{
								"visible": true,
							}, EggProbability)
						case p >= 0.6 && p < 0.8:
							createWithProbability(e, "plant/grass_plant", x, y, gameArea.Id(), map[string]interface{}{
								"visible": true,
							}, GrassProbability)
						case p >= 0.8 && p < 1.0:
							bat := createWithProbability(e, "mob/bat", x, y, gameArea.Id(), nil, BatProbability)
							if bat == nil {
								createWithProbability(e, "mob/zombie", x, y, gameArea.Id(), nil, ZombieProbability)
							}
						}
					}
				}
			e.CreateGameObject(surfaceKind, x, y, 0.0, gameArea.Id(), nil)
			gaMap.Cells <- &world_maps.WorldCell{x, y, surfaceKind}
		}
	}
	gaMap.Finish <- true // save map to img
	generateTown(e, gameArea.Id())
}

func createWithProbability(e entity.IEngine, objPath string, x, y float64, gameAreaId string, additionalProps map[string]interface{}, objProbability float64) entity.IGameObject {
	probability := rand.Float64()
	if probability <= objProbability {
		return e.CreateGameObject(objPath, x, y, 0.0, gameAreaId, additionalProps)
	}
	return nil
}

func generateTown(e entity.IEngine, gameAreaId string) {
	wallSize := 3.0
	townCenter := constants.SurfaceSize / 2.0
	townHalfSize := constants.TownSize / 2.0
	// vertical walls
	e.CreateGameObject("wall/wooden_wall", townCenter + townHalfSize + 1.0, townCenter - townHalfSize, 0.0, gameAreaId, nil)
	e.CreateGameObject("wall/wooden_wall", townCenter + townHalfSize + 1.0, townCenter + townHalfSize - wallSize, 0.0, gameAreaId, nil)
	e.CreateGameObject("wall/wooden_wall", townCenter - townHalfSize + 1.0, townCenter - townHalfSize, math.Pi, gameAreaId, nil)
	e.CreateGameObject("wall/wooden_wall", townCenter - townHalfSize + 1.0, townCenter + townHalfSize - wallSize, math.Pi, gameAreaId, nil)
	// horizontal walls
	e.CreateGameObject("wall/wooden_wall", townCenter + townHalfSize - wallSize + 1.0, townCenter - townHalfSize, math.Pi / 2.0, gameAreaId, nil)
	e.CreateGameObject("wall/wooden_wall", townCenter + townHalfSize - wallSize + 1.0, townCenter + townHalfSize, math.Pi * 3 / 2.0, gameAreaId, nil)
	e.CreateGameObject("wall/wooden_wall", townCenter - townHalfSize + 1.0, townCenter - townHalfSize, math.Pi / 2.0, gameAreaId, nil)
	e.CreateGameObject("wall/wooden_wall", townCenter - townHalfSize + 1.0, townCenter + townHalfSize, math.Pi * 3 / 2.0, gameAreaId, nil)
	
	// npc
	e.CreateGameObject("npc/town_keeper", townCenter - townHalfSize + 2.0, townCenter + townHalfSize - 2.0, math.Pi * 2 - math.Pi / 4, gameAreaId, nil)
	e.CreateGameObject("npc/dungeon_keeper", townCenter - townHalfSize + 4.0, townCenter + townHalfSize - 2.0, math.Pi * 2 - math.Pi / 4, gameAreaId, nil)
}
