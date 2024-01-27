package game_test

import (
	"os"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/redis/go-redis/v9"
)

func TestMain(m *testing.M) {
	server := miniredis.RunT(&testing.T{})
	rdb := redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})
	storage.SetClient(rdb)
	os.Exit(m.Run())
}
