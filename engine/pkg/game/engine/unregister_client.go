package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func UnregisterClient(e IEngine, client entity.IClient) {
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		close(client.GetSendChannel())
		player.Client = nil
		e.Floors()[0].FilteredRemove(e.GameObjects()[player.VisionAreaGameObjectId], func(b utils.IBounds) bool {
			return player.VisionAreaGameObjectId == b.(entity.GameObject).Id
		})
		e.GameObjects()[player.VisionAreaGameObjectId] = nil
		e.GameObjects()[player.CharacterGameObjectId].Properties["visible"] = false
		e.GameObjects()[player.CharacterGameObjectId].Properties["speed_x"] = 0
		e.GameObjects()[player.CharacterGameObjectId].Properties["speed_y"] = 0
		player.VisibleObjects = nil
	}
}
