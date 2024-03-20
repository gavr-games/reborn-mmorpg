package character_object

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

func (charGameObj *CharacterObject) Move(e entity.IEngine, newX float64, newY float64, floor int) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		if visionAreaGameObj, visionAreaOk := e.GameObjects().Load(player.VisionAreaGameObjectId); visionAreaOk {
			e.Floors()[charGameObj.Floor()].FilteredRemove(charGameObj, func(b utils.IBounds) bool {
				return charGameObj.Id() == b.(entity.IGameObject).Id()
			})
			charGameObj.SetX(newX)
			charGameObj.SetY(newY)
			charGameObj.SetFloor(floor)
			e.Floors()[floor].Insert(charGameObj)

			// Update lifted item
			liftedObjectId := charGameObj.GetProperty("lifted_object_id")
			if liftedObjectId != nil {
				if liftedObj, liftOk := e.GameObjects().Load(liftedObjectId.(string)); liftOk {
					if liftedObj != nil {
						e.Floors()[liftedObj.Floor()].FilteredRemove(liftedObj, func(b utils.IBounds) bool {
							return liftedObj.Id() == b.(entity.IGameObject).Id()
						})
						liftedObj.SetX(charGameObj.X())
						liftedObj.SetY(charGameObj.Y())
						liftedObj.SetFloor(floor)
						e.Floors()[floor].Insert(liftedObj)
					}
				}
			}

			// Update vision area game object
			newVisionAreaX := charGameObj.GetVisionAreaX()
			newVisionAreaY := charGameObj.GetVisionAreaY()
			if visionAreaGameObj.X() != newVisionAreaX || visionAreaGameObj.Y() != newVisionAreaY {
				e.Floors()[visionAreaGameObj.Floor()].FilteredRemove(visionAreaGameObj, func(b utils.IBounds) bool {
					return visionAreaGameObj.Id() == b.(entity.IGameObject).Id()
				})
				visionAreaDx := newVisionAreaX - visionAreaGameObj.X()
				visionAreaDy := newVisionAreaY - visionAreaGameObj.Y()
				visionAreaGameObj.SetX(visionAreaGameObj.X() + visionAreaDx)
				visionAreaGameObj.SetY(visionAreaGameObj.Y() + visionAreaDy)
				e.Floors()[visionAreaGameObj.Floor()].Insert(visionAreaGameObj)
				go updateVisibleObjects(e, player, visionAreaDx, visionAreaDy, visionAreaGameObj)
			}
		}
	}
}

// Determine new and old visible objects, send updates to client
// In engine system of coords looks like this:
// .----> x (East)
// |
// |
// V
// y (North)
func updateVisibleObjects(e entity.IEngine, player *entity.Player, dx float64, dy float64, visionAreaGameObj entity.IGameObject) {
	var addObjects []utils.IBounds
	var removeObjects []utils.IBounds
	if dy > 0 && dx == 0 { // North
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y() + visionAreaGameObj.Height() - dy,
			Width:  visionAreaGameObj.Width(),
			Height: dy,
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y() - dy,
			Width:  visionAreaGameObj.Width(),
			Height: dy,
		})
	} else if dy < 0 && dx == 0 { // South
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y(),
			Width:  visionAreaGameObj.Width(),
			Height: math.Abs(dy),
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y() + visionAreaGameObj.Height(),
			Width:  visionAreaGameObj.Width(),
			Height: math.Abs(dy),
		})
	} else if dy == 0 && dx > 0 { // East
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() + visionAreaGameObj.Width() - dx,
			Y:      visionAreaGameObj.Y(),
			Width:  dx,
			Height: visionAreaGameObj.Height(),
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y(),
			Width:  dx,
			Height: visionAreaGameObj.Height(),
		})
	} else if dy == 0 && dx < 0 { // West
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y(),
			Width:  math.Abs(dx),
			Height: visionAreaGameObj.Height(),
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() + visionAreaGameObj.Width(),
			Y:      visionAreaGameObj.Y(),
			Width:  math.Abs(dx),
			Height: visionAreaGameObj.Height(),
		})
	} else if dy > 0 && dx > 0 { // North-East
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y() + visionAreaGameObj.Height() - dy,
			Width:  visionAreaGameObj.Width(),
			Height: dy,
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y() - dy,
			Width:  visionAreaGameObj.Width(),
			Height: dy,
		})
		addObjects = append(addObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() + visionAreaGameObj.Width() - dx,
			Y:      visionAreaGameObj.Y(),
			Width:  dx,
			Height: visionAreaGameObj.Height() - dy,
		})...)
		removeObjects = append(removeObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y(),
			Width:  dx,
			Height: visionAreaGameObj.Height() - dy,
		})...)
	} else if dy < 0 && dx < 0 { // South-West
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y(),
			Width:  visionAreaGameObj.Width(),
			Height: math.Abs(dy),
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y() + visionAreaGameObj.Height(),
			Width:  visionAreaGameObj.Width(),
			Height: math.Abs(dy),
		})
		addObjects = append(addObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y() - dy,
			Width:  math.Abs(dx),
			Height: visionAreaGameObj.Height() + dy,
		})...)
		removeObjects = append(removeObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() + visionAreaGameObj.Width(),
			Y:      visionAreaGameObj.Y() - dy,
			Width:  math.Abs(dx),
			Height: visionAreaGameObj.Height() + dy,
		})...)
	} else if dy < 0 && dx > 0 { // South-West
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y(),
			Width:  visionAreaGameObj.Width(),
			Height: math.Abs(dy),
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y() + visionAreaGameObj.Height(),
			Width:  visionAreaGameObj.Width(),
			Height: math.Abs(dy),
		})
		addObjects = append(addObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() + visionAreaGameObj.Width() - dx,
			Y:      visionAreaGameObj.Y() - dy,
			Width:  dx,
			Height: visionAreaGameObj.Height() + dy,
		})...)
		removeObjects = append(removeObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y() - dy,
			Width:  dx,
			Height: visionAreaGameObj.Height() + dy,
		})...)
	} else if dy > 0 && dx < 0 { // North-West
		addObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y() + visionAreaGameObj.Height() - dy,
			Width:  visionAreaGameObj.Width(),
			Height: dy,
		})
		removeObjects = game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() - dx,
			Y:      visionAreaGameObj.Y() - dy,
			Width:  visionAreaGameObj.Width(),
			Height: dy,
		})
		addObjects = append(addObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X(),
			Y:      visionAreaGameObj.Y(),
			Width:  math.Abs(dx),
			Height: visionAreaGameObj.Height() - dy,
		})...)
		removeObjects = append(removeObjects, game_objects.GetVisibleObjects(e, visionAreaGameObj.Floor(), utils.Bounds{
			X:      visionAreaGameObj.X() + visionAreaGameObj.Width(),
			Y:      visionAreaGameObj.Y(),
			Width:  math.Abs(dx),
			Height: visionAreaGameObj.Height() - dy,
		})...)
	}
	if len(addObjects) > 0 {
		e.SendResponse("add_objects", map[string]interface{}{
			"objects": addObjects,
		}, player)
	}
	if len(removeObjects) > 0 {
		e.SendResponse("remove_objects", map[string]interface{}{
			"objects": removeObjects,
		}, player)
	}
}