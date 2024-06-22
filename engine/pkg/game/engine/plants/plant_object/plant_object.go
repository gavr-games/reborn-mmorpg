package plant_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/destroyable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type PlantObject struct {
	destroyable_object.DestroyableObject
	entity.GameObject
}

func NewPlantObject(gameObj entity.IGameObject) *PlantObject {
	plant := &PlantObject{
		destroyable_object.DestroyableObject{},
		*gameObj.(*entity.GameObject),
	}
	plant.InitDestroyableObject(plant)
	return plant
}
