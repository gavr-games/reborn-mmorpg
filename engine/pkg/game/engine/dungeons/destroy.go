package dungeons

import (
	"errors"
	"strings"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

// Destroy dungeon, teleport dragons to owner
func Destroy(e entity.IEngine, charGameObj entity.IGameObject) (bool, error) {
	dungeonId := charGameObj.GetProperty("current_dungeon_id")
	if dungeonId == nil {
		return false, errors.New("Dungeon does not exist")
	}

	if dungeon, ok := e.GameAreas().Load(dungeonId.(string)); ok {
		// Remove current dungeon id
		charGameObj.SetProperty("current_dungeon_id", nil)
		storage.GetClient().Updates <- charGameObj.Clone()

		// Remove all game objects
		gameObjects := dungeon.GetAllGameObjects()
		for _, val := range gameObjects {
			gameObj := val.(entity.IGameObject)
			performRemove := true
			if strings.Contains(gameObj.Kind(), "_dragon") { // Teleport dragons to owner
				ownerId := gameObj.GetProperty("owner_id")
				if ownerId != nil && ownerId.(string) == charGameObj.Id() {
					performRemove = false
					e.PerformTask(func() { // send this code to engine main loop
						if ok, _ := gameObj.(entity.IDragonObject).TeleportToOwner(charGameObj); !ok {
							// If dragon is dead, just move it to same area as character
							gameObj.SetGameAreaId(charGameObj.GameAreaId())
							storage.GetClient().Updates <- gameObj.Clone()
						}
					})
				}
			}
			if performRemove {
				e.RemoveGameObject(gameObj) // TODO: improve performance by removing without cleaning quadtree and sending updates to players
			}
		}
		

		// Remove GameArea
		e.GameAreas().Delete(dungeon.Id())

		return true, nil
	} else {
		return false, errors.New("Dungeon does not exist")
	}
}
