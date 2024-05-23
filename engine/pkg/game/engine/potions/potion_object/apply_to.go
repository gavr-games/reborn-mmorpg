package potion_object

import (
	"github.com/satori/go.uuid"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/effects"
)

func (item *PotionObject) ApplyToPlayer(e entity.IEngine, player *entity.Player) bool {
	var (
		obj, container entity.IGameObject
		ok, contOk bool
	)
	if obj, ok = e.GameObjects().Load(player.CharacterGameObjectId); !ok {
		return false
	}
	
	containerId := item.GetProperty("container_id")
	if containerId != nil {
		if container, contOk = e.GameObjects().Load(containerId.(string)); !contOk {
			return false
		}
		// check container belongs to character
		if !container.(entity.IContainerObject).CheckAccess(e, player) {
			e.SendSystemMessage("You don't have access to this container", player)
			return false
		}

		// Remove from container
		if containerId != nil {
			if !container.(entity.IContainerObject).Remove(e, player, item.Id()) {
				return false
			}
		}

		// Check same group effect is already applied and remove
		itemEffectMap := item.GetProperty("effect").(map[string]interface{})
		effectGroup := itemEffectMap["group"].(string)
		for effectId, effect := range obj.Effects() {
			if effect.(map[string]interface{})["group"].(string) == effectGroup {
				effects.Remove(e, effectId, obj)
			}
		}

		// Apply effect
		effectId := uuid.NewV4().String()
		obj.SetEffect(effectId, utils.CopyMap(itemEffectMap))
		effectMap := utils.CopyMap(itemEffectMap)
		effectMap["id"] = effectId
		effectMap["target_id"] = obj.Id()
		e.Effects().Store(effectId, effectMap)
		e.SendGameObjectUpdate(obj, "update_object")

		// Remove item
		e.RemoveGameObject(item)

		return true
	} else {
		e.SendSystemMessage("You can use items only from container", player)
		return false
	}
}
