package mob_object

import (
	"context"
)

func (mob *MobObject) StopEverything() {
	mob.FSM.Event(context.Background(), "stop")
}
