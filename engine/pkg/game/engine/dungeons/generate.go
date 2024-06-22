package dungeons

import (
	"fmt"
	"math"

	"pgregory.net/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	MinSize = 30.0
	SizePerLevel = 5.0
	MinSplitCount = 2
	IncSplitPerLevel = 3
	MinRoomSize = 6.0
	MaxRoomSize = 12.0
	MinMobsPerRoom = 1
	MaxMobsPerRoom = 3
	MobHealthPerLevel = 5.0
	MobDamagePerLevel = 1.0
	CorridorSize = 2.0
	CharPosMargin = 1.0
	ExitMargin = 2.0
	WallMargin = 0.2
	DoorMargin = 0.6
	ChestMargin = 2.5
	ColumnMargin = 0.2
)

func generate(e entity.IEngine, charGameObj entity.IGameObject, level float64, dragonIds []interface{}) (*entity.GameArea, float64, float64) {
	width := MinSize + level * (SizePerLevel - 1)
	height := MinSize + level * (SizePerLevel - 1)
	dungeon := entity.NewGameArea(
		fmt.Sprintf("dungeon_%s", charGameObj.Id()),
		0,
		0,
		width,
		height,
	)
	dungeonId := dungeon.Id()
	e.GameAreas().Store(dungeonId, dungeon)
	storage.GetClient().GameAreasUpdates <- dungeon

	// Generate rooms
	splitsCount := MinSplitCount + math.Floor(level / IncSplitPerLevel)
	rootCont := BSPGenerate(width, height, MinRoomSize, MaxRoomSize, int(splitsCount))
	generateRooms(e, rootCont, level, dungeonId)

	// Generate chest, exit, char pos
	rooms := rootCont.GetThreeMostDistantRooms()
	generateChest(e, rooms[1], dungeonId)
	generateExit(e, rooms[2], dungeonId, charGameObj, level, dragonIds)
	charX := rooms[0].X + CharPosMargin
	charY := rooms[0].Y + CharPosMargin

	// Genrate floor and walls
	generateFloorAndWalls(e, rootCont, dungeonId, int(width), int(height))

	return dungeon, charX, charY
}

func generateRooms(e entity.IEngine, cont *Container, level float64, dungeonId string) {
	if cont.Child1 == nil && cont.Child2 == nil {
		minWidth := math.Max(cont.Width / 2.0, MinRoomSize)
		maxWidth := cont.Width
		minHeight := math.Max(cont.Height / 2.0, MinRoomSize)
		maxHeight := cont.Height
		width := minWidth
		if maxWidth - minWidth > 0 {
			width = float64(rand.Intn(int(maxWidth - minWidth))) + minWidth
		}
		width = width - float64(int(width) % 2)
		height := minHeight
		if maxHeight - minHeight > 0 {
			height = float64(rand.Intn(int(maxHeight - minHeight))) + minHeight
		}
		height = height - float64(int(height) % 2)
		minX := cont.X
		maxX := cont.X + cont.Width - width
		x := minX
		if maxX - minX > 0 {
			x = float64(rand.Intn(int(maxX - minX))) + minX
		}
		x = x - float64(int(x) % 2)
		minY := cont.Y
		maxY := cont.Y + cont.Height - height
		y := minY
		if maxY - minY > 0 {
			y = float64(rand.Intn(int(maxY - minY))) + minY
		}
		y = y - float64(int(y) % 2)
		cont.Room = &Container{x, y, width, height, nil, nil, nil}

		// Generate Mobs
		mobsCount := rand.Intn(MaxMobsPerRoom - MinMobsPerRoom) + MinMobsPerRoom
		for i := 0; i < mobsCount; i++ {
			mobX := float64(rand.Intn(int(width))) + x
			mobY := float64(rand.Intn(int(height))) + y
			mob := e.CreateGameObject("mob/zombie", mobX, mobY, 0.0, dungeonId, nil)
			health := mob.GetProperty("health").(float64)
			damage := mob.GetProperty("damage").(float64)
			mob.SetProperty("health", health + MobHealthPerLevel * level)
			mob.SetProperty("max_health", health + MobHealthPerLevel * level)
			mob.SetProperty("damage", damage + MobDamagePerLevel * level)
		}
	} else {
		generateRooms(e, cont.Child1, level, dungeonId)
		generateRooms(e, cont.Child2, level, dungeonId)
	}
}

