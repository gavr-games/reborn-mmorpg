package engine

func LoadGameObjects(e IEngine, floorSize float64) {
	for x := 0; x < 100; x++ {
    for y := 0; y < 100; y++ {
			// + 0.5 because we want to place the center point
			gameObj := CreateGameObject("grass", float64(x) + 0.5, float64(y) + 0.5, nil)
			gameObj.Floor = 0
			e.GameObjects()[gameObj.Id] = gameObj
			e.Floors()[gameObj.Floor].Insert(gameObj)
		}
	}
}