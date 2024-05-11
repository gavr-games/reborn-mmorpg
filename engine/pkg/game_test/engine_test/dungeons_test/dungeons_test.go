package dungeons_test

import (
	"os"
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func TestMain(m *testing.M) {
	game_test.Setup()
	os.Exit(m.Run())
}

func getDungeonExit(e entity.IEngine, charGameObj entity.IGameObject) entity.IGameObject {
	dungeonId := charGameObj.GetProperty("current_dungeon_id")
	dungeon, _ := e.GameAreas().Load(dungeonId.(string))
	gameObjects := dungeon.GetAllGameObjects()

	for _, val := range gameObjects {
		gameObj := val.(entity.IGameObject)
		if gameObj.Kind() == "dungeon_exit" {
			return gameObj
		}
	}

	return nil
}