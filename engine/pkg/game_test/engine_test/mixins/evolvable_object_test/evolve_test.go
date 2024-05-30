package evolvable_object_test

import (
	"testing"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test"
	"github.com/gavr-games/reborn-mmorpg/pkg/game_test/factories"
	"github.com/stretchr/testify/assert"
)

func TestEvolve(t *testing.T) {
	testFunction = callEvolve
	e := game_test.GetEngine()
	gameObjectFactory := factories.NewGameObjectFactory()
	charGameObj := gameObjectFactory.CreateCharGameObject(game_test.GetEngine())
	player = gameObjectFactory.CreatePlayer(e, charGameObj)
	evolvableObject = gameObjectFactory.CreateObjectKeyXYArea(e, EvolvableObjectKey, charGameObj.X() + AwayCoord, charGameObj.Y(), charGameObj.GameAreaId())

	t.Run("Character needs to be closer", testNotClose)

	charGameObj.SetX(evolvableObject.X())
	charGameObj.SetY(evolvableObject.Y())
	t.Run("Not enough food to evolve", testNotEnoughFood)

	evolvableObject.SetProperty("alive", false)
	evolvableObject.SetProperty("food_to_evolve", 0.0)
	t.Run("Evolvable object is dead", testAlive)

	evolvableObject.SetProperty("alive", true)
	t.Run("Evolve", testSuccess)

	t.Run("Object evolved", func(t *testing.T) {
		obj, objOk := e.GameObjects().Load(evolvableObject.Id())
		area, _ := e.GameAreas().Load(evolvableObject.GameAreaId())
		gameObjects := area.GetAllGameObjects()
		assert.Equal(t, EvolveToKind, gameObjects[len(gameObjects)-1].(entity.IGameObject).Kind())
		assert.False(t, objOk)
		assert.Nil(t, obj)
	})
}
