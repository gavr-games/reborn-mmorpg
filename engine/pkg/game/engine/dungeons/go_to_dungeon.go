package dungeons

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	KeeperDistance = 1.0
	MaxDragons = 3
)

// Create dungeon and teleport character with dragons to it
func GoToDungeon(e entity.IEngine, charGameObj entity.IGameObject, level float64, dragonIds []interface{}) (bool, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		gameArea, gaOk := e.GameAreas().Load(charGameObj.GameAreaId())
		if !gaOk {
			return false, errors.New("Game Area does not exist")
		}

		// Is standing near dungeon keeper
		possibleKeeperObjects := gameArea.RetrieveIntersections(utils.Bounds{
			X:      charGameObj.X() - KeeperDistance,
			Y:      charGameObj.Y() - KeeperDistance,
			Width:  charGameObj.Width() + KeeperDistance * 2,
			Height: charGameObj.Height() + KeeperDistance * 2,
		})

		foundKeeper := false
		for _, val := range possibleKeeperObjects {
			gameObj := val.(entity.IGameObject)
			if gameObj.Kind() == "dungeon_keeper" {
				foundKeeper = true
				break
			}
		}
		if !foundKeeper {
			e.SendSystemMessage("You need to stand near the Dungeon Keeper.", player)
			return false, errors.New("Character is not near the Dungeon Keeper")
		}

		// Validate level
		maxLevel := charGameObj.GetProperty("max_dungeon_lvl")
		if maxLevel == nil || maxLevel.(float64) < level {
			e.SendSystemMessage("Invalid dungeon level.", player)
			return false, errors.New("Invalid dungeon level")
		}

		// Validate only 3 dragons can be telported
		if len(dragonIds) > MaxDragons {
			e.SendSystemMessage("Too many dragons selected.", player)
			return false, errors.New("Too many dragons selected")
		}

		// Validate dragons exist and alive and belong to this char
		for _, id := range dragonIds {
			dragonId := id.(string)
			if dragon, dOk := e.GameObjects().Load(dragonId); dOk {
				if alive := dragon.GetProperty("alive"); alive == nil || !alive.(bool) {
					e.SendSystemMessage("Dead dragon selected, ressurect it first.", player)
					return false, errors.New("Dead dragon selected")
				}
				ownerId := dragon.GetProperty("owner_id")
				if ownerId == nil || charGameObj.Id() != ownerId.(string) {
					e.SendSystemMessage("Invalid dragon selected.", player)
					return false, errors.New("Invalid dragon selected")
				}
			} else {
				e.SendSystemMessage("Invalid dragon selected.", player)
				return false, errors.New("Invalid dragon selected")
			}
		}

		e.SendSystemMessage("You will be teleported to dungeon soon. Please wait.", player)
		generateAndTeleport(e, charGameObj, level, dragonIds)

		return true, nil
	}

	return false, errors.New("Player does not exist")
}

func generateAndTeleport(e entity.IEngine, charGameObj entity.IGameObject, level float64, dragonIds []interface{}) {
	// Generate dungeon
	dungeon := generate(e, charGameObj, level, dragonIds)
	
	e.PerformTask(func() { // send this code to engine main loop
		// Teleport to dungeon
		charGameObj.(entity.ICharacterObject).DeselectTarget(e)
		charGameObjClone := charGameObj.Clone()
		e.SendResponseToVisionAreas(charGameObjClone, "remove_object", map[string]interface{}{
			"object": charGameObjClone,
		})
		charGameObj.(entity.ICharacterObject).Move(e, 1.0, 1.0, dungeon.Id())

		// Teleport dragons to dungeon
		for _, id := range dragonIds {
			dragonId := id.(string)
			if dragon, dOk := e.GameObjects().Load(dragonId); dOk {
				dragon.(entity.IDragonObject).TeleportToOwner(charGameObj)
			}
		}

		// Set current dungeon
		charGameObj.SetProperty("current_dungeon_id", dungeon.Id())
		e.SendGameObjectUpdate(charGameObj, "update_object")
	})
}