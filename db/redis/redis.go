package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var DB = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

func Push(key string, value string) error {
	err := DB.Set(Ctx, key, value, 0).Err()
	return err
}

func Get(key string) (string, error) {
	v, err := DB.Get(Ctx, key).Result()
	return v, err
}
