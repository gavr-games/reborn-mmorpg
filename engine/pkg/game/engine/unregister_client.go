package engine

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Process when player disconnects from the game
func UnregisterClient(e entity.IEngine, client entity.IClient) {
	if player, ok := e.Players()[client.GetCharacter().Id]; ok {
		if e.GameObjects()[player.VisionAreaGameObjectId] != nil && player.Client != nil {
			//TODO: handle issue this gives panic when closing closed channel
			//close(client.GetSendChannel())
			player.Client = nil
			e.Floors()[e.GameObjects()[player.VisionAreaGameObjectId].Floor].FilteredRemove(e.GameObjects()[player.VisionAreaGameObjectId], func(b utils.IBounds) bool {
				return player.VisionAreaGameObjectId == b.(*entity.GameObject).Id
			})
			e.GameObjects()[player.VisionAreaGameObjectId] = nil
			charObj := e.GameObjects()[player.CharacterGameObjectId]
			charObj.Properties["visible"] = false
			charObj.Properties["speed_x"] = 0.0
			charObj.Properties["speed_y"] = 0.0
			e.SendResponseToVisionAreas(charObj, "remove_object", map[string]interface{}{
				"object": charObj,
			})
		}
	}
}
