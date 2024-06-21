package generate_world

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func GenerateWorld(e entity.IEngine) {
	GenerateSurface(e)
	GenerateTown(e)
}
