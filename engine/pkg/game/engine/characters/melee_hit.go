package characters

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/targets"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/melee_weapons"
)

// Tries to hit target with the melee weapon
func MeleeHit(e entity.IEngine, obj entity.IGameObject, player *entity.Player) bool {
	targetId, ok := obj.Properties()["target_id"]
	if !ok {
		if player != nil {
			e.SendSystemMessage("No target to hit.", player)
		}
		return false
	}

	if targetId == nil {
		if player != nil {
			e.SendSystemMessage("No target to hit.", player)
		}
		return false
	}

	targetObj := e.GameObjects()[targetId.(string)]
	if targetObj == nil {
		targets.Deselect(e, obj)
		if player != nil {
			e.SendSystemMessage("No target to hit.", player)
		}
		return false
	}

	// Check has melee_weapon equipped
	weapon, equipped := HasTypeEquipped(e, obj, "melee_weapon")
	if !equipped {
		if player != nil {
			e.SendSystemMessage("You need to equip weapon to hit.", player)
		}
		return false
	}

	// Check Cooldown
	// here we cast everything to float64, because go restores from json everything as float64
	lastHitAt, hitted := weapon.Properties()["last_hit_at"]
	if hitted {
		if float64(utils.MakeTimestamp()) - lastHitAt.(float64) >= weapon.Properties()["cooldown"].(float64) {
			weapon.Properties()["last_hit_at"] = float64(utils.MakeTimestamp())
		} else {
			return false
		}
	} else {
		weapon.Properties()["last_hit_at"] = float64(utils.MakeTimestamp())
	}

	// Send hit attempt to client
	e.SendResponseToVisionAreas(obj, "melee_hit_attempt", map[string]interface{}{
		"object": obj,
		"weapon": weapon,
	})

	// check collision with target
	if !melee_weapons.CanHit(obj, weapon, targetObj) {
		return false
	}

	// deduct health and update object
	targetObj.Properties()["health"] = targetObj.Properties()["health"].(float64) - weapon.Properties()["damage"].(float64)
	if targetObj.Properties()["health"].(float64) <= 0.0 {
		targetObj.Properties()["health"] = 0.0
	}
	// Trigger mob to aggro
	if targetObj.Properties()["type"].(string) == "mob" {
		e.Mobs()[targetObj.Id()].Attack(obj.Id())
	}
	e.SendGameObjectUpdate(targetObj, "update_object")

	if player != nil {
		e.SendSystemMessage(fmt.Sprintf("You dealt %d damage to %s.", int(weapon.Properties()["damage"].(float64)), targetObj.Properties()["kind"].(string)), player)
	}

	// die if health < 0
	if targetObj.Properties()["health"].(float64) == 0.0 {
		targets.Deselect(e, obj)
		if targetObj.Properties()["type"].(string) == "mob" {
			if player != nil {
				e.SendSystemMessage(fmt.Sprintf("You killed %s.", targetObj.Properties()["kind"].(string)), player)
			}
			e.Mobs()[targetObj.Id()].Die()
		} else {
			// for characters
			Reborn(e, targetObj)
		}
	}

	return true
}
