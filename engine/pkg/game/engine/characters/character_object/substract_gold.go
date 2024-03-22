package character_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (charGameObj *CharacterObject) SubstractGold(e entity.IEngine, amount float64) (bool, error) {
	playerId := charGameObj.GetProperty("player_id")
	if playerId == nil {
		return false, errors.New("Player does not exist")
	}

	player, pOk := e.Players().Load(playerId.(int))
	if !pOk {
		return false, errors.New("Player does not exist")
	}

	resourceKey := "gold"
	resourceAmount := amount
	slots := charGameObj.GetProperty("slots").(map[string]interface{})
	var (
		container entity.IGameObject
		contOk bool
	)
	if slots["back"] == nil {
		return false, errors.New("Player does not have container")
	}
	if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
		return false, errors.New("Player does not have container")
	}
	// check container has items
	if !container.(entity.IContainerObject).HasItemsKinds(e, map[string]interface{}{
		(resourceKey): resourceAmount,
	}) {
		e.SendSystemMessage("You don't have required resources.", player)
		return false, errors.New("Player does not have required resources")
	}

	// substract resources/money
	if !container.(entity.IContainerObject).RemoveItemsKinds(e, player, map[string]interface{}{
		(resourceKey): resourceAmount,
	}) {
		e.SendSystemMessage("Can't remove required resources.", player)
		return false, errors.New("Player can not remove required resources")
	}

	return true, nil
}
