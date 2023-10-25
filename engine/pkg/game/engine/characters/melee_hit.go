package characters

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mobs"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/targets"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/melee_weapons"
)

// Tries to hit target with the melee weapon
func MeleeHit(e entity.IEngine, obj *entity.GameObject, player *entity.Player) bool {
	targetId, ok := obj.Properties["target_id"]
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
	lastHitAt, hitted := weapon.Properties["last_hit_at"]
	if hitted {
		if utils.MakeTimestamp() - lastHitAt.(int64) >= int64(weapon.Properties["cooldown"].(float64)) {
			weapon.Properties["last_hit_at"] = utils.MakeTimestamp()
		} else {
			if player != nil {
				e.SendSystemMessage("Your weapon is on cooldown.", player)
			}
			return false
		}
	} else {
		weapon.Properties["last_hit_at"] = utils.MakeTimestamp()
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
	targetObj.Properties["health"] = targetObj.Properties["health"].(float64) - weapon.Properties["damage"].(float64)
	if targetObj.Properties["health"].(float64) <= 0.0 {
		targetObj.Properties["health"] = 0.0
	}
	storage.GetClient().Updates <- targetObj
	e.SendGameObjectUpdate(targetObj, "update_object")

	if player != nil {
		e.SendSystemMessage(fmt.Sprintf("You dealt %d damage to %s.", int(weapon.Properties["damage"].(float64)), targetObj.Properties["kind"].(string)), player)
	}

	// die if health < 0
	if targetObj.Properties["health"].(float64) == 0.0 {
		targets.Deselect(e, obj)
		if targetObj.Properties["type"].(string) == "mob" {
			if player != nil {
				e.SendSystemMessage(fmt.Sprintf("You killed %s.", targetObj.Properties["kind"].(string)), player)
			}
			mobs.Die(e, targetObj.Id)
		} else {
			//TODO: do something for character
		}
	}

	return true
}
