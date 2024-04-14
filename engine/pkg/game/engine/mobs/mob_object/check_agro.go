package mob_object

import (
	"context"
	"strings"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Agressive mob finds first available target in radius and attacks
func (mob *MobObject) CheckAgro() {
	if ar := mob.GetProperty("agro_radius"); ar != nil {
		agroRadius := ar.(float64)
		// check potential attack objects in agro_raius
		e := mob.Engine
		targetObjId := ""
		gameArea, gaOk := e.GameAreas().Load(mob.GameAreaId())
		if !gaOk {
			return
		}
		possibleAttackObjects := gameArea.RetrieveIntersections(utils.Bounds{
			X:      mob.X() - agroRadius / 2 + mob.Width() / 2,
			Y:      mob.Y() - agroRadius / 2 + mob.Height() / 2,
			Width:  agroRadius * 2,
			Height: agroRadius * 2,
		})

		if len(possibleAttackObjects) > 0 {
			for _, val := range possibleAttackObjects {
				obj := val.(entity.IGameObject)
				checkDistance := false
				visible := obj.GetProperty("visible")
				isVisible := visible != nil && visible.(bool)
				// Attack visible players
				if obj.Kind() == "player" && isVisible {
					checkDistance = true
				}
				// Attack alive player dragons
				if strings.Contains(obj.Kind(), "_dragon") && isVisible {
					if alive := obj.GetProperty("alive"); alive != nil && alive.(bool) {
						if obj.GetProperty("owner_id") != nil {
							checkDistance = true
						}
					}
				}

				// Check distance
				// We need to check the distance, because we found all objects in square, but the agro zone is circle
				if checkDistance && mob.GetDistance(obj) <= agroRadius {
					targetObjId = obj.Id()
					break
				}
			}
		}

		// Start attack
		if targetObjId != "" {
			ctx := context.WithValue(context.Background(), "targetObjId", targetObjId)
			mob.FSM.Event(ctx, "attack")
		}
	}
}
