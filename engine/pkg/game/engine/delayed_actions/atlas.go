package delayed_actions
import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/trees"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/rocks"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/plants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/hatcheries"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/craft"
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
		"CutCactus": map[string]interface{}{
			"func": plants.CutCactus,
			"duration": 3000.0, // ms
		},
		"Craft": map[string]interface{}{
			"func": craft.Craft,
			"duration": 0.0, // will be taken from craft atlas
		},
		"HatchFireDragon": map[string]interface{}{
			"func": hatcheries.HatchFireDragon,
			"duration": 60000.0,
		},
	}

	return delayedActionsAtlas
}
