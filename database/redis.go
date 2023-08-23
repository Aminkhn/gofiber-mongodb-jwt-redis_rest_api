package database

import (
	"log"

	"github.com/aminkhn/mongo-rest-api/config"
	"github.com/go-redis/redis"
)

type RedisDbInstance struct {
	Db *redis.Client
}

var RedisDb RedisDbInstance

func RedisConnectDb(config *config.Configuration) {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	log.Println("connection to Redis established")
	// redis instance
	RedisDb = RedisDbInstance{Db: rdb}
}
