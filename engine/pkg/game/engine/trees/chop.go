package trees

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// This func is called via delayed action mechanism
// params: characterId, treeId
func Chop(e entity.IEngine, params map[string]interface{}) bool {
	tree := e.GameObjects()[params["treeId"].(string)].(entity.ITreeObject)
	character := e.GameObjects()[params["characterId"].(string)]
	return tree.Chop(e, character)
}