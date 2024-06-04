package fishing_rod_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestCatch(t *testing.T) {
	testFunction = callCatch
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player := gameObjectFactory.CreatePlayer(e, charGameObj)
	fishingRodObject = gameObjectFactory.CreateObjectKeyXYArea(e, FishingRodObjectKey, charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId())
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	backPack := gameObjectFactory.CreateBackpackGameObject(e, charGameObj)
	slots["back"] = backPack.Id()
	charGameObj.SetProperty("slots", slots)

	t.Run("Test fishing rod is not equipped", testNotEquipped)

	backPack.(entity.IContainerObject).Put(e, player, fishingRodObject.Id(), -1)
	fishingRodObject.(entity.IEquipableObject).Equip(e, player)
	t.Run("Test not near water", testNotNearWater)

	gameObjectFactory.CreateObjectKeyXYArea(e, "surface/water", charGameObj.X() + CloseCoord, charGameObj.Y() + CloseCoord, charGameObj.GameAreaId())
	t.Run("Caught fish", testSuccess)

	t.Run("Caught fish", func(t *testing.T) {
		hasFish := backPack.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
			"fish": 1.0,
		})
		assert.True(t, hasFish)
	})
}
