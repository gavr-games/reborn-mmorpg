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
		"Chop": map[string]interface{}{
			"func": trees.Chop,
			"duration": 3000.0, // ms
		},
		"Chip": map[string]interface{}{
			"func": rocks.Chip,
			"duration": 3000.0, // ms
		},
		"CutPlant": map[string]interface{}{
			"func": plants.Cut,
			"duration": 3000.0, // ms
		},
		"Craft": map[string]interface{}{
			"func": craft.Craft,
			"duration": 0.0, // will be taken from craft atlas
		},
		"Hatch": map[string]interface{}{
			"func": hatcheries.Hatch,
			"duration": 60000.0,
		},
		"TownTeleport": map[string]interface{}{
			"func": characters.TownTeleport,
			"duration": 10000.0,
		},
		"ClaimTeleport": map[string]interface{}{
			"func": characters.ClaimTeleport,
			"duration": 10000.0,
		},
		"InitClaim": map[string]interface{}{
			"func": claims.Init,
			"duration": 1.0,
		},
		"ExpireClaim": map[string]interface{}{
			"func": claims.Expire,
			"duration": constants.ClaimRentDuration,
		},
		"Dig": map[string]interface{}{
			"func": shovels.Dig,
			"duration": 500.0,
		},
		"GrowGrass": map[string]interface{}{
			"func": surfaces.GrowGrass,
			"duration": 60000.0,
		},
	}

	return delayedActionsAtlas
}
