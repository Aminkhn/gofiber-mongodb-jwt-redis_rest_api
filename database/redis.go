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

// Establishing Redis Connection
func RedisConnectDb(config *config.Configuration) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.RedisUrl,
		Password:     "", // no password set
		DB:           0,  // use default DB
		ReadTimeout:  -1,
		WriteTimeout: -1,
	})
	// ping pong error check
	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("Somethig is wrong with Redis!")
	} else {
		log.Printf("connection to Redis established '%s'", pong)
	}
	// redis instance
	RedisDb = RedisDbInstance{Db: rdb}
}
