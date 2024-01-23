package delayed_actions
import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/trees"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/rocks"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/plants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/hatcheries"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/characters"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/claims"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/shovels"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/surfaces"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
)


func GetDelayedActionsAtlas() map[string]map[string]interface{} {
	delayedActionsAtlas:= map[string]map[string]interface{}{
		"Chip": map[string]interface{}{
			"func": rocks.Chip,
			"duration": 3000.0, // ms
		},
		"Chop": map[string]interface{}{
			"func": trees.Chop,
			"duration": 3000.0, // ms
		},
		"ClaimTeleport": map[string]interface{}{
			"func": characters.ClaimTeleport,
			"duration": 10000.0,
		},
		"Craft": map[string]interface{}{
			"func": craft.Craft,
			"duration": 0.0, // will be taken from craft atlas
		},
		"CutPlant": map[string]interface{}{
			"func": plants.Cut,
			"duration": 3000.0, // ms
		},
		"Dig": map[string]interface{}{
			"func": shovels.Dig,
			"duration": 500.0,
		},
		"ExpireClaim": map[string]interface{}{
			"func": claims.Expire,
			"duration": constants.ClaimRentDuration,
		},
		"GrowGrass": map[string]interface{}{
			"func": surfaces.GrowGrass,
			"duration": 60000.0,
		},
		"GrowPlant": map[string]interface{}{
			"func": plants.Grow,
			"duration": 60000.0,
		},
		"HarvestPlant": map[string]interface{}{
			"func": plants.Harvest,
			"duration": 500.0, // ms
		},
		"Hatch": map[string]interface{}{
			"func": hatcheries.Hatch,
			"duration": 60000.0,
		},
		"InitClaim": map[string]interface{}{
			"func": claims.Init,
			"duration": 1.0,
		},
		"TownTeleport": map[string]interface{}{
			"func": characters.TownTeleport,
			"duration": 10000.0,
		},
	}

	return delayedActionsAtlas
}
