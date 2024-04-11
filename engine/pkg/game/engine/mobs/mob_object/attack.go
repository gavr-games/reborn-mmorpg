package mob_object

import (
	"context"
)

func (mob *MobObject) Attack(targetObjId string) {
	ctx := context.WithValue(context.Background(), "targetObjId", targetObjId)
	mob.FSM.Event(ctx, "attack")
}