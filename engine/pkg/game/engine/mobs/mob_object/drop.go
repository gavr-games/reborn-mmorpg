package mob_object

import (
	"math"
	"math/rand"
)

func (mob *MobObject) drop() {
	if drops := mob.GetProperty("drop"); drops != nil {
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
				dropItem := mob.Engine.CreateGameObject(name, mob.X(), mob.Y(), 0.0, mob.Floor(), additionalProps)
				mob.Engine.SendGameObjectUpdate(dropItem, "add_object")
			}
		}
	}
}