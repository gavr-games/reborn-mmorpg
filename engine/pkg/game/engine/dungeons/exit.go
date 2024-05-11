package dungeons

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

const (
	ExitDistance = 0.5
)

// Exit the dungeon, destroy it, teleport char and dragons back
func Exit(e entity.IEngine, charGameObj entity.IGameObject, dungeonExitObj entity.IGameObject) (bool, error) {
	playerId := charGameObj.GetProperty("player_id").(int)
	if player, ok := e.Players().Load(playerId); ok {
		slots := charGameObj.GetProperty("slots").(map[string]interface{})

		if dungeonExitObj.GetDistance(charGameObj) > ExitDistance {
			e.SendSystemMessage("You need to be closer.", player)
			return false, errors.New("Player need to be closer to exit")
		}

		if charId := dungeonExitObj.GetProperty("character_id"); charId == nil || charId.(string) != charGameObj.Id() {
			e.SendSystemMessage("You don't have permission to exit this dungeon", player)
			return false, errors.New("Player does not have permission to exit this dungeon")
		}

		if slots["back"] == nil {
			e.SendSystemMessage("You don't have container", player)
			return false, errors.New("Player does not have container")
		}

		var (
			container entity.IGameObject
			contOk bool
		)
		if container, contOk = e.GameObjects().Load(slots["back"].(string)); !contOk {
			e.SendSystemMessage("You don't have container", player)
			return false, errors.New("Player does not have container")
		}

		// Check player has dungeon key
		dungeonKey := container.(entity.IContainerObject).GetItemKind(e, "dungeon_key")
		if dungeonKey == nil {
			e.SendSystemMessage("You don't have dungeon exit key", player)
			return false, errors.New("Player does not have dungeon exit key")
		}

		// Remove key
		if !container.(entity.IContainerObject).Remove(e, player, dungeonKey.Id()) {
			e.SendSystemMessage("Can't remove dungeon key", player)
			return false, errors.New("Can't remove dungeon key")
		}
		e.RemoveGameObject(dungeonKey)

		// Add exp to char and dragons
		dungeonLevel := dungeonExitObj.GetProperty("level").(float64)
		dragonIds := dungeonExitObj.GetProperty("dragon_ids").([]interface{})
		charGameObj.(entity.ILevelingObject).AddDungeonExperience(e, dungeonLevel)
		for _, id := range dragonIds {
			dragonId := id.(string)
			if dragon, dOk := e.GameObjects().Load(dragonId); dOk {
				dragon.(entity.ILevelingObject).AddDungeonExperience(e, dungeonLevel)
			}
		}

		// Increase allowed dungeon lvl
		charMaxLevel := charGameObj.GetProperty("max_dungeon_lvl").(float64)
		if dungeonLevel == charMaxLevel {
			charGameObj.SetProperty("max_dungeon_lvl", charMaxLevel + 1.0)
			storage.GetClient().Updates <- charGameObj.Clone()
		}

		// Teleport character to town
		charGameObj.(entity.ICharacterObject).TownTeleport(e)

		// Destroy dungeon (remove current_dungeon_id, teleport dragons)
		go Destroy(e, charGameObj)

		return true, nil
	}

	return false, errors.New("Player does not exist")
}
