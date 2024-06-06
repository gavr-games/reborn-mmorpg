package game_objects

import (
	"errors"
	"fmt"
	"math"
	"strings"

	uuid "github.com/satori/go.uuid"

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
func CreateFromTemplate(e entity.IEngine, objPath string, x float64, y float64, rotation float64) (entity.IGameObject, error) {
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

	gameObj := &entity.GameObject{}
	gameObj.InitGameObject()
	gameObj.SetProperties(utils.CopyMap(objTemplate))
	gameObj.SetEffects(make(map[string]interface{}))
	gameObj.SetId(id)
	gameObj.SetX(x)
	gameObj.SetY(y)
	gameObj.SetWidth(width)
	gameObj.SetHeight(height)
	gameObj.SetType(objTemplate["type"].(string))
	gameObj.SetGameAreaId("")
	gameObj.SetRotation(rotation)

	if (gameObj.Type() == "container") {
		gameObj.SetProperty("items_ids", make([]interface{}, int(gameObj.GetProperty("max_capacity").(float64))))
	}

	// Some templates might have effects to be created with the object
	if effects := gameObj.GetProperty("effects"); effects != nil {
		gameObj.SetEffects(utils.CopyMap(effects.(map[string]interface{})))
		gameObj.SetProperty("effects", nil)
	}

	return e.CreateGameObjectStruct(gameObj), nil
}
