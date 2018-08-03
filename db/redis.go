package db

import (
	"os"

	"github.com/andrysds/clarity"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	options, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	clarity.PanicIfError(err, "error on parsing redis url")
	client := redis.NewClient(options)
	_, err = client.Ping().Result()
	clarity.PanicIfError(err, "error on connecting to redis")
	Redis = client
}
