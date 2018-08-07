package db

import (
	"log"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	options, err := redis.ParseURL(config.RedisUrl)
	clarity.PanicIfError(err, "error on parsing redis url")

	client := redis.NewClient(options)
	_, err = client.Ping().Result()
	clarity.PanicIfError(err, "error on connecting to redis")

	Redis = client
	log.Println("* Redis initialized")
}
