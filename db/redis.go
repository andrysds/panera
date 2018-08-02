package db

import (
	"os"

	"github.com/andrysds/clarity/errutil"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	_, err := client.Ping().Result()
	errutil.PanicIfError(err, "error on connecting to redis")
	Redis = client
}
