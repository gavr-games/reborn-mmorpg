package character_object

import (
	"fmt"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

// Tries to hit target with the melee weapon
func (obj *CharacterObject) MeleeHit(e entity.IEngine) bool {
	playerId := obj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		targetId := obj.GetProperty("target_id")

		if targetId == nil {
			e.SendSystemMessage("No target to hit.", player)
			return false
		}

		var (
			targetObj entity.IGameObject
			targetOk bool
		)
		if targetObj, targetOk = e.GameObjects().Load(targetId.(string)); !targetOk {
			obj.DeselectTarget(e)
			e.SendSystemMessage("No target to hit.", player)
			return false
		}

		// Check has melee_weapon equipped
		weapon, equipped := obj.HasTypeEquipped(e, "melee_weapon")
		if !equipped {
			e.SendSystemMessage("You need to equip weapon to hit.", player)
			return false
		}

		// Check Cooldown
		// here we cast everything to float64, because go restores from json everything as float64
		lastHitAt := weapon.GetProperty("last_hit_at")
		if lastHitAt != nil {
			if float64(utils.MakeTimestamp())-lastHitAt.(float64) >= weapon.GetProperty("cooldown").(float64) {
				weapon.SetProperty("last_hit_at", float64(utils.MakeTimestamp()))
			} else {
				return false
			}
		} else {
			weapon.SetProperty("last_hit_at", float64(utils.MakeTimestamp()))
		}

		// Send hit attempt to client
		e.SendResponseToVisionAreas(obj, "melee_hit_attempt", map[string]interface{}{
			"object": obj.Clone(),
			"weapon": weapon.Clone(),
		})

		// check collision with target
		if !weapon.(entity.IMeleeWeaponObject).CanHit(obj, targetObj) {
			return false
		}

		// deduct health and update object
		damage := weapon.GetProperty("damage").(float64)
		targetObj.SetProperty("health", targetObj.GetProperty("health").(float64) - damage)
		if targetObj.GetProperty("health").(float64) <= 0.0 {
			targetObj.SetProperty("health", 0.0)
		}
		// Trigger mob to aggro
		if targetObj.Type() == "mob" {
			if mob, ok := e.Mobs().Load(targetObj.Id()); ok {
				mob.Attack(obj.Id())
			}
		}

		e.SendSystemMessage(fmt.Sprintf("You dealt %d damage to %s.", int(damage), targetObj.Kind()), player)

		// die if health == 0
		if targetObj.GetProperty("health").(float64) == 0.0 {
			obj.DeselectTarget(e)
			if targetObj.Type() == "mob" {
				e.SendSystemMessage(fmt.Sprintf("You killed %s.", targetObj.Kind()), player)
				if mob, ok := e.Mobs().Load(targetObj.Id()); ok {
					mob.Die()
					obj.AddExperience(e, fmt.Sprintf("kill_mob/%s", targetObj.Kind()))
				}
			} else {
				// for characters
				targetObj.(entity.ICharacterObject).Reborn(e)
			}
		} else {
			e.SendGameObjectUpdate(targetObj, "update_object")
		}
	} else {
		return false
	}
	return true
}
