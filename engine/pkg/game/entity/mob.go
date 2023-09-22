package entity

import (
	"math/rand"
)

const (
	IdleState = 0
	MovingState = 1
)

const (
	IdleTime = 40000.0
	MovingTime = 5000.0
)

const (
	RemoveCmd = 0
)

type MobCommand struct {
	// command to perform
	command int

	// command params
	params map[string]interface{}
}

type Mob struct {
	Id string
	Engine IEngine
	TickTime int64
	State int
	ticks chan int64
	commands chan *MobCommand
}

func NewMob(e IEngine, id string) *Mob {
	mob := &Mob{
		Id:       id,
		Engine:   e,
		TickTime: e.CurrentTickTime(),
		State:    IdleState,
		ticks:    make(chan int64),
		commands: make(chan *MobCommand),
	}

	go mob.Run()

	return mob
}

func (mob *Mob) GetTicksChannel() chan int64 {
	return mob.ticks
}

func (mob *Mob) GetCommandsChannel() chan *MobCommand {
	return mob.commands
}

func (mob *Mob) Run() {
	defer close(mob.ticks)
	defer close(mob.commands)
	for {
		select {
		case newTickTime := <-mob.ticks:
			// Mob logic processing goes here:
			// Stop moving.
			if mob.State == MovingState && (newTickTime - mob.TickTime) >= MovingTime {
				mobObj := mob.Engine.GameObjects()[mob.Id]
				mobObj.Properties["speed_x"] = 0.0
				mobObj.Properties["speed_y"] = 0.0
				mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
				mob.TickTime = newTickTime
				mob.State = IdleState
			}
			// Start moving.
			if mob.State == IdleState && (newTickTime - mob.TickTime) >= IdleTime {
				mobObj := mob.Engine.GameObjects()[mob.Id]
				mobSpeed := mobObj.Properties["speed"].(float64)
				possibleSpeeds := [3]float64{-mobSpeed, 0.0, mobSpeed}
				mobObj.Properties["speed_x"] = possibleSpeeds[rand.Intn(len(possibleSpeeds))]
				mobObj.Properties["speed_y"] = possibleSpeeds[rand.Intn(len(possibleSpeeds))]
				mob.Engine.SendGameObjectUpdate(mobObj, "update_object")
				mob.TickTime = newTickTime
				mob.State = MovingState
			}
		case cmd := <-mob.commands:
			// We should always first remove mob from e.Mobs() and then send RemoveCmd
			if cmd.command == RemoveCmd {
				return
			}
		}
	}
}
