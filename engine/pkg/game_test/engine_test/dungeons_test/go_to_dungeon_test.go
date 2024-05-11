package dungeons_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/dungeons"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

const (
	awayCoord  = 30.0
)

func TestGoToDungeon(t *testing.T) {
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(e)
	gameObjectFactory.CreatePlayer(e, charGameObj)
	gameObjectFactory.CreateVisionArea(e, charGameObj)

	// Fail if not close to dungeon keeper
	dungeonKeeper := gameObjectFactory.CreateObjectKeyXYArea(e, "npc/dungeon_keeper", charGameObj.X() + awayCoord, charGameObj.Y(), charGameObj.GameAreaId())
	
	res, err := dungeons.GoToDungeon(e, charGameObj, 1.0, []interface{}{})
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Character is not near the Dungeon Keeper", err.Error())

	dungeonKeeper.SetX(charGameObj.X())

	// Fail if level is invalid
	res, err = dungeons.GoToDungeon(e, charGameObj, 2.0, []interface{}{})
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid dungeon level", err.Error())

	// Fail if dragon does not belong to character
	dragon := gameObjectFactory.CreateObjectKeyXYArea(e, "mob/fire_dragon", charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	var dragonId interface{} = dragon.Id()

	res, err = dungeons.GoToDungeon(e, charGameObj, 1.0, []interface{}{dragonId})
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid dragon selected", err.Error())

	// Fail if dragon does not belong to character
	dragon.SetProperty("owner_id", charGameObj.Id())
	dragon.SetProperty("alive", false)

	res, err = dungeons.GoToDungeon(e, charGameObj, 1.0, []interface{}{dragonId})
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Dead dragon selected", err.Error())

	// Fail too many dragons
	res, err = dungeons.GoToDungeon(e, charGameObj, 1.0, []interface{}{dragonId, dragonId, dragonId, dragonId})
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Too many dragons selected", err.Error())

	// Success
	oldGameAreaId := dragon.GameAreaId()
	dragon.SetProperty("alive", true)

	res, err = dungeons.GoToDungeon(e, charGameObj, 1.0, []interface{}{dragonId})
	assert.True(t, res)
	assert.Nil(t, err)

	// Check char and dragon are teleported
	assert.NotEqual(t, oldGameAreaId, dragon.GameAreaId())
	assert.NotEqual(t, oldGameAreaId, charGameObj.GameAreaId())

	// Check dungeon id is set
	dungeonId := charGameObj.GetProperty("current_dungeon_id")
	assert.NotNil(t, dungeonId)

	// Check dungeon exists
	dungeon, dOk := e.GameAreas().Load(dungeonId.(string))
	assert.True(t, dOk)

	// Check dungeon has exit
	hasExit := false
	hasKey := false
	gameObjects := dungeon.GetAllGameObjects()
	for _, val := range gameObjects {
		gameObj := val.(entity.IGameObject)
		if gameObj.Kind() == "dungeon_exit" {
			hasExit = true
		} else if gameObj.Kind() == "dungeon_key" {
			hasKey = true
		}
		if hasExit && hasKey {
			break
		}
	}
	assert.True(t, hasExit)
	assert.True(t, hasKey)

	// Cleanup
	e.RemoveGameObject(dungeonKeeper)
}
