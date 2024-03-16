package npc_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestSellItem(t *testing.T) {
	testFunction = callSellItem
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(e)
	npcObj = gameObjectFactory.CreateNpcGameObject(e)

	t.Run("Player does not exist", testPlayerDoesNotExist)

	// Create a new player
	player := gameObjectFactory.CreatePlayer(e, charGameObj)

	t.Run("Player does not have container", testPlayerDoesNotHaveContainer)

	// Create a container for player
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	slots["back"] = gameObjectFactory.CreateBackpackGameObject(e, charGameObj).Id()
	charGameObj.SetProperty("slots", slots)

	// Place player far from NPC
	charGameObj.SetX(2.0)
	charGameObj.SetY(2.0)

	t.Run("Player need to be closer to NPC", testPlayerNeedToBeCloser)

	// Place player closer to NPC
	charGameObj.SetX(0)
	charGameObj.SetY(0)

	t.Run("Player does not have required resources", testPlayerDoesNotHaveRequiredResources)

	// Give player enough resources
	itemToSellObj := gameObjectFactory.CreateResourceGameObject(e, charGameObj, itemToSellObjKey)
	container, _ := e.GameObjects().Load(slots["back"].(string))
	container.(entity.IContainerObject).Put(e, player, itemToSellObj.Id(), -1)

	t.Run("Player sold item successfully", testSuccess)
	t.Run("Player got rid of item", func(t *testing.T) {
		hasItem := container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
			itemToSellKind: amount,
		})
		assert.False(t, hasItem)
	})
	t.Run("Player received resources", func(t *testing.T) {
		hasItem := container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
			resourceKind: sellingPrice,
		})
		assert.True(t, hasItem)
	})
}
