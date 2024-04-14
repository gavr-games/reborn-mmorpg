package claim_obelisk_object_test


import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestFindKindInArea(t *testing.T) {
	testFunction = callFindKindInArea
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	claimObeliskObj = gameObjectFactory.CreateClaimObeliskObject(game_test.GetEngine(), charGameObj)

	t.Run("Claim area does not exist", testClaimAreaDoesNotExist)

	claimObeliskObj.(entity.IClaimObeliskObject).Init(game_test.GetEngine())
	kind = objectKind
	t.Run("Test object with kind not found", testObjectNotFound)

	// Create object far away
	gameObjectFactory.CreateObjectKeyXYArea(game_test.GetEngine(), objectKey, claimObeliskObj.X() + awayCoord, claimObeliskObj.Y(), claimObeliskObj.GameAreaId())
	t.Run("Test object with kind not found", testObjectNotFound)

	// Create object in claim area
	findObj := gameObjectFactory.CreateObjectKeyXYArea(game_test.GetEngine(), objectKey, claimObeliskObj.X() + nearCoord, claimObeliskObj.Y(), claimObeliskObj.GameAreaId())
	t.Run("Found object with kind", testSuccess)
	t.Run("Found right object with kind", func(t *testing.T) {
		assert.Equal(t, findObj.Id(), res.(entity.IGameObject).Id())
	})
}
