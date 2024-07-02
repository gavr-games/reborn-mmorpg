package generate_world

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/world_maps"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const(
	TownSize = 50.0
	WallSize = 2.0
	WindmillX = 30
	WindmillY = 30
	BellTowerX = 18
	BellTowerY = 40
	BlacksmithX = 8
	BlacksmithY = 16
	SawmillX = 4
	SawmillY = 6
	InnX = 6
	InnY = 28
	HouseX = 39
	HouseY = 31
	TownKeeperX = 33
	TownKeeperY = 28
	DungeonKeeperX = 44
	DungeonKeeperY = 28
	WellX = 36
	WellY = 20
	BonfireX = 36
	BonfireY = 17
	TablesX = 30
	TablesY = 20
	FarmX = 32.0
	FarmY = 4.0
	FarmWidth = 14.0
	FarmHeight = 6.0
)

var TreesX = [...]float64 {4.0, 18.0, 38.0, 46.0, 28.0}
var TreesY = [...]float64 {3.0, 37.0, 28.0, 13.0, 5.0}

func GenerateTown(e entity.IEngine) {
	// Create town game area
	town := entity.NewGameArea("town", 0, 0, TownSize, TownSize)
	townId := town.Id()
	e.GameAreas().Store(townId, town)
	storage.GetClient().GameAreasUpdates <- town
	gaMap := world_maps.NewGameAreaMap(town)

	for x := 0.0; x < TownSize; x++ {
		// Create moat
		e.CreateGameObject("surface/water", x, 0.0, 0.0, townId, nil)
		e.CreateGameObject("surface/water", x, TownSize - 1.0, 0.0, townId, nil)
		gaMap.Cells <- &world_maps.WorldCell{X: x, Y: 0.0, SurfaceKind: "surface/water"}
		gaMap.Cells <- &world_maps.WorldCell{X: x, Y: TownSize - 1.0, SurfaceKind: "surface/water"}
		// Create walls
		if x > 0.0 && x < TownSize - WallSize && int64(x) % WallSize == 0.0 {
			e.CreateGameObject("wall/brick_wall", x, WallSize, 0.0, townId, nil)
			e.CreateGameObject("wall/brick_wall", x, TownSize - WallSize, 0.0, townId, nil)
		}
		// Create main horizontal road
		if x >= WallSize && x < TownSize - WallSize {
			e.CreateGameObject("surface/stone_road", x, TownSize / 2.0, 0.0, townId, nil)
			e.CreateGameObject("surface/stone_road", x, TownSize / 2.0 + 1.0, 0.0, townId, nil)
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: TownSize / 2.0, SurfaceKind: "surface/stone_road"}
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: TownSize / 2.0 + 1.0, SurfaceKind: "surface/stone_road"}
		}
	}
	for y := 1.0; y < TownSize - 1.0; y++ {
		// Create moat
		e.CreateGameObject("surface/water", 0.0, y, 0.0, townId, nil)
		e.CreateGameObject("surface/water", TownSize - 1.0, y, 0.0, townId, nil)
		gaMap.Cells <- &world_maps.WorldCell{X: 0.0, Y: y, SurfaceKind: "surface/water"}
		gaMap.Cells <- &world_maps.WorldCell{X: TownSize - 1.0, Y: 0.0, SurfaceKind: "surface/water"}
		// Create walls
		if y > 0.0 && y < TownSize - WallSize && int64(y) % WallSize == 0.0 {
			e.CreateGameObject("wall/brick_wall", WallSize, y, math.Pi / 2.0, townId, nil)
			e.CreateGameObject("wall/brick_wall", TownSize - WallSize, y, math.Pi / 2.0, townId, nil)
		}
		// Create main vertical road
		if y >= WallSize && y < TownSize - WallSize {
			e.CreateGameObject("surface/stone_road", TownSize / 2.0, y, 0.0, townId, nil)
			e.CreateGameObject("surface/stone_road", TownSize / 2.0 + 1.0, y, 0.0, townId, nil)
			gaMap.Cells <- &world_maps.WorldCell{X: TownSize / 2.0, Y: y, SurfaceKind: "surface/stone_road"}
			gaMap.Cells <- &world_maps.WorldCell{X: TownSize / 2.0 + 1.0, Y: y, SurfaceKind: "surface/stone_road"}
		}
	}
	// Create grass
	for x := 1.0; x < TownSize - 2.0; x++ {
		e.CreateGameObject("surface/grass", x, 1.0, 0.0, townId, nil)
		e.CreateGameObject("surface/grass", x, TownSize - 2.0, 0.0, townId, nil)
		gaMap.Cells <- &world_maps.WorldCell{X: x, Y: 1.0, SurfaceKind: "surface/grass"}
		gaMap.Cells <- &world_maps.WorldCell{X: x, Y: TownSize - 2.0, SurfaceKind: "surface/grass"}
	}
	for y := 1.0; y < TownSize - 2.0; y++ {
		e.CreateGameObject("surface/grass", 1.0, y, 0.0, townId, nil)
		e.CreateGameObject("surface/grass", TownSize - 2.0, y, 0.0, townId, nil)
		gaMap.Cells <- &world_maps.WorldCell{X: 1.0, Y: y, SurfaceKind: "surface/grass"}
		gaMap.Cells <- &world_maps.WorldCell{X: TownSize - 2.0, Y: 0.0, SurfaceKind: "surface/grass"}
	}

	// Create town floor
	for x := WallSize; x < TownSize / 2.0; x++ {
		for y := WallSize; y < TownSize / 2.0; y++ {
			e.CreateGameObject("surface/town_floor", x, y, 0.0, townId, nil)
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: y, SurfaceKind: "surface/town_floor"}
		}
		for y := TownSize / 2.0 + 2.0; y < TownSize - WallSize; y++ {
			e.CreateGameObject("surface/town_floor", x, y, 0.0, townId, nil)
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: y, SurfaceKind: "surface/town_floor"}
		}
	}
	for x := TownSize / 2.0 + 2.0; x < TownSize - WallSize; x++ {
		for y := WallSize; y < TownSize / 2.0; y++ {
			if !(x >= FarmX && x < FarmX + FarmWidth && y >= FarmY && y < FarmY + FarmHeight) {
				e.CreateGameObject("surface/town_floor", x, y, 0.0, townId, nil)
				gaMap.Cells <- &world_maps.WorldCell{X: x, Y: y, SurfaceKind: "surface/town_floor"}
			}
		}
		for y := TownSize / 2.0 + 2.0; y < TownSize - WallSize; y++ {
			e.CreateGameObject("surface/town_floor", x, y, 0.0, townId, nil)
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: y, SurfaceKind: "surface/town_floor"}
		}
	}

	// Create columns
	e.CreateGameObject("wall/brick_column", WallSize, WallSize, 0.0, townId, nil)
	e.CreateGameObject("wall/brick_column", TownSize - WallSize, WallSize, 0.0, townId, nil)
	e.CreateGameObject("wall/brick_column", WallSize, TownSize - WallSize, 0.0, townId, nil)
	e.CreateGameObject("wall/brick_column", TownSize - WallSize, TownSize - WallSize, 0.0, townId, nil)

	// Create gates
	e.CreateGameObject("teleport/town_gate", TownSize / 2, 1.9, 0.0, townId, nil)
	e.CreateGameObject("teleport/town_gate", TownSize / 2, TownSize - 2.1, 0.0, townId, nil)
	e.CreateGameObject("teleport/town_gate", 2.2, TownSize / 2, math.Pi / 2, townId, nil)
	e.CreateGameObject("teleport/town_gate", TownSize - 1.8, TownSize / 2, math.Pi / 2, townId, nil)

	// Create buildings
	e.CreateGameObject("town/windmill", WindmillX, WindmillY, math.Pi, townId, nil)
	e.CreateGameObject("town/bell_tower", BellTowerX, BellTowerY, math.Pi, townId, nil)
	e.CreateGameObject("town/house", HouseX, HouseY, math.Pi, townId, nil)
	e.CreateGameObject("town/sawmill", SawmillX, SawmillY, 0.0, townId, nil)
	e.CreateGameObject("town/inn", InnX, InnY, math.Pi, townId, nil)
	e.CreateGameObject("town/blacksmith", BlacksmithX, BlacksmithY, math.Pi, townId, nil)
	e.CreateGameObject("equipment/anvil", BlacksmithX + 11.0, BlacksmithY + 2.0, 0.0, townId, nil)

	// Create Well area
	e.CreateGameObject("well/well", WellX, WellY, 0.0, townId, nil)
	e.CreateGameObject("bonfire/bonfire", BonfireX, BonfireY, 0.0, townId, nil)
	e.CreateGameObject("container/wooden_table", TablesX, TablesY, math.Pi / 2, townId, nil)
	e.CreateGameObject("sit/wooden_bench", TablesX, TablesY + 1.5, 0.0, townId, nil)
	e.CreateGameObject("sit/wooden_bench", TablesX, TablesY - 1.0, math.Pi, townId, nil)
	e.CreateGameObject("container/wooden_table", TablesX, TablesY - 5.5, math.Pi / 2, townId, nil)
	e.CreateGameObject("sit/wooden_bench", TablesX, TablesY - 4.0, 0.0, townId, nil)
	e.CreateGameObject("sit/wooden_bench", TablesX, TablesY - 6.5, math.Pi, townId, nil)

	// Create farm
	for x := FarmX; x < FarmX + FarmWidth; x++ {
		// Create fence
		if int64(x) % WallSize == 0.0 {
			e.CreateGameObject("wall/wooden_fence", x, FarmY, 0.0, townId, nil)
			e.CreateGameObject("wall/wooden_fence", x, FarmY + FarmHeight, 0.0, townId, nil)
		}
		for y := FarmY; y < FarmY + FarmHeight; y++ {
			e.CreateGameObject("surface/dirt", x, y, 0.0, townId, map[string]interface{}{
				"current_action": nil,
			})
			gaMap.Cells <- &world_maps.WorldCell{X: x, Y: y, SurfaceKind: "surface/dirt"}
			// Create plants
			if y < FarmY + FarmHeight / 2 {
				e.CreateGameObject("plant/carrot_ripe", x, y, 0.0, townId, nil)
			} else {
				e.CreateGameObject("plant/tomato_ripe", x, y, 0.0, townId, nil)
			}
			// Create fence
			if int64(y) % WallSize == 0.0 {
				e.CreateGameObject("wall/wooden_fence", FarmX, y, math.Pi / 2.0, townId, nil)
				e.CreateGameObject("wall/wooden_fence", FarmX + FarmWidth, y, math.Pi / 2.0, townId, nil)
			}
		}
	}

	// Create NPC
	e.CreateGameObject("npc/town_keeper", TownKeeperX, TownKeeperY, math.Pi * 2 - math.Pi / 4, townId, nil)
	e.CreateGameObject("container/market_stand", TownKeeperX - 4.0, TownKeeperY, math.Pi / 2, townId, nil)
	e.CreateGameObject("npc/dungeon_keeper", DungeonKeeperX, DungeonKeeperY, math.Pi * 2 - math.Pi / 2, townId, nil)
	e.CreateGameObject("town/trapdoor", DungeonKeeperX - 2.0, DungeonKeeperY, 0.0, townId, nil)


	// Create trees
	for i := 0; i < len(TreesX); i++ {
		e.CreateGameObject("tree/tree_5", TreesX[i], TreesY[i], math.Pi / 4 * float64(i), townId, nil)
	}

	gaMap.Finish <- true // save map to img
}
