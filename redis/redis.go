package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "37.152.183.146:6379",
	Password: "",
	DB:       0,
})

// Set value to redis
func Set(key string, value []byte) {
	rdb.Set(key, value, time.Duration(30)*time.Minute).Err()
}

// Get value from redis
func Get(key string) (string, error) {
	val, error := rdb.Get(key).Result()
	return val, error
}
