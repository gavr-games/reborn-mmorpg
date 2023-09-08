package game_objects

import (
	"errors"
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func findTemplate(objPath string) (map[string]interface{}, error) {
	gameObjectsAtlas := GetObjectsAtlas()
	if strings.Contains(objPath, "/") {
		res := strings.Split(objPath, "/")
		objType := res[0] // like "tree"
		objKind := res[1] // like "pine_5"
		if _, ok := gameObjectsAtlas[objType]; !ok {
			return nil, errors.New(fmt.Sprintf("Object type %s not found", objType))
		}
		if _, ok := gameObjectsAtlas[objType][objKind]; !ok {
			return nil, errors.New(fmt.Sprintf("Object kind %s not found", objKind))
		}
		return gameObjectsAtlas[objType][objKind].(map[string]interface{}), nil
	} else {
		objType := objPath // like "tree", "rock"
		objKinds, ok := gameObjectsAtlas[objType]
		if !ok {
			return nil, errors.New(fmt.Sprintf("Object type %s not found", objType))
		}
		return utils.PickRandomInMap(objKinds).(map[string]interface{}), nil
	}
}

// objPath - "tree/pine_5", "rock/rock_moss". 
// If just type provided "tree", "rock" it chooses random object kind
func CreateFromTemplate(objPath string, x float64, y float64) (*entity.GameObject, error) {
	//TODO: return error if place is occupied for collidable objects like trees and rocks

	//Get object template from GameObjectAtlas
	objTemplate, err := findTemplate(objPath)
	if err != nil {
		return nil, err
	}

	id := uuid.NewV4().String()
	width := objTemplate["width"].(float64)
	height := objTemplate["height"].(float64)
	objX := x - width / 2
	objY := y - height / 2

	gameObj := &entity.GameObject{
		X: objX,
		Y: objY,
		Width: width,
		Height: height,
		Id: id,
		Type: objTemplate["type"].(string),
		Properties: make(map[string]interface{}),
	}
	gameObj.Properties = utils.CopyMap(objTemplate)
	gameObj.Properties["x"] = objX
	gameObj.Properties["y"] = objY
	gameObj.Properties["id"] = id

	if (gameObj.Properties["type"].(string) == "container") {
		gameObj.Properties["items_ids"] = make([]string, gameObj.Properties["max_capacity"].(int))
	}

	return gameObj, nil
}