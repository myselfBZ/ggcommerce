package redis

import (
	"context"
    "log"
	"github.com/go-redis/redis/v8"
)


var ctx = context.Background()

var Client *redis.Client

func InitRedis() {
    Client = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := Client.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }
}

