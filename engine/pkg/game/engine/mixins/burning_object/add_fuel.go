package burning_object

import (
	"errors"
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (obj *BurningObject) AddFuel(e entity.IEngine, itemId string, player *entity.Player) (bool, error) {
	var (
		charGameObj, item, container entity.IGameObject
		charOk, itemOk, contOk bool
	)
	gameObj := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false, errors.New("character not found")
	}

	if item, itemOk = e.GameObjects().Load(itemId); !itemOk {
		return false, errors.New("item not found")
	}

	// Check character is close enough
	if !gameObj.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer.", player)
		return false, errors.New("character needs to be closer")
	}

	// Check item is in allowed fuels
	allowedFuels := gameObj.GetProperty("allowed_fuels")
	fuelValue := 0.0
	itemKind := item.Kind()
	if allowedFuels == nil {
		e.SendSystemMessage("This fuel is not allowed.", player)
		return false, errors.New("this fuel is not allowed")
	}
	for fuelKind := range allowedFuels.(map[string]interface{}) {
    if fuelKind == itemKind {
			fuelValue = allowedFuels.(map[string]interface{})[fuelKind].(float64)
		}
	}
	if fuelValue == 0.0 {
		e.SendSystemMessage("This fuel is not allowed.", player)
		return false, errors.New("this fuel is not allowed")
	}

	// Check max fuel
	fuel := gameObj.GetProperty("fuel")
	maxFuel := gameObj.GetProperty("max_fuel")
	if fuel.(float64) == maxFuel.(float64) {
		e.SendSystemMessage("No space for more fuel.", player)
		return false, errors.New("no space for more fuel")
	}

	//check in container
	containerId := item.GetProperty("container_id")
	if (containerId == nil) {
		e.SendSystemMessage("First pickup item to feed it.", player)
		return false, errors.New("item is not in container")
	}

	// check container belongs to character
	if (containerId != nil) {
		if container, contOk = e.GameObjects().Load(containerId.(string)); !contOk {
			return false, errors.New("container not found")
		}
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false, errors.New("no access to container")
		}
		// remove from container if in container
		if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
			return false, errors.New("cannot remove item from container")
		}

		// Update fuel, check max
		fuel = fuel.(float64) + fuelValue
		if fuel.(float64) > maxFuel.(float64) {
			fuel = maxFuel
		}
		gameObj.SetProperty("fuel", fuel)
		
		e.SendGameObjectUpdate(gameObj, "update_object")
		e.RemoveGameObject(item)
		e.SendSystemMessage(fmt.Sprintf("You put %s to %s", item.Kind(), gameObj.Kind()), player)
	}

	return true, nil
}
