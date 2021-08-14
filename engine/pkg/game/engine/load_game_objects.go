package engine

import (
	"math/rand"
)

func LoadGameObjects(e IEngine, floorSize float64) {
	// grass
	for x := 0; x < 100; x++ {
    for y := 0; y < 100; y++ {
			// + 0.5 because we want to place the center point
			gameObj := CreateGameObject("grass", float64(x) + 0.5, float64(y) + 0.5, nil)
			gameObj.Floor = 0
			e.GameObjects()[gameObj.Id] = gameObj
			e.Floors()[gameObj.Floor].Insert(gameObj)
		}
	}
	// rocks
	for i := 0; i < 20; i++ {
		x := 1.0 + rand.Float64() * (99.0 - 1.0)
		y := 1.0 + rand.Float64() * (99.0 - 1.0)
		gameObj := CreateGameObject("rock_moss", x, y, nil)
		gameObj.Floor = 0
		e.GameObjects()[gameObj.Id] = gameObj
		e.Floors()[gameObj.Floor].Insert(gameObj)
	}
}