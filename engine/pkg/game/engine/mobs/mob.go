package mobs

import (
	"math"
	"math/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/melee_weapons"
)

const (
	IdleState = 0
	MovingState = 1
	StartFollowState = 2
	FollowingState = 3
	StopFollowingState = 4
	StartAttackingState = 5
	AttackingState = 6
	StopAttackingingState = 7
	RenewAttackingState = 8
)

const (
	IdleTime = 40000.0 // stays idle during this time
	MovingTime = 5000.0 // randomly moves during this time
	FollowingTime = 40000.0 // stops following after this time
	FollowingDistance = 0.1 // stops when in range of the target
	FollowingDirectionChangeTime = 2000 // change direction only once per this time
	AttackSpeedUp = 1.5 // increases the mob speed during attack
	AttackingTime = 20000.0 // during this time the mob attacks until it calms down if not hitted back
	AttackingDirectionChangeTime = 500 // change direction only once per this time
)

type Mob struct {
	Id string // GameObjId
	Engine entity.IEngine
	TickTime int64
	State int
	TargetObjectId string
	directionTickTime int64 // when direction was last time changed
}

func NewMob(e entity.IEngine, id string) *Mob {
	mob := &Mob{
		Id:       id,
		Engine:   e,
		TickTime: e.CurrentTickTime(),
		State:    IdleState,
		TargetObjectId: "", // for following and attack
		directionTickTime: e.CurrentTickTime(),
	}

	return mob
}

func (mob *Mob) GetId() string {
	return mob.Id
}

func (mob *Mob) Run(newTickTime int64) {
	// Mob logic processing goes here:
	// Stop moving.
	if mob.State == MovingState && (newTickTime - mob.TickTime) >= MovingTime {
		mob.stop()
		mob.TickTime = newTickTime
	} else
	// Start random moving.
	if mob.State == IdleState && (newTickTime - mob.TickTime) >= IdleTime {
		mob.moveInRandomDirection()
		mob.TickTime = newTickTime
	} else
	// Stop following
	if mob.State == StopFollowingState {
		mob.TargetObjectId = ""
		mob.stop()
		mob.TickTime = newTickTime
	} else 
	// Start Following
	if mob.State == StartFollowState {
		mob.State = FollowingState
		mob.TickTime = newTickTime
		mob.directionTickTime = newTickTime
	} else
	// Perform following
	if mob.State == FollowingState {
		if (newTickTime - mob.TickTime) >= FollowingTime {
			mob.Unfollow()
		} else { // Perform actual following
			targetObj, ok := mob.Engine.GameObjects()[mob.TargetObjectId]
			if ok {
				mob.performFollowing(newTickTime, targetObj, FollowingDirectionChangeTime)
			} else {
				mob.Unfollow()
			}
		}
	} else
	// Start Attacking
	if mob.State == StartAttackingState {
		mob.State = AttackingState
		mob.TickTime = newTickTime
		mob.directionTickTime = newTickTime
		mobObj := mob.Engine.GameObjects()[mob.Id]
		mobObj.Properties()["speed"] = mobObj.Properties()["speed"].(float64) * AttackSpeedUp
	} else
	// Renew Attacking
	if mob.State == RenewAttackingState {
		mob.State = AttackingState
		mob.TickTime = newTickTime
		mob.directionTickTime = newTickTime
	} else
	// Stop attacking
	if mob.State == StopAttackingingState {
		mob.TargetObjectId = ""
		mobObj := mob.Engine.GameObjects()[mob.Id]
		mobObj.Properties()["speed"] = mobObj.Properties()["speed"].(float64) / AttackSpeedUp
		mob.stop()
		mob.TickTime = newTickTime
	} else
	// Perform attacking
	if mob.State == AttackingState {
		if (newTickTime - mob.TickTime) >= AttackingTime {
			mob.StopAttacking()
		} else { // Perform actual following before hit
			targetObj, ok := mob.Engine.GameObjects()[mob.TargetObjectId]
			if ok {
				mob.performFollowing(newTickTime, targetObj, AttackingDirectionChangeTime)
				mob.MeleeHit(targetObj)
			} else {
				mob.StopAttacking()
			}
		}
	} 
}

func (mob *Mob) stop() {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	mobObj.Properties()["speed_x"] = 0.0
	mobObj.Properties()["speed_y"] = 0.0
	mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
	mob.State = IdleState
}

func (mob *Mob) moveInRandomDirection() {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	mobDirection := game_objects.PossibleDirections[rand.Intn(len(game_objects.PossibleDirections))]
	game_objects.SetXYSpeeds(mobObj, mobDirection)
	mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
	mob.State = MovingState
}

func (mob *Mob) performFollowing(newTickTime int64, targetObj entity.IGameObject, directionChangeTime int64) {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	if mobObj.GetDistance(targetObj) <= FollowingDistance {
		// Stop the mob
		if mobObj.Properties()["speed_x"].(float64) != 0.0 || mobObj.Properties()["speed_y"].(float64) != 0.0 {
			mobObj.Properties()["speed_x"] = 0.0
			mobObj.Properties()["speed_y"] = 0.0
			mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
		}
		mob.turnToTarget(targetObj) // TODO: send only on change
	} else {
		if (newTickTime - mob.directionTickTime >= directionChangeTime) {
			mob.directionTickTime = newTickTime
			game_objects.SetXYSpeeds(mobObj, mob.getDirectionToTarget(targetObj))
			mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
		}
	}
}

