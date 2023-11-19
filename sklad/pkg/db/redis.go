package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Cache_Set(key, value string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func Cache_Get(key string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println(key, " does not exist")
		return "", err
	} else if err != nil {
		panic(err)
	} else {
		return value, err
	}
}
