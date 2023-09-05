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
//TODO: add floor parameter here and boolean addToFloor
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
	case "tree_5":
		gameObj = &entity.GameObject{
			X: x - 0.5,
			Y: y - 0.5,
			Width: 1.0,
			Height: 1.0,
			Id: id,
			Type: "tree",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["width"] = 1.0
		gameObj.Properties["height"] = 1.0
		gameObj.Properties["x"] = x
		gameObj.Properties["y"] = y
		gameObj.Properties["shape"] = "circle"
		gameObj.Properties["kind"] = objKind
		gameObj.Properties["collidable"] = true
	case "axe":
		gameObj = &entity.GameObject{
			X: x - 0.5,
			Y: y - 0.5,
			Width: 1.0,
			Height: 1.0,
			Id: id,
			Type: "tool",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["width"] = 1
		gameObj.Properties["height"] = 1
		gameObj.Properties["x"] = x
		gameObj.Properties["y"] = y
		gameObj.Properties["shape"] = "rectangle"
		gameObj.Properties["kind"] = objKind
		// pickable
		// equipable
		// possible actions ?
	case "backpack":
		gameObj = &entity.GameObject{
			X: x - 0.5,
			Y: y - 0.5,
			Width: 1.0,
			Height: 1.0,
			Id: id,
			Type: "container",
			Properties: make(map[string]interface{}),
		}
		gameObj.Properties["width"] = 1
		gameObj.Properties["height"] = 1
		gameObj.Properties["x"] = x
		gameObj.Properties["y"] = y
		gameObj.Properties["shape"] = "rectangle"
		gameObj.Properties["kind"] = objKind
		gameObj.Properties["max_capacity"] = 16
		gameObj.Properties["free_capacity"] = 16
		gameObj.Properties["parent_container_id"] = nil
		gameObj.Properties["visible"] = false
		gameObj.Properties["items_ids"] = make([]string, gameObj.Properties["max_capacity"].(int))
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
		gameObj.Properties["slots"] = make(map[string]interface{})
		slots := gameObj.Properties["slots"].(map[string]interface{})
		slots["left_arm"] = nil
		slots["back"] = nil
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
	}
	if gameObj != nil && additionalProps != nil {
		for k, v := range additionalProps {
			gameObj.Properties[k] = v
		}
	}
	if gameObj != nil && objKind != "player_vision_area" {
		storage.GetClient().Updates <- gameObj
	}
	return gameObj
}
