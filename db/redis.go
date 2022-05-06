package db

import (
	"go_blog/common"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRdb() {
	configset := common.GetConfig()
	redisHost := configset.RedisDB.Host
	redisPort := configset.RedisDB.Port
	redisPassword := configset.RedisDB.Password

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})
	Rdb = rdb
}
