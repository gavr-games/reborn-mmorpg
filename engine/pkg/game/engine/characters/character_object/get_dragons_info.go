package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func (charGameObj *CharacterObject) GetDragonsInfo(e entity.IEngine) map[string]interface{} {
	dragonsInfo := make(map[string]interface{})
	dragonsInfo["max_dragons"] = charGameObj.GetProperty("max_dragons")
	dragonsIds := charGameObj.GetProperty("dragons_ids").([]interface{})
	dragonsInfo["dragons"] = make([]map[string]interface{}, len(dragonsIds))
	i := 0
	for _, dragonId := range dragonsIds {
		if dragon, ok := e.GameObjects().Load(dragonId.(string)); ok {
			dragonsInfo["dragons"].([]map[string]interface{})[i] = serializers.GetInfo(e, dragon)
			i++
		}
	}
	return dragonsInfo
}
