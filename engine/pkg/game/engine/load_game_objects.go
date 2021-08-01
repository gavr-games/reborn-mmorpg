package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func LoadGameObjects(floors []*utils.Quadtree, gameObjects map[string]*entity.GameObject, floorSize float64) {
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
			// + 0.5 because we want to place the center point
			gameObj := CreateGameObject("grass", float64(x) + 0.5, float64(y) + 0.5, nil)
			gameObj.Floor = 0
			gameObjects[gameObj.Id] = gameObj
			floors[0].Insert(gameObj)
		}
	}
}