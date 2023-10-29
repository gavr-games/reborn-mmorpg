package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/targets"
)

func Reborn(e entity.IEngine, charGameObj *entity.GameObject) {
	charGameObj.Properties["health"] = charGameObj.Properties["max_health"]
	charGameObj.X = constants.InitialPlayerX
	charGameObj.Y = constants.InitialPlayerY
	charGameObj.Properties["x"] = constants.InitialPlayerX
	charGameObj.Properties["y"] = constants.InitialPlayerY
	targets.Deselect(e, charGameObj)
	e.SendGameObjectUpdate(charGameObj, "update_object")
}