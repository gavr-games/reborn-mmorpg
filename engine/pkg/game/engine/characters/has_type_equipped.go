package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func HasTypeEquipped(e entity.IEngine, charGameObj *entity.GameObject, itemType string) (*entity.GameObject, bool) {
	slots := charGameObj.Properties["slots"].(map[string]interface{})
	
	for _, slotItemId := range slots {
		if slotItemId != nil {
			slotItem := e.GameObjects()[slotItemId.(string)]
			if slotItem.Properties["type"].(string) == itemType {
				return slotItem, true
			}
		}
	}

	return nil, false
}