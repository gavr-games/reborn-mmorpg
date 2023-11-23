package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func UpdateAll(e entity.IEngine, tickDelta int64) {
	for _, gameObj := range e.GameObjects() {
		if gameObj != nil && gameObj.CurrentAction() != nil {
			gameObj.CurrentAction().TimeLeft = gameObj.CurrentAction().TimeLeft - float64(tickDelta)
			if (gameObj.CurrentAction().TimeLeft <= 0.0) {
				Finish(e, gameObj)
			}
		}
	}
}

