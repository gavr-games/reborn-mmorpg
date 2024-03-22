package character_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestSubstractGold(t *testing.T) {
	testFunction = callSubstractGold
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	amount = substractGoldAmount

	t.Run("Player does not exist", testPlayerDoesNotExist)

	// Create a new player
	player := gameObjectFactory.CreatePlayer(game_test.GetEngine(), charGameObj)

	t.Run("Player does not have container", testPlayerDoesNotHaveContainer)

	// Create a container for player
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	slots["back"] = gameObjectFactory.CreateBackpackGameObject(game_test.GetEngine(), charGameObj).Id()
	charGameObj.SetProperty("slots", slots)

	t.Run("Player does not have required resources", testPlayerDoesNotHaveRequiredResources)

	// Give player enough resources
	resourceObj := gameObjectFactory.CreateStackableResourceGameObject(game_test.GetEngine(), charGameObj, "resource/gold", totalGoldAmount)
	container, _ := game_test.GetEngine().GameObjects().Load(slots["back"].(string))
	container.(entity.IContainerObject).Put(game_test.GetEngine(), player, resourceObj.Id(), -1)
	t.Run("Substracted gold", testSuccess)
	t.Run("Substracted gold amount", func(t *testing.T) {
		amountLeft := resourceObj.GetProperty("amount").(float64)
		assert.Equal(t, totalGoldAmount - substractGoldAmount, amountLeft)
	})

	// Try to substract more gold, than player has
	amount = totalGoldAmount
	t.Run("Player does not have required resources", testPlayerDoesNotHaveRequiredResources)

	// Substract last gold
	amount = substractGoldAmount
	t.Run("Substracted gold", testSuccess)
	t.Run("Substracted last gold amount", func(t *testing.T) {
		_, resOk := game_test.GetEngine().GameObjects().Load(resourceObj.Id())
		assert.False(t, resOk)
	})
}
