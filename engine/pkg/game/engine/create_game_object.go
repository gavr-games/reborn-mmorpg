package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func CreateGameObject(objKind string, id int, x float64, y float64) *entity.GameObject {
	switch t := objKind; t {
	case "grass":
		gameObj := &entity.GameObject{
			X: x,
			Y: y,
			Width: 1,
			Height: 1,
			Id: id,
			Type: "surface",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["width"] = 1
		gameObj.Properties["height"] = 1
		gameObj.Properties["x"] = x
		gameObj.Properties["y"] = y
		gameObj.Properties["shape"] = "rectangle"
		gameObj.Properties["kind"] = objKind
		return gameObj
	default:
		return nil
	}
}
