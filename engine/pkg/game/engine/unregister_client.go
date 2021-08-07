package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Process when player disconnects from the game
func UnregisterClient(e IEngine, client entity.IClient) {
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		close(client.GetSendChannel())
		player.Client = nil
		e.Floors()[0].FilteredRemove(e.GameObjects()[player.VisionAreaGameObjectId], func(b utils.IBounds) bool {
			return player.VisionAreaGameObjectId == b.(*entity.GameObject).Id
		})
		e.GameObjects()[player.VisionAreaGameObjectId] = nil
		e.GameObjects()[player.CharacterGameObjectId].Properties["visible"] = false
		e.GameObjects()[player.CharacterGameObjectId].Properties["speed_x"] = 0.0
		e.GameObjects()[player.CharacterGameObjectId].Properties["speed_y"] = 0.0
		player.VisibleObjects = nil
		SendGameObjectUpdate(e, e.GameObjects()[player.CharacterGameObjectId], "remove_object")
	}
}
