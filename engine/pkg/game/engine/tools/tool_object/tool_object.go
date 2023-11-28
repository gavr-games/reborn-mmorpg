package tool_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/equipable_object"
)

type ToolObject struct {
	pickable_object.PickableObject
	equipable_object.EquipableObject
	entity.GameObject
}

func NewToolObject(gameObj entity.IGameObject) *ToolObject {
	tool := &ToolObject{pickable_object.PickableObject{}, equipable_object.EquipableObject{}, *gameObj.(*entity.GameObject)}
	tool.InitPickableObject(tool)
	tool.InitEquipableObject(tool)
	return tool
}
