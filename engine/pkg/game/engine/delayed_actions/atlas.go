package delayed_actions
import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/trees"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/rocks"
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
	}

	return delayedActionsAtlas
}
