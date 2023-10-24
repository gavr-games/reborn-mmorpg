package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Clones the object
func Clone(obj *entity.GameObject) *entity.GameObject {
	clone := &entity.GameObject{
		X: obj.X,
		Y: obj.Y,
		Width: obj.Width,
		Height: obj.Height,
		Id: obj.Id,
		Type: obj.Type,
		Floor: obj.Floor,
		Rotation: obj.Rotation,
		Properties: make(map[string]interface{}),
	}
	clone.Properties = utils.CopyMap(obj.Properties)
	return clone
}
