package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (charGameObj *CharacterObject) HasTypeEquipped(e entity.IEngine, itemType string) (entity.IGameObject, bool) {
	slots := charGameObj.Properties()["slots"].(map[string]interface{})
	
	for _, slotItemId := range slots {
		if slotItemId != nil {
			if slotItem, slotOk := e.GameObjects().Load(slotItemId.(string)); slotOk {
				if slotItem.Type() == itemType {
					return slotItem, true
				}
			}
		}
	}

	return nil, false
}