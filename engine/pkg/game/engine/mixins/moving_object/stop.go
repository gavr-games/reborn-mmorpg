package moving_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Stop the object
func (mObj *MovingObject) Stop(e entity.IEngine) {
	obj := mObj.gameObj
	obj.SetProperty("speed_x", 0.0)
	obj.SetProperty("speed_y", 0.0)
	obj.SetMoveToCoords(nil)
	e.SendGameObjectUpdate(obj, "update_object")
}
