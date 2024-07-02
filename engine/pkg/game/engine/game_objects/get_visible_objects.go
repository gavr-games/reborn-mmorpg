package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// get visible objects for bounds
func GetVisibleObjects(e entity.IEngine, gameAreaId string, bounds utils.Bounds) []utils.IBounds {
	if gameArea, gaOk := e.GameAreas().Load(gameAreaId); gaOk {
		visibleObjects := gameArea.RetrieveIntersections(bounds)
		// Filter visible objects
		n := 0
		for _, val := range visibleObjects {
			gameObj := val.(entity.IGameObject)
			if visible := gameObj.GetProperty("visible"); visible != nil {
				if visible.(bool) {
					visibleObjects[n] = val
					n++
				}
			} else {
				visibleObjects[n] = val
				n++
			}
		}
		visibleObjects = visibleObjects[:n]
		return visibleObjects
	} else {
		return nil
	}
}
