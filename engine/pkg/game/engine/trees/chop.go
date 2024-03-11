package trees

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, treeId
func Chop(e entity.IEngine, params map[string]interface{}) bool {
	var (
		tree, character entity.IGameObject
		treeOk, charOk bool
	)
	if tree, treeOk = e.GameObjects().Load(params["treeId"].(string)); !treeOk {
		return false
	}
	if character, charOk = e.GameObjects().Load(params["characterId"].(string)); !charOk {
		return false
	}
	return tree.(entity.ITreeObject).Chop(e, character)
}