package db

import (
	"os"

	"github.com/andrysds/clarity/errutil"
	"github.com/gomodule/redigo/redis"
)

var Redis *redis.Conn

func InitRedis() {
	redisConn, err := redis.DialURL(os.Getenv("REDIS_URL"))
	errutil.PanicIfError(err, "error on connecting to redis")
	Redis = &redisConn
}
