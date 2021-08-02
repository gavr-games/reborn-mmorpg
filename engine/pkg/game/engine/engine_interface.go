package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

type IEngine interface {
	Floors() []*utils.Quadtree
	Players() map[int]*entity.Player
	GameObjects() map[string]*entity.GameObject
}
