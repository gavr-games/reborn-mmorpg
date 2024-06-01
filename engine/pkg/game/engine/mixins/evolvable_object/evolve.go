package evolvable_object

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
)

func (obj *EvolvableObject) Evolve(e entity.IEngine, player *entity.Player) (bool, error) {
	var (
		charGameObj entity.IGameObject
		charOk bool
	)
	gameObj := obj.gameObj
	if charGameObj, charOk = e.GameObjects().Load(player.CharacterGameObjectId); !charOk {
		return false, errors.New("Character not found")
	}

	// Check character is close enough
	if !gameObj.IsCloseTo(charGameObj) {
		e.SendSystemMessage("You need to be closer.", player)
		return false, errors.New("Character needs to be closer")
	}

	// Check enough food
	foodToEvolve := gameObj.GetProperty("food_to_evolve")
	if foodToEvolve != nil && foodToEvolve != 0.0 {
		e.SendSystemMessage(fmt.Sprintf("The %s needs %d more food to evolve.", gameObj.Kind(), int(foodToEvolve.(float64))), player)
		return false, errors.New("Object needs more food to evolve")
	}

	// Check alive
	if alive := gameObj.GetProperty("alive"); alive != nil && !alive.(bool) {
		e.SendSystemMessage(fmt.Sprintf("The %s is dead, ressurect it first.", gameObj.Kind()), player)
		return false, errors.New("The evolvable object is dead")
	}

	// Check claim access
	if !claims.CheckAccess(e, charGameObj, gameObj) {
		e.SendSystemMessage("You don't have an access to this claim.", player)
		return false, errors.New("No access to evolvable object")
	}

	// Evolve
	if evolveTo := gameObj.GetProperty("evolve_to"); evolveTo != nil {
		ownerId := gameObj.GetProperty("owner_id")
		evolveGameObj := e.CreateGameObject(
			evolveTo.(string), gameObj.X(), gameObj.Y(), 0.0, gameObj.GameAreaId(), 
			map[string]interface{}{
				"owner_id": ownerId,
				"level": gameObj.GetProperty("level"),
				"experience": gameObj.GetProperty("experience"),
			})
		// Update dragons list for character
		if strings.Contains(gameObj.Kind(), "_dragon") && ownerId != nil && ownerId.(string) == charGameObj.Id() {
			dragonIds := charGameObj.GetProperty("dragons_ids").([]interface{})
			index := slices.IndexFunc(dragonIds, func(id interface{}) bool { return id != nil && id.(string) == gameObj.Id() })
			dragonIds[index] = evolveGameObj.Id()
			charGameObj.SetProperty("dragons_ids", dragonIds) 
			storage.GetClient().Updates <- charGameObj.Clone()
		}
		e.SendResponseToVisionAreas(evolveGameObj, "add_object", map[string]interface{}{
			"object": evolveGameObj.Clone(),
		})
		e.RemoveGameObject(gameObj)

		return true, nil
	}

	return false, errors.New("Object cannot evolve")
}