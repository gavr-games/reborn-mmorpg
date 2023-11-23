package effects

import (
	"github.com/satori/go.uuid"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/containers"
)

func ApplyPlayer(e entity.IEngine, itemId string, player *entity.Player) bool {
	item := e.GameObjects()[itemId]
	obj := e.GameObjects()[player.CharacterGameObjectId]

	if item == nil {
		e.SendSystemMessage("Wrong item.", player)
		return false
	}
	
	// check container belongs to character
	if (item.Properties()["container_id"] != nil) {
		if !containers.CheckAccess(e, player, e.GameObjects()[item.Properties()["container_id"].(string)]) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}
	} else {
		e.SendSystemMessage("You can use items only from container", player)
		return false
	}

	// Remove from container
	if (item.Properties()["container_id"] != nil) {
		if !containers.Remove(e, player, item.Properties()["container_id"].(string), itemId) {
			return false
		}
	}

	// Check same group effect is already applied and remove
	effectGroup := item.Properties()["effect"].(map[string]interface{})["group"].(string)
	for effectId, effect := range obj.Effects() {
		if effect.(map[string]interface{})["group"].(string) == effectGroup {
			Remove(e, effectId, obj)
		}
	}

	// Apply effect
	effectId := uuid.NewV4().String()
	obj.Effects()[effectId] = utils.CopyMap(item.Properties()["effect"].(map[string]interface{}))
	e.Effects()[effectId] = utils.CopyMap(item.Properties()["effect"].(map[string]interface{}))
	e.Effects()[effectId]["id"] = effectId
	e.Effects()[effectId]["target_id"] = obj.Id()
	e.SendGameObjectUpdate(obj, "update_object")

	// Remove item
	e.GameObjects()[itemId] = nil
	delete(e.GameObjects(), itemId)
	storage.GetClient().Deletes <- itemId

	return true
}
