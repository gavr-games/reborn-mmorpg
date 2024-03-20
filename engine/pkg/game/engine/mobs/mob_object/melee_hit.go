package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) MeleeHit(targetObj entity.IGameObject) bool {
	// Check Cooldown
	// here we cast everything to float64, because go restores from json everything as float64
	if lastHitAt := mob.GetProperty("last_hit_at"); lastHitAt != nil {
		if float64(utils.MakeTimestamp()) - lastHitAt.(float64) >= mob.GetProperty("cooldown").(float64) {
			mob.SetProperty("last_hit_at", float64(utils.MakeTimestamp()))
		} else {
			return false
		}
	} else {
		mob.SetProperty("last_hit_at", float64(utils.MakeTimestamp()))
	}

	// check collision with target
	if !mob.CanHit(mob, targetObj) {
		return false
	}

	// Send hit attempt to client
	mobClone := mob.Clone()
	mob.Engine.SendResponseToVisionAreas(mob, "melee_hit_attempt", map[string]interface{}{
		"object": mobClone,
		"weapon": mobClone, // mob has all required weapon attributes itself to act like weapon
	})

	// deduct health and update object
	targetObj.SetProperty("health", targetObj.GetProperty("health").(float64) - mob.GetProperty("damage").(float64))
	if targetObj.GetProperty("health").(float64) <= 0.0 {
		targetObj.SetProperty("health", 0.0)
	}
	// Trigger mob to aggro
	if targetObj.Type() == "mob" {
		if targetMob, ok := mob.Engine.Mobs().Load(targetObj.Id()); ok {
			targetMob.Attack(mob.Id())
		}
	}
	mob.Engine.SendGameObjectUpdate(targetObj, "update_object")

	// die if health < 0
	if targetObj.GetProperty("health").(float64) == 0.0 {
		mob.StopAttacking()
		if targetObj.Type() == "mob" {
			if targetMob, ok := mob.Engine.Mobs().Load(targetObj.Id()); ok {
				targetMob.Die()
			}
		} else {
			// for characters
			targetObj.(entity.ICharacterObject).Reborn(mob.Engine)
		}
	}

	return true
}