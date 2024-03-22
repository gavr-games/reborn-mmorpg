package dragon_object_test


import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestResurrect(t *testing.T) {
	testFunction = callResurrect
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player := gameObjectFactory.CreatePlayer(game_test.GetEngine(), charGameObj)
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	slots["back"] = gameObjectFactory.CreateBackpackGameObject(game_test.GetEngine(), charGameObj).Id()
	charGameObj.SetProperty("slots", slots)
	dragonObj = gameObjectFactory.CreateObjectKeyXYFloor(game_test.GetEngine(), dragonKey, charGameObj.X(), charGameObj.Y(), charGameObj.Floor())

	t.Run("Not dragon owner", testNotOwner)

	dragonObj.SetProperty("owner_id", charGameObj.Id())

	t.Run("Dragon already alive", testAlreadyAlive)

	dragonObj.(entity.IDragonObject).Die()

	t.Run("Player does not have gold", testPlayerDoesNotHaveRequiredResources)

	// Give player enough resources
	resourceObj := gameObjectFactory.CreateStackableResourceGameObject(game_test.GetEngine(), charGameObj, "resource/gold", goldAmount)
	container, _ := game_test.GetEngine().GameObjects().Load(slots["back"].(string))
	container.(entity.IContainerObject).Put(game_test.GetEngine(), player, resourceObj.Id(), -1)

	t.Run("Player has no claim", testClaimNotExists)

	// Create claim
	claimObeliskObj = gameObjectFactory.CreateClaimObeliskObject(game_test.GetEngine(), charGameObj)
	claimObeliskObj.(entity.IClaimObeliskObject).Init(game_test.GetEngine())

	t.Run("PLayer has no dragon altar", testDragonAltarNotExists)
	
	// Create dragon altar
	altarObj := gameObjectFactory.CreateObjectKeyXYFloor(game_test.GetEngine(), altarKey, claimObeliskObj.X() + nearCoord, claimObeliskObj.Y(), claimObeliskObj.Floor())
	
	t.Run("Resurrected dragon", testSuccess)
	t.Run("Resurrected dragon on altar", func(t *testing.T) {
		assert.True(t, dragonObj.GetProperty("alive").(bool))
		assert.Equal(t, dragonObj.GetProperty("max_health"), dragonObj.GetProperty("health"))
		assert.Equal(t, dragonObj.X(), altarObj.X())
		assert.Equal(t, dragonObj.Y(), altarObj.Y())
		assert.Equal(t, dragonObj.Floor(), altarObj.Floor())
	})
}
