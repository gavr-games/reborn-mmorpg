package entity

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

type IEngine interface {
	Floors() []*utils.Quadtree
	Players() map[int]*Player
	GameObjects() map[string]*GameObject
	SendResponse(responseType string, responseData map[string]interface{}, player *Player)
	SendResponseToVisionAreas(gameObj *GameObject, responseType string, responseData map[string]interface{})
	SendGameObjectUpdate(gameObj *GameObject, updateType string)
	SendSystemMessage(message string, player *Player)
}
