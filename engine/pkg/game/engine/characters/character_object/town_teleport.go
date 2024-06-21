package character_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (charGameObj *CharacterObject) TownTeleport(e entity.IEngine) bool {
	charGameObj.DeselectTarget(e)
	charGameObjClone := charGameObj.Clone()
	e.SendResponseToVisionAreas(charGameObjClone, "remove_object", map[string]interface{}{
		"object": charGameObjClone,
	})
	charGameObj.Move(e, constants.InitialPlayerX, constants.InitialPlayerY, e.GetGameAreaByKey("town").Id())
	e.SendGameObjectUpdate(charGameObj, "update_object")

	return true
}