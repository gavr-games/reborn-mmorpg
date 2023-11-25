package moving_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Stop the object
func (mObj *MovingObject) Stop(e entity.IEngine) {
	obj := mObj.gameObj
	obj.Properties()["speed_x"] = 0.0
	obj.Properties()["speed_y"] = 0.0
	e.SendGameObjectUpdate(obj, "update_object")
}
