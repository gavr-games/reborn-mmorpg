package vision_area_updater

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	ChanelCapacity = 1000
)

type VisionAreaUpdate struct {
	GameObj entity.IGameObject
	ResponseType string
	ResponseData map[string]interface{}
}

type VisionAreaUpdater struct {
	e       entity.IEngine
	Updates chan *VisionAreaUpdate
}

var instance *VisionAreaUpdater = nil
var ctx = context.Background()
var once sync.Once

func GetUpdater(e entity.IEngine) *VisionAreaUpdater {
	once.Do(func() {
		instance = &VisionAreaUpdater{
			e:       e,
			Updates: make(chan *VisionAreaUpdate, ChanelCapacity),
		}
	})
	return instance
}

func (vau *VisionAreaUpdater) updatesWorker(updatesChan <-chan *VisionAreaUpdate) {
	for visionAreaUpdate := range updatesChan {
		e := vau.e
		gameObj := visionAreaUpdate.GameObj
		intersectingObjects := e.Floors()[gameObj.Floor()].RetrieveIntersections(utils.Bounds{
			X:      gameObj.X(),
			Y:      gameObj.Y(),
			Width:  gameObj.Width(),
			Height: gameObj.Height(),
		})
		resp := entity.EngineResponse{
			ResponseType: visionAreaUpdate.ResponseType,
			ResponseData: visionAreaUpdate.ResponseData,
		}
		message, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, obj := range intersectingObjects {
			if obj.(entity.IGameObject).Type() == "player" && obj.(entity.IGameObject).Kind() == "player_vision_area" {
				playerId := obj.(entity.IGameObject).GetProperty("player_id").(int)
				if player, ok := e.Players().Load(playerId); ok {
					if player.Client != nil {
						select {
						case player.Client.GetSendChannel() <- message:
						default:
							engine.UnregisterClient(e, player.Client)
						}
					}
				}
			}
		}
	}
}

func (vau *VisionAreaUpdater) Run() {
	go vau.updatesWorker(vau.Updates)
}
