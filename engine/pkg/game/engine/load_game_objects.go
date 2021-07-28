package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func LoadGameObjects(floors []*utils.Quadtree, gameObjects map[int]*entity.GameObject, gameObjectsId *int, floorSize float64) {
	floors[0] = &utils.Quadtree{
		Bounds: utils.Bounds{
			X:      0,
			Y:      0,
			Width:  floorSize,
			Height: floorSize,
		},
		MaxObjects: 30,
		MaxLevels:  10,
		Level:      0,
		Objects:    make([]utils.IBounds, 0),
		Nodes:      make([]utils.Quadtree, 0),
	}

	for x := 0; x < 100; x++ {
    for y := 0; y < 100; y++ {
			*gameObjectsId++
			gameObj := CreateGameObject("grass", *gameObjectsId, float64(x), float64(y))
			gameObjects[*gameObjectsId] = gameObj
			floors[0].Insert(gameObj)
		}
	}
}