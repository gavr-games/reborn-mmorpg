package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// get visible objects for bounds
func GetVisibleObjects(e entity.IEngine, floor int, bounds utils.Bounds) []utils.IBounds {
	visibleObjects := e.Floors()[floor].RetrieveIntersections(bounds)
	// Filter visible objects
	n := 0
	for _, val := range visibleObjects {
		if visible := val.(entity.IGameObject).GetProperty("visible"); visible != nil {
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
}
