package npc_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (npcObj *NpcObject) GetDungeonsInfo(e entity.IEngine, charGameObj entity.IGameObject) (map[string]interface{}, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		actions := npcObj.GetProperty("actions")

		if actions == nil || actions.(map[string]interface{})["dungeons"] == nil {
			e.SendSystemMessage("You can't teleport to dungeons with this NPC", player)
			return nil, errors.New("NPC cannot teleport to dungeons")
		}

		if npcObj.GetDistance(charGameObj) > TradeDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return nil, errors.New("Player need to be closer to NPC")
		}

		dungeonsInfo := charGameObj.(entity.ICharacterObject).GetDragonsInfo(e)
		dungeonsInfo["max_dungeon_lvl"] = charGameObj.GetProperty("max_dungeon_lvl")
		dungeonsInfo["current_dungeon_id"] = charGameObj.GetProperty("current_dungeon_id")

		return dungeonsInfo, nil
	}

	return nil, errors.New("Player does not exist")
}
