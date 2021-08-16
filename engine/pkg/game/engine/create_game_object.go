package engine

import (
	"github.com/satori/go.uuid"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)
const (
	PlayerRadius = 0.5
	PlayerVisionArea = 50.0
	PlayerSpeed = 2.0
)

//TODO: optimize the memory by using integers instead of string constants
func CreateGameObject(objKind string, x float64, y float64, additionalProps map[string]interface{}) *entity.GameObject {
	id := uuid.NewV4().String()
	var gameObj *entity.GameObject
	gameObj = nil
	switch t := objKind; t {
	case "grass":
		gameObj = &entity.GameObject{
			X: x - 0.5,
			Y: y - 0.5,
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
		if additionalProps != nil {
			for k, v := range additionalProps {
				gameObj.Properties[k] = v
			}
		}
	case "rock_moss":
		gameObj = &entity.GameObject{
			X: x - 0.438,
			Y: y - 0.549,
			Width: 0.876,
			Height: 1.098,
			Id: id,
			Type: "rock",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["width"] = 0.876
		gameObj.Properties["height"] = 1.098
		gameObj.Properties["x"] = x
		gameObj.Properties["y"] = y
		gameObj.Properties["shape"] = "rectangle"
		gameObj.Properties["kind"] = objKind
		gameObj.Properties["collidable"] = true
		if additionalProps != nil {
			for k, v := range additionalProps {
				gameObj.Properties[k] = v
			}
		}
	case "player":
		gameObj = &entity.GameObject{
			X: x - PlayerRadius,
			Y: y - PlayerRadius,
			Width: 1,
			Height: 1,
			Id: id,
			Type: "player",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["radius"] = PlayerRadius
		gameObj.Properties["x"] = x
		gameObj.Properties["y"] = y
		gameObj.Properties["shape"] = "circle"
		gameObj.Properties["kind"] = objKind
		gameObj.Properties["speed"] = PlayerSpeed
		gameObj.Properties["speed_x"] = 0.0
		gameObj.Properties["speed_y"] = 0.0
		gameObj.Properties["visible"] = true
		if additionalProps != nil {
			for k, v := range additionalProps {
				gameObj.Properties[k] = v
			}
		}
	case "player_vision_area":
		gameObj = &entity.GameObject{
			X: x - PlayerVisionArea / 2,
			Y: y - PlayerVisionArea / 2,
			Width: PlayerVisionArea,
			Height: PlayerVisionArea,
			Id: id,
			Type: "player_vision_area",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["width"] = PlayerVisionArea
		gameObj.Properties["height"] = PlayerVisionArea
		gameObj.Properties["x"] = x - PlayerVisionArea / 2
		gameObj.Properties["y"] = y - PlayerVisionArea / 2
		gameObj.Properties["shape"] = "rectangle"
		gameObj.Properties["kind"] = objKind
		gameObj.Properties["visible"] = false
		if additionalProps != nil {
			for k, v := range additionalProps {
				gameObj.Properties[k] = v
			}
		}
	}
	if gameObj != nil && objKind != "player_vision_area" {
		storage.GetClient().Updates <- gameObj
	}
	return gameObj
}
