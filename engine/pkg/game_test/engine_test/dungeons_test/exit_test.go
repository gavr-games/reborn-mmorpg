package dungeons_test

import (
	"testing"
	"time"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/dungeons"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

const (
	Retries = 3
)

func TestExit(t *testing.T) {
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(e)
	player := gameObjectFactory.CreatePlayer(e, charGameObj)
	gameObjectFactory.CreateVisionArea(e, charGameObj)
	
	// Put char and dragon to the dungeon
	dungeonKeeper := gameObjectFactory.CreateObjectKeyXYArea(e, "npc/dungeon_keeper", charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId())
	dragon := gameObjectFactory.CreateObjectKeyXYArea(e, "mob/fire_dragon", charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	dragon.SetProperty("owner_id", charGameObj.Id())
	var dragonId interface{} = dragon.Id()

	res, err := dungeons.GoToDungeon(e, charGameObj, 1.0, []interface{}{dragonId})
	assert.True(t, res)
	assert.Nil(t, err)

	// Fail - not close to dungeon exit
	dungeonExit := getDungeonExit(e, charGameObj)

	res, err = dungeons.Exit(e, charGameObj, dungeonExit)
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player need to be closer to exit", err.Error())

	// Fail - no key
	charGameObj.SetX(dungeonExit.X())
	charGameObj.SetY(dungeonExit.Y())
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	backPack := gameObjectFactory.CreateBackpackGameObject(e, charGameObj)
	slots["back"] = backPack.Id()
	charGameObj.SetProperty("slots", slots)

	res, err = dungeons.Exit(e, charGameObj, dungeonExit)
	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, "Player does not have dungeon exit key", err.Error())

	// Success
	oldGameAreaId := dragon.GameAreaId()
	keyObj := gameObjectFactory.CreateObjectKeyXYArea(e, "resource/dungeon_key", charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	backPack.(entity.IContainerObject).Put(e, player, keyObj.Id(), -1)

	res, err = dungeons.Exit(e, charGameObj, dungeonExit)
	assert.True(t, res)
	assert.Nil(t, err)

	// Check gaiend exp
	charExp := charGameObj.GetProperty("experience").(float64)
	dragonExp := charGameObj.GetProperty("experience").(float64)
	assert.NotEqual(t, 0.0, charExp)
	assert.NotEqual(t, 0.0, dragonExp)

	// Check increased Max dungeon lvl
	maxLvl := charGameObj.GetProperty("max_dungeon_lvl").(float64)
	assert.Equal(t, 2.0, maxLvl)

	// Check removed key
	_, keyOk := e.GameObjects().Load(keyObj.Id())
	assert.False(t, keyOk)

	// Wait for goroutine to finish
	// TODO: maybe there is a better way? retry package?
	for i := 1; i <= Retries; i++ {
		if oldGameAreaId == dragon.GameAreaId() {
			time.Sleep(200 * time.Millisecond) 
		} else {
			break
		}
	}

	// Check teleported back to town
	assert.NotEqual(t, oldGameAreaId, dragon.GameAreaId())
	assert.NotEqual(t, oldGameAreaId, charGameObj.GameAreaId())

	// Check removed current dungeon
	dungeonId := charGameObj.GetProperty("current_dungeon_id")
	assert.Nil(t, dungeonId)

	// Cleanup
	e.RemoveGameObject(dungeonKeeper)
}