func (mob *Mob) turnToTarget(targetObj entity.IGameObject) {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	direction := mob.getDirectionToTarget(targetObj)
	if mobObj.Rotation() != game_objects.GetRotation(direction) {
		game_objects.SetRotation(mobObj, direction)
		mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
	}
}

func (mob *Mob) getDirectionToTarget(targetObj entity.IGameObject) string {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	// Calclate angle between mob and target
	// Choose the closest direction by angle by calculatin index in PossibleDirections slice
	dx := targetObj.X() - mobObj.X()
	dy := targetObj.Y() - mobObj.Y()
	angle := math.Atan2(dy, dx) // range (-PI, PI)
	if angle < 0.0 {
		angle = angle + math.Pi * 2
	}
	quotient := math.Floor(angle / (math.Pi / 4)) // math.Pi / 4 - is the angle between movement directions
	remainder := angle - (math.Pi / 4) * quotient
	if (remainder > math.Pi / 8) {
		quotient = quotient + 1.0
	}
	directionIndex := int(quotient)
	if (directionIndex == len(game_objects.PossibleDirections)) {
		directionIndex = 0
	}
	return game_objects.PossibleDirections[directionIndex]
}

func (mob *Mob) Follow(targetObjId string) {
	mob.State = StartFollowState
	mob.TargetObjectId = targetObjId
}

func (mob *Mob) Unfollow() {
	mob.State = StopFollowingState
}

func (mob *Mob) Attack(targetObjId string) {
	if mob.State == AttackingState {
		mob.State = RenewAttackingState
	} else {
		mob.State = StartAttackingState
		mob.TargetObjectId = targetObjId
	}
}

func (mob *Mob) StopAttacking() {
	mob.State = StopAttackingingState
}

func (mob *Mob) MeleeHit(targetObj entity.IGameObject) bool {
	mobObj := mob.Engine.GameObjects()[mob.Id]

	// Check Cooldown
	// here we cast everything to float64, because go restores from json everything as float64
	lastHitAt, hitted := mobObj.Properties()["last_hit_at"]
	if hitted {
		if float64(utils.MakeTimestamp()) - lastHitAt.(float64) >= mobObj.Properties()["cooldown"].(float64) {
			mobObj.Properties()["last_hit_at"] = float64(utils.MakeTimestamp())
		} else {
			return false
		}
	} else {
		mobObj.Properties()["last_hit_at"] = float64(utils.MakeTimestamp())
	}

	// check collision with target
	if !melee_weapons.CanHit(mobObj, mobObj, targetObj) {
		return false
	}

	// Send hit attempt to client
	mob.Engine.SendResponseToVisionAreas(mobObj, "melee_hit_attempt", map[string]interface{}{
		"object": mobObj,
		"weapon": mobObj, // mob has all required weapon attributes itself to act like weapon
	})

	// deduct health and update object
	targetObj.Properties()["health"] = targetObj.Properties()["health"].(float64) - mobObj.Properties()["damage"].(float64)
	if targetObj.Properties()["health"].(float64) <= 0.0 {
		targetObj.Properties()["health"] = 0.0
	}
	// Trigger mob to aggro
	if targetObj.Properties()["type"].(string) == "mob" {
		mob.Engine.Mobs()[targetObj.Id()].Attack(mob.Id)
	}
	mob.Engine.SendGameObjectUpdate(targetObj, "update_object")

	// die if health < 0
	if targetObj.Properties()["health"].(float64) == 0.0 {
		mob.StopAttacking()
		if targetObj.Properties()["type"].(string) == "mob" {
			mob.Engine.Mobs()[targetObj.Id()].Die()
		} else {
			// for characters
			characters.Reborn(mob.Engine, targetObj)
		}
	}

	return true
}

func (mob *Mob) Die() {
	mob.drop()

	// remove from world
	mobObj := mob.Engine.GameObjects()[mob.Id]
	mob.Engine.Floors()[mobObj.Floor()].FilteredRemove(mobObj, func(b utils.IBounds) bool {
			return mob.Id == b.(entity.IGameObject).Id()
	})
	mob.Engine.GameObjects()[mob.Id] = nil
	delete(mob.Engine.GameObjects(), mob.Id)

	mob.Engine.SendGameObjectUpdate(mobObj, "remove_object")
}

func (mob *Mob) drop() {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	if drops, ok := mobObj.Properties()["drop"]; ok {
		for name, dropProperties := range drops.(map[string]interface{}) {
			probability := rand.Float64()
			if probability <= dropProperties.(map[string]interface{})["probability"].(float64) {
				additionalProps := make(map[string]interface{})
				additionalProps["visible"] = true
				if _, stackable := dropProperties.(map[string]interface{})["min"]; stackable {
					min := dropProperties.(map[string]interface{})["min"].(float64)
					max := dropProperties.(map[string]interface{})["max"].(float64)
					additionalProps["amount"] = math.Ceil((rand.Float64() * (max - min)) + min)
				}
				dropItem := mob.Engine.CreateGameObject(name, mobObj.X(), mobObj.Y(), 0.0, mobObj.Floor(), additionalProps)
				mob.Engine.SendGameObjectUpdate(dropItem, "add_object")
			}
		}
	}
}
