package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetToken(ctx context.Context, key string) (error, string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return err, ""
	}
	return nil, value

}

func GetToken(ctx context.Context, key string, value string) (error, string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	status := client.Set(ctx, key, value, time.Duration(time.Now().Add(time.Hour*24).Unix()))
	result, err := status.Result()
	if err != nil {
		return err, ""
	}
	return nil, result
}
