package dungeons

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	SizePerLevel = 10.0
	ChestX = 5.0
	ChestY = 5.0
	ExitX  = 8.0
	ExitY  = 8.0
	ZombieX = 1.0
	ZombieY = 8.0
)

func generate(e entity.IEngine, charGameObj entity.IGameObject, level float64, dragonIds []interface{}) *entity.GameArea {
	dungeon := entity.NewGameArea(
		fmt.Sprintf("dungeon_%s", charGameObj.Id()),
		0,
		0,
		level * SizePerLevel,
		level * SizePerLevel,
	)
	e.GameAreas().Store(dungeon.Id(), dungeon)
	storage.GetClient().GameAreasUpdates <- dungeon

	generateFloor(e, dungeon)
	generateChest(e, dungeon, ChestX, ChestY)
	generateExit(e, dungeon, charGameObj, ExitX, ExitY, level, dragonIds)
	e.CreateGameObject("mob/zombie", ZombieX, ZombieY, 0.0, dungeon.Id(), nil)

	return dungeon
}

func generateFloor(e entity.IEngine, dungeon *entity.GameArea) {
	for x := 0.; x < dungeon.Width(); x++ {
		for y := 0.; y < dungeon.Height(); y++ {
			e.CreateGameObject("surface/dungeon_floor", x, y, 0.0, dungeon.Id(), nil)
		}
	}
}

func generateChest(e entity.IEngine, dungeon *entity.GameArea, x, y float64) {
	chest := e.CreateGameObject("container/dungeon_chest", x, y, 0.0, dungeon.Id(), nil)
	key := e.CreateGameObject("resource/dungeon_key", x, y, 0.0, dungeon.Id(), nil)

	// Put key to the chest
	contItemsIds := chest.GetProperty("items_ids").([]interface{})
	contItemsIds[0] = key.Id()
	chest.SetProperty("items_ids", contItemsIds)
	chest.SetProperty("free_capacity", chest.GetProperty("free_capacity").(float64) - 1.0)
	key.SetProperty("container_id", chest.Id())

	storage.GetClient().Updates <- chest.Clone()
	storage.GetClient().Updates <- key.Clone()
}

func generateExit(e entity.IEngine, dungeon *entity.GameArea, charGameObj entity.IGameObject, x, y, level float64, dragonIds []interface{}) {
	e.CreateGameObject("wall/dungeon_exit", x, y, 0.0, dungeon.Id(), map[string]interface{}{
		"character_id": charGameObj.Id(),
		"level": level,
		"dragon_ids": dragonIds,
	})
}