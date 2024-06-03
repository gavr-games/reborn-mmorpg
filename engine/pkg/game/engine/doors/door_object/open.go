package door_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (door *DoorObject) Open(e entity.IEngine, player *entity.Player) (bool, error) {
	var (
		charGameObj entity.IGameObject
		charOk bool
	)
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false, errors.New("character not found")
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, door) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false, errors.New("no access")
	}

	// Check near building
	if !door.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer to the claim.", player)
		return false, errors.New("character needs to be closer")
	}

	// Open
	door.SetProperty("state", "opened")
	door.SetProperty("collidable", false)
	e.SendGameObjectUpdate(door, "update_object")

	return true, nil
}
