package game_objects

import (
	"math"
	"errors"
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

func searchAtlas(gameObjectsAtlas map[string]map[string]interface{}, objKind string) (map[string]interface{}, error) {
	for _, objects := range gameObjectsAtlas {
		for _, obj := range objects {
			if obj.(map[string]interface{})["kind"].(string) == objKind {
				return obj.(map[string]interface{}), nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("Object kind %s not found", objKind))
}

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
			// try to find gameObject by kind, not type. Like "stone_hammer"
			objTemplate, err := searchAtlas(gameObjectsAtlas, objPath)
			return objTemplate, err
		}
		return utils.PickRandomInMap(objKinds).(map[string]interface{}), nil
	}
}

// objPath - "tree/pine_5", "rock/rock_moss". 
// If just type provided "tree", "rock" it chooses random object kind
func CreateFromTemplate(objPath string, x float64, y float64, rotation float64) (*entity.GameObject, error) {
	//TODO: return error if place is occupied for collidable objects like trees and rocks

	//Get object template from GameObjectAtlas
	objTemplate, err := findTemplate(objPath)
	if err != nil {
		return nil, err
	}

	id := uuid.NewV4().String()

	width := objTemplate["width"].(float64)
	height := objTemplate["height"].(float64)
	if int(math.Ceil(rotation / (math.Pi / 2.0))) % 2 == 1 {
		tempWidth := width
		width = height
		height = tempWidth
	}

	gameObj := &entity.GameObject{
		X: x,
		Y: y,
		Width: width,
		Height: height,
		Id: id,
		Type: objTemplate["type"].(string),
		Floor: -1, // -1 for does not belong to any floor
		Rotation: rotation,
		Properties: make(map[string]interface{}),
		Effects: make(map[string]interface{}),
	}
	gameObj.Properties = utils.CopyMap(objTemplate)
	gameObj.Properties["x"] = x
	gameObj.Properties["y"] = y
	gameObj.Properties["id"] = id

	if (gameObj.Properties["type"].(string) == "container") {
		gameObj.Properties["items_ids"] = make([]interface{}, gameObj.Properties["max_capacity"].(int))
	}

	return gameObj, nil
}
