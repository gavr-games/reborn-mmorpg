package game_objects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Check object can move and does not collide with other objects
// TODO: implement accurate check for circles
func CanMove(floor *utils.Quadtree, obj *entity.GameObject, dx float64, dy float64) (float64, float64) {
	possibleCollidableObjects := floor.RetrieveIntersections(utils.Bounds{
		X:      obj.X + dx,
		Y:      obj.Y + dy,
		Width:  obj.Width,
		Height: obj.Height,
	})
	// Filter collidable objects
	n := 0
	for _, val := range possibleCollidableObjects {
		if collidable, ok := val.(*entity.GameObject).Properties["collidable"]; ok {
			if collidable.(bool) {
				possibleCollidableObjects[n] = val
				n++
			}
		}
	}
	possibleCollidableObjects = possibleCollidableObjects[:n]

	if len(possibleCollidableObjects) == 0 {
		return dx, dy
	} else {
		return 0.0, 0.0
	}
	
}
