package config

import "github.com/go-redis/redis/v8"

var Redis *redis.Client

func CreateRedisClient() {
	opt, err := redis.ParseURL("redis://172.30.160.1:6379/0")
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(opt)
	Redis = redis
}
