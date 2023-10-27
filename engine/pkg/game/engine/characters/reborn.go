package characters

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

const InitialPlayerX = 4.0
const InitialPlayerY = 4.0

func Reborn(e entity.IEngine, charGameObj *entity.GameObject) {
	charGameObj.Properties["health"] = charGameObj.Properties["max_health"]
	charGameObj.X = InitialPlayerX
	charGameObj.Y = InitialPlayerY
	charGameObj.Properties["x"] = InitialPlayerX
	charGameObj.Properties["y"] = InitialPlayerY
	e.SendGameObjectUpdate(charGameObj, "update_object")
}