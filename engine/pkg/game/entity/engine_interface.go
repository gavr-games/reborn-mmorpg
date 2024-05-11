package entity

import (
	"github.com/puzpuzpuz/xsync/v3"
)

type Task func()

type IEngine interface {
	GameAreas() *xsync.MapOf[string, *GameArea]
	Players() *xsync.MapOf[int, *Player]
	GameObjects() *xsync.MapOf[string, IGameObject]
	Mobs() *xsync.MapOf[string, IMobObject]
	Effects() *xsync.MapOf[string, map[string]interface{}]
	CurrentTickTime() int64
	PerformTask(f func())
	EnableTestingMode()
	GetGameAreaByKey(key string) *GameArea
	SendResponse(responseType string, responseData map[string]interface{}, player *Player)
	SendResponseToVisionAreas(gameObj IGameObject, responseType string, responseData map[string]interface{})
	SendGameObjectUpdate(gameObj IGameObject, updateType string)
	SendSystemMessage(message string, player *Player)
	CreateGameObjectStruct(gameObj IGameObject) IGameObject
	CreateGameObject(objPath string, x float64, y float64, rotation float64, gameAreaId string, additionalProps map[string]interface{}) IGameObject
	RemoveGameObject(gameObj IGameObject)
}
