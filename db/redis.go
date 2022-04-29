package db

import (
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "password123", // no password set
		DB:       0,             // use default DB
	})
	Rdb = rdb
}
