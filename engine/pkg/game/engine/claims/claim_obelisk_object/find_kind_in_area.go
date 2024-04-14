package claim_obelisk_object

import (
	"errors"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (claimObelisk *ClaimObeliskObject) FindKindInArea(e entity.IEngine, kind string) (entity.IGameObject, error) {
	// get claim area
	claimAreaId := claimObelisk.GetProperty("claim_area_id")
	if claimAreaId == nil {
		return nil, errors.New("Claim area does not exist")
	}
	claimAreaObj, areaOk := e.GameObjects().Load(claimAreaId.(string))
	if !areaOk {
		return nil, errors.New("Claim area does not exist")
	}

	gameArea, gaOk := e.GameAreas().Load(claimAreaObj.GameAreaId())
	if !gaOk {
		return nil, errors.New("GameArea does not exist")
	}
	possibleObjects := gameArea.RetrieveIntersections(utils.Bounds{
		X:      claimAreaObj.X(),
		Y:      claimAreaObj.Y(),
		Width:  claimAreaObj.Width(),
		Height: claimAreaObj.Height(),
	})

	for _, val := range possibleObjects {
		gameObj := val.(entity.IGameObject)
		if gameObj.Kind() == kind {
			return gameObj, nil
		}
	}

	return nil, errors.New("No object found on claim area")
}
