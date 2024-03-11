package entity

import (
	"github.com/puzpuzpuz/xsync/v3"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

type IEngine interface {
	Floors() []*utils.Quadtree
	Players() *xsync.MapOf[int, *Player]
	GameObjects() map[string]IGameObject
	Mobs() *xsync.MapOf[string, IMobObject]
	Effects() *xsync.MapOf[string, map[string]interface{}]
	CurrentTickTime() int64
	SendResponse(responseType string, responseData map[string]interface{}, player *Player)
	SendResponseToVisionAreas(gameObj IGameObject, responseType string, responseData map[string]interface{})
	SendGameObjectUpdate(gameObj IGameObject, updateType string)
	SendSystemMessage(message string, player *Player)
	CreateGameObjectStruct(gameObj IGameObject) IGameObject
	CreateGameObject(objPath string, x float64, y float64, rotation float64, floor int, additionalProps map[string]interface{}) IGameObject
}
