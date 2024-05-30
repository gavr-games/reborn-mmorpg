package feedable_object

import (
	"errors"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (obj *FeedableObject) Feed(e entity.IEngine, itemId string, player *entity.Player) (bool, error) {
	var (
		charGameObj, item, container entity.IGameObject
		charOk, itemOk, contOk bool
	)
	gameObj := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false, errors.New("Character not found")
	}

	if item, itemOk = e.GameObjects().Load(itemId); !itemOk {
		return false, errors.New("Item not found")
	}

	// Check character is close enough
	if !gameObj.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer.", player)
		return false, errors.New("Character needs to be closer")
	}

	// Check item is eatable
	if eatable := item.GetProperty("eatable"); eatable == nil || !eatable.(bool) {
		e.SendSystemMessage("You cannot feed this item.", player)
		return false, errors.New("Wrong item to feed")
	}

	// Check fullness
	fullness := gameObj.GetProperty("fullness")
	maxFullness := gameObj.GetProperty("max_fullness")
	if fullness.(float64) == maxFullness.(float64) {
		e.SendSystemMessage("No space for more food.", player)
		return false, errors.New("No space for more food")
	}

	//check in container
	containerId := item.GetProperty("container_id")
	if (containerId == nil) {
		e.SendSystemMessage("First pickup item to feed it.", player)
		return false, errors.New("Item is not in container")
	}

	// check container belongs to character
	if (containerId != nil) {
		if container, contOk = e.GameObjects().Load(containerId.(string)); !contOk {
			return false, errors.New("Container not found")
		}
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false, errors.New("No access to container")
		}
		// remove from container if in container
		if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
			return false, errors.New("Cannot remove item from container")
		}

		// Update fullness, check max
		fullness = fullness.(float64) + item.GetProperty("fullness").(float64)
		if fullness.(float64) > maxFullness.(float64) {
			fullness = maxFullness
		}
		gameObj.SetProperty("fullness", fullness)

		// Update food_to_evolve
		foodToEvolve := gameObj.GetProperty("food_to_evolve")
		if foodToEvolve != nil && foodToEvolve != 0.0 {
			foodToEvolve = foodToEvolve.(float64) - item.GetProperty("fullness").(float64)
			if foodToEvolve.(float64) < 0.0 {
				foodToEvolve = 0.0
			}
			gameObj.SetProperty("food_to_evolve", foodToEvolve)
		}
		
		e.SendGameObjectUpdate(gameObj, "update_object")
		e.RemoveGameObject(item)
		e.SendSystemMessage(fmt.Sprintf("You gave %s some %s", gameObj.Kind(), item.Kind()), player)
	}

	return true, nil
}
