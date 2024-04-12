package mob_object_test


import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs/mob_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
)

func TestResurrect(t *testing.T) {
	gameObjectFactory := factories.NewGameObjectFactory()
	targetObj = gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	mob = gameObjectFactory.CreateObjectKeyXYFloor(game_test.GetEngine(), mobKey, targetObj.X() + noAgroDistance, targetObj.Y(), targetObj.Floor()).(*mob_object.MobObject)

	t.Run("No target to agro", testNoCheckAgroTarget)

	targetObj.SetX(targetObj.X() + noAgroDistance)

	t.Run("Found target to agro", testCheckAgroTarget)
}
