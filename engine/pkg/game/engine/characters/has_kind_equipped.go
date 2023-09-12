package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func HasKindEquipped(e entity.IEngine, charGameObj *entity.GameObject, kind string) bool {
	slots := charGameObj.Properties["slots"].(map[string]interface{})
	
	for _, slotItemId := range slots {
		if slotItemId != nil {
			slotItem := e.GameObjects()[slotItemId.(string)]
			if slotItem.Properties["kind"].(string) == kind {
				return true
			}
		}
	}

	return false
}