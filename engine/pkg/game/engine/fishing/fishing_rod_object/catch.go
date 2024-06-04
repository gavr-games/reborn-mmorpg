package fishing_rod_object

import (
	"errors"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (fishingRod *FishingRodObject) Catch(e entity.IEngine, charGameObj entity.IGameObject) (bool, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false, errors.New("no container found")
		}

		if ok, err := fishingRod.CheckCatch(e, charGameObj); !ok {
			return false, err
		}

		// Create fish resource
		// TODO: create fish atlas for different fishes
		resourceKey := "fish"
		resourceObj := e.CreateGameObject(fmt.Sprintf("resource/%s", resourceKey), charGameObj.X(), charGameObj.Y(), 0.0, "", nil)

		// put resource to container or drop it to the ground
		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			return false, errors.New("no container found")
		}
		container.(entity.IContainerObject).PutOrDrop(e, charGameObj, resourceObj.Id(), -1)

		charGameObj.(entity.ILevelingObject).AddExperience(e, "fishing")

		e.SendSystemMessage(fmt.Sprintf("You received a %s.", resourceObj.Kind()), player)
		return true, nil
	} else {
		return false, errors.New("player not found")
	}
}
