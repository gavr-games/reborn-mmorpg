package players

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// get visible objects for player
func GetVisibleObjects(e entity.IEngine, player *entity.Player) []utils.IBounds {
	visionArea := e.GameObjects()[player.VisionAreaGameObjectId]
	visibleObjects := e.Floors()[visionArea.Floor()].RetrieveIntersections(utils.Bounds{
		X:      visionArea.X(),
		Y:      visionArea.Y(),
		Width:  visionArea.Width(),
		Height: visionArea.Height(),
	})
	// Filter visible objects
	n := 0
	for _, val := range visibleObjects {
		if visible, ok := val.(entity.IGameObject).Properties()["visible"]; ok {
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
