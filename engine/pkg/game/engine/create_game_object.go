package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)
const (
	PlayerRadius = 0.5
	PlayerVisionArea = 50.0
	PlayerSpeed = 2.0
)

//TODO: add floor parameter here and boolean addToFloor
func CreateGameObject(objPath string, x float64, y float64, additionalProps map[string]interface{}) *entity.GameObject {
	gameObj, err := game_objects.CreateFromTemplate(objPath, x, y)
	if err != nil {
		//TODO: handle error
	}
	if additionalProps != nil {
		for k, v := range additionalProps {
			gameObj.Properties[k] = v
		}
	}
	if gameObj.Properties["kind"].(string) != "player_vision_area" {
		storage.GetClient().Updates <- gameObj
	}
	return gameObj
}
