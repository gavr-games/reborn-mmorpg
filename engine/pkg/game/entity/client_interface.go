package entity

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

type IClient interface {
	GetSendChannel() chan []byte
	GetCharacter() *utils.Character
}
