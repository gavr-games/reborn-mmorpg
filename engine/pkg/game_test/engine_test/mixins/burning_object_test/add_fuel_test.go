package burning_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestAddFuel(t *testing.T) {
	testFunction = callAddFuel
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player = gameObjectFactory.CreatePlayer(e, charGameObj)
	burningObject = gameObjectFactory.CreateObjectKeyXYArea(e, BurningObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())
	notFuelObj := gameObjectFactory.CreateObjectKeyXYArea(e, NotFuelObjectKey, charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	itemId = notFuelObj.Id()

	t.Run("Character needs to be closer", testNotClose)

	charGameObj.SetX(burningObject.X())
	charGameObj.SetY(burningObject.Y())
	t.Run("Not fuel item", testNotFuel)

	fuel := gameObjectFactory.CreateObjectKeyXYArea(e, FuelObjectKey, charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	itemId = fuel.Id()
	t.Run("Fuel not in container", testNotInContainer)

	initialFuel := burningObject.GetProperty("fuel")
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	backPack := gameObjectFactory.CreateBackpackGameObject(e, charGameObj)
	slots["back"] = backPack.Id()
	charGameObj.SetProperty("slots", slots)
	backPack.(entity.IContainerObject).Put(e, player, fuel.Id(), -1)
	burningObject.SetProperty("fuel", burningObject.GetProperty("max_fuel"))
	t.Run("Fuel is full", testFuelIsFull)

	burningObject.SetProperty("fuel", initialFuel.(float64))
	t.Run("Add fuel", testSuccess)

	t.Run("Fuel has changed", func(t *testing.T) {
		fuel := burningObject.GetProperty("fuel")
		assert.Equal(t, FuelValue, fuel.(float64) - initialFuel.(float64))
	})
}
