package entity

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

type IEngine interface {
	Floors() []*utils.Quadtree
	Players() map[int]*Player
	GameObjects() map[string]*GameObject
	Mobs() map[string] IMob
	Effects() map[string]map[string]interface{}
	CurrentTickTime() int64
	SendResponse(responseType string, responseData map[string]interface{}, player *Player)
	SendResponseToVisionAreas(gameObj *GameObject, responseType string, responseData map[string]interface{})
	SendGameObjectUpdate(gameObj *GameObject, updateType string)
	SendSystemMessage(message string, player *Player)
	CreateGameObject(objPath string, x float64, y float64, floor int, additionalProps map[string]interface{}) *GameObject
}
