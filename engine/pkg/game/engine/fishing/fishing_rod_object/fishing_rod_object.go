package fishing_rod_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/tools/tool_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// TODO: add container for bait
type FishingRodObject struct {
	tool_object.ToolObject
}

func NewFishingRodObject(gameObj entity.IGameObject) *FishingRodObject {
	fishingRod := &FishingRodObject{
		*tool_object.NewToolObject(gameObj),
	}
	return fishingRod
}
