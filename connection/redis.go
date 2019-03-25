package connection

import (
	"log"
	"os"

	"github.com/andrysds/clarity"
	"github.com/go-redis/redis"
)

// Redis represents Redis connection
var Redis *redis.Client

// InitRedis initializes Redis connection
func InitRedis() {
	options, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	clarity.PanicIfError(err, "error on parsing redis url")

	client := redis.NewClient(options)
	_, err = client.Ping().Result()
	clarity.PanicIfError(err, "error on connecting to redis")

	Redis = client
	log.Println("* Redis initialized")
}
