package feedable_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestFeed(t *testing.T) {
	testFunction = callFeed
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player = gameObjectFactory.CreatePlayer(e, charGameObj)
	feedableObject = gameObjectFactory.CreateObjectKeyXYArea(e, FeedableObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())
	notEatableObj := gameObjectFactory.CreateObjectKeyXYArea(e, NotEatableObjectKey, charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	itemId = notEatableObj.Id()

	t.Run("Character needs to be closer", testNotClose)

	charGameObj.SetX(feedableObject.X())
	charGameObj.SetY(feedableObject.Y())
	t.Run("Not eatable item", testNotEatable)

	foodObj := gameObjectFactory.CreateObjectKeyXYArea(e, EatableObjectKey, charGameObj.X(), charGameObj.Y(), charGameObj.GameAreaId()) 
	itemId = foodObj.Id()
	foodObjFullness := foodObj.GetProperty("fullness")
	t.Run("Food not in container", testNotInContainer)

	initialFoodToEvolve := feedableObject.GetProperty("food_to_evolve")
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	backPack := gameObjectFactory.CreateBackpackGameObject(e, charGameObj)
	slots["back"] = backPack.Id()
	charGameObj.SetProperty("slots", slots)
	backPack.(entity.IContainerObject).Put(e, player, foodObj.Id(), -1)
	t.Run("Consume food", testSuccess)

	t.Run("Fullness has changed", func(t *testing.T) {
		fullness := feedableObject.GetProperty("fullness")
		foodToEvolve := feedableObject.GetProperty("food_to_evolve")
		assert.Equal(t, foodObjFullness, fullness)
		assert.Equal(t, initialFoodToEvolve.(float64) - foodToEvolve.(float64), foodObjFullness)
	})
}
