package mobs

import (
	"math"
	"math/rand"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
)

const (
	IdleState = 0
	MovingState = 1
	StartFollowState = 2
	FollowingState = 3
	StopFollowingState = 4
)

const (
	IdleTime = 40000.0 // stays idle during this time
	MovingTime = 5000.0 // randomly moves during this time
	FollowingTime = 40000.0 // stops following after this time
	FollowingDistance = 1.0 // stops when in range of the target
	FollowingDirectionChangeTime = 2000 // change direction only
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
			mob.performFollowing(newTickTime)
		}
	}
}

func (mob *Mob) stop() {
	mobObj := mob.Engine.GameObjects()[mob.Id]
	mobObj.Properties["speed_x"] = 0.0
	mobObj.Properties["speed_y"] = 0.0
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

func (mob *Mob) performFollowing(newTickTime int64) {
	targetObj, ok := mob.Engine.GameObjects()[mob.TargetObjectId]
	if ok {
		mobObj := mob.Engine.GameObjects()[mob.Id]
		if game_objects.GetDistance(mobObj, targetObj) <= FollowingDistance {
			mobObj.Properties["speed_x"] = 0.0
			mobObj.Properties["speed_y"] = 0.0
			mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
		} else {
			if (newTickTime - mob.directionTickTime >= FollowingDirectionChangeTime) {
				mob.directionTickTime = newTickTime
				// Calclate angle between mob and target
				// Choose the closest direction by angle by calculatin index in PossibleDirections slice
				dx := targetObj.X - mobObj.X
				dy := targetObj.Y - mobObj.Y
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
				mobDirection := game_objects.PossibleDirections[directionIndex]
				game_objects.SetXYSpeeds(mobObj, mobDirection)
				mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
			}
		}
	} else {
		mob.Unfollow()
	}
}

func (mob *Mob) Follow(targetObjId string) {
	mob.State = StartFollowState
	mob.TargetObjectId = targetObjId
}

func (mob *Mob) Unfollow() {
	mob.State = StopFollowingState
}
