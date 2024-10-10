// Path: project_root/db/redis.go

package db

import (
	"Go_lang_Microservice/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() error {
	opt, err := redis.ParseURL(config.Get("REDIS_URL"))
	if err != nil {
		return err
	}

	RedisClient = redis.NewClient(opt)
	return nil
}
