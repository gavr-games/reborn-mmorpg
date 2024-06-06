package storage

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/redis/go-redis/v9"
)

const (
	ChanelCapacity = 1000
)

// declaration defined type
type StorageClient struct {
	redisClient      *redis.Client
	Updates          chan entity.IGameObject
	GameAreasUpdates chan *entity.GameArea
	Deletes          chan string
}

var instance *StorageClient = nil
var ctx = context.Background()
var once sync.Once

func GetClient() *StorageClient {
	if instance != nil {
		return instance
	}

	once.Do(func() {
		opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
		rdb := redis.NewClient(opt)
		SetClient(rdb)
	})
	return instance
}

func SetClient(rdb *redis.Client) {
	instance = &StorageClient{
		redisClient:      rdb,
		Updates:          make(chan entity.IGameObject, ChanelCapacity),
		GameAreasUpdates: make(chan *entity.GameArea, ChanelCapacity),
		Deletes:          make(chan string, ChanelCapacity),
	}
}

func (sc *StorageClient) SaveGameObject(obj entity.IGameObject) {
	message, err := json.Marshal(obj.(*entity.GameObject))
	if err != nil {
		panic(err)
	}

	setErr := sc.redisClient.Set(ctx, obj.Id(), message, 0).Err()
	if setErr != nil {
		panic(setErr)
	}
}

func (sc *StorageClient) RemoveGameObject(objId string) {
	if err := sc.redisClient.Del(ctx, objId).Err(); err != nil {
		panic(err)
	}
}

func (sc *StorageClient) GetGameObject(id string) entity.IGameObject {
	val, redisErr := sc.redisClient.Get(ctx, id).Result()
	var obj *entity.GameObject
	if redisErr != nil {
		panic(redisErr)
	}
	err := json.Unmarshal([]byte(val), &obj)
	if err != nil {
		panic(err)
	}
	return obj
}

func (sc *StorageClient) ReadAllGameObjects(process func(entity.IGameObject)) int {
	i := 0
	iter := sc.redisClient.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		obj := sc.GetGameObject(iter.Val())
		if obj.Type() != "" { // exclude GameArea and other objects stored in Redis
			process(obj)
			i++
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return i
}

func (sc *StorageClient) GetGameArea(id string) *entity.GameArea {
	val, redisErr := sc.redisClient.Get(ctx, id).Result()
	var ga *entity.GameArea
	if redisErr != nil {
		panic(redisErr)
	}
	err := json.Unmarshal([]byte(val), &ga)
	if err != nil {
		panic(err)
	}
	return ga
}

func (sc *StorageClient) SaveGameArea(obj *entity.GameArea) {
	message, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	setErr := sc.redisClient.Set(ctx, obj.Id(), message, 0).Err()
	if setErr != nil {
		panic(setErr)
	}
}

func (sc *StorageClient) Run() {
	for {
		select {
		case obj := <- sc.Updates:
			sc.SaveGameObject(obj)
		case ga := <- sc.GameAreasUpdates:
			sc.SaveGameArea(ga)
		case objId := <- sc.Deletes:
			sc.RemoveGameObject(objId)
		}
	}
}
