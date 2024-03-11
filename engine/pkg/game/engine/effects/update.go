package effects

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Goes through all effects and tries to apply them
func Update(e entity.IEngine, tickDelta int64) {
	e.Effects().Range(func(effectId string, effect map[string]interface{}) bool {
		var (
			obj entity.IGameObject
			ok bool
		)
		if obj, ok = e.GameObjects().Load(effect["target_id"].(string)); !ok {
			// remove effect if game object is gone
			Remove(e, effectId, nil)
		}

		// apply effect
		/*
		Example:
		"effect": map[string]interface{}{
					"type": "periodic", 
					"attribute": "health",
					"value": 5.0,
					"cooldown": 2000.0,
					"current_cooldown": 0.0,
					"number": 10.0,
					"group": "potion_healing",
				},
		*/
		if effect["type"].(string) == "periodic" { // once per cooldown
			effect["current_cooldown"] = effect["current_cooldown"].(float64) + float64(tickDelta)
			if effect["current_cooldown"].(float64) >= effect["cooldown"].(float64) {
				effect["current_cooldown"] = 0.0
				obj.Properties()[effect["attribute"].(string)] = obj.Properties()[effect["attribute"].(string)].(float64) + effect["value"].(float64)
				effect["number"] = effect["number"].(float64) - 1.0
				if effect["number"].(float64) <= 0.0 {
					Remove(e, effectId, obj)
				} else {
					e.SendGameObjectUpdate(obj, "update_object")
				}
				// Health
				if effect["attribute"].(string) == "health" {
					if obj.Properties()["health"].(float64) > obj.Properties()["max_health"].(float64) {
						obj.Properties()["health"] = obj.Properties()["max_health"].(float64)
					} else 
					// die if health < 0
					if obj.Properties()["health"].(float64) <= 0.0 {
						if obj.Properties()["type"].(string) == "mob" {
							if mob, ok := e.Mobs().Load(obj.Id()); ok {
								mob.Die()
							}
						} else {
							// for characters
							obj.(entity.ICharacterObject).Reborn(e)
						}
					}
				}
			}
		}
		return true
	})
}
