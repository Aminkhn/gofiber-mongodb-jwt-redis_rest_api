package database

import (
	"log"

	"github.com/aminkhn/mongo-rest-api/config"
	"github.com/go-redis/redis"
)

type RedisDbInstance struct {
	Db *redis.Client
}

var (
	RedisDb RedisDbInstance
)

// ping pong error check
func RedisConnectDb(config *config.Configuration) {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
		DB:   15,
	})
	// checking Redis Connection
	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("Somethig is wrong with Redis!")
	} else {
		log.Printf("connection to Redis established '%s'", pong)
	}
	// redis instance
	RedisDb = RedisDbInstance{Db: rdb}
}