func generateFloorAndWalls(e entity.IEngine, cont *Container, dungeonId string, width, height int) {
	allRooms := []*Container{}
	allRooms = cont.GetAllRooms(allRooms)
	allCorridors := []*Container{}
	allCorridors = cont.GetAllCorridors(allCorridors, CorridorSize)

	// initializing dungeon map
	var dungeonMap [][]byte
	dungeonMap = make([][]byte, height)
	for i := range dungeonMap {
		dungeonMap[i] = make([]byte, width)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dungeonMap[x][y] = ' '
		}
	}

	// Fill rooms
	for i := range allRooms {
		for x := int(allRooms[i].X); x < int(allRooms[i].X + allRooms[i].Width); x++ {
			for y := int(allRooms[i].Y); y < int(allRooms[i].Y + allRooms[i].Height); y++ {
				dungeonMap[x][y] = 'x'
			}
		}
	}
	// Fill corridors
	for i := range allCorridors {
		for x := int(allCorridors[i].X); x < int(allCorridors[i].X + allCorridors[i].Width); x++ {
			for y := int(allCorridors[i].Y); y < int(allCorridors[i].Y + allCorridors[i].Height); y++ {
				dungeonMap[x][y] = 'x'
			}
		}
	}

	// Create floor
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if dungeonMap[x][y] == 'x' {
				e.CreateGameObject("surface/dungeon_floor", float64(x), float64(y), 0.0, dungeonId, nil)
			}
		}
	}

	// Create Walls
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if dungeonMap[x][y] == 'x' {
				// Check neighbour cells are empty and we need to place wall
				// Check north cell
				if y - 1 < 0 || dungeonMap[x][y - 1] == ' ' {
					e.CreateGameObject("wall/dungeon_wall", float64(x), float64(y) - WallMargin, 0.0, dungeonId, nil)
					// Check corners for column
					if x - 1 < 0 || dungeonMap[x - 1][y] == ' ' {
						e.CreateGameObject("wall/dungeon_column", float64(x) - ColumnMargin, float64(y) - ColumnMargin, 0.0, dungeonId, nil)
					}
					if x + 2 >= width || dungeonMap[x + 2][y] == ' ' {
						e.CreateGameObject("wall/dungeon_column", float64(x + 2) - ColumnMargin, float64(y) - ColumnMargin, 0.0, dungeonId, nil)
					}
				}
				// Check south cell
				if y + 1 >= height || dungeonMap[x][y + 1] == ' ' {
					e.CreateGameObject("wall/dungeon_wall", float64(x), float64(y + 1), 0.0, dungeonId, nil)
					// Check corners for column
					if x - 1 < 0 || dungeonMap[x - 1][y] == ' ' {
						e.CreateGameObject("wall/dungeon_column", float64(x) - ColumnMargin, float64(y + 1), 0.0, dungeonId, nil)
					}
					if x + 2 >= width || dungeonMap[x + 2][y] == ' ' {
						e.CreateGameObject("wall/dungeon_column", float64(x + 2) - ColumnMargin, float64(y + 1), 0.0, dungeonId, nil)
					}
				}
				x++
			}
		}
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if dungeonMap[x][y] == 'x' {
				// Check neighbour cells are empty and we need to place wall
				// Check west cell
				if x - 1 < 0 || dungeonMap[x - 1][y] == ' ' {
					e.CreateGameObject("wall/dungeon_wall", float64(x) - WallMargin, float64(y), math.Pi / 2.0, dungeonId, nil)
				}
				// Check east cell
				if x + 1 >= width || dungeonMap[x + 1][y] == ' ' {
					e.CreateGameObject("wall/dungeon_wall", float64(x + 1), float64(y), math.Pi / 2.0, dungeonId, nil)
				}
				y++
			}
		}
	}
	// Create columns, which left
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if dungeonMap[x][y] == ' ' {
				// o__
				// |
				if x - 1 >= 0 && dungeonMap[x - 1][y] == 'x' && y - 1 >= 0 && dungeonMap[x][y - 1] == 'x' {
					e.CreateGameObject("wall/dungeon_column", float64(x), float64(y), 0.0, dungeonId, nil)
				}
				// __o
				//   |
				if x + 1 < width && dungeonMap[x + 1][y] == 'x' && y - 1 >= 0 && dungeonMap[x][y - 1] == 'x' {
					e.CreateGameObject("wall/dungeon_column", float64(x + 1), float64(y), 0.0, dungeonId, nil)
				}
				// |
				// o--
				if x - 1 >= 0 && dungeonMap[x - 1][y] == 'x' && y + 1 < height && dungeonMap[x][y + 1] == 'x' {
					e.CreateGameObject("wall/dungeon_column", float64(x), float64(y + 1), 0.0, dungeonId, nil)
				}
				//   |
				// --o
				if x + 1 < width && dungeonMap[x + 1][y] == 'x' && y + 1 < height && dungeonMap[x][y + 1] == 'x' {
					e.CreateGameObject("wall/dungeon_column", float64(x + 1), float64(y + 1), 0.0, dungeonId, nil)
				}
			}
		}
	}
}

func generateChest(e entity.IEngine, room *Container, dungeonId string) {
	x := room.X + ChestMargin
	y := room.Y + room.Height - ChestMargin - 2.0
	chest := e.CreateGameObject("container/dungeon_chest", x, y, 0.0, dungeonId, nil)
	key := e.CreateGameObject("resource/dungeon_key", x, y, 0.0, dungeonId, nil)

	// Put key to the chest
	contItemsIds := chest.GetProperty("items_ids").([]interface{})
	contItemsIds[0] = key.Id()
	chest.SetProperty("items_ids", contItemsIds)
	chest.SetProperty("free_capacity", chest.GetProperty("free_capacity").(float64) - 1.0)
	key.SetProperty("container_id", chest.Id())

	storage.GetClient().Updates <- chest.Clone()
	storage.GetClient().Updates <- key.Clone()
}

func generateExit(e entity.IEngine, room *Container, dungeonId string, charGameObj entity.IGameObject, level float64, dragonIds []interface{}) {
	x := room.X + ExitMargin
	y := room.Y + room.Height - DoorMargin
	e.CreateGameObject("wall/dungeon_exit", x, y, 0.0, dungeonId, map[string]interface{}{
		"character_id": charGameObj.Id(),
		"level": level,
		"dragon_ids": dragonIds,
	})
}