package dragon_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (dragon *DragonObject) Resurrect(charGameObj entity.IGameObject) (bool, error) {
	e := dragon.Engine
	// Check only owner can ask dragon to teleport
	if playerId := charGameObj.GetProperty("player_id"); playerId != nil {
		playerIdInt := playerId.(int)
		if player, ok := e.Players().Load(playerIdInt); ok {
			if dragon.GetProperty("owner_id") != nil && charGameObj.Id() == dragon.GetProperty("owner_id").(string) {
				if alive := dragon.GetProperty("alive"); alive != nil && alive.(bool) {
					e.SendSystemMessage("The dragon is already alive.", player)
					return false, errors.New("The dragon is already alive")
				}

				// check character has money
				if substracted, substractionErr := charGameObj.(entity.ICharacterObject).SubstractGold(e, RESURRECT_COST_PER_LEVEL * (dragon.GetProperty("level").(float64) + 1)); !substracted {
					return false, substractionErr
				}

				// check altar exists on claim
				obeliskId := charGameObj.GetProperty("claim_obelisk_id")
				if obeliskId == nil {
					e.SendSystemMessage("You need claim to resurrect dragon.", player)
					return false, errors.New("Claim does not exist")
				}
				claimObelisk, obeliskOk := e.GameObjects().Load(obeliskId.(string))
				if !obeliskOk {
					e.SendSystemMessage("You need claim to resurrect dragon.", player)
					return false, errors.New("Claim does not exist")
				}
				altar, err := claimObelisk.(entity.IClaimObeliskObject).FindKindInArea(e, "dragon_altar")
				if err != nil {
					e.SendSystemMessage("You need dragon altar on the claim to resurrect dragon.", player)
					return false, errors.New("Dragon altar does not exist")
				}

				// Resurrect
				dragon.StopEverything()
				dragon.SetProperty("alive", true)
				dragon.SetProperty("visible", true)
				dragon.SetProperty("health", dragon.GetProperty("max_health"))
				dragon.SetX(altar.X())
				dragon.SetY(altar.Y())
				dragon.SetFloor(altar.Floor())
				e.Floors()[dragon.Floor()].Insert(dragon)
				e.Mobs().Store(dragon.Id(), dragon)
				e.SendGameObjectUpdate(dragon, "add_object")
				dragon.Engine.SendResponse("dragons_info", charGameObj.(entity.ICharacterObject).GetDragonsInfo(dragon.Engine), player)
				
				return true, nil
			} else {
				e.SendSystemMessage("You are not the owner of this creature.", player)
				return false, errors.New("Player is not the owner of this creature")
			}
		} else {
			return false, errors.New("Player not found")
		}
	} else {
		return false, errors.New("Player not found")
	}
}
