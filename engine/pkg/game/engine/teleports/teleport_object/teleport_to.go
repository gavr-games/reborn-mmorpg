package teleport_object

import (
	"errors"

	"pgregory.net/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	MaxRetries = 10
	TeleportDistance = 0.5
)

func (teleport *TeleportObject) TeleportTo(e entity.IEngine, charGameObj entity.IGameObject) (bool, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		teleportTo := teleport.GetProperty("teleport_to").(map[string]interface{})
		
		// Check Area
		area := e.GetGameAreaByKey(teleportTo["area"].(string))
		if area == nil {
			e.SendSystemMessage("You cannot teleport to this area.", player)
			return false, errors.New("area not found")
		}

		// Check near the teleport
		if teleport.GetDistance(charGameObj) > TeleportDistance {
			e.SendSystemMessage("You need to be closer to the teleport.", player)
			return false, errors.New("too far from teleport")
		}

		// Check coords
		x := 0.0
		y := 0.0
		randomCoords := teleportTo["random_coords"]
		if randomCoords != nil && randomCoords.(bool) {
			retries := 0
			canTeleport := true
			for {
				x = float64(rand.Intn(int(area.Width())))
				y = float64(rand.Intn(int(area.Height())))
				// Check not intersecting with another objects
				possibleCollidableObjects := area.RetrieveIntersections(utils.Bounds{
					X:      x,
					Y:      y,
					Width:  charGameObj.Width(),
					Height: charGameObj.Height(),
				})
				if len(possibleCollidableObjects) > 0 {
					for _, val := range possibleCollidableObjects {
						gameObj := val.(entity.IGameObject)
						collidable := gameObj.GetProperty("collidable")
						if (collidable != nil && collidable.(bool)) || gameObj.Kind() == "claim_area" {
							canTeleport = false
							break;
						}
					}
				}

				if canTeleport {
					break
				}
				if retries > MaxRetries {
					return false, errors.New("random coords failed")
				}
				retries++
			}
		} else {
			x = teleportTo["x"].(float64)
			y = teleportTo["y"].(float64)
		}

		// Move character
		charGameObj.(entity.ICharacterObject).DeselectTarget(e)
		charGameObjClone := charGameObj.Clone()
		e.SendResponseToVisionAreas(charGameObjClone, "remove_object", map[string]interface{}{
			"object": charGameObjClone,
		})
		charGameObj.(entity.ICharacterObject).Move(e, x, y, area.Id())
		e.SendGameObjectUpdate(charGameObj, "update_object")

		e.SendSystemMessage("You has been teleported.", player)
		return true, nil
	} else {
		return false, errors.New("player not found")
	}
}